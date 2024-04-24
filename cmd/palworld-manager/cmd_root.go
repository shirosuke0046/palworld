package main

import (
	"github.com/urfave/cli/v3"
)

var rootCommand = &cli.Command{
	Name:      "palworld-manager",
	Usage:     "Palworld REST API client for CLI",
	UsageText: "palworld-manager <command> [options] [arguments]",
	Commands: []*cli.Command{
		announceCommand,
		forceStopCommand,
		kickCommand,
		printConfigTemplateCommand,
		saveCommand,
		showInformationCommand,
		showMetricsCommand,
		showPlayersCommand,
		showSettingsCommand,
		shutdownCommand,
		versionCommand,
	},
}
