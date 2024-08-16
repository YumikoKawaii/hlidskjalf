package consume

import (
	"context"
	"elysium.com/shared/clickhouse"
	"elysium.com/shared/redis"
	"elysium.com/shared/types"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	batchSize = 100
)

type Consumer interface {
	Consume(ctx context.Context, topic string)
	Close(ctx context.Context) error
}

type Batch struct {
	subject    string
	prototypes []types.Prototype
}

func InformMessageFn() {
	logrus.Infof("[Valgrind] - Message is processed")
}

func NewConsumer(redisConsumer redis.Consumer, chClient clickhouse.Client, chCfg clickhouse.Config) Consumer {
	return &consumerImpl{
		redisConsumer: redisConsumer,
		chClient:      chClient,
		chCfg:         chCfg,
		eventChan:     make(chan types.Prototype),
		batchChan:     make(chan Batch),
	}
}

type consumerImpl struct {
	redisConsumer redis.Consumer
	chClient      clickhouse.Client
	chCfg         clickhouse.Config
	eventChan     chan types.Prototype
	batchChan     chan Batch
}

// receive events, divide by batch, push to clickhouse
// events will be store in ram, accept falling events
func (c *consumerImpl) Consume(ctx context.Context, topic string) {

	go c.redisConsumer.Consume(ctx, topic, c.handleMessage)
	go c.forkEventChan()
	go c.forkBatchChan(ctx)

}

func (c *consumerImpl) handleMessage(bytes []byte) error {

	prototype, err := types.Analyze(bytes)
	if err != nil {
		return err
	}

	c.eventChan <- prototype
	return nil
}

func (c *consumerImpl) forkEventChan() {

	ticker := time.NewTicker(time.Duration(3) * time.Second)

	bag := make([]types.Prototype, 0)

	for {
		select {
		case <-ticker.C:
			if len(bag) == 0 {
				continue
			}
			batch := Batch{
				prototypes: bag,
			}
			c.batchChan <- batch
			bag = make([]types.Prototype, 0)
			logrus.Infof("[Valgrind] - Flush all events")
		case event := <-c.eventChan:
			bag = append(bag, event)
			if len(bag) == batchSize {
				batch := Batch{
					prototypes: bag,
				}
				c.batchChan <- batch
				bag = make([]types.Prototype, 0)
				logrus.Infof("[Valgrind] - Flush all events")
			}
		}
	}

}

func (c *consumerImpl) forkBatchChan(ctx context.Context) {
	for {
		select {
		case batch := <-c.batchChan:
			logrus.Infof("[Valgrind] - Received batch, length: %d", len(batch.prototypes))
			batches := divideBatchIntoBatchesBySubject(batch)
			for _, b := range batches {
				if err := c.handleBatch(ctx, b); err != nil {
					logrus.Errorf("error handling batch: %s", err.Error())
				}
			}
		}
	}
}

func (c *consumerImpl) handleBatch(ctx context.Context, batch Batch) error {

	// divide batch by subject
	// determine tables
	table := fmt.Sprintf("%s.%s", c.chCfg.Database, batch.prototypes[0].ToClickhouseTableName())
	distributedTable := clickhouse.GetDistributedTable(table)

	// generate query
	values := make([]string, 0)
	for _, prototype := range batch.prototypes {
		values = append(values, types.ToInsertValue(prototype))
	}

	stmt := fmt.Sprintf(
		""+
			"insert into %s "+
			"values %s",
		distributedTable,
		strings.Join(values, ","),
	)

	if err := c.chClient.Exec(ctx, stmt); err != nil {
		return err
	}

	return nil
}

func (c *consumerImpl) Close(ctx context.Context) error {
	return c.redisConsumer.Close(ctx)
}
