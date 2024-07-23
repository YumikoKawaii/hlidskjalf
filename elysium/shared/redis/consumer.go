package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type HandleMessageFn func(bytes []byte) error

type Consumer interface {
	Consume(ctx context.Context, topic string, fn HandleMessageFn) error
	Close(ctx context.Context) error
}

type consumerImpl struct {
	redisClient *redis.Client
	subscriber  *redis.PubSub
}

func NewConsumer(client *redis.Client) Consumer {
	return &consumerImpl{
		redisClient: client,
	}
}

func (c *consumerImpl) Consume(ctx context.Context, topic string, fn HandleMessageFn) error {

	subscriber := c.redisClient.Subscribe(ctx, topic)
	c.subscriber = subscriber
	return nil
}

func (c *consumerImpl) Close(ctx context.Context) error {
	return c.subscriber.Close()
}
