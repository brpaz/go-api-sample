package testutil

import "github.com/brpaz/go-api-sample/internal/config"

func GetMockConfig() config.Config {
	return config.Config{
		Env:   "test",
		Port:  0,
		Debug: true,
	}
}
