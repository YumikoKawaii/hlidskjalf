package main

import (
	"context"
	"elysium.com/applications/yggdrasil/config"
	"elysium.com/applications/yggdrasil/pkg/processor"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg, _ := config.Initialize()

	looper := processor.NewLooper(cfg.ValgrindGRPCAddress, cfg.FrequencyInSec, cfg.Throughput)
	logrus.Infof("[Yggdrasil] - Serving...")
	looper.Resurrect(context.Background())
}
