package txm

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const maxQueuedTransactions = 250

type Broadcaster struct {
	services.StateMachine
	lggr           logger.SugaredLogger
	chainID        *big.Int
	client         Client
	attemptBuilder AttemptBuilder
	keystore       keys.AddressLister
	config         Config

	transactionQueueMu sync.RWMutex
	transactionQueue   map[common.Address][]*types.Transaction

	triggerCh chan struct{}
	stopCh    services.StopChan
	wg        sync.WaitGroup
}

func NewBroadcaster(lggr logger.Logger, chainID *big.Int, client Client, attemptBuilder AttemptBuilder, config Config, keystore keys.AddressLister) *Broadcaster {
	return &Broadcaster{
		lggr:             logger.Sugared(logger.Named(lggr, "Broadcaster")),
		keystore:         keystore,
		chainID:          chainID,
		client:           client,
		attemptBuilder:   attemptBuilder,
		config:           config,
		transactionQueue: make(map[common.Address][]*types.Transaction),
		triggerCh:        make(chan struct{}),
	}
}

func (b *Broadcaster) Start(ctx context.Context) error {
	return b.StartOnce("Broadcaster", func() error {
		b.stopCh = make(chan struct{})

		b.wg.Add(1)
		go b.broadcastLoop()
		return nil
	})
}

func (b *Broadcaster) Close() error {
	return b.StopOnce("Broadcaster", func() error {
		close(b.stopCh)
		b.wg.Wait()
		return nil
	})
}

func (b *Broadcaster) HealthReport() map[string]error {
	return map[string]error{b.lggr.Name(): b.Healthy()}
}

func (b *Broadcaster) Trigger(address common.Address) {
	if !b.IfStarted(func() {
		b.triggerCh <- struct{}{}
	}) {
		b.lggr.Error("Broadcaster unstarted")
	}
}

func (b *Broadcaster) CreateTransaction(ctx context.Context, txRequest *types.TxRequest) (*types.Transaction, error) {
	b.transactionQueueMu.Lock()
	defer b.transactionQueueMu.Unlock()

	tx := &types.Transaction{
		// TODO: add ID
		//ID:                m.txIDCount,
		IdempotencyKey:    txRequest.IdempotencyKey,
		ChainID:           b.chainID,
		FromAddress:       txRequest.FromAddress,
		ToAddress:         txRequest.ToAddress,
		Value:             txRequest.Value,
		Data:              txRequest.Data,
		SpecifiedGasLimit: txRequest.SpecifiedGasLimit,
		CreatedAt:         time.Now(),
		Meta:              txRequest.Meta,
	}

	fromAddress := txRequest.FromAddress
	uLen := len(b.transactionQueue[fromAddress])
	if uLen >= maxQueuedTransactions {
		b.lggr.Warnw(fmt.Sprintf("Transaction queue for address: %v reached max limit of: %d. Dropping oldest transactions", fromAddress, maxQueuedTransactions),
			"txs", b.transactionQueue[fromAddress][0:uLen-maxQueuedTransactions+1]) // need to make room for the new tx
		for i := range b.transactionQueue[fromAddress][0 : uLen-maxQueuedTransactions+1] {
			b.transactionQueue[fromAddress][i] = nil // avoid memory leaks
		}
		b.transactionQueue[fromAddress] = b.transactionQueue[fromAddress][uLen-maxQueuedTransactions+1:]
	}

	b.transactionQueue[txRequest.FromAddress] = append(b.transactionQueue[txRequest.FromAddress], tx)
	return tx, nil
}

func (b *Broadcaster) broadcastLoop() {
	defer b.wg.Done()
	ctx, cancel := b.stopCh.NewCtx()
	defer cancel()

	ticker := services.NewTicker(b.config.BlockTime)
	defer ticker.Stop()

	for {
		start := time.Now()
		err := b.broadcastTransaction(ctx)
		if err != nil {
			b.lggr.Errorw("Error during transaction broadcasting", "err", err)
		} else {
			b.lggr.Debug("Transaction broadcasting time elapsed: ", time.Since(start))
		}
		select {
		case <-ctx.Done():
			return
		case <-b.triggerCh:
			continue
		case <-ticker.C:
			continue
		}
	}
}

func (b *Broadcaster) broadcastTransaction(ctx context.Context) error {
	b.transactionQueueMu.Lock()
	defer b.transactionQueueMu.Unlock()
	for address, transactions := range b.transactionQueue {
		if len(transactions) == 0 {
			continue // no transactions
		}

		latestNonce, err := b.client.NonceAt(ctx, address, nil)
		if err != nil {
			b.lggr.Error("Error getting latest nonce", err)
			continue
		}

		pendingNonce, err := b.client.PendingNonceAt(ctx, address)
		if err != nil {
			b.lggr.Error("Error getting pending nonce", err)
			continue
		}

		if latestNonce > pendingNonce {
			b.lggr.Warnw("Nonce out of sync, skipping address.", "address", address, "latestNonce", latestNonce, "pendingNonce", pendingNonce)
			continue
		}

		// Optimistically send all the subsequent transactions if one or more fail.
		// Best cases scenario, the next transmissions will cover the nonce gap.
		// Worst case scenario, the transactions will be transmitted with gaps, which
		// will be filled on the next transmission cycle.
		var transactionsLeft []*types.Transaction
		for i, tx := range transactions {
			//nolint:gosec // transactions length is fixed
			for nonce := latestNonce; nonce <= pendingNonce+uint64(i); nonce++ {
				tx.Nonce = new(uint64)
				*tx.Nonce = nonce
				attempt, err := b.attemptBuilder.NewAttempt(ctx, b.lggr, tx, b.config.EIP1559)
				if err != nil {
					b.lggr.Error("Error creating new attempt", "transaction", tx, "err", err)
					continue
				}
				start := time.Now()
				txErr := b.client.SendTransaction(ctx, tx, attempt)
				b.lggr.Infow("Broadcasted attempt", "tx", tx, "attempt", attempt, "duration", time.Since(start), "txErr: ", txErr)
			}
			transactionsLeft = append(transactionsLeft, transactions[i:]...)
		}
		b.transactionQueue[address] = transactionsLeft
	}
	return nil
}
