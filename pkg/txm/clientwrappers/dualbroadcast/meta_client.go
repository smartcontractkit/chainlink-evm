package dualbroadcast

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	timeout = time.Second * 300
)

type MetaClientKeystore interface {
	SignMessage(ctx context.Context, address common.Address, data []byte) ([]byte, error)
	SignTx(ctx context.Context, fromAddress common.Address, tx *evmtypes.Transaction) (*evmtypes.Transaction, error)
}

type MetaClient struct {
	lggr      logger.SugaredLogger
	c         client.Client
	ks        MetaClientKeystore
	customURL *url.URL
	chainID   *big.Int
}

func NewMetaClient(lggr logger.Logger, c client.Client, ks MetaClientKeystore, customURL *url.URL, chainID *big.Int) *MetaClient {
	return &MetaClient{
		lggr:      logger.Sugared(logger.Named(lggr, "MetaClient")),
		c:         c,
		ks:        ks,
		customURL: customURL,
		chainID:   chainID,
	}
}

func (a *MetaClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return a.c.NonceAt(ctx, address, blockNumber)
}

func (a *MetaClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return a.c.PendingNonceAt(ctx, address)
}

func (a *MetaClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	if meta != nil && meta.DualBroadcast != nil && *meta.DualBroadcast && !tx.IsPurgeable {
		meta, err := a.SendRequest(ctx, tx, attempt)
		if err != nil {
			return fmt.Errorf("error sending request for transactionID(%d): %w", tx.ID, err)
		}
		if meta != nil {
			if err := a.SendOperation(ctx, tx, attempt, *meta); err != nil {
				return fmt.Errorf("failed to send operation for transactionID(%d): %w", tx.ID, err)
			}
		}
	}
	return a.c.SendTransaction(ctx, attempt.SignedTransaction)
}

type Parameters struct {
	ChainID      *hexutil.Uint64 `json:"chainId"`
	ToAddress    common.Address  `json:"adapter"`
	Payload      hexutil.Bytes   `json:"updatePayload"`
	ER           bool            `json:"earlyReturn"`
	FromAddress  common.Address  `json:"bundlerEoa"`
	MaxFeePerGas *hexutil.Big    `json:"maxFeePerGas"`
	Signature    hexutil.Bytes   `json:"signature"`
}

type requestResponse struct {
	Result *Metacalldata `json:"result,omitempty"`
	Error  struct {
		ErrorMessage string `json:"message,omitempty"`
	}
}

type Metacalldata struct {
	ToAddress    common.Address `json:"metacallDestination"`
	GasLimit     uint64         `json:"metacallGasLimit"`
	MaxFeePerGas *big.Int       `json:"metacallMaxFeePerGas"`
	CallData     []byte         `json:"metacallCallData"`
}

func (a *MetaClient) SendRequest(parentCtx context.Context, tx *types.Transaction, attempt *types.Attempt) (*Metacalldata, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	m := []byte{97, 116, 108, 97, 115, 95, 111, 101, 118, 65, 117, 99, 116, 105, 111, 110}

	cid := hexutil.Uint64(a.chainID.Uint64())
	var fee hexutil.Big
	if attempt.Fee.ValidDynamic() {
		fee = hexutil.Big(*attempt.Fee.GasFeeCap.ToInt())
	} else {
		fee = hexutil.Big(*attempt.Fee.GasPrice.ToInt())
	}
	auctionParams := Parameters{
		ChainID:      &cid,
		ToAddress:    tx.ToAddress,
		Payload:      tx.Data,
		ER:           true,
		FromAddress:  tx.FromAddress,
		MaxFeePerGas: &fee,
	}

	payload := fmt.Sprintf(
		"%s:%s:%s:%t:%s:%s",
		auctionParams.ChainID.String(),
		auctionParams.ToAddress.Hex(),
		auctionParams.Payload.String(),
		auctionParams.ER,
		auctionParams.FromAddress.Hex(),
		auctionParams.MaxFeePerGas.String(),
	)

	signature, err := a.ks.SignMessage(parentCtx, tx.FromAddress, []byte(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}
	auctionParams.Signature = signature
	marshalledParamsExtended, err := json.Marshal(auctionParams)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed auction params: %w", err)
	}
	body := fmt.Appendf(nil, `{"jsonrpc":"2.0","method":"%s","params":[%s], "id":1}`, string(m), marshalledParamsExtended)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.customURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %v failed with status: %d", req, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response requestResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.ErrorMessage != "" {
		return nil, errors.New(response.Error.ErrorMessage)
	}

	return response.Result, nil
}

func (a *MetaClient) SendOperation(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, meta Metacalldata) error {
	if tx.Nonce == nil {
		return fmt.Errorf("failed to create attempt for txID: %v: nonce empty", tx.ID)
	}

	// TODO: fastest way to avoid overpaying, but might require additional checks.
	tip := meta.MaxFeePerGas
	if attempt.Fee.ValidDynamic() && meta.MaxFeePerGas.Cmp(attempt.Fee.GasTipCap.ToInt()) >= 0 {
		tip = attempt.Fee.GasTipCap.ToInt()
	}
	dynamicTx := evmtypes.DynamicFeeTx{
		Nonce:     *tx.Nonce,
		To:        &meta.ToAddress,
		Gas:       meta.GasLimit,
		GasTipCap: tip,
		GasFeeCap: meta.MaxFeePerGas,
		Data:      meta.CallData,
	}

	signedTx, err := a.ks.SignTx(ctx, tx.FromAddress, evmtypes.NewTx(&dynamicTx))
	if err != nil {
		return fmt.Errorf("failed to sign attempt for txID: %v, err: %w", tx.ID, err)
	}
	a.lggr.Infow("Intercepted attempt for tx", "txID", tx.ID, "toAddress", meta.ToAddress, "gasLimit", meta.GasLimit,
		"TipCap", tip, "FeeCap", meta.MaxFeePerGas)
	return a.c.SendTransaction(ctx, signedTx)
}
