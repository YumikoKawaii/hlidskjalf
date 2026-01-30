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
		Server: &server.Config{
			GRPC: "localhost:10443",
			HTTP: "localhost:10080",
		},
		TracerConfig: tracer.DefaultConfig(),
	}
}
