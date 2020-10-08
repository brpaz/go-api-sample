// +build unit

package http_test

import (
	"encoding/json"
	"errors"
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getMockErrorHandler() *appHttp.ErrorHandler {
	return appHttp.NewErrorHandler(zap.NewNop())
}

func TestErrorHandler_Handle_GenericError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	getMockErrorHandler().Handle(errors.New("some-error"), ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code)
	assert.Equal(t, "Internal Error", respBody.Message)
}

func TestErrorHandler_Handler_ApplicationError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	getMockErrorHandler().Handle(appErrors.
		NewApplicationError(
			appErrors.ErrCodeInternalError,
			"error"),
		ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code)
}

func TestErrorHandler_Handle_echoHTTPError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	getMockErrorHandler().Handle(echo.NewHTTPError(http.StatusServiceUnavailable, "message"), ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
	assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code)
	assert.Equal(t, "message", respBody.Message)
}

func TestErrorHandler_Handle_ValidatorErrors(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	type mystruct struct {
		Value string `json:"value" validate:"required"`
	}

	validatorInstance := appHttp.NewRequestValidator(validator.New())
	err := validatorInstance.Validate(&mystruct{})

	getMockErrorHandler().Handle(err, ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	assert.Len(t, respBody.Fields, 1)
	assert.Equal(t, respBody.Code, appErrors.ErrCodeValidationFailed)
}
