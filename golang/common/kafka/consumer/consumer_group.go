package consumer

import (
	kafkaConfig "YumikoKawaii/hlidskjalf/golang/common/kafka/config"
	"context"
	"github.com/IBM/sarama"
	"golang.org/x/xerrors"
	"log"
	"os"
	"strings"
	"time"
)

type Group struct {
	group sarama.ConsumerGroup

	closeChan chan struct{}
	batchSize int
	cycle     time.Duration
}

type Option func(group *Group)

func WithBatchSize(batchSize int) Option {
	return func(group *Group) {
		group.batchSize = batchSize
	}
}

func WithCycle(t time.Duration) Option {
	return func(group *Group) {
		group.cycle = t
	}
}

func NewConsumerGroup(cfg kafkaConfig.ConsumerConfig, opts ...Option) (consumerGroup *Group, err error) {
	if cfg.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)
	}

	var version sarama.KafkaVersion
	if version, err = sarama.ParseKafkaVersion(cfg.Version); err != nil {
		return nil, xerrors.Errorf("error while parsing kafka version: %w", err)
	}

	saramaCfg := sarama.NewConfig()
	saramaCfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaCfg.Version = version

	var group sarama.ConsumerGroup
	if group, err = sarama.NewConsumerGroup(cfg.Brokers, cfg.GroupID, saramaCfg); err != nil {
		return nil, xerrors.Errorf("error while initializing new consumer group: %w", err)
	}

	consumerGroup = &Group{
		group:     group,
		closeChan: make(chan struct{}),
		batchSize: 20,
		cycle:     50 * time.Millisecond,
	}

	for _, opt := range opts {
		opt(consumerGroup)
	}

	return
}

func (g *Group) Consume(ctx context.Context, topic string, fn HandleRecordFn) <-chan error {
	log.Printf("consuming message from topic: %s\n", topic)

	errChan := make(chan error)
	handler := NewConsumerGroupHandler(fn)

	go func(errChan <-chan error) {
		for {
			select {
			case <-g.closeChan:
				log.Println("closing consumer")
				return
			default:
				if err := g.group.Consume(ctx, strings.Split(topic, ","), handler); err != nil {
					log.Printf("error while consunming topic %s: %w\n", topic, err)
				}

				if ctx.Err() != nil {
					log.Printf("context was canceled: %w\n", ctx.Err())
				}
			}
		}
	}(errChan)

	return errChan
}

func (g *Group) Close(ctx context.Context) error {
	close(g.closeChan)
	if err := g.group.Close(); err != nil {
		log.Printf("error while closing consumer group: %w\n", err)
	}
	return nil
}
