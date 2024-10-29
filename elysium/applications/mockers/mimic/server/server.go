package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/mockers/mimic/config"
	"elysium.com/applications/mockers/mimic/server/handler"
	"elysium.com/applications/mockers/mimic/service"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	logger := interceptor.Logger{}

	sv := service.NewServer(
		service.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			logger.Unary("Mimic"),
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
