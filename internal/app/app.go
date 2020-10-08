package app

import (
	"context"
	"fmt"
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/internal/config"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/internal/http/healthcheck"
	appMiddleware "github.com/brpaz/go-api-sample/internal/http/middleware"
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type App struct {
	config config.Config
	logger *zap.Logger
	dic    di.Container
	server *echo.Echo
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {
	dic := di.BuildContainer(config, logger)
	return &App{
		config: config,
		logger: logger,
		dic:    dic,
		//server: http.NewServer(config, dic),
	}
}

// Start This function is responsible to start the application server
func (app *App) Start() error {

	e := echo.New()

	// we will treat logs by ourselves
	e.HideBanner = true
	e.HidePort = true

	e.Debug = app.config.Debug
	e.Validator = appHttp.NewRequestValidator(validator.New())
	e.HTTPErrorHandler = appHttp.NewErrorHandler(app.logger).Handle
	e.Use(appMiddleware.ZapLogger(app.logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
		if strings.HasPrefix(c.Path(), "/_heatlh") {
			return true
		}
		return false
	})
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	hh := app.dic.Get(di.ServiceHealcheckHandler).(*healthcheck.Handler)
	e.GET("/_health", hh.Handle)

	todosCreateHandler := app.dic.Get(di.ServiceTodoCreateHandler).(*todo.CreateHandler)
	e.POST("/todos", todosCreateHandler.Handle)

	app.server = e

	port := fmt.Sprintf(":%d", app.config.Port)

	app.logger.Info(fmt.Sprintf("Starting application on port %d", app.config.Port))
	return app.server.Start(port)
}

func (app *App) Shutdown(ctx context.Context) error {
	if app.server != nil {
		return app.server.Shutdown(ctx)
	}

	return nil
}
