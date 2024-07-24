package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/appengine/log"
)

type HandleMessageFn func(bytes []byte) error

type Consumer interface {
	Consume(ctx context.Context, topic string, fn HandleMessageFn)
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

func (c *consumerImpl) Consume(ctx context.Context, topic string, fn HandleMessageFn) {

	subscriber := c.redisClient.Subscribe(ctx, topic)
	c.subscriber = subscriber
	for {
		b, err := c.subscriber.Receive(ctx)
		if err != nil {
			log.Errorf(ctx, "error while receiving message: %s", err.Error())
			continue
		}

		bytes, ok := b.([]byte)
		if !ok {
			log.Errorf(ctx, "error invalid message payload: %v", b)
			continue
		}

		if err := fn(bytes); err != nil {
			log.Errorf(ctx, "error handling message: %s", err.Error())
		}
	}
	//return nil
}

func (c *consumerImpl) Close(ctx context.Context) error {
	return c.subscriber.Close()
}
