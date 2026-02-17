package config

import (
	"github.com/YumikoKawaii/shared/tracer"
)

type ServerConfig struct {
	HTTP string `json:"http" mapstructure:"http" yaml:"http"`
}

type TransportConfig struct {
	MaxIdleConns        int `json:"max_idle_conns" mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxIdleConnsPerHost int `json:"max_idle_conns_per_host" mapstructure:"max_idle_conns_per_host" yaml:"max_idle_conns_per_host"`
}

type Application struct {
	Server       *ServerConfig         `json:"server" mapstructure:"server" yaml:"server"`
	Transport    *TransportConfig      `json:"transport" mapstructure:"transport" yaml:"transport"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	Namespace    string                `json:"namespace" mapstructure:"namespace" yaml:"namespace"`
}

func loadDefault() *Application {
	return &Application{
		Server:       &ServerConfig{HTTP: "0.0.0.0:10080"},
		Transport:    &TransportConfig{MaxIdleConns: 200, MaxIdleConnsPerHost: 50},
		TracerConfig: tracer.DefaultConfig(),
		Namespace:    "hlidskjalf",
	}
}
