package gas_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

// TestRefreshDynamicPrice_EmptyBaseFee is a regression test for a production
// crash: FeeHistoryEstimator.RefreshDynamicPrice panicked with
// "index out of range [-1]" whenever the upstream eth_feeHistory response had
// an empty BaseFee ([]) array. Real RPC providers do return this shape (e.g.
// on a deep/archive range some providers can't serve, or on chains without
// full EIP-1559 support), and it reached this code path unmodified from the
// underlying client — no error, just an empty slice.
func TestRefreshDynamicPrice_EmptyBaseFee(t *testing.T) {
	client := mocks.NewFeeHistoryEstimatorClient(t)
	chainID := testutils.FixtureChainID

	feeHistoryResult := &ethereum.FeeHistory{
		OldestBlock:  big.NewInt(100),
		Reward:       [][]*big.Int{{big.NewInt(1), big.NewInt(2)}},
		BaseFee:      []*big.Int{}, // reproduction: empty BaseFee array, err == nil
		GasUsedRatio: nil,
	}
	client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

	cfg := gas.FeeHistoryEstimatorConfig{BlockHistorySize: 2, RewardPercentile: 60}
	u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)

	assert.NotPanics(t, func() {
		err := u.RefreshDynamicPrice()
		assert.ErrorContains(t, err, "empty baseFeePerGas array")
	})
}
