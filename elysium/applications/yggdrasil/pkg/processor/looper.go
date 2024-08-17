package processor

import (
	"context"
	pb "elysium.com/pb/valgrind"
	"elysium.com/shared/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

type Looper interface {
	Resurrect(ctx context.Context)
}

type LooperConfig struct {
	FrequencyInSec int
	Throughput     int
}

func NewLooper(valgrindAddress string, frequency, throughput int) Looper {

	conn, err := grpc.NewClient(valgrindAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	valgrindClient := pb.NewEntryServiceClient(conn)
	return &looperImpl{
		valgrindClient: valgrindClient,
		cfg: LooperConfig{
			FrequencyInSec: frequency,
			Throughput:     throughput,
		},
		generator: NewGenerator(),
	}
}

type looperImpl struct {
	valgrindClient pb.EntryServiceClient
	cfg            LooperConfig
	generator      Generator
}

func (l *looperImpl) Resurrect(ctx context.Context) {

	activeTicker := time.NewTicker(time.Duration(l.cfg.FrequencyInSec) * time.Second)
	logTicker := time.NewTicker(time.Duration(5) * time.Second)

	for {
		select {
		case <-activeTicker.C:
			for _ = range l.cfg.Throughput {
				event := l.generator.Generate(types.SyntheticStudent{}.Subject)
				l.sendEvent(ctx, event)
			}
		case <-logTicker.C:
			logrus.Infof("[Yggdrasil] - Running...")
		}
	}

}

func (l *looperImpl) sendEvent(ctx context.Context, prototype types.Prototype) {

	resp, err := l.valgrindClient.Entry(ctx, buildEntryRequest(prototype.ToBytes()))
	if err != nil {
		logrus.Errorf("[Yggdrasil] - Error sending event to Valgrind: %s", err.Error())
	}

	if resp != nil && resp.Code != http.StatusOK {
		logrus.Infof("[Yggdrasil] - Message from Valgrind: %s", resp.Message)
	}

}

func buildEntryRequest(data []byte) *pb.EntryRequest {
	return &pb.EntryRequest{
		Payload: data,
	}
}
