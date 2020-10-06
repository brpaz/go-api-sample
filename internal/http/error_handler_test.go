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

func TestErrorHandler_ValidationErrors(t *testing.T) {
	t.Skip("todo")
	/*var jsonBody = []byte(`{}`)

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(jsonBody))

	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := testutil.CreateEchoTestContext(req, rec)
	ctx.Echo().POST("/test", func(c echo.Context) error {
		fmt.Println("heello")
		type T struct {
			Description string `json:"description" validate:"required"`
		}
		var request T

		if err := c.Bind(&request); err != nil {
			c.Error(err)
			return err
		}


		if err := c.Validate(&request); err != nil {
			c.Error(err)
			return err
		}
		return c.String(http.StatusServiceUnavailable, "OK")
	})

	var respBody appHttp.ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&respBody); err != nil {
		assert.Fail(t, "cannot parse response", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	//assert.Equal(t, appErrors.ErrCodeInternalError, respBody.Code)
	//assert.Equal(t, "message", respBody.Message)*/
}
