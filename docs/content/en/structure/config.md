---
title: 'Configuration'
position: 3
category: 'Application Structure'
fullscreen: false
---


Configuration can come from many places (environment variables, config files) etc. 

The `Config` type represents as code all the configuration that the application needs.

```go
// app/config/config.go
type Config struct {
	Env   string `env:"APP_ENV,default=prod"`
	Port  int    `env:"APP_PORT,default=1234"`
	Debug bool   `env:"APP_DEBUG,default=false"`
	DB    struct {
		Host     string `env:"DB_HOST,required"`
		Port     uint   `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required"`
		Password string `env:"DB_PASSWORD,required"`
		Database string `env:"DB_DATABASE,required"`
		Driver   string `env:"DB_DRIVER,default=postgres"`
	}
	LogLevel       string `env:"LOG_LEVEL,default=info"`
	MetricsEnabled bool   `env:"METRICS_ENABLED,default=false"`
}
```

Why having this struct and not accessing environment variables directly for example?

* Easy to change the source of the configurations (using a file for example) without having to change all the application code.
* Can be mocked for unit tests
* It´s more explicit when you pass the config as an argument of a function.
* Easy to run validations that are the required variables are present on the application startup.
* Strongly Typed and autocomplete in the editor.

## The Envconfig library

[Envconfig](https://github.com/sethvargo/go-envconfig) is a library that populates struct field values based on environment variables or arbitrary lookup functions. 

It supports pre-setting mutations, which is useful for things like converting values to uppercase, trimming whitespace, or looking up secrets.

It also provides validations, default values and Unmarshal of complex types.

The `config.Load()` function, that will uses envconfig to populate the config struct and return.

```go
// Load Loads the application config
func Load() (Config, error) {
	var c Config

	ctx := context.Background()

	err := envconfig.Process(ctx, &c)

	return c, err
}
``
