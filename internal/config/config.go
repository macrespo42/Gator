package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Can't get home directory")
	}

	configFilePath := filepath.Join(homePath, configFileName)
	return configFilePath, nil
}

func write(cfg Config) error {
	raw_cfg, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, raw_cfg, 0666)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var cfg Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Can't retrieve config file path")
	}

	jsonFile, err := os.Open(configFilePath)

	defer jsonFile.Close()
	if err != nil {
		return Config{}, fmt.Errorf("Can't read config file")
	}

	jsonBody, err := io.ReadAll(jsonFile)

	err = json.Unmarshal(jsonBody, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Can't parse config file")
	}

	return cfg, nil
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	write(*c)
}
