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

	"github.com/mitchellh/mapstructure"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	timeout = time.Second * 5
	metaABI = `[
  {
    "type": "function",
    "name": "metacall",
    "stateMutability": "payable",
    "inputs": [
      {
        "name": "userOp",
        "type": "tuple",
        "components": [
          { "name": "from", "type": "address" },
          { "name": "to", "type": "address" },
          { "name": "value", "type": "uint256" },
          { "name": "gas", "type": "uint256" },
          { "name": "maxFeePerGas", "type": "uint256" },
          { "name": "nonce", "type": "uint256" },
          { "name": "deadline", "type": "uint256" },
          { "name": "dapp", "type": "address" },
          { "name": "control", "type": "address" },
          { "name": "callConfig", "type": "uint32" },
          { "name": "dappGasLimit", "type": "uint32" },
          { "name": "solverGasLimit", "type": "uint32" },
          { "name": "bundlerSurchargeRate", "type": "uint24" },
          { "name": "sessionKey", "type": "address" },
          { "name": "data", "type": "bytes" },
          { "name": "signature", "type": "bytes" }
        ]
      },
      {
        "name": "solverOps",
        "type": "tuple[]",
        "components": [
          { "name": "from", "type": "address" },
          { "name": "to", "type": "address" },
          { "name": "value", "type": "uint256" },
          { "name": "gas", "type": "uint256" },
          { "name": "maxFeePerGas", "type": "uint256" },
          { "name": "deadline", "type": "uint256" },
          { "name": "solver", "type": "address" },
          { "name": "control", "type": "address" },
          { "name": "userOpHash", "type": "bytes32" },
          { "name": "bidToken", "type": "address" },
          { "name": "bidAmount", "type": "uint256" },
          { "name": "data", "type": "bytes" },
          { "name": "signature", "type": "bytes" }
        ]
      },
      {
        "name": "dAppOp",
        "type": "tuple",
        "components": [
          { "name": "from", "type": "address" },
          { "name": "to", "type": "address" },
          { "name": "nonce", "type": "uint256" },
          { "name": "deadline", "type": "uint256" },
          { "name": "control", "type": "address" },
          { "name": "bundler", "type": "address" },
          { "name": "userOpHash", "type": "bytes32" },
          { "name": "callChainHash", "type": "bytes32" },
          { "name": "signature", "type": "bytes" }
        ]
      },
      { "name": "gasRefundBeneficiary", "type": "address" }
    ],
    "outputs": [ { "name": "auctionWon", "type": "bool" } ]
  }
]`
	ABI = `[
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

var _ txm.Client = &MetaClient{}

type MetaClientKeystore interface {
	SignMessage(ctx context.Context, address common.Address, data []byte) ([]byte, error)
	SignTx(ctx context.Context, fromAddress common.Address, tx *evmtypes.Transaction) (*evmtypes.Transaction, error)
}

type MetaClientRPC interface {
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	PendingNonceAt(context.Context, common.Address) (uint64, error)
	SendTransaction(context.Context, *evmtypes.Transaction) error
}

type MetaClient struct {
	lggr      logger.SugaredLogger
	c         MetaClientRPC
	ks        MetaClientKeystore
	customURL *url.URL
	chainID   *big.Int
	metrics   *MetaMetrics
}

func NewMetaClient(lggr logger.Logger, c MetaClientRPC, ks MetaClientKeystore, customURL *url.URL, chainID *big.Int) (*MetaClient, error) {
	metrics, err := NewMetaMetrics(chainID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create Meta metrics: %w", err)
	}

	return &MetaClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.Txm.MetaClient")),
		c:         c,
		ks:        ks,
		customURL: customURL,
		chainID:   chainID,
		metrics:   metrics,
	}, nil
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
		meta, err := a.SendRequest(ctx, tx, attempt, *meta.DualBroadcastParams, tx.ToAddress)
		if err != nil {
			a.metrics.RecordSendRequestError(ctx)
			return fmt.Errorf("error sending request for transactionID(%d): %w", tx.ID, err)
		}
		if meta != nil {
			if err := a.SendOperation(ctx, tx, attempt, *meta); err != nil {
				a.metrics.RecordSendOperationError(ctx)
				return fmt.Errorf("failed to send operation for transactionID(%d): %w", tx.ID, err)
			}
			return nil
		}
		a.lggr.Infof("No bids for transactionID(%d): ", tx.ID)
		return nil
	}
	a.lggr.Infow("Broadcasting attempt to public mempool", "tx", tx)
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
	UO  *UORaw `json:"userOperation,omitempty"`
	SOS []*SO  `json:"solverOperations,omitempty"`
	DO  *DO    `json:"dAppOperation,omitempty"`
	MetacalldataResponse
}

type MetacalldataResponse struct {
	ToAddress    common.Address `json:"metacallDestination"`
	GasLimit     *hexutil.Big   `json:"metacallGasLimit"`
	MaxFeePerGas *hexutil.Big   `json:"metacallMaxFeePerGas"`
	CallData     hexutil.Bytes  `json:"metacallCallData"`
}

type UO struct {
	To           common.Address
	MaxFeePerGas *big.Int
	Dapp         common.Address
	Control      common.Address
	Data         []byte
}

type UORaw struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Nonce        *hexutil.Big   `json:"nonce"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	CallConfig   *hexutil.Big   `json:"callConfig"`
	DappGasLimit *hexutil.Big   `json:"dappGasLimit,omitempty"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

type SO struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Solver       common.Address `json:"solver"`
	Control      common.Address `json:"control"`
	UserOpHash   common.Hash    `json:"userOpHash"`
	BidToken     common.Address `json:"bidToken"`
	BidAmount    *hexutil.Big   `json:"bidAmount"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

