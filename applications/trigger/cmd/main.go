package main

import (
	"github.com/YumikoKawaii/hlidskjalf/applications/trigger/serve"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "rpc-runtime",
	}

	cmd.AddCommand(&cobra.Command{
		Use: "serve",
		Run: serve.Server,
	})

	if err := cmd.Execute(); err != nil {
		logger.Fatalf("failed to execute command: %v", err)
	}
}
