package discovery

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Service struct {
	mu sync.RWMutex

	ips   []string
	port  int
	rrIdx atomic.Uint64

	cancel context.CancelFunc
}

func NewService(port int, cancel context.CancelFunc) *Service {
	return &Service{
		port:   port,
		cancel: cancel,
	}
}

// GetTarget returns an address (ip:port) using round-robin.
func (s *Service) GetTarget() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	idx := s.rrIdx.Add(1) - 1
	ip := s.ips[idx%uint64(len(s.ips))]
	return fmt.Sprintf("%s:%d", ip, s.port)
}

func (s *Service) Cancel() {
	s.cancel()
}

func (s *Service) UpdateEndpoints(ips []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ips = ips
}