type DO struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Nonce         *hexutil.Big   `json:"nonce"`
	Deadline      *hexutil.Big   `json:"deadline"`
	Control       common.Address `json:"control"`
	Bundler       common.Address `json:"bundler"`
	UserOpHash    common.Hash    `json:"userOpHash"`
	CallChainHash common.Hash    `json:"callChainHash"`
	Signature     hexutil.Bytes  `json:"signature"`
}

type Metacalldata struct {
	UOP                  UO
	SOPs                 []SO
	DOP                  DO
	GasRefundBeneficiary common.Address
}

func (a *MetaClient) SendRequest(parentCtx context.Context, tx *types.Transaction, attempt *types.Attempt, dualBroadcastParams string, fwdrDestAddress common.Address) (*MetacalldataResponse, error) {
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

	// Start timing for endpoint latency measurement
	startTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
	latency := time.Since(startTime)

	// Record latency
	a.metrics.RecordLatency(ctx, latency)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	// Record status code
	a.metrics.RecordStatusCode(ctx, resp.StatusCode)

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
		if strings.Contains(response.Error.ErrorMessage, "no solver operations received") {
			a.metrics.RecordBidsReceived(ctx, 0)
			return nil, nil
		}
		return nil, errors.New(response.Error.ErrorMessage)
	}

	if response.Result == nil {
		return nil, nil
	}

	// Record bid count (number of solver operations received)
	a.metrics.RecordBidsReceived(ctx, len(response.Result.SOS))

	if r, err := json.MarshalIndent(response.Result, "", "  "); err == nil {
		a.lggr.Info("Response: ", string(r))
	}

	return VerifyResponse(response.Result.MetacalldataResponse, dualBroadcastParams, tx.Data, tx.FromAddress, fwdrDestAddress)
}

