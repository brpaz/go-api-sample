---
title: 'Logging'
position: 4
category: 'Application Structure'
fullscreen: false
---

This application uses [Zap logger](https://github.com/uber-go/zap) which is one of the most performant Golang loggers.

In the `ìnternal/logging/logging.go` you can find the `BuildLogger` function, which is responsible to build a new instance of the Zap logger, based on the application configuration.

```go
// BuildLogger creates a new instance of ZAP Logger based on confgi
func BuildLogger(cfg config.Config) (*zap.Logger, error) {
	var level zapcore.Level

	err := (&level).UnmarshalText([]byte(cfg.LogLevel))

	if err != nil {
		log.Warnf("Unrecognized Log level '%s'. falling back to INFO", cfg.LogLevel)
		level = zapcore.InfoLevel
	}

	conf := zap.NewProductionConfig()

	if cfg.Env == config.EnvDev {
		conf = zap.NewDevelopmentConfig()
	}

	conf.Level.SetLevel(level)

	return conf.Build()
}
```

The logger should be constructed on the main function then passing as dependency to the application and subsequent services.

Why not a global logger? Personal preference. In general global state is bad, and I  like to have explicit dependencies in my code. It is also easier to mock in tests.

I recommend reading [Go, without package scoped variables](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables) by Dave Cheney.
