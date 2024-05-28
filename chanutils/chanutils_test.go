package chanutils

import (
	"testing"
)

func TestChannelUtils(t *testing.T) {
	ch := CreateChannel(1)

	// Test WriteToChannel
	go WriteToChannel(ch, "test message")

	// Test ReadFromChannel
	msg := ReadFromChannel(ch)
	if msg != "test message" {
		t.Errorf("Expected message 'test message', but got %v", msg)
	}
}
