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

func doShowMetricsPre(_ context.Context, cmd *cli.Command) error {
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

func printMetricsJSON(metrics *palworld.Metrics) error {
	b, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func printMetricsTable(metrics *palworld.Metrics) {
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

	t.SetHeader([]string{"current player num", "max player num", "server fps", "server frame time", "uptime"})
	t.Append([]string{
		fmt.Sprintf("%d", metrics.CurrentPlayerNum),
		fmt.Sprintf("%d", metrics.MaxPlayerNum),
		fmt.Sprintf("%d", metrics.ServerFPS),
		fmt.Sprintf("%f", metrics.ServerFrameTime),
		fmt.Sprintf("%d", metrics.Uptime),
	})

	t.Render()
}

func doShowMetrics(_ context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	metrics, err := client.Metrics()
	if err != nil {
		return err
	}

	switch cmd.String("output") {
	case "table":
		printMetricsTable(metrics)
	case "json":
		err = printMetricsJSON(metrics)
		if err != nil {
			return err
		}
	}

	return nil
}

var showMetricsCommand = &cli.Command{
	Name:      "show-metrics",
	UsageText: "palworld-manager show-metrics [options]",
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
	Before: doShowMetricsPre,
	Action: doShowMetrics,
}
