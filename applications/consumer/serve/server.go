package serve

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/YumikoKawaii/hlidskjalf/applications/consumer/config"
	"github.com/YumikoKawaii/hlidskjalf/applications/consumer/handlers"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

func Server(_ *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	_ = logger.Initialize(cfg.Logger)

	logger.Info("[こんしゅーまー] - しょきかちゅう...")

	subscriber, err := newSubscriber(cfg.Kafka)
	if err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("さぶすくらいばーきどうえらー")
	}
	defer func() { _ = subscriber.Close() }()

	consumer := handlers.Initialize(cfg.Kafka)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go serveHTTP(cfg.HTTPPort)

	logger.Info("[こんしゅーまー] - しょうひかいし...")
	if err := consumer.Run(ctx, subscriber); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatalf("しょうひえらー")
	}
	logger.Info("[こんしゅーまー] - しゅうりょう")
}

// newSubscriber builds the Watermill Kafka subscriber.
//
// ChannelBufferSize is the "stream capacity" knob. It is set on the underlying
// sarama config here. Note: it does NOT add processing lanes — see the comment
// in config/application.go. The lane count lives in cfg.Concurrency (the worker
// pool), not here.
func newSubscriber(cfg *config.KafkaConfig) (*kafka.Subscriber, error) {
	saramaCfg := kafka.DefaultSaramaSubscriberConfig()
	if cfg.ChannelBufferSize > 0 {
		saramaCfg.ChannelBufferSize = cfg.ChannelBufferSize
	}

	return kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               cfg.Brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaCfg,
			ConsumerGroup:         cfg.ConsumerGroup,
		},
		watermill.NewStdLogger(false, false),
	)
}

func serveHTTP(port int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/api/v1/health/liveness", health)
	mux.HandleFunc("/api/v1/health/readiness", health)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	addr := fmt.Sprintf(":%d", port)
	logger.Infof("[こんしゅーまー] - HTTP/めとりくす: %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Errorf("HTTPさーばーえらー")
		os.Exit(1)
	}
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
