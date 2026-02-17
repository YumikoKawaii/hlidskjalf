package limiter

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/YumikoKawaii/hlidskjalf/applications/skidbladnir/constants"
	api "github.com/YumikoKawaii/rpc.com/protobuf/limiter"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

// Operator is the per-pod rate limiter. Every Skidbladnir instance (including
// the leader) runs an Operator. It owns a local token bucket that is checked
// synchronously on every proxied request (no network hop).
//
// A background goroutine periodically fetches its fair-share capacity from
// the Coordinator (leader) via gRPC, updating the local limiter's rate and
// burst. This fetch also serves as a heartbeat so the Coordinator knows
// this operator is alive.
type Operator struct {
	limiter *rate.Limiter // local token bucket, rate/burst set by coordinator

	podIP      string       // this pod's IP, used as identity in capacity requests
	leaderIP   atomic.Value // current leader IP, updated on election changes
	leaderPort int          // gRPC port of the leader's shared server

	limiterClient api.LimiterClient
}

func NewOperator(podIP string, leaderPort int) *Operator {
	return &Operator{
		limiter:    rate.NewLimiter(0, 0), // start at zero; coordinator sets fair-share on first fetch
		podIP:      podIP,
		leaderPort: leaderPort,
	}
}

// Allow performs a pure in-memory token bucket check. Called on every proxied request.
func (o *Operator) Allow() bool {
	return o.limiter.Allow()
}

// SetLeader updates the leader IP and establishes a gRPC connection to the
// new leader's Coordinator. Called by Manager on election state changes.
// The connection is created once and reused for all subsequent RPCs.
func (o *Operator) SetLeader(ip string) error {
	target := fmt.Sprintf("%s:%d", ip, o.leaderPort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("[operator] failed to connect to coordinator %s: %v", target, err)
		return err
	}
	o.limiterClient = api.NewLimiterClient(conn)
	o.leaderIP.Store(ip)
	return nil
}

// Fetch periodically queries the Coordinator for this operator's fair-share capacity.
// Runs an immediate fetch on startup (so the limiter is configured before traffic
// arrives), then ticks at FollowerInterval. Each fetch also serves as a heartbeat
// so the Coordinator knows this operator is alive.
func (o *Operator) Fetch(ctx context.Context) {
	ticker := time.NewTicker(constants.FollowerInterval)
	defer ticker.Stop()
	o.fetch(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			o.fetch(ctx)
		}
	}
}

// Unregister notifies the Coordinator that this operator is shutting down.
// The Coordinator removes it from the tracking map immediately, allowing
// rebalancing without waiting for TTL expiry.
func (o *Operator) Unregister() {
	ctx, cancel := context.WithTimeout(context.Background(), constants.FollowerTTL)
	defer cancel()
	_, err := o.limiterClient.Unregister(ctx, &api.UnregisterRequest{
		Ip: o.podIP,
	})
	if err != nil {
		logger.Errorf("[operator] unregister: failed to notify coordinator %s: %v", o.leaderIP.Load(), err)
		return
	}
	logger.Info("[operator] unregistered from coordinator")
}

// fetch requests the current fair-share RPS and burst from the Coordinator,
// then updates the local token bucket accordingly.
func (o *Operator) fetch(ctx context.Context) {
	if o.limiterClient == nil {
		return // leader not yet known
	}
	capacity, err := o.limiterClient.Capacity(ctx, &api.CapacityRequest{
		Ip: o.podIP,
	})
	if err != nil {
		logger.Errorf("[operator] failed to fetch capacity from coordinator %s: %v", o.leaderIP.Load(), err)
		return
	}

	o.limiter.SetLimit(rate.Limit(capacity.Rps))
	o.limiter.SetBurst(int(capacity.Burst))
	logger.Infof("[operator] capacity: rps=%.2f burst=%d", capacity.Rps, capacity.Burst)
}
