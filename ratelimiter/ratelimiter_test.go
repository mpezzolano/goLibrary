package ratelimiter

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rate := 5
	bucketSize := 10

	rl := NewRateLimiter(rate, bucketSize)
	defer rl.Stop()

	// Test if initial requests are allowed
	for i := 0; i < bucketSize; i++ {
		if !rl.Allow() {
			t.Errorf("Expected request %d to be allowed", i)
		}
	}

	// Test if additional requests are blocked
	if rl.Allow() {
		t.Errorf("Expected request to be blocked")
	}

	// Wait for the bucket to refill and test again
	time.Sleep(time.Duration(bucketSize/rate+1) * time.Second)
	for i := 0; i < bucketSize; i++ {
		if !rl.Allow() {
			t.Errorf("Expected request %d to be allowed after refill", i)
		}
	}
}
