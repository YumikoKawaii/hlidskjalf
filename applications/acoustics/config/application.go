package config

import (
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type ErrorEmitterConfig struct {
	Interval int `json:"interval" mapstructure:"interval" yaml:"interval"`
}

type Application struct {
	Server       *server.Config        `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	ErrorEmitter *ErrorEmitterConfig   `json:"error_emitter" mapstructure:"error_emitter" yaml:"error_emitter"`
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
		ErrorEmitter: &ErrorEmitterConfig{Interval: 5},
	}
}
