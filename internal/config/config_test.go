// +build unit

package config_test

import (
	"github.com/brpaz/go-api-sample/internal/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	_ = os.Setenv("APP_PORT", "1000")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASSWORD", "123456")
	_ = os.Unsetenv("APP_DEBUG")

	defer func() {
		_ = os.Unsetenv("APP_PORT")
		_ = os.Unsetenv("DB_HOST")
		_ = os.Unsetenv("DB_PORT")
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
	}()

	cfg, err := config.Load()

	assert.Nil(t, err)
	assert.Equal(t, 1000, cfg.Port)
	assert.False(t, cfg.Debug)
}

func TestLoad_WithError(t *testing.T) {

	_ = os.Setenv("APP_DEBUG", "invalid-value")
	defer func() {
		_ = os.Unsetenv("APP_DEBUG")
	}()

	_, err := config.Load()

	assert.NotNil(t, err)
}
