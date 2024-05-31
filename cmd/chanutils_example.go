package main

import (
	"fmt"
	"goLibrary/chanutils"
	_ "goLibrary/chanutils"
)

func runChanUtilsExamples() {
	ch := chanutils.CreateChannel(1)

	go chanutils.WriteToChannel(ch, "Hello, Channel!")

	message := chanutils.ReadFromChannel(ch)
	fmt.Println("Received from channel:", message)
}
