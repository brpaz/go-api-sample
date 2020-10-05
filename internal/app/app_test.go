// +build unit

package app_test

import (
	"github.com/brpaz/go-api-sample/internal/app"
	"github.com/brpaz/go-api-sample/internal/config"
	"go.uber.org/zap"
	"testing"
)

func TestStartApp_Success(t *testing.T) {
	cfg := config.Config{
		Env:   config.EnvDev,
		Port:  0,
		Debug: true,
	}

	logger := zap.NewNop()

	// TODO look for better ways to really check if the server has started
	go func() {
		_ = app.New(cfg, logger).StartServer()
	}()
}
