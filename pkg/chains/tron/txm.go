package tron

import (
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	tronkeystore "github.com/smartcontractkit/chainlink-tron/relayer/keystore"
	"github.com/smartcontractkit/chainlink-tron/relayer/sdk"
	trontxm "github.com/smartcontractkit/chainlink-tron/relayer/txm"
)

type TxmConfig interface {
	LimitDefault() uint64
}

func ConstructTxm(logger logger.Logger, cfg TxmConfig, nodes []*toml.Node, keystore keys.Store) (*trontxm.TronTxm, error) {
	if len(nodes) == 0 {
		return nil, fmt.Errorf("Tron chain requires at least one node")
	}

	fullNodeURL := nodes[0].HTTPURLExtraWrite.URL()

	// This is only used for CCIP 1.5, which doesn't need to poll for finality.
	// By using the same URL for both solidity and full node, transactions will
	// be marked as finalized upon confirmation, which is acceptable in this case.
	combinedClient, err := sdk.CreateCombinedClient(fullNodeURL, fullNodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create combined client: %w", err)
	}

	fixedEnergyValue := new(big.Int).SetUint64(cfg.LimitDefault()).Int64()

	return trontxm.New(logger, tronkeystore.NewLoopKeystoreAdapter(keystore), combinedClient, trontxm.TronTxmConfig{
		// Overrides the energy estimator to always use the fixed energy
		FixedEnergyValue: fixedEnergyValue,
		// Maximum number of transactions to buffer in the broadcast channel.
		BroadcastChanSize: 100,
		// Number of seconds to wait between polling the blockchain for transaction confirmation.
		ConfirmPollSecs: 5,
		// How long transactions are kept in the txm
		RetentionPeriod: 0,
		// How often to reap the txm of finished transactions
		ReapInterval: 1 * time.Minute,
	}), nil
}
