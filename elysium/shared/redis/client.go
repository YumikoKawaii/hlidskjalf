package redis

import "github.com/go-redis/redis/v8"

func Initialize(cfg Config) *redis.Client {

	client := redis.NewClient(
		&redis.Options{
			Addr: cfg.Address,
		},
	)
	return client
}
