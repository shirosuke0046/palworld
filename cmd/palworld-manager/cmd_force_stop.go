package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/shirosuke0046/palworld"
	"github.com/urfave/cli/v3"
)

func doForceStopPre(_ context.Context, cmd *cli.Command) error {
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

func doForceStop(ctx context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	fmt.Println("Are you sure you want to force-stop the palworld server?")
	sc := bufio.NewScanner(os.Stdin)

Q:
	for {
		fmt.Print("(Y/n): ")

		sc.Scan()
		input := strings.TrimSpace(sc.Text())

		switch input {
		case "y", "Y":
			break Q
		case "n", "N":
			return nil
		default:
		}
	}

	err := client.ForceStop()
	if err != nil {
		return err
	}

	fmt.Println("The palworld server was stopped.")

	return nil
}

var forceStopCommand = &cli.Command{
	Name:      "force-stop",
	Usage:     "Force stop the palworld server (not recommended)",
	UsageText: "palworld-manager force-stop [options]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Usage:   "specify the configration file",
			Aliases: []string{"c"},
			Value:   "",
		},
	},
	Before: doForceStopPre,
	Action: doForceStop,
}
