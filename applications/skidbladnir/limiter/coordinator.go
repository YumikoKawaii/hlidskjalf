package limiter

import (
	"sync"
	"time"
)

// Coordinator is the leader-side capacity allocator. Only the pod that holds the
// Kubernetes Lease creates a Coordinator. It does not handle traffic itself —
// it purely tracks registered operators and computes fair-share allocations.
//
// Each Operator (including the one on the leader pod) periodically calls
// Capacity(), which doubles as a heartbeat. Operators that stop calling
// (crash or scale-down) are evicted after the configured TTL.
type Coordinator struct {
	mu          sync.Mutex
	globalRPS   float64              // configured service-wide RPS cap
	globalBurst int                  // configured service-wide burst cap
	ttl         time.Duration        // evict operators that haven't checked in within this window
	operators   map[string]time.Time // operatorIP -> last seen timestamp
}

func NewCoordinator(rps float64, burst int, ttl time.Duration) *Coordinator {
	return &Coordinator{
		globalRPS:   rps,
		globalBurst: burst,
		ttl:         ttl,
		operators:   make(map[string]time.Time),
	}
}

// Capacity computes the fair-share allocation for a requesting operator.
// It doubles as a heartbeat: the operator's last-seen timestamp is updated,
// and stale operators (no heartbeat within TTL) are evicted.
//
// The global budget is split evenly: rps / operatorCount. Every pod
// (including the leader) runs an Operator that registers here, so there
// is no special +1 accounting for the leader.
func (c *Coordinator) Capacity(operatorIP string) Capacity {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	c.operators[operatorIP] = now

	// Evict stale operators that haven't checked in within TTL
	for ip, lastSeen := range c.operators {
		if now.Sub(lastSeen) > c.ttl {
			delete(c.operators, ip)
		}
	}

	totalOperators := len(c.operators)
	if totalOperators == 0 {
		totalOperators = 1
	}

	rps := c.globalRPS / float64(totalOperators)
	burst := int(float64(c.globalBurst) / float64(totalOperators))
	if burst < 1 {
		burst = 1
	}

	return Capacity{
		Burst: int32(burst),
		RPS:   rps,
	}
}

// Unregister removes an operator immediately from the tracking map.
// Called during graceful shutdown so the coordinator can rebalance
// remaining operators without waiting for TTL expiry.
func (c *Coordinator) Unregister(operatorIP string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.operators, operatorIP)
}
