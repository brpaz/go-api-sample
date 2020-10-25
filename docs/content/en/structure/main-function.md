---
title: 'Main function'
position: 1
category: 'Application Structure'
fullscreen: false
---

The main function located in `cmd/server/main.go` is responsible for loading the application configuration, setup the logger and other global dependencies and boot the application.

The main function is tiny and should contain no application logic:

```go
func main() {
    
    // loads .env files
	dotenv()

	// Load config into the application
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load application configuration", err)
	}

	// Setup the application logger
	logger, err := logging.BuildLogger(cfg)

	if err != nil {
		log.Fatalf("Failed to configure application logger", err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	appInstance := app.New(cfg, logger)

	go func() {
		if err := appInstance.Start(); err != nil {
			logger.Fatal("Failed to start application server:" + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := appInstance.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal(err)
	}
```

The `appInstance := app.New(cfg, logger)` instantiates our application. This code is placed on the `internal/app`package.





