package config

import (
	"github.com/alecthomas/kong"
)

type StatsRange struct {
	Upper int `env:"STATS_RANGE_UPPER" name:"stats-range-upper"`
	Lower int `env:"STATS_RANGE_LOWER" name:"stats-range-lower"`
}

type StringRange struct {
	Upper int `env:"STRING_RANGE_UPPER" name:"string-range-upper"`
	Lower int `env:"STRING_RANGE_LOWER" name:"string-range-lower"`
}

type PrimaryRange struct {
	Upper int `env:"PRIMARY_RANGE_UPPER" name:"primary-range-upper"`
	Lower int `env:"PRIMARY_RANGE_LOWER" name:"primary-range-lower"`
}

type SecondaryRange struct {
	Upper int `env:"SECONDARY_RANGE_UPPER" name:"secondary-range-upper"`
	Lower int `env:"SECONDARY_RANGE_LOWER" name:"secondary-range-lower"`
}

type Application struct {
	Serve struct{} `kong:"cmd"`

	HTTPPort int `env:"HTTP_PORT" default:"8080"`
	GRPCPort int `env:"GRPC_PORT" default:"8090"`

	StatsRange     StatsRange     `kong:"embed"`
	StringRange    StringRange    `kong:"embed"`
	PrimaryRange   PrimaryRange   `kong:"embed"`
	SecondaryRange SecondaryRange `kong:"embed"`

	TrafficDelayInMilliSec int `env:"TRAFFIC_DELAY_IN_SEC" default:"200"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
