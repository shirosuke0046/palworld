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

func doKickPre(ctx context.Context, cmd *cli.Command) error {
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

func doKick(ctx context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	user_id := cmd.String("user-id")

	players, err := client.Players()
	if err != nil {
		return err
	}

	var target *palworld.Player
	for _, p := range players.Players {
		if p.UserId == user_id {
			target = p
			break
		}
	}

	if target == nil {
		return fmt.Errorf("the player with the user id '%s' does not exist", user_id)
	}

	yes := cmd.Bool("yes")
	if !yes {
		fmt.Printf("Are you sure you want to kick %s (id: %s) ?\n", target.Name, target.UserId)
		sc := bufio.NewScanner(os.Stdin)

	Q:
		for {
			fmt.Print("(Y/n): ")

			sc.Scan()
			input := strings.TrimSpace(sc.Text())

			switch input {
			case "Y":
				break Q
			case "n", "N":
				return nil
			default:
			}
		}
	}

	err = client.Kick(target.UserId, "You are kicked.")
	if err != nil {
		return err
	}

	return nil
}

var kickCommand = &cli.Command{
	Name:      "kick",
	Usage:     "Kick the specified logged-in player",
	UsageText: "palworld-manager kick [options] -u USER",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:     "user-id",
			Aliases:  []string{"u"},
			Value:    "",
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "yes",
			Aliases: []string{"y"},
			Value:   false,
		},
	},
	Before: doKickPre,
	Action: doKick,
}
