package main

import (
	"fmt"
	"goLibrary/chanutils"
	_ "goLibrary/chanutils"
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
}
