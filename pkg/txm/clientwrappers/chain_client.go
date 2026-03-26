package clientwrappers

import (
	"context"
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
	metrics                     *chainClientMetrics
}

func NewChainClient(lggr logger.Logger, client client.Client, readRequestsToMultipleNodes bool) (*ChainClient, error) {
	chainClientLogger := logger.Sugared(logger.Named(lggr, "Txm.ChainClient"))

	metrics, err := newChainClientMetrics(client.ConfiguredChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize chain client metrics: %w", err)
	}
	if metrics == nil {
		return nil, fmt.Errorf("failed to initialize chain client metrics: nil metrics recorder")
	}

	return &ChainClient{
		lggr:                        chainClientLogger,
		c:                           client,
		readRequestsToMultipleNodes: readRequestsToMultipleNodes,
		metrics:                     metrics,
	}, nil
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

		ctx, cancel := context.WithTimeout(ctx, MultiCallMaxTimeout)
		defer cancel()

		startedAt := time.Now()
		nonce, err := c.c.NonceAtWithFallback(ctx, address, blockNumber)
		callDuration := time.Since(startedAt)
		if err != nil {
			err = fmt.Errorf("error calling NonceAtWithFallback: %w", err)
			c.metrics.recordMultiCallDuration(ctx, "eth_getTransactionCount", blockTag, callDuration, err)
			return 0, err
		}

		c.metrics.recordMultiCallDuration(ctx, "eth_getTransactionCount", blockTag, callDuration, nil)
		c.lggr.Debugw("eth_getTransactionCount", "address", address, "nonce", nonce, "callDuration", callDuration)
		return nonce, nil
	}
	return c.c.NonceAt(ctx, address, blockNumber)
}

func (c *ChainClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if c.readRequestsToMultipleNodes {
		ctx, cancel := context.WithTimeout(ctx, MultiCallMaxTimeout)
		defer cancel()

		startedAt := time.Now()
		nonce, err := c.c.PendingNonceAtWithFallback(ctx, address)
		callDuration := time.Since(startedAt)
		if err != nil {
			err = fmt.Errorf("error calling PendingNonceAtWithFallback: %w", err)
			c.metrics.recordMultiCallDuration(ctx, "eth_getTransactionCount", "pending", callDuration, err)
			return 0, err
		}

		c.metrics.recordMultiCallDuration(ctx, "eth_getTransactionCount", "pending", callDuration, nil)
		c.lggr.Debugw("eth_getTransactionCount", "address", address, "blockTag", "pending", "nonce", nonce, "callDuration", callDuration)
		return nonce, nil
	}
	return c.c.PendingNonceAt(ctx, address)
}

func (c *ChainClient) SendTransaction(ctx context.Context, _ *types.Transaction, attempt *types.Attempt) error {
	return c.c.SendTransaction(ctx, attempt.SignedTransaction)
}
