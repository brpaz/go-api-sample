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
	_ = os.Unsetenv("APP_DEBUG")
	defer func() {
		_ = os.Unsetenv("APP_PORT")
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
