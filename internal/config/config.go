package config

import (
	"encoding/json"
	"io"
	"log"
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
		return "", err
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

	err = os.WriteFile(configFilePath, raw_cfg, 0600)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var cfg Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	jsonBody, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(jsonBody, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		log.Fatal("Can't update gator config")
	}
}
