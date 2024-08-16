package redis

type Config struct {
	Address string `env:"REDIS_ADDRESS" default:"localhost:6379"`
}
