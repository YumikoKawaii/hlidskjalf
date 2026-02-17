package serve

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/discovery"
	"github.com/YumikoKawaii/hlidskjalf/applications/bifrost/handler"
	"github.com/YumikoKawaii/shared/health"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Server(_ *cobra.Command, _ []string) {
	logger.Info("[bifrost] starting...")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	watcher, err := discovery.NewWatcher(cfg.Namespace)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	go watcher.DiscoverServices(ctx)

	instance := server.Initialize(cfg.Server)

	healthHandler := health.Initialize()
	if err := healthHandler.Register(instance); err != nil {
		panic(err)
	}

	processor := handler.Initialize(watcher, cfg.Transport)

	// Wrap the shared mux: registered paths (health, metrics) go to the mux,
	// everything else falls through to the reverse proxy. Then wrap with h2c
	// so Bifrost serves both HTTP/1.1 and HTTP/2 cleartext (gRPC).
	h2s := &http2.Server{}
	h2cHandler := h2c.NewHandler(processor.Wrap(instance.HttpMux()), h2s)
	instance.SetHttpHandler(&h2cHandler)

	logger.Infof("[bifrost] serving: %s...", cfg.Server.HTTP)
	if err := instance.Serve(); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
