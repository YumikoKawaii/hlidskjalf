package server

import (
	"elysium.com/applications/users/config"
	"elysium.com/applications/users/pkg/repository"
	"elysium.com/applications/users/pkg/user_service"
	"elysium.com/applications/users/server/handler/user"
	server "elysium.com/applications/users/service"
	"elysium.com/shared/mysql"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	sv := server.NewServer(
		server.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
	)

	gormDB := mysql.Initialize(&cfg.MysqlCfg)
	repo := repository.NewRepository(gormDB)
	userService := user_service.NewService(repo)
	handler := user.NewHandler(userService)

	if err := sv.Register(handler); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
