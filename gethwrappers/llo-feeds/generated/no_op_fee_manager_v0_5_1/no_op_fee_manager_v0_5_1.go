// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package no_op_fee_manager_v0_5_1

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

type CommonAddressAndWeight struct {
	Addr   common.Address
	Weight uint64
}

type CommonAsset struct {
	AssetAddress common.Address
	Amount       *big.Int
}

var NoOpFeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_linkAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nativeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFeeAndReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"reward\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"appliedDiscount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"i_linkAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_nativeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_rewardManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkAvailableForPayment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"payLinkDeficit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processFee\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"processFeeBulk\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"s_globalDiscounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"s_nativeSurcharge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"s_subscriberDiscounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setFeeRecipients\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCommon.AddressAndWeight[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNativeSurcharge\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSubscriberDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSubscriberGlobalDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610cd4380380610cd483398101604081905261002f91610068565b6001600160a01b0392831660805290821660a0521660c0526100ab565b80516001600160a01b038116811461006357600080fd5b919050565b60008060006060848603121561007d57600080fd5b6100868461004c565b92506100946020850161004c565b91506100a26040850161004c565b90509250925092565b60805160a05160c051610bfa6100da600039600061022f015260006102a6015260006104190152610bfa6000f3fe6080604052600436106101295760003560e01c806376cf3187116100a5578063dba45fe011610074578063e389d9a411610059578063e389d9a4146103ec578063ea4b861b14610407578063f65df9621461043b57600080fd5b8063dba45fe014610342578063e03dab1a1461035057600080fd5b806376cf3187146102db57806387d6d843146102f6578063ce7817d114610321578063d09dc3391461020957600080fd5b806332f5f746116100fc57806350538094116100e1578063505380941461027657806363878668146102945780636c2f1a17146102c857600080fd5b806332f5f746146102095780633aa5ac071461021d57600080fd5b806301ffc9a71461012e578063181f5a77146101635780631cc7f2d8146101af5780631d4d84a2146101e7575b600080fd5b34801561013a57600080fd5b5061014e610149366004610687565b610456565b60405190151581526020015b60405180910390f35b34801561016f57600080fd5b50604080518082018252601481527f4e6f4f704665654d616e6167657220302e352e310000000000000000000000006020820152905161015a91906106d0565b3480156101bb57600080fd5b506101d96101ca366004610765565b670de0b6b3a764000092915050565b60405190815260200161015a565b3480156101f357600080fd5b50610207610202366004610798565b505050565b005b34801561021557600080fd5b5060006101d9565b34801561022957600080fd5b506102517f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b34801561028257600080fd5b50610207610291366004610814565b50565b3480156102a057600080fd5b506102517f000000000000000000000000000000000000000000000000000000000000000081565b6102076102d6366004610878565b6105d3565b3480156102e757600080fd5b50610207610202366004610927565b34801561030257600080fd5b506101d961031136600461096a565b670de0b6b3a76400009392505050565b34801561032d57600080fd5b5061020761033c36600461099d565b50505050565b6102076102d63660046109ea565b34801561035c57600080fd5b506103a161036b366004610a6e565b505060408051808201825260008082526020808301829052835180850190945281845283015292909150670de0b6b3a764000090565b60408051845173ffffffffffffffffffffffffffffffffffffffff9081168252602095860151868301528451169181019190915292909101516060830152608082015260a00161015a565b3480156103f857600080fd5b50610207610291366004610b55565b34801561041357600080fd5b506102517f000000000000000000000000000000000000000000000000000000000000000081565b34801561044757600080fd5b50610207610202366004610b6e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a70000000000000000000000000000000000000000000000000000000014806104e957507fffffffff0000000000000000000000000000000000000000000000000000000082167f268093e700000000000000000000000000000000000000000000000000000000145b8061053557507fffffffff0000000000000000000000000000000000000000000000000000000082167f41d6bc9500000000000000000000000000000000000000000000000000000000145b8061058157507fffffffff0000000000000000000000000000000000000000000000000000000082167fdba45fe000000000000000000000000000000000000000000000000000000000145b806105cd57507fffffffff0000000000000000000000000000000000000000000000000000000082167f6c2f1a1700000000000000000000000000000000000000000000000000000000145b92915050565b6105dc816105e3565b5050505050565b34156102915760008173ffffffffffffffffffffffffffffffffffffffff163460405160006040518083038185875af1925050503d8060008114610643576040519150601f19603f3d011682016040523d82523d6000602084013e610648565b606091505b5050905080610683576040517ff0c49d4400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60006020828403121561069957600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146106c957600080fd5b9392505050565b600060208083528351808285015260005b818110156106fd578581018301518582016040015282016106e1565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461076057600080fd5b919050565b6000806040838503121561077857600080fd5b6107818361073c565b915061078f6020840161073c565b90509250929050565b6000806000606084860312156107ad57600080fd5b6107b68461073c565b92506107c46020850161073c565b9150604084013577ffffffffffffffffffffffffffffffffffffffffffffffff811681146107f157600080fd5b809150509250925092565b803567ffffffffffffffff8116811461076057600080fd5b60006020828403121561082657600080fd5b6106c9826107fc565b60008083601f84011261084157600080fd5b50813567ffffffffffffffff81111561085957600080fd5b60208301915083602082850101111561087157600080fd5b9250929050565b60008060008060006060868803121561089057600080fd5b853567ffffffffffffffff808211156108a857600080fd5b818801915088601f8301126108bc57600080fd5b8135818111156108cb57600080fd5b8960208260051b85010111156108e057600080fd5b6020928301975095509087013590808211156108fb57600080fd5b506109088882890161082f565b909450925061091b90506040870161073c565b90509295509295909350565b60008060006060848603121561093c57600080fd5b6109458461073c565b92506109536020850161073c565b9150610961604085016107fc565b90509250925092565b60008060006060848603121561097f57600080fd5b6109888461073c565b9250602084013591506109616040850161073c565b600080600080608085870312156109b357600080fd5b6109bc8561073c565b9350602085013592506109d16040860161073c565b91506109df606086016107fc565b905092959194509250565b600080600080600060608688031215610a0257600080fd5b853567ffffffffffffffff80821115610a1a57600080fd5b610a2689838a0161082f565b909750955060208801359150808211156108fb57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215610a8357600080fd5b610a8c8461073c565b9250602084013567ffffffffffffffff80821115610aa957600080fd5b818601915086601f830112610abd57600080fd5b813581811115610acf57610acf610a3f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610b1557610b15610a3f565b81604052828152896020848701011115610b2e57600080fd5b8260208601602083013760006020848301015280965050505050506109616040850161073c565b600060208284031215610b6757600080fd5b5035919050565b600080600060408486031215610b8357600080fd5b83359250602084013567ffffffffffffffff80821115610ba257600080fd5b818601915086601f830112610bb657600080fd5b813581811115610bc557600080fd5b8760208260061b8501011115610bda57600080fd5b602083019450809350505050925092509256fea164736f6c6343000813000a",
}

