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

func doShutdownPre(_ context.Context, cmd *cli.Command) error {
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

func doShutdown(_ context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	yes := cmd.Bool("yes")
	if !yes {
		fmt.Println("Are you sure you want to shutdown the palworld server?")
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
	}

	waitTime := int(cmd.Int("second"))
	message := cmd.String("message")

	err := client.Shutdown(waitTime, message)
	if err != nil {
		return err
	}

	return nil
}

var shutdownCommand = &cli.Command{
	Name:      "shutdown",
	Usage:     "Shutdown the palworld server",
	UsageText: "palworld-manager shutdown [options]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Usage:   "specify the configration file",
			Aliases: []string{"c"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "message",
			Usage:   "a message to the logged-in players",
			Aliases: []string{"m"},
			Value:   "This server will shutdown in 600 seconds.",
		},
		&cli.IntFlag{
			Name:    "second",
			Usage:   "seconds until shutdown",
			Aliases: []string{"s"},
			Value:   600,
			Validator: func(v int64) error {
				if v < 1 {
					return fmt.Errorf("must be a positive integer")
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "yes",
			Usage:   "skip confirmation if flagged",
			Aliases: []string{"y"},
			Value:   false,
		},
	},
	Before: doShutdownPre,
	Action: doShutdown,
}
