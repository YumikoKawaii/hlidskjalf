package serve

import (
	"context"
	"elysium.com/applications/valgrind/config"
	"elysium.com/applications/valgrind/pkg/consume"
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/redis"
	"elysium.com/shared/topics"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func Consume(cfg *config.Config) {

	ctx := context.Background()

	redisCli := redis.Initialize(cfg.RedisCfg)
	redisConsumer := redis.NewConsumer(redisCli, consume.InformMessageFn)
	chCli := clickhouse.Initialize(ctx, cfg.ClickhouseCfg)

	consumer := consume.NewConsumer(redisConsumer, chCli, cfg.ClickhouseCfg)

	consumer.Consume(ctx, topics.EntryTopic)

	logrus.Infof("[Valgrind] - Serving...")
	// trap here until get terminated

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	<-signals // Blocks here until either SIGINT or SIGTERM is received.

	closeCtx, closeFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer closeFn()

	// we need to close consumer group first to avoid sending to closed worker pool.
	if err := consumer.Close(closeCtx); err != nil {
		logrus.Errorf("error closing consumer: %s", err.Error())
	}

	logrus.Infof("shuting down...")
}
