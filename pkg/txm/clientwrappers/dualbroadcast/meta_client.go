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
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	timeout      = time.Second * 300
	metaMethodID = "0x4317ca01"
	ABI          = `[
  {
    "inputs": [
      {
        "internalType": "address",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "type": "bytes"
      }
    ],
    "name": "update",
    "type": "function"
  }
]`
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

	if meta != nil && meta.DualBroadcast != nil && *meta.DualBroadcast && !tx.IsPurgeable && meta.DualBroadcastParams != nil && meta.FwdrDestAddress != nil {
		meta, err := a.SendRequest(ctx, tx, attempt, meta.DualBroadcastParams, *meta.FwdrDestAddress)
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

type RequestResponse struct {
	Result *ResponseResult `json:"result"`
	Error  struct {
		ErrorMessage string `json:"message,omitempty"`
	}
}

type ResponseResult struct {
	UOP  *UO   `json:"userOperation"`
	SOPs []*SO `json:"solverOperations"`
	DOP  *DO   `json:"dAppOperation"`
	Metacalldata
}

type UO struct {
	To           common.Address `json:"to"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	Data         hexutil.Bytes  `json:"data"`
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
	GasLimit     *hexutil.Big   `json:"metacallGasLimit"`
	MaxFeePerGas *hexutil.Big   `json:"metacallMaxFeePerGas"`
	CallData     hexutil.Bytes  `json:"metacallCallData"`
}

func (a *MetaClient) SendRequest(parentCtx context.Context, tx *types.Transaction, attempt *types.Attempt, dualBroadcastParams *string, fwdrDestAddress common.Address) (*Metacalldata, error) {
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

	var response RequestResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.Error.ErrorMessage != "" {
		return nil, errors.New(response.Error.ErrorMessage)
	}

	if response.Result == nil {
		return nil, errors.New("empty response")
	}

	return VerifyResponse(tx.Data, tx.FromAddress, response.Result, dualBroadcastParams, fwdrDestAddress)
}

func VerifyResponse(txData []byte, fromAddress common.Address, result *ResponseResult, dualBroadcastParams *string, fwdrDestAddress common.Address) (*Metacalldata, error) {
	params, err := url.ParseQuery(*dualBroadcastParams)
	if err != nil {
		return nil, err
	}

	destination := params["destination"]
	dapp := params["dapp"]
	if len(destination) != 1 || len(dapp) != 1 {
		return nil, fmt.Errorf("incorrect size for info params: %v - %v", destination, dapp)
	}
	to := common.HexToAddress(destination[0]) // metacall address
	dApp := common.HexToAddress(dapp[0])
	abi, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, fmt.Errorf("couldn't read ABI: %w", err)
	}

	if result.UOP == nil {
		return nil, errors.New("nil UOP for metacall")
	}
	updateFn, ok := abi.Methods["update"]
	if !ok {
		return nil, errors.New("update method not found in ABI")
	}
	if !bytes.HasPrefix(result.UOP.Data, updateFn.ID) {
		return nil, fmt.Errorf("incorrect method id in uop.Data: %v", result.UOP.Data)
	}

	args, err := updateFn.Inputs.UnpackValues(result.UOP.Data[4:]) // remove function selector
	if err != nil || len(args) < 2 {
		return nil, fmt.Errorf("unpack failed, data: %v, err: %w", result.UOP.Data, err)
	}

	destinationAddress, ok := args[0].(common.Address)
	if !ok {
		return nil, fmt.Errorf("incorrect type for update.from: %v", args[0])
	}

	updateCalldata, ok := args[1].([]byte)
	if !ok {
		return nil, fmt.Errorf("incorrect type for update.calldata: %v", args[1])
	}

	// Metacall
	if result.Metacalldata.ToAddress != to || !strings.HasPrefix(result.Metacalldata.CallData.String(), metaMethodID) || !bytes.Contains(result.Metacalldata.CallData, txData) {
		return nil, fmt.Errorf("incorrect metacall: metacall.ToAddress: %v, metacall.CallData: %v, to: %v, metaMethodID: %v, txData: %v",
			result.Metacalldata.ToAddress, result.Metacalldata.CallData.String(), to, metaMethodID, txData)
	}

	// DOP
	if result.DOP == nil {
		return nil, errors.New("nil DOP for metacall")
	}
	if result.DOP.To != to || result.DOP.Control != dApp || result.DOP.Bundler != fromAddress {
		return nil, fmt.Errorf("incorrect DOP: dop.To: %v, dop.Control: %v, dop.Bundler: %v, to: %v, dApp: %v, fromAddress: %v",
			result.DOP.To, result.DOP.Control, result.DOP.Bundler, to, dApp, fromAddress)
	}

	// SOP
	atLeastOne := false
	for _, sop := range result.SOPs {
		if sop != nil {
			if sop.To != to || sop.Control != dApp {
				// Exit early
				return nil, fmt.Errorf("incorrect SOP: sop.To: %v, sop.Control: %v, to: %v, dApp: %v", sop.To, sop.Control, to, dApp)
			}
			atLeastOne = true
		}
	}
	if !atLeastOne {
		return nil, errors.New("no valid sop")
	}

	// UOP
	if result.UOP.To != to ||
		result.UOP.MaxFeePerGas == nil || result.Metacalldata.MaxFeePerGas == nil || result.UOP.MaxFeePerGas.ToInt().Cmp(result.Metacalldata.MaxFeePerGas.ToInt()) != 0 ||
		result.UOP.Dapp != dApp ||
		result.UOP.Control != dApp ||
		destinationAddress != fwdrDestAddress || !bytes.Equal(updateCalldata, txData) {
		return nil, fmt.Errorf("incorrect UOP: uop.To: %v, uop.MaxFeePerGas: %v, uop.Dapp: %v, uop.update.destinationAddress: %v, uop.update.calldata: %v, to: %v, metacall.MaxFeePerGas: %v, dApp: %v, fwdrDestAddress: %v, txData: %v",
			result.UOP.To, result.UOP.MaxFeePerGas, result.UOP.Dapp, destinationAddress, updateCalldata, to, result.Metacalldata.MaxFeePerGas, dApp, fwdrDestAddress, txData)
	}

	return &result.Metacalldata, nil
}

func (a *MetaClient) SendOperation(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, meta Metacalldata) error {
	if tx.Nonce == nil {
		return fmt.Errorf("failed to create attempt for txID: %v: nonce empty", tx.ID)
	}

	// TODO: fastest way to avoid overpaying, but might require additional checks.
	tip := meta.MaxFeePerGas.ToInt()
	if attempt.Fee.ValidDynamic() && meta.MaxFeePerGas.ToInt().Cmp(attempt.Fee.GasTipCap.ToInt()) >= 0 {
		tip = attempt.Fee.GasTipCap.ToInt()
	}
	gas := meta.GasLimit.ToInt()
	if !gas.IsUint64() {
		return fmt.Errorf("gas value does not fit in uint64: %s", gas)
	}
	dynamicTx := evmtypes.DynamicFeeTx{
		Nonce:     *tx.Nonce,
		To:        &meta.ToAddress,
		Gas:       gas.Uint64(),
		GasTipCap: tip,
		GasFeeCap: meta.MaxFeePerGas.ToInt(),
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
