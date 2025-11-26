package transmitter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-evm/pkg/chains/legacyevm/mocks"
	configmocks "github.com/smartcontractkit/chainlink-evm/pkg/config/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
)

func TestGetGasLimitFrom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		defaultGasLimit  uint64
		ocr2GasLimit     *uint32
		pluginGasLimit   *uint32
		expectedGasLimit uint64
		description      string
	}{
		{
			name:             "uses default gas limit when no overrides",
			defaultGasLimit:  500000,
			ocr2GasLimit:     nil,
			pluginGasLimit:   nil,
			expectedGasLimit: 500000,
			description:      "Should use default gas limit from chain config",
		},
		{
			name:             "uses OCR2 override when set",
			defaultGasLimit:  500000,
			ocr2GasLimit:     uint32Ptr(300000),
			pluginGasLimit:   nil,
			expectedGasLimit: 300000,
			description:      "Should use OCR2 gas limit override from chain config",
		},
		{
			name:             "plugin gas limit overrides OCR2",
			defaultGasLimit:  500000,
			ocr2GasLimit:     uint32Ptr(400000),
			pluginGasLimit:   uint32Ptr(100000),
			expectedGasLimit: 100000,
			description:      "Plugin gas limit should override OCR2 limit",
		},
		{
			name:             "plugin gas limit overrides default",
			defaultGasLimit:  500000,
			ocr2GasLimit:     nil,
			pluginGasLimit:   uint32Ptr(150000),
			expectedGasLimit: 150000,
			description:      "Plugin gas limit should override default limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockChain := mocks.NewChain(t)
			mockChainConfig := configmocks.NewChainScopedConfig(t)
			mockEVM := configmocks.NewEVM(t)
			mockGasEstimator := configmocks.NewGasEstimator(t)

			// Setup chain mock
			mockChain.EXPECT().Config().Return(mockChainConfig)

			// Setup chain config mock
			mockChainConfig.EXPECT().EVM().Return(mockEVM)

			// Setup EVM config mock
			mockEVM.EXPECT().GasEstimator().Return(mockGasEstimator)

			// Setup gas estimator mock
			mockGasEstimator.EXPECT().LimitDefault().Return(tt.defaultGasLimit)
			mockGasEstimator.EXPECT().LimitJobType().Return(&mockJobTypeConfig{gasLimit: tt.ocr2GasLimit})

			gasLimit := getGasLimitFrom(mockChain, ConfigTransmitterOpts{
				PluginGasLimit: tt.pluginGasLimit,
			})

			if !assert.Equal(t, tt.expectedGasLimit, gasLimit, tt.description) {
				t.Errorf("expected gas limit: %d, got: %d", tt.expectedGasLimit, gasLimit) // to print in decimal format
			}
		})
	}
}

type mockJobTypeConfig struct {
	txmgr.TestLimitJobTypeConfig
	gasLimit *uint32
}

func (x mockJobTypeConfig) OCR2() *uint32 {
	return x.gasLimit
}

func uint32Ptr(v uint32) *uint32 {
	return &v
}
