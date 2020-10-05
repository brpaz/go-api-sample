package app

import (
	"context"
	"fmt"
	"github.com/brpaz/echozap"
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	config config.Config
	logger *zap.Logger
	di     di.Container
	server *echo.Echo
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {

	app := &App{
		config: config,
		logger: logger,
	}

	app.bootstrap()
	return app
}

// bootstraps the application. This function is responsible for building the DI container and to configure the Echo web server
func (app *App) bootstrap() {
	app.buildContainer()

	e := echo.New()
	e.HideBanner = true
	e.Debug = app.config.Debug

	e.Use(echozap.ZapLogger(app.logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Routes
	app.registerRoutes(e)

	app.server = e
}

// StartServer This function is responsible to start the application server
func (app *App) StartServer() error {
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

func (app *App) buildContainer() {
	// TODO handle error
	// TODO this logic can be moved outsite of app when we start having many services definitions
	builder, _ := di.NewBuilder()

	app.di = builder.Build()
}
