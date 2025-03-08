package config

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	configuration *Configuration
	once          sync.Once
)

type Option func(cfg *Configuration)

// All get all config
func All(opts ...Option) *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AllowEmptyEnv(true)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}

		for _, opt := range opts {
			opt(configuration)
		}

		if configuration.App.ENV == "local" {
			log.Println("config : ", configuration)
		}
	})

	return configuration
}

func Get() *Configuration {
	return configuration
}

// for testing purpose
func Set(conf *Configuration) {
	configuration = conf
}
