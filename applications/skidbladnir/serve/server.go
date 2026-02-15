package serve

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/inbound"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/outbound"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler/rate_limiter"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/ratelimit"
	"github.com/YumikoKawaii/shared/health"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Server(_ *cobra.Command, _ []string) {
	logger.Info("[skidbladnir] starting...")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	var rlManager *ratelimit.Manager
	if cfg.RateLimit != nil && cfg.RateLimit.Enabled {
		podIP := cfg.Peer.PodIP
		rlManager = ratelimit.NewManager(podIP)
		ctx := context.Background()
		go rlManager.Start(ctx, cfg.Peer.ServiceName, cfg.Peer.Namespace, cfg.RateLimit.RPS, cfg.RateLimit.Burst, cfg.Peer.LeaderPort)
		logger.Infof("[skidbladnir] rate limiting enabled: rps=%.2f burst=%d", cfg.RateLimit.RPS, cfg.RateLimit.Burst)
	}

	// Shared server handles gRPC (rate limiter protocol) + HTTP (egress proxy, health, metrics)
	instance := server.Initialize(cfg.Server)

	healthHandler := health.Initialize()
	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	rateLimiterHandler := rate_limiter.Initialize(rlManager)
	if err := rateLimiterHandler.Register(instance); err != nil {
		panic(err)
	}

	// Outbound egress proxy wraps the shared mux: internal requests (health,
	// metrics, rate limiter) go to the mux, everything else is proxied.
	processor := outbound.Initialize()
	proxyHandler := processor.Register(instance.HttpMux(), cfg.Server.HTTP)

	h2s := &http2.Server{}
	h2cHandler := h2c.NewHandler(proxyHandler, h2s)
	instance.SetHttpHandler(&h2cHandler)

	// Inbound proxy listeners — rate limit inbound traffic before forwarding to local app
	if cfg.Inbound != nil {
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
	}

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
