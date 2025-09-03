package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func MustGetABI(json string) abi.ABI {
	abi, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		panic("could not parse ABI: " + err.Error())
	}
	return abi
}

// NullClientChainID is set to a chainID that is unlikely to be used in production.
// It cannot be zero due to a breaking change in go-ethereum:
// https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction_signing.go#L193
const NullClientChainID = 1399100
