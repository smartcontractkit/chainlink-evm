// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_emitter

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

var MessageEmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"emitMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastMessage\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessage\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessageEmitted\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
	Bin: "0x6080806040523461001657610a8d908161001c8239f35b600080fdfe608060408181526004908136101561001657600080fd5b600092833560e01c908163181f5a7714610830575080632ac0df261461036e578063570c537e146102105763e0c01bfe1461005057600080fd5b3461020c576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102085773ffffffffffffffffffffffffffffffffffffffff61009d6109ba565b16808552600193600183526100b4848720546109e2565b156101875750845260018152818420825193858254926100d3846109e2565b8088529360018116908115610145575060011461010a575b61010687876100fc828c03836108e4565b5191829182610954565b0390f35b9080949750528583205b8284106101325750505082610106946100fc928201019438806100eb565b8054868501880152928601928101610114565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168887015250505050151560051b83010192506100fc8261010638806100eb565b608490838551917f08c379a00000000000000000000000000000000000000000000000000000000083528201526024808201527f4e6f206c617374206d65737361676520666f722074686520676976656e20736560448201527f6e646572000000000000000000000000000000000000000000000000000000006064820152fd5b8380fd5b8280fd5b503461020c57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020c5761025361024a6109ba565b60243590610a35565b808452602092848452610268838620546109e2565b156102ec575083528282528083209181519284815490610287826109e2565b8087529260019280841690811561014557506001146102b15761010687876100fc828c03836108e4565b9080949750528583205b8284106102d95750505082610106946100fc928201019438806100eb565b80548685018801529286019281016102bb565b608490848451917f08c379a0000000000000000000000000000000000000000000000000000000008352820152603960248201527f4d65737361676520646f6573206e6f7420657869737420666f7220746865206760448201527f6976656e2073656e64657220616e642074696d657374616d70000000000000006064820152fd5b5091903461082c57602092837ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020c5781359367ffffffffffffffff918286116108285736602387011215610828578584013592831161082857602493368585890101116108245783156107ca576103ea4233610a35565b908187528684526103fd838820546109e2565b610749575085528482528085209361041585546109e2565b94601f95868111610709575b508585119085838960018514610661578a92610654575b50508660011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8860031b1c19161790555b33875260018085528388209161048183546109e2565b888111610611575b50889060011461054c579186808a9b9386829b9c967fc799f359194674b273986b8c03283265390f642b631c04e6526b99d0d8f4c38d9b98979461053f575b50501b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8660031b1c19161790555b82845195808752860152018284013785818584010152817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe04296339601168101030190a380f35b86010135925085386104c8565b828952858920907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe088168a5b8181106105fa57509188999a9b9391897fc799f359194674b273986b8c03283265390f642b631c04e6526b99d0d8f4c38d9a979694106105c0575b505084811b0190556104f8565b857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88960031b161c19918601013516905538806105b3565b8c8301870135845592840192918801918801610578565b838a52868a2089808a0160051c820192898b1061064b575b0160051c019083905b828110610640575050610489565b8b8155018390610632565b92508192610629565b8b01013590508338610438565b838b52878b20929091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081168b89888f5b8484106106ed5750505050106106b3575b5050600186811b01905561046b565b847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88a60031b161c19918c01013516905538806106a4565b860101358655600190950194938401938b93500189888f610693565b8188528488208780880160051c820192878910610740575b0160051c01905b8181106107355750610421565b888155600101610728565b92508192610721565b608490603e87868651937f08c379a00000000000000000000000000000000000000000000000000000000085528401528201527f4d65737361676520616c72656164792065786973747320666f7220746865207360448201527f616d652073656e64657220616e6420626c6f636b2074696d657374616d7000006064820152fd5b6017858460649451937f08c379a00000000000000000000000000000000000000000000000000000000085528401528201527f4d6573736167652063616e6e6f7420626520656d7074790000000000000000006044820152fd5b8580fd5b8480fd5b5080fd5b92939050346108e157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108e1578183019083821067ffffffffffffffff8311176108b5575061010693508152601882527f4d657373616765456d697474657220312e302e302d646576000000000000000060208301525191829182610954565b806041867f4e487b71000000000000000000000000000000000000000000000000000000006024945252fd5b80fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761092557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60208082528251818301819052939260005b8581106109a6575050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b818101830151848201604001528201610966565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036109dd57565b600080fd5b90600182811c92168015610a2b575b60208310146109fc57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916109f1565b906040519073ffffffffffffffffffffffffffffffffffffffff60208301931683526040820152604081526060810181811067ffffffffffffffff821117610925576040525190209056fea164736f6c6343000818000a",
}

