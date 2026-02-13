package config

import (
	"github.com/YumikoKawaii/shared/tracer"
)

type ServerConfig struct {
	HTTP string `json:"http" mapstructure:"http" yaml:"http"`
}

type Application struct {
	Server       *ServerConfig         `json:"skidbladnir_server" mapstructure:"skidbladnir_server" yaml:"skidbladnir_server"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
}

func loadDefault() *Application {
	return &Application{
		Server:       &ServerConfig{HTTP: "0.0.0.0:15001"},
		TracerConfig: tracer.DefaultConfig(),
	}
}
