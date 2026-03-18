package txm

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	MaxInFlightTransactions int = 16
	MaxInFlightSubset       int = 5

	broadcastInterval           time.Duration = 30 * time.Second
	maxAttemptsThreshold        uint16        = 10
	pendingNonceDefaultTimeout  time.Duration = 30 * time.Second
	pendingNonceRecheckInterval time.Duration = 1 * time.Second
)

type Client interface {
	PendingNonceAt(context.Context, common.Address) (uint64, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
}

type TxStore interface {
	AbandonPendingTransactions(context.Context, common.Address) error
	AppendAttemptToTransaction(context.Context, uint64, common.Address, *types.Attempt) (attempts []*types.Attempt, err error)
	CreateEmptyUnconfirmedTransaction(context.Context, common.Address, uint64, uint64) (*types.Transaction, error)
	CreateTransaction(context.Context, *types.TxRequest) (*types.Transaction, error)
	FetchUnconfirmedTransactions(context.Context, common.Address) ([]*types.Transaction, error)
	FetchUnconfirmedTransactionAtNonceWithCount(context.Context, uint64, common.Address) (*types.Transaction, int, error)
	MarkConfirmedAndReorgedTransactions(context.Context, uint64, common.Address) ([]*types.Transaction, []uint64, error)
	MarkUnconfirmedTransactionPurgeable(context.Context, uint64, common.Address) error
	UpdateSignedAttempt(context.Context, uint64, uint64, *evmtypes.Transaction, common.Address) error
	UpdateTransactionBroadcast(context.Context, uint64, uint64, common.Hash, common.Address) error
	UpdateUnstartedTransactionWithNonce(context.Context, common.Address, uint64) (*types.Transaction, error)

	// ErrorHandler
	DeleteAttemptForUnconfirmedTx(context.Context, uint64, *types.Attempt, common.Address) error
	MarkTxFatal(context.Context, *types.Transaction, common.Address) error
}

type AttemptBuilder interface {
	NewAgnosticBumpAttempt(ctx context.Context, lggr logger.Logger, tx *types.Transaction, dynamic bool) (*types.Attempt, error)
}

type ErrorHandler interface {
	HandleError(context.Context, *types.Transaction, error, TxStore, func(common.Address, uint64), bool) (err error)
}

type StuckTxDetector interface {
	DetectStuckTransaction(ctx context.Context, tx *types.Transaction) (bool, error)
}

type Keystore interface {
	EnabledAddressesForChain(ctx context.Context, chainID *big.Int) (addresses []common.Address, err error)
}

type Config struct {
	EIP1559             bool
	BlockTime           time.Duration
	RetryBlockThreshold uint16
	EmptyTxLimitDefault uint64
}

type Txm struct {
	services.StateMachine
	lggr            logger.SugaredLogger
	chainID         *big.Int
	client          Client
	attemptBuilder  AttemptBuilder
	errorHandler    ErrorHandler
	stuckTxDetector StuckTxDetector
	txStore         TxStore
	keystore        keys.AddressLister
	config          Config
	Metrics         *txmMetrics

	nonceMapMu sync.RWMutex
	nonceMap   map[common.Address]uint64

	triggerCh map[common.Address]chan struct{}
	stopCh    services.StopChan
	wg        sync.WaitGroup
}

func NewTxm(lggr logger.Logger, chainID *big.Int, client Client, attemptBuilder AttemptBuilder, txStore TxStore, stuckTxDetector StuckTxDetector, config Config, keystore keys.AddressLister, errorHandler ErrorHandler) *Txm {
	return &Txm{
		lggr:            logger.Sugared(logger.Named(lggr, "Txm")),
		keystore:        keystore,
		chainID:         chainID,
		client:          client,
		attemptBuilder:  attemptBuilder,
		txStore:         txStore,
		stuckTxDetector: stuckTxDetector,
		config:          config,
		errorHandler:    errorHandler,
		nonceMap:        make(map[common.Address]uint64),
		triggerCh:       make(map[common.Address]chan struct{}),
	}
}

func (t *Txm) Start(ctx context.Context) error {
	return t.StartOnce("Txm", func() error {
		tm, err := NewTxmMetrics(t.chainID)
		if err != nil {
			return err
		}
		t.Metrics = tm
		t.stopCh = make(chan struct{})

		addresses, err := t.keystore.EnabledAddresses(ctx)
		if err != nil {
			return err
		}
		for _, address := range addresses {
			t.startAddress(address)
		}
		t.lggr.Infof("Started Txm")
		return nil
	})
}

func (t *Txm) startAddress(address common.Address) {
	triggerCh := make(chan struct{}, 1)
	t.triggerCh[address] = triggerCh

	t.wg.Add(1)
	go t.loop(address, triggerCh)
}

func (t *Txm) initializeNonce(ctx context.Context, address common.Address) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, pendingNonceDefaultTimeout)
	defer cancel()
	for {
		pendingNonce, err := t.client.PendingNonceAt(ctxWithTimeout, address)
		if err != nil {
			t.lggr.Errorw("Error when fetching initial nonce", "address", address, "err", err)
			select {
			case <-time.After(pendingNonceRecheckInterval):
			case <-ctx.Done():
				t.lggr.Errorw("context error", "err", context.Cause(ctx))
				return
			}
			continue
		}
		t.SetNonce(address, pendingNonce)
		t.lggr.Debugf("Set initial nonce for address: %v to %d", address, pendingNonce)
		return
	}
}

