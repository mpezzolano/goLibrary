package logger

import (
	"fmt"
	"time"
)

// Logger principal structure
type Logger struct {
	level LogLevel
}

// NewLogger create new instance,  Logger with specific level
func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

// Log print a message with a specific level
func (l *Logger) Log(level LogLevel, msg string) {
	if level >= l.level {
		fmt.Printf("%s [%s] %s\n", time.Now().Format(time.RFC3339), level.String(), msg)
	}
}

// Info methods to log with specific level
func (l *Logger) Info(msg string) {
	l.Log(INFO, msg)
}

func (l *Logger) Warning(msg string) {
	l.Log(WARNING, msg)
}

func (l *Logger) Error(msg string) {
	l.Log(ERROR, msg)
}

func (l *Logger) Debug(msg string) {
	l.Log(DEBUG, msg)
}
