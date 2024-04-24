package palworld

type Settings struct {
	Difficulty                            string  `json:"Difficulty"`
	DayTimeSpeedRate                      float64 `json:"DayTimeSpeedRate"`
	NightTimeSpeedRate                    float64 `json:"NightTimeSpeedRate"`
	ExpRate                               float64 `json:"ExpRate"`
	PalCaptureRate                        float64 `json:"PalCaptureRate"`
	PalSpawnNumRate                       float64 `json:"PalSpawnNumRate"`
	PalDamageRateAttack                   float64 `json:"PalDamageRateAttack"`
	PalDamageRateDefense                  float64 `json:"PalDamageRateDefense"`
	PlayerDamageRateAttack                float64 `json:"PlayerDamageRateAttack"`
	PlayerDamageRateDefense               float64 `json:"PlayerDamageRateDefense"`
	PlayerStomachDecreaceRate             float64 `json:"PlayerStaminaDecreaceRate"`
	PlayerAutoHPRegeneRate                float64 `json:"PlayerAutoHPRegeneRate"`
	PlayerAutoHpRegeneRateInSleep         float64 `json:"PlayerAutoHpRegeneRateInSleep"`
	PalStomachDecreaceRate                float64 `json:"PalStomachDecreaceRate"`
	PalStaminaDecreaceRate                float64 `json:"PalStaminaDecreaceRate"`
	PalAutoHPRegeneRate                   float64 `json:"PalAutoHPRegeneRate"`
	PalAutoHpRegeneRateInSleep            float64 `json:"PalAutoHpRegeneRateInSleep"`
	BuildObjectDamageRate                 float64 `json:"BuildObjectDamageRate"`
	BuildObjectDeteriorationDamageRate    float64 `json:"BuildObjectDeteriorationDamageRate"`
	CollectionDropRate                    float64 `json:"CollectionDropRate"`
	CollectionObjectHpRate                float64 `json:"CollectionObjectHpRate"`
	CollectionObjectRespawnSpeedRate      float64 `json:"CollectionObjectRespawnSpeedRate"`
	EnemyDropItemRate                     float64 `json:"EnemyDropItemRate"`
	DeathPenalty                          string  `json:"DeathPenalty"`
	B_EnablePlayerToPlayerDamage          bool    `json:"bEnablePlayerToPlayerDamage"`
	B_EnableFriendlyFire                  bool    `json:"bEnableFriendlyFire"`
	B_EnableInvaderEnemy                  bool    `json:"bEnableInvaderEnemy"`
	B_ActiveUNKO                          bool    `json:"bActiveUNKO"`
	B_EnableAimAssistPad                  bool    `json:"bEnableAimAssistPad"`
	B_EnableAimAssistKeyboard             bool    `json:"bEnableAimAssistKeyboard"`
	DropItemMaxNum                        int     `json:"DropItemMaxNum"`
	DropItemMaxNum_UNKO                   int     `json:"DropItemMaxNum_UNKO"`
	BaseCampMaxNum                        int     `json:"BaseCampMaxNum"`
	BaseCampWorkerMaxNum                  int     `json:"BaseCampWorkerMaxNum"`
	DropItemAliveMaxHours                 float64 `json:"DropItemAliveMaxHours"`
	B_AutoResetGuildNoOnlinePlayers       bool    `json:"bAutoResetGuildNoOnlinePlayers"`
	AutoResetGuildTimeNoOnlinePlayers     float64 `json:"AutoResetGuildTimeNoOnlinePlayers"`
	GuildPlayerMaxNum                     int     `json:"GuildPlayerMaxNum"`
	PalEggDefaultHatchingTime             float64 `json:"PalEggDefaultHatchingTime"`
	WorkSpeedRate                         float64 `json:"WorkSpeedRate"`
	B_IsMultiplay                         bool    `json:"bIsMultiplay"`
	B_IsPvP                               bool    `json:"bIsPvP"`
	B_CanPickupOtherGuildDeathPenaltyDrop bool    `json:"bCanPickupOtherGuildDeathPenaltyDrop"`
	B_EnableNonLoginPenalty               bool    `json:"bEnableNonLoginPenalty"`
	B_EnableFastTravel                    bool    `json:"bEnableFastTravel"`
	B_IsStartLocationSelectByMap          bool    `json:"bIsStartLocationSelectByMap"`
	B_ExistPlayerAfterLogout              bool    `json:"bExistPlayerAfterLogout"`
	B_EnableDefenseOtherGuildPlayer       bool    `json:"bEnableDefenseOtherGuildPlayer"`
	CoopPlayerMaxNum                      int     `json:"CoopPlayerMaxNum"`
	ServerPlayerMaxNum                    int     `json:"ServerPlayerMaxNum"`
	ServerName                            string  `json:"ServerName"`
	ServerDescription                     string  `json:"ServerDescription"`
	PublicPort                            int     `json:"PublicPort"`
	PublicIP                              string  `json:"PublicIP"`
	RCONEnabled                           bool    `json:"RCONEnabled"`
	RCONPort                              int     `json:"RCONPort"`
	Region                                string  `json:"Region"`
	B_UseAuth                             bool    `json:"bUseAuth"`
	BanListURL                            string  `json:"BanListURL"`
	RESTAPIEnabled                        bool    `json:"RESTAPIEnabled"`
	RESTAPIPort                           int     `json:"RESTAPIPort"`
	B_ShowPlayerList                      bool    `json:"bShowPlayerList"`
	AllowConnectPlatform                  string  `json:"AllowConnectPlatform"`
	B_IsUseBackupSaveData                 bool    `json:"bIsUseBackupSaveData"`
	LogFormatType                         string  `json:"LogFormatType"`
}

func (c *Client) Settings() (*Settings, error) {
	apiPath := "/v1/api/settings"

	var v Settings
	if err := c.get(apiPath, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
