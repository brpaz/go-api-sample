package app

import (
	"context"
	"errors"
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/internal/config"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	config config.Config
	logger *zap.Logger
	dic    di.Container
	server *appHttp.Server
}

// New Creates a new instance of the application
func New(config config.Config, logger *zap.Logger) *App {
	dic := di.BuildContainer(config, logger)

	return &App{
		config: config,
		logger: logger,
		dic:    dic,
		server: appHttp.NewServer(config, logger, dic),
	}
}

// Start This function is responsible to start the application server
func (app *App) Start() error {
	if app.server == nil {
		return errors.New("server is not initialized")
	}
	return app.server.Start()
}

func (app *App) Shutdown(ctx context.Context) error {
	if app.server == nil {
		return errors.New("server is not initialized")
	}
	return app.server.Shutdown(ctx)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.server.ServeHTTP(w, r)
}
