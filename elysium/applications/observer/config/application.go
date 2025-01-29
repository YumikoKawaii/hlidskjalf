package config

import (
	"elysium.com/shared/redis"
	"github.com/alecthomas/kong"
)

type Application struct {
	Serve struct{} `kong:"cmd"`

	HTTPPort int          `env:"HTTP_PORT" default:"8080"`
	GRPCPort int          `env:"GRPC_PORT" default:"8090"`
	RedisCfg redis.Config `kong:"embed"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
