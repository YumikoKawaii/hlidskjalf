package consumer

import (
	kafkaTypes "YumikoKawaii/hlidskjalf/golang/common/kafka/types"
	"github.com/IBM/sarama"
)

type RecordHeaders []*sarama.RecordHeader

type GroupHandler struct {
	HandleRecordFn HandleRecordFn
}

func NewConsumerGroupHandler(fn HandleRecordFn) *GroupHandler {
	return &GroupHandler{
		HandleRecordFn: fn,
	}
}

func (g *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (g *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (g *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			kafkaRecord := &kafkaTypes.Message{
				Key:     message.Key,
				Value:   message.Value,
				Headers: RecordHeaders(message.Headers).ToRecordHeader(),
			}
			g.HandleRecordFn(kafkaRecord)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func (h RecordHeaders) ToRecordHeader() map[string][]byte {
	headers := make(map[string][]byte)
	for _, header := range h {
		headers[string(header.Key)] = header.Value
	}
	return headers
}
