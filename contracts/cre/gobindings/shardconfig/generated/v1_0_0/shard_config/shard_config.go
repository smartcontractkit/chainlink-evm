// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shard_config

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

var ShardConfigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_desiredShardCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mcmsAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"desiredShardCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDesiredShardCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mcmsAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDesiredShardCount\",\"inputs\":[{\"name\":\"_newCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMCMSAddress\",\"inputs\":[{\"name\":\"_newMcmsAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MCMSAddressUpdated\",\"inputs\":[{\"name\":\"newMcmsAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShardCountUpdated\",\"inputs\":[{\"name\":\"newCount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106a93803806106a983398101604081905261002f91610168565b6001600160a01b03811661008a5760405162461bcd60e51b815260206004820152601460248201527f496e76616c6964204d434d53206164647265737300000000000000000000000060448201526064015b60405180910390fd5b600082116100e55760405162461bcd60e51b815260206004820152602260248201527f536861726420636f756e74206d7573742062652067726561746572207468616e604482015261020360f41b6064820152608401610081565b6000828155600180546001600160a01b0319166001600160a01b03841617905560405183917f14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b877991a26040516001600160a01b038216907f5f2b0e3c4978c02818ac390c030d739e08f2c0d30f28c6fac9a793feb9ed497d90600090a250506101a5565b6000806040838503121561017b57600080fd5b825160208401519092506001600160a01b038116811461019a57600080fd5b809150509250929050565b6104f5806101b46000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80636bc66efb116100505780636bc66efb146100ec57806384deb990146100f5578063ae8f25961461010857600080fd5b80630164a01014610077578063181f5a771461008e5780632e3cd5aa146100d7575b600080fd5b6000545b6040519081526020015b60405180910390f35b6100ca6040518060400160405280601181526020017f5368617264436f6e66696720312e302e3000000000000000000000000000000081525081565b6040516100859190610425565b6100ea6100e5366004610492565b61014d565b005b61007b60005481565b6100ea6101033660046104ab565b610293565b6001546101289073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610085565b60015473ffffffffffffffffffffffffffffffffffffffff1633146101d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f6e6c79204d434d532063616e2075706461746520736861726420636f756e7460448201526064015b60405180910390fd5b60008111610263576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f536861726420636f756e74206d7573742062652067726561746572207468616e60448201527f203000000000000000000000000000000000000000000000000000000000000060648201526084016101ca565b600081815560405182917f14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b877991a250565b60015473ffffffffffffffffffffffffffffffffffffffff163314610339576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4f6e6c79204d434d532063616e2075706461746520697473206f776e2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016101ca565b73ffffffffffffffffffffffffffffffffffffffff81166103b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e76616c6964204d434d53206164647265737300000000000000000000000060448201526064016101ca565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f5f2b0e3c4978c02818ac390c030d739e08f2c0d30f28c6fac9a793feb9ed497d90600090a250565b60006020808352835180602085015260005b8181101561045357858101830151858201604001528201610437565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000602082840312156104a457600080fd5b5035919050565b6000602082840312156104bd57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146104e157600080fd5b939250505056fea164736f6c6343000818000a",
}

var ShardConfigABI = ShardConfigMetaData.ABI

var ShardConfigBin = ShardConfigMetaData.Bin

func DeployShardConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _desiredShardCount *big.Int, _mcmsAddress common.Address) (common.Address, *types.Transaction, *ShardConfig, error) {
	parsed, err := ShardConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShardConfigBin), backend, _desiredShardCount, _mcmsAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShardConfig{address: address, abi: *parsed, ShardConfigCaller: ShardConfigCaller{contract: contract}, ShardConfigTransactor: ShardConfigTransactor{contract: contract}, ShardConfigFilterer: ShardConfigFilterer{contract: contract}}, nil
}

type ShardConfig struct {
	address common.Address
	abi     abi.ABI
	ShardConfigCaller
	ShardConfigTransactor
	ShardConfigFilterer
}

type ShardConfigCaller struct {
	contract *bind.BoundContract
}

type ShardConfigTransactor struct {
	contract *bind.BoundContract
}

type ShardConfigFilterer struct {
	contract *bind.BoundContract
}

