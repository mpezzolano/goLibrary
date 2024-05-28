package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	log := NewLogger(INFO)

	if log.level != INFO {
		t.Errorf("Expected log level to be INFO, but got %v", log.level)
	}

}
