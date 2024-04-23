/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli/v3"
)

var (
	appname  = "palworld-manager"
	version  string
	revision string
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

	return filepath.Join(cfgDir, appname, "config.json"), nil
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

func main() {
	cmd := &cli.Command{
		Name:      appname,
		Usage:     "Palworld REST API client for CLI",
		UsageText: "palworld-manager <command> [options] [arguments]",
		Commands: []*cli.Command{
			announceCommand,
			forceStopCommand,
			kickCommand,
			saveCommand,
			showInformationCommand,
			showMetricsCommand,
			showPlayersCommand,
			showSettingsCommand,
			shutdownCommand,
			versionCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
