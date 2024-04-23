package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func doVersion(_ context.Context, _ *cli.Command) error {
	fmt.Printf("palworld-manager v%s (%s)\n", version, revision)
	return nil
}

var versionCommand = &cli.Command{
	Name:   "version",
	Usage:  "Print command version",
	Action: doVersion,
}
