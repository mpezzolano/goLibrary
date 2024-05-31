package main

import (
	"fmt"
	"goLibrary/config"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

type AppConfig struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

func runConfigExamples() {
	log := logger.NewLogger(logger.INFO)
	var appConfig AppConfig

	err := config.LoadConfig("config/test_config.json", &appConfig)
	if err != nil {
		log.Error("Failed to load config: " + err.Error())
	} else {
		log.Info("Config loaded successfully")
	}

	fmt.Printf("App Name: %s, Port: %d\n", appConfig.AppName, appConfig.Port)

	envConfig := map[string]string{
		"APP_NAME": "",
		"PORT":     "",
	}
	config.LoadConfigFromEnv(envConfig)
	fmt.Printf("Env Config - App Name: %s, Port: %s\n", envConfig["APP_NAME"], envConfig["PORT"])
}
