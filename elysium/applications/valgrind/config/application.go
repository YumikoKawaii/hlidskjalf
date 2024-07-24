package config

import (
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/redis"
	"github.com/alecthomas/kong"
)

type Config struct {
	MigrateSchema struct{} `kong:"cmd"`
	Serve         struct{} `kong:"cmd"`
	Consume       struct{} `kong:"cmd"`

	HTTPPort      int               `env:"HTTP_PORT" default:"8080"`
	GRPCPort      int               `env:"GRPC_PORT" default:"8090"`
	ClickhouseCfg clickhouse.Config `kong:"embed"`
	RedisCfg      redis.Config      `kong:"embed"`
}

func Initialize() (*Config, *kong.Context) {
	cfg := Config{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
