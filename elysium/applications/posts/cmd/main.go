package main

import (
	"elysium.com/applications/posts/config"
	"elysium.com/applications/posts/server"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, kongCtx := config.Initialize()
	switch cmd := kongCtx.Command(); cmd {
	case "serve":
		server.Serve(cfg)
	default:
		logrus.Fatalf("unknown option: %s", cmd)
	}
}
