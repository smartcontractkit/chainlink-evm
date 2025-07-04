package tron

import (
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	tronkeystore "github.com/smartcontractkit/chainlink-tron/relayer/keystore"
	tronclient "github.com/smartcontractkit/chainlink-tron/relayer/sdk"
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
	tronClient, err := tronclient.CreateFullNodeClient(fullNodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create tron client: %w", err)
	}

	fixedEnergyValue := new(big.Int).SetUint64(cfg.LimitDefault()).Int64()

	return trontxm.New(logger, tronkeystore.NewLoopKeystoreAdapter(keystore), tronClient, trontxm.TronTxmConfig{
		// Overrides the energy estimator to always use the fixed energy
		FixedEnergyValue: fixedEnergyValue,
		// Maximum number of transactions to buffer in the broadcast channel.
		BroadcastChanSize: 100,
		// Number of seconds to wait between polling the blockchain for transaction confirmation.
		ConfirmPollSecs: 5,
	}), nil
}
