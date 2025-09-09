// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_burn_mint_erc20_upgradeable

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

var IBurnMintERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

var IBurnMintERC20UpgradeableABI = IBurnMintERC20UpgradeableMetaData.ABI

type IBurnMintERC20Upgradeable struct {
	address common.Address
	abi     abi.ABI
	IBurnMintERC20UpgradeableCaller
	IBurnMintERC20UpgradeableTransactor
	IBurnMintERC20UpgradeableFilterer
}

type IBurnMintERC20UpgradeableCaller struct {
	contract *bind.BoundContract
}

type IBurnMintERC20UpgradeableTransactor struct {
	contract *bind.BoundContract
}

type IBurnMintERC20UpgradeableFilterer struct {
	contract *bind.BoundContract
}

type IBurnMintERC20UpgradeableSession struct {
	Contract     *IBurnMintERC20Upgradeable
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IBurnMintERC20UpgradeableCallerSession struct {
	Contract *IBurnMintERC20UpgradeableCaller
	CallOpts bind.CallOpts
}

type IBurnMintERC20UpgradeableTransactorSession struct {
	Contract     *IBurnMintERC20UpgradeableTransactor
	TransactOpts bind.TransactOpts
}

type IBurnMintERC20UpgradeableRaw struct {
	Contract *IBurnMintERC20Upgradeable
}

type IBurnMintERC20UpgradeableCallerRaw struct {
	Contract *IBurnMintERC20UpgradeableCaller
}

type IBurnMintERC20UpgradeableTransactorRaw struct {
	Contract *IBurnMintERC20UpgradeableTransactor
}

func NewIBurnMintERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*IBurnMintERC20Upgradeable, error) {
	abi, err := abi.JSON(strings.NewReader(IBurnMintERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindIBurnMintERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20Upgradeable{address: address, abi: abi, IBurnMintERC20UpgradeableCaller: IBurnMintERC20UpgradeableCaller{contract: contract}, IBurnMintERC20UpgradeableTransactor: IBurnMintERC20UpgradeableTransactor{contract: contract}, IBurnMintERC20UpgradeableFilterer: IBurnMintERC20UpgradeableFilterer{contract: contract}}, nil
}

func NewIBurnMintERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IBurnMintERC20UpgradeableCaller, error) {
	contract, err := bindIBurnMintERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20UpgradeableCaller{contract: contract}, nil
}

func NewIBurnMintERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IBurnMintERC20UpgradeableTransactor, error) {
	contract, err := bindIBurnMintERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20UpgradeableTransactor{contract: contract}, nil
}

func NewIBurnMintERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IBurnMintERC20UpgradeableFilterer, error) {
	contract, err := bindIBurnMintERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20UpgradeableFilterer{contract: contract}, nil
}

func bindIBurnMintERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBurnMintERC20UpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnMintERC20Upgradeable.Contract.IBurnMintERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.IBurnMintERC20UpgradeableTransactor.contract.Transfer(opts)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.IBurnMintERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnMintERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.contract.Transfer(opts)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBurnMintERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.Allowance(&_IBurnMintERC20Upgradeable.CallOpts, owner, spender)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.Allowance(&_IBurnMintERC20Upgradeable.CallOpts, owner, spender)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBurnMintERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.BalanceOf(&_IBurnMintERC20Upgradeable.CallOpts, account)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.BalanceOf(&_IBurnMintERC20Upgradeable.CallOpts, account)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBurnMintERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.TotalSupply(&_IBurnMintERC20Upgradeable.CallOpts)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IBurnMintERC20Upgradeable.Contract.TotalSupply(&_IBurnMintERC20Upgradeable.CallOpts)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "approve", spender, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Approve(&_IBurnMintERC20Upgradeable.TransactOpts, spender, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Approve(&_IBurnMintERC20Upgradeable.TransactOpts, spender, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "burn", amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Burn(&_IBurnMintERC20Upgradeable.TransactOpts, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Burn(&_IBurnMintERC20Upgradeable.TransactOpts, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "burn0", account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Burn0(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Burn0(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "burnFrom", account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.BurnFrom(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.BurnFrom(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "mint", account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Mint(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Mint(&_IBurnMintERC20Upgradeable.TransactOpts, account, amount)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "transfer", to, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Transfer(&_IBurnMintERC20Upgradeable.TransactOpts, to, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.Transfer(&_IBurnMintERC20Upgradeable.TransactOpts, to, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.TransferFrom(&_IBurnMintERC20Upgradeable.TransactOpts, from, to, value)
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IBurnMintERC20Upgradeable.Contract.TransferFrom(&_IBurnMintERC20Upgradeable.TransactOpts, from, to, value)
}

type IBurnMintERC20UpgradeableApprovalIterator struct {
	Event *IBurnMintERC20UpgradeableApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IBurnMintERC20UpgradeableApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnMintERC20UpgradeableApproval)
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
		it.Event = new(IBurnMintERC20UpgradeableApproval)
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

func (it *IBurnMintERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

func (it *IBurnMintERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IBurnMintERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IBurnMintERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBurnMintERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20UpgradeableApprovalIterator{contract: _IBurnMintERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IBurnMintERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBurnMintERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IBurnMintERC20UpgradeableApproval)
				if err := _IBurnMintERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) ParseApproval(log types.Log) (*IBurnMintERC20UpgradeableApproval, error) {
	event := new(IBurnMintERC20UpgradeableApproval)
	if err := _IBurnMintERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IBurnMintERC20UpgradeableTransferIterator struct {
	Event *IBurnMintERC20UpgradeableTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IBurnMintERC20UpgradeableTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnMintERC20UpgradeableTransfer)
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
		it.Event = new(IBurnMintERC20UpgradeableTransfer)
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

func (it *IBurnMintERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

func (it *IBurnMintERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IBurnMintERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBurnMintERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBurnMintERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBurnMintERC20UpgradeableTransferIterator{contract: _IBurnMintERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBurnMintERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBurnMintERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IBurnMintERC20UpgradeableTransfer)
				if err := _IBurnMintERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_IBurnMintERC20Upgradeable *IBurnMintERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*IBurnMintERC20UpgradeableTransfer, error) {
	event := new(IBurnMintERC20UpgradeableTransfer)
	if err := _IBurnMintERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20Upgradeable) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _IBurnMintERC20Upgradeable.abi.Events["Approval"].ID:
		return _IBurnMintERC20Upgradeable.ParseApproval(log)
	case _IBurnMintERC20Upgradeable.abi.Events["Transfer"].ID:
		return _IBurnMintERC20Upgradeable.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (IBurnMintERC20UpgradeableApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (IBurnMintERC20UpgradeableTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_IBurnMintERC20Upgradeable *IBurnMintERC20Upgradeable) Address() common.Address {
	return _IBurnMintERC20Upgradeable.address
}

type IBurnMintERC20UpgradeableInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IBurnMintERC20UpgradeableApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *IBurnMintERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*IBurnMintERC20UpgradeableApproval, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBurnMintERC20UpgradeableTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBurnMintERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*IBurnMintERC20UpgradeableTransfer, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