func (t *Txm) Close() error {
	return t.StopOnce("Txm", func() error {
		close(t.stopCh)
		t.wg.Wait()
		return nil
	})
}

func (t *Txm) HealthReport() map[string]error {
	return map[string]error{t.lggr.Name(): t.Healthy()}
}

func (t *Txm) CreateTransaction(ctx context.Context, txRequest *types.TxRequest) (tx *types.Transaction, err error) {
	tx, err = t.txStore.CreateTransaction(ctx, txRequest)
	if err == nil {
		t.lggr.Infow("Created transaction", "tx", tx)
	}
	return
}

func (t *Txm) Trigger(address common.Address) {
	if !t.IfStarted(func() {
		triggerCh, exists := t.triggerCh[address]
		if !exists {
			return
		}
		triggerCh <- struct{}{}
	}) {
		t.lggr.Error("Txm unstarted")
	}
}

func (t *Txm) Abandon(address common.Address) error {
	// TODO: restart txm
	t.lggr.Infof("Dropping unstarted and unconfirmed transactions for address: %v", address)
	return t.txStore.AbandonPendingTransactions(context.TODO(), address)
}

func (t *Txm) GetNonce(address common.Address) uint64 {
	t.nonceMapMu.RLock()
	defer t.nonceMapMu.RUnlock()
	return t.nonceMap[address]
}

func (t *Txm) SetNonce(address common.Address, nonce uint64) {
	t.nonceMapMu.Lock()
	defer t.nonceMapMu.Unlock()
	t.nonceMap[address] = nonce
}

func newBackoff(minDuration time.Duration) backoff.Backoff {
	return backoff.Backoff{
		Min:    minDuration,
		Max:    1 * time.Minute,
		Jitter: true,
	}
}

func (t *Txm) loop(address common.Address, triggerCh chan struct{}) {
	defer t.wg.Done()
	ctx, cancel := t.stopCh.NewCtx()
	defer cancel()
	broadcastWithBackoff := newBackoff(1 * time.Second)
	var broadcastCh <-chan time.Time
	backfillTicker := services.TickerConfig{Initial: t.config.BlockTime, JitterPct: services.DefaultJitter}.NewTicker(t.config.BlockTime)
	defer backfillTicker.Stop()

	t.initializeNonce(ctx, address)

	for {
		start := time.Now()
		bo, err := t.BroadcastTransaction(ctx, address)
		if err != nil {
			t.lggr.Errorw("Error during transaction broadcasting", "err", err)
		} else {
			t.lggr.Debug("Transaction broadcasting time elapsed: ", time.Since(start))
		}
		if bo {
			broadcastCh = time.After(broadcastWithBackoff.Duration())
		} else {
			broadcastWithBackoff.Reset()
			broadcastCh = time.After(utils.WithJitter(broadcastInterval))
		}
		select {
		case <-ctx.Done():
			return
		case <-triggerCh:
			continue
		case <-broadcastCh:
			continue
		case <-backfillTicker.C:
			start := time.Now()
			err := t.BackfillTransactions(ctx, address)
			if err != nil {
				t.lggr.Errorw("Error during backfill", "err", err)
			} else {
				t.lggr.Debug("Backfill time elapsed: ", time.Since(start))
			}
		}
	}
}

