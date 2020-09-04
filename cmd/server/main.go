package main

import (
	"github.com/brpaz/go-api-sample/internal/app"
	"github.com/labstack/gommon/log"
	"os"

	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Loads environment variables from ".env file" in dev mode.
func dotenv() {
	if os.Getenv("APP_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Errorf("Failed to load dotenv file", err)
		}
	}
}

func setupLogger(cfg config.Config) (*zap.Logger, error) {
	// Handle output format in dev plus debug mode.
	if cfg.Env == config.EnvDev {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}

// Main function
func main() {

	dotenv()

	// Load config into the application
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load application configuration", err)
	}

	// Setups the application logger
	logger, err := setupLogger(cfg)

	if err != nil {
		log.Fatalf("Failed to configure application logger", err)
	}

	appInstance := app.New(cfg, logger)

	if err := appInstance.Boot(); err != nil {
		logger.Fatal("Failed to start application server:" + err.Error())
	}
}
