package redis

import (
	"context"
	"elysium.com/shared/types"
	"github.com/go-redis/redis/v8"
	"google.golang.org/appengine/log"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, prototype types.Prototype) error
}

type publisherImpl struct {
	redisClient *redis.Client
}

func NewPublisher(client *redis.Client) Publisher {
	return &publisherImpl{
		redisClient: client,
	}
}

func (p *publisherImpl) Publish(ctx context.Context, topic string, prototype types.Prototype) error {

	if err := p.redisClient.Publish(ctx, topic, prototype.ToBytes()); err != nil {
		log.Errorf(ctx, "error while publish message to topic %s: %s", topic, err.Err())
	}

	return nil

}
