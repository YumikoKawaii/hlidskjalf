package config

type ProducerConfig struct {
	KafkaConfig `kong:"-"`
	MaxRetries  int    `name:"retries" env:"KAFKA_MAX_RETRIES" default:"3" yaml:"max_retries"`
	Topic       string `name:"topic"`
	Idempotent  bool   `name:"idempotent" env:"KAFKA_PRODUCER_IDEMPOTENT" default:"true" yaml:"idempotent"`
}
