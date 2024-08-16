package redis

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type HandleMessageFn func(bytes []byte) error

type InformMessageHandled func()

type Consumer interface {
	Consume(ctx context.Context, topic string, fn HandleMessageFn)
	Close(ctx context.Context) error
}

type consumerImpl struct {
	redisClient *v8.Client
	subscriber  *v8.PubSub
	informFn    InformMessageHandled
}

func NewConsumer(client *v8.Client, informFn InformMessageHandled) Consumer {
	return &consumerImpl{
		redisClient: client,
		informFn:    informFn,
	}
}

func (c *consumerImpl) Consume(ctx context.Context, topic string, fn HandleMessageFn) {

	subscriber := c.redisClient.Subscribe(ctx, topic)
	c.subscriber = subscriber
	for {
		message, err := c.subscriber.ReceiveMessage(ctx)
		if err != nil {
			logrus.Errorf("error while receiving message: %s", err.Error())
			continue
		}

		bytes := []byte(message.Payload)

		if err := fn(bytes); err != nil {
			logrus.Errorf("error handling message: %s", err.Error())
		}

		c.informFn()
	}
}

func (c *consumerImpl) Close(ctx context.Context) error {
	return c.subscriber.Close()
}
