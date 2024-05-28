package errorutils

import (
	"fmt"
)

type CustomError struct {
	Message string
	Code    int
}

func New(message string, code int) error {
	return &CustomError{
		Message: message,
		Code:    code,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

func Unwrap(err error) error {
	return fmt.Errorf("%w", err)
}
