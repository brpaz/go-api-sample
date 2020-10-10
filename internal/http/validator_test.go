// +build unit

package http_test

import (
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/brpaz/go-api-sample/internal/http"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type MockValidationError struct {
	FieldName  string
	FieldTag   string
	FieldValue interface{}
}

func (e *MockValidationError) Tag() string {
	return e.FieldTag
}

func (e *MockValidationError) ActualTag() string {
	return e.FieldTag
}

func (e *MockValidationError) Namespace() string {
	return e.FieldName
}

func (e *MockValidationError) StructNamespace() string {
	return e.FieldName
}

func (e *MockValidationError) Field() string {
	return e.FieldName
}

func (e *MockValidationError) StructField() string {
	return e.FieldName
}

func (e *MockValidationError) Value() interface{} {
	return e.FieldValue
}

func (e *MockValidationError) Param() string {
	return ""
}

func (e *MockValidationError) Kind() reflect.Kind {
	return reflect.String
}

func (e *MockValidationError) Type() reflect.Type {
	return reflect.TypeOf(e.Value())
}

func (e *MockValidationError) Translate(ut ut.Translator) string {
	return "error"
}

func (e *MockValidationError) Error() string {
	return "error"
}

func TestRequestValidator_Validate(t *testing.T) {
	type mystruct struct {
		Value string `json:"description" validate:"required"`
	}

	validatorInstance := http.NewRequestValidator(validator.New())
	err := validatorInstance.Validate(&mystruct{
		Value: "test",
	})

	assert.Nil(t, err)
}

func TestRequestValidator_MapErrorCodeFromTag(t *testing.T) {
	type test struct {
		tag  string
		code string
	}

	tests := []test{
		{tag: "required", code: appErrors.ErrCodeValidationRequiredField},
		{tag: "alpha", code: appErrors.ErrCodeValidationAlpha},
		{tag: "custom", code: appErrors.ErrCodeValidationFailed},
	}

	validatorInstance := http.NewRequestValidator(validator.New())
	for _, tc := range tests {
		code := validatorInstance.MapErrorCodeFromTag(tc.tag)
		assert.Equal(t, code, tc.code)
	}
}

func TestRequestValidator_NormalizeFieldName(t *testing.T) {

	err := &MockValidationError{
		FieldName: "Description",
	}

	validatorInstance := http.NewRequestValidator(validator.New())
	name := validatorInstance.NormalizeFieldName(err)

	assert.Equal(t, "description", name)
}

func TestRequestValidator_FormatErrorMessage(t *testing.T) {
	type test struct {
		input  string
		output string
	}

	tests := []test{
		{input: "required", output: "Description is a required field"},
	}

	validatorInstance := http.NewRequestValidator(validator.New())
	err := &MockValidationError{
		FieldName: "Description",
		FieldTag:  "required",
	}
	for _, tc := range tests {
		msg := validatorInstance.FormatErrorMessage(err)
		assert.Equal(t, tc.output, msg)
	}
}
