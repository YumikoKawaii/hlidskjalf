package main

import (
	"pawn.com/config"
	"pawn.com/server"
)

func main() {

	cfg := config.LoadAppConfig()

	server.Serve(cfg)
}
