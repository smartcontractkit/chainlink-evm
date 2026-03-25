package ocr2transmit

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
)

func TestIsTransmitCalldata(t *testing.T) {
	parsed, err := ocr2aggregator.OCR2AggregatorMetaData.GetAbi()
	require.NoError(t, err)
	m := parsed.Methods["transmit"]
	require.NotNil(t, m)

	// Minimal non-empty args for selector match only (Pack may fail on full args); we only need first 4 bytes.
	data := append([]byte{}, m.ID...)
	require.True(t, IsTransmitCalldata(data))

	require.False(t, IsTransmitCalldata([]byte{1, 2, 3}))
	require.False(t, IsTransmitCalldata(nil))
}

func TestContractLabel(t *testing.T) {
	to := common.HexToAddress("0x0000000000000000000000000000000000000001")
	dest := common.HexToAddress("0x0000000000000000000000000000000000000002")
	require.Equal(t, dest.Hex(), ContractLabel(to, &dest))
	require.Equal(t, to.Hex(), ContractLabel(to, nil))
}
