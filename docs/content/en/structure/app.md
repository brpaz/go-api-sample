---
title: 'The application'
position: 5
category: 'Application Structure'
fullscreen: false
---

The `internal/app` package is where the core application code is placed.

The `App` struct represents the Application with the main dependencies as properties.

```go
// internal/app/app.go
type App struct {
	config config.Config
	logger *zap.Logger
	dic    di.Container
	server *echo.Echo
}
```

This `Ǹew` function creates a new instance of the application.

```go
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
```

The `di.BuildContainer(config, logger)` is responsible to wire up the application services dependencies. For a more detailed explanation of how it works, check the [Depedency Injection](/structure/di) page.


## Bootstrap function

The `bootstrap` function will configure the http server, including middleware and route registration.

```go
// bootstraps the application server and configure routes and middlewares
func (app *App) bootstrap() {
	app.server.HideBanner = true
	app.server.HidePort = true
	app.server.Validator = validator.NewRequestValidator()
	app.server.HTTPErrorHandler = appHttp.NewErrorHandler(app.logger).Handle

	app.registerMiddlewares()
	app.registerRoutes()
}
```

## Start function

The `Start` function can be called to start the application server.

```go
// Start start the application server on the configured port
func (app *App) Start() error {
	port := fmt.Sprintf(":%d", app.config.Port)

	app.logger.Info(fmt.Sprintf("Starting application on port %s", port))
	return app.server.Start(port)
}
```


<alert>
No other package should include code from the app package.
The main function, creates a new app instance, which then bootstraps the http server and pass any dependencies nedeed to the downstream services.
</alert>
