package retry

import (
	"time"
)

// RetryConfig holds the configuration for retry logic
type RetryConfig struct {
	Attempts int
	Delay    time.Duration
}

// DoWithRetry performs the given function with retry logic based on the provided config
func DoWithRetry(fn func() error, config RetryConfig) error {
	var err error
	for i := 0; i < config.Attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(config.Delay)
	}
	return err
}
