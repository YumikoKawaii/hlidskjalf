package config

import (
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/adapters/acoustics"
	"github.com/YumikoKawaii/shared/metrics"
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type ErrorEmitterConfig struct {
	Interval int `json:"interval" mapstructure:"interval" yaml:"interval"`
}

type Application struct {
	Server        *server.Config         `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig  *tracer.Configuration  `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	MetricsConfig *metrics.Configuration `json:"metrics_config" mapstructure:"metrics_config" yaml:"metrics_config"`
	Acoustics     *acoustics.Config      `json:"acoustics" mapstructure:"acoustics" yaml:"acoustics"`
	ErrorEmitter  *ErrorEmitterConfig    `json:"error_emitter" mapstructure:"error_emitter" yaml:"error_emitter"`
}

func loadDefault() *Application {
	return &Application{
		Server:        server.DefaultConfig(),
		TracerConfig:  tracer.DefaultConfig(),
		MetricsConfig: metrics.DefaultConfig(),
		Acoustics:     &acoustics.Config{},
		ErrorEmitter:  &ErrorEmitterConfig{Interval: 5},
	}
}
