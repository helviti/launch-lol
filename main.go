package main

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
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

func main() {
	var config Config

	LeaguePath := "C:\\Riot Games\\League of Legends\\"
	ClientPath := filepath.Join(LeaguePath, "LeagueClient.exe")
	configPath := filepath.Join(LeaguePath, "Config\\LeagueClientSettings.yaml")

	data, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	config.Install.Globals.Locale = "en_GB"
	config.Install.Globals.Region = "TR"

	writeData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(configPath, writeData, 0)
	if err != nil {
		log.Fatal(err)
	}

	//
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{nil, nil, nil}
	_, err = os.StartProcess(ClientPath, []string{"--launch-product=league_of_legends", "--launch-patchline=live"}, procAttr)

	if err != nil {
		log.Fatal(err)
	}

}
