package ratelimit

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	api "github.com/YumikoKawaii/rpc.com/protobuf/rate_limiter"
	"github.com/YumikoKawaii/shared/logger"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Peer is the follower-side rate limiter. Each non-leader Skidbladnir instance
// uses a Peer to enforce its local share of the global rate limit. The local
// token bucket is checked synchronously on every proxied request (no network hop).
// A background goroutine periodically requests refills from the leader via gRPC
// when the local bucket drops below 20% capacity.
type Peer struct {
	limiter    *rate.Limiter // local token bucket with sub-budget of the global limit
	podIP      string        // this pod's IP, used as identity in token requests
	leaderIP   atomic.Value  // current leader IP, updated on election changes
	burst      int           // max burst size (used to calculate refill threshold)
	leaderPort int           // gRPC port of the leader's shared server
}

func NewPeer(rps float64, burst int, podIP string, leaderPort int) *Peer {
	return &Peer{
		limiter:    rate.NewLimiter(rate.Limit(rps), burst),
		podIP:      podIP,
		burst:      burst,
		leaderPort: leaderPort,
	}
}

// Allow performs a pure in-memory token bucket check. Called on every proxied request.
func (p *Peer) Allow() bool {
	return p.limiter.Allow()
}

// SetLeader updates the leader IP atomically. Called by Manager when election state changes.
func (p *Peer) SetLeader(ip string) {
	p.leaderIP.Store(ip)
}

// RunRefillLoop checks token levels every 500ms and requests refills from the
// leader via gRPC when tokens drop below 20% of burst capacity.
func (p *Peer) RunRefillLoop(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.maybeRefill(ctx)
		}
	}
}

// Unregister sends a quantity=0 token request to the leader, signaling that
// this peer is shutting down. The leader removes the peer from its allocation
// map immediately, allowing it to rebalance fair-share RPS across the
// remaining peers without waiting for TTL expiry.
func (p *Peer) Unregister() {
	leaderVal := p.leaderIP.Load()
	if leaderVal == nil {
		return
	}
	leaderIP, ok := leaderVal.(string)
	if !ok || leaderIP == "" || leaderIP == p.podIP {
		return // no leader known, or we are the leader
	}

	target := fmt.Sprintf("%s:%d", leaderIP, p.leaderPort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("[ratelimit] unregister: failed to connect to leader %s: %v", target, err)
		return
	}
	defer conn.Close()

	client := api.NewRateLimiterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = client.ServeTokens(ctx, &api.ServeTokensRequest{
		Ip:       p.podIP,
		Quantity: 0, // signal: unregister
	})
	if err != nil {
		logger.Errorf("[ratelimit] unregister: failed to notify leader %s: %v", target, err)
		return
	}
	logger.Info("[ratelimit] unregistered from leader")
}

// maybeRefill requests tokens from the leader if the local bucket is below 20% capacity.
func (p *Peer) maybeRefill(ctx context.Context) {
	tokens := p.limiter.Tokens()
	capacity := float64(p.burst)
	if capacity == 0 || tokens/capacity > 0.2 {
		return // still have enough tokens, skip refill
	}

	leaderVal := p.leaderIP.Load()
	if leaderVal == nil {
		return
	}
	leaderIP, ok := leaderVal.(string)
	if !ok || leaderIP == "" {
		return
	}

	// Request enough tokens to fill back to burst capacity
	want := p.burst - int(tokens)
	if want <= 0 {
		want = p.burst
	}

	target := fmt.Sprintf("%s:%d", leaderIP, p.leaderPort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("[ratelimit] failed to connect to leader %s: %v", target, err)
		return
	}
	defer conn.Close()

	client := api.NewRateLimiterClient(conn)
	reqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := client.ServeTokens(reqCtx, &api.ServeTokensRequest{
		Ip:       p.podIP,
		Quantity: int32(want),
	})
	if err != nil {
		logger.Errorf("[ratelimit] failed to request tokens from leader %s: %v", target, err)
		return
	}

	// Update local limiter RPS to match the leader's fair-share allocation
	p.limiter.SetLimit(rate.Limit(resp.Rps))
	logger.Infof("[ratelimit] refill: granted=%d rps=%.2f", resp.Quantity, resp.Rps)
}
