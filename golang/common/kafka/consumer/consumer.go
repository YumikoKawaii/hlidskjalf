package consumer

import (
	kafkaTypes "YumikoKawaii/hlidskjalf/golang/common/kafka/types"
	"context"
)

type HandleRecordFn func(record *kafkaTypes.Message)

type Consumer interface {
	Consume(ctx context.Context, topic, fn HandleRecordFn)

	Close(ctx context.Context)
}