var NoOpFeeManagerABI = NoOpFeeManagerMetaData.ABI

var NoOpFeeManagerBin = NoOpFeeManagerMetaData.Bin

func DeployNoOpFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _linkAddress common.Address, _nativeAddress common.Address, _rewardManagerAddress common.Address) (common.Address, *types.Transaction, *NoOpFeeManager, error) {
	parsed, err := NoOpFeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoOpFeeManagerBin), backend, _linkAddress, _nativeAddress, _rewardManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpFeeManager{address: address, abi: *parsed, NoOpFeeManagerCaller: NoOpFeeManagerCaller{contract: contract}, NoOpFeeManagerTransactor: NoOpFeeManagerTransactor{contract: contract}, NoOpFeeManagerFilterer: NoOpFeeManagerFilterer{contract: contract}}, nil
}

type NoOpFeeManager struct {
	address common.Address
	abi     abi.ABI
	NoOpFeeManagerCaller
	NoOpFeeManagerTransactor
	NoOpFeeManagerFilterer
}

type NoOpFeeManagerCaller struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerTransactor struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerFilterer struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerSession struct {
	Contract     *NoOpFeeManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NoOpFeeManagerCallerSession struct {
	Contract *NoOpFeeManagerCaller
	CallOpts bind.CallOpts
}

type NoOpFeeManagerTransactorSession struct {
	Contract     *NoOpFeeManagerTransactor
	TransactOpts bind.TransactOpts
}

type NoOpFeeManagerRaw struct {
	Contract *NoOpFeeManager
}

type NoOpFeeManagerCallerRaw struct {
	Contract *NoOpFeeManagerCaller
}

type NoOpFeeManagerTransactorRaw struct {
	Contract *NoOpFeeManagerTransactor
}

func NewNoOpFeeManager(address common.Address, backend bind.ContractBackend) (*NoOpFeeManager, error) {
	abi, err := abi.JSON(strings.NewReader(NoOpFeeManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNoOpFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManager{address: address, abi: abi, NoOpFeeManagerCaller: NoOpFeeManagerCaller{contract: contract}, NoOpFeeManagerTransactor: NoOpFeeManagerTransactor{contract: contract}, NoOpFeeManagerFilterer: NoOpFeeManagerFilterer{contract: contract}}, nil
}

func NewNoOpFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*NoOpFeeManagerCaller, error) {
	contract, err := bindNoOpFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerCaller{contract: contract}, nil
}

func NewNoOpFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOpFeeManagerTransactor, error) {
	contract, err := bindNoOpFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerTransactor{contract: contract}, nil
}

func NewNoOpFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOpFeeManagerFilterer, error) {
	contract, err := bindNoOpFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerFilterer{contract: contract}, nil
}

func bindNoOpFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoOpFeeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpFeeManager.Contract.NoOpFeeManagerCaller.contract.Call(opts, result, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.NoOpFeeManagerTransactor.contract.Transfer(opts)
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.NoOpFeeManagerTransactor.contract.Transact(opts, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpFeeManager.Contract.contract.Call(opts, result, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.contract.Transfer(opts)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.contract.Transact(opts, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) GetFeeAndReward(opts *bind.CallOpts, arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "getFeeAndReward", arg0, arg1, arg2)

	outstruct := new(GetFeeAndReward)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(CommonAsset)).(*CommonAsset)
	outstruct.Reward = *abi.ConvertType(out[1], new(CommonAsset)).(*CommonAsset)
	outstruct.AppliedDiscount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) GetFeeAndReward(arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	return _NoOpFeeManager.Contract.GetFeeAndReward(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) GetFeeAndReward(arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	return _NoOpFeeManager.Contract.GetFeeAndReward(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) ILinkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "i_linkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) ILinkAddress() (common.Address, error) {
	return _NoOpFeeManager.Contract.ILinkAddress(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) ILinkAddress() (common.Address, error) {
	return _NoOpFeeManager.Contract.ILinkAddress(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) INativeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "i_nativeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) INativeAddress() (common.Address, error) {
	return _NoOpFeeManager.Contract.INativeAddress(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) INativeAddress() (common.Address, error) {
	return _NoOpFeeManager.Contract.INativeAddress(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) IRewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "i_rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) IRewardManager() (common.Address, error) {
	return _NoOpFeeManager.Contract.IRewardManager(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) IRewardManager() (common.Address, error) {
	return _NoOpFeeManager.Contract.IRewardManager(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _NoOpFeeManager.Contract.LinkAvailableForPayment(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _NoOpFeeManager.Contract.LinkAvailableForPayment(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SGlobalDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "s_globalDiscounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SGlobalDiscounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SGlobalDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SGlobalDiscounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SGlobalDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SNativeSurcharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "s_nativeSurcharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SNativeSurcharge() (*big.Int, error) {
	return _NoOpFeeManager.Contract.SNativeSurcharge(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SNativeSurcharge() (*big.Int, error) {
	return _NoOpFeeManager.Contract.SNativeSurcharge(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SSubscriberDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "s_subscriberDiscounts", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SSubscriberDiscounts(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SSubscriberDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SSubscriberDiscounts(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SSubscriberDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NoOpFeeManager.Contract.SupportsInterface(&_NoOpFeeManager.CallOpts, interfaceId)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NoOpFeeManager.Contract.SupportsInterface(&_NoOpFeeManager.CallOpts, interfaceId)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) TypeAndVersion() (string, error) {
	return _NoOpFeeManager.Contract.TypeAndVersion(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) TypeAndVersion() (string, error) {
	return _NoOpFeeManager.Contract.TypeAndVersion(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) PayLinkDeficit(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "payLinkDeficit", arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) PayLinkDeficit(arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.PayLinkDeficit(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) PayLinkDeficit(arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.PayLinkDeficit(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) ProcessFee(opts *bind.TransactOpts, arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "processFee", arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) ProcessFee(arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFee(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) ProcessFee(arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFee(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) ProcessFeeBulk(opts *bind.TransactOpts, arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "processFeeBulk", arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) ProcessFeeBulk(arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFeeBulk(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) ProcessFeeBulk(arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFeeBulk(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) SetFeeRecipients(opts *bind.TransactOpts, arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "setFeeRecipients", arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) SetFeeRecipients(arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetFeeRecipients(&_NoOpFeeManager.TransactOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) SetFeeRecipients(arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetFeeRecipients(&_NoOpFeeManager.TransactOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) SetNativeSurcharge(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "setNativeSurcharge", arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) SetNativeSurcharge(arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetNativeSurcharge(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) SetNativeSurcharge(arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetNativeSurcharge(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) UpdateSubscriberDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "updateSubscriberDiscount", arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) UpdateSubscriberDiscount(arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) UpdateSubscriberDiscount(arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) UpdateSubscriberGlobalDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "updateSubscriberGlobalDiscount", arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) UpdateSubscriberGlobalDiscount(arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberGlobalDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) UpdateSubscriberGlobalDiscount(arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberGlobalDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) Withdraw(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "withdraw", arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) Withdraw(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.Withdraw(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) Withdraw(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.Withdraw(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

type GetFeeAndReward struct {
	Fee             CommonAsset
	Reward          CommonAsset
	AppliedDiscount *big.Int
}

func (_NoOpFeeManager *NoOpFeeManager) Address() common.Address {
	return _NoOpFeeManager.address
}

type NoOpFeeManagerInterface interface {
	GetFeeAndReward(opts *bind.CallOpts, arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

		error)

	ILinkAddress(opts *bind.CallOpts) (common.Address, error)

	INativeAddress(opts *bind.CallOpts) (common.Address, error)

	IRewardManager(opts *bind.CallOpts) (common.Address, error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	SGlobalDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	SNativeSurcharge(opts *bind.CallOpts) (*big.Int, error)

	SSubscriberDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	PayLinkDeficit(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error)

	ProcessFee(opts *bind.TransactOpts, arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error)

	ProcessFeeBulk(opts *bind.TransactOpts, arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error)

	SetFeeRecipients(opts *bind.TransactOpts, arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error)

	SetNativeSurcharge(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error)

	UpdateSubscriberDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error)

	UpdateSubscriberGlobalDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error)

	Address() common.Address
}
