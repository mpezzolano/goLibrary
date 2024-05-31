package main

import (
	"fmt"
	_ "goLibrary/chanutils"
	_ "goLibrary/fileutils"
)

func main() {
	fmt.Println("Running logger examples")
	runLoggerExamples()

	fmt.Println("Running file utils examples")
	runFileUtilsExamples()

	fmt.Println("Running channel utils examples")
	runChanUtilsExamples()

	fmt.Println("Running error utils examples")
	runErrorUtilsExamples()

	fmt.Println("Running config examples")
	runConfigExamples()

	fmt.Println("Running HTTP client examples")
	runHTTPClientExamples()

}
