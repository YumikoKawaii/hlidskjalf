package ratelimit

import (
	"context"
	"sync/atomic"

	"github.com/YumikoKawaii/shared/logger"
)

// Manager is the top-level coordinator for distributed rate limiting. It is
// created by the server when rate limiting is enabled and passed to handlers.
//
// On the leader pod, Allow() checks the global token bucket directly.
// On follower pods, Allow() checks the local Peer token bucket.
type Manager struct {
	isLeader atomic.Bool            // true if this pod currently holds the lease
	leader   atomic.Pointer[Leader] // non-nil only while leading
	peer     *Peer                  // local rate limiter for follower mode
	podIP    string                 // this pod's IP (election identity)
	leaderIP atomic.Value           // current leader's IP
	rps      float64                // configured global RPS
	burst    int                    // configured burst size
}

func NewManager(podIP string) *Manager {
	return &Manager{
		podIP: podIP,
	}
}

// Allow checks whether the current request should be permitted.
// Leader: checks the global token bucket directly (no network call).
// Follower: checks the local peer token bucket (refilled in background).
func (m *Manager) Allow() bool {
	if m.isLeader.Load() {
		if l := m.leader.Load(); l != nil {
			allowed := l.Allow()
			logger.Infof("[ratelimit] leader allow=%v tokens=%.1f", allowed, l.limiter.Tokens())
			return allowed
		}
		return true
	}
	if m.peer != nil {
		allowed := m.peer.Allow()
		logger.Infof("[ratelimit] peer allow=%v tokens=%.1f", allowed, m.peer.limiter.Tokens())
		return allowed
	}
	return true
}

// Start launches the leader election goroutine and the peer refill loop.
// On election win, it creates a Leader with the global budget.
// On election loss, it clears the Leader and reverts to peer-only mode.
// The peer refill loop runs regardless, requesting tokens from whichever
// pod currently holds the lease.
func (m *Manager) Start(ctx context.Context, serviceName, namespace string, rps float64, burst int, leaderPort int) {
	m.rps = rps
	m.burst = burst
	m.peer = NewPeer(rps, burst, m.podIP, leaderPort)

	go RunElection(ctx, serviceName, namespace, m.podIP, ElectionCallbacks{
		OnStartedLeading: func(ctx context.Context) {
			logger.Infof("[ratelimit] became leader (identity=%s)", m.podIP)
			l := NewLeader(rps, burst)
			m.leader.Store(l)
			m.isLeader.Store(true)
			m.leaderIP.Store(m.podIP)
			<-ctx.Done() // block until leadership is lost
		},
		OnStoppedLeading: func() {
			logger.Info("[ratelimit] stopped leading")
			m.isLeader.Store(false)
			m.leader.Store(nil)
		},
		OnNewLeader: func(identity string) {
			logger.Infof("[ratelimit] new leader: %s", identity)
			m.leaderIP.Store(identity)
			m.peer.SetLeader(identity)
		},
	})

	go m.peer.RunRefillLoop(ctx)
}

// GrantTokens delegates a token request to the Leader. Called by the gRPC handler.
// Returns zero grant if this instance is not the leader.
func (m *Manager) GrantTokens(peerIP string, want int) TokenGrant {
	if !m.isLeader.Load() {
		return TokenGrant{}
	}
	l := m.leader.Load()
	if l == nil {
		return TokenGrant{}
	}
	return l.GrantTokens(peerIP, want)
}

// UnregisterPeer removes a peer from the leader's allocation map immediately.
// Called by the gRPC handler when a peer sends a quantity=0 token request
// (graceful shutdown signal).
func (m *Manager) UnregisterPeer(peerIP string) {
	if !m.isLeader.Load() {
		return
	}
	l := m.leader.Load()
	if l == nil {
		return
	}
	l.Unregister(peerIP)
}

// Stop performs graceful shutdown: notifies the leader that this peer is leaving
// so it can immediately rebalance fair-share RPS across remaining peers.
func (m *Manager) Stop() {
	if m.peer != nil {
		m.peer.Unregister()
	}
}

// IsLeader returns whether this instance currently holds the lease.
func (m *Manager) IsLeader() bool {
	return m.isLeader.Load()
}

// GetStatus returns the current state for the debug/status endpoint.
func (m *Manager) GetStatus() Status {
	leaderIP := ""
	if v := m.leaderIP.Load(); v != nil {
		leaderIP, _ = v.(string)
	}

	resp := Status{
		IsLeader:  m.isLeader.Load(),
		LeaderIP:  leaderIP,
		PodIP:     m.podIP,
		GlobalRPS: m.rps,
	}

	//if resp.IsLeader {
	//	if l := m.leader.Load(); l != nil {
	//		resp.PeerAllocations = l.GetAllocations()
	//	}
	//}

	return resp
}
