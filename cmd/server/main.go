package main

import (
	"context"
	"github.com/brpaz/go-api-sample/internal/app"
	"github.com/brpaz/go-api-sample/internal/logging"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"time"

	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/joho/godotenv"
)

// Loads environment variables from ".env file" in dev mode.
func dotenv() {
	if os.Getenv("APP_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Infof("Failed to load dotenv file", err)
		}
	}
}

// Main entry point of the application
func main() {

	dotenv()

	// Load config into the application
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load application configuration", err)
	}

	// Setup the application logger
	logger, err := logging.BuildLogger(cfg)

	if err != nil {
		log.Fatalf("Failed to configure application logger", err)
	}

	appInstance := app.New(cfg, logger)

	go func() {
		if err := appInstance.Start(); err != nil {
			logger.Fatal("Failed to start application server:" + err.Error())
		}
	} ()


	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := appInstance.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal(err)
	}
}
