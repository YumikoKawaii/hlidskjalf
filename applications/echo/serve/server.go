package serve

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/echo/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/handlers/echo"
	"github.com/YumikoKawaii/hlidskjalf/applications/echo/workers"
	"github.com/YumikoKawaii/shared/adapters/acoustics"
	"github.com/YumikoKawaii/shared/health"
	"github.com/YumikoKawaii/shared/interceptors"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func Server(_ *cobra.Command, _ []string) {
	logger.Info("[えこー] - しょきかちゅう...")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcprometheus.EnableHandlingTimeHistogram()

	instance := server.Initialize(
		cfg.Server,
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryLoggingInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcvalidator.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamLoggingInterceptor(),
			grpcprometheus.StreamServerInterceptor,
			grpcvalidator.StreamServerInterceptor(),
			grpcrecovery.StreamServerInterceptor(),
		),
	)

	healthHandler := health.Initialize()

	acousticsClient, err := acoustics.Initialize(cfg.Acoustics)
	if err != nil {
		panic(err)
	}

	echoHandler := echo.Initialize(acousticsClient, cfg.ErrorRate, cfg.RandomDelay.Rate, cfg.RandomDelay.Value)

	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	if err := echoHandler.Register(instance); err != nil {
		panic(err)
	}

	errorEmitter := &workers.ErrorEmitter{Interval: cfg.ErrorEmitter.Interval}
	errorEmitter.Start(ctx)

	logger.Info("[えこー] - せつぞくをまっています...")

	if err := instance.Serve(); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("さーばーきどうえらー")
	}
}
