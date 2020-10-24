package logging

import (
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
