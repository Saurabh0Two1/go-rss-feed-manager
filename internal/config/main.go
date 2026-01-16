package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const CONFIG_FILENAME = ".gatorconfig.json"

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("failed to access config file")
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read the file")
	}
	defer file.Close()

	cfg := Config{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to access home directory")
	}

	fullPath := filepath.Join(homeDir, CONFIG_FILENAME)

	return fullPath, nil
}

func write(cfg *Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to access home directory")
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
