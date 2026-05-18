package txm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	maxBumpThreshold   = 5              // maxBumpThreshold controls the maximum number of bumps for an attempt.
	hederaWeiToTinybar = 10_000_000_001 // hederaWeiToTinybar is the minimum allowed value for a transfer in Hedera plus 1. Hedera uses HBAR instead of ETH
)

type attemptBuilder struct {
	gas.EvmFeeEstimator
	priceMaxKey         func(common.Address) *assets.Wei
	keystore            keys.TxSigner
	emptyTxLimitDefault uint64
	feeBoost            bool
	chainType           chaintype.ChainType
}

func NewAttemptBuilder(priceMaxKey func(common.Address) *assets.Wei, estimator gas.EvmFeeEstimator, keystore keys.TxSigner, emptyTxLimitDefault uint64, feeBoost bool, chainType chaintype.ChainType) *attemptBuilder {
	return &attemptBuilder{
		priceMaxKey:         priceMaxKey,
		EvmFeeEstimator:     estimator,
		keystore:            keystore,
		emptyTxLimitDefault: emptyTxLimitDefault,
		feeBoost:            feeBoost,
		chainType:           chainType,
	}
}

func (a *attemptBuilder) NewAttempt(ctx context.Context, lggr logger.Logger, tx *types.Transaction, dynamic bool) (*types.Attempt, error) {
	var fee gas.EvmFee
	var estimatedGasLimit uint64
	var err error
	if tx.IsPurgeable || a.feeBoost {
		gasLimit := tx.SpecifiedGasLimit
		if tx.IsPurgeable {
			gasLimit = a.emptyTxLimitDefault
		}
		fee, estimatedGasLimit, err = a.GetMaxFee(ctx, tx.Data, gasLimit, a.priceMaxKey(tx.FromAddress), &tx.FromAddress, &tx.ToAddress)
		if err != nil {
			return nil, err
		}
	} else {
		fee, estimatedGasLimit, err = a.GetFee(ctx, tx.Data, tx.SpecifiedGasLimit, a.priceMaxKey(tx.FromAddress), &tx.FromAddress, &tx.ToAddress)
		if err != nil {
			return nil, err
		}
	}
	txType := evmtypes.LegacyTxType
	if dynamic {
		txType = evmtypes.DynamicFeeTxType
	}
	return a.newCustomAttempt(ctx, tx, fee, estimatedGasLimit, byte(txType), lggr)
}

func (a *attemptBuilder) NewBumpAttempt(ctx context.Context, lggr logger.Logger, tx *types.Transaction, previousAttempt types.Attempt) (*types.Attempt, error) {
	gasLimit := tx.SpecifiedGasLimit
	if tx.IsPurgeable {
		gasLimit = a.emptyTxLimitDefault
	}
	bumpedFee, bumpedFeeLimit, err := a.EvmFeeEstimator.BumpFee(ctx, previousAttempt.Fee, gasLimit, a.priceMaxKey(tx.FromAddress), nil)
	if err != nil {
		return nil, err
	}
	return a.newCustomAttempt(ctx, tx, bumpedFee, bumpedFeeLimit, previousAttempt.Type, lggr)
}

func (a *attemptBuilder) NewAgnosticBumpAttempt(ctx context.Context, lggr logger.Logger, tx *types.Transaction, dynamic bool) (attempt *types.Attempt, err error) {
	if a.chainType == chaintype.ChainHedera {
		return a.newHederaAttempt(ctx, lggr, tx, a.priceMaxKey(tx.FromAddress), dynamic)
	}
	// if the transaction is purgeable or feeBoost is enabled, NewAttempt will return the max fee instantly, so there is no need to bump
	attempt, err = a.NewAttempt(ctx, lggr, tx, dynamic)
	if err != nil {
		return
	}

	if tx.IsPurgeable || a.feeBoost {
		return
	}

	bumps := min(maxBumpThreshold, tx.AttemptCount)
	for range bumps {
		bumpedAttempt, err := a.NewBumpAttempt(ctx, lggr, tx, *attempt)
		if err != nil {
			lggr.Errorf("error bumping attempt: %v for txID: %v", err, tx.ID)
			return attempt, nil
		}
		attempt = bumpedAttempt
	}

	return attempt, nil
}

// newHederaAttempt is used to build a new attempt for Hedera.
// Hedera is a special case. It doesn't have a mempool but can reject an attempt for unknown reasons, even though the RPC returns success.
// The network binds transactions with unique IDs and a timestamp. If the timestamp exceeds a threshold it will auto-reject the
// transaction no matter how many times we retry. To bypass this case, we fetch a new market price and bump the fee by 1 per attempt
// to forcefully generate a new hash. We avoid max pricing purgeable transactions for the same reason.
func (a *attemptBuilder) newHederaAttempt(ctx context.Context, lggr logger.Logger, tx *types.Transaction, maxPrice *assets.Wei, dynamic bool) (*types.Attempt, error) {
	gasLimit := tx.SpecifiedGasLimit
	if tx.IsPurgeable {
		gasLimit = a.emptyTxLimitDefault
	}
	fee, estimatedGasLimit, err := a.GetFee(ctx, tx.Data, gasLimit, maxPrice, &tx.FromAddress, &tx.ToAddress)
	if err != nil {
		return nil, err
	}
	txType := evmtypes.LegacyTxType
	if dynamic {
		txType = evmtypes.DynamicFeeTxType
	}

	attempt, err := a.newCustomAttempt(ctx, tx, fee, estimatedGasLimit, byte(txType), lggr)
	if err != nil {
		return nil, err
	}
	for range tx.AttemptCount {
		if attempt.Fee.ValidDynamic() && maxPrice.Cmp(attempt.Fee.GasFeeCap) > 0 {
			fee.GasFeeCap = attempt.Fee.GasFeeCap.Add(assets.NewWeiI(1)) // Hedera doesn't have a mempool so maxPriorityFeePerGas is always 0.
		} else if attempt.Fee.GasPrice != nil && maxPrice.Cmp(attempt.Fee.GasPrice) > 0 {
			fee.GasPrice = attempt.Fee.GasPrice.Add(assets.NewWeiI(1))
		} else {
			break
		}
		attempt, err = a.newCustomAttempt(ctx, tx, fee, estimatedGasLimit, byte(txType), lggr)
		if err != nil {
			return nil, err
		}
	}
	return attempt, nil
}

