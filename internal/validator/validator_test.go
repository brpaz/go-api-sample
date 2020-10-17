// +build unit

package validator_test

import (
	appErrors "github.com/brpaz/go-api-sample/internal/errors"
	"github.com/brpaz/go-api-sample/internal/validator"
	ut "github.com/go-playground/universal-translator"
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

	validatorInstance := validator.NewRequestValidator()
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

	validatorInstance := validator.NewRequestValidator()
	for _, tc := range tests {
		code := validatorInstance.MapErrorCodeFromTag(tc.tag)
		assert.Equal(t, code, tc.code)
	}
}

func TestRequestValidator_NormalizeFieldName(t *testing.T) {

	err := &MockValidationError{
		FieldName: "Description",
	}

	validatorInstance := validator.NewRequestValidator()
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

	validatorInstance := validator.NewRequestValidator()
	err := &MockValidationError{
		FieldName: "Description",
		FieldTag:  "required",
	}
	for _, tc := range tests {
		msg := validatorInstance.FormatErrorMessage(err)
		assert.Equal(t, tc.output, msg)
	}
}
