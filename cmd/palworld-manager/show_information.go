package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/shirosuke0046/palworld"
	"github.com/urfave/cli/v3"
)

func doShowInformationPre(_ context.Context, cmd *cli.Command) error {
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

	printType := cmd.String("output")
	switch printType {
	case "table":
	case "json":
	default:
		return fmt.Errorf("unknown output mode '%s': choose 'table' or 'json'", printType)
	}

	return nil
}

func printInformationJSON(information *palworld.Information) error {
	b, err := json.MarshalIndent(information, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func printInformationTable(information *palworld.Information) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoWrapText(false)
	t.SetAutoFormatHeaders(true)
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetRowSeparator("")
	t.SetHeaderLine(false)
	t.SetBorder(false)
	t.SetTablePadding("\t")
	t.SetNoWhiteSpace(true)

	t.SetHeader([]string{"server name", "version", "description"})
	t.Append([]string{information.ServerName, information.Version, information.Description})

	t.Render()
}

func doShowInformation(_ context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	information, err := client.Information()
	if err != nil {
		return err
	}

	switch cmd.String("output") {
	case "table":
		printInformationTable(information)
	case "json":
		err = printInformationJSON(information)
		if err != nil {
			return err
		}
	}

	return nil
}

var showInformationCommand = &cli.Command{
	Name:      "show-information",
	UsageText: "palworld-manager show-information [options]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   "table",
		},
	},
	Before: doShowInformationPre,
	Action: doShowInformation,
}
