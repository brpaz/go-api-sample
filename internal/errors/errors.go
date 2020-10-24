package errors

import (
	"fmt"
)

const (
	ErrCodeInternalError           = "INTERNAL_ERROR"
	ErrCodeValidationFailed        = "VALIDATION_FAILED"
	ErrCodeValidationRequiredField = "REQUIRED"
	ErrCodeValidationAlpha         = "ALPHA"
	ErrMessageInternalError        = "Internal Error"
	ErrMessageValidationFailed     = "Validation Failed"
)

// ApplicationError struct that holds an Application error.
type ApplicationError struct {
	Code        string
	Message     string
	OriginalErr error
}

// NewApplicationError Creates a new application error
func NewApplicationError(code string, message string) *ApplicationError {
	err := &ApplicationError{
		Code:    code,
		Message: message,
	}

	return err
}

// WithOriginalError Add the original Error to the Application error instance
func (e *ApplicationError) WithOriginalError(err error) *ApplicationError {
	e.OriginalErr = err
	return e
}

// Error Prints the error
func (e *ApplicationError) Error() string {
	if e.OriginalErr != nil {
		return fmt.Sprintf("ERROR: %s (%s) - %v", e.Message, e.Code, e.OriginalErr)
	}

	return fmt.Sprintf("ERROR: %s (%s)", e.Message, e.Code)
}
