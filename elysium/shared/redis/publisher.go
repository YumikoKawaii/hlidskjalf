package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/appengine/log"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, bytes []byte) error
}

type publisherImpl struct {
	redisClient *redis.Client
}

func NewPublisher(client *redis.Client) Publisher {
	return &publisherImpl{
		redisClient: client,
	}
}

func (p *publisherImpl) Publish(ctx context.Context, topic string, bytes []byte) error {

	if err := p.redisClient.Publish(ctx, topic, bytes); err != nil {
		log.Errorf(ctx, "error while publish message to topic %s: %s", topic, err.Err())
	}

	return nil

}
