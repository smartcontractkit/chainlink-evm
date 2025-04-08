package keys

import (
	"errors"
	"sync"
)

type Mutex struct {
	mu          sync.Mutex
	serviceType ServiceType
	count       int // Tracks active users per service type
}
type ServiceType int

const (
	TXMv1 ServiceType = iota
	TXMv2
)

// TryLock attempts to lock the resource for the specified service type.
// It returns an error if the resource is locked by a different service type.
func (m *Mutex) TryLock(serviceType ServiceType) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.count == 0 {
		m.serviceType = serviceType
	}

	// Check if other service types are using the resource
	if m.serviceType != serviceType && m.count > 0 {
		return errors.New("resource is locked by another service type")
	}

	// Increment active count for the current service type
	m.count++
	return nil
}

// Unlock releases the lock for the service type
func (m *Mutex) Unlock(serviceType ServiceType) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the service type has an active lock
	if m.count == 0 {
		return errors.New("no active lock")
	}

	if m.serviceType != serviceType {
		return errors.New("no active lock for this service type")
	}

	// Decrement active count for the service type
	m.count--
	return nil
}

// IsLocked checks if the resource is locked by a specific service type.
func (m *Mutex) IsLocked(serviceType ServiceType) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.count == 0 || m.serviceType != serviceType {
		return false
	}

	return true
}
