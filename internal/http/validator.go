package http

import "github.com/go-playground/validator/v10"

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{
		validator: validator,
	}
}

func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
