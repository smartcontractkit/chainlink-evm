package client

import (
	"context"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// contractBackendV2 adapts a Client to bind/v2 ContractBackend, whose
// TransactionByHash gained a "pending" bool return in go-ethereum v1.17.8.
type contractBackendV2 struct {
	Client
}

func (a contractBackendV2) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	tx, err := a.Client.TransactionByHash(ctx, txHash)
	return tx, false, err
}

// AsContractBackend returns the Client as a bind/v2 ContractBackend, adapting
// the TransactionByHash signature changed in go-ethereum v1.17.8.
func AsContractBackend(c Client) bind.ContractBackend {
	return contractBackendV2{c}
}
