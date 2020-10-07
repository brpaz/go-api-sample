package app

import (
	"context"
	"fmt"
	"github.com/brpaz/echozap"
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/internal/config"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/internal/http/healthcheck"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	config config.Config
	logger *zap.Logger
	dic    di.Container
	server *echo.Echo
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {
	return &App{
		config: config,
		logger: logger,
		dic:    di.BuildContainer(config),
	}
}

// StartServer This function is responsible to start the application server
func (app *App) StartServer() error {

	e := echo.New()
	e.HideBanner = true
	e.Debug = app.config.Debug
	e.Validator = appHttp.NewRequestValidator(validator.New())
	e.HTTPErrorHandler = appHttp.ErrorHandler

	e.Use(echozap.ZapLogger(app.logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	hh := app.dic.Get(di.ServiceHealcheckHandler).(healthcheck.Handler)
	e.GET("/_health", hh.Handle)

	/*todosCreateHandler :=  app.dic.Get(di.ServiceTodoCreateUseCase).(*todo.CreateHandler)
	e.POST("/todos", todosCreateHandler.Handle)*/

	app.server = e


	port := fmt.Sprintf(":%d", app.config.Port)

	return app.server.Start(port)
}

func (app *App) Shutdown() error {
	if app.server != nil {
		return app.server.Shutdown(context.Background())
	}

	return nil
}

// ServeHTTP Implements the ServeHTTP interface. This function is mostly useful for testing because it allows
// integration with the httptest server, to easy start the application in a testing scenario.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.server.ServeHTTP(w, r)
}
