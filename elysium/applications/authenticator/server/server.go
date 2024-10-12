package server

import (
	"context"
	"elysium.com/applications/authenticator/config"
	"elysium.com/applications/authenticator/pkg/api_key"
	"elysium.com/applications/authenticator/pkg/auth"
	"elysium.com/applications/authenticator/pkg/discovery"
	"elysium.com/applications/authenticator/pkg/jwt"
	"elysium.com/applications/authenticator/pkg/repository"
	"elysium.com/applications/authenticator/server/handler/authenticator"
	server "elysium.com/applications/authenticator/service"
	"elysium.com/shared/mysql"
	"elysium.com/shared/redis"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"time"
)

func Serve(cfg *config.Application) {

	sv := server.NewServer(
		server.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
	)

	gormDB := mysql.Initialize(&cfg.MysqlCfg)
	repo := repository.NewRepository(gormDB)
	authService := auth.NewService(repo, cfg.Secret)
	jwtResolver := jwt.NewResolver(&cfg.JWTConfig)
	apiKeyResolver := api_key.NewResolver(cfg.ApiKeyMapConfigPath)
	authHandler := authenticator.NewService(authService, jwtResolver, apiKeyResolver)

	redisClient := redis.Initialize(cfg.RedisCfg)
	storage := discovery.NewStorage(redisClient, repo)

	if err := sv.Register(authHandler); err != nil {
		panic(err)
	}

	go storage.Operate(context.Background(), time.Duration(cfg.OperateIntervalInSec)*time.Second)

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
