package main

import (
	"fmt"
	"os"

	"github.com/brpaz/echozap"
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/brpaz/go-api-sample/internal/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Logger The application logger
var Logger *zap.Logger

func main() {

	// Loads environment variables from ".env file" in dev mode.
	if os.Getenv("APP_ENV") == "dev" {
		_ = godotenv.Load()
	}

	// Load config into the application
	if err := config.Load(); err != nil {
		panic(err)
	}

	// Setups the application logger
	setupLogger()

	// Configures and starts the application
	startServer()
}

func setupLogger() {

	if config.Get().Env == "dev" {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}
}

func startServer() {
	e := echo.New()
	e.HideBanner = true
	e.Debug = config.Get().Debug

	e.Use(echozap.ZapLogger(Logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Routes
	e.GET("/hello", handlers.Hello)
	e.GET("/_health", handlers.Health)

	port := fmt.Sprintf(":%d", config.Get().Port)
	e.Logger.Fatal(e.Start(port))
}
