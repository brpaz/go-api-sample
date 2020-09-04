package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

const (
	envPrefix = "APP_"
	EnvDev    = "dev"
)

// Config struct that holds application configuration
type Config struct {
	Env   string `env:"ENV,default=prod"`
	Port  int    `env:"PORT,default=1234"`
	Debug bool   `env:"DEBUG,default=false"`
}

// Load Loads the application config
func Load() (Config, error) {
	var c Config

	ctx := context.Background()

	l := envconfig.PrefixLookuper(envPrefix, envconfig.OsLookuper())
	err := envconfig.ProcessWith(ctx, &c, l)

	return c, err
}
