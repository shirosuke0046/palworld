package main

import (
	"context"
	"fmt"

	"github.com/shirosuke0046/palworld"
	"github.com/urfave/cli/v3"
)

func doAnnouncePre(_ context.Context, cmd *cli.Command) error {
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

func doAnnounce(_ context.Context, cmd *cli.Command) error {
	message := cmd.String("message")

	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	err := client.Announce(message)
	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}

var announceCommand = &cli.Command{
	Name:      "announce",
	Usage:     "Announce the message to logged-in players",
	UsageText: "palworld-manager announce [options] -m MESSAGE",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:     "message",
			Aliases:  []string{"m"},
			Value:    "",
			Required: true,
		},
	},
	Before: doAnnouncePre,
	Action: doAnnounce,
}
