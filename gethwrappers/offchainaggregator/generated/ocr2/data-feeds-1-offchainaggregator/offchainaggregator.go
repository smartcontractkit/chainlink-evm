// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datafeeds1offchainaggregator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OffchainAggregatorMetaData contains all meta data concerning the OffchainAggregator contract.
var OffchainAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_link\",\"type\":\"address\",\"internalType\":\"contractLinkTokenInterface\"},{\"name\":\"_minAnswer\",\"type\":\"int192\",\"internalType\":\"int192\"},{\"name\":\"_maxAnswer\",\"type\":\"int192\",\"internalType\":\"int192\"},{\"name\":\"_billingAccessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"},{\"name\":\"_requesterAccessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptPayeeship\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAnswer\",\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBilling\",\"inputs\":[],\"outputs\":[{\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountingGas\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBillingAccessController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkToken\",\"inputs\":[],\"outputs\":[{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"contractLinkTokenInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequesterAccessController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoundData\",\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransmitters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"contractAggregatorValidatorInterface\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestAnswer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestConfigDigestAndEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"scanLogs\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTransmissionDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"round\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"latestAnswer_\",\"type\":\"int192\",\"internalType\":\"int192\"},{\"name\":\"latestTimestamp_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkAvailableForPayment\",\"inputs\":[],\"outputs\":[{\"name\":\"availableBalance\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxAnswer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int192\",\"internalType\":\"int192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minAnswer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int192\",\"internalType\":\"int192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleObservationCount\",\"inputs\":[{\"name\":\"transmitterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owedPayment\",\"inputs\":[{\"name\":\"transmitterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestNewRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBilling\",\"inputs\":[{\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"accountingGas\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBillingAccessController\",\"inputs\":[{\"name\":\"_billingAccessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLinkToken\",\"inputs\":[{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"contractLinkTokenInterface\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayees\",\"inputs\":[{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"payees\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequesterAccessController\",\"inputs\":[{\"name\":\"requesterAccessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorConfig\",\"inputs\":[{\"name\":\"newValidator\",\"type\":\"address\",\"internalType\":\"contractAggregatorValidatorInterface\"},{\"name\":\"newGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferPayeeship\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transmit\",\"inputs\":[{\"name\":\"reportContext\",\"type\":\"bytes32[3]\",\"internalType\":\"bytes32[3]\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"ss\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"rawVs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFunds\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawPayment\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AnswerUpdated\",\"inputs\":[{\"name\":\"current\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BillingAccessControllerSet\",\"inputs\":[{\"name\":\"old\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BillingSet\",\"inputs\":[{\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"accountingGas\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"configCount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LinkTokenSet\",\"inputs\":[{\"name\":\"oldLinkToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\"},{\"name\":\"newLinkToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"startedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTransmission\",\"inputs\":[{\"name\":\"aggregatorRoundId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"answer\",\"type\":\"int192\",\"indexed\":false,\"internalType\":\"int192\"},{\"name\":\"transmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"observationsTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"observations\",\"type\":\"int192[]\",\"indexed\":false,\"internalType\":\"int192[]\"},{\"name\":\"observers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"juelsPerFeeCoin\",\"type\":\"int192\",\"indexed\":false,\"internalType\":\"int192\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"epochAndRound\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OraclePaid\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"linkToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayeeshipTransferRequested\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposed\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayeeshipTransferred\",\"inputs\":[{\"name\":\"transmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"previous\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequesterAccessControllerSet\",\"inputs\":[{\"name\":\"old\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundRequested\",\"inputs\":[{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"round\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transmitted\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorConfigSet\",\"inputs\":[{\"name\":\"previousValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\"},{\"name\":\"previousGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"currentValidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractAggregatorValidatorInterface\"},{\"name\":\"currentGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005edf38038062005edf8339810160408190526200003491620004a4565b868686868686863380600081620000925760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c557620000c58162000179565b5050601180546001600160a01b0319166001600160a01b038a169081179091556040519091506000907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a908290a36200011e8462000224565b60ff821660c052601062000133828262000686565b506200013f836200029d565b6200014c60008062000318565b50505050601791820b608052900b60a05250506015805460ff191660011790555062000752945050505050565b336001600160a01b03821603620001d35760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000089565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b0390811690821681146200029957601280546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d4891291015b60405180910390a15b5050565b620002a7620003ff565b600f546001600160a01b0390811690821681146200029957600f80546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634910162000290565b62000322620003ff565b60408051808201909152600e546001600160a01b03808216808452600160a01b90920463ffffffff16602084015284161415806200037057508163ffffffff16816020015163ffffffff1614155b15620003fa576040805180820182526001600160a01b0385811680835263ffffffff8681166020948501819052600e80546001600160c01b0319168417600160a01b830217905586518786015187519316835294820152909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541910160405180910390a35b505050565b6000546001600160a01b031633146200045b5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640162000089565b565b6001600160a01b03811681146200047357600080fd5b50565b8051601781900b81146200048957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600080600080600080600060e0888a031215620004c057600080fd5b8751620004cd816200045d565b96506020620004de89820162000476565b9650620004ee60408a0162000476565b9550606089015162000500816200045d565b60808a015190955062000513816200045d565b60a08a015190945060ff811681146200052b57600080fd5b60c08a01519093506001600160401b03808211156200054957600080fd5b818b0191508b601f8301126200055e57600080fd5b8151818111156200057357620005736200048e565b604051601f8201601f19908116603f011681019083821181831017156200059e576200059e6200048e565b816040528281528e86848701011115620005b757600080fd5b600093505b82841015620005db5784840186015181850187015292850192620005bc565b600086848301015280965050505050505092959891949750929550565b600181811c908216806200060d57607f821691505b6020821081036200062e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620003fa57600081815260208120601f850160051c810160208610156200065d5750805b601f850160051c820191505b818110156200067e5782815560010162000669565b505050505050565b81516001600160401b03811115620006a257620006a26200048e565b620006ba81620006b38454620005f8565b8462000634565b602080601f831160018114620006f25760008415620006d95750858301515b600019600386901b1c1916600185901b1785556200067e565b600085815260208120601f198616915b82811015620007235788860151825594840194600190910190840162000702565b5085821015620007425787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c0516157416200079e6000396000610468015260008181610517015281816126ad0152613ccf015260008181610364015281816126850152613ca201526157416000f3fe608060405234801561001057600080fd5b50600436106103155760003560e01c80639a6fc8f5116101a7578063d09dc339116100ee578063e76d516811610097578063f2fde38b11610071578063f2fde38b1461084f578063fbffd2c114610862578063feaf968c1461087557600080fd5b8063e76d516814610818578063eb45716314610829578063eb5dcd6c1461083c57600080fd5b8063e3d0e712116100c8578063e3d0e71214610793578063e4902f82146107a6578063e5fe4577146107ce57600080fd5b8063d09dc3391461076d578063daffc4b514610775578063dc7f01241461078657600080fd5b8063b121e14711610150578063b633620c1161012a578063b633620c14610736578063c107532914610749578063c4c92b371461075c57600080fd5b8063b121e147146106fd578063b1dc65a414610710578063b5ab58dc1461072357600080fd5b80639e3ceeab116101815780639e3ceeab146106a6578063a118f249146106b9578063afcb95d7146106cc57600080fd5b80639a6fc8f5146105f75780639bd2c0b1146106415780639c849b301461069357600080fd5b8063668a0f021161026b57806381ff7048116102145780638ac28d5a116101ee5780638ac28d5a1461059c5780638da5cb5b146105af57806398e5b12a146105d457600080fd5b806381ff7048146105515780638205bf6a146105815780638823da6c1461058957600080fd5b80637284e416116102455780637284e4161461053957806379ba5097146105415780638038e4a11461054957600080fd5b8063668a0f02146104e75780636b14daf8146104ef57806370da2f671461051257600080fd5b8063313ce567116102cd57806354fd4d50116102a757806354fd4d50146104b7578063643dc105146104bf578063666cab8d146104d257600080fd5b8063313ce567146104635780634fb174701461049c57806350d25bcd146104af57600080fd5b8063181f5a77116102fe578063181f5a771461034a57806322adbc781461035f578063299372681461039957600080fd5b80630a7569831461031a5780630eafb25b14610324575b600080fd5b61032261087d565b005b610337610332366004614b17565b6108e4565b6040519081526020015b60405180910390f35b610352610a05565b6040516103419190614b7a565b6103867f000000000000000000000000000000000000000000000000000000000000000081565b60405160179190910b8152602001610341565b610427600b546a0100000000000000000000810463ffffffff908116926e010000000000000000000000000000830482169272010000000000000000000000000000000000008104831692760100000000000000000000000000000000000000000000820416917a01000000000000000000000000000000000000000000000000000090910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a001610341565b61048a7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff9091168152602001610341565b6103226104aa366004614b8d565b610a25565b610337610ca5565b610337600681565b6103226104cd366004614bd8565b610d63565b6104da61101e565b6040516103419190614c95565b610337611080565b6105026104fd366004614d78565b611127565b6040519015158152602001610341565b6103867f000000000000000000000000000000000000000000000000000000000000000081565b61035261114f565b6103226111e6565b6103226112af565b600d54600a546040805163ffffffff80851682526401000000009094049093166020840152820152606001610341565b610337611317565b610322610597366004614b17565b6113ef565b6103226105aa366004614b17565b61148f565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610341565b6105dc611501565b60405169ffffffffffffffffffff9091168152602001610341565b61060a610605366004614dc8565b611678565b6040805169ffffffffffffffffffff968716815260208101959095528401929092526060830152909116608082015260a001610341565b604080518082018252600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602092830181905283519182529181019190915201610341565b6103226106a1366004614e40565b61172a565b6103226106b4366004614b17565b611920565b6103226106c7366004614b17565b6119b7565b600a54600b546040805160008152602081019390935261010090910460081c63ffffffff1690820152606001610341565b61032261070b366004614b17565b611a51565b61032261071e366004614eac565b611b45565b610337610731366004614f91565b612090565b610337610744366004614f91565b612130565b610322610757366004614faa565b6121c8565b6012546001600160a01b03166105bc565b61033761249c565b600f546001600160a01b03166105bc565b6015546105029060ff1681565b6103226107a1366004615097565b612545565b6107b96107b4366004614b17565b612dc1565b60405163ffffffff9091168152602001610341565b6107d6612e7f565b6040805195865263ffffffff909416602086015260ff9092169284019290925260179190910b606083015267ffffffffffffffff16608082015260a001610341565b6011546001600160a01b03166105bc565b610322610837366004615164565b612f44565b61032261084a366004614b8d565b613061565b61032261085d366004614b17565b6131b1565b610322610870366004614b17565b6131c2565b61060a6131d3565b610885613308565b60155460ff16156108e257601580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f53963890600090a15b565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff16918101919091529061094b5750600092915050565b600b5460208201516000917201000000000000000000000000000000000000900463ffffffff169060069060ff16601f811061098957610989615192565b600881049190910154600b546109bf926007166004026101000a90910463ffffffff9081169166010000000000009004166151f0565b63ffffffff166109cf9190615214565b6109dd90633b9aca00615214565b905081604001516bffffffffffffffffffffffff16816109fd919061522b565b949350505050565b60606040518060600160405280602481526020016156e860249139905090565b610a2d613308565b6011546001600160a01b03908116908316819003610a4a57505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa158015610aa7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610acb919061523e565b50610ad4613362565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038316906370a0823190602401602060405180830381865afa158015610b34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b58919061523e565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018390529192509083169063a9059cbb906044016020604051808303816000875af1158015610bc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be89190615257565b610c395760405162461bcd60e51b815260206004820152601f60248201527f7472616e736665722072656d61696e696e672066756e6473206661696c65640060448201526064015b60405180910390fd5b601180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925560405190918416907f4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a90600090a350505b5050565b6000610ce8336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b610d345760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b600b546601000000000000900463ffffffff166000908152600c602052604090205460170b905090565b905090565b6012546001600160a01b0316610d816000546001600160a01b031690565b6001600160a01b0316336001600160a01b03161480610e2657506040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b03821690636b14daf890610de59033906000903690600401615279565b602060405180830381865afa158015610e02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e269190615257565b610e725760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c6044820152606401610c30565b610e7a613362565b600b80547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff166a010000000000000000000063ffffffff8981169182027fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff16929092176e010000000000000000000000000000898416908102919091177fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff1672010000000000000000000000000000000000008985169081027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1691909117760100000000000000000000000000000000000000000000948916948502177fffffff000000ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000062ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b6060600580548060200260200160405190810160405280929190818152602001828054801561107657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611058575b5050505050905090565b60006110c3336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b61110f5760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b600b546601000000000000900463ffffffff16905090565b60006111338383613746565b8061114657506001600160a01b03831632145b90505b92915050565b6060611192336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b6111de5760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b610d5e613776565b6001546001600160a01b031633146112405760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610c30565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6112b7613308565b60155460ff166108e257601580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c348090600090a1565b600061135a336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b6113a65760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b50600b5463ffffffff660100000000000090910481166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b6113f7613308565b6001600160a01b03811660009081526016602052604090205460ff161561148c576001600160a01b03811660008181526016602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d191015b60405180910390a15b50565b6001600160a01b038181166000908152601360205260409020541633146114f85760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e2077697468647261770000000000000000006044820152606401610c30565b61148c816137ff565b600080546001600160a01b03163314806115a55750600f546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf8906115649033906000903690600401615279565b602060405180830381865afa158015611581573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a59190615257565b6115f15760405162461bcd60e51b815260206004820152601d60248201527f4f6e6c79206f776e6572267265717565737465722063616e2063616c6c0000006044820152606401610c30565b600b54600a546040805191825263ffffffff6101008404600881901c8216602085015260ff811684840152915164ffffffffff9092169366010000000000009004169133917f41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c9181900360600190a261166b8160016152b8565b63ffffffff169250505090565b60008060008060006116c1336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b61170d5760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b61171686613a4b565b945094509450945094505b91939590929450565b611732613308565b8281146117815760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a656044820152606401610c30565b60005b838110156119195760008585838181106117a0576117a0615192565b90506020020160208101906117b59190614b17565b905060008484848181106117cb576117cb615192565b90506020020160208101906117e09190614b17565b6001600160a01b03808416600090815260136020526040902054919250168015808061181d5750826001600160a01b0316826001600160a01b0316145b6118695760405162461bcd60e51b815260206004820152601160248201527f706179656520616c7265616479207365740000000000000000000000000000006044820152606401610c30565b6001600160a01b03848116600090815260136020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168583169081179091559083161461190257826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080611911906152d5565b915050611784565b5050505050565b611928613308565b600f546001600160a01b039081169082168114610ca157600f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691821790925560408051928416835260208301919091527f27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae63491015b60405180910390a15050565b6119bf613308565b6001600160a01b03811660009081526016602052604090205460ff1661148c576001600160a01b03811660008181526016602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101611483565b6001600160a01b03818116600090815260146020526040902054163314611aba5760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e20616363657074006044820152606401610c30565b6001600160a01b0381811660008181526013602090815260408083208054337fffffffffffffffffffffffff000000000000000000000000000000000000000080831682179093556014909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600b5460ff8116835290810464ffffffffff90811660208085018290526601000000000000840463ffffffff908116968601969096526a01000000000000000000008404861660608601526e01000000000000000000000000000084048616608086015272010000000000000000000000000000000000008404861660a0860152760100000000000000000000000000000000000000000000840490951660c08501527a01000000000000000000000000000000000000000000000000000090920462ffffff1660e08401529394509092918c013591821611611c795760405162461bcd60e51b815260206004820152600c60248201527f7374616c65207265706f727400000000000000000000000000000000000000006044820152606401610c30565b3360009081526002602052604090205460ff16611cd85760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c30565b600a548b3514611d2a5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c30565b611d388a8a8a8a8a8a613b05565b8151611d4590600161530d565b60ff168714611d965760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c30565b868514611de55760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c30565b60008a8a604051611df7929190615326565b604051908190038120611e0e918e90602001615336565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a811015611fb45760006001858a8460208110611e5d57611e5d615192565b611e6a91901a601b61530d565b8f8f86818110611e7c57611e7c615192565b905060200201358e8e87818110611e9557611e95615192565b9050602002013560405160008152602001604052604051611ed2949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611ef4573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526003602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050611f8d5760405162461bcd60e51b815260206004820152600f60248201527f7369676e6174757265206572726f7200000000000000000000000000000000006044820152606401610c30565b826020015160080260ff166001901b84019350508080611fac906152d5565b915050611e3e565b5081827e0101010101010101010101010101010101010101010101010101010101010116146120255760405162461bcd60e51b815260206004820152601060248201527f6475706c6963617465207369676e6572000000000000000000000000000000006044820152606401610c30565b50600091506120749050838d836020020135848e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613ba292505050565b90506120828382863361409f565b505050505050505050505050565b60006120d3336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b61211f5760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b612128826141ea565b90505b919050565b6000612173336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b6121bf5760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b6121288261421c565b6000546001600160a01b031633148061226b57506012546040517f6b14daf80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636b14daf89061222a9033906000903690600401615279565b602060405180830381865afa158015612247573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061226b9190615257565b6122b75760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c6044820152606401610c30565b60006122c161426e565b6011546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529192506000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015612328573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061234c919061523e565b90508181101561239e5760405162461bcd60e51b815260206004820152601460248201527f696e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610c30565b6011546001600160a01b031663a9059cbb856123c36123bd868661534c565b8761444f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015612426573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061244a9190615257565b6124965760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742066756e647300000000000000000000000000006044820152606401610c30565b50505050565b6011546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015612502573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612526919061523e565b9050600061253261426e565b905061253e818361535f565b9250505090565b61254d613308565b601f8651111561259f5760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79206f7261636c6573000000000000000000000000000000006044820152606401610c30565b84518651146125f05760405162461bcd60e51b815260206004820152601660248201527f6f7261636c65206c656e677468206d69736d61746368000000000000000000006044820152606401610c30565b85516125fd85600361537f565b60ff161061264d5760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c30565b6126598460ff16614466565b604080517f010000000000000000000000000000000000000000000000000000000000000060208201527f0000000000000000000000000000000000000000000000000000000000000000821b60218201527f000000000000000000000000000000000000000000000000000000000000000090911b6039820152605101604051602081830303815290604052805190602001208380519060200120146127425760405162461bcd60e51b815260206004820152601560248201527f696e76616c6964206f6e636861696e436f6e66696700000000000000000000006044820152606401610c30565b6040805160c0810182528781526020810187905260ff8616918101919091526060810184905267ffffffffffffffff8316608082015260a08101829052600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000ff1690556127af613362565b60045460005b8181101561288e576000600482815481106127d2576127d2615192565b6000918252602082200154600580546001600160a01b03909216935090849081106127ff576127ff615192565b60009182526020808320909101546001600160a01b03948516835260038252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905594168252600290529190912080547fffffffffffffffffffffffffffffffffffff00000000000000000000000000001690555080612886816152d5565b9150506127b5565b5061289b600460006149bf565b6128a7600560006149bf565b60005b825151811015612bb45760036000846000015183815181106128ce576128ce615192565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156129425760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c30565b604080518082019091526001815260ff82166020820152835180516003916000918590811061297357612973615192565b6020908102919091018101516001600160a01b0316825281810192909252604001600090812083518154948401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009095169015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff161761010060ff90951694909402939093179092558401518051600292919084908110612a1857612a18615192565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615612a8c5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c30565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff168152506002600085602001518481518110612ad657612ad6615192565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516bffffffffffffffffffffffff1662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff60ff95909516610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff931515939093167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090941693909317919091179290921617905580612bac816152d5565b9150506128aa565b5081518051612bcb916004916020909101906149dd565b506020808301518051612be29260059201906149dd565b506040820151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612c749286929082169116176152b8565b92506101000a81548163ffffffff021916908363ffffffff160217905550612cd34630600d60009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a001516144b6565b600a819055600d5484516020860151604080880151606089015160808a015160a08b015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612d3c988b98919763ffffffff90911696919590949093909290919061539b565b60405180910390a1600b546601000000000000900463ffffffff1660005b845151811015612db45781600682601f8110612d7857612d78615192565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080612dac906152d5565b915050612d5a565b5050505050505050505050565b6001600160a01b03811660009081526002602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046bffffffffffffffffffffffff169181019190915290612e285750600092915050565b6006816020015160ff16601f8110612e4257612e42615192565b600881049190910154600b54612e78926007166004026101000a90910463ffffffff9081169166010000000000009004166151f0565b9392505050565b600080808080333214612ed45760405162461bcd60e51b815260206004820152601460248201527f4f6e6c792063616c6c61626c6520627920454f410000000000000000000000006044820152606401610c30565b5050600a54600b5463ffffffff6601000000000000820481166000908152600c60205260409020549296610100909204600881901c8216965064ffffffffff169450601783900b93507c010000000000000000000000000000000000000000000000000000000090920490911690565b612f4c613308565b60408051808201909152600e546001600160a01b038082168084527401000000000000000000000000000000000000000090920463ffffffff1660208401528416141580612faa57508163ffffffff16816020015163ffffffff1614155b1561305c576040805180820182526001600160a01b0385811680835263ffffffff8681166020948501819052600e80547fffffffffffffffff00000000000000000000000000000000000000000000000016841774010000000000000000000000000000000000000000830217905586518786015187519316835294820152909392909116917fb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541910160405180910390a35b505050565b6001600160a01b038281166000908152601360205260409020541633146130ca5760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e207570646174650000006044820152606401610c30565b6001600160a01b03811633036131225760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c30565b6001600160a01b03808316600090815260146020526040902080548383167fffffffffffffffffffffffff00000000000000000000000000000000000000008216811790925590911690811461305c576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b6131b9613308565b61148c81614544565b6131ca613308565b61148c81614605565b600080600080600061321c336000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112792505050565b6132685760405162461bcd60e51b815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610c30565b6132f7600b546601000000000000900463ffffffff9081166000818152600c602090815260409182902082516060810184529054601781900b8083527801000000000000000000000000000000000000000000000000820487169383018490527c0100000000000000000000000000000000000000000000000000000000909104909516920182905291938490565b945094509450945094509091929394565b6000546001600160a01b031633146108e25760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c30565b601154600b54604080516103e08101918290526001600160a01b0390931692660100000000000090920463ffffffff1691600091600690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116133a0579050505050505090506000600580548060200260200160405190810160405280929190818152602001828054801561343b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161341d575b5050505050905060005b81518110156137385760006002600084848151811061346657613466615192565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff1690506000600260008585815181106134d2576134d2615192565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060008483601f811061353f5761353f615192565b6020020151600b5490870363ffffffff90811692507201000000000000000000000000000000000000909104168102633b9aca00028201801561372d5760006013600087878151811061359457613594615192565b6020908102919091018101516001600160a01b0390811683529082019290925260409081016000205490517fa9059cbb00000000000000000000000000000000000000000000000000000000815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af1158015613621573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136459190615257565b6136915760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742066756e647300000000000000000000000000006044820152606401610c30565b878786601f81106136a4576136a4615192565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b03168787815181106136e1576136e1615192565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8560405161372391815260200190565b60405180910390a4505b505050600101613445565b50611919600683601f614a5a565b6001600160a01b03821660009081526016602052604081205460ff168061114657505060155460ff161592915050565b60606010805461378590615431565b80601f01602080910402602001604051908101604052809291908181526020018280546137b190615431565b80156110765780601f106137d357610100808354040283529160200191611076565b820191906000526020600020905b8154815290600101906020018083116137e157509395945050505050565b6001600160a01b0381166000908152600260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046bffffffffffffffffffffffff1692810192909252613861575050565b600061386c836108e4565b9050801561305c576001600160a01b03838116600090815260136020526040908190205460115491517fa9059cbb000000000000000000000000000000000000000000000000000000008152908316600482018190526024820185905292919091169063a9059cbb906044016020604051808303816000875af11580156138f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061391b9190615257565b6139675760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742066756e647300000000000000000000000000006044820152606401610c30565b600b60000160069054906101000a900463ffffffff166006846020015160ff16601f811061399757613997615192565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b0384811660008181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff169055601154915186815291841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b60008080808063ffffffff69ffffffffffffffffffff87161115613a7d57506000935083925082915081905080611721565b5050505063ffffffff8281166000908152600c602090815260409182902082516060810184529054601781900b8083527801000000000000000000000000000000000000000000000000820486169383018490527c01000000000000000000000000000000000000000000000000000000009091049094169201829052939491939250908490565b6000613b12826020615214565b613b1d856020615214565b613b298861014461522b565b613b33919061522b565b613b3d919061522b565b613b4890600061522b565b9050368114613b995760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c30565b50505050505050565b600080613bae8361468c565b9050601f8160400151511115613c065760405162461bcd60e51b815260206004820152601e60248201527f6e756d206f62736572766174696f6e73206f7574206f6620626f756e647300006044820152606401610c30565b604081015151865160ff1610613c5e5760405162461bcd60e51b815260206004820152601e60248201527f746f6f206665772076616c75657320746f207472757374206d656469616e00006044820152606401610c30565b64ffffffffff841660208701526040810151805160009190613c829060029061547e565b81518110613c9257613c92615192565b602002602001015190508060170b7f000000000000000000000000000000000000000000000000000000000000000060170b13158015613cf857507f000000000000000000000000000000000000000000000000000000000000000060170b8160170b13155b613d445760405162461bcd60e51b815260206004820152601e60248201527f6d656469616e206973206f7574206f66206d696e2d6d61782072616e676500006044820152606401610c30565b60408701805190613d54826154b9565b63ffffffff1663ffffffff168152505060405180606001604052808260170b8152602001836000015163ffffffff1681526020014263ffffffff16815250600c6000896040015163ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a81548177ffffffffffffffffffffffffffffffffffffffffffffffff021916908360170b77ffffffffffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555090505086600b60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050866040015163ffffffff167fc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a823385600001518660400151876020015188606001518d8d604051613fe89897969594939291906154dc565b60405180910390a26040808801518351915163ffffffff9283168152600092909116907f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac602719060200160405180910390a3866040015163ffffffff168160170b7f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f4260405161407891815260200190565b60405180910390a361409187604001518260170b614731565b506060015195945050505050565b60008360170b126124965760006140d1633b9aca003a04866080015163ffffffff16876060015163ffffffff16614881565b90506010360260005a905060006140fa8663ffffffff1685858b60e0015162ffffff16866148a7565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526002602052604090205460c08c01519290910492506201000090046bffffffffffffffffffffffff9081169163ffffffff16633b9aca0002828401019081168211156141855750505050505050612496565b6001600160a01b038816600090815260026020526040902080546bffffffffffffffffffffffff90921662010000027fffffffffffffffffffffffffffffffffffff000000000000000000000000ffff90921691909117905550505050505050505050565b600063ffffffff82111561420057506000919050565b5063ffffffff166000908152600c602052604090205460170b90565b600063ffffffff82111561423257506000919050565b5063ffffffff9081166000908152600c60205260409020547c010000000000000000000000000000000000000000000000000000000090041690565b60008060058054806020026020016040519081016040528092919081815260200182805480156142c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116142a9575b50508351600b54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150600690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116143035790505050505050905060005b83811015614396578181601f811061436357614363615192565b602002015161437290846151f0565b6143829063ffffffff168761522b565b95508061438e816152d5565b915050614349565b50600b546143c4907201000000000000000000000000000000000000900463ffffffff16633b9aca00615214565b6143ce9086615214565b945060005b8381101561444757600260008683815181106143f1576143f1615192565b6020908102919091018101516001600160a01b0316825281019190915260400160002054614433906201000090046bffffffffffffffffffffffff168761522b565b95508061443f816152d5565b9150506143d3565b505050505090565b600081831015614460575081611149565b50919050565b8060001061148c5760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c30565b6000808a8a8a8a8a8a8a8a8a6040516020016144da99989796959493929190615584565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150505b9998505050505050505050565b336001600160a01b0382160361459c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c30565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114610ca157601280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d4891291016119ab565b6146c06040518060800160405280600063ffffffff1681526020016060815260200160608152602001600060170b81525090565b60008060606000858060200190518101906146db919061561e565b929650909450925090506146ef868361490b565b81516040805160208082019690965281519082018252918252805160808101825263ffffffff969096168652938501529183015260170b606082015292915050565b60408051808201909152600e546001600160a01b0381168083527401000000000000000000000000000000000000000090910463ffffffff16602083015261477857505050565b60006147856001856151f0565b63ffffffff8181166000818152600c6020908152604091829020549087015187519251602481019490945260179190910b6044840181905289851660648501526084840189905294955061483593169160a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fbeed9b5100000000000000000000000000000000000000000000000000000000179052614983565b6119195760405162461bcd60e51b815260206004820152601060248201527f696e73756666696369656e7420676173000000000000000000000000000000006044820152606401610c30565b6000838381101561489457600285850304015b61489e818461444f565b95945050505050565b6000818610156148f95760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c4761736044820152606401610c30565b50633b9aca0094039190910101020290565b60008151602061491b9190615214565b6149269060a061522b565b61493190600061522b565b90508083511461305c5760405162461bcd60e51b815260206004820152601660248201527f7265706f7274206c656e677468206d69736d61746368000000000000000000006044820152606401610c30565b60005a61138881106149b757611388810390508460408204820311156149b7576000808451602086016000888af150600191505b509392505050565b508054600082559060005260206000209081019061148c9190614aed565b828054828255906000526020600020908101928215614a4a579160200282015b82811115614a4a57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039091161782556020909201916001909101906149fd565b50614a56929150614aed565b5090565b600483019183908215614a4a5791602002820160005b83821115614ab457835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302614a70565b8015614ae45782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614ab4565b5050614a569291505b5b80821115614a565760008155600101614aee565b6001600160a01b038116811461148c57600080fd5b600060208284031215614b2957600080fd5b8135612e7881614b02565b6000815180845260005b81811015614b5a57602081850181015186830182015201614b3e565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006111466020830184614b34565b60008060408385031215614ba057600080fd5b8235614bab81614b02565b91506020830135614bbb81614b02565b809150509250929050565b63ffffffff8116811461148c57600080fd5b600080600080600060a08688031215614bf057600080fd5b8535614bfb81614bc6565b94506020860135614c0b81614bc6565b93506040860135614c1b81614bc6565b92506060860135614c2b81614bc6565b9150608086013562ffffff81168114614c4357600080fd5b809150509295509295909350565b600081518084526020808501945080840160005b83811015614c8a5781516001600160a01b031687529582019590820190600101614c65565b509495945050505050565b6020815260006111466020830184614c51565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614d0057614d00614ca8565b604052919050565b600082601f830112614d1957600080fd5b813567ffffffffffffffff811115614d3357614d33614ca8565b614d466020601f19601f84011601614cd7565b818152846020838601011115614d5b57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215614d8b57600080fd5b8235614d9681614b02565b9150602083013567ffffffffffffffff811115614db257600080fd5b614dbe85828601614d08565b9150509250929050565b600060208284031215614dda57600080fd5b813569ffffffffffffffffffff81168114612e7857600080fd5b60008083601f840112614e0657600080fd5b50813567ffffffffffffffff811115614e1e57600080fd5b6020830191508360208260051b8501011115614e3957600080fd5b9250929050565b60008060008060408587031215614e5657600080fd5b843567ffffffffffffffff80821115614e6e57600080fd5b614e7a88838901614df4565b90965094506020870135915080821115614e9357600080fd5b50614ea087828801614df4565b95989497509550505050565b60008060008060008060008060e0898b031215614ec857600080fd5b606089018a811115614ed957600080fd5b8998503567ffffffffffffffff80821115614ef357600080fd5b818b0191508b601f830112614f0757600080fd5b813581811115614f1657600080fd5b8c6020828501011115614f2857600080fd5b6020830199508098505060808b0135915080821115614f4657600080fd5b614f528c838d01614df4565b909750955060a08b0135915080821115614f6b57600080fd5b50614f788b828c01614df4565b999c989b50969995989497949560c00135949350505050565b600060208284031215614fa357600080fd5b5035919050565b60008060408385031215614fbd57600080fd5b8235614fc881614b02565b946020939093013593505050565b600067ffffffffffffffff821115614ff057614ff0614ca8565b5060051b60200190565b600082601f83011261500b57600080fd5b8135602061502061501b83614fd6565b614cd7565b82815260059290921b8401810191818101908684111561503f57600080fd5b8286015b8481101561506357803561505681614b02565b8352918301918301615043565b509695505050505050565b803560ff8116811461212b57600080fd5b803567ffffffffffffffff8116811461212b57600080fd5b60008060008060008060c087890312156150b057600080fd5b863567ffffffffffffffff808211156150c857600080fd5b6150d48a838b01614ffa565b975060208901359150808211156150ea57600080fd5b6150f68a838b01614ffa565b965061510460408a0161506e565b9550606089013591508082111561511a57600080fd5b6151268a838b01614d08565b945061513460808a0161507f565b935060a089013591508082111561514a57600080fd5b5061515789828a01614d08565b9150509295509295509295565b6000806040838503121561517757600080fd5b823561518281614b02565b91506020830135614bbb81614bc6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff82811682821603908082111561520d5761520d6151c1565b5092915050565b8082028115828204841417611149576111496151c1565b80820180821115611149576111496151c1565b60006020828403121561525057600080fd5b5051919050565b60006020828403121561526957600080fd5b81518015158114612e7857600080fd5b6001600160a01b038416815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b63ffffffff81811683821601908082111561520d5761520d6151c1565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615306576153066151c1565b5060010190565b60ff8181168382160190811115611149576111496151c1565b8183823760009101908152919050565b8281526080810160608360208401379392505050565b81810381811115611149576111496151c1565b818103600083128015838313168383128216171561520d5761520d6151c1565b60ff818116838216029081169081811461520d5761520d6151c1565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526153cb8184018a614c51565b905082810360808401526153df8189614c51565b905060ff871660a084015282810360c08401526153fc8187614b34565b905067ffffffffffffffff851660e08401528281036101008401526154218185614b34565b9c9b505050505050505050505050565b600181811c9082168061544557607f821691505b602082108103614460577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000826154b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600063ffffffff8083168181036154d2576154d26151c1565b6001019392505050565b600061010080830160178c810b855260206001600160a01b038d168187015263ffffffff8c1660408701528360608701528293508a5180845261012087019450818c01935060005b81811015615542578451840b86529482019493820193600101615524565b505050505082810360808401526155598188614b34565b91505061556b60a083018660170b9052565b8360c083015261453760e083018464ffffffffff169052565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526155be8285018b614c51565b915083820360808501526155d2828a614c51565b915060ff881660a085015283820360c08501526155ef8288614b34565b90861660e085015283810361010085015290506154218185614b34565b8051601781900b811461212b57600080fd5b6000806000806080858703121561563457600080fd5b845161563f81614bc6565b809450506020808601519350604086015167ffffffffffffffff81111561566557600080fd5b8601601f8101881361567657600080fd5b805161568461501b82614fd6565b81815260059190911b8201830190838101908a8311156156a357600080fd5b928401925b828410156156c8576156b98461560c565b825292840192908401906156a8565b80965050505050506156dc6060860161560c565b90509295919450925056fe416363657373436f6e74726f6c6c65644f43523241676772656761746f7220312e302e30a2646970667358221220fed0cd68c5bfb721dbfe295c92e43bb48a83a2d5f53ba8e3528a83638fb3c6ad64736f6c63430008130033",
}

// OffchainAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OffchainAggregatorMetaData.ABI instead.
var OffchainAggregatorABI = OffchainAggregatorMetaData.ABI

// OffchainAggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OffchainAggregatorMetaData.Bin instead.
var OffchainAggregatorBin = OffchainAggregatorMetaData.Bin

// DeployOffchainAggregator deploys a new Ethereum contract, binding an instance of OffchainAggregator to it.
func DeployOffchainAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _link common.Address, _minAnswer *big.Int, _maxAnswer *big.Int, _billingAccessController common.Address, _requesterAccessController common.Address, _decimals uint8, description string) (common.Address, *types.Transaction, *OffchainAggregator, error) {
	parsed, err := OffchainAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffchainAggregatorBin), backend, _link, _minAnswer, _maxAnswer, _billingAccessController, _requesterAccessController, _decimals, description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffchainAggregator{OffchainAggregatorCaller: OffchainAggregatorCaller{contract: contract}, OffchainAggregatorTransactor: OffchainAggregatorTransactor{contract: contract}, OffchainAggregatorFilterer: OffchainAggregatorFilterer{contract: contract}}, nil
}

// OffchainAggregator is an auto generated Go binding around an Ethereum contract.
type OffchainAggregator struct {
	OffchainAggregatorCaller     // Read-only binding to the contract
	OffchainAggregatorTransactor // Write-only binding to the contract
	OffchainAggregatorFilterer   // Log filterer for contract events
}

// OffchainAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OffchainAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OffchainAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OffchainAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OffchainAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OffchainAggregatorSession struct {
	Contract     *OffchainAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OffchainAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OffchainAggregatorCallerSession struct {
	Contract *OffchainAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OffchainAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OffchainAggregatorTransactorSession struct {
	Contract     *OffchainAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OffchainAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OffchainAggregatorRaw struct {
	Contract *OffchainAggregator // Generic contract binding to access the raw methods on
}

// OffchainAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OffchainAggregatorCallerRaw struct {
	Contract *OffchainAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// OffchainAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OffchainAggregatorTransactorRaw struct {
	Contract *OffchainAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOffchainAggregator creates a new instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregator(address common.Address, backend bind.ContractBackend) (*OffchainAggregator, error) {
	contract, err := bindOffchainAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregator{OffchainAggregatorCaller: OffchainAggregatorCaller{contract: contract}, OffchainAggregatorTransactor: OffchainAggregatorTransactor{contract: contract}, OffchainAggregatorFilterer: OffchainAggregatorFilterer{contract: contract}}, nil
}

// NewOffchainAggregatorCaller creates a new read-only instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorCaller(address common.Address, caller bind.ContractCaller) (*OffchainAggregatorCaller, error) {
	contract, err := bindOffchainAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorCaller{contract: contract}, nil
}

// NewOffchainAggregatorTransactor creates a new write-only instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OffchainAggregatorTransactor, error) {
	contract, err := bindOffchainAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorTransactor{contract: contract}, nil
}

