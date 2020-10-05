// +build unit

package healthcheck_test

import (
	"encoding/json"
	"github.com/brpaz/go-api-sample/internal/healthcheck"
	healthecklib "github.com/brpaz/go-healthcheck"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/_health", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	// Assertions
	if assert.NoError(t, healthcheck.HealthHandler(c)) {
		assert.Equal(t, http.StatusOK, resp.Code)

		var response healthecklib.Health
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			assert.Fail(t, "Error decoding response body", err)
		}

		assert.Equal(t, healthecklib.Pass, response.Status)
	}
}
