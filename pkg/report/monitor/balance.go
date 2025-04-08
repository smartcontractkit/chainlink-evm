package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"
)

// TODO: Move this package to a more appropriate location
// ChainInfo contains the chain information (used as execution context)
type ChainInfo struct {
	ChainFamilyName string
	ChainID         string
	NetworkName     string
	NetworkNameFull string
}

// Config defines the balance monitor configuration.
type Config struct {
	BalancePollPeriod config.Duration
}

// BalanceClient defines the interface for getting account balances.
type BalanceClient interface {
	GetAccountBalance(addr string) (float64, error)
}

// BalanceMonitorOpts contains the options for creating a new balance monitor.
type BalanceMonitorOpts struct {
	ChainInfo           ChainInfo
	ChainNativeCurrency string

	Config           Config
	Logger           logger.Logger
	Keystore         core.Keystore
	NewBalanceClient func() (BalanceClient, error)

	// Maps a public key to an account address (optional, can return key as is)
	KeyToAccountMapper func(context.Context, string) (string, error)
}

// NewBalanceMonitor returns a balance monitoring services.Service which reports the balance of all Keystore accounts.
func NewBalanceMonitor(opts BalanceMonitorOpts) (services.Service, error) {
	return newBalanceMonitor(opts)
}

func newBalanceMonitor(opts BalanceMonitorOpts) (*balanceMonitor, error) {
	// Try to create a new gauge for account balance
	gauge, err := NewGaugeAccBalance(opts.ChainNativeCurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to create gauge: %w", err)
	}

	lggr := logger.Named(opts.Logger, "BalanceMonitor")
	return &balanceMonitor{
		cfg:  opts.Config,
		lggr: lggr,
		ks:   opts.Keystore,

		newReader:          opts.NewBalanceClient,
		keyToAccountMapper: opts.KeyToAccountMapper,
		updateFn: func(ctx context.Context, acc string, balance float64) {
			lggr.Infow("Account balance updated", "unit", opts.ChainNativeCurrency, "account", acc, "balance", balance)
			gauge.Record(ctx, balance, acc, opts.ChainInfo)
		},

		stop: make(chan struct{}),
		done: make(chan struct{}),
	}, nil
}

type balanceMonitor struct {
	services.StateMachine
	cfg  Config
	lggr logger.Logger
	ks   core.Keystore

	// Returns a new BalanceClient
	newReader func() (BalanceClient, error)
	// Maps a public key to an account address (optional, can return key as is)
	keyToAccountMapper func(context.Context, string) (string, error)
	// Updates the balance metric
	updateFn func(ctx context.Context, acc string, balance float64) // overridable for testing

	// Cached instance, intermitently reset to nil.
	reader BalanceClient

	stop services.StopChan
	done chan struct{}
}

func (m *balanceMonitor) Name() string {
	return m.lggr.Name()
}

func (m *balanceMonitor) Start(context.Context) error {
	return m.StartOnce(m.Name(), func() error {
		go m.start()
		return nil
	})
}

func (m *balanceMonitor) Close() error {
	return m.StopOnce(m.Name(), func() error {
		close(m.stop)
		<-m.done
		return nil
	})
}

func (m *balanceMonitor) HealthReport() map[string]error {
	return map[string]error{m.Name(): m.Healthy()}
}

// monitor fn continuously updates balances, until stop signal is received.
func (m *balanceMonitor) start() {
	defer close(m.done)
	ctx, cancel := m.stop.NewCtx()
	defer cancel()

	period := m.cfg.BalancePollPeriod.Duration()
	tick := time.After(utils.WithJitter(period))
	for {
		select {
		case <-m.stop:
			return
		case <-tick:
			m.updateBalances(ctx)
			tick = time.After(utils.WithJitter(period))
		}
	}
}

// getReader returns the stored BalanceClient, creating a new one if necessary.
func (m *balanceMonitor) getReader() (BalanceClient, error) {
	if m.reader == nil {
		var err error
		m.reader, err = m.newReader()
		if err != nil {
			return nil, err
		}
	}
	return m.reader, nil
}

// updateBalances updates the balances of all accounts in the keystore, using the provided BalanceClient and the updateFn.
func (m *balanceMonitor) updateBalances(ctx context.Context) {
	m.lggr.Debug("Updating account balances")
	keys, err := m.ks.Accounts(ctx)
	if err != nil {
		m.lggr.Errorw("Failed to get keys", "err", err)
		return
	}
	if len(keys) == 0 {
		return
	}
	reader, err := m.getReader()
	if err != nil {
		m.lggr.Errorw("Failed to get client", "err", err)
		return
	}

	var gotSomeBals bool
	for _, pk := range keys {
		// Check for shutdown signal, since Balance blocks and may be slow.
		select {
		case <-m.stop:
			return
		default:
		}

		// Account address can always be derived from the public key currently
		// TODO: if we need to support key rotation, the keystore should store the address explicitly
		// Notice: this is chain-specific key to account mapping injected (e.g., relevant for Aptos key management)
		accAddr, err := m.keyToAccountMapper(ctx, pk)
		if err != nil {
			m.lggr.Errorw("Failed to convert public key to account address", "err", err)
			continue
		}

		balance, err := reader.GetAccountBalance(accAddr)
		if err != nil {
			m.lggr.Errorw("Failed to get balance", "account", accAddr, "err", err)
			continue
		}
		gotSomeBals = true
		m.updateFn(ctx, accAddr, balance)
	}

	// Try a new client next time. // TODO: This is for multinode
	if !gotSomeBals {
		m.reader = nil
	}
}
