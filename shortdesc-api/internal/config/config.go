package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type env struct {
	ApiPort           string             `yaml:"api_port"`
	MainMWInstance    string             `yaml:"main_instance"`
	MWInstanceConfigs []MWInstanceConfig `yaml:"mw_instances"`
}

// MWInstanceConfig is a struct to hold the API config of a MediaWiki instance.
type MWInstanceConfig struct {
	ApiUrl string `yaml:"api_url"`
	ID     string `yaml:"id"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
}

// Env holds the environment config.
var Env env

// Load loads the environment config file.
func Load(envFile string) error {
	file, err := os.ReadFile(envFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &Env)
	if err != nil {
		return err
	}

	return nil
}
