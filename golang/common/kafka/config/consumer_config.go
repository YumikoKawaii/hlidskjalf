package config

type ConsumerConfig struct {
	KafkaConfig `kong:"-"`
	Topic       string `name:"consumer-topic" env:"KAFKA_CONSUMER_TOPIC" default:"yumiko" yaml:"topic"`
	ClientID    string `name:"consumer-client-id" env:"KAFKA_CONSUMER_CLIENT_ID" yaml:"client-id"`
	GroupID     string `name:"consumer-group-id" env:"KAFKA_CONSUMER_GROUP_ID" yaml:"group-id"`
}
