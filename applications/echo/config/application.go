package config

import (
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/adapters/acoustics"
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type Application struct {
	Server       *server.Config        `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	Acoustics    *acoustics.Config     `json:"acoustics" mapstructure:"acoustics" yaml:"acoustics"`
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
		Acoustics:    &acoustics.Config{},
	}
}
