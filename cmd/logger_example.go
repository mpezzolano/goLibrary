package main

import (
	"github.com/mpezzolano/goLibrary/logger"
)

func RunLoggerExamples() {
	log := logger.NewLogger(logger.INFO)
	log.Info("This is an info message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")
}