var MessageEmitterABI = MessageEmitterMetaData.ABI

var MessageEmitterBin = MessageEmitterMetaData.Bin

func DeployMessageEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageEmitter, error) {
	parsed, err := MessageEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageEmitter{address: address, abi: *parsed, MessageEmitterCaller: MessageEmitterCaller{contract: contract}, MessageEmitterTransactor: MessageEmitterTransactor{contract: contract}, MessageEmitterFilterer: MessageEmitterFilterer{contract: contract}}, nil
}

type MessageEmitter struct {
	address common.Address
	abi     abi.ABI
	MessageEmitterCaller
	MessageEmitterTransactor
	MessageEmitterFilterer
}

type MessageEmitterCaller struct {
	contract *bind.BoundContract
}

type MessageEmitterTransactor struct {
	contract *bind.BoundContract
}

type MessageEmitterFilterer struct {
	contract *bind.BoundContract
}

type MessageEmitterSession struct {
	Contract     *MessageEmitter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageEmitterCallerSession struct {
	Contract *MessageEmitterCaller
	CallOpts bind.CallOpts
}

type MessageEmitterTransactorSession struct {
	Contract     *MessageEmitterTransactor
	TransactOpts bind.TransactOpts
}

type MessageEmitterRaw struct {
	Contract *MessageEmitter
}

type MessageEmitterCallerRaw struct {
	Contract *MessageEmitterCaller
}

type MessageEmitterTransactorRaw struct {
	Contract *MessageEmitterTransactor
}

func NewMessageEmitter(address common.Address, backend bind.ContractBackend) (*MessageEmitter, error) {
	abi, err := abi.JSON(strings.NewReader(MessageEmitterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageEmitter{address: address, abi: abi, MessageEmitterCaller: MessageEmitterCaller{contract: contract}, MessageEmitterTransactor: MessageEmitterTransactor{contract: contract}, MessageEmitterFilterer: MessageEmitterFilterer{contract: contract}}, nil
}

func NewMessageEmitterCaller(address common.Address, caller bind.ContractCaller) (*MessageEmitterCaller, error) {
	contract, err := bindMessageEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterCaller{contract: contract}, nil
}

func NewMessageEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageEmitterTransactor, error) {
	contract, err := bindMessageEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterTransactor{contract: contract}, nil
}

func NewMessageEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageEmitterFilterer, error) {
	contract, err := bindMessageEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterFilterer{contract: contract}, nil
}

func bindMessageEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MessageEmitter *MessageEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageEmitter.Contract.MessageEmitterCaller.contract.Call(opts, result, method, params...)
}

func (_MessageEmitter *MessageEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageEmitter.Contract.MessageEmitterTransactor.contract.Transfer(opts)
}

func (_MessageEmitter *MessageEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageEmitter.Contract.MessageEmitterTransactor.contract.Transact(opts, method, params...)
}

func (_MessageEmitter *MessageEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageEmitter.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageEmitter *MessageEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageEmitter.Contract.contract.Transfer(opts)
}

func (_MessageEmitter *MessageEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageEmitter.Contract.contract.Transact(opts, method, params...)
}

func (_MessageEmitter *MessageEmitterCaller) GetLastMessage(opts *bind.CallOpts, emitter common.Address) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "getLastMessage", emitter)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) GetLastMessage(emitter common.Address) (string, error) {
	return _MessageEmitter.Contract.GetLastMessage(&_MessageEmitter.CallOpts, emitter)
}

func (_MessageEmitter *MessageEmitterCallerSession) GetLastMessage(emitter common.Address) (string, error) {
	return _MessageEmitter.Contract.GetLastMessage(&_MessageEmitter.CallOpts, emitter)
}

func (_MessageEmitter *MessageEmitterCaller) GetMessage(opts *bind.CallOpts, emitter common.Address, timestamp *big.Int) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "getMessage", emitter, timestamp)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) GetMessage(emitter common.Address, timestamp *big.Int) (string, error) {
	return _MessageEmitter.Contract.GetMessage(&_MessageEmitter.CallOpts, emitter, timestamp)
}

