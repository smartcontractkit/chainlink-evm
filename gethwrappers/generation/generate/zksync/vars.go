package zksyncwrapper

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// these are mocks for the placeholders used in the template
type PlaceholderContractName struct{}

type IPlaceholderContractNameMetaData struct {
	GetAbi func() (*abi.ABI, error)
}

var PlaceholderContractNameMetaData = IPlaceholderContractNameMetaData{
	GetAbi: func() (*abi.ABI, error) {
		return nil, nil
	},
}

var ZkBytecode = []byte{}

func NewPlaceholderContractName(address common.Address, backend bind.ContractBackend) (*PlaceholderContractName, error) {
	return nil, nil
}
