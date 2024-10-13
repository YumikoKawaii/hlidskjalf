package server

import (
	"elysium.com/applications/posts/config"
	"elysium.com/applications/posts/pkg/post_service"
	"elysium.com/applications/posts/pkg/repository"
	"elysium.com/applications/posts/server/handler/posts"
	server "elysium.com/applications/posts/service"
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
	postService := post_service.NewService(repo)
	handler := posts.NewHandler(postService)

	if err := sv.Register(handler); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
