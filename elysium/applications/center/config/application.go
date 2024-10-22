package config

import (
	"elysium.com/applications/center/pkg/adapter/interactions"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/center/pkg/adapter/users"
	"github.com/alecthomas/kong"
)

type Application struct {
	Serve struct{} `kong:"cmd"`

	HTTPPort int `env:"HTTP_PORT" default:"8080"`
	GRPCPort int `env:"GRPC_PORT" default:"8090"`

	PostServiceCfg        posts.Config        `kong:"embed"`
	InteractionServiceCfg interactions.Config `kong:"embed"`
	UserServiceCfg        users.Config        `kong:"embed"`
	AuthenticatorHost     string              `env:"AUTHENTICATOR_HOST"`

	UseHTTPProtocol bool `env:"USE_GRPC_PROTOCOL" default:"false"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
