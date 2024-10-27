package server

import (
	"elysium.com/applications/interceptor"
	"elysium.com/applications/posts/config"
	"elysium.com/applications/posts/pkg/post_service"
	"elysium.com/applications/posts/pkg/repository"
	"elysium.com/applications/posts/server/handler/posts"
	server "elysium.com/applications/posts/service"
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
			logger.Unary("Posts"),
		),
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
