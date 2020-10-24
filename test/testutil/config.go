package testutil

import (
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/brpaz/go-api-sample/internal/db"
)

func GetMockConfig() config.Config {
	cfg := config.Config{
		Env:   "test",
		Port:  0,
		Debug: true,
	}

	cfg.DB.Driver = db.MockDriverName

	return cfg
}
