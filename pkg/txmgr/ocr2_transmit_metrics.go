package txmgr

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-evm/pkg/ocr2transmit"
)

func recordOCR2TransmitOutcomeV1(chainID *big.Int, tx *Tx, outcome string) {
	if tx == nil {
		return
	}
	var fwdr *common.Address
	if meta, err := tx.GetMeta(); err == nil && meta != nil && meta.FwdrDestAddress != nil {
		fwdr = meta.FwdrDestAddress
	}
	cid := chainID
	if cid == nil {
		cid = tx.ChainID
	}
	ocr2transmit.RecordOutcome(cid, tx.FromAddress, tx.ToAddress, tx.EncodedPayload, fwdr, outcome)
}
