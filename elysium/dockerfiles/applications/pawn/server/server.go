package server

import (
	"elysium.com/applications/pawn/config"
	"elysium.com/applications/pawn/server/handler/greet"
	"elysium.com/applications/pawn/services"
	"golang.org/x/xerrors"
	"log"
)

func Serve(cfg *config.Config) {

	svCfg := services.NewConfig(cfg.GRPCPort, cfg.HTTPPort)
	sv := services.NewServer(svCfg)

	greetSv := greet.NewServiceServer()
	if err := sv.Register(greetSv); err != nil {
		panic(xerrors.Errorf("error register greet server: %w", err))
	}

	if err := sv.Serve(); err != nil {
		log.Fatalf(err.Error())
	}
}
