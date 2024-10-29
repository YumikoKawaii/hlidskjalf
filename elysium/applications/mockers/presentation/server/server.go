package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/mockers/presentation/config"
	"elysium.com/applications/mockers/presentation/pkg/adapters/echo"
	"elysium.com/applications/mockers/presentation/pkg/adapters/faux"
	"elysium.com/applications/mockers/presentation/pkg/adapters/fictio"
	"elysium.com/applications/mockers/presentation/pkg/adapters/mimic"
	"elysium.com/applications/mockers/presentation/pkg/adapters/virtual"
	"elysium.com/applications/mockers/presentation/server/handler"
	"elysium.com/applications/mockers/presentation/service"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {

	logger := interceptor.Logger{}

	sv := service.NewServer(
		service.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			logger.Unary("Presentation"),
		),
	)

	echoClient := echo.NewClient(cfg.EchoHost, cfg.UseGrpcProtocol)
	fauxClient := faux.NewClient(cfg.FauxHost, cfg.UseGrpcProtocol)
	fictioClient := fictio.NewClient(cfg.FictioHost, cfg.UseGrpcProtocol)
	mimicClient := mimic.NewClient(cfg.MimicHost, cfg.UseGrpcProtocol)
	virtualClient := virtual.NewClient(cfg.VirtualHost, cfg.UseGrpcProtocol)

	h := handler.NewHandler(*cfg, echoClient, fauxClient, fictioClient, mimicClient, virtualClient)

	if err := sv.Register(h); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
