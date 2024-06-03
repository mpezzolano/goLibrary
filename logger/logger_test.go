package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// TestNewLogger checks the creation of a new logger instance
func TestNewLogger(t *testing.T) {
	log := NewLogger(INFO)
	if log.level != INFO {
		t.Errorf("Expected logger level to be %v, got %v", INFO, log.level)
	}
}

// captureOutput captures the output of a function that prints to stdout
func captureOutput(f func()) string {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// Capture output in a goroutine
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Execute the function
	f()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = stdOut

	return <-outC
}

// TestLog checks the Log method with different levels
func TestLog(t *testing.T) {
	log := NewLogger(INFO)

	tests := []struct {
		level LogLevel
		msg   string
	}{
		{INFO, "This is an info message"},
		{WARNING, "This is a warning message"},
		{ERROR, "This is an error message"},
		{DEBUG, "This is a debug message"},
	}

	for _, tt := range tests {
		output := captureOutput(func() {
			log.Log(tt.level, tt.msg)
		})

		if tt.level >= log.level {
			// To handle time differences, we check if the output contains the expected string
			if !containsLogMessage(output, tt.level.String(), tt.msg) {
				t.Errorf("Expected '%s' to contain '%s [%s] %s'", output, time.Now().Format(time.RFC3339), tt.level.String(), tt.msg)
			}
		} else if output != "" {
			t.Errorf("Expected no output, but got '%s'", output)
		}
	}
}

// TestInfo checks the Info method
func TestInfo(t *testing.T) {
	log := NewLogger(INFO)
	output := captureOutput(func() {
		log.Info("This is an info message")
	})

	expected := fmt.Sprintf("%s [INFO] This is an info message\n", time.Now().Format(time.RFC3339))
	if !containsLogMessage(output, "INFO", "This is an info message") {
		t.Errorf("Expected output to contain '%s', but got '%s'", expected, output)
	}
}

// TestWarning checks the Warning method
func TestWarning(t *testing.T) {
	log := NewLogger(INFO)
	output := captureOutput(func() {
		log.Warning("This is a warning message")
	})

	expected := fmt.Sprintf("%s [WARNING] This is a warning message\n", time.Now().Format(time.RFC3339))
	if !containsLogMessage(output, "WARNING", "This is a warning message") {
		t.Errorf("Expected output to contain '%s', but got '%s'", expected, output)
	}
}

// TestError checks the Error method
func TestError(t *testing.T) {
	log := NewLogger(INFO)
	output := captureOutput(func() {
		log.Error("This is an error message")
	})

	expected := fmt.Sprintf("%s [ERROR] This is an error message\n", time.Now().Format(time.RFC3339))
	if !containsLogMessage(output, "ERROR", "This is an error message") {
		t.Errorf("Expected output to contain '%s', but got '%s'", expected, output)
	}
}

// TestDebug checks the Debug method
func TestDebug(t *testing.T) {
	log := NewLogger(DEBUG)
	output := captureOutput(func() {
		log.Debug("This is a debug message")
	})

	expected := fmt.Sprintf("%s [DEBUG] This is a debug message\n", time.Now().Format(time.RFC3339))
	if !containsLogMessage(output, "DEBUG", "This is a debug message") {
		t.Errorf("Expected output to contain '%s', but got '%s'", expected, output)
	}
}

// containsLogMessage checks if the log output contains the expected level and message
func containsLogMessage(output, level, msg string) bool {
	return strings.Contains(output, fmt.Sprintf("[%s] %s", level, msg))
}
