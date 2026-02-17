package config

import (
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type TransportConfig struct {
	MaxIdleConns        int `json:"max_idle_conns" mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxIdleConnsPerHost int `json:"max_idle_conns_per_host" mapstructure:"max_idle_conns_per_host" yaml:"max_idle_conns_per_host"`
}

type Application struct {
	Server       *server.Config        `json:"server" mapstructure:"server" yaml:"server"`
	Transport    *TransportConfig      `json:"transport" mapstructure:"transport" yaml:"transport"`
	Logger       *logger.Configuration `json:"logger" mapstructure:"logger" yaml:"logger"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	Namespace    string                `json:"namespace" mapstructure:"namespace" yaml:"namespace"`
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		Transport:    &TransportConfig{MaxIdleConns: 200, MaxIdleConnsPerHost: 50},
		Logger:       logger.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
		Namespace:    "hlidskjalf",
	}
}
