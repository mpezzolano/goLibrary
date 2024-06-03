package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate   int
	bucket int
	mutex  sync.Mutex
	ticker *time.Ticker
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(rate int, bucketSize int) *RateLimiter {
	rl := &RateLimiter{
		rate:   rate,
		bucket: bucketSize,
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
	}

	// Refill the bucket with tokens at a constant rate
	go func() {
		for range rl.ticker.C {
			rl.mutex.Lock()
			if rl.bucket < bucketSize {
				rl.bucket++
			}
			rl.mutex.Unlock()
		}
	}()

	return rl
}

// Allow checks if a request is allowed based on the current rate limit
func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	if rl.bucket > 0 {
		rl.bucket--
		return true
	}
	return false
}

// Stop stops the rate limiter ticker
func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}
