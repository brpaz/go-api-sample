// +build unit

package http_test

import (
	"encoding/json"
	"errors"
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/test/testutil"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorHandler_GenericError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	appHttp.ErrorHandler(errors.New("some-error"), ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code )
	assert.Equal(t, "internal error", respBody.Message)
}

func TestErrorHandler_EchoHTTPError(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)

	appHttp.ErrorHandler(echo.NewHTTPError(http.StatusServiceUnavailable, "message"), ctx)

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
	assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code)
	assert.Equal(t, "message", respBody.Message)
}
