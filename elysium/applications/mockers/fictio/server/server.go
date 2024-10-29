package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/mockers/fictio/config"
	"elysium.com/applications/mockers/fictio/server/handler"
	"elysium.com/applications/mockers/fictio/service"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	logger := interceptor.Logger{}

	sv := service.NewServer(
		service.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			logger.Unary("Fictio"),
		),
	)

	h := handler.NewHandler(*cfg)

	if err := sv.Register(h); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
