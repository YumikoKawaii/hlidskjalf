package server

import (
	"elysium.com/applications/interactions/config"
	"elysium.com/applications/interactions/pkg/interaction_service"
	"elysium.com/applications/interactions/pkg/repository"
	"elysium.com/applications/interactions/server/handler/interactions"
	"elysium.com/applications/interactions/service"
	"elysium.com/applications/interceptor"
	"elysium.com/shared/mysql"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	logger := interceptor.Logger{}

	sv := service.NewServer(
		service.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			logger.Unary("Interactions"),
		),
	)

	gormDB := mysql.Initialize(&cfg.MysqlCfg)
	repo := repository.NewRepository(gormDB)
	interactionService := interaction_service.NewService(repo)
	handler := interactions.NewHandler(interactionService)

	if err := sv.Register(handler); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
