package txm

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type StuckTxDetectorConfig struct {
	BlockTime             time.Duration
	StuckTxBlockThreshold uint32
	DetectionURL          string
	DualBroadcast         bool
}

type stuckTxDetector struct {
	lggr         logger.Logger
	chainType    chaintype.ChainType
	config       StuckTxDetectorConfig
	mu           sync.Mutex
	lastPurgeMap map[common.Address]time.Time
}

func NewStuckTxDetector(lggr logger.Logger, chaintype chaintype.ChainType, config StuckTxDetectorConfig) *stuckTxDetector {
	return &stuckTxDetector{
		lggr:         lggr,
		chainType:    chaintype,
		config:       config,
		lastPurgeMap: make(map[common.Address]time.Time),
	}
}

func (s *stuckTxDetector) DetectStuckTransaction(ctx context.Context, tx *types.Transaction) (bool, error) {
	//nolint:gocritic //placeholder for upcoming chaintypes
	switch s.chainType {
	default:
		return s.timeBasedDetection(tx), nil
	}
}

// timeBasedDetection marks a transaction if:
//   - LastBroadcastAt is nil
//   - Total attempt count is equal or greater than the maxAttemptsThreshold
//
// or all the following conditions are met:
//   - LastBroadcastAt is not nil
//   - Time since last broadcast is above the threshold
//   - Time since last purge is above threshold
//
// NOTE: Potentially we can use a subset of threhsold for last purge check, because the transaction would have already been broadcasted to the mempool
// so it is more likely to be picked up compared to a transaction that hasn't been broadcasted before. This would avoid slowing down TXM for sebsequent transactions
// in case the current one is stuck.
func (s *stuckTxDetector) timeBasedDetection(tx *types.Transaction) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	threshold := (s.config.BlockTime * time.Duration(s.config.StuckTxBlockThreshold))
	if tx.LastBroadcastAt == nil {
		if tx.AttemptCount >= maxAttemptsThreshold {
			s.lggr.Debugf("TxID: %v reached max attempts threshold: %d. Transaction is now considered stuck and will be purged.",
				tx.ID, maxAttemptsThreshold)
			s.lastPurgeMap[tx.FromAddress] = time.Now()
			return true
		}
		return false
	}

	if last := s.lastPurgeMap[tx.FromAddress]; min(time.Since(*tx.LastBroadcastAt), time.Since(last)) > threshold {
		s.lggr.Debugf("TxID: %v last broadcast was: %v and last purge: %v which is more than the max configured duration: %v. Transaction is now considered stuck and will be purged.",
			tx.ID, tx.LastBroadcastAt, last, threshold)
		s.lastPurgeMap[tx.FromAddress] = time.Now()
		return true
	}
	return false
}
