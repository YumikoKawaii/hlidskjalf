package config

import (
	"elysium.com/applications/mockers/utilities"
	"github.com/alecthomas/kong"
)

type Application struct {
	Serve struct{} `kong:"cmd"`

	HTTPPort int `env:"HTTP_PORT" default:"8080"`
	GRPCPort int `env:"GRPC_PORT" default:"8090"`

	StatsRange   utilities.StatsRange   `kong:"embed"`
	StringRange  utilities.StringRange  `kong:"embed"`
	PrimaryRange utilities.PrimaryRange `kong:"embed"`

	TrafficDelayConfig utilities.TrafficDelayConfig `kong:"embed"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
