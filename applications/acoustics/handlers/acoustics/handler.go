package acoustics

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/YumikoKawaii/hlidskjalf/applications/acoustics/config"
	api "github.com/YumikoKawaii/rpc.com/protobuf/acoustics"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Handler struct {
	api.UnimplementedAcousticsServer
	errorRate   float64
	randomDelay *config.RandomDelayConfig
}

func Initialize(errorRate float64, randomDelay *config.RandomDelayConfig) *Handler {
	logger.Info("[あこーすてぃくすはんどらー] - はんどらーをしょきか")
	return &Handler{
		errorRate:   errorRate,
		randomDelay: randomDelay,
	}
}

func (h *Handler) simulateDelay() {
	if h.randomDelay.Base > 0 {
		time.Sleep(time.Duration(h.randomDelay.Base))
	}
	if h.randomDelay.Value > 0 && rand.Float64() < h.randomDelay.Rate {
		time.Sleep(time.Duration(rand.Int64N(h.randomDelay.Value)))
	}
}

func (h *Handler) Entry(ctx context.Context, request *api.EntryRequest) (*api.EntryResponse, error) {
	logger.Debug("[あこーすてぃくすはんどらー] - えんとりーりくえすとをじゅしん")

	h.simulateDelay()

	if h.errorRate > 0 && rand.Float64() < h.errorRate {
		logger.Error("[あこーすてぃくすはんどらー] - えんとりーしみゅれーとえらー")
		return nil, status.Error(codes.Internal, "simulated error")
	}

	logger.Debug("[あこーすてぃくすはんどらー] - えんとりーかんりょう")
	return &api.EntryResponse{}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterAcousticsServer(instance.Instance(), h)
	return api.RegisterAcousticsHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
}
