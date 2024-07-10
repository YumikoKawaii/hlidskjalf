package server

import (
	"golang.org/x/xerrors"
	"log"
	"pawn.com/config"
	"pawn.com/server/handler/greet"
	"pawn.com/services"
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
