package echo

import (
	"context"
	"math/rand/v2"

	api "github.com/YumikoKawaii/rpc.com/protobuf/echo"
	"github.com/YumikoKawaii/shared/adapters/acoustics"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Handler struct {
	api.UnimplementedEchoServer
	acousticsClient acoustics.Client
	errorRate       float64
}

func Initialize(acousticsClient acoustics.Client, errorRate float64) *Handler {
	logger.Info("[えこーはんどらー] - はんどらーをしょきか")
	return &Handler{
		acousticsClient: acousticsClient,
		errorRate:       errorRate,
	}
}

func (h *Handler) Charge(ctx context.Context, request *api.ChargeRequest) (*api.ChargeResponse, error) {
	logger.Info("[えこーはんどらー] - ちゃーじりくえすとをじゅしん")

	if h.errorRate > 0 && rand.Float64() < h.errorRate {
		logger.Error("[えこーはんどらー] - ちゃーじしみゅれーとえらー")
		return nil, status.Error(codes.Internal, "simulated error")
	}

	if _, err := h.acousticsClient.Entry(ctx, acoustics.EntryRequest{}); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Error("[えこーはんどらー] - あこーすてぃくすふぉわーどしっぱい")
		return nil, err
	}

	logger.Info("[えこーはんどらー] - ちゃーじかんりょう")
	return &api.ChargeResponse{}, nil
}

func (h *Handler) Discharge(ctx context.Context, request *api.DischargeRequest) (*api.DischargeResponse, error) {
	logger.Info("[えこーはんどらー] - でぃすちゃーじりくえすとをじゅしん")

	if h.errorRate > 0 && rand.Float64() < h.errorRate {
		logger.Error("[えこーはんどらー] - でぃすちゃーじしみゅれーとえらー")
		return nil, status.Error(codes.Internal, "simulated error")
	}

	logger.Info("[えこーはんどらー] - でぃすちゃーじかんりょう")
	return &api.DischargeResponse{}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterEchoServer(instance.Instance(), h)
	if err := api.RegisterEchoHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Error("[えこーはんどらー] - げーとうぇいとうろくしっぱい")
		return err
	}
	return nil
}
