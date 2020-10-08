// +build unit

package app_test

import (
	"github.com/brpaz/go-api-sample/internal/app"
	"github.com/brpaz/go-api-sample/test/testutil"
	"go.uber.org/zap"
	"testing"
)

func TestApp_Creeate(t *testing.T) {
	cfg := testutil.GetMockConfig()
	logger := zap.NewNop()

	app.New(cfg, logger)
}
