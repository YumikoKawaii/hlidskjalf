package main

import (
	"elysium.com/applications/pawn/config"
	"elysium.com/applications/pawn/server"
)

func main() {

	cfg := config.LoadAppConfig()

	server.Serve(cfg)
}
