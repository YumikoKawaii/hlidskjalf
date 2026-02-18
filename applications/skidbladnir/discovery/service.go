package discovery

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrNoEndpoints = errors.New("no endpoints available")

// Service tracks the set of pod IPs for a single upstream service
// and round-robins across them using an atomic counter.
type Service struct {
	mu sync.RWMutex

	ips   []string
	rrIdx atomic.Uint64
}

func NewService() *Service {
	return &Service{}
}

// GetTarget returns a "podIP:port" address using round-robin selection.
// Returns ErrNoEndpoints if no pod IPs have been discovered yet.
func (s *Service) GetTarget(port string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.ips) == 0 {
		return "", ErrNoEndpoints
	}

	idx := s.rrIdx.Add(1) - 1
	ip := s.ips[idx%uint64(len(s.ips))]
	return fmt.Sprintf("%s:%s", ip, port), nil
}

// UpdateEndpoints replaces the current pod IP list. Called by the Resolver
// when the Kubernetes Endpoints watch detects changes (scale up/down, pod restarts).
func (s *Service) UpdateEndpoints(ips []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ips = ips
}
