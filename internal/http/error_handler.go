package http

import (
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ErrorResponse struct for the ErrorResponse
type ErrorResponse struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty" `
}

// FieldError struct for the FieldError. used in validation errors
type FieldError struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

// ErrorHandler custom HTTP error handler
func ErrorHandler(err error, c echo.Context) {
	httpCode := http.StatusInternalServerError
	errCode := appErrors.ErrCodeInternalError
	errMessage := appErrors.ErrMessageInternalError
	fieldErrors := make([]FieldError, 0)

	if he, ok := err.(*echo.HTTPError); ok {
		httpCode = he.Code
		errMessage = he.Message.(string)
		c.Logger().Warn(he)
	} else if ae, ok := err.(*appErrors.ApplicationError); ok {
		errMessage = ae.Message
		c.Logger().Error(ae)
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		httpCode = http.StatusUnprocessableEntity
		errCode = appErrors.ErrCodeValidationFailed
		errMessage = appErrors.ErrMessageValidationFailed

		validatorInstance := c.Echo().Validator.(*RequestValidator)
		fieldErrors = mapValidationErrors(ve, validatorInstance)

		c.Logger().Info(ve)
	}

	response := ErrorResponse{
		Code:    errCode,
		Message: errMessage,
		Fields:  fieldErrors,
	}

	c.JSON(httpCode, response)
}

func mapValidationErrors(ve validator.ValidationErrors, validator *RequestValidator) []FieldError {
	mappedErrors := make([]FieldError, 0)
	for _, e := range ve {
		mappedErrors = append(mappedErrors, FieldError{
			Code:     validator.MapErrorCodeFromTag(e.Tag()),
			Message:  validator.MapErrorMessage(e),
			Location: validator.NormalizeFieldName(e),
		})
	}

	return mappedErrors
}
