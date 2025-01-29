package observer

import (
	"context"
	"elysium.com/applications/observer/pkg/schemas"
	pb "elysium.com/pb/observer"
	"elysium.com/shared/redis"
	"encoding/json"
	v8 "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"net/http"
)

const topic = "board_casting_information"

type Handler struct {
	*pb.UnimplementedObserverServer
	redisCli  *v8.Client
	publisher redis.Publisher
}

func NewService(redisCli *v8.Client) *Handler {
	return &Handler{
		redisCli:  redisCli,
		publisher: redis.NewPublisher(redisCli),
	}
}

func (h *Handler) Entry(ctx context.Context, request *pb.EntryRequest) (*pb.EntryResponse, error) {
	// increase
	if _, err := h.redisCli.TxPipelined(ctx, func(pipeliner v8.Pipeliner) error {
		if err := pipeliner.Incr(ctx, request.CurrentSession).Err(); err != nil {
			return err
		}
		if err := pipeliner.Decr(ctx, request.PreviousSession).Err(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// retrieve latest count and publish messages ??
	currentSessionView, err := h.redisCli.Get(ctx, request.CurrentSession).Int64()
	if err != nil {
		return nil, err
	}

	prevSessionView, err := h.redisCli.Get(ctx, request.PreviousSession).Int64()
	if err != nil {
		return nil, err
	}

	messages := []*schemas.Message{
		{
			SessionId: request.CurrentSession,
			View:      currentSessionView,
		},
		{
			SessionId: request.PreviousSession,
			View:      prevSessionView,
		},
	}

	for _, m := range messages {
		if err := h.publisher.Publish(ctx, topic, m.ToByte()); err != nil {
			logrus.Errorf("error publishing message: %s", err.Error())
		}
	}

	return &pb.EntryResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}

func (h *Handler) ViewStream(request *pb.ViewStreamRequest, server pb.Observer_ViewStreamServer) error {

	ctx := server.Context()
	// get current view
	view, err := h.redisCli.Get(ctx, request.Session).Int64()
	if err != nil {
		return err
	}
	if err := server.Send(&pb.ViewStreamResponse{
		View: view,
	}); err != nil {
		return err
	}

	consumer := redis.NewConsumer(h.redisCli, func() {
		logrus.Infof("[Observer] - Received message!")
	})

	go func() {
		consumer.Consume(ctx, topic, func(bytes []byte) error {
			message := &schemas.Message{}
			if err := json.Unmarshal(bytes, message); err != nil {
				return err
			}
			if message.SessionId == request.Session {
				return server.Send(&pb.ViewStreamResponse{
					View: message.View,
				})
			} else {
				return nil
			}
		})
	}()

	<-ctx.Done()
	logrus.Infof("[Observer] - Session terminated")
	return nil
}
