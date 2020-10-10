package http

import (
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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

type ErrorHandler struct {
	logger *zap.Logger
}

func NewErrorHandler(logger *zap.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

// ErrorHandler custom HTTP error handler
func (h *ErrorHandler) Handle(err error, c echo.Context) {
	httpCode := http.StatusInternalServerError
	errCode := appErrors.ErrCodeInternalError
	errMessage := appErrors.ErrMessageInternalError
	fieldErrors := make([]FieldError, 0)

	if he, ok := err.(*echo.HTTPError); ok {
		httpCode = he.Code
		errMessage = he.Message.(string)
	} else if ae, ok := err.(*appErrors.ApplicationError); ok {
		errMessage = ae.Message
		h.logger.Error("Http Error", zap.String("message", errMessage), zap.Error(err))
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		httpCode = http.StatusUnprocessableEntity
		errCode = appErrors.ErrCodeValidationFailed
		errMessage = appErrors.ErrMessageValidationFailed

		validatorInstance := c.Echo().Validator.(*RequestValidator)
		fieldErrors = mapValidationErrors(ve, validatorInstance)

		h.logger.Warn("Validation Error", zap.Error(err))
	}

	response := ErrorResponse{
		Code:    errCode,
		Message: errMessage,
		Fields:  fieldErrors,
	}

	_ = c.JSON(httpCode, response)
}

func mapValidationErrors(ve validator.ValidationErrors, validator *RequestValidator) []FieldError {
	mappedErrors := make([]FieldError, 0)
	for _, e := range ve {
		mappedErrors = append(mappedErrors, FieldError{
			Code:     validator.MapErrorCodeFromTag(e.Tag()),
			Message:  validator.FormatErrorMessage(e),
			Location: validator.NormalizeFieldName(e),
		})
	}

	return mappedErrors
}
