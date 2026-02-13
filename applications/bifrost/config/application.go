package config

import (
	"github.com/YumikoKawaii/shared/tracer"
)

type ServerConfig struct {
	HTTP string `json:"http" mapstructure:"http" yaml:"http"`
}

type Application struct {
	Server       *ServerConfig        `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	Namespace    string                `json:"namespace" mapstructure:"namespace" yaml:"namespace"`
}

func loadDefault() *Application {
	return &Application{
		Server:       &ServerConfig{HTTP: "0.0.0.0:10080"},
		TracerConfig: tracer.DefaultConfig(),
		Namespace:    "hlidskjalf",
	}
}
