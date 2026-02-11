package discovery

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Service struct {
	name string
	mu   sync.RWMutex

	ips   []string
	ports map[string]int32 // port name → port number
	rrIdx atomic.Uint64

	cancel context.CancelFunc
}

func NewService(name string, cancel context.CancelFunc) *Service {
	return &Service{
		name:   name,
		cancel: cancel,
	}
}

// GetTarget returns an address (ip:port) using round-robin for the given port name.
// Returns ("", false) if no endpoints are available or the port doesn't exist.
func (s *Service) GetTarget(portName string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	port, ok := s.ports[portName]
	if !ok || len(s.ips) == 0 {
		return "", false
	}

	idx := s.rrIdx.Add(1) - 1
	ip := s.ips[idx%uint64(len(s.ips))]
	return fmt.Sprintf("%s:%d", ip, port), true
}

func (s *Service) Cancel() {
	s.cancel()
}

func (s *Service) UpdateEndpoints(ips []string, ports map[string]int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ips = ips
	s.ports = ports
}