func (a *attemptBuilder) newCustomAttempt(
	ctx context.Context,
	tx *types.Transaction,
	fee gas.EvmFee,
	estimatedGasLimit uint64,
	txType byte,
	lggr logger.Logger,
) (attempt *types.Attempt, err error) {
	switch txType {
	case 0x0:
		if fee.GasPrice == nil {
			err = fmt.Errorf("tried to create attempt of type %v for txID: %v but estimator did not return legacy fee", txType, tx.ID)
			logger.Sugared(lggr).AssumptionViolation(err.Error())
			return
		}
		return a.newLegacyAttempt(ctx, tx, fee.GasPrice, estimatedGasLimit)
	case 0x2:
		if !fee.ValidDynamic() {
			err = fmt.Errorf("tried to create attempt of type %v for txID: %v but estimator did not return dynamic fee", txType, tx.ID)
			logger.Sugared(lggr).AssumptionViolation(err.Error())
			return
		}
		return a.newDynamicFeeAttempt(ctx, tx, fee.DynamicFee, estimatedGasLimit)
	default:
		return nil, fmt.Errorf("cannot build attempt, unrecognized transaction type: %v", txType)
	}
}

func (a *attemptBuilder) newLegacyAttempt(ctx context.Context, tx *types.Transaction, gasPrice *assets.Wei, estimatedGasLimit uint64) (*types.Attempt, error) {
	var data []byte
	var toAddress common.Address
	value := big.NewInt(0)
	if !tx.IsPurgeable {
		data = tx.Data
		toAddress = tx.ToAddress
		value = tx.Value
	}
	if a.chainType == chaintype.ChainHedera && tx.IsPurgeable {
		value = big.NewInt(hederaWeiToTinybar)
		toAddress = tx.FromAddress
	}
	if tx.Nonce == nil {
		return nil, fmt.Errorf("failed to create attempt for txID: %v: nonce empty", tx.ID)
	}
	legacyTx := evmtypes.LegacyTx{
		Nonce:    *tx.Nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      estimatedGasLimit,
		GasPrice: gasPrice.ToInt(),
		Data:     data,
	}

	signedTx, err := a.keystore.SignTx(ctx, tx.FromAddress, evmtypes.NewTx(&legacyTx))
	if err != nil {
		return nil, fmt.Errorf("failed to sign attempt for txID: %v, err: %w", tx.ID, err)
	}

	attempt := &types.Attempt{
		TxID:              tx.ID,
		Fee:               gas.EvmFee{GasPrice: gasPrice},
		Hash:              signedTx.Hash(),
		GasLimit:          estimatedGasLimit,
		Type:              evmtypes.LegacyTxType,
		SignedTransaction: signedTx,
	}

	return attempt, nil
}

func (a *attemptBuilder) newDynamicFeeAttempt(ctx context.Context, tx *types.Transaction, dynamicFee gas.DynamicFee, estimatedGasLimit uint64) (*types.Attempt, error) {
	var data []byte
	var toAddress common.Address
	value := big.NewInt(0)
	if !tx.IsPurgeable {
		data = tx.Data
		toAddress = tx.ToAddress
		value = tx.Value
	}
	if a.chainType == chaintype.ChainHedera && tx.IsPurgeable {
		value = big.NewInt(10_000_000_001)
		toAddress = tx.FromAddress
	}
	if tx.Nonce == nil {
		return nil, fmt.Errorf("failed to create attempt for txID: %v: nonce empty", tx.ID)
	}
	dynamicTx := evmtypes.DynamicFeeTx{
		Nonce:     *tx.Nonce,
		To:        &toAddress,
		Value:     value,
		Gas:       estimatedGasLimit,
		GasFeeCap: dynamicFee.GasFeeCap.ToInt(),
		GasTipCap: dynamicFee.GasTipCap.ToInt(),
		Data:      data,
	}

	signedTx, err := a.keystore.SignTx(ctx, tx.FromAddress, evmtypes.NewTx(&dynamicTx))
	if err != nil {
		return nil, fmt.Errorf("failed to sign attempt for txID: %v, err: %w", tx.ID, err)
	}

	attempt := &types.Attempt{
		TxID:              tx.ID,
		Fee:               gas.EvmFee{DynamicFee: gas.DynamicFee{GasFeeCap: dynamicFee.GasFeeCap, GasTipCap: dynamicFee.GasTipCap}},
		Hash:              signedTx.Hash(),
		GasLimit:          estimatedGasLimit,
		Type:              evmtypes.DynamicFeeTxType,
		SignedTransaction: signedTx,
	}

	return attempt, nil
}
