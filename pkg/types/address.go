package types

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/keystore/corekeys/ethkey"
)

type EIP55Address = ethkey.EIP55Address

type EIP55AddressCollection = ethkey.EIP55AddressCollection

func NewEIP55Address(s string) (EIP55Address, error) {
	return ethkey.NewEIP55Address(s)
}

func MustEIP55Address(s string) EIP55Address {
	return ethkey.MustEIP55Address(s)
}

func EIP55AddressFromAddress(a common.Address) EIP55Address {
	return ethkey.EIP55AddressFromAddress(a)
}
