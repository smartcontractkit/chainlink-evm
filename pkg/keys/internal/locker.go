package internal

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type Locker[M any] struct {
	mu  sync.Mutex
	mus map[common.Address]*M
}

func (l *Locker[M]) GetMutex(address common.Address) *M {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.mus == nil {
		l.mus = make(map[common.Address]*M)
	}

	mu, exists := l.mus[address]
	if !exists {
		mu = new(M)
		l.mus[address] = mu
	}
	return mu
}
