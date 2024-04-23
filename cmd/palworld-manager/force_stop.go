package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var forceStopCommand = &cli.Command{
	Name: "force-stop",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
		&cli.BoolFlag{
			Name:    "yes",
			Aliases: []string{"y"},
			Value:   false,
		},
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		fmt.Println(cmd.Bool("yes"))
		return nil
	},
}
