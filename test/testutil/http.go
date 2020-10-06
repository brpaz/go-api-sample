package testutil

import (
	appHttp "github.com/brpaz/go-api-sample/internal/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	http "net/http"
	"net/http/httptest"
)

func CreateEchoTestContext(req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	e := echo.New()
	e.Validator = appHttp.NewRequestValidator(validator.New())
	e.HTTPErrorHandler = appHttp.ErrorHandler
	c := e.NewContext(req, rec)

	return c
}
