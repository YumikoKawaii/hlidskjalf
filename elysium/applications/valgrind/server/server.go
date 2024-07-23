package server

import (
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/server/handler/entry"
	"elysium.com/applications/valgrind/services"
	"golang.org/x/xerrors"
	"log"
)

func Serve(cfg *config.Config) {

	//ctx := context.Background()

	svCfg := services.NewConfig(cfg.GRPCPort, cfg.HTTPPort)
	sv := services.NewServer(svCfg)

	entryService := entry.NewService()

	if err := sv.Register(entryService); err != nil {
		panic(xerrors.Errorf("error register greet server: %w", err))
	}

	if err := sv.Serve(); err != nil {
		log.Fatalf(err.Error())
	}
}
