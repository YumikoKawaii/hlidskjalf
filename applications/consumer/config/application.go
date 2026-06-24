package config

import (
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/tracer"
)

// KafkaConfig holds the Kafka / Watermill subscriber knobs.
//
// Mapping to env vars (Viper replaces "." with "__"):
//
//	KAFKA__BROKERS, KAFKA__TOPIC, KAFKA__CONSUMER_GROUP,
//	KAFKA__CHANNEL_BUFFER_SIZE, KAFKA__CONCURRENCY,
//	KAFKA__ACK_AFTER_WORK, KAFKA__WORK_DELAY_MS
type KafkaConfig struct {
	Brokers       []string `json:"brokers" mapstructure:"brokers" yaml:"brokers"`
	Topic         string   `json:"topic" mapstructure:"topic" yaml:"topic"`
	ConsumerGroup string   `json:"consumer_group" mapstructure:"consumer_group" yaml:"consumer_group"`

	// ChannelBufferSize maps to sarama's ChannelBufferSize — the size of the
	// internal buffer between Kafka fetch and message delivery.
	//
	// IMPORTANT: this is the "stream capacity" knob. Raising it does NOT add
	// processing lanes. It only lets more already-fetched messages wait in
	// memory. If the bottleneck is the ack-gated handler loop, a bigger buffer
	// changes nothing — the messages just queue earlier. It is exposed here so
	// the effect (or lack of it) can be measured directly.
	ChannelBufferSize int `json:"channel_buffer_size" mapstructure:"channel_buffer_size" yaml:"channel_buffer_size"`

	// Concurrency is the number of handler goroutines draining the subscriber
	// channel. THIS is the real lever. With 1, every partition collapses into a
	// single ack-gated lane (steady, low CPU). With N, up to N messages are
	// processed in parallel and CPU scales with N (bounded by partitions).
	Concurrency int `json:"concurrency" mapstructure:"concurrency" yaml:"concurrency"`

	// AckAfterWork controls ack ordering relative to processing.
	//   true  (default) -> ack AFTER work completes  => at-least-once, but a
	//                       Concurrency of 1 serializes on each ack round-trip.
	//   false -> ack immediately, hand work to the pool => at-most-once on crash.
	AckAfterWork bool `json:"ack_after_work" mapstructure:"ack_after_work" yaml:"ack_after_work"`

	// WorkIterations is the number of sha256 rounds run per message. This is the
	// CPU-bound knob: bigger => each message holds a core longer => the CPU graph
	// responds to Concurrency. Tune so one message ~saturates a core for a few ms.
	WorkIterations int `json:"work_iterations" mapstructure:"work_iterations" yaml:"work_iterations"`

	// WorkDelayMs adds an optional I/O-bound tail (e.g. a gRPC call to Acoustics
	// via Bifrost). A single lane spending its time here is CPU-idle — the
	// opposite regime from WorkIterations. 0 by default (pure CPU experiment).
	WorkDelayMs int `json:"work_delay_ms" mapstructure:"work_delay_ms" yaml:"work_delay_ms"`
}

type Application struct {
	Logger       *logger.Configuration `json:"logger" mapstructure:"logger" yaml:"logger"`
	TracerConfig *tracer.Configuration `json:"tracer_config" mapstructure:"tracer_config" yaml:"tracer_config"`
	Kafka        *KafkaConfig          `json:"kafka" mapstructure:"kafka" yaml:"kafka"`

	// HTTPPort serves /debug/pprof and /metrics so the consumer can be profiled
	// the same way as the gRPC services.
	HTTPPort int `json:"http_port" mapstructure:"http_port" yaml:"http_port"`
}

func loadDefault() *Application {
	return &Application{
		Logger:       logger.DefaultConfig(),
		TracerConfig: tracer.DefaultConfig(),
		Kafka: &KafkaConfig{
			Brokers:           []string{"kafka:9092"},
			Topic:             "events",
			ConsumerGroup:     "consumer",
			ChannelBufferSize: 256,
			Concurrency:       1,
			AckAfterWork:      true,
			WorkIterations:    50000,
			WorkDelayMs:       0,
		},
		HTTPPort: 10080,
	}
}
