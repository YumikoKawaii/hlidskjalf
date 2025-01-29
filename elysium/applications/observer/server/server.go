package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/observer/config"
	"elysium.com/applications/observer/server/handler/observer"
	server "elysium.com/applications/observer/service"
	"elysium.com/shared/redis"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {
	prometheus := grpc_prometheus.NewServerMetrics()
	//jwtInterceptor := jwt.NewInterceptor(jwt.NewResolver(&cfg.JWTConfig))

	zapLogger := zap.S().Desugar()
	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
	logger := interceptor.Logger{}

	sv := server.NewServer(
		server.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			prometheus.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			logger.Unary("observer"),
		),
	)

	redisClient := redis.Initialize(cfg.RedisCfg)
	obsHandler := observer.NewService(redisClient)

	if err := sv.Register(obsHandler); err != nil {
		panic(err)
	}

	if err := sv.Serve(); err != nil {
		panic(err)
	}
}
