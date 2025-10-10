// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logger_tester

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

var LoggerTesterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowedEmitters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emitLog\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"logCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owners\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllowedEmitter\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedEmitters\",\"inputs\":[{\"name\":\"users\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isOwner\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwners\",\"inputs\":[{\"name\":\"users\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"isOwner\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"LogEmitted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"manual\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedEmitter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b503360808190525f908152602081905260409020805460ff191660011790556080516109536100465f395f61019e01526109535ff3fe608060405234801561000f575f80fd5b50600436106100c4575f3560e01c80638034caf71161007d578063d5f3948811610058578063d5f3948814610199578063e3539993146101e5578063f42c13bf146101f8575f80fd5b80638034caf71461014d578063805f21321461016f578063ac42100f14610182575f80fd5b8063516c731c116100ad578063516c731c146101125780635a37c1b0146101275780636cdabe911461013a575f80fd5b806301ffc9a7146100c8578063022914a7146100f0575b5f80fd5b6100db6100d63660046106cd565b610200565b60405190151581526020015b60405180910390f35b6100db6100fe36600461073b565b5f6020819052908152604090205460ff1681565b610125610120366004610763565b610298565b005b610125610135366004610794565b61033c565b610125610148366004610794565b61041b565b6100db61015b36600461073b565b60016020525f908152604090205460ff1681565b61012561017d366004610856565b6104f3565b61018b60025481565b6040519081526020016100e7565b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e7565b6101256101f3366004610763565b610592565b610125610631565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f805f213200000000000000000000000000000000000000000000000000000000148061029257507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b335f9081526020819052604090205460ff166102e7576040517f245aecd30000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff919091165f90815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b335f9081526020819052604090205460ff16610386576040517f245aecd30000000000000000000000000000000000000000000000000000000081523360048201526024016102de565b5f5b82811015610415578160015f8686858181106103a6576103a66108bd565b90506020020160208101906103bb919061073b565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101610388565b50505050565b335f9081526020819052604090205460ff16610465576040517f245aecd30000000000000000000000000000000000000000000000000000000081523360048201526024016102de565b5f5b8281101561041557815f80868685818110610484576104846108bd565b9050602002016020810190610499919061073b565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101610467565b335f9081526001602052604090205460ff1661053d576040517f0667a3460000000000000000000000000000000000000000000000000000000081523360048201526024016102de565b60028054905f61054c836108ea565b9091555050600254604080519182525f60208301527fbe6d999fb18a595cd3c30b8f4eaa7461197546c42aef0599d395b58590f865b0910160405180910390a150505050565b335f9081526020819052604090205460ff166105dc576040517f245aecd30000000000000000000000000000000000000000000000000000000081523360048201526024016102de565b73ffffffffffffffffffffffffffffffffffffffff919091165f90815260016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b335f9081526001602052604090205460ff1661067b576040517f0667a3460000000000000000000000000000000000000000000000000000000081523360048201526024016102de565b60028054905f61068a836108ea565b909155505060025460408051918252600160208301527fbe6d999fb18a595cd3c30b8f4eaa7461197546c42aef0599d395b58590f865b0910160405180910390a1565b5f602082840312156106dd575f80fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461070c575f80fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610736575f80fd5b919050565b5f6020828403121561074b575f80fd5b61070c82610713565b80358015158114610736575f80fd5b5f8060408385031215610774575f80fd5b61077d83610713565b915061078b60208401610754565b90509250929050565b5f805f604084860312156107a6575f80fd5b833567ffffffffffffffff808211156107bd575f80fd5b818601915086601f8301126107d0575f80fd5b8135818111156107de575f80fd5b8760208260051b85010111156107f2575f80fd5b6020928301955093506108089186019050610754565b90509250925092565b5f8083601f840112610821575f80fd5b50813567ffffffffffffffff811115610838575f80fd5b60208301915083602082850101111561084f575f80fd5b9250929050565b5f805f8060408587031215610869575f80fd5b843567ffffffffffffffff80821115610880575f80fd5b61088c88838901610811565b909650945060208701359150808211156108a4575f80fd5b506108b187828801610811565b95989497509550505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361093f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fea164736f6c6343000818000a",
}

