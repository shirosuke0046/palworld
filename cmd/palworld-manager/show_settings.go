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

func doShowSettingsPre(_ context.Context, cmd *cli.Command) error {
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

func printSettingsJSON(settings *palworld.Settings) error {
	b, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func printSettingsTable(settings *palworld.Settings) {
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

	t.Append([]string{"Difficulty", settings.Difficulty})
	t.Append([]string{"DayTimeSpeedRate", fmt.Sprintf("%f", settings.DayTimeSpeedRate)})
	t.Append([]string{"NightTimeSpeedRate", fmt.Sprintf("%f", settings.NightTimeSpeedRate)})
	t.Append([]string{"ExpRate", fmt.Sprintf("%f", settings.ExpRate)})
	t.Append([]string{"PalCaptureRate", fmt.Sprintf("%f", settings.PalCaptureRate)})
	t.Append([]string{"PalSpawnNumRate", fmt.Sprintf("%f", settings.PalSpawnNumRate)})
	t.Append([]string{"PalDamageRateAttack", fmt.Sprintf("%f", settings.PalDamageRateAttack)})
	t.Append([]string{"PalDamageRateDefense", fmt.Sprintf("%f", settings.PalDamageRateDefense)})
	t.Append([]string{"PlayerDamageRateAttack", fmt.Sprintf("%f", settings.PlayerDamageRateAttack)})
	t.Append([]string{"PlayerDamageRateDefense", fmt.Sprintf("%f", settings.PlayerDamageRateDefense)})
	t.Append([]string{"PlayerStaminaDecreaceRate", fmt.Sprintf("%f", settings.PlayerStomachDecreaceRate)})
	t.Append([]string{"PlayerAutoHPRegeneRate", fmt.Sprintf("%f", settings.PlayerAutoHPRegeneRate)})
	t.Append([]string{"PlayerAutoHpRegeneRateInSleep", fmt.Sprintf("%f", settings.PlayerAutoHpRegeneRateInSleep)})
	t.Append([]string{"PalStomachDecreaceRate", fmt.Sprintf("%f", settings.PalStomachDecreaceRate)})
	t.Append([]string{"PalStaminaDecreaceRate", fmt.Sprintf("%f", settings.PalStaminaDecreaceRate)})
	t.Append([]string{"PalAutoHPRegeneRate", fmt.Sprintf("%f", settings.PalAutoHPRegeneRate)})
	t.Append([]string{"PalAutoHpRegeneRateInSleep", fmt.Sprintf("%f", settings.PalAutoHpRegeneRateInSleep)})
	t.Append([]string{"BuildObjectDamageRate", fmt.Sprintf("%f", settings.BuildObjectDamageRate)})
	t.Append([]string{"BuildObjectDeteriorationDamageRate", fmt.Sprintf("%f", settings.BuildObjectDeteriorationDamageRate)})
	t.Append([]string{"CollectionDropRate", fmt.Sprintf("%f", settings.CollectionDropRate)})
	t.Append([]string{"CollectionObjectHpRate", fmt.Sprintf("%f", settings.CollectionObjectHpRate)})
	t.Append([]string{"CollectionObjectRespawnSpeedRate", fmt.Sprintf("%f", settings.CollectionObjectRespawnSpeedRate)})
	t.Append([]string{"EnemyDropItemRate", fmt.Sprintf("%f", settings.EnemyDropItemRate)})
	t.Append([]string{"DeathPenalty", settings.DeathPenalty})
	t.Append([]string{"bEnablePlayerToPlayerDamage", fmt.Sprintf("%t", settings.B_EnablePlayerToPlayerDamage)})
	t.Append([]string{"bEnableFriendlyFire", fmt.Sprintf("%t", settings.B_EnableFriendlyFire)})
	t.Append([]string{"bEnableInvaderEnemy", fmt.Sprintf("%t", settings.B_EnableInvaderEnemy)})
	t.Append([]string{"bActiveUNKO", fmt.Sprintf("%t", settings.B_ActiveUNKO)})
	t.Append([]string{"bEnableAimAssistPad", fmt.Sprintf("%t", settings.B_EnableAimAssistPad)})
	t.Append([]string{"bEnableAimAssistKeyboard", fmt.Sprintf("%t", settings.B_EnableAimAssistKeyboard)})
	t.Append([]string{"DropItemMaxNum", fmt.Sprintf("%d", settings.DropItemMaxNum)})
	t.Append([]string{"DropItemMaxNum_UNKO", fmt.Sprintf("%d", settings.DropItemMaxNum_UNKO)})
	t.Append([]string{"BaseCampMaxNum", fmt.Sprintf("%d", settings.BaseCampMaxNum)})
	t.Append([]string{"BaseCampWorkerMaxNum", fmt.Sprintf("%d", settings.BaseCampWorkerMaxNum)})
	t.Append([]string{"DropItemAliveMaxHours", fmt.Sprintf("%f", settings.DropItemAliveMaxHours)})
	t.Append([]string{"bAutoResetGuildNoOnlinePlayers", fmt.Sprintf("%t", settings.B_AutoResetGuildNoOnlinePlayers)})
	t.Append([]string{"AutoResetGuildTimeNoOnlinePlayers", fmt.Sprintf("%f", settings.AutoResetGuildTimeNoOnlinePlayers)})
	t.Append([]string{"GuildPlayerMaxNum", fmt.Sprintf("%d", settings.GuildPlayerMaxNum)})
	t.Append([]string{"PalEggDefaultHatchingTime", fmt.Sprintf("%f", settings.PalEggDefaultHatchingTime)})
	t.Append([]string{"WorkSpeedRate", fmt.Sprintf("%f", settings.WorkSpeedRate)})
	t.Append([]string{"bIsMultiplay", fmt.Sprintf("%t", settings.B_IsMultiplay)})
	t.Append([]string{"bIsPvP", fmt.Sprintf("%t", settings.B_IsPvP)})
	t.Append([]string{"bCanPickupOtherGuildDeathPenaltyDrop", fmt.Sprintf("%t", settings.B_CanPickupOtherGuildDeathPenaltyDrop)})
	t.Append([]string{"bEnableNonLoginPenalty", fmt.Sprintf("%t", settings.B_EnableNonLoginPenalty)})
	t.Append([]string{"bEnableFastTravel", fmt.Sprintf("%t", settings.B_EnableFastTravel)})
	t.Append([]string{"bIsStartLocationSelectByMap", fmt.Sprintf("%t", settings.B_IsStartLocationSelectByMap)})
	t.Append([]string{"bExistPlayerAfterLogout", fmt.Sprintf("%t", settings.B_ExistPlayerAfterLogout)})
	t.Append([]string{"bEnableDefenseOtherGuildPlayer", fmt.Sprintf("%t", settings.B_EnableDefenseOtherGuildPlayer)})
	t.Append([]string{"CoopPlayerMaxNum", fmt.Sprintf("%d", settings.CoopPlayerMaxNum)})
	t.Append([]string{"ServerPlayerMaxNum", fmt.Sprintf("%d", settings.ServerPlayerMaxNum)})
	t.Append([]string{"ServerName", settings.ServerName})
	t.Append([]string{"ServerDescription", settings.ServerDescription})
	t.Append([]string{"PublicPort", fmt.Sprintf("%d", settings.PublicPort)})
	t.Append([]string{"PublicIP", settings.PublicIP})
	t.Append([]string{"RCONEnabled", fmt.Sprintf("%t", settings.RCONEnabled)})
	t.Append([]string{"RCONPort", fmt.Sprintf("%d", settings.RCONPort)})
	t.Append([]string{"Region", settings.Region})
	t.Append([]string{"bUseAuth", fmt.Sprintf("%t", settings.B_UseAuth)})
	t.Append([]string{"BanListURL", settings.BanListURL})
	t.Append([]string{"RESTAPIEnabled", fmt.Sprintf("%t", settings.RESTAPIEnabled)})
	t.Append([]string{"RESTAPIPort", fmt.Sprintf("%d", settings.RESTAPIPort)})
	t.Append([]string{"bShowPlayerList", fmt.Sprintf("%t", settings.B_ShowPlayerList)})
	t.Append([]string{"AllowConnectPlatform", settings.AllowConnectPlatform})
	t.Append([]string{"bIsUseBackupSaveData", fmt.Sprintf("%t", settings.B_IsUseBackupSaveData)})
	t.Append([]string{"LogFormatType", settings.LogFormatType})

	t.Render()
}

func doShowSettings(_ context.Context, cmd *cli.Command) error {
	cfg := cmd.Metadata["config"].(*config)
	client := palworld.New(cfg.BaseURL, cfg.Password)

	settings, err := client.Settings()
	if err != nil {
		return err
	}

	switch cmd.String("output") {
	case "table":
		printSettingsTable(settings)
	case "json":
		err = printSettingsJSON(settings)
		if err != nil {
			return err
		}
	}

	return nil
}

var showSettingsCommand = &cli.Command{
	Name:      "show-settings",
	UsageText: "palworld-manager show-settings [options]",
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
	Before: doShowSettingsPre,
	Action: doShowSettings,
}
