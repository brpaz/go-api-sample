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
		log.Info("loading .env")
		if err := godotenv.Load(); err != nil {
			log.Info("Failed to load dotenv file:" + err.Error())
		}
	}
}

func setupLogger(cfg config.Config) (*zap.Logger, error) {
	// TODO Handle output format in dev plus debug mode.
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

	if err := app.New(cfg, logger).StartServer(); err != nil {
		logger.Fatal("Failed to start application server:" + err.Error())
	}
}
