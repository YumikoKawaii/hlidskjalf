package redis

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type HandleMessageFn func(bytes []byte) error

type Consumer interface {
	Consume(ctx context.Context, topic string, fn HandleMessageFn)
	Close(ctx context.Context) error
}

type consumerImpl struct {
	redisClient *v8.Client
	subscriber  *v8.PubSub
}

func NewConsumer(client *v8.Client) Consumer {
	return &consumerImpl{
		redisClient: client,
	}
}

func (c *consumerImpl) Consume(ctx context.Context, topic string, fn HandleMessageFn) {

	subscriber := c.redisClient.Subscribe(ctx, topic)
	c.subscriber = subscriber
	for {
		b, err := c.subscriber.Receive(ctx)
		if err != nil {
			logrus.Errorf("error while receiving message: %s", err.Error())
			continue
		}

		bytes, ok := b.([]byte)
		if !ok {
			logrus.Errorf("error invalid message payload: %v", b)
			continue
		}

		if err := fn(bytes); err != nil {
			logrus.Errorf("error handling message: %s", err.Error())
		}
	}
	//return nil
}

func (c *consumerImpl) Close(ctx context.Context) error {
	return c.subscriber.Close()
}
