package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/users/config"
	"elysium.com/applications/users/pkg/repository"
	"elysium.com/applications/users/pkg/user_service"
	"elysium.com/applications/users/server/handler/user"
	server "elysium.com/applications/users/service"
	"elysium.com/shared/mysql"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Application) {
	prometheus := grpc_prometheus.NewServerMetrics()
	zapLogger := zap.S().Desugar()
	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
	logger := interceptor.Logger{}

	sv := server.NewServer(
		server.NewConfig(cfg.GRPCPort, cfg.HTTPPort),
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			prometheus.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
			logger.Unary("Users"),
		),
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