var LoggerTesterABI = LoggerTesterMetaData.ABI

var LoggerTesterBin = LoggerTesterMetaData.Bin

func DeployLoggerTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoggerTester, error) {
	parsed, err := LoggerTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoggerTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoggerTester{address: address, abi: *parsed, LoggerTesterCaller: LoggerTesterCaller{contract: contract}, LoggerTesterTransactor: LoggerTesterTransactor{contract: contract}, LoggerTesterFilterer: LoggerTesterFilterer{contract: contract}}, nil
}

type LoggerTester struct {
	address common.Address
	abi     abi.ABI
	LoggerTesterCaller
	LoggerTesterTransactor
	LoggerTesterFilterer
}

type LoggerTesterCaller struct {
	contract *bind.BoundContract
}

type LoggerTesterTransactor struct {
	contract *bind.BoundContract
}

type LoggerTesterFilterer struct {
	contract *bind.BoundContract
}

type LoggerTesterSession struct {
	Contract     *LoggerTester
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LoggerTesterCallerSession struct {
	Contract *LoggerTesterCaller
	CallOpts bind.CallOpts
}

type LoggerTesterTransactorSession struct {
	Contract     *LoggerTesterTransactor
	TransactOpts bind.TransactOpts
}

type LoggerTesterRaw struct {
	Contract *LoggerTester
}

type LoggerTesterCallerRaw struct {
	Contract *LoggerTesterCaller
}

type LoggerTesterTransactorRaw struct {
	Contract *LoggerTesterTransactor
}

func NewLoggerTester(address common.Address, backend bind.ContractBackend) (*LoggerTester, error) {
	abi, err := abi.JSON(strings.NewReader(LoggerTesterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLoggerTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoggerTester{address: address, abi: abi, LoggerTesterCaller: LoggerTesterCaller{contract: contract}, LoggerTesterTransactor: LoggerTesterTransactor{contract: contract}, LoggerTesterFilterer: LoggerTesterFilterer{contract: contract}}, nil
}

func NewLoggerTesterCaller(address common.Address, caller bind.ContractCaller) (*LoggerTesterCaller, error) {
	contract, err := bindLoggerTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTesterCaller{contract: contract}, nil
}

func NewLoggerTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggerTesterTransactor, error) {
	contract, err := bindLoggerTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTesterTransactor{contract: contract}, nil
}

func NewLoggerTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggerTesterFilterer, error) {
	contract, err := bindLoggerTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggerTesterFilterer{contract: contract}, nil
}

func bindLoggerTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoggerTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LoggerTester *LoggerTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoggerTester.Contract.LoggerTesterCaller.contract.Call(opts, result, method, params...)
}

func (_LoggerTester *LoggerTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoggerTester.Contract.LoggerTesterTransactor.contract.Transfer(opts)
}

func (_LoggerTester *LoggerTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoggerTester.Contract.LoggerTesterTransactor.contract.Transact(opts, method, params...)
}

func (_LoggerTester *LoggerTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoggerTester.Contract.contract.Call(opts, result, method, params...)
}

func (_LoggerTester *LoggerTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoggerTester.Contract.contract.Transfer(opts)
}

func (_LoggerTester *LoggerTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoggerTester.Contract.contract.Transact(opts, method, params...)
}

func (_LoggerTester *LoggerTesterCaller) AllowedEmitters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LoggerTester.contract.Call(opts, &out, "allowedEmitters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LoggerTester *LoggerTesterSession) AllowedEmitters(arg0 common.Address) (bool, error) {
	return _LoggerTester.Contract.AllowedEmitters(&_LoggerTester.CallOpts, arg0)
}

func (_LoggerTester *LoggerTesterCallerSession) AllowedEmitters(arg0 common.Address) (bool, error) {
	return _LoggerTester.Contract.AllowedEmitters(&_LoggerTester.CallOpts, arg0)
}

func (_LoggerTester *LoggerTesterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoggerTester.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LoggerTester *LoggerTesterSession) Deployer() (common.Address, error) {
	return _LoggerTester.Contract.Deployer(&_LoggerTester.CallOpts)
}

