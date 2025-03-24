package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	bytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error marshaling struct: %w", err)
	}

	configFilePath, err := configFilePath()
	if err != nil {
		return fmt.Errorf("Error fetching config file path: %w", err)
	}

	if err := os.WriteFile(configFilePath, bytes, 0644); err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	return nil
}

func Read() (Config, error) {
	var config Config

	configFilePath, err := configFilePath()
	if err != nil {
		return config, fmt.Errorf("Error fetching config file path: %w", err)
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return config, fmt.Errorf("Error opening file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return config, fmt.Errorf("Error reading file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("Error marshalling data: %w", err)
	}

	return config, nil
}

func configFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error fetching home dir: %w", err)
	}

	return fmt.Sprintf("%s/%s", dir, configFileName), nil
}