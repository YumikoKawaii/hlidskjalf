package config

import (
	"github.com/YumikoKawaii/shared/server"
	"github.com/YumikoKawaii/shared/tracer"
)

type Application struct {
	Server       *server.Config
	TracerConfig *tracer.Configuration
}

func loadDefault() *Application {
	return &Application{
		Server:       server.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
	}
}
