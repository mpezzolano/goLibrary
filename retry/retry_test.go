package retry

import (
	"errors"
	"testing"
	"time"
)

func TestDoWithRetry_Success(t *testing.T) {
	attempts := 0
	fn := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	}

	config := RetryConfig{Attempts: 5, Delay: 100 * time.Millisecond}
	err := DoWithRetry(fn, config)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if attempts != 3 {
		t.Fatalf("Expected 3 attempts, but got %d", attempts)
	}
}

func TestDoWithRetry_Failure(t *testing.T) {
	fn := func() error {
		return errors.New("permanent error")
	}

	config := RetryConfig{Attempts: 3, Delay: 100 * time.Millisecond}
	err := DoWithRetry(fn, config)
	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}
}