func (t *Txm) BroadcastTransaction(ctx context.Context, address common.Address) (bool, error) {
	for {
		_, unconfirmedCount, err := t.txStore.FetchUnconfirmedTransactionAtNonceWithCount(ctx, 0, address)
		if err != nil {
			return false, err
		}

		// Optimistically send up to MaxInFlightSubset of the MaxInFlightTransactions. After that threshold, broadcast more cautiously
		// by checking the pending nonce so no more than MaxInFlightSubset can get stuck simultaneously i.e. due
		// to insufficient balance. We're making this trade-off to avoid storing stuck transactions and making unnecessary
		// RPC calls. The upper limit is always MaxInFlightTransactions regardless of the pending nonce.
		if unconfirmedCount >= MaxInFlightSubset {
			if unconfirmedCount > MaxInFlightTransactions {
				t.lggr.Warnf("Reached transaction limit: %d for unconfirmed transactions", MaxInFlightTransactions)
				return true, nil
			}
			pendingNonce, e := t.client.PendingNonceAt(ctx, address)
			if e != nil {
				return false, e
			}
			nonce := t.GetNonce(address)
			if nonce > pendingNonce {
				t.lggr.Warnf("Reached transaction limit. LocalNonce: %d, PendingNonce %d, unconfirmedCount: %d",
					nonce, pendingNonce, unconfirmedCount)
				return true, nil
			}
		}

		nonce := t.GetNonce(address)
		tx, err := t.txStore.UpdateUnstartedTransactionWithNonce(ctx, address, nonce)
		if err != nil {
			return false, err
		}
		if tx == nil {
			return false, nil
		}
		t.SetNonce(address, nonce+1)

		if err := t.createAndSendAttempt(ctx, tx, address); err != nil {
			return false, err
		}
	}
}

func (t *Txm) createAndSendAttempt(ctx context.Context, tx *types.Transaction, address common.Address) error {
	attempt, err := t.attemptBuilder.NewAgnosticBumpAttempt(ctx, t.lggr, tx, t.config.EIP1559)
	if err != nil {
		return err
	}

	if tx.Nonce == nil {
		return fmt.Errorf("nonce for txID: %v is empty", tx.ID)
	}
	attempts, err := t.txStore.AppendAttemptToTransaction(ctx, *tx.Nonce, address, attempt)
	if err != nil {
		return err
	}
	tx.AttemptCount++
	tx.Attempts = attempts

	return t.sendTransactionWithError(ctx, tx, attempt, address)
}

func (t *Txm) sendTransactionWithError(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, fromAddress common.Address) (err error) {
	if tx.Nonce == nil {
		return fmt.Errorf("nonce for txID: %v is empty", tx.ID)
	}
	start := time.Now()
	txErr := t.client.SendTransaction(ctx, tx, attempt)
	t.lggr.Infow("Broadcasted attempt", "tx", tx, "attempt", attempt, "duration", time.Since(start), "txErr: ", txErr)
	if txErr != nil && t.errorHandler != nil {
		if err = t.errorHandler.HandleError(ctx, tx, txErr, t.txStore, t.SetNonce, false); err != nil {
			return
		}
	} else if txErr != nil {
		pendingNonce, pErr := t.client.PendingNonceAt(ctx, fromAddress)
		if pErr != nil {
			return pErr
		}
		if pendingNonce <= *tx.Nonce {
			return fmt.Errorf("pending nonce for txID: %v didn't increase. PendingNonce: %d, TxNonce: %d. TxErr: %w", tx.ID, pendingNonce, *tx.Nonce, txErr)
		}
	}

	t.Metrics.IncrementNumBroadcastedTxs(ctx)
	if err = t.Metrics.EmitTxMessage(ctx, attempt.Hash, fromAddress, tx); err != nil {
		t.lggr.Errorw("Beholder error emitting tx message", "err", err)
	}

	return t.txStore.UpdateTransactionBroadcast(ctx, attempt.TxID, *tx.Nonce, attempt.Hash, fromAddress)
}

