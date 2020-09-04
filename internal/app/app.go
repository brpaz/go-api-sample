package app

import (
	"fmt"
	"github.com/brpaz/echozap"
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

type App struct {
	config config.Config
	logger *zap.Logger
	di     di.Container
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {
	app := &App{
		config: config,
		logger: logger,
	}

	app.buildContainer()

	return app
}

// Boot This function is responsible to start the Application server and configure all the routes and middlewares
func (app *App) Boot() error {
	e := echo.New()
	e.HideBanner = true
	e.Debug = app.config.Debug

	e.Use(echozap.ZapLogger(app.logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Routes
	app.registerRoutes(e)

	port := fmt.Sprintf(":%d", app.config.Port)

	return e.Start(port)
}

func (app *App) buildContainer() {
	// TODO handle error
	// TODO this logic can be moved outsite of app when we start having many services definitions
	builder, _ := di.NewBuilder()

	app.di = builder.Build()
}
