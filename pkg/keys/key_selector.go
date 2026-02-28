package keys

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// Strategy defines key selection strategies.
type Strategy string

const (
	StrategyHealthBasedFallback Strategy = "HealthBasedFallback"
	StrategyRoundRobin          Strategy = "RoundRobin"
)

// KeyHealth tracks health status of a sending key.
type KeyHealth struct {
	Address   common.Address
	Healthy   bool
	Reason    string
	UpdatedAt time.Time
}

// Selector picks a FromAddress for each write operation.
type Selector interface {
	// SelectKey returns the address to use for the next write.
	SelectKey(ctx context.Context) (common.Address, error)
	// MarkUnhealthy flags a key as temporarily unavailable.
	MarkUnhealthy(addr common.Address, reason string)
	// MarkHealthy restores a key to the available pool.
	MarkHealthy(addr common.Address)
	// ActiveAddresses returns all configured addresses.
	ActiveAddresses() []common.Address
	// HealthStatus returns health info for all keys.
	HealthStatus() []KeyHealth
}

// NewSelector creates a Selector from a list of addresses and a strategy.
// Falls back to single-key behavior if only one address is provided.
func NewSelector(lggr logger.Logger, addresses []common.Address, strategy Strategy) (Selector, error) {
	if len(addresses) == 0 {
		return nil, fmt.Errorf("at least one sending key address is required")
	}
	switch strategy {
	case StrategyRoundRobin:
		return newRoundRobinSelector(lggr, addresses), nil
	case StrategyHealthBasedFallback:
		return newHealthBasedSelector(lggr, addresses), nil
	default:
		return nil, fmt.Errorf("unknown key selection strategy: %s", strategy)
	}
}

// healthBasedSelector uses the primary key and falls back to secondary only when primary is unhealthy.
type healthBasedSelector struct {
	lggr      logger.Logger
	addresses []common.Address
	health    map[common.Address]*KeyHealth
	mu        sync.RWMutex
}

func newHealthBasedSelector(lggr logger.Logger, addresses []common.Address) *healthBasedSelector {
	health := make(map[common.Address]*KeyHealth, len(addresses))
	for _, addr := range addresses {
		health[addr] = &KeyHealth{Address: addr, Healthy: true, UpdatedAt: time.Now()}
	}
	return &healthBasedSelector{
		lggr:      logger.Named(lggr, "HealthBasedKeySelector"),
		addresses: addresses,
		health:    health,
	}
}

func (s *healthBasedSelector) SelectKey(_ context.Context) (common.Address, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Return first healthy key (primary preferred)
	for _, addr := range s.addresses {
		if h, ok := s.health[addr]; ok && h.Healthy {
			return addr, nil
		}
	}
	// All keys unhealthy — return primary anyway and let TxManager handle it
	s.lggr.Warnw("All sending keys are unhealthy, falling back to primary", "primary", s.addresses[0])
	return s.addresses[0], nil
}

func (s *healthBasedSelector) MarkUnhealthy(addr common.Address, reason string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h, ok := s.health[addr]; ok {
		h.Healthy = false
		h.Reason = reason
		h.UpdatedAt = time.Now()
		s.lggr.Warnw("Sending key marked unhealthy", "address", addr, "reason", reason)
	}
}

func (s *healthBasedSelector) MarkHealthy(addr common.Address) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h, ok := s.health[addr]; ok {
		if !h.Healthy {
			s.lggr.Infow("Sending key recovered", "address", addr, "previousReason", h.Reason)
		}
		h.Healthy = true
		h.Reason = ""
		h.UpdatedAt = time.Now()
	}
}

func (s *healthBasedSelector) ActiveAddresses() []common.Address {
	return s.addresses
}

func (s *healthBasedSelector) HealthStatus() []KeyHealth {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]KeyHealth, 0, len(s.addresses))
	for _, addr := range s.addresses {
		if h, ok := s.health[addr]; ok {
			result = append(result, *h)
		}
	}
	return result
}

// roundRobinSelector cycles through healthy keys.
type roundRobinSelector struct {
	lggr      logger.Logger
	addresses []common.Address
	health    map[common.Address]*KeyHealth
	index     uint64
	mu        sync.Mutex
}

func newRoundRobinSelector(lggr logger.Logger, addresses []common.Address) *roundRobinSelector {
	health := make(map[common.Address]*KeyHealth, len(addresses))
	for _, addr := range addresses {
		health[addr] = &KeyHealth{Address: addr, Healthy: true, UpdatedAt: time.Now()}
	}
	return &roundRobinSelector{
		lggr:      logger.Named(lggr, "RoundRobinKeySelector"),
		addresses: addresses,
		health:    health,
	}
}

func (s *roundRobinSelector) SelectKey(_ context.Context) (common.Address, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := uint64(len(s.addresses))
	// Try each address starting from current index
	for i := uint64(0); i < n; i++ {
		idx := (s.index + i) % n
		addr := s.addresses[idx]
		if h, ok := s.health[addr]; ok && h.Healthy {
			s.index = idx + 1
			return addr, nil
		}
	}
	// All unhealthy — use next in rotation anyway
	addr := s.addresses[s.index%n]
	s.index++
	s.lggr.Warnw("All sending keys are unhealthy, using next in rotation", "address", addr)
	return addr, nil
}

func (s *roundRobinSelector) MarkUnhealthy(addr common.Address, reason string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h, ok := s.health[addr]; ok {
		h.Healthy = false
		h.Reason = reason
		h.UpdatedAt = time.Now()
		s.lggr.Warnw("Sending key marked unhealthy", "address", addr, "reason", reason)
	}
}

func (s *roundRobinSelector) MarkHealthy(addr common.Address) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h, ok := s.health[addr]; ok {
		if !h.Healthy {
			s.lggr.Infow("Sending key recovered", "address", addr, "previousReason", h.Reason)
		}
		h.Healthy = true
		h.Reason = ""
		h.UpdatedAt = time.Now()
	}
}

func (s *roundRobinSelector) ActiveAddresses() []common.Address {
	return s.addresses
}

func (s *roundRobinSelector) HealthStatus() []KeyHealth {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]KeyHealth, 0, len(s.addresses))
	for _, addr := range s.addresses {
		if h, ok := s.health[addr]; ok {
			result = append(result, *h)
		}
	}
	return result
}
