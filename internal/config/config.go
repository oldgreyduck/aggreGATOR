package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func Read() (Config, error) {
	var cfg Config

	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c *Config) SetUser(name string) error {
    c.CurrentUserName = name

    path, err := getConfigFilePath()
    if err != nil {
        return err
    }

    b, err := json.MarshalIndent(c, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(path, b, 0600)
}
