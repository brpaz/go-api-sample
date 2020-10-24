package validator

import (
	"fmt"
	"github.com/brpaz/go-api-sample/internal/errors"
	goValidator "github.com/go-playground/validator/v10"
	"strings"
)

type RequestValidator struct {
	validator *goValidator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validator: goValidator.New(),
	}
}

func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NormalizeFieldName Normalizes a field name from a field error. Ex: lowercase.
func (cv *RequestValidator) NormalizeFieldName(fieldError goValidator.FieldError) string {
	return strings.ToLower(fieldError.StructField())
}

// MapErrorCodeFromTag converts the "tag" field from the go-validator FieldError into an application error code
func (cv *RequestValidator) MapErrorCodeFromTag(tag string) string {
	// add your validator mappings to the switch.
	// see https://github.com/go-playground/validator#baked-in-validations for list of supported tags
	switch tag {
	case "required":
		return errors.ErrCodeValidationRequiredField
	case "alpha":
		return errors.ErrCodeValidationAlpha
	default:
		return errors.ErrCodeValidationFailed
	}
}

// FormatErrorMessage Formats the error message from the Validation error.
// we could probably use the built in translator (https://github.com/go-playground/validator/blob/master/_examples/translations/main.go)
// but since we dont need i18n in the app for now, let´s keep it simple.
func (cv *RequestValidator) FormatErrorMessage(vE goValidator.FieldError) string {
	message := vE.Error()
	switch vE.Tag() {
	case "required":
		message = fmt.Sprintf("%s is a required field", vE.StructField())
	}

	return message
}
