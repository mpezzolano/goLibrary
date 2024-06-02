package main

import (
	"fmt"
	"github.com/mpezzolano/goLibrary/chanutils"
	_ "github.com/mpezzolano/goLibrary/chanutils"
	_ "github.com/mpezzolano/goLibrary/fileutils"
	"github.com/mpezzolano/goLibrary/logger"
	_ "github.com/mpezzolano/goLibrary/logger"
)

func main() {

	log := logger.NewLogger(logger.INFO)
	log.Info("This is an info message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")

	fmt.Println("another test")
	ch := chanutils.CreateChannel(1)

	go chanutils.WriteToChannel(ch, "Hello, Channel!")

	message := chanutils.ReadFromChannel(ch)
	fmt.Println("Received from channel:", message)

}