func (t *Txm) BackfillTransactions(ctx context.Context, address common.Address) error {
	latestNonce, err := t.client.NonceAt(ctx, address, nil)
	if err != nil {
		return err
	}

	confirmedTransactions, unconfirmedTransactionIDs, err := t.txStore.MarkConfirmedAndReorgedTransactions(ctx, latestNonce, address)
	if err != nil {
		return err
	}
	if len(confirmedTransactions) > 0 || len(unconfirmedTransactionIDs) > 0 {
		t.Metrics.IncrementNumConfirmedTxs(ctx, len(confirmedTransactions))
		confirmedTransactionIDs := t.extractMetrics(ctx, confirmedTransactions)
		t.lggr.Infof("Confirmed transaction IDs: %v . Re-orged transaction IDs: %v", confirmedTransactionIDs, unconfirmedTransactionIDs)
	}

	tx, unconfirmedCount, err := t.txStore.FetchUnconfirmedTransactionAtNonceWithCount(ctx, latestNonce, address)
	if err != nil {
		return err
	}
	if unconfirmedCount == 0 {
		t.lggr.Debugf("All transactions confirmed for address: %v", address)
		t.Metrics.ReachedMaxAttempts(ctx, false)
		return nil
	}

	if tx == nil || *tx.Nonce != latestNonce {
		t.lggr.Warnf("Nonce gap at nonce: %d - address: %v. Creating a new transaction\n", latestNonce, address)
		t.Metrics.IncrementNumNonceGaps(ctx)
		return t.createAndSendEmptyTx(ctx, latestNonce, address)
	} else { //nolint:revive //easier to read
		if !tx.IsPurgeable && t.stuckTxDetector != nil {
			isStuck, err := t.stuckTxDetector.DetectStuckTransaction(ctx, tx)
			if err != nil {
				return err
			}
			if isStuck {
				tx.IsPurgeable = true
				err = t.txStore.MarkUnconfirmedTransactionPurgeable(ctx, *tx.Nonce, address)
				if err != nil {
					return err
				}
				t.lggr.Infof("Marked tx as purgeable. Sending purge attempt for txID: %d", tx.ID)
				return t.createAndSendAttempt(ctx, tx, address)
			}
		}

		if tx.AttemptCount >= maxAttemptsThreshold {
			t.Metrics.ReachedMaxAttempts(ctx, true)
			t.lggr.Warnf("Reached max attempts threshold for txID: %d. TXM will broadcast more attempts  but if this"+
				" error persists, it means the transaction won't likely be confirmed and there is an issue with the transaction."+
				"Look for any error messages from previous broadcasted attempts that may indicate why this happened, i.e. wallet is out of funds. Tx: %v", tx.ID,
				tx.PrintWithAttempts())
		} else {
			t.Metrics.ReachedMaxAttempts(ctx, false)
		}

		// Rebroadcast if at least one of the following conditions is met:
		// - The transaction has never been broadcasted successfully before
		// - The last broadcast was more than RetryBlockThreshold blocks ago
		// - The transaction is purgeable
		if tx.LastBroadcastAt == nil || time.Since(*tx.LastBroadcastAt) > (t.config.BlockTime*time.Duration(t.config.RetryBlockThreshold)) || tx.IsPurgeable {
			t.lggr.Info("Rebroadcasting attempt for txID: ", tx.ID)
			return t.createAndSendAttempt(ctx, tx, address)
		}
	}
	return nil
}

func (t *Txm) createAndSendEmptyTx(ctx context.Context, latestNonce uint64, address common.Address) error {
	tx, err := t.txStore.CreateEmptyUnconfirmedTransaction(ctx, address, latestNonce, t.config.EmptyTxLimitDefault)
	if err != nil {
		return err
	}
	return t.createAndSendAttempt(ctx, tx, address)
}

func (t *Txm) extractMetrics(ctx context.Context, txs []*types.Transaction) []uint64 {
	confirmedTxIDs := make([]uint64, 0, len(txs))
	for _, tx := range txs {
		confirmedTxIDs = append(confirmedTxIDs, tx.ID)
		if tx.InitialBroadcastAt != nil {
			t.Metrics.RecordTimeUntilTxConfirmed(ctx, float64(time.Since(*tx.InitialBroadcastAt)))
		}
	}
	return confirmedTxIDs
}
