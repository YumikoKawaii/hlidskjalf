package server

import (
	"elysium.com/applications/center/config"
	"elysium.com/applications/center/pkg/adapter/interactions"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/center/pkg/adapter/users"
	"elysium.com/applications/center/server/handler/center"
	"elysium.com/applications/center/service"
	"elysium.com/applications/interceptor"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	authInterceptor := interceptor.NewInterceptor(cfg.AuthenticatorHost, cfg.UseGRPCProtocol)
	logger := interceptor.Logger{}

	sv := service.NewServer(
		service.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			logger.Unary("Center"),
			authInterceptor.Unary(),
		),
	)

	postClient := posts.NewClient(cfg.PostServiceCfg, cfg.UseGRPCProtocol)
	interactionClient := interactions.NewClient(cfg.InteractionServiceCfg, cfg.UseGRPCProtocol)
	userClient := users.NewClient(cfg.UserServiceCfg, cfg.UseGRPCProtocol)
	handler := center.NewHandler(postClient, interactionClient, userClient)

	if err := sv.Register(handler); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