type ShardConfigSession struct {
	Contract     *ShardConfig
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ShardConfigCallerSession struct {
	Contract *ShardConfigCaller
	CallOpts bind.CallOpts
}

type ShardConfigTransactorSession struct {
	Contract     *ShardConfigTransactor
	TransactOpts bind.TransactOpts
}

type ShardConfigRaw struct {
	Contract *ShardConfig
}

type ShardConfigCallerRaw struct {
	Contract *ShardConfigCaller
}

type ShardConfigTransactorRaw struct {
	Contract *ShardConfigTransactor
}

func NewShardConfig(address common.Address, backend bind.ContractBackend) (*ShardConfig, error) {
	abi, err := abi.JSON(strings.NewReader(ShardConfigABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindShardConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShardConfig{address: address, abi: abi, ShardConfigCaller: ShardConfigCaller{contract: contract}, ShardConfigTransactor: ShardConfigTransactor{contract: contract}, ShardConfigFilterer: ShardConfigFilterer{contract: contract}}, nil
}

func NewShardConfigCaller(address common.Address, caller bind.ContractCaller) (*ShardConfigCaller, error) {
	contract, err := bindShardConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShardConfigCaller{contract: contract}, nil
}

func NewShardConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ShardConfigTransactor, error) {
	contract, err := bindShardConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShardConfigTransactor{contract: contract}, nil
}

func NewShardConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ShardConfigFilterer, error) {
	contract, err := bindShardConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShardConfigFilterer{contract: contract}, nil
}

func bindShardConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShardConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ShardConfig *ShardConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardConfig.Contract.ShardConfigCaller.contract.Call(opts, result, method, params...)
}

func (_ShardConfig *ShardConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardConfig.Contract.ShardConfigTransactor.contract.Transfer(opts)
}

func (_ShardConfig *ShardConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardConfig.Contract.ShardConfigTransactor.contract.Transact(opts, method, params...)
}

func (_ShardConfig *ShardConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardConfig.Contract.contract.Call(opts, result, method, params...)
}

func (_ShardConfig *ShardConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardConfig.Contract.contract.Transfer(opts)
}

func (_ShardConfig *ShardConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardConfig.Contract.contract.Transact(opts, method, params...)
}

