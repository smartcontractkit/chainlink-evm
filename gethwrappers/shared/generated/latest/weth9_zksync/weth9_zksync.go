// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weth9_zksync

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generated"
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

var WETH9ZKSyncMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"guy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d3565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57600081815260208120601f850160051c810160208610156100ec5750805b601f850160051c820191505b8181101561010b578281556001016100f8565b5050505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b602080601f831160018114610176576000841561015e5750858301515b600019600386901b1c1916600185901b17855561010b565b600085815260208120601f198616915b828110156101a557888601518255948401946001909101908401610186565b50858210156101c35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610935806101e26000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f6919061071e565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107b3565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f6101673660046107dd565b6103bd565b34801561017857600080fd5b506100cd610187366004610819565b6105c4565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d3366004610832565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e96106f3565b34801561020657600080fd5b5061011f6102153660046107b3565b610700565b6100cd610714565b34801561022e57600080fd5b5061013e61023d36600461084d565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108af565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108c2565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108c2565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120548211156103ef57600080fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610455575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020546fffffffffffffffffffffffffffffffff14155b156104dd5773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561049757600080fd5b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152812080548492906104d7908490610915565b90915550505b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604081208054849290610512908490610915565b909155505073ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120805484929061054c9084906108af565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516105b291815260200190565b60405180910390a35060019392505050565b336000908152600360205260409020548111156105e057600080fd5b33600090815260036020526040812080548392906105ff908490610915565b9091555050604051600090339083908381818185875af1925050503d8060008114610646576040519150601f19603f3d011682016040523d82523d6000602084013e61064b565b606091505b50509050806106ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5472616e73666572206661696c65640000000000000000000000000000000000604482015260640160405180910390fd5b60405182815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a25050565b600180546102c2906108c2565b600061070d3384846103bd565b9392505050565b61071c61025a565b565b600060208083528351808285015260005b8181101561074b5785810183015185820160400152820161072f565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107ae57600080fd5b919050565b600080604083850312156107c657600080fd5b6107cf8361078a565b946020939093013593505050565b6000806000606084860312156107f257600080fd5b6107fb8461078a565b92506108096020850161078a565b9150604084013590509250925092565b60006020828403121561082b57600080fd5b5035919050565b60006020828403121561084457600080fd5b61070d8261078a565b6000806040838503121561086057600080fd5b6108698361078a565b91506108776020840161078a565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b7610880565b600181811c908216806108d657607f821691505b60208210810361090f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b761088056fea164736f6c6343000813000a",
}

var WETH9ZKSyncABI = WETH9ZKSyncMetaData.ABI

var WETH9ZKSyncBin = WETH9ZKSyncMetaData.Bin

func DeployWETH9ZKSync(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WETH9ZKSync, error) {
	parsed, err := WETH9ZKSyncMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WETH9ZKSyncBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETH9ZKSync{address: address, abi: *parsed, WETH9ZKSyncCaller: WETH9ZKSyncCaller{contract: contract}, WETH9ZKSyncTransactor: WETH9ZKSyncTransactor{contract: contract}, WETH9ZKSyncFilterer: WETH9ZKSyncFilterer{contract: contract}}, nil
}

type WETH9ZKSync struct {
	address common.Address
	abi     abi.ABI
	WETH9ZKSyncCaller
	WETH9ZKSyncTransactor
	WETH9ZKSyncFilterer
}

type WETH9ZKSyncCaller struct {
	contract *bind.BoundContract
}

type WETH9ZKSyncTransactor struct {
	contract *bind.BoundContract
}

type WETH9ZKSyncFilterer struct {
	contract *bind.BoundContract
}

type WETH9ZKSyncSession struct {
	Contract     *WETH9ZKSync
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WETH9ZKSyncCallerSession struct {
	Contract *WETH9ZKSyncCaller
	CallOpts bind.CallOpts
}

type WETH9ZKSyncTransactorSession struct {
	Contract     *WETH9ZKSyncTransactor
	TransactOpts bind.TransactOpts
}

type WETH9ZKSyncRaw struct {
	Contract *WETH9ZKSync
}

type WETH9ZKSyncCallerRaw struct {
	Contract *WETH9ZKSyncCaller
}

type WETH9ZKSyncTransactorRaw struct {
	Contract *WETH9ZKSyncTransactor
}

func NewWETH9ZKSync(address common.Address, backend bind.ContractBackend) (*WETH9ZKSync, error) {
	abi, err := abi.JSON(strings.NewReader(WETH9ZKSyncABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWETH9ZKSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSync{address: address, abi: abi, WETH9ZKSyncCaller: WETH9ZKSyncCaller{contract: contract}, WETH9ZKSyncTransactor: WETH9ZKSyncTransactor{contract: contract}, WETH9ZKSyncFilterer: WETH9ZKSyncFilterer{contract: contract}}, nil
}

func NewWETH9ZKSyncCaller(address common.Address, caller bind.ContractCaller) (*WETH9ZKSyncCaller, error) {
	contract, err := bindWETH9ZKSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncCaller{contract: contract}, nil
}

func NewWETH9ZKSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*WETH9ZKSyncTransactor, error) {
	contract, err := bindWETH9ZKSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncTransactor{contract: contract}, nil
}

func NewWETH9ZKSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*WETH9ZKSyncFilterer, error) {
	contract, err := bindWETH9ZKSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncFilterer{contract: contract}, nil
}

func bindWETH9ZKSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETH9ZKSyncMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WETH9ZKSync *WETH9ZKSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9ZKSync.Contract.WETH9ZKSyncCaller.contract.Call(opts, result, method, params...)
}

