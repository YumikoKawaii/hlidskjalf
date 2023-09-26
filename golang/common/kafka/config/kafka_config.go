package config

type KafkaConfig struct {
	Brokers        []string `name:"brokers" env:"KAFKA_BROKERS" default:"localhost:9092" yaml:"brokers"`
	Version        string   `name:"kafka-version" env:"KAFKA_VERSION" default:"3.3.1" yaml:"version"`
	Verbose        bool     `name:"verbose" env:"KAFKA_VERBOSE" default:"false" yaml:"verbose"`
	UseCompression bool     `name:"use-compression" env:"KAFKA_USE_COMPRESSION" default:"true" yaml:"use-compression"`
}
