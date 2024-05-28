package config

import (
	"encoding/json"
	"os"
)

// LoadConfig loads configuration from a JSON file
func LoadConfig(filepath string, config interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv(config map[string]string) {
	for key := range config {
		if value, exists := os.LookupEnv(key); exists {
			config[key] = value
		}
	}
}
