package testutil

import (
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/brpaz/go-api-sample/internal/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	http "net/http"
	"net/http/httptest"
)

func CreateEchoTestContext(req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	e := echo.New()
	e.Validator = validator.NewRequestValidator()
	e.HTTPErrorHandler = appHttp.NewErrorHandler(zap.NewNop()).Handle
	c := e.NewContext(req, rec)

	return c
}