func (_WETH9ZKSync *WETH9ZKSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.WETH9ZKSyncTransactor.contract.Transfer(opts)
}

func (_WETH9ZKSync *WETH9ZKSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.WETH9ZKSyncTransactor.contract.Transact(opts, method, params...)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9ZKSync.Contract.contract.Call(opts, result, method, params...)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.contract.Transfer(opts)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.contract.Transact(opts, method, params...)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH9ZKSync.Contract.Allowance(&_WETH9ZKSync.CallOpts, arg0, arg1)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH9ZKSync.Contract.Allowance(&_WETH9ZKSync.CallOpts, arg0, arg1)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH9ZKSync.Contract.BalanceOf(&_WETH9ZKSync.CallOpts, arg0)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH9ZKSync.Contract.BalanceOf(&_WETH9ZKSync.CallOpts, arg0)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) Decimals() (uint8, error) {
	return _WETH9ZKSync.Contract.Decimals(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) Decimals() (uint8, error) {
	return _WETH9ZKSync.Contract.Decimals(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) Name() (string, error) {
	return _WETH9ZKSync.Contract.Name(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) Name() (string, error) {
	return _WETH9ZKSync.Contract.Name(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) Symbol() (string, error) {
	return _WETH9ZKSync.Contract.Symbol(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) Symbol() (string, error) {
	return _WETH9ZKSync.Contract.Symbol(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH9ZKSync.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9ZKSync *WETH9ZKSyncSession) TotalSupply() (*big.Int, error) {
	return _WETH9ZKSync.Contract.TotalSupply(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncCallerSession) TotalSupply() (*big.Int, error) {
	return _WETH9ZKSync.Contract.TotalSupply(&_WETH9ZKSync.CallOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.Transact(opts, "approve", guy, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Approve(&_WETH9ZKSync.TransactOpts, guy, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Approve(&_WETH9ZKSync.TransactOpts, guy, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.Transact(opts, "deposit")
}

func (_WETH9ZKSync *WETH9ZKSyncSession) Deposit() (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Deposit(&_WETH9ZKSync.TransactOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) Deposit() (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Deposit(&_WETH9ZKSync.TransactOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.Transact(opts, "transfer", dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Transfer(&_WETH9ZKSync.TransactOpts, dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Transfer(&_WETH9ZKSync.TransactOpts, dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.Transact(opts, "transferFrom", src, dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.TransferFrom(&_WETH9ZKSync.TransactOpts, src, dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.TransferFrom(&_WETH9ZKSync.TransactOpts, src, dst, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.Transact(opts, "withdraw", wad)
}

func (_WETH9ZKSync *WETH9ZKSyncSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Withdraw(&_WETH9ZKSync.TransactOpts, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Withdraw(&_WETH9ZKSync.TransactOpts, wad)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9ZKSync.contract.RawTransact(opts, nil)
}

func (_WETH9ZKSync *WETH9ZKSyncSession) Receive() (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Receive(&_WETH9ZKSync.TransactOpts)
}

func (_WETH9ZKSync *WETH9ZKSyncTransactorSession) Receive() (*types.Transaction, error) {
	return _WETH9ZKSync.Contract.Receive(&_WETH9ZKSync.TransactOpts)
}

type WETH9ZKSyncApprovalIterator struct {
	Event *WETH9ZKSyncApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9ZKSyncApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9ZKSyncApproval)
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

	select {
	case log := <-it.logs:
		it.Event = new(WETH9ZKSyncApproval)
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

func (it *WETH9ZKSyncApprovalIterator) Error() error {
	return it.fail
}

func (it *WETH9ZKSyncApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9ZKSyncApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WETH9ZKSyncApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncApprovalIterator{contract: _WETH9ZKSync.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9ZKSyncApproval)
				if err := _WETH9ZKSync.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_WETH9ZKSync *WETH9ZKSyncFilterer) ParseApproval(log types.Log) (*WETH9ZKSyncApproval, error) {
	event := new(WETH9ZKSyncApproval)
	if err := _WETH9ZKSync.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9ZKSyncDepositIterator struct {
	Event *WETH9ZKSyncDeposit

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9ZKSyncDepositIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9ZKSyncDeposit)
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

	select {
	case log := <-it.logs:
		it.Event = new(WETH9ZKSyncDeposit)
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

func (it *WETH9ZKSyncDepositIterator) Error() error {
	return it.fail
}

func (it *WETH9ZKSyncDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9ZKSyncDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WETH9ZKSyncDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncDepositIterator{contract: _WETH9ZKSync.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9ZKSyncDeposit)
				if err := _WETH9ZKSync.contract.UnpackLog(event, "Deposit", log); err != nil {
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

func (_WETH9ZKSync *WETH9ZKSyncFilterer) ParseDeposit(log types.Log) (*WETH9ZKSyncDeposit, error) {
	event := new(WETH9ZKSyncDeposit)
	if err := _WETH9ZKSync.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9ZKSyncTransferIterator struct {
	Event *WETH9ZKSyncTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9ZKSyncTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9ZKSyncTransfer)
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

	select {
	case log := <-it.logs:
		it.Event = new(WETH9ZKSyncTransfer)
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

func (it *WETH9ZKSyncTransferIterator) Error() error {
	return it.fail
}

func (it *WETH9ZKSyncTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9ZKSyncTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WETH9ZKSyncTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncTransferIterator{contract: _WETH9ZKSync.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9ZKSyncTransfer)
				if err := _WETH9ZKSync.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_WETH9ZKSync *WETH9ZKSyncFilterer) ParseTransfer(log types.Log) (*WETH9ZKSyncTransfer, error) {
	event := new(WETH9ZKSyncTransfer)
	if err := _WETH9ZKSync.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9ZKSyncWithdrawalIterator struct {
	Event *WETH9ZKSyncWithdrawal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9ZKSyncWithdrawalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9ZKSyncWithdrawal)
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

	select {
	case log := <-it.logs:
		it.Event = new(WETH9ZKSyncWithdrawal)
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

func (it *WETH9ZKSyncWithdrawalIterator) Error() error {
	return it.fail
}

func (it *WETH9ZKSyncWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9ZKSyncWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WETH9ZKSyncWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WETH9ZKSyncWithdrawalIterator{contract: _WETH9ZKSync.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

func (_WETH9ZKSync *WETH9ZKSyncFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WETH9ZKSync.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9ZKSyncWithdrawal)
				if err := _WETH9ZKSync.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

func (_WETH9ZKSync *WETH9ZKSyncFilterer) ParseWithdrawal(log types.Log) (*WETH9ZKSyncWithdrawal, error) {
	event := new(WETH9ZKSyncWithdrawal)
	if err := _WETH9ZKSync.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_WETH9ZKSync *WETH9ZKSync) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WETH9ZKSync.abi.Events["Approval"].ID:
		return _WETH9ZKSync.ParseApproval(log)
	case _WETH9ZKSync.abi.Events["Deposit"].ID:
		return _WETH9ZKSync.ParseDeposit(log)
	case _WETH9ZKSync.abi.Events["Transfer"].ID:
		return _WETH9ZKSync.ParseTransfer(log)
	case _WETH9ZKSync.abi.Events["Withdrawal"].ID:
		return _WETH9ZKSync.ParseWithdrawal(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WETH9ZKSyncApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (WETH9ZKSyncDeposit) Topic() common.Hash {
	return common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c")
}

func (WETH9ZKSyncTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (WETH9ZKSyncWithdrawal) Topic() common.Hash {
	return common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65")
}

func (_WETH9ZKSync *WETH9ZKSync) Address() common.Address {
	return _WETH9ZKSync.address
}

type WETH9ZKSyncInterface interface {
	Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Name(opts *bind.CallOpts) (string, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WETH9ZKSyncApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncApproval, src []common.Address, guy []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*WETH9ZKSyncApproval, error)

	FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WETH9ZKSyncDepositIterator, error)

	WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncDeposit, dst []common.Address) (event.Subscription, error)

	ParseDeposit(log types.Log) (*WETH9ZKSyncDeposit, error)

	FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WETH9ZKSyncTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncTransfer, src []common.Address, dst []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*WETH9ZKSyncTransfer, error)

	FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WETH9ZKSyncWithdrawalIterator, error)

	WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETH9ZKSyncWithdrawal, src []common.Address) (event.Subscription, error)

	ParseWithdrawal(log types.Log) (*WETH9ZKSyncWithdrawal, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
