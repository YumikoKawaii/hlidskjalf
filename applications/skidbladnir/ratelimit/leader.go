package ratelimit

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const peerTTL = 5 * time.Second // evict peers that haven't requested tokens within this window

// Leader holds the global rate limiter and tracks per-peer allocations.
// Only the elected leader instance creates a Leader; it processes token
// requests from followers and enforces the service-wide rate limit.
// Peers that stop sending refill requests (due to scale-down or crash)
// are automatically evicted after peerTTL.
type Leader struct {
	mu        sync.Mutex
	limiter   *rate.Limiter         // global token bucket shared across all pods
	globalRPS float64               // configured service-wide RPS cap
	peers     map[string]time.Time  // peerIP -> last seen timestamp
}

func NewLeader(rps float64, burst int) *Leader {
	return &Leader{
		limiter:   rate.NewLimiter(rate.Limit(rps), burst),
		globalRPS: rps,
		peers:     make(map[string]time.Time),
	}
}

// GrantTokens reserves tokens from the global pool for a requesting peer.
// It evicts stale peers, recalculates fair-share RPS, and drains tokens.
// The leader itself is counted as a participant (+1) since it also consumes
// from the global pool directly.
// Peer state is not explicitly migrated on leader change — the new leader
// rebuilds its peer map as followers send their next refill request (~500ms).
func (l *Leader) GrantTokens(peerIP string, want int) TokenGrant {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	// Update last-seen timestamp for requesting peer
	l.peers[peerIP] = now

	// Evict stale peers that haven't checked in within TTL
	for ip, lastSeen := range l.peers {
		if now.Sub(lastSeen) > peerTTL {
			delete(l.peers, ip)
		}
	}

	// +1 accounts for the leader pod itself
	totalParticipants := len(l.peers) + 1
	perPeerRPS := l.globalRPS / float64(totalParticipants)

	// Drain tokens from the global bucket up to the requested amount
	granted := 0
	for i := 0; i < want; i++ {
		if l.limiter.Allow() {
			granted++
		} else {
			break
		}
	}

	return TokenGrant{
		Granted: granted,
		RPS:     perPeerRPS,
	}
}

// Unregister removes a peer immediately from the allocation map.
// Called during graceful shutdown so the leader can rebalance the remaining
// peers' fair-share RPS without waiting for TTL expiry.
func (l *Leader) Unregister(peerIP string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.peers, peerIP)
}

// Allow checks the global limiter directly. Used by the leader pod
// for its own inbound requests (no need to talk to itself over gRPC).
func (l *Leader) Allow() bool {
	return l.limiter.Allow()
}