func (_LoggerTester *LoggerTesterCallerSession) Deployer() (common.Address, error) {
	return _LoggerTester.Contract.Deployer(&_LoggerTester.CallOpts)
}

func (_LoggerTester *LoggerTesterCaller) LogCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoggerTester.contract.Call(opts, &out, "logCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LoggerTester *LoggerTesterSession) LogCounter() (*big.Int, error) {
	return _LoggerTester.Contract.LogCounter(&_LoggerTester.CallOpts)
}

func (_LoggerTester *LoggerTesterCallerSession) LogCounter() (*big.Int, error) {
	return _LoggerTester.Contract.LogCounter(&_LoggerTester.CallOpts)
}

func (_LoggerTester *LoggerTesterCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LoggerTester.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LoggerTester *LoggerTesterSession) Owners(arg0 common.Address) (bool, error) {
	return _LoggerTester.Contract.Owners(&_LoggerTester.CallOpts, arg0)
}

func (_LoggerTester *LoggerTesterCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _LoggerTester.Contract.Owners(&_LoggerTester.CallOpts, arg0)
}

func (_LoggerTester *LoggerTesterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LoggerTester.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LoggerTester *LoggerTesterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LoggerTester.Contract.SupportsInterface(&_LoggerTester.CallOpts, interfaceId)
}

func (_LoggerTester *LoggerTesterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LoggerTester.Contract.SupportsInterface(&_LoggerTester.CallOpts, interfaceId)
}

func (_LoggerTester *LoggerTesterTransactor) EmitLog(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "emitLog")
}

func (_LoggerTester *LoggerTesterSession) EmitLog() (*types.Transaction, error) {
	return _LoggerTester.Contract.EmitLog(&_LoggerTester.TransactOpts)
}

func (_LoggerTester *LoggerTesterTransactorSession) EmitLog() (*types.Transaction, error) {
	return _LoggerTester.Contract.EmitLog(&_LoggerTester.TransactOpts)
}

func (_LoggerTester *LoggerTesterTransactor) OnReport(opts *bind.TransactOpts, arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "onReport", arg0, arg1)
}

func (_LoggerTester *LoggerTesterSession) OnReport(arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _LoggerTester.Contract.OnReport(&_LoggerTester.TransactOpts, arg0, arg1)
}

func (_LoggerTester *LoggerTesterTransactorSession) OnReport(arg0 []byte, arg1 []byte) (*types.Transaction, error) {
	return _LoggerTester.Contract.OnReport(&_LoggerTester.TransactOpts, arg0, arg1)
}

func (_LoggerTester *LoggerTesterTransactor) SetAllowedEmitter(opts *bind.TransactOpts, user common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "setAllowedEmitter", user, isAllowed)
}

func (_LoggerTester *LoggerTesterSession) SetAllowedEmitter(user common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetAllowedEmitter(&_LoggerTester.TransactOpts, user, isAllowed)
}

func (_LoggerTester *LoggerTesterTransactorSession) SetAllowedEmitter(user common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetAllowedEmitter(&_LoggerTester.TransactOpts, user, isAllowed)
}

func (_LoggerTester *LoggerTesterTransactor) SetAllowedEmitters(opts *bind.TransactOpts, users []common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "setAllowedEmitters", users, isAllowed)
}

func (_LoggerTester *LoggerTesterSession) SetAllowedEmitters(users []common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetAllowedEmitters(&_LoggerTester.TransactOpts, users, isAllowed)
}

func (_LoggerTester *LoggerTesterTransactorSession) SetAllowedEmitters(users []common.Address, isAllowed bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetAllowedEmitters(&_LoggerTester.TransactOpts, users, isAllowed)
}

func (_LoggerTester *LoggerTesterTransactor) SetOwner(opts *bind.TransactOpts, user common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "setOwner", user, isOwner)
}

func (_LoggerTester *LoggerTesterSession) SetOwner(user common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetOwner(&_LoggerTester.TransactOpts, user, isOwner)
}

