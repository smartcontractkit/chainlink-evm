package transmitter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configmocks "github.com/smartcontractkit/chainlink-evm/pkg/config/mocks"
)

func TestGetGasLimitFrom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		defaultGasLimit     uint64
		ocr2GasLimit        *uint32
		relayConfigGasLimit *uint32
		pluginGasLimit      *uint32
		expectedGasLimit    uint64
		description         string
	}{
		{
			name:                "uses default gas limit when no overrides",
			defaultGasLimit:     500000,
			ocr2GasLimit:        nil,
			relayConfigGasLimit: nil,
			pluginGasLimit:      nil,
			expectedGasLimit:    500000,
			description:         "Should use default gas limit from chain config",
		},
		{
			name:                "uses OCR2 override when set",
			defaultGasLimit:     500000,
			ocr2GasLimit:        uint32Ptr(400000),
			relayConfigGasLimit: nil,
			pluginGasLimit:      nil,
			expectedGasLimit:    400000,
			description:         "Should use OCR2 gas limit override from chain config",
		},
		{
			name:                "relayconfig gas limit overrides OCR2 gas limit",
			defaultGasLimit:     500000,
			ocr2GasLimit:        uint32Ptr(400000),
			relayConfigGasLimit: uint32Ptr(300000),
			pluginGasLimit:      nil,
			expectedGasLimit:    300000,
			description:         "Relay config gas limit should override OCR2 gas limit",
		},
		{
			name:                "plugin gas limit overrides relay config gas limit",
			defaultGasLimit:     500000,
			ocr2GasLimit:        uint32Ptr(400000),
			relayConfigGasLimit: uint32Ptr(300000),
			pluginGasLimit:      uint32Ptr(200000),
			expectedGasLimit:    200000,
			description:         "Plugin gas limit should override relay config gas limit",
		},
		{
			name:                "relay config gas limit overrides default",
			defaultGasLimit:     500000,
			ocr2GasLimit:        nil,
			relayConfigGasLimit: uint32Ptr(300000),
			pluginGasLimit:      nil,
			expectedGasLimit:    300000,
			description:         "Relay config gas limit should override default limit",
		},
		{
			name:                "plugin gas limit overrides default",
			defaultGasLimit:     500000,
			ocr2GasLimit:        nil,
			relayConfigGasLimit: nil,
			pluginGasLimit:      uint32Ptr(200000),
			expectedGasLimit:    200000,
			description:         "Plugin gas limit should override default limit",
		},
		{
			name:                "plugin gas limit overrides OCR2 gas limit",
			defaultGasLimit:     500000,
			ocr2GasLimit:        uint32Ptr(400000),
			relayConfigGasLimit: nil,
			pluginGasLimit:      uint32Ptr(200000),
			expectedGasLimit:    200000,
			description:         "Plugin gas limit should override OCR2 limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockGasEstimator := configmocks.NewGasEstimator(t)
			mockLimitJobType := configmocks.NewLimitJobType(t)

			// Setup gas estimator mock
			mockGasEstimator.EXPECT().LimitDefault().Return(tt.defaultGasLimit)
			mockGasEstimator.EXPECT().LimitJobType().Return(mockLimitJobType)
			mockLimitJobType.EXPECT().OCR2().Return(tt.ocr2GasLimit)

			gasLimit := getGasLimitFrom(mockGasEstimator, ConfigTransmitterOpts{
				PluginGasLimit: tt.pluginGasLimit,
			}, tt.relayConfigGasLimit)

			if !assert.Equal(t, tt.expectedGasLimit, gasLimit, tt.description) {
				t.Errorf("expected gas limit: %d, got: %d", tt.expectedGasLimit, gasLimit) // to print in decimal format
			}
		})
	}
}

func uint32Ptr(v uint32) *uint32 {
	return &v
}
