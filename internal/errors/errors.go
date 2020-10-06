package errors

import (
	"fmt"
)

type ErrorCode int

const (
	ErrCodeInternalError ErrorCode = iota
	ErrCodeValidationFailed
)

type ApplicationError struct {
	Code ErrorCode
	Message string
	OriginalErr error
}

func NewApplicationError(code ErrorCode, message string) *ApplicationError {
	err := &ApplicationError{
		Code: code,
		Message: message,
	}

	return err
}


func (e *ApplicationError) WithOriginalError(err error) *ApplicationError {
	e.OriginalErr = err
	return e
}

func (e *ApplicationError) Error()  string {
	if e.OriginalErr != nil {
		return fmt.Sprintf("ERROR: %s (%d) - %v", e.Message, e.Code, e.OriginalErr)
	}

	return fmt.Sprintf("ERROR: %s (%d)", e.Message, e.Code)
}
