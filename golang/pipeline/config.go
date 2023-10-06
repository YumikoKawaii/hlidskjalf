package pipeline

import (
	kafkaConfig "YumikoKawaii/hlidskjalf/golang/common/kafka/config"
)

type Config struct {
	KafkaConfig    kafkaConfig.KafkaConfig    `kong:"embed" yaml:"kafka-config"`
	ConsumerConfig kafkaConfig.ConsumerConfig `kong:"embed" yaml:"consumer-config"`
	ProducerConfig kafkaConfig.ProducerConfig `kong:"embed" yaml:"producer-config"`

	HTTPPort int `env:"HTTP_PORT" default:"8076" yaml:"http-port"`
}
