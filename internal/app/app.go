// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/kjj49/test-go-go-ahead/config"
	"github.com/kjj49/test-go-go-ahead/internal/controller/http"
	"github.com/kjj49/test-go-go-ahead/internal/usecase"
	"github.com/kjj49/test-go-go-ahead/internal/usecase/repository"
	"github.com/kjj49/test-go-go-ahead/internal/usecase/webapi"
	"github.com/kjj49/test-go-go-ahead/pkg/httpserver"
	"github.com/kjj49/test-go-go-ahead/pkg/logger"
	"github.com/kjj49/test-go-go-ahead/pkg/postgres"
	"github.com/robfig/cron/v3"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	currencyUseCase := usecase.New(
		repository.New(pg),
		webapi.New(),
	)

	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, l, currencyUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Cron
	c := cron.New()
	// Start the GetAllCurrency feature at 10:00 UTC+ daily
	c.AddFunc("0 10 * * *", func() {
		currencyUseCase.GetAllCurrency(context.Background())
	})
	c.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
