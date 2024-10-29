package config

import (
	"elysium.com/applications/mockers/utilities"
	"github.com/alecthomas/kong"
)

type Application struct {
	Serve struct{} `kong:"cmd"`

	HTTPPort int `env:"HTTP_PORT" default:"8080"`
	GRPCPort int `env:"GRPC_PORT" default:"8090"`

	TrafficDelayConfig utilities.TrafficDelayConfig `kong:"embed"`

	EchoHost    string `env:"ECHO_HOST"`
	FauxHost    string `env:"FAUX_HOST"`
	FictioHost  string `env:"FICTIO_HOST"`
	MimicHost   string `env:"MIMIC_HOST"`
	VirtualHost string `env:"VIRTUAL_HOST"`

	UseGrpcProtocol     bool `env:"USE_GRPC_PROTOCOL" default:"true"`
	SkipMarshalResponse bool `env:"SKIP_MARSHAL_RESPONSE" default:"false"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
