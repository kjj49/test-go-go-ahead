package main

import (
	"log"

	"github.com/kjj49/test-go-go-ahead/config"
	"github.com/kjj49/test-go-go-ahead/internal/app"
)

// @title Test Go Go Ahead
// @version 1.0
// @description API Service for Test Go Go Ahead
// @host localhost:8080
func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
