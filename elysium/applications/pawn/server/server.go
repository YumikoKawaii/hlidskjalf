package server

import (
	"context"
	"elysium.com/applications/pawn/config"
	"elysium.com/applications/pawn/pkg/students"
	"elysium.com/applications/pawn/server/handler/greet"
	"elysium.com/applications/pawn/server/handler/performance"
	"elysium.com/applications/pawn/services"
	"elysium.com/shared/clickhouse"
	"golang.org/x/xerrors"
	"log"
)

func Serve(cfg *config.Config) {

	ctx := context.Background()

	svCfg := services.NewConfig(cfg.GRPCPort, cfg.HTTPPort)
	sv := services.NewServer(svCfg)

	greetSv := greet.NewServiceServer()
	if err := sv.Register(greetSv); err != nil {
		panic(xerrors.Errorf("error register greet server: %w", err))
	}

	chClient := clickhouse.Initialize(ctx)
	chCfg := clickhouse.LoadConfig()
	perfSv := performance.NewServer(students.NewService(students.NewRepository(chClient, chCfg)))
	if err := sv.Register(perfSv); err != nil {
		panic(xerrors.Errorf("error register perf server: %w", err))
	}

	if err := sv.Serve(); err != nil {
		log.Fatalf(err.Error())
	}
}
