package redis

import (
	"context"
	v8 "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, bytes []byte) error
}

type publisherImpl struct {
	redisClient *v8.Client
}

func NewPublisher(client *v8.Client) Publisher {
	return &publisherImpl{
		redisClient: client,
	}
}

func (p *publisherImpl) Publish(ctx context.Context, topic string, bytes []byte) error {

	if err := p.redisClient.Publish(ctx, topic, bytes).Err(); err != nil {
		logrus.Errorf("error while publish message to topic %s: %s", topic, err.Error())
	}
	logrus.Infof("[Valgrind] - Event was sent to topic %s", topic)
	return nil

}
