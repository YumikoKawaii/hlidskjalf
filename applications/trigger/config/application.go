package config

import (
	"github.com/YumikoKawaii/hlidskjalf/applications/trigger/workers/chaos"
	"github.com/YumikoKawaii/shared/adapters/acoustics"
	"github.com/YumikoKawaii/shared/adapters/echo"
	"github.com/YumikoKawaii/shared/metrics"
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type Application struct {
	Server        *server.Config         `json:"server" mapstructure:"server" yaml:"server"`
	TracerConfig  *tracer.Configuration  `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	MetricsConfig *metrics.Configuration `json:"metrics_config" mapstructure:"metrics_config" yaml:"metrics_config"`
	Acoustics     *acoustics.Config      `json:"acoustics" mapstructure:"acoustics" yaml:"acoustics"`
	Echo          *echo.Config           `json:"echo" mapstructure:"echo" yaml:"echo"`
	Chaos         *chaos.Config          `json:"chaos" mapstructure:"chaos" yaml:"chaos"`
}

func loadDefault() *Application {
	return &Application{
		Server:        server.DefaultConfig(),
		TracerConfig:  tracer.DefaultConfig(),
		MetricsConfig: metrics.DefaultConfig(),
		Acoustics:     &acoustics.Config{},
		Echo:          &echo.Config{},
	}
}
