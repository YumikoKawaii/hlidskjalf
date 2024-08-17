package config

import "github.com/alecthomas/kong"

type Config struct {
	ValgrindGRPCAddress string `env:"VALGRIND_GRPC_ADDRESS" default:"localhost:8090"`
	FrequencyInSec      int    `env:"FREQUENCY_IN_SEC" default:"30"`
	Throughput          int    `env:"THROUGHPUT" default:"100"`
}

func Initialize() (*Config, *kong.Context) {
	cfg := &Config{}
	kongCtx := kong.Parse(cfg)
	return cfg, kongCtx
}
