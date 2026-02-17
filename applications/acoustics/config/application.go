package config

import (
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type ErrorEmitterConfig struct {
	Interval int `json:"interval" mapstructure:"interval" yaml:"interval"`
}

type RandomDelayConfig struct {
	Base  int64   `json:"base" mapstructure:"base" yaml:"base"`
	Rate  float64 `json:"rate" mapstructure:"rate" yaml:"rate"`
	Value int64   `json:"value" mapstructure:"value" yaml:"value"`
}

type Application struct {
	Server       *server.Config        `json:"server" mapstructure:"server" yaml:"server"`
	Logger       *logger.Configuration `json:"logger" mapstructure:"logger" yaml:"logger"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	ErrorEmitter *ErrorEmitterConfig   `json:"error_emitter" mapstructure:"error_emitter" yaml:"error_emitter"`
	RandomDelay  *RandomDelayConfig    `json:"random_delay" mapstructure:"random_delay" yaml:"random_delay"`
	ErrorRate    float64               `json:"error_rate" mapstructure:"error_rate" yaml:"error_rate"`
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		Logger:       logger.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
		ErrorEmitter: &ErrorEmitterConfig{Interval: 5},
		RandomDelay:  &RandomDelayConfig{},
		ErrorRate:    0,
	}
}
