package serve

import (
	"net/http"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/handler"
	"github.com/YumikoKawaii/shared/logger"
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

	processor := handler.Initialize()

	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    cfg.Server.GRPC,
		Handler: h2c.NewHandler(processor.Handler(), h2s),
	}

	logger.Infof("serving: %s...", cfg.Server.GRPC)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
