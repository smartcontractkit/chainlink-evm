// Package ocr2transmit detects OCR2 aggregator transmit calldata for metrics (DF-22761 / DF-22643).
package ocr2transmit

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
)

var (
	transmitMethodID     []byte
	transmitMethodIDOnce sync.Once
)

func transmitSig() []byte {
	transmitMethodIDOnce.Do(func() {
		parsed, err := ocr2aggregator.OCR2AggregatorMetaData.GetAbi()
		if err != nil {
			return
		}
		m, ok := parsed.Methods["transmit"]
		if !ok {
			return
		}
		transmitMethodID = m.ID
	})
	return transmitMethodID
}

// IsTransmitCalldata returns true if data begins with the OCR2Aggregator transmit function selector.
func IsTransmitCalldata(data []byte) bool {
	sig := transmitSig()
	if len(sig) != 4 || len(data) < 4 {
		return false
	}
	return data[0] == sig[0] && data[1] == sig[1] && data[2] == sig[2] && data[3] == sig[3]
}

// ContractLabel returns the logical aggregator address for metrics: meta forwarder dest if set, else to address.
func ContractLabel(to common.Address, fwdrDest *common.Address) string {
	if fwdrDest != nil && *fwdrDest != (common.Address{}) {
		return fwdrDest.Hex()
	}
	return to.Hex()
}
