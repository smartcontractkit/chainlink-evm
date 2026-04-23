package dualbroadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const maxBlockDiff = 24

// Flashbots bundle submission (mev_sendBundle). Kept separate from generic eth_sendRawTransaction flow in ofa_backend.go.

func (d *ofaTXClient) sendBundle(ctx context.Context, fromAddress common.Address, meta *types.TxMeta) error {
	var urlParams string
	if meta != nil && meta.DualBroadcastParams != nil {
		urlParams = *meta.DualBroadcastParams
	}
	unconfirmedTxs, err := d.txStore.FetchUnconfirmedTransactions(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch unconfirmed transactions: %w", err)
	}

	attempts := make([]*types.Attempt, 0, len(unconfirmedTxs))
	attemptIDs := make([]uint64, 0, len(unconfirmedTxs))
	nonces := make([]uint64, 0, len(unconfirmedTxs))
	ids := make([]uint64, 0, len(unconfirmedTxs))
	for _, unconfirmedTx := range unconfirmedTxs {
		if len(unconfirmedTx.Attempts) > 0 && unconfirmedTx.Nonce != nil && unconfirmedTx.Attempts[len(unconfirmedTx.Attempts)-1].SignedTransaction != nil {
			latestAttempt := unconfirmedTx.Attempts[len(unconfirmedTx.Attempts)-1]
			attempts = append(attempts, latestAttempt)
			attemptIDs = append(attemptIDs, latestAttempt.ID)
			ids = append(ids, unconfirmedTx.ID)
		}
	}

	if len(attempts) < 2 {
		return nil
	}

	prevNonce := attempts[0].SignedTransaction.Nonce()
	nonces = append(nonces, prevNonce)
	for _, attempt := range attempts[1:] {
		nonce := attempt.SignedTransaction.Nonce()
		nonces = append(nonces, nonce)
		expectedNonce := prevNonce + 1
		if nonce != expectedNonce {
			return fmt.Errorf("bundle attempts must be contiguous and strictly increasing: expected nonce %d, got nonce %d", expectedNonce, nonce)
		}
		prevNonce = nonce
	}

	currentBlock, err := d.c.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get current block height: %w", err)
	}
	maxBlock := currentBlock.NumberU64() + maxBlockDiff

	bodyItems := make([]map[string]any, 0, len(attempts))
	for _, attempt := range attempts {
		txData, err := attempt.SignedTransaction.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal transaction for attempt ID %d: %w", attempt.ID, err)
		}

		bodyItems = append(bodyItems, map[string]any{
			"tx":         hexutil.Encode(txData),
			"revertMode": "allow",
		})
	}
	privacy, refundConfig, err := parseURLParams(urlParams)
	if err != nil {
		return err
	}

	bundleParams := map[string]any{
		"body": bodyItems,
		"inclusion": map[string]any{
			"block":    hexutil.EncodeBig(new(big.Int).SetUint64(currentBlock.NumberU64())),
			"maxBlock": hexutil.EncodeBig(new(big.Int).SetUint64(maxBlock)),
		},
		"privacy": privacy,
		"version": "v0.1",
	}
	if refundConfig.Address != "" {
		bundleParams["validity"] = map[string]any{
			"refundConfig": []any{refundConfig},
		}
	}

	requestBody := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "mev_sendBundle",
		"params":  []any{bundleParams},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal bundle request: %w", err)
	}

	raw, err := d.postJSONRPC(ctx, fromAddress, bodyBytes, nil)
	if err != nil {
		return err
	}

	var bundleResult struct {
		BundleHash string `json:"bundleHash"`
	}
	if err := json.Unmarshal(raw, &bundleResult); err != nil {
		return fmt.Errorf("failed to decode response %s into bundle result: %w", string(raw), err)
	}
	d.lggr.Infow("Broadcasted transaction bundle", "txIDs", ids, "attemptIDs", attemptIDs, "nonces", nonces, "bundleHash", bundleResult.BundleHash)
	return nil
}

func parseURLParams(params string) (privacy, refundConfig, error) {
	values, err := url.ParseQuery(params)
	if err != nil {
		return privacy{}, refundConfig{}, fmt.Errorf("unable to parse params: %w", err)
	}

	pvc := privacy{}
	if timeout, err := strconv.Atoi(values.Get("auctionTimeout")); err == nil {
		pvc.AuctionTimeout = timeout
	}

	pvc.Builders = append(pvc.Builders, values["builder"]...)

	pvc.Hints = append(pvc.Hints, values["hint"]...)

	refundCfg := refundConfig{}
	refundRaw := values.Get("refund")
	if refundRaw != "" {
		parts := strings.Split(refundRaw, ":")
		if len(parts) != 2 {
			return privacy{}, refundConfig{}, fmt.Errorf("unable to parse refund: %s. Expected format: address:percent", refundRaw)
		}
		address := parts[0]
		percentVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return privacy{}, refundConfig{}, fmt.Errorf("unable to parse percentage: %w", err)
		}

		pvc.WantRefund = percentVal
		refundCfg = refundConfig{
			Address: address,
			Percent: 100,
		}
	}
	return pvc, refundCfg, nil
}

type privacy struct {
	WantRefund     int      `json:"wantRefund"`
	AuctionTimeout int      `json:"auctionTimeout"`
	Builders       []string `json:"builders"`
	Hints          []string `json:"hints"`
}

type refundConfig struct {
	Address string `json:"address"`
	Percent int    `json:"percent"`
}
