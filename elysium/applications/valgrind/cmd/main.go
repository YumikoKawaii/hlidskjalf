package main

import (
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/serve"
	"elysium.com/applications/valgrind/server"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, kongCtx := config.Initialize()
	switch kongCtx.Command() {
	case "serve":
		server.Serve(cfg)
	case "migrate-schema":
		fmt.Println("migrating schema...")
		serve.Migrate(cfg)
	case "consume":
		serve.Consume(cfg)
	default:
		logrus.Fatalf("unexpected command: %v", kongCtx.Command())
	}

}
