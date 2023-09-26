package types

import (
	"encoding/binary"
	"errors"
	"github.com/IBM/sarama"
	"golang.org/x/xerrors"
)

type (
	SchemaType   string
	EncodingType string
)

var (
	ErrorUnsupportedSchemaType   = errors.New("unsupported schema type")
	ErrorUnsupportedEncodingType = errors.New("unsupported encoding type")
	ErrorNoSchemaType            = errors.New("no schema type")
)

type HeaderOption string

func (h HeaderOption) String() string {
	return HeaderPrefix + string(h)
}

const (
	HeaderPrefix = "hlidskjalf."

	EncodingTypeOption HeaderOption = "encoding-type"
	SchemaTypeOption   HeaderOption = "schema-type"
)

const (
	AvroType     SchemaType = "AVRO"
	ProtobufType SchemaType = "PROTOBUF"

	ConfluentEncodingType EncodingType = "confluent"
	NativeEncodingType    EncodingType = "native"
)

type KafkaResponse struct {
	Partition int32
	Offset    int64
}

type SchemaConfig struct {
	SchemaID     int
	EncodingType EncodingType
	SchemaType   SchemaType
}

type Message struct {
	Key, Value   []byte
	Headers      map[string][]byte
	SchemaConfig *SchemaConfig
}

func (m *Message) PutToHeader(key string, value []byte) {
	if m.Headers == nil {
		m.Headers = make(map[string][]byte)
	}
	m.Headers[key] = value
}

func (m *Message) GetFromHeader(key string) (value []byte, exists bool) {
	value, exists = m.Headers[key]
	return
}

func (m *Message) ResolveSerializationPayload() error {
	if m.SchemaConfig == nil {
		return ErrorNoSchemaType
	}
	if m.SchemaConfig.EncodingType == ConfluentEncodingType {
		return m.resolveConfluentEncodedPayload()
	}
	return ErrorUnsupportedEncodingType
}

func (m *Message) resolveConfluentEncodedPayload() error {
	switch m.SchemaConfig.SchemaType {
	case ProtobufType:
		m.Value = m.resolveProtobufSerializationPayload()
		return nil
	default:
		return ErrorUnsupportedSchemaType
	}
}

func (m *Message) resolveProtobufSerializationPayload() []byte {
	var recordValue []byte

	recordValue = append(recordValue, byte(0))

	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(m.SchemaConfig.SchemaID))
	recordValue = append(recordValue, schemaIDBytes...)

	return recordValue
}

func (m *Message) AdjustMessageBySchemaType(encodingType EncodingType, schemaType SchemaType) error {

	et, encodingFound := m.GetFromHeader(EncodingTypeOption.String())
	st, schemaFound := m.GetFromHeader(SchemaTypeOption.String())

	if encodingFound {
		encodingType = EncodingType(et)
	}

	if schemaFound {
		schemaType = SchemaType(st)
	}

	if encodingType == NativeEncodingType {
		return nil
	}

	switch schemaType {
	case ProtobufType:
		schemaID := binary.BigEndian.Uint32(m.Value[1:5])
		m.SchemaConfig = &SchemaConfig{
			SchemaID:   int(schemaID),
			SchemaType: schemaType,
		}
		m.Value = m.Value[6:]
		return nil
	default:
		return ErrorUnsupportedSchemaType
	}
}

func (m *Message) ToSaramaProducerMessage() (*sarama.ProducerMessage, error) {
	kafkaHeaders := make([]sarama.RecordHeader, 0)

	for key, value := range m.Headers {
		kafkaHeaders = append(kafkaHeaders, sarama.RecordHeader{
			Key:   []byte(key),
			Value: value,
		})
	}

	if err := m.ResolveSerializationPayload(); err != nil {
		return nil, xerrors.Errorf("error while resolving serialization payload for message with key: %s: %w", string(m.Key), err)
	}

	return &sarama.ProducerMessage{
		Key:     sarama.ByteEncoder(m.Key),
		Value:   sarama.ByteEncoder(m.Value),
		Headers: kafkaHeaders,
	}, nil
}
