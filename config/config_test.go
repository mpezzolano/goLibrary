package config

import (
	"os"
	"testing"
)

type TestConfig struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

func TestLoadConfig(t *testing.T) {
	config := TestConfig{}
	err := LoadConfig("config/test_config.json", &config)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.AppName != "MyApp" || config.Port != 8080 {
		t.Errorf("Expected config {AppName: MyApp, Port: 8080}, but got {%v, %v}", config.AppName, config.Port)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("APP_NAME", "EnvApp")
	os.Setenv("PORT", "9090")
	config := map[string]string{
		"APP_NAME": "",
		"PORT":     "",
	}
	LoadConfigFromEnv(config)
	if config["APP_NAME"] != "EnvApp" || config["PORT"] != "9090" {
		t.Errorf("Expected config {APP_NAME: EnvApp, PORT: 9090}, but got {%v, %v}", config["APP_NAME"], config["PORT"])
	}
}
