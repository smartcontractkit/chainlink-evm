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

	if meta != nil && meta.DualBroadcast != nil && *meta.DualBroadcast && !tx.IsPurgeable && meta.DualBroadcastParams != nil {
		meta, err := a.SendRequest(ctx, tx, attempt, meta.DualBroadcastParams)
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
	Result *ResponseResult `json:"result"`
	Error  struct {
		ErrorMessage string `json:"message,omitempty"`
	}
}

type ResponseResult struct {
	UOP *UO `json:"userOperation"`
	SOP *SO `json:"solverOperations"`
	DOP *DO `json:"dAppOperation"`
	Metacalldata
}

type UO struct {
	To        common.Address `json:"to"`
	Dapp      common.Address `json:"dapp"`
	Control   common.Address `json:"control"`
	Data      hexutil.Bytes  `json:"data"`
	Signature hexutil.Bytes  `json:"signature"`
}

type SO struct {
	To      common.Address `json:"to"`
	Control common.Address `json:"control"`
}

type DO struct {
	To      common.Address `json:"to"`
	Control common.Address `json:"control"`
	Bundler common.Address `json:"bundler"`
}

type Metacalldata struct {
	ToAddress    common.Address `json:"metacallDestination"`
	GasLimit     uint64         `json:"metacallGasLimit"`
	MaxFeePerGas *big.Int       `json:"metacallMaxFeePerGas"`
	CallData     []byte         `json:"metacallCallData"`
}

func (a *MetaClient) SendRequest(parentCtx context.Context, tx *types.Transaction, attempt *types.Attempt, dualBroadcastParams *string) (*Metacalldata, error) {
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
	params := Parameters{
		ChainID:      &cid,
		ToAddress:    tx.ToAddress,
		Payload:      tx.Data,
		ER:           true,
		FromAddress:  tx.FromAddress,
		MaxFeePerGas: &fee,
	}

	payload := fmt.Sprintf(
		"%s:%s:%s:%t:%s:%s",
		params.ChainID.String(),
		params.ToAddress.Hex(),
		params.Payload.String(),
		params.ER,
		params.FromAddress.Hex(),
		params.MaxFeePerGas.String(),
	)

	signature, err := a.ks.SignMessage(parentCtx, tx.FromAddress, []byte(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}
	params.Signature = signature
	marshalledParamsExtended, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed params: %w", err)
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

	return verifyResponse(tx.Data, signature, tx.FromAddress, response.Result, dualBroadcastParams)
}

func verifyResponse(txData []byte, signature []byte, fromAddress common.Address, result *ResponseResult, dualBroadcastParams *string) (*Metacalldata, error) {
	params, err := url.ParseQuery(*dualBroadcastParams)
	if err != nil {
		return nil, err
	}

	info := params["info"]
	if len(info) != 2 {
		return nil, fmt.Errorf("incorrect size for info: %v", info)
	}
	to := common.HexToAddress(info[0])
	dApp := common.HexToAddress(info[1])
	if result.Metacalldata.ToAddress != to ||
		result.UOP.To != to ||
		result.SOP.To != to ||
		result.DOP.To != to {
		return nil, fmt.Errorf("incorrect destination address, metacall.To: %v, uOp.To: %v, sOp.To: %v, dOp.To: %v, info.To: %v",
			result.Metacalldata.ToAddress, result.UOP.To, result.SOP.To, result.DOP.To, to)
	}

	if result.UOP.Control != dApp || result.UOP.Dapp != dApp || result.DOP.Control != dApp || result.SOP.Control != dApp {
		return nil, fmt.Errorf("incorrect dApp address, uOp.To: %v, uOp.dApp: %v, dOp.Control: %v, sOp.Control: %v, info.dApp: %v",
			result.UOP.Control, result.UOP.Dapp, result.DOP.Control, result.SOP.Control, dApp)
	}

	if !bytes.Equal(result.UOP.Data, txData) || !bytes.Contains(result.Metacalldata.CallData, txData) {
		return nil, fmt.Errorf("incorrect calldata, uOp.Data: %v, metacall.CallData: %v, txData: %v", result.UOP.Data, result.Metacalldata.CallData, txData)
	}

	if !bytes.Equal(result.UOP.Signature, signature) {
		return nil, fmt.Errorf("incorrect signature, uOp.Signature: %v, signature: %v", result.UOP.Signature, signature)
	}

	if result.DOP.Bundler != fromAddress {
		return nil, fmt.Errorf("incorrect bundler, dOp.Bundler: %v, fromAddress: %v", result.DOP.Bundler, fromAddress)
	}

	return &result.Metacalldata, nil
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
