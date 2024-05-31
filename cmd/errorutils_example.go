package main

import (
	"errors"
	"fmt"
	"goLibrary/errorutils"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

func runErrorUtilsExamples() {
	log := logger.NewLogger(logger.INFO)

	errH := errorutils.New("Something went wrong", 500)
	log.Error(errH.Error())

	wrappedErr := errorutils.Wrap(errH, "Additional context")
	log.Error(wrappedErr.Error())

	unwrappedErr := errorutils.Unwrap(wrappedErr)
	if errors.Is(unwrappedErr, errH) {
		log.Info("The original error was unwrapped successfully")
	} else {
		log.Error("Failed to unwrap the original error")
	}

	fmt.Println("Finished error handling example.")
}
