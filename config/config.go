package config

import (
	"github.com/allentom/harukap/config"
)

var DefaultConfigProvider *config.Provider

func InitConfigProvider() error {
	var err error
	DefaultConfigProvider, err = config.NewProvider(func(provider *config.Provider) {
		ReadConfig(provider)
	})
	return err
}

var Instance Config

type Config struct {
	EnableAuth bool
}

func ReadConfig(provider *config.Provider) {
	configer := provider.Manager
	configer.SetDefault("youplus.auth", false)

	Instance = Config{
		EnableAuth: configer.GetBool("youplus.auth"),
	}
}
