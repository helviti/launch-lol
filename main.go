package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var RiotGamesPath = "C:\\Riot Games"
var RiotClientServicesPath = filepath.Join(RiotGamesPath, "RiotClientServices.exe")
var LeaguePath = filepath.Join(RiotGamesPath, "League of Legends")
var LeagueClientPath = filepath.Join(LeaguePath, "LeagueClient.exe")
var configPath = filepath.Join(LeaguePath, "Config", "LeagueClientSettings.yaml")

func EditLeageueClientConfig() {
	var config LeagueClientConfig

	data, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	if config.Install.Globals.Locale == "en_GB" {
		return
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

}

func EditRiotClientSettings() {
	CacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	RiotClientSettingsPath := filepath.Join(CacheDir, "Riot Games", "Riot Client", "Config", "RiotClientSettings.yaml")

	var config RiotClientSettings

	data, err := os.ReadFile(RiotClientSettingsPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, config)

	if config.Install.Globals.Locale == "en_US" {
		return
	}

	config.Install.Globals.Locale = "en_US"

	writeData, err := yaml.Marshal(&config)

	err = os.WriteFile(RiotClientSettingsPath, writeData, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func KillRiotClient() {
	kill := exec.Command("taskkill", "/im", "RiotClientServices.exe", "/T", "/F")
	err := kill.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func LaunchLeague() {
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{nil, nil, nil}
	_, err := os.StartProcess(LeagueClientPath, []string{"--launch-product=league_of_legends", "--launch-patchline=live"}, procAttr)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	KillRiotClient()
	EditRiotClientSettings()
	EditLeageueClientConfig()
	LaunchLeague()
}
