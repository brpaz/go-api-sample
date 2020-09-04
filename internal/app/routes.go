package app

import (
	"github.com/brpaz/go-api-sample/internal/healthcheck"
	"github.com/labstack/echo/v4"
)

func (app *App) registerRoutes(e *echo.Echo) {
	e.GET("/hello", healthcheck.HealthHandler)
	e.GET("/_health", healthcheck.HealthHandler)
}
