package clientwrappers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// MultiCallMaxTimeout is the maximum timeout for the multi-call operation.
// Given how frequently reads are made, we're making a tradeoff between latency and read availability.
const MultiCallMaxTimeout = 1500 * time.Millisecond

type ChainClient struct {
	lggr                        logger.SugaredLogger
	c                           client.Client
	readRequestsToMultipleNodes bool
}

func NewChainClient(lggr logger.Logger, client client.Client, readRequestsToMultipleNodes bool) *ChainClient {
	return &ChainClient{lggr: logger.Sugared(logger.Named(lggr, "Txm.ChainClient")), c: client, readRequestsToMultipleNodes: readRequestsToMultipleNodes}
}

func (c *ChainClient) BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error) {
	return c.c.BlockByNumber(ctx, number)
}

func (c *ChainClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	if c.readRequestsToMultipleNodes {
		blockTag := "latest"
		if blockNumber != nil {
			blockTag = hexutil.EncodeBig(blockNumber)
		}
		return getTransactionCountMultiCall(ctx, c.c, c.lggr, address, blockTag)
	}
	return c.c.NonceAt(ctx, address, blockNumber)
}

func (c *ChainClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if c.readRequestsToMultipleNodes {
		return getTransactionCountMultiCall(ctx, c.c, c.lggr, address, "pending")
	}
	return c.c.PendingNonceAt(ctx, address)
}

func (c *ChainClient) SendTransaction(ctx context.Context, _ *types.Transaction, attempt *types.Attempt) error {
	return c.c.SendTransaction(ctx, attempt.SignedTransaction)
}

type decodeMultiplexedResultFunc[T any] func(raw json.RawMessage) (T, error)

func multiCallSequential[T any](
	parentCtx context.Context,
	c client.Client,
	method string,
	args []any,
	decode decodeMultiplexedResultFunc[T],
) (result T, callDuration time.Duration, err error) {
	ctx, cancel := context.WithTimeout(parentCtx, MultiCallMaxTimeout)
	defer cancel()

	startedAt := time.Now()
	raw, err := c.CallContextAllSequential(ctx, method, args...)
	callDuration = time.Since(startedAt)
	if err != nil {
		return result, callDuration, fmt.Errorf("error calling %s: %w", method, err)
	}

	decoded, decodeErr := decode(raw)
	if decodeErr != nil {
		return result, callDuration, fmt.Errorf("error decoding %s result: %w", method, decodeErr)
	}

	return decoded, callDuration, nil
}

func getTransactionCountMultiCall(parentCtx context.Context, c client.Client, lggr logger.SugaredLogger, address common.Address, blockTag string) (uint64, error) {
	nonce, callDuration, err := multiCallSequential(
		parentCtx,
		c,
		"eth_getTransactionCount",
		[]any{address, blockTag},
		func(raw json.RawMessage) (uint64, error) {
			var nonce hexutil.Uint64
			if unmarshalErr := json.Unmarshal(raw, &nonce); unmarshalErr != nil {
				return 0, unmarshalErr
			}

			return uint64(nonce), nil
		},
	)
	if err != nil {
		return 0, err
	}

	lggr.Debugw("eth_getTransactionCount", "address", address, "nonce", nonce, "callDuration", callDuration)
	return nonce, nil
}
