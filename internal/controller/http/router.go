// Package http implements routing paths. Each services in own file.
package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/kjj49/test-go-go-ahead/docs"

	"github.com/kjj49/test-go-go-ahead/internal/usecase"
	"github.com/kjj49/test-go-go-ahead/pkg/logger"
)

// NewRouter -.
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Currency) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/currency")
	{
		newCurrencyRoutes(h, t, l)
	}
}
