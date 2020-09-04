package app

import (
	"github.com/brpaz/go-api-sample/internal/handlers"
	"github.com/labstack/echo/v4"
)

func (app *App) registerRoutes(e *echo.Echo) {
	e.GET("/hello", handlers.Hello)
	e.GET("/_health", handlers.Health)
}
