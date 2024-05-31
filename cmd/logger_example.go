package main

import (
	"goLibrary/logger"
	_ "goLibrary/logger"
)

func runLoggerExamples() {
	log := logger.NewLogger(logger.INFO)
	log.Info("This is an info message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")
}
