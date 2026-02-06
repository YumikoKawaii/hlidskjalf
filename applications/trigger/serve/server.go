package serve

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/trigger/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/trigger/workers/chaos"
	"github.com/YumikoKawaii/shared/adapters/acoustics"
	"github.com/YumikoKawaii/shared/adapters/echo"
	"github.com/YumikoKawaii/shared/health"
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
	logger.Info("[とりがー] - しょきかちゅう...")

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
			grpcprometheus.UnaryServerInterceptor,
			grpcvalidator.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpcprometheus.StreamServerInterceptor,
			grpcvalidator.StreamServerInterceptor(),
			grpcrecovery.StreamServerInterceptor(),
		),
	)

	acousticsClient, err := acoustics.Initialize(cfg.Acoustics)
	if err != nil {
		panic(err)
	}

	echoClient, err := echo.Initialize(cfg.Echo)
	if err != nil {
		panic(err)
	}

	chaosWorker := chaos.Initialize(cfg.Chaos)
	chaosWorker.Register(
		chaos.ToTriggerFunction[acoustics.EntryRequest, acoustics.EntryResponse]("acoustics.Entry", acousticsClient.Entry),
		chaos.ToTriggerFunction[echo.ChargeRequest, echo.ChargeResponse]("echo.Charge", echoClient.Charge),
		chaos.ToTriggerFunction[echo.DischargeRequest, echo.DischargeResponse]("echo.Discharge", echoClient.Discharge),
	)

	healthHandler := health.Initialize()
	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	chaosWorker.Start(ctx)

	logger.Info("[とりがー] - せつぞくをまっています...")

	if err := instance.Serve(); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("さーばーきどうえらー")
	}
}
