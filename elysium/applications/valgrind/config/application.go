package config

import (
	"elysium.com/shared/clickhouse"
	"github.com/alecthomas/kong"
)

type Config struct {
	MigrateSchema struct{} `kong:"cmd"`

	ClickhouseCfg clickhouse.Config `kong:"embed"`
}

func Initialize() (*Config, *kong.Context) {
	cfg := Config{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
