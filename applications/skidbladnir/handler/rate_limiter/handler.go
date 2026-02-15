package rate_limiter

import (
	"context"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/ratelimit"
	api "github.com/YumikoKawaii/rpc.com/protobuf/rate_limiter"
	"github.com/YumikoKawaii/shared/logger"
	"github.com/YumikoKawaii/shared/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Handler struct {
	api.UnimplementedRateLimiterServer
	manager *ratelimit.Manager
}

func Initialize(manager *ratelimit.Manager) *Handler {
	return &Handler{
		manager: manager,
	}
}

// ServeTokens handles token requests from follower peers.
// Only the leader serves tokens; followers return Unavailable.
func (h *Handler) ServeTokens(_ context.Context, request *api.ServeTokensRequest) (*api.ServeTokensResponse, error) {
	if h.manager == nil {
		return nil, status.Error(codes.Unavailable, "rate limiting not enabled")
	}

	if !h.manager.IsLeader() {
		return nil, status.Error(codes.Unavailable, "not the leader")
	}

	// quantity <= 0 is the unregister signal: peer is shutting down gracefully
	if request.Quantity <= 0 {
		h.manager.UnregisterPeer(request.Ip)
		logger.Infof("[ratelimit] peer %s unregistered", request.Ip)
		return &api.ServeTokensResponse{}, nil
	}

	grant := h.manager.GrantTokens(request.Ip, int(request.Quantity))
	return &api.ServeTokensResponse{
		Quantity: int32(grant.Granted),
		Rps:      float32(grant.RPS),
	}, nil
}

// Status returns the current rate limiter state for debugging.
func (h *Handler) Status(_ context.Context, _ *api.StatusRequest) (*api.StatusResponse, error) {
	if h.manager == nil {
		return &api.StatusResponse{}, nil
	}

	s := h.manager.GetStatus()
	return &api.StatusResponse{
		LeaderIp: s.LeaderIP,
		Ip:       s.PodIP,
		Rps:      float32(s.GlobalRPS),
	}, nil
}

func (h *Handler) Register(instance *server.Server) error {
	api.RegisterRateLimiterServer(instance.Instance(), h)
	if err := api.RegisterRateLimiterHandlerFromEndpoint(
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
