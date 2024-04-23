package main

import (
	"context"
	"fmt"

	"github.com/shirosuke0046/palworld"
	"github.com/urfave/cli/v3"
)

func doSavePre(ctx context.Context, cmd *cli.Command) error {
	cfgFile := cmd.String("config")
	if cfgFile == "" {
		f, err := defaultConfigFile()
		if err != nil {
			return err
		}
		cfgFile = f
	}

	cfg, err := loadConfig(cfgFile)
	if err != nil {
		return err
	}

	cmd.Metadata["config"] = cfg

	return nil
}

func doSave(ctx context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	err := client.Save()
	if err != nil {
		return err
	}

	fmt.Println("The world data was saved successfully.")

	return nil
}

var saveCommand = &cli.Command{
	Name: "save",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
	},
	Before: doSavePre,
	Action: doSave,
}
