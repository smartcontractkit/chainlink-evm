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

// MultiplexingMaxTimeout is the maximum timeout for the multiplexing operation.
// Given how frequently reads are made, we're making a tradeoff between latency and read availability.
const MultiplexingMaxTimeout = 1 * time.Second

type ChainClient struct {
	lggr         logger.SugaredLogger
	c            client.Client
	multiplexing bool
}

func NewChainClient(lggr logger.Logger, client client.Client, multiplexing bool) *ChainClient {
	return &ChainClient{lggr: logger.Sugared(logger.Named(lggr, "Txm.ChainClient")), c: client, multiplexing: multiplexing}
}

func (c *ChainClient) BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error) {
	return c.c.BlockByNumber(ctx, number)
}

func (c *ChainClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	if c.multiplexing {
		blockTag := "latest"
		if blockNumber != nil {
			blockTag = hexutil.EncodeBig(blockNumber)
		}
		return GetTransactionCountMultiplexed(ctx, c.c, c.lggr, address, blockTag)
	}
	return c.c.NonceAt(ctx, address, blockNumber)
}

func (c *ChainClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if c.multiplexing {
		return GetTransactionCountMultiplexed(ctx, c.c, c.lggr, address, "pending")
	}
	return c.c.PendingNonceAt(ctx, address)
}

func (c *ChainClient) SendTransaction(ctx context.Context, _ *types.Transaction, attempt *types.Attempt) error {
	return c.c.SendTransaction(ctx, attempt.SignedTransaction)
}

type DecodeMultiplexedResultFunc[T any] func(raw json.RawMessage) (T, error)
type IsBetterMultiplexedResultFunc[T any] func(candidate T, current T) bool

func multiplexCallBest[T any](
	parentCtx context.Context,
	c client.Client,
	method string,
	args []interface{},
	decode DecodeMultiplexedResultFunc[T],
	isBetter IsBetterMultiplexedResultFunc[T],
) (best T, allSuccessful []T, callDuration time.Duration, err error) {
	ctx := parentCtx
	cancel := func() {}
	ctx, cancel = context.WithTimeout(parentCtx, MultiplexingMaxTimeout)
	defer cancel()

	startedAt := time.Now()
	results, err := c.CallContextAll(ctx, method, args...)
	callDuration = time.Since(startedAt)
	if err != nil {
		return best, nil, callDuration, fmt.Errorf("error multiplexing %s call: %w", method, err)
	}

	allSuccessful = make([]T, 0, len(results))
	found := false
	for _, result := range results {
		if result.Err != nil {
			continue
		}

		decoded, decodeErr := decode(result.Result)
		if decodeErr != nil {
			continue
		}

		allSuccessful = append(allSuccessful, decoded)
		if !found || isBetter(decoded, best) {
			best = decoded
			found = true
		}
	}

	if !found {
		return best, nil, callDuration, fmt.Errorf("%s returned no successful results from primary nodes", method)
	}

	return best, allSuccessful, callDuration, nil
}

func GetTransactionCountMultiplexed(parentCtx context.Context, c client.Client, lggr logger.SugaredLogger, address common.Address, blockTag string) (uint64, error) {
	highest, nonces, callDuration, err := multiplexCallBest(
		parentCtx,
		c,
		"eth_getTransactionCount",
		[]interface{}{address, blockTag},
		func(raw json.RawMessage) (uint64, error) {
			var nonce hexutil.Uint64
			if unmarshalErr := json.Unmarshal(raw, &nonce); unmarshalErr != nil {
				return 0, unmarshalErr
			}

			return uint64(nonce), nil
		},
		func(candidate uint64, current uint64) bool {
			return candidate > current
		},
	)
	if err != nil {
		return 0, err
	}

	lggr.Debugw("TransactionCount", "address", address, "nonces", nonces, "callDuration", callDuration)
	return highest, nil
}
