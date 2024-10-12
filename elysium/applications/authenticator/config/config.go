package config

import (
	"elysium.com/applications/authenticator/pkg/auth"
	"elysium.com/applications/authenticator/pkg/jwt"
	"elysium.com/shared/mysql"
	"elysium.com/shared/redis"
	"github.com/alecthomas/kong"
)

type Application struct {
	Serve   struct{} `kong:"cmd"`
	Consume struct{} `kong:"cmd"`

	HTTPPort int          `env:"HTTP_PORT" default:"8080"`
	GRPCPort int          `env:"GRPC_PORT" default:"8090"`
	MysqlCfg mysql.Config `kong:"embed"`
	RedisCfg redis.Config `kong:"embed"`

	ApiKeyMapConfigPath string      `env:"API_KEY_MAP_CONFIG_PATH"`
	Secret              auth.Secret `kong:"embed"`
	JWTConfig           jwt.Config  `kong:"embed"`

	OperateIntervalInSec int `env:"OPERATE_INTERVAL_IN_SEC"`
}

func Initialize() (*Application, *kong.Context) {
	cfg := Application{}
	kongCtx := kong.Parse(&cfg)
	return &cfg, kongCtx
}