// NewOffchainAggregatorFilterer creates a new log filterer instance of OffchainAggregator, bound to a specific deployed contract.
func NewOffchainAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OffchainAggregatorFilterer, error) {
	contract, err := bindOffchainAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorFilterer{contract: contract}, nil
}

// bindOffchainAggregator binds a generic wrapper to an already deployed contract.
func bindOffchainAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OffchainAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainAggregator *OffchainAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainAggregator.Contract.OffchainAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainAggregator *OffchainAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.OffchainAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainAggregator *OffchainAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.OffchainAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OffchainAggregator *OffchainAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffchainAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OffchainAggregator *OffchainAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OffchainAggregator *OffchainAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.contract.Transact(opts, method, params...)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorSession) CheckEnabled() (bool, error) {
	return _OffchainAggregator.Contract.CheckEnabled(&_OffchainAggregator.CallOpts)
}

// CheckEnabled is a free data retrieval call binding the contract method 0xdc7f0124.
//
// Solidity: function checkEnabled() view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCallerSession) CheckEnabled() (bool, error) {
	return _OffchainAggregator.Contract.CheckEnabled(&_OffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorSession) Decimals() (uint8, error) {
	return _OffchainAggregator.Contract.Decimals(&_OffchainAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Decimals() (uint8, error) {
	return _OffchainAggregator.Contract.Decimals(&_OffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorSession) Description() (string, error) {
	return _OffchainAggregator.Contract.Description(&_OffchainAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Description() (string, error) {
	return _OffchainAggregator.Contract.Description(&_OffchainAggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetAnswer(&_OffchainAggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetAnswer(&_OffchainAggregator.CallOpts, _roundId)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPriceGwei       uint32
		ReasonableGasPriceGwei    uint32
		ObservationPaymentGjuels  uint32
		TransmissionPaymentGjuels uint32
		AccountingGas             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPriceGwei = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceGwei = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ObservationPaymentGjuels = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TransmissionPaymentGjuels = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetBilling(&_OffchainAggregator.CallOpts)
}

// GetBilling is a free data retrieval call binding the contract method 0x29937268.
//
// Solidity: function getBilling() view returns(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetBilling(&_OffchainAggregator.CallOpts)
}

// GetBillingAccessController is a free data retrieval call binding the contract method 0xc4c92b37.
//
// Solidity: function getBillingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBillingAccessController is a free data retrieval call binding the contract method 0xc4c92b37.
//
// Solidity: function getBillingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) GetBillingAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.GetBillingAccessController(&_OffchainAggregator.CallOpts)
}

// GetBillingAccessController is a free data retrieval call binding the contract method 0xc4c92b37.
//
// Solidity: function getBillingAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.GetBillingAccessController(&_OffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorCaller) GetLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorSession) GetLinkToken() (common.Address, error) {
	return _OffchainAggregator.Contract.GetLinkToken(&_OffchainAggregator.CallOpts)
}

// GetLinkToken is a free data retrieval call binding the contract method 0xe76d5168.
//
// Solidity: function getLinkToken() view returns(address linkToken)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetLinkToken() (common.Address, error) {
	return _OffchainAggregator.Contract.GetLinkToken(&_OffchainAggregator.CallOpts)
}

// GetRequesterAccessController is a free data retrieval call binding the contract method 0xdaffc4b5.
//
// Solidity: function getRequesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) GetRequesterAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getRequesterAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRequesterAccessController is a free data retrieval call binding the contract method 0xdaffc4b5.
//
// Solidity: function getRequesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) GetRequesterAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.GetRequesterAccessController(&_OffchainAggregator.CallOpts)
}

// GetRequesterAccessController is a free data retrieval call binding the contract method 0xdaffc4b5.
//
// Solidity: function getRequesterAccessController() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetRequesterAccessController() (common.Address, error) {
	return _OffchainAggregator.Contract.GetRequesterAccessController(&_OffchainAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetRoundData(&_OffchainAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.GetRoundData(&_OffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetTimestamp(&_OffchainAggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _OffchainAggregator.Contract.GetTimestamp(&_OffchainAggregator.CallOpts, _roundId)
}

// GetTransmitters is a free data retrieval call binding the contract method 0x666cab8d.
//
// Solidity: function getTransmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTransmitters is a free data retrieval call binding the contract method 0x666cab8d.
//
// Solidity: function getTransmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorSession) GetTransmitters() ([]common.Address, error) {
	return _OffchainAggregator.Contract.GetTransmitters(&_OffchainAggregator.CallOpts)
}

// GetTransmitters is a free data retrieval call binding the contract method 0x666cab8d.
//
// Solidity: function getTransmitters() view returns(address[])
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetTransmitters() ([]common.Address, error) {
	return _OffchainAggregator.Contract.GetTransmitters(&_OffchainAggregator.CallOpts)
}

// GetValidatorConfig is a free data retrieval call binding the contract method 0x9bd2c0b1.
//
// Solidity: function getValidatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorCaller) GetValidatorConfig(opts *bind.CallOpts) (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "getValidatorConfig")

	outstruct := new(struct {
		Validator common.Address
		GasLimit  uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetValidatorConfig is a free data retrieval call binding the contract method 0x9bd2c0b1.
//
// Solidity: function getValidatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorSession) GetValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _OffchainAggregator.Contract.GetValidatorConfig(&_OffchainAggregator.CallOpts)
}

// GetValidatorConfig is a free data retrieval call binding the contract method 0x9bd2c0b1.
//
// Solidity: function getValidatorConfig() view returns(address validator, uint32 gasLimit)
func (_OffchainAggregator *OffchainAggregatorCallerSession) GetValidatorConfig() (struct {
	Validator common.Address
	GasLimit  uint32
}, error) {
	return _OffchainAggregator.Contract.GetValidatorConfig(&_OffchainAggregator.CallOpts)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_OffchainAggregator *OffchainAggregatorSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _OffchainAggregator.Contract.HasAccess(&_OffchainAggregator.CallOpts, _user, _calldata)
}

// HasAccess is a free data retrieval call binding the contract method 0x6b14daf8.
//
// Solidity: function hasAccess(address _user, bytes _calldata) view returns(bool)
func (_OffchainAggregator *OffchainAggregatorCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _OffchainAggregator.Contract.HasAccess(&_OffchainAggregator.CallOpts, _user, _calldata)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestAnswer(&_OffchainAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestAnswer(&_OffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_OffchainAggregator *OffchainAggregatorSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDetails(&_OffchainAggregator.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDetails(&_OffchainAggregator.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDigestAndEpoch(&_OffchainAggregator.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _OffchainAggregator.Contract.LatestConfigDigestAndEpoch(&_OffchainAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestRound() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestRound(&_OffchainAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestRound(&_OffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestRoundData(&_OffchainAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _OffchainAggregator.Contract.LatestRoundData(&_OffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestTimestamp(&_OffchainAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _OffchainAggregator.Contract.LatestTimestamp(&_OffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes32 configDigest, uint32 epoch, uint8 round, int192 latestAnswer_, uint64 latestTimestamp_)
func (_OffchainAggregator *OffchainAggregatorCaller) LatestTransmissionDetails(opts *bind.CallOpts) (struct {
	ConfigDigest    [32]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "latestTransmissionDetails")

	outstruct := new(struct {
		ConfigDigest    [32]byte
		Epoch           uint32
		Round           uint8
		LatestAnswer    *big.Int
		LatestTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigDigest = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Round = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.LatestAnswer = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LatestTimestamp = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes32 configDigest, uint32 epoch, uint8 round, int192 latestAnswer_, uint64 latestTimestamp_)
func (_OffchainAggregator *OffchainAggregatorSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [32]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _OffchainAggregator.Contract.LatestTransmissionDetails(&_OffchainAggregator.CallOpts)
}

// LatestTransmissionDetails is a free data retrieval call binding the contract method 0xe5fe4577.
//
// Solidity: function latestTransmissionDetails() view returns(bytes32 configDigest, uint32 epoch, uint8 round, int192 latestAnswer_, uint64 latestTimestamp_)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LatestTransmissionDetails() (struct {
	ConfigDigest    [32]byte
	Epoch           uint32
	Round           uint8
	LatestAnswer    *big.Int
	LatestTimestamp uint64
}, error) {
	return _OffchainAggregator.Contract.LatestTransmissionDetails(&_OffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OffchainAggregator.Contract.LinkAvailableForPayment(&_OffchainAggregator.CallOpts)
}

// LinkAvailableForPayment is a free data retrieval call binding the contract method 0xd09dc339.
//
// Solidity: function linkAvailableForPayment() view returns(int256 availableBalance)
func (_OffchainAggregator *OffchainAggregatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _OffchainAggregator.Contract.LinkAvailableForPayment(&_OffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorSession) MaxAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MaxAnswer(&_OffchainAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MaxAnswer(&_OffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorSession) MinAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MinAnswer(&_OffchainAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_OffchainAggregator *OffchainAggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _OffchainAggregator.Contract.MinAnswer(&_OffchainAggregator.CallOpts)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address transmitterAddress) view returns(uint32)
func (_OffchainAggregator *OffchainAggregatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address transmitterAddress) view returns(uint32)
func (_OffchainAggregator *OffchainAggregatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _OffchainAggregator.Contract.OracleObservationCount(&_OffchainAggregator.CallOpts, transmitterAddress)
}

// OracleObservationCount is a free data retrieval call binding the contract method 0xe4902f82.
//
// Solidity: function oracleObservationCount(address transmitterAddress) view returns(uint32)
func (_OffchainAggregator *OffchainAggregatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _OffchainAggregator.Contract.OracleObservationCount(&_OffchainAggregator.CallOpts, transmitterAddress)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address transmitterAddress) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address transmitterAddress) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _OffchainAggregator.Contract.OwedPayment(&_OffchainAggregator.CallOpts, transmitterAddress)
}

// OwedPayment is a free data retrieval call binding the contract method 0x0eafb25b.
//
// Solidity: function owedPayment(address transmitterAddress) view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _OffchainAggregator.Contract.OwedPayment(&_OffchainAggregator.CallOpts, transmitterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorSession) Owner() (common.Address, error) {
	return _OffchainAggregator.Contract.Owner(&_OffchainAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Owner() (common.Address, error) {
	return _OffchainAggregator.Contract.Owner(&_OffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorSession) TypeAndVersion() (string, error) {
	return _OffchainAggregator.Contract.TypeAndVersion(&_OffchainAggregator.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_OffchainAggregator *OffchainAggregatorCallerSession) TypeAndVersion() (string, error) {
	return _OffchainAggregator.Contract.TypeAndVersion(&_OffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OffchainAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorSession) Version() (*big.Int, error) {
	return _OffchainAggregator.Contract.Version(&_OffchainAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_OffchainAggregator *OffchainAggregatorCallerSession) Version() (*big.Int, error) {
	return _OffchainAggregator.Contract.Version(&_OffchainAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptOwnership(&_OffchainAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptOwnership(&_OffchainAggregator.TransactOpts)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptPayeeship(&_OffchainAggregator.TransactOpts, transmitter)
}

// AcceptPayeeship is a paid mutator transaction binding the contract method 0xb121e147.
//
// Solidity: function acceptPayeeship(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AcceptPayeeship(&_OffchainAggregator.TransactOpts, transmitter)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "addAccess", _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AddAccess(&_OffchainAggregator.TransactOpts, _user)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.AddAccess(&_OffchainAggregator.TransactOpts, _user)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "disableAccessCheck")
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.DisableAccessCheck(&_OffchainAggregator.TransactOpts)
}

// DisableAccessCheck is a paid mutator transaction binding the contract method 0x0a756983.
//
// Solidity: function disableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.DisableAccessCheck(&_OffchainAggregator.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "enableAccessCheck")
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.EnableAccessCheck(&_OffchainAggregator.TransactOpts)
}

// EnableAccessCheck is a paid mutator transaction binding the contract method 0x8038e4a1.
//
// Solidity: function enableAccessCheck() returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.EnableAccessCheck(&_OffchainAggregator.TransactOpts)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "removeAccess", _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RemoveAccess(&_OffchainAggregator.TransactOpts, _user)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address _user) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RemoveAccess(&_OffchainAggregator.TransactOpts, _user)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RequestNewRound(&_OffchainAggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_OffchainAggregator *OffchainAggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _OffchainAggregator.Contract.RequestNewRound(&_OffchainAggregator.TransactOpts)
}

// SetBilling is a paid mutator transaction binding the contract method 0x643dc105.
//
// Solidity: function setBilling(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

// SetBilling is a paid mutator transaction binding the contract method 0x643dc105.
//
// Solidity: function setBilling(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBilling(&_OffchainAggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

// SetBilling is a paid mutator transaction binding the contract method 0x643dc105.
//
// Solidity: function setBilling(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBilling(&_OffchainAggregator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBillingAccessController(&_OffchainAggregator.TransactOpts, _billingAccessController)
}

// SetBillingAccessController is a paid mutator transaction binding the contract method 0xfbffd2c1.
//
// Solidity: function setBillingAccessController(address _billingAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetBillingAccessController(&_OffchainAggregator.TransactOpts, _billingAccessController)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetConfig(&_OffchainAggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetConfig(&_OffchainAggregator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address linkToken, address recipient) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetLinkToken(opts *bind.TransactOpts, linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setLinkToken", linkToken, recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address linkToken, address recipient) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLinkToken(&_OffchainAggregator.TransactOpts, linkToken, recipient)
}

// SetLinkToken is a paid mutator transaction binding the contract method 0x4fb17470.
//
// Solidity: function setLinkToken(address linkToken, address recipient) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetLinkToken(linkToken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetLinkToken(&_OffchainAggregator.TransactOpts, linkToken, recipient)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] transmitters, address[] payees) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setPayees", transmitters, payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] transmitters, address[] payees) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetPayees(&_OffchainAggregator.TransactOpts, transmitters, payees)
}

// SetPayees is a paid mutator transaction binding the contract method 0x9c849b30.
//
// Solidity: function setPayees(address[] transmitters, address[] payees) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetPayees(&_OffchainAggregator.TransactOpts, transmitters, payees)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetRequesterAccessController(opts *bind.TransactOpts, requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setRequesterAccessController", requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetRequesterAccessController(&_OffchainAggregator.TransactOpts, requesterAccessController)
}

// SetRequesterAccessController is a paid mutator transaction binding the contract method 0x9e3ceeab.
//
// Solidity: function setRequesterAccessController(address requesterAccessController) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetRequesterAccessController(requesterAccessController common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetRequesterAccessController(&_OffchainAggregator.TransactOpts, requesterAccessController)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address newValidator, uint32 newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) SetValidatorConfig(opts *bind.TransactOpts, newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "setValidatorConfig", newValidator, newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address newValidator, uint32 newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetValidatorConfig(&_OffchainAggregator.TransactOpts, newValidator, newGasLimit)
}

// SetValidatorConfig is a paid mutator transaction binding the contract method 0xeb457163.
//
// Solidity: function setValidatorConfig(address newValidator, uint32 newGasLimit) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) SetValidatorConfig(newValidator common.Address, newGasLimit uint32) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.SetValidatorConfig(&_OffchainAggregator.TransactOpts, newValidator, newGasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OffchainAggregator *OffchainAggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferOwnership(&_OffchainAggregator.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferOwnership(&_OffchainAggregator.TransactOpts, to)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_OffchainAggregator *OffchainAggregatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferPayeeship(&_OffchainAggregator.TransactOpts, transmitter, proposed)
}

// TransferPayeeship is a paid mutator transaction binding the contract method 0xeb5dcd6c.
//
// Solidity: function transferPayeeship(address transmitter, address proposed) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.TransferPayeeship(&_OffchainAggregator.TransactOpts, transmitter, proposed)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.Transmit(&_OffchainAggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.Transmit(&_OffchainAggregator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient, uint256 amount) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient, uint256 amount) returns()
func (_OffchainAggregator *OffchainAggregatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawFunds(&_OffchainAggregator.TransactOpts, recipient, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient, uint256 amount) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawFunds(&_OffchainAggregator.TransactOpts, recipient, amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.contract.Transact(opts, "withdrawPayment", transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawPayment(&_OffchainAggregator.TransactOpts, transmitter)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x8ac28d5a.
//
// Solidity: function withdrawPayment(address transmitter) returns()
func (_OffchainAggregator *OffchainAggregatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _OffchainAggregator.Contract.WithdrawPayment(&_OffchainAggregator.TransactOpts, transmitter)
}

// OffchainAggregatorAddedAccessIterator is returned from FilterAddedAccess and is used to iterate over the raw logs and unpacked data for AddedAccess events raised by the OffchainAggregator contract.
type OffchainAggregatorAddedAccessIterator struct {
	Event *OffchainAggregatorAddedAccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorAddedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorAddedAccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorAddedAccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorAddedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorAddedAccess represents a AddedAccess event raised by the OffchainAggregator contract.
type OffchainAggregatorAddedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedAccess is a free log retrieval operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*OffchainAggregatorAddedAccessIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorAddedAccessIterator{contract: _OffchainAggregator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

// WatchAddedAccess is a free log subscription operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorAddedAccess)
				if err := _OffchainAggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedAccess is a log parse operation binding the contract event 0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4.
//
// Solidity: event AddedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseAddedAccess(log types.Log) (*OffchainAggregatorAddedAccess, error) {
	event := new(OffchainAggregatorAddedAccess)
	if err := _OffchainAggregator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the OffchainAggregator contract.
type OffchainAggregatorAnswerUpdatedIterator struct {
	Event *OffchainAggregatorAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorAnswerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorAnswerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorAnswerUpdated represents a AnswerUpdated event raised by the OffchainAggregator contract.
type OffchainAggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*OffchainAggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorAnswerUpdatedIterator{contract: _OffchainAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorAnswerUpdated)
				if err := _OffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*OffchainAggregatorAnswerUpdated, error) {
	event := new(OffchainAggregatorAnswerUpdated)
	if err := _OffchainAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorBillingAccessControllerSetIterator is returned from FilterBillingAccessControllerSet and is used to iterate over the raw logs and unpacked data for BillingAccessControllerSet events raised by the OffchainAggregator contract.
type OffchainAggregatorBillingAccessControllerSetIterator struct {
	Event *OffchainAggregatorBillingAccessControllerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorBillingAccessControllerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorBillingAccessControllerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorBillingAccessControllerSet represents a BillingAccessControllerSet event raised by the OffchainAggregator contract.
type OffchainAggregatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBillingAccessControllerSet is a free log retrieval operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*OffchainAggregatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorBillingAccessControllerSetIterator{contract: _OffchainAggregator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchBillingAccessControllerSet is a free log subscription operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorBillingAccessControllerSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBillingAccessControllerSet is a log parse operation binding the contract event 0x793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912.
//
// Solidity: event BillingAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*OffchainAggregatorBillingAccessControllerSet, error) {
	event := new(OffchainAggregatorBillingAccessControllerSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorBillingSetIterator is returned from FilterBillingSet and is used to iterate over the raw logs and unpacked data for BillingSet events raised by the OffchainAggregator contract.
type OffchainAggregatorBillingSetIterator struct {
	Event *OffchainAggregatorBillingSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorBillingSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorBillingSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorBillingSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorBillingSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorBillingSet represents a BillingSet event raised by the OffchainAggregator contract.
type OffchainAggregatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBillingSet is a free log retrieval operation binding the contract event 0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f.
//
// Solidity: event BillingSet(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*OffchainAggregatorBillingSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorBillingSetIterator{contract: _OffchainAggregator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

// WatchBillingSet is a free log subscription operation binding the contract event 0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f.
//
// Solidity: event BillingSet(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorBillingSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBillingSet is a log parse operation binding the contract event 0x0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f.
//
// Solidity: event BillingSet(uint32 maximumGasPriceGwei, uint32 reasonableGasPriceGwei, uint32 observationPaymentGjuels, uint32 transmissionPaymentGjuels, uint24 accountingGas)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseBillingSet(log types.Log) (*OffchainAggregatorBillingSet, error) {
	event := new(OffchainAggregatorBillingSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorCheckAccessDisabledIterator is returned from FilterCheckAccessDisabled and is used to iterate over the raw logs and unpacked data for CheckAccessDisabled events raised by the OffchainAggregator contract.
type OffchainAggregatorCheckAccessDisabledIterator struct {
	Event *OffchainAggregatorCheckAccessDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorCheckAccessDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorCheckAccessDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorCheckAccessDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorCheckAccessDisabled represents a CheckAccessDisabled event raised by the OffchainAggregator contract.
type OffchainAggregatorCheckAccessDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessDisabled is a free log retrieval operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*OffchainAggregatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorCheckAccessDisabledIterator{contract: _OffchainAggregator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessDisabled is a free log subscription operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorCheckAccessDisabled)
				if err := _OffchainAggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckAccessDisabled is a log parse operation binding the contract event 0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638.
//
// Solidity: event CheckAccessDisabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseCheckAccessDisabled(log types.Log) (*OffchainAggregatorCheckAccessDisabled, error) {
	event := new(OffchainAggregatorCheckAccessDisabled)
	if err := _OffchainAggregator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorCheckAccessEnabledIterator is returned from FilterCheckAccessEnabled and is used to iterate over the raw logs and unpacked data for CheckAccessEnabled events raised by the OffchainAggregator contract.
type OffchainAggregatorCheckAccessEnabledIterator struct {
	Event *OffchainAggregatorCheckAccessEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorCheckAccessEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorCheckAccessEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorCheckAccessEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorCheckAccessEnabled represents a CheckAccessEnabled event raised by the OffchainAggregator contract.
type OffchainAggregatorCheckAccessEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCheckAccessEnabled is a free log retrieval operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*OffchainAggregatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorCheckAccessEnabledIterator{contract: _OffchainAggregator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

// WatchCheckAccessEnabled is a free log subscription operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorCheckAccessEnabled)
				if err := _OffchainAggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckAccessEnabled is a log parse operation binding the contract event 0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480.
//
// Solidity: event CheckAccessEnabled()
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseCheckAccessEnabled(log types.Log) (*OffchainAggregatorCheckAccessEnabled, error) {
	event := new(OffchainAggregatorCheckAccessEnabled)
	if err := _OffchainAggregator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the OffchainAggregator contract.
type OffchainAggregatorConfigSetIterator struct {
	Event *OffchainAggregatorConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorConfigSet represents a ConfigSet event raised by the OffchainAggregator contract.
type OffchainAggregatorConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffchainAggregatorConfigSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorConfigSetIterator{contract: _OffchainAggregator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorConfigSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigSet is a log parse operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseConfigSet(log types.Log) (*OffchainAggregatorConfigSet, error) {
	event := new(OffchainAggregatorConfigSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorLinkTokenSetIterator is returned from FilterLinkTokenSet and is used to iterate over the raw logs and unpacked data for LinkTokenSet events raised by the OffchainAggregator contract.
type OffchainAggregatorLinkTokenSetIterator struct {
	Event *OffchainAggregatorLinkTokenSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorLinkTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorLinkTokenSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorLinkTokenSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorLinkTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorLinkTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorLinkTokenSet represents a LinkTokenSet event raised by the OffchainAggregator contract.
type OffchainAggregatorLinkTokenSet struct {
	OldLinkToken common.Address
	NewLinkToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLinkTokenSet is a free log retrieval operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed oldLinkToken, address indexed newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterLinkTokenSet(opts *bind.FilterOpts, oldLinkToken []common.Address, newLinkToken []common.Address) (*OffchainAggregatorLinkTokenSetIterator, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorLinkTokenSetIterator{contract: _OffchainAggregator.contract, event: "LinkTokenSet", logs: logs, sub: sub}, nil
}

// WatchLinkTokenSet is a free log subscription operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed oldLinkToken, address indexed newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchLinkTokenSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorLinkTokenSet, oldLinkToken []common.Address, newLinkToken []common.Address) (event.Subscription, error) {

	var oldLinkTokenRule []interface{}
	for _, oldLinkTokenItem := range oldLinkToken {
		oldLinkTokenRule = append(oldLinkTokenRule, oldLinkTokenItem)
	}
	var newLinkTokenRule []interface{}
	for _, newLinkTokenItem := range newLinkToken {
		newLinkTokenRule = append(newLinkTokenRule, newLinkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "LinkTokenSet", oldLinkTokenRule, newLinkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorLinkTokenSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLinkTokenSet is a log parse operation binding the contract event 0x4966a50c93f855342ccf6c5c0d358b85b91335b2acedc7da0932f691f351711a.
//
// Solidity: event LinkTokenSet(address indexed oldLinkToken, address indexed newLinkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseLinkTokenSet(log types.Log) (*OffchainAggregatorLinkTokenSet, error) {
	event := new(OffchainAggregatorLinkTokenSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "LinkTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the OffchainAggregator contract.
type OffchainAggregatorNewRoundIterator struct {
	Event *OffchainAggregatorNewRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorNewRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorNewRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorNewRound represents a NewRound event raised by the OffchainAggregator contract.
type OffchainAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OffchainAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorNewRoundIterator{contract: _OffchainAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorNewRound)
				if err := _OffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseNewRound(log types.Log) (*OffchainAggregatorNewRound, error) {
	event := new(OffchainAggregatorNewRound)
	if err := _OffchainAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorNewTransmissionIterator is returned from FilterNewTransmission and is used to iterate over the raw logs and unpacked data for NewTransmission events raised by the OffchainAggregator contract.
type OffchainAggregatorNewTransmissionIterator struct {
	Event *OffchainAggregatorNewTransmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorNewTransmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorNewTransmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorNewTransmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorNewTransmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorNewTransmission represents a NewTransmission event raised by the OffchainAggregator contract.
type OffchainAggregatorNewTransmission struct {
	AggregatorRoundId     uint32
	Answer                *big.Int
	Transmitter           common.Address
	ObservationsTimestamp uint32
	Observations          []*big.Int
	Observers             []byte
	JuelsPerFeeCoin       *big.Int
	ConfigDigest          [32]byte
	EpochAndRound         *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewTransmission is a free log retrieval operation binding the contract event 0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, int192 answer, address transmitter, uint32 observationsTimestamp, int192[] observations, bytes observers, int192 juelsPerFeeCoin, bytes32 configDigest, uint40 epochAndRound)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*OffchainAggregatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorNewTransmissionIterator{contract: _OffchainAggregator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

// WatchNewTransmission is a free log subscription operation binding the contract event 0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, int192 answer, address transmitter, uint32 observationsTimestamp, int192[] observations, bytes observers, int192 juelsPerFeeCoin, bytes32 configDigest, uint40 epochAndRound)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorNewTransmission)
				if err := _OffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTransmission is a log parse operation binding the contract event 0xc797025feeeaf2cd924c99e9205acb8ec04d5cad21c41ce637a38fb6dee6016a.
//
// Solidity: event NewTransmission(uint32 indexed aggregatorRoundId, int192 answer, address transmitter, uint32 observationsTimestamp, int192[] observations, bytes observers, int192 juelsPerFeeCoin, bytes32 configDigest, uint40 epochAndRound)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseNewTransmission(log types.Log) (*OffchainAggregatorNewTransmission, error) {
	event := new(OffchainAggregatorNewTransmission)
	if err := _OffchainAggregator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOraclePaidIterator is returned from FilterOraclePaid and is used to iterate over the raw logs and unpacked data for OraclePaid events raised by the OffchainAggregator contract.
type OffchainAggregatorOraclePaidIterator struct {
	Event *OffchainAggregatorOraclePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorOraclePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOraclePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorOraclePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorOraclePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOraclePaid represents a OraclePaid event raised by the OffchainAggregator contract.
type OffchainAggregatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePaid is a free log retrieval operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*OffchainAggregatorOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOraclePaidIterator{contract: _OffchainAggregator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

// WatchOraclePaid is a free log subscription operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOraclePaid)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOraclePaid is a log parse operation binding the contract event 0xd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c.
//
// Solidity: event OraclePaid(address indexed transmitter, address indexed payee, uint256 amount, address indexed linkToken)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOraclePaid(log types.Log) (*OffchainAggregatorOraclePaid, error) {
	event := new(OffchainAggregatorOraclePaid)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferRequestedIterator struct {
	Event *OffchainAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffchainAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOwnershipTransferRequestedIterator{contract: _OffchainAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOwnershipTransferRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffchainAggregatorOwnershipTransferRequested, error) {
	event := new(OffchainAggregatorOwnershipTransferRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferredIterator struct {
	Event *OffchainAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the OffchainAggregator contract.
type OffchainAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffchainAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorOwnershipTransferredIterator{contract: _OffchainAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorOwnershipTransferred)
				if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*OffchainAggregatorOwnershipTransferred, error) {
	event := new(OffchainAggregatorOwnershipTransferred)
	if err := _OffchainAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorPayeeshipTransferRequestedIterator is returned from FilterPayeeshipTransferRequested and is used to iterate over the raw logs and unpacked data for PayeeshipTransferRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferRequestedIterator struct {
	Event *OffchainAggregatorPayeeshipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorPayeeshipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorPayeeshipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorPayeeshipTransferRequested represents a PayeeshipTransferRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferRequested is a free log retrieval operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*OffchainAggregatorPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorPayeeshipTransferRequestedIterator{contract: _OffchainAggregator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferRequested is a free log subscription operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorPayeeshipTransferRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayeeshipTransferRequested is a log parse operation binding the contract event 0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367.
//
// Solidity: event PayeeshipTransferRequested(address indexed transmitter, address indexed current, address indexed proposed)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*OffchainAggregatorPayeeshipTransferRequested, error) {
	event := new(OffchainAggregatorPayeeshipTransferRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorPayeeshipTransferredIterator is returned from FilterPayeeshipTransferred and is used to iterate over the raw logs and unpacked data for PayeeshipTransferred events raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferredIterator struct {
	Event *OffchainAggregatorPayeeshipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorPayeeshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorPayeeshipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorPayeeshipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorPayeeshipTransferred represents a PayeeshipTransferred event raised by the OffchainAggregator contract.
type OffchainAggregatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayeeshipTransferred is a free log retrieval operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*OffchainAggregatorPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorPayeeshipTransferredIterator{contract: _OffchainAggregator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

// WatchPayeeshipTransferred is a free log subscription operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorPayeeshipTransferred)
				if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayeeshipTransferred is a log parse operation binding the contract event 0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3.
//
// Solidity: event PayeeshipTransferred(address indexed transmitter, address indexed previous, address indexed current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParsePayeeshipTransferred(log types.Log) (*OffchainAggregatorPayeeshipTransferred, error) {
	event := new(OffchainAggregatorPayeeshipTransferred)
	if err := _OffchainAggregator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorRemovedAccessIterator is returned from FilterRemovedAccess and is used to iterate over the raw logs and unpacked data for RemovedAccess events raised by the OffchainAggregator contract.
type OffchainAggregatorRemovedAccessIterator struct {
	Event *OffchainAggregatorRemovedAccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorRemovedAccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorRemovedAccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorRemovedAccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorRemovedAccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorRemovedAccess represents a RemovedAccess event raised by the OffchainAggregator contract.
type OffchainAggregatorRemovedAccess struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedAccess is a free log retrieval operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*OffchainAggregatorRemovedAccessIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorRemovedAccessIterator{contract: _OffchainAggregator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

// WatchRemovedAccess is a free log subscription operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorRemovedAccess)
				if err := _OffchainAggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedAccess is a log parse operation binding the contract event 0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1.
//
// Solidity: event RemovedAccess(address user)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseRemovedAccess(log types.Log) (*OffchainAggregatorRemovedAccess, error) {
	event := new(OffchainAggregatorRemovedAccess)
	if err := _OffchainAggregator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorRequesterAccessControllerSetIterator is returned from FilterRequesterAccessControllerSet and is used to iterate over the raw logs and unpacked data for RequesterAccessControllerSet events raised by the OffchainAggregator contract.
type OffchainAggregatorRequesterAccessControllerSetIterator struct {
	Event *OffchainAggregatorRequesterAccessControllerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorRequesterAccessControllerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorRequesterAccessControllerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorRequesterAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorRequesterAccessControllerSet represents a RequesterAccessControllerSet event raised by the OffchainAggregator contract.
type OffchainAggregatorRequesterAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequesterAccessControllerSet is a free log retrieval operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterRequesterAccessControllerSet(opts *bind.FilterOpts) (*OffchainAggregatorRequesterAccessControllerSetIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorRequesterAccessControllerSetIterator{contract: _OffchainAggregator.contract, event: "RequesterAccessControllerSet", logs: logs, sub: sub}, nil
}

// WatchRequesterAccessControllerSet is a free log subscription operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchRequesterAccessControllerSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorRequesterAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "RequesterAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorRequesterAccessControllerSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequesterAccessControllerSet is a log parse operation binding the contract event 0x27b89aede8b560578baaa25ee5ce3852c5eecad1e114b941bbd89e1eb4bae634.
//
// Solidity: event RequesterAccessControllerSet(address old, address current)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseRequesterAccessControllerSet(log types.Log) (*OffchainAggregatorRequesterAccessControllerSet, error) {
	event := new(OffchainAggregatorRequesterAccessControllerSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "RequesterAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorRoundRequestedIterator is returned from FilterRoundRequested and is used to iterate over the raw logs and unpacked data for RoundRequested events raised by the OffchainAggregator contract.
type OffchainAggregatorRoundRequestedIterator struct {
	Event *OffchainAggregatorRoundRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorRoundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorRoundRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorRoundRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorRoundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorRoundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorRoundRequested represents a RoundRequested event raised by the OffchainAggregator contract.
type OffchainAggregatorRoundRequested struct {
	Requester    common.Address
	ConfigDigest [32]byte
	Epoch        uint32
	Round        uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRoundRequested is a free log retrieval operation binding the contract event 0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c.
//
// Solidity: event RoundRequested(address indexed requester, bytes32 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterRoundRequested(opts *bind.FilterOpts, requester []common.Address) (*OffchainAggregatorRoundRequestedIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorRoundRequestedIterator{contract: _OffchainAggregator.contract, event: "RoundRequested", logs: logs, sub: sub}, nil
}

// WatchRoundRequested is a free log subscription operation binding the contract event 0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c.
//
// Solidity: event RoundRequested(address indexed requester, bytes32 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchRoundRequested(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorRoundRequested, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "RoundRequested", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorRoundRequested)
				if err := _OffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoundRequested is a log parse operation binding the contract event 0x41e3990591fd372502daa15842da15bc7f41c75309ab3ff4f56f1848c178825c.
//
// Solidity: event RoundRequested(address indexed requester, bytes32 configDigest, uint32 epoch, uint8 round)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseRoundRequested(log types.Log) (*OffchainAggregatorRoundRequested, error) {
	event := new(OffchainAggregatorRoundRequested)
	if err := _OffchainAggregator.contract.UnpackLog(event, "RoundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorTransmittedIterator is returned from FilterTransmitted and is used to iterate over the raw logs and unpacked data for Transmitted events raised by the OffchainAggregator contract.
type OffchainAggregatorTransmittedIterator struct {
	Event *OffchainAggregatorTransmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorTransmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorTransmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorTransmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorTransmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorTransmitted represents a Transmitted event raised by the OffchainAggregator contract.
type OffchainAggregatorTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransmitted is a free log retrieval operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OffchainAggregatorTransmittedIterator, error) {

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorTransmittedIterator{contract: _OffchainAggregator.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

// WatchTransmitted is a free log subscription operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorTransmitted) (event.Subscription, error) {

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorTransmitted)
				if err := _OffchainAggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransmitted is a log parse operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseTransmitted(log types.Log) (*OffchainAggregatorTransmitted, error) {
	event := new(OffchainAggregatorTransmitted)
	if err := _OffchainAggregator.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OffchainAggregatorValidatorConfigSetIterator is returned from FilterValidatorConfigSet and is used to iterate over the raw logs and unpacked data for ValidatorConfigSet events raised by the OffchainAggregator contract.
type OffchainAggregatorValidatorConfigSetIterator struct {
	Event *OffchainAggregatorValidatorConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OffchainAggregatorValidatorConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffchainAggregatorValidatorConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OffchainAggregatorValidatorConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OffchainAggregatorValidatorConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OffchainAggregatorValidatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OffchainAggregatorValidatorConfigSet represents a ValidatorConfigSet event raised by the OffchainAggregator contract.
type OffchainAggregatorValidatorConfigSet struct {
	PreviousValidator common.Address
	PreviousGasLimit  uint32
	CurrentValidator  common.Address
	CurrentGasLimit   uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorConfigSet is a free log retrieval operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) FilterValidatorConfigSet(opts *bind.FilterOpts, previousValidator []common.Address, currentValidator []common.Address) (*OffchainAggregatorValidatorConfigSetIterator, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OffchainAggregator.contract.FilterLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return &OffchainAggregatorValidatorConfigSetIterator{contract: _OffchainAggregator.contract, event: "ValidatorConfigSet", logs: logs, sub: sub}, nil
}

// WatchValidatorConfigSet is a free log subscription operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) WatchValidatorConfigSet(opts *bind.WatchOpts, sink chan<- *OffchainAggregatorValidatorConfigSet, previousValidator []common.Address, currentValidator []common.Address) (event.Subscription, error) {

	var previousValidatorRule []interface{}
	for _, previousValidatorItem := range previousValidator {
		previousValidatorRule = append(previousValidatorRule, previousValidatorItem)
	}

	var currentValidatorRule []interface{}
	for _, currentValidatorItem := range currentValidator {
		currentValidatorRule = append(currentValidatorRule, currentValidatorItem)
	}

	logs, sub, err := _OffchainAggregator.contract.WatchLogs(opts, "ValidatorConfigSet", previousValidatorRule, currentValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OffchainAggregatorValidatorConfigSet)
				if err := _OffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorConfigSet is a log parse operation binding the contract event 0xb04e3a37abe9c0fcdfebdeae019a8e2b12ddf53f5d55ffb0caccc1bedaca1541.
//
// Solidity: event ValidatorConfigSet(address indexed previousValidator, uint32 previousGasLimit, address indexed currentValidator, uint32 currentGasLimit)
func (_OffchainAggregator *OffchainAggregatorFilterer) ParseValidatorConfigSet(log types.Log) (*OffchainAggregatorValidatorConfigSet, error) {
	event := new(OffchainAggregatorValidatorConfigSet)
	if err := _OffchainAggregator.contract.UnpackLog(event, "ValidatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
