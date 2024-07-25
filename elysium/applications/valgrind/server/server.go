package server

import (
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/pkg/publish"
	"elysium.com/applications/valgrind/server/handler/entry"
	"elysium.com/applications/valgrind/services"
	"elysium.com/shared/redis"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

func Serve(cfg *config.Config) {

	//ctx := context.Background()

	svCfg := services.NewConfig(cfg.GRPCPort, cfg.HTTPPort)
	sv := services.NewServer(svCfg)

	redisCli := redis.Initialize(cfg.RedisCfg)
	redisPublisher := redis.NewPublisher(redisCli)

	entryService := entry.NewService(publish.NewPublisher(redisPublisher))

	if err := sv.Register(entryService); err != nil {
		panic(xerrors.Errorf("error register greet server: %w", err))
	}

	if err := sv.Serve(); err != nil {
		logrus.Fatalf(err.Error())
	}
}
