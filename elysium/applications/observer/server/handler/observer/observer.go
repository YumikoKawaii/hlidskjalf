package observer

import (
	"context"
	pb "elysium.com/pb/observer"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type Handler struct {
	*pb.UnimplementedObserverServer
	redisCli *redis.Client
}

func NewService(redisCli *redis.Client) *Handler {
	return &Handler{
		redisCli: redisCli,
	}
}

func (h *Handler) Entry(ctx context.Context, request *pb.EntryRequest) (*pb.EntryResponse, error) {
	// increase
	if _, err := h.redisCli.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
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

	return &pb.EntryResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
