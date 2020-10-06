package http

import (
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Code    appErrors.ErrorCode `json:"code"`
	Message string              `json:"message"`
	Fields  []ErrorField        `json:"fields,omitempty" `
}

type ErrorField struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

// ErrorHandler custom HTTP error handler
func ErrorHandler(err error, c echo.Context) {
	httpCode := http.StatusInternalServerError
	errCode := appErrors.ErrCodeInternalError
	errMessage := "internal error"

	if he, ok := err.(*echo.HTTPError); ok {
		httpCode = he.Code
		errMessage = he.Message.(string)
	} else if ae, ok := err.(*appErrors.ApplicationError); ok {
		errMessage = ae.Message

	} else if ve, ok := err.(validator.ValidationErrors); ok {
		errCode := appErrors.ErrCodeValidationFailed
		errMessage = ve.Error()

		// TODO handle field messages
		httpCode = http.StatusUnprocessableEntity
	}

	response := ErrorResponse{
		Code: errCode,
		Message: errMessage,
	}

	c.JSON(httpCode, response)
}
