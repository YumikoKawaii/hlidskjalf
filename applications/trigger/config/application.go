package config

import (
	"github.com/YumikoKawaii/shared/server"
)

type Application struct {
	Server *server.Config `json:"server" mapstructure:"server" yaml:"server"`
}

func loadDefault() *Application {
	return &Application{
		Server: server.DefaultConfig(),
	}
}
