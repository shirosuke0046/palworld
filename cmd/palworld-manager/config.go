package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type config struct {
	BaseURL  string `json:"base_url"`
	Password string `json:"password"`
}

func defaultConfigFile() (string, error) {
	var cfgDir string

	switch runtime.GOOS {
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		cfgDir = filepath.Join(home, ".config")
	default:
		d, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		cfgDir = d
	}

	return filepath.Join(cfgDir, "palworld-manager", "config.json"), nil
}

func loadConfig(cfgFile string) (*config, error) {
	b, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	var cfg config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "http://127.0.0.1:8212"
	}

	return &cfg, nil
}
