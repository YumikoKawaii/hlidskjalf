package producer

import (
	"YumikoKawaii/hlidskjalf/golang/common/kafka/config"
	kafkaTypes "YumikoKawaii/hlidskjalf/golang/common/kafka/types"
	"context"
	"errors"
	"github.com/IBM/sarama"
	"golang.org/x/xerrors"
	"log"
	"os"
	"time"
)

const defaultReconnectDelay = time.Second

var ErrorKafkaNotAvailable = errors.New("kafka is not available")

type SyncProducer struct {
	dataCollector sarama.SyncProducer

	config config.ProducerConfig
}

func NewSyncProducer(config config.ProducerConfig) (*SyncProducer, error) {
	if config.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama]", log.LstdFlags)
	}

	dataCollector, err := newDataCollector(config)
	producer := &SyncProducer{
		dataCollector: dataCollector,
		config:        config,
	}

	if err != nil {
		go producer.fetchDataCollector(defaultReconnectDelay)
	}

	return producer, err
}

func (p *SyncProducer) Send(topic string, message *kafkaTypes.Message) (*kafkaTypes.KafkaResponse, error) {
	producerMessage, err := message.ToSaramaProducerMessage()
	if err != nil {
		return nil, xerrors.Errorf("error while converting to producer message: %w", err)
	}
	producerMessage.Topic = topic
	return p.send(producerMessage)
}

func (p *SyncProducer) send(message *sarama.ProducerMessage) (*kafkaTypes.KafkaResponse, error) {
	if p.dataCollector == nil {
		return nil, ErrorKafkaNotAvailable
	}

	partition, offset, err := p.dataCollector.SendMessage(message)
	if err != nil {
		return nil, err
	}
	return &kafkaTypes.KafkaResponse{
		Partition: partition,
		Offset:    offset,
	}, err
}

func (p *SyncProducer) Close(ctx context.Context) error {
	if p.dataCollector != nil {
		if err := p.dataCollector.Close(); err != nil {
			return xerrors.Errorf("failed to shutdown data collector cleanly: %w", err)
		}
	}
	return nil
}

func (p *SyncProducer) fetchDataCollector(delay time.Duration) {
	timer := time.NewTimer(delay)
	for range timer.C {
		timer.Reset(delay)
		dataCollector, err := newDataCollector(p.config)
		if err == nil {
			p.dataCollector = dataCollector
		} else {
			log.Println("can't fetch data collector from kafka: %w", err)
		}
	}
}

func newDataCollector(producerConfig config.ProducerConfig) (sarama.SyncProducer, error) {
	saramaCfg := sarama.NewConfig()

	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Retry.Max = producerConfig.MaxRetries
	saramaCfg.Producer.Return.Successes = true

	saramaCfg.Version = sarama.V3_1_0_0
	saramaCfg.Net.MaxOpenRequests = 1
	saramaCfg.Producer.Idempotent = producerConfig.Idempotent

	if producerConfig.UseCompression {
		saramaCfg.Producer.Compression = sarama.CompressionLZ4
	}

	return sarama.NewSyncProducer(producerConfig.Brokers, saramaCfg)
}
