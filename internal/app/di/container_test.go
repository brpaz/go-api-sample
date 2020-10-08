// +build unit

package di_test

import (
	"github.com/brpaz/go-api-sample/internal/app/di"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestBuildContainer(t *testing.T) {
	cnt := di.BuildContainer(testutil.GetMockConfig(), zap.NewNop())

	assert.NotEmpty(t, cnt.Definitions())
}
