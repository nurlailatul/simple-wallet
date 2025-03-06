package config

import (
	"simple-wallet/pkg/db"
)

type Configuration struct {
	App      AppConfiguration      `mapstructure:"app"`
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
}

type AppConfiguration struct {
	Name string `mapstructure:"name"`
	ENV  string `mapstructure:"env"`
}

type ServerConfiguration struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfiguration struct {
	Main db.Config `mapstructure:"main"`
}
