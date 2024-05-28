package main

import (
	"errors"
	"fmt"
	"goLibrary/chanutils"
	_ "goLibrary/chanutils"
	"goLibrary/config"
	"goLibrary/errorutils"
	"goLibrary/fileutils"
	_ "goLibrary/fileutils"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

type AppConfig struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

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

	fmt.Println("---------Configurations----------")
	var appConfig AppConfig
	errC := config.LoadConfig("config/test_config.json", &appConfig)
	if errC != nil {
		log.Error("Failed to load config: " + errC.Error())
	} else {
		log.Info("Config loaded successfully")
	}

	fmt.Printf("App Name: %s, Port: %d\n", appConfig.AppName, appConfig.Port)

	// Example of loading config from environment variables
	envConfig := map[string]string{
		"APP_NAME": "",
		"PORT":     "",
	}
	config.LoadConfigFromEnv(envConfig)
	fmt.Printf("Env Config - App Name: %s, Port: %s\n", envConfig["APP_NAME"], envConfig["PORT"])

	fmt.Println("--------- Finish Configurations----------")
}
