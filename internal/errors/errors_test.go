// +build unit

package errors_test

import (
	"errors"
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplicationError_withOriginalError(t *testing.T) {
	err := appErrors.NewApplicationError(appErrors.ErrCodeInternalError, "custom error message").WithOriginalError(errors.New("some error"))
	assert.Equal(t, appErrors.ErrCodeInternalError, err.Code)
	assert.Equal(t, "custom error message", err.Message)
	assert.Error(t, err.OriginalErr)
	assert.Equal(t, "ERROR: custom error message (INTERNAL_ERROR) - some error", err.Error())
}

func TestNewApplicationError_withoutOriginalError(t *testing.T) {
	err := appErrors.NewApplicationError(appErrors.ErrCodeInternalError, "custom error message")
	assert.Equal(t, appErrors.ErrCodeInternalError, err.Code)
	assert.Equal(t, "custom error message", err.Message)
	assert.Nil(t, err.OriginalErr)
	assert.Equal(t, "ERROR: custom error message (INTERNAL_ERROR)", err.Error())

}
