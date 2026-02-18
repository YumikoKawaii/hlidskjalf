package serve

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/discovery"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/inbound"
	limiterHandler "github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/limiter"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/outbound"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/limiter"
	"github.com/YumikoKawaii/shared/health"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Server(_ *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	_ = logger.Initialize(cfg.Logger)

	logger.Info("[skidbladnir] starting...")

	var rlManager *limiter.Manager
	if cfg.Limiter.Enabled {
		rlManager, err = limiter.NewManager(cfg)
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		go rlManager.Start(ctx)
		logger.Infof("[skidbladnir] rate limiting enabled: rps=%.2f burst=%d", cfg.Limiter.RPS, cfg.Limiter.Burst)
	}

	// Shared server handles gRPC (limiter protocol) + HTTP (egress proxy, health, metrics)
	instance := server.Initialize(cfg.Server)

	healthHandler := health.Initialize()
	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	lh := limiterHandler.Initialize(rlManager)
	if err := lh.Register(instance); err != nil {
		panic(err)
	}

	// Outbound load balancer: lazily resolves pod IPs for upstream services
	resolver, err := discovery.NewResolver(cfg.Namespace)
	if err != nil {
		panic(err)
	}
	resolver.Start(context.Background())

	// Outbound egress proxy wraps the shared mux: internal requests (health,
	// metrics, limiter) go to the mux, everything else is proxied.
	processor := outbound.Initialize(resolver)
	proxyHandler := processor.Register(instance.HttpMux(), cfg.Server.HTTP)

	h2s := &http2.Server{}
	h2cHandler := h2c.NewHandler(proxyHandler, h2s)
	instance.SetHttpHandler(&h2cHandler)

	// Inbound proxy listeners — rate limit inbound traffic before forwarding to local app
	grpcInbound := inbound.NewHandler(cfg.Inbound.GRPC.TargetPort, rlManager)
	go func() {
		if err := grpcInbound.ListenAndServe(cfg.Inbound.GRPC.Endpoint); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("[skidbladnir] inbound gRPC proxy error: %v", err)
		}
	}()

	httpInbound := inbound.NewHandler(cfg.Inbound.HTTP.TargetPort, rlManager)
	go func() {
		if err := httpInbound.ListenAndServe(cfg.Inbound.HTTP.Endpoint); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("[skidbladnir] inbound HTTP proxy error: %v", err)
		}
	}()

	// Graceful shutdown: unregister from leader before exiting so
	// the leader can immediately rebalance fair-share RPS.
	if rlManager != nil {
		go func() {
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
			<-sigCh
			logger.Info("[skidbladnir] shutting down, unregistering from leader...")
			rlManager.Stop()
		}()
	}

	logger.Infof("[skidbladnir] serving: %s...", cfg.Server.HTTP)
	if err := instance.Serve(); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
