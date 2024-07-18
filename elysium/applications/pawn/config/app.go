package config

import "github.com/alecthomas/kong"

type Config struct {
	HTTPPort int `env:"HTTP_PORT" default:"8080"`
	GRPCPort int `env:"GRPC_PORT" default:"8090"`
}

func LoadAppConfig() *Config {
	config := &Config{}
	_ = kong.Parse(config)
	return config
}
