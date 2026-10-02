package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

// Export a config struct that represensts the JSON file structure, including struct tags
type Config struct {
	DB_URL            string `json:"db_url"`
	Current_User_Name string `json:"current_user_name"`
}

// Export a read function that reads the JSON found at ~/.gatorconfig.json and returns a Config struct
// It should read the file from the HOME directory, then decode the JSON string into a new Config struct
// os.UserHomeDir to get the location of home

func ReadConfig() (Config, error) {
	var config Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("failed to get config file path: %v", err)
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	return config, nil
}

// Export a SetUser method on the Config struct that writes the config struct to the JSON file
// after setting the current_user_name field

func (cfg *Config) SetUser(userName string) error {
	cfg.Current_User_Name = userName

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %v", err)
	}

	// Marshal the config struct to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config data: %v", err)
	}

	// Write the marshaled JSON to the config file
	file, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write config data to file: %v", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %v", err)
	}
	configFilePath := homeDir + "/" + configFileName
	return configFilePath, nil
}
