package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
	EnvTest = "test"
)

// Config struct that holds application configuration
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

// Load Loads the application config
func Load() (Config, error) {
	var c Config

	ctx := context.Background()

	err := envconfig.Process(ctx, &c)

	return c, err
}
