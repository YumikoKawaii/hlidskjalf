package publish

import (
	"context"
	"elysium.com/shared/redis"
	"elysium.com/shared/types"
)

type Publisher interface {
	Publish(ctx context.Context, topic string, bytes []byte) error
}

type publisherImpl struct {
	redisPublisher redis.Publisher
}

func NewPublisher(publisher redis.Publisher) Publisher {
	return &publisherImpl{redisPublisher: publisher}
}

func (p *publisherImpl) Publish(ctx context.Context, topic string, bytes []byte) error {

	// analyze if bytes is valid
	_, err := types.Analyze(bytes)
	if err != nil {
		return err
	}

	return p.redisPublisher.Publish(ctx, topic, bytes)
}
