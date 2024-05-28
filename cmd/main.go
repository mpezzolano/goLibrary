package main

import (
	"errors"
	"fmt"
	"goLibrary/chanutils"
	_ "goLibrary/chanutils"
	"goLibrary/errorutils"
	"goLibrary/fileutils"
	_ "goLibrary/fileutils"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

func main() {
	log := logger.NewLogger(logger.INFO)
	log.Info("This is an info message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")

	// Writing to a file
	err := fileutils.WriteFile("example.txt", "This is an example content")
	if err != nil {
		log.Error("Failed to write file: " + err.Error())
	} else {
		log.Info("File written successfully")
	}

	// channels example
	ch := chanutils.CreateChannel(1)

	// Write to the channel
	go chanutils.WriteToChannel(ch, "Hello, Channel!")

	// Read from the channel
	message := chanutils.ReadFromChannel(ch)
	fmt.Println("Received from channel:", message)

	// Create a custom error
	errH := errorutils.New("Something went wrong", 500)
	log.Error(errH.Error())

	// Wrap an existing error
	wrappedErr := errorutils.Wrap(errH, "Additional context")
	log.Error(wrappedErr.Error())

	// Unwrap an error
	unwrappedErr := errorutils.Unwrap(wrappedErr)
	if errors.Is(unwrappedErr, errH) {
		log.Info("The original error was unwrapped successfully")
	} else {
		log.Error("Failed to unwrap the original error")
	}

	fmt.Println("Finished error handling example.")

}
