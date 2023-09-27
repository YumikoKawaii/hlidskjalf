package config

import (
	"github.com/alecthomas/kong"
	"log"
)

type AppConfig struct {
	AppMode           string                `name:"mode" env:"APP_MODE" default:"development"`
	SchemaRegistryCfg *SchemaRegistryConfig `kong:"embed"`
	GRPCPort          int
	HTTPPort          int
	SecretKey         string `name:"secret-key" env:"SECRET_KEY" default:"yumiko"`
	SecretHeaderKey   string `name:"secret-header-key" env:"SECRET_HEADER_KEY" default:"yumiko-key"`
}

func Load() *AppConfig {
	var appCfg *AppConfig
	err := kong.Parse(appCfg)
	if err != nil {
		log.Printf("error while parsing config: %w", err)
		return nil
	}
	return appCfg
}
