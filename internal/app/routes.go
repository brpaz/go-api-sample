// routes.go
// Register all application routes in this file.

package app

import (
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/internal/http/healthcheck"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Registers Application rotues
func (app *App) registerRoutes() {
	app.server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	hh := app.dic.Get(di.ServiceHealcheckHandler).(*healthcheck.Handler)
	app.server.GET("/_health", hh.Handle)

	todosCreateHandler := app.dic.Get(di.ServiceTodoCreateHandler).(*todo.CreateHandler)
	app.server.POST("/todos", todosCreateHandler.Handle)

	todosListHandler := app.dic.Get(di.ServiceTodoListHandler).(*todo.ListTodoHandler)
	app.server.GET("/todos", todosListHandler.Handle)
}
