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

type Server struct {
	e           *echo.Echo
	logger      *zap.Logger
	config      config.Config
	diContainer di.Container
}

// NewServer Creates a new instance of the application server
func NewServer(cfg config.Config, logger *zap.Logger, dic di.Container) *Server {
	srv := Server{
		e:      echo.New(),
		logger: logger,
		config: cfg,
		diContainer:  dic,
	}

	srv.configure()
	srv.registerMiddlewares()
	srv.registerRoutes()

	return &srv
}

func (srv *Server) configure() {
	srv.e.HideBanner = true
	srv.e.HidePort = true
	srv.e.Validator = appHttp.NewRequestValidator(validator.New())
	srv.e.HTTPErrorHandler = appHttp.NewErrorHandler(srv.logger).Handle
}

func (srv *Server) registerMiddlewares() {
	srv.e.Use(appMiddleware.ZapLogger(srv.logger))
	srv.e.Use(middleware.RequestID())
	srv.e.Use(middleware.Recover())
	srv.e.Use(middleware.Gzip())

	if srv.config.MetricsEnabled {
		p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/_heatlh") {
				return true
			}
			return false
		})
		p.Use(srv.e)
	}
}

func (srv *Server) registerRoutes() {
	srv.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	hh := srv.diContainer.Get(di.ServiceHealcheckHandler).(*healthcheck.Handler)
	srv.e.GET("/_health", hh.Handle)

	todosCreateHandler := srv.diContainer.Get(di.ServiceTodoCreateHandler).(*todo.CreateHandler)
	srv.e.POST("/todos", todosCreateHandler.Handle)
}

func (srv *Server) Start() error {
	port := fmt.Sprintf(":%d", srv.config.Port)

	srv.logger.Info(fmt.Sprintf("Starting application on port %s", port))
	return srv.e.Start(port)
}

func (srv *Server) Shutdown(ctx context.Context) error {
	srv.logger.Info("stopping server")
	return srv.e.Shutdown(ctx)
}
