package errorutils

import (
	"errors"
	"testing"
)

func TestCustomError(t *testing.T) {
	err := New("Test error", 404)
	if err.Error() != "Error 404: Test error" {
		t.Errorf("Expected error message 'Error 404: Test error', but got '%s'", err.Error())
	}
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "Context")
	if wrappedErr.Error() != "Context: Original error" {
		t.Errorf("Expected wrapped error message 'Context: Original error', but got '%s'", wrappedErr.Error())
	}
}

func TestUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "Context")
	unwrappedErr := Unwrap(wrappedErr)
	if !errors.Is(unwrappedErr, originalErr) {
		t.Errorf("Expected unwrapped error to be 'Original error', but got '%s'", unwrappedErr.Error())
	}
}
