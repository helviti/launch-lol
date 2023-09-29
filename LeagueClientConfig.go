package main

type LeagueClientConfig struct {
	Install struct {
		CrashReporting struct {
			Enabled bool   `yaml:"enabled"`
			Type    string `yaml:"type"`
		} `yaml:"crash_reporting"`
		GameSettings struct {
			AccountID int  `yaml:"accountId"`
			Modified  bool `yaml:"modified"`
			Timestamp int  `yaml:"timestamp"`
		} `yaml:"game-settings"`
		GameflowPatcherLock           int64       `yaml:"gameflow-patcher-lock"`
		GameflowProcessInfo           interface{} `yaml:"gameflow-process-info"`
		GameflowSpectateReconnectInfo interface{} `yaml:"gameflow-spectate-reconnect-info"`
		Globals                       struct {
			Locale string `yaml:"locale"`
			Region string `yaml:"region"`
		} `yaml:"globals"`
		LcuSettings struct {
			AccountID int  `yaml:"accountId"`
			Modified  bool `yaml:"modified"`
			Timestamp int  `yaml:"timestamp"`
		} `yaml:"lcu-settings"`
		Patcher struct {
			ClientMigrated         bool     `yaml:"client_migrated"`
			ClientPatcherAvailable bool     `yaml:"client_patcher_available"`
			GameMigrated           bool     `yaml:"game_migrated"`
			GamePatcherAvailable   bool     `yaml:"game_patcher_available"`
			Locales                []string `yaml:"locales"`
			Toggles                struct {
				NewClientPatcher int `yaml:"new_client_patcher"`
				NewGamePatcher   int `yaml:"new_game_patcher"`
			} `yaml:"toggles"`
		} `yaml:"patcher"`
		PerksSettings struct {
			AccountID int  `yaml:"accountId"`
			Modified  bool `yaml:"modified"`
			Timestamp int  `yaml:"timestamp"`
		} `yaml:"perks-settings"`
		RiotclientUpgrade struct {
			SeqSuccessCount int `yaml:"seq-success-count"`
			SuccessCount    int `yaml:"success-count"`
		} `yaml:"riotclient-upgrade"`
		RsoAuth struct {
			InstallIdentifier string `yaml:"install-identifier"`
		} `yaml:"rso-auth"`
	} `yaml:"install"`
}
