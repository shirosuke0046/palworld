package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

const configTemplate = `{
  "base_url": "http://127.0.0.1:8212",
  "password": "ENTER ADMIN PASSWORD"
}`

func doPrintConfigTemplate(_ context.Context, _ *cli.Command) error {
	fmt.Println(configTemplate)
	return nil
}

var printConfigTemplateCommand = &cli.Command{
	Name:      "print-config-template",
	Usage:     "Print config template",
	UsageText: "palworld-manager print-config-template",
	Action:    doPrintConfigTemplate,
}
