package main

import "github.com/urfave/cli/v3"

var saveCommand = &cli.Command{
	Name: "save",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
	},
}
