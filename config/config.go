package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	Gin      Gin      `toml:"gin"`
	Mongo    Mongo    `toml:"mongo"`
	Factions Factions `toml:"factions"`
}

type Gin struct {
	Port    string   `toml:"port"`
	Mode    string   `toml:"mode"`
	Domain  string   `toml:"domain"`
	Origins []string `toml:"origins"`
}

type Mongo struct {
	URI          string `toml:"uri"`
	DatabaseName string `toml:"database_name"`
}

type Factions struct {
	MapNumber uint8 `toml:"map_number"`
}

// Prepare attempts to read the config.toml file in to memory
func Prepare() *Configuration {
	f := "config.toml"

	if _, err := os.Stat(f); err != nil {
		f = "default_config.toml"
		fmt.Println("unable to read config.toml, skipping and reading defaults")
	}

	var conf Configuration
	_, err := toml.DecodeFile(f, &conf)

	if err != nil {
		panic("failed to decode config file: " + err.Error())
	}

	return &conf
}
