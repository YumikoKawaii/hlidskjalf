package config

import (
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type Application struct {
	Server       *server.Config        `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
	}
}
