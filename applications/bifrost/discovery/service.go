package discovery

import (
	"context"
	"sync"
	"sync/atomic"
)

type Service struct {
	name string
	mu   sync.RWMutex

	ips   []string
	rrIdx atomic.Uint64

	cancel context.CancelFunc
}

func NewService(name string, cancel context.CancelFunc) *Service {
	return &Service{
		name:   name,
		cancel: cancel,
	}
}

// GetTarget returns an IP using round-robin. Returns ("", false) if no endpoints are available.
func (s *Service) GetTarget() (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.ips) == 0 {
		return "", false
	}

	idx := s.rrIdx.Add(1) - 1
	return s.ips[idx%uint64(len(s.ips))], true
}

func (s *Service) Cancel() {
	s.cancel()
}

func (s *Service) UpdateIPs(ips []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ips = ips
}