func (_MessageEmitter *MessageEmitterCallerSession) GetMessage(emitter common.Address, timestamp *big.Int) (string, error) {
	return _MessageEmitter.Contract.GetMessage(&_MessageEmitter.CallOpts, emitter, timestamp)
}

func (_MessageEmitter *MessageEmitterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) TypeAndVersion() (string, error) {
	return _MessageEmitter.Contract.TypeAndVersion(&_MessageEmitter.CallOpts)
}

func (_MessageEmitter *MessageEmitterCallerSession) TypeAndVersion() (string, error) {
	return _MessageEmitter.Contract.TypeAndVersion(&_MessageEmitter.CallOpts)
}

func (_MessageEmitter *MessageEmitterTransactor) EmitMessage(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _MessageEmitter.contract.Transact(opts, "emitMessage", message)
}

func (_MessageEmitter *MessageEmitterSession) EmitMessage(message string) (*types.Transaction, error) {
	return _MessageEmitter.Contract.EmitMessage(&_MessageEmitter.TransactOpts, message)
}

func (_MessageEmitter *MessageEmitterTransactorSession) EmitMessage(message string) (*types.Transaction, error) {
	return _MessageEmitter.Contract.EmitMessage(&_MessageEmitter.TransactOpts, message)
}

type MessageEmitterMessageEmittedIterator struct {
	Event *MessageEmitterMessageEmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageEmitterMessageEmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageEmitterMessageEmitted)
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
		it.Event = new(MessageEmitterMessageEmitted)
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

func (it *MessageEmitterMessageEmittedIterator) Error() error {
	return it.fail
}

func (it *MessageEmitterMessageEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageEmitterMessageEmitted struct {
	Emitter   common.Address
	Timestamp *big.Int
	Message   string
	Raw       types.Log
}

func (_MessageEmitter *MessageEmitterFilterer) FilterMessageEmitted(opts *bind.FilterOpts, emitter []common.Address, timestamp []*big.Int) (*MessageEmitterMessageEmittedIterator, error) {

	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _MessageEmitter.contract.FilterLogs(opts, "MessageEmitted", emitterRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterMessageEmittedIterator{contract: _MessageEmitter.contract, event: "MessageEmitted", logs: logs, sub: sub}, nil
}

func (_MessageEmitter *MessageEmitterFilterer) WatchMessageEmitted(opts *bind.WatchOpts, sink chan<- *MessageEmitterMessageEmitted, emitter []common.Address, timestamp []*big.Int) (event.Subscription, error) {

	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _MessageEmitter.contract.WatchLogs(opts, "MessageEmitted", emitterRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageEmitterMessageEmitted)
				if err := _MessageEmitter.contract.UnpackLog(event, "MessageEmitted", log); err != nil {
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

func (_MessageEmitter *MessageEmitterFilterer) ParseMessageEmitted(log types.Log) (*MessageEmitterMessageEmitted, error) {
	event := new(MessageEmitterMessageEmitted)
	if err := _MessageEmitter.contract.UnpackLog(event, "MessageEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (MessageEmitterMessageEmitted) Topic() common.Hash {
	return common.HexToHash("0xc799f359194674b273986b8c03283265390f642b631c04e6526b99d0d8f4c38d")
}

func (_MessageEmitter *MessageEmitter) Address() common.Address {
	return _MessageEmitter.address
}

type MessageEmitterInterface interface {
	GetLastMessage(opts *bind.CallOpts, emitter common.Address) (string, error)

	GetMessage(opts *bind.CallOpts, emitter common.Address, timestamp *big.Int) (string, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	EmitMessage(opts *bind.TransactOpts, message string) (*types.Transaction, error)

	FilterMessageEmitted(opts *bind.FilterOpts, emitter []common.Address, timestamp []*big.Int) (*MessageEmitterMessageEmittedIterator, error)

	WatchMessageEmitted(opts *bind.WatchOpts, sink chan<- *MessageEmitterMessageEmitted, emitter []common.Address, timestamp []*big.Int) (event.Subscription, error)

	ParseMessageEmitted(log types.Log) (*MessageEmitterMessageEmitted, error)

	Address() common.Address
}
