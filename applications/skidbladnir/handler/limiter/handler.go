package limiter

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/limiter"
	api "github.com/YumikoKawaii/rpc.com/protobuf/limiter"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Handler struct {
	api.UnimplementedLimiterServer
	manager *limiter.Manager
}

func Initialize(manager *limiter.Manager) *Handler {
	return &Handler{
		manager: manager,
	}
}

// Capacity handles capacity requests from followers.
// Delegates to Manager which returns fair-share RPS and burst for the requesting follower.
func (h *Handler) Capacity(_ context.Context, request *api.CapacityRequest) (*api.CapacityResponse, error) {
	if h.manager == nil {
		return nil, status.Error(codes.Unavailable, "unavailable")
	}
	capacity, err := h.manager.Capacity(request.Ip)
	if err != nil {
		logger.Errorf("error capacity allocation: %s", err.Error())
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	return &api.CapacityResponse{
		Rps:   capacity.RPS,
		Burst: capacity.Burst,
	}, nil
}

// Unregister removes a follower from the leader's tracking map on graceful shutdown.
func (h *Handler) Unregister(_ context.Context, request *api.UnregisterRequest) (*api.UnregisterResponse, error) {
	if h.manager == nil {
		return nil, status.Error(codes.Unavailable, "unavailable")
	}
	if err := h.manager.Unregister(request.Ip); err != nil {
		logger.Errorf("error unregister follower: %s", err.Error())
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	return &api.UnregisterResponse{}, nil
}

// Status returns the current rate limiter state for debugging.
func (h *Handler) Status(_ context.Context, _ *api.StatusRequest) (*api.StatusResponse, error) {
	if h.manager == nil {
		return nil, status.Error(codes.Unavailable, "unavailable")
	}
	s := h.manager.GetStatus()
	return &api.StatusResponse{
		LeaderIp: s.LeaderIP,
		Ip:       s.PodIP,
		Rps:      float32(s.GlobalRps),
		Burst:    s.GlobalBurst,
	}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterLimiterServer(instance.Instance(), h)
	if err := api.RegisterLimiterHandlerFromEndpoint(
		context.Background(),
		instance.Mux(),
		instance.GRPCHost(),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Error("[ratelimit] gateway registration failed")
		return err
	}
	return nil
}