func VerifyResponse(metacalldata MetacalldataResponse, dualBroadcastParams string, txData []byte, fromAddress common.Address, fwdrDestAddress common.Address) (*MetacalldataResponse, error) {
	params, err := url.ParseQuery(dualBroadcastParams)
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

	if metacalldata.ToAddress != to {
		return nil, fmt.Errorf("incorrect metacall: metacall.ToAddress: %v, to: %v",
			metacalldata.ToAddress, to)
	}

	// Calldata verification
	if len(metacalldata.CallData) < 4 {
		return nil, errors.New("calldata too short")
	}

	mABI, err := abi.JSON(strings.NewReader(metaABI))
	if err != nil {
		return nil, fmt.Errorf("parse ABI: %w", err)
	}
	method, ok := mABI.Methods["metacall"]
	if !ok {
		return nil, errors.New("method not found in ABI")
	}

	if !bytes.Equal(metacalldata.CallData[:4], method.ID) {
		return nil, fmt.Errorf("selector mismatch: got %x, expected %x", metacalldata.CallData[:4], method.ID)
	}
	args, err := method.Inputs.Unpack(metacalldata.CallData[4:])
	if err != nil || len(args) != 4 {
		return nil, fmt.Errorf("unpack error(argsLen:%d): %w", len(args), err)
	}

	result := new(Metacalldata)
	err = mapstructure.Decode(args[0], &result.UOP)
	if err != nil {
		return nil, fmt.Errorf("error unpacking UOP: %w", err)
	}
	err = mapstructure.Decode(args[1], &result.SOPs)
	if err != nil {
		return nil, fmt.Errorf("error unpacking SOPs: %w", err)
	}
	err = mapstructure.Decode(args[2], &result.DOP)
	if err != nil {
		return nil, fmt.Errorf("error unpacking DOP: %w", err)
	}
	return VerifyMetadata(txData, fromAddress, *result, fwdrDestAddress, dApp, to, metacalldata)
}

func VerifyMetadata(txData []byte, fromAddress common.Address, result Metacalldata, fwdrDestAddress common.Address, dApp common.Address, to common.Address, metacalldata MetacalldataResponse) (*MetacalldataResponse, error) {
	abi, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, fmt.Errorf("couldn't read ABI: %w", err)
	}

	updateFn, ok := abi.Methods["update"]
	if !ok {
		return nil, errors.New("update method not found in ABI")
	}
	if len(result.UOP.Data) < 4 || !bytes.HasPrefix(result.UOP.Data, updateFn.ID) {
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

	// DOP
	if result.DOP.To != to || result.DOP.Control != dApp || result.DOP.Bundler != fromAddress {
		return nil, fmt.Errorf("incorrect DOP: dop.To: %v, dop.Control: %v, dop.Bundler: %v, to: %v, dApp: %v, fromAddress: %v",
			result.DOP.To, result.DOP.Control, result.DOP.Bundler, to, dApp, fromAddress)
	}

	// SOP
	atLeastOne := false
	for _, sop := range result.SOPs {
		if sop.To != to || sop.Control != dApp {
			// Exit early
			return nil, fmt.Errorf("incorrect SOP: sop.To: %v, sop.Control: %v, to: %v, dApp: %v", sop.To, sop.Control, to, dApp)
		}
		atLeastOne = true
	}
	if !atLeastOne {
		return nil, nil
	}

	// UOP
	if result.UOP.To != to ||
		result.UOP.MaxFeePerGas == nil || metacalldata.MaxFeePerGas == nil || result.UOP.MaxFeePerGas.Cmp(metacalldata.MaxFeePerGas.ToInt()) != 0 ||
		result.UOP.Dapp != dApp ||
		result.UOP.Control != dApp ||
		destinationAddress != fwdrDestAddress || !bytes.Equal(updateCalldata, txData) {
		return nil, fmt.Errorf("incorrect UOP: uop.To: %v, uop.MaxFeePerGas: %v, uop.Dapp: %v, uop.update.destinationAddress: %v, uop.update.calldata: %v, to: %v, metacall.MaxFeePerGas: %v, dApp: %v, fwdrDestAddress: %v, txData: %v",
			result.UOP.To, result.UOP.MaxFeePerGas, result.UOP.Dapp, destinationAddress, updateCalldata, to, metacalldata.MaxFeePerGas, dApp, fwdrDestAddress, txData)
	}

	return &metacalldata, nil
}

func (a *MetaClient) SendOperation(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, meta MetacalldataResponse) error {
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
