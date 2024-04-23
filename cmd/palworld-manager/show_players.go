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

func doShowPlayersPre(_ context.Context, cmd *cli.Command) error {
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

func printPlayersJSON(players *palworld.Players) error {
	b, err := json.MarshalIndent(players, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func printPlayersTable(players *palworld.Players) {
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

	t.SetHeader([]string{"name", "level", "location", "player id", "user id", "ip", "ping"})
	for _, p := range players.Players {
		t.Append([]string{
			p.Name,
			fmt.Sprintf("%d", p.Level),
			fmt.Sprintf("%f,%f", p.LocationX, p.LocationY),
			p.PlayerId,
			p.UserId,
			p.IP,
			fmt.Sprintf("%f", p.Ping),
		})
	}

	t.Render()
}

func doShowPlayers(ctx context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	players, err := client.Players()
	if err != nil {
		return err
	}

	switch cmd.String("output") {
	case "json":
		err = printPlayersJSON(players)
		if err != nil {
			return err
		}
	case "table":
		printPlayersTable(players)
	}

	return nil
}

var showPlayersCommand = &cli.Command{
	Name:      "show-players",
	UsageText: "palworld-manager show-players [options]",
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
	Before: doShowPlayersPre,
	Action: doShowPlayers,
}
