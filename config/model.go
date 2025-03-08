package config

import (
	"simple-wallet/pkg/db"
)

type Configuration struct {
	App      AppConfiguration      `mapstructure:"app"`
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
	Swagger  SwaggerConfiguration  `mapstructure:"swagger" yaml:"swagger"`
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

type SwaggerConfiguration struct {
	Description string   `mapstructure:"description" yaml:"description"`
	Title       string   `mapstructure:"title" yaml:"title"`
	Version     string   `mapstructure:"version" yaml:"version"`
	Schemes     []string `mapstructure:"schemes" yaml:"schemes"`
	Host        string   `mapstructure:"host" yaml:"host"`
	ApiKey      string   `mapstructure:"apiKey" yaml:"apiKey"`
}