func (_ShardConfig *ShardConfigCaller) DesiredShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "desiredShardCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) DesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.DesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) DesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.DesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) GetDesiredShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "getDesiredShardCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) GetDesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.GetDesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) GetDesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.GetDesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) McmsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "mcmsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) McmsAddress() (common.Address, error) {
	return _ShardConfig.Contract.McmsAddress(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) McmsAddress() (common.Address, error) {
	return _ShardConfig.Contract.McmsAddress(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) TypeAndVersion() (string, error) {
	return _ShardConfig.Contract.TypeAndVersion(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) TypeAndVersion() (string, error) {
	return _ShardConfig.Contract.TypeAndVersion(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigTransactor) SetDesiredShardCount(opts *bind.TransactOpts, _newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.contract.Transact(opts, "setDesiredShardCount", _newCount)
}

func (_ShardConfig *ShardConfigSession) SetDesiredShardCount(_newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetDesiredShardCount(&_ShardConfig.TransactOpts, _newCount)
}

func (_ShardConfig *ShardConfigTransactorSession) SetDesiredShardCount(_newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetDesiredShardCount(&_ShardConfig.TransactOpts, _newCount)
}

func (_ShardConfig *ShardConfigTransactor) SetMCMSAddress(opts *bind.TransactOpts, _newMcmsAddress common.Address) (*types.Transaction, error) {
	return _ShardConfig.contract.Transact(opts, "setMCMSAddress", _newMcmsAddress)
}

func (_ShardConfig *ShardConfigSession) SetMCMSAddress(_newMcmsAddress common.Address) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetMCMSAddress(&_ShardConfig.TransactOpts, _newMcmsAddress)
}

func (_ShardConfig *ShardConfigTransactorSession) SetMCMSAddress(_newMcmsAddress common.Address) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetMCMSAddress(&_ShardConfig.TransactOpts, _newMcmsAddress)
}

type ShardConfigMCMSAddressUpdatedIterator struct {
	Event *ShardConfigMCMSAddressUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ShardConfigMCMSAddressUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardConfigMCMSAddressUpdated)
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
		it.Event = new(ShardConfigMCMSAddressUpdated)
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

func (it *ShardConfigMCMSAddressUpdatedIterator) Error() error {
	return it.fail
}

func (it *ShardConfigMCMSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ShardConfigMCMSAddressUpdated struct {
	NewMcmsAddress common.Address
	Raw            types.Log
}

func (_ShardConfig *ShardConfigFilterer) FilterMCMSAddressUpdated(opts *bind.FilterOpts, newMcmsAddress []common.Address) (*ShardConfigMCMSAddressUpdatedIterator, error) {

	var newMcmsAddressRule []interface{}
	for _, newMcmsAddressItem := range newMcmsAddress {
		newMcmsAddressRule = append(newMcmsAddressRule, newMcmsAddressItem)
	}

	logs, sub, err := _ShardConfig.contract.FilterLogs(opts, "MCMSAddressUpdated", newMcmsAddressRule)
	if err != nil {
		return nil, err
	}
	return &ShardConfigMCMSAddressUpdatedIterator{contract: _ShardConfig.contract, event: "MCMSAddressUpdated", logs: logs, sub: sub}, nil
}

func (_ShardConfig *ShardConfigFilterer) WatchMCMSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigMCMSAddressUpdated, newMcmsAddress []common.Address) (event.Subscription, error) {

	var newMcmsAddressRule []interface{}
	for _, newMcmsAddressItem := range newMcmsAddress {
		newMcmsAddressRule = append(newMcmsAddressRule, newMcmsAddressItem)
	}

	logs, sub, err := _ShardConfig.contract.WatchLogs(opts, "MCMSAddressUpdated", newMcmsAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ShardConfigMCMSAddressUpdated)
				if err := _ShardConfig.contract.UnpackLog(event, "MCMSAddressUpdated", log); err != nil {
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

func (_ShardConfig *ShardConfigFilterer) ParseMCMSAddressUpdated(log types.Log) (*ShardConfigMCMSAddressUpdated, error) {
	event := new(ShardConfigMCMSAddressUpdated)
	if err := _ShardConfig.contract.UnpackLog(event, "MCMSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ShardConfigShardCountUpdatedIterator struct {
	Event *ShardConfigShardCountUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ShardConfigShardCountUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardConfigShardCountUpdated)
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
		it.Event = new(ShardConfigShardCountUpdated)
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

func (it *ShardConfigShardCountUpdatedIterator) Error() error {
	return it.fail
}

func (it *ShardConfigShardCountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ShardConfigShardCountUpdated struct {
	NewCount *big.Int
	Raw      types.Log
}

func (_ShardConfig *ShardConfigFilterer) FilterShardCountUpdated(opts *bind.FilterOpts, newCount []*big.Int) (*ShardConfigShardCountUpdatedIterator, error) {

	var newCountRule []interface{}
	for _, newCountItem := range newCount {
		newCountRule = append(newCountRule, newCountItem)
	}

	logs, sub, err := _ShardConfig.contract.FilterLogs(opts, "ShardCountUpdated", newCountRule)
	if err != nil {
		return nil, err
	}
	return &ShardConfigShardCountUpdatedIterator{contract: _ShardConfig.contract, event: "ShardCountUpdated", logs: logs, sub: sub}, nil
}

func (_ShardConfig *ShardConfigFilterer) WatchShardCountUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigShardCountUpdated, newCount []*big.Int) (event.Subscription, error) {

	var newCountRule []interface{}
	for _, newCountItem := range newCount {
		newCountRule = append(newCountRule, newCountItem)
	}

	logs, sub, err := _ShardConfig.contract.WatchLogs(opts, "ShardCountUpdated", newCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ShardConfigShardCountUpdated)
				if err := _ShardConfig.contract.UnpackLog(event, "ShardCountUpdated", log); err != nil {
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

func (_ShardConfig *ShardConfigFilterer) ParseShardCountUpdated(log types.Log) (*ShardConfigShardCountUpdated, error) {
	event := new(ShardConfigShardCountUpdated)
	if err := _ShardConfig.contract.UnpackLog(event, "ShardCountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ShardConfigMCMSAddressUpdated) Topic() common.Hash {
	return common.HexToHash("0x5f2b0e3c4978c02818ac390c030d739e08f2c0d30f28c6fac9a793feb9ed497d")
}

func (ShardConfigShardCountUpdated) Topic() common.Hash {
	return common.HexToHash("0x14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b8779")
}

func (_ShardConfig *ShardConfig) Address() common.Address {
	return _ShardConfig.address
}

type ShardConfigInterface interface {
	DesiredShardCount(opts *bind.CallOpts) (*big.Int, error)

	GetDesiredShardCount(opts *bind.CallOpts) (*big.Int, error)

	McmsAddress(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SetDesiredShardCount(opts *bind.TransactOpts, _newCount *big.Int) (*types.Transaction, error)

	SetMCMSAddress(opts *bind.TransactOpts, _newMcmsAddress common.Address) (*types.Transaction, error)

	FilterMCMSAddressUpdated(opts *bind.FilterOpts, newMcmsAddress []common.Address) (*ShardConfigMCMSAddressUpdatedIterator, error)

	WatchMCMSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigMCMSAddressUpdated, newMcmsAddress []common.Address) (event.Subscription, error)

	ParseMCMSAddressUpdated(log types.Log) (*ShardConfigMCMSAddressUpdated, error)

	FilterShardCountUpdated(opts *bind.FilterOpts, newCount []*big.Int) (*ShardConfigShardCountUpdatedIterator, error)

	WatchShardCountUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigShardCountUpdated, newCount []*big.Int) (event.Subscription, error)

	ParseShardCountUpdated(log types.Log) (*ShardConfigShardCountUpdated, error)

	Address() common.Address
}
