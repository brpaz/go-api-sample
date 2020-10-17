package app

import (
	"context"
	"fmt"
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/internal/config"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	appMiddleware "github.com/brpaz/go-api-sample/internal/http/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// App this struct represents the application.
type App struct {
	config config.Config
	logger *zap.Logger
	dic    di.Container
	server *echo.Echo
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {
	dic := di.BuildContainer(config, logger)

	app := App{
		config: config,
		logger: logger,
		dic:    dic,
		server: echo.New(),
	}

	app.bootstrap()

	return &app
}

// bootstraps the application server and configure routes and middlewares
func (app *App) bootstrap() {
	app.server.HideBanner = true
	app.server.HidePort = true
	app.server.Validator = appHttp.NewRequestValidator(validator.New())
	app.server.HTTPErrorHandler = appHttp.NewErrorHandler(app.logger).Handle

	app.registerMiddlewares()
	app.registerRoutes()
}

// Registers application middlewares
func (app *App) registerMiddlewares() {
	app.server.Use(appMiddleware.ZapLogger(app.logger))
	app.server.Use(middleware.RequestID())
	app.server.Use(middleware.Recover())
	app.server.Use(middleware.Gzip())

	if app.config.MetricsEnabled {
		p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/_heatlh")
		})
		p.Use(app.server)
	}
}

// Start start the application server on the configured port
func (app *App) Start() error {
	port := fmt.Sprintf(":%d", app.config.Port)

	app.logger.Info(fmt.Sprintf("Starting application on port %s", port))
	return app.server.Start(port)
}

// Shutdown stops the application server
func (app *App) Shutdown(ctx context.Context) error {
	app.logger.Info("stopping server")

	app.dic.Delete()

	return app.server.Shutdown(ctx)
}

// ServeHTTP Implements ServeHTTP interface. this allows the app to be compatible with httptest server for testing.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.server.ServeHTTP(w, r)
}