func (_LoggerTester *LoggerTesterTransactorSession) SetOwner(user common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetOwner(&_LoggerTester.TransactOpts, user, isOwner)
}

func (_LoggerTester *LoggerTesterTransactor) SetOwners(opts *bind.TransactOpts, users []common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.contract.Transact(opts, "setOwners", users, isOwner)
}

func (_LoggerTester *LoggerTesterSession) SetOwners(users []common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetOwners(&_LoggerTester.TransactOpts, users, isOwner)
}

func (_LoggerTester *LoggerTesterTransactorSession) SetOwners(users []common.Address, isOwner bool) (*types.Transaction, error) {
	return _LoggerTester.Contract.SetOwners(&_LoggerTester.TransactOpts, users, isOwner)
}

type LoggerTesterLogEmittedIterator struct {
	Event *LoggerTesterLogEmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LoggerTesterLogEmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggerTesterLogEmitted)
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
		it.Event = new(LoggerTesterLogEmitted)
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

func (it *LoggerTesterLogEmittedIterator) Error() error {
	return it.fail
}

func (it *LoggerTesterLogEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LoggerTesterLogEmitted struct {
	Id     *big.Int
	Manual bool
	Raw    types.Log
}

func (_LoggerTester *LoggerTesterFilterer) FilterLogEmitted(opts *bind.FilterOpts) (*LoggerTesterLogEmittedIterator, error) {

	logs, sub, err := _LoggerTester.contract.FilterLogs(opts, "LogEmitted")
	if err != nil {
		return nil, err
	}
	return &LoggerTesterLogEmittedIterator{contract: _LoggerTester.contract, event: "LogEmitted", logs: logs, sub: sub}, nil
}

func (_LoggerTester *LoggerTesterFilterer) WatchLogEmitted(opts *bind.WatchOpts, sink chan<- *LoggerTesterLogEmitted) (event.Subscription, error) {

	logs, sub, err := _LoggerTester.contract.WatchLogs(opts, "LogEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LoggerTesterLogEmitted)
				if err := _LoggerTester.contract.UnpackLog(event, "LogEmitted", log); err != nil {
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

func (_LoggerTester *LoggerTesterFilterer) ParseLogEmitted(log types.Log) (*LoggerTesterLogEmitted, error) {
	event := new(LoggerTesterLogEmitted)
	if err := _LoggerTester.contract.UnpackLog(event, "LogEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_LoggerTester *LoggerTester) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LoggerTester.abi.Events["LogEmitted"].ID:
		return _LoggerTester.ParseLogEmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LoggerTesterLogEmitted) Topic() common.Hash {
	return common.HexToHash("0xbe6d999fb18a595cd3c30b8f4eaa7461197546c42aef0599d395b58590f865b0")
}

func (_LoggerTester *LoggerTester) Address() common.Address {
	return _LoggerTester.address
}

type LoggerTesterInterface interface {
	AllowedEmitters(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	Deployer(opts *bind.CallOpts) (common.Address, error)

	LogCounter(opts *bind.CallOpts) (*big.Int, error)

	Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	EmitLog(opts *bind.TransactOpts) (*types.Transaction, error)

	OnReport(opts *bind.TransactOpts, arg0 []byte, arg1 []byte) (*types.Transaction, error)

	SetAllowedEmitter(opts *bind.TransactOpts, user common.Address, isAllowed bool) (*types.Transaction, error)

	SetAllowedEmitters(opts *bind.TransactOpts, users []common.Address, isAllowed bool) (*types.Transaction, error)

	SetOwner(opts *bind.TransactOpts, user common.Address, isOwner bool) (*types.Transaction, error)

	SetOwners(opts *bind.TransactOpts, users []common.Address, isOwner bool) (*types.Transaction, error)

	FilterLogEmitted(opts *bind.FilterOpts) (*LoggerTesterLogEmittedIterator, error)

	WatchLogEmitted(opts *bind.WatchOpts, sink chan<- *LoggerTesterLogEmitted) (event.Subscription, error)

	ParseLogEmitted(log types.Log) (*LoggerTesterLogEmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
