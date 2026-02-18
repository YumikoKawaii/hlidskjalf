package discovery

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrNoEndpoints = errors.New("no endpoints available")

type Service struct {
	mu sync.RWMutex

	ips   []string
	rrIdx atomic.Uint64
}

func NewService() *Service {
	return &Service{}
}

// GetTarget returns an address (ip:port) using round-robin.
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

func (s *Service) UpdateEndpoints(ips []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ips = ips
}
