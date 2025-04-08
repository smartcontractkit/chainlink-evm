// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package don_id_claimer

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

var DonIDClaimerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_capabilitiesRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAuthorizedDeployer\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAuthorizedDeployer\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncNextDONIdWithOffset\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorizedDeployerSet\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DonIDClaimed\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DonIDSynced\",\"inputs\":[{\"name\":\"newDONId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461014c57602081610d86803803809161001f8285610151565b83398101031261014c57516001600160a01b0381169081900361014c57331561013b57600180546001600160a01b03191633179055801561012a576080526100663361018a565b50608051604051637e6e477f60e11b815290602090829060049082906001600160a01b03165afa90811561011e576000916100d8575b506001805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055604051610b68908161021e8239608051816104d20152f35b6020813d602011610116575b816100f160209383610151565b8101031261011257519063ffffffff8216820361010f57503861009c565b80fd5b5080fd5b3d91506100e4565b6040513d6000823e3d90fd5b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761017457604052565b634e487b7160e01b600052604160045260246000fd5b806000526003602052604060002054156000146102175760025468010000000000000000811015610174576001810180600255811015610201577f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0181905560025460009182526003602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806309c43a58146107cd578063181f5a77146106bb578063381780d31461064c57806375c82b151461043357806379ba50971461034a57806385d60fa51461020b5780638da5cb5b146101b9578063f2fde38b146100c95763fcdc8efe1461007f57600080fd5b346100c45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c457602063ffffffff60015460a01c16604051908152f35b600080fd5b346100c45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c45773ffffffffffffffffffffffffffffffffffffffff6101156108b0565b61011d6108d3565b1633811461018f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100c45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100c45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c457610251336000526003602052604060002054151590565b1561031c5760015463ffffffff8160a01c166040518181527f5c1797088dea65fb15359febc2804939866c2a62e82aaecd8cf0032beb88996960203392a263ffffffff81146102ed576020917fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff00000000000000000000000000000000000000006001840160a01b16911617600155604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346100c45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c45760005473ffffffffffffffffffffffffffffffffffffffff81163303610409577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100c45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c45760043563ffffffff81168091036100c45761048a336000526003602052604060002054151590565b1561031c576040517ffcdc8efe00000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561064057600091610595575b5063ffffffff160163ffffffff81116102ed5760207f8bc6bcd971d85963c2ab42056c4c2ff253723ce5a6ba5358d54e62ce7cf70b1d917fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff00000000000000000000000000000000000000006001549260a01b169116178060015563ffffffff6040519160a01c168152a1005b60203d602011610639575b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f820116820182811067ffffffffffffffff82111761060c5760209183916040528101031261060857519063ffffffff82168203610605575063ffffffff610503565b80fd5b5080fd5b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503d6105a0565b6040513d6000823e3d90fd5b346100c45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c45760206106b173ffffffffffffffffffffffffffffffffffffffff61069d6108b0565b166000526003602052604060002054151590565b6040519015158152f35b346100c45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c4576040516040810181811067ffffffffffffffff82111761079e57604052601281527f446f6e4944436c61696d657220312e362e310000000000000000000000000000602082015260405190602082528181519182602083015260005b8381106107865750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610746565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346100c45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c4576108046108b0565b60243590811515908183036100c45773ffffffffffffffffffffffffffffffffffffffff906108316108d3565b16918215610886577f016cd377780c9e2bbe411bc110907cdac96e95bd4bdc86f007cb3668ded01f4291602091156108775761086c84610afb565b505b604051908152a2005b61088084610965565b5061086e565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100c457565b73ffffffffffffffffffffffffffffffffffffffff6001541633036108f457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80548210156109365760005260206000200190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000818152600360205260409020548015610af4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116102ed57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116102ed57818103610a85575b5050506002548015610a56577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610a1381600261091e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610adc610a96610aa793600261091e565b90549060031b1c928392600261091e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260036020526040600020553880806109da565b5050600090565b80600052600360205260406000205415600014610b55576002546801000000000000000081101561079e57610b3c610aa7826001859401600255600261091e565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var DonIDClaimerABI = DonIDClaimerMetaData.ABI

var DonIDClaimerBin = DonIDClaimerMetaData.Bin

func DeployDonIDClaimer(auth *bind.TransactOpts, backend bind.ContractBackend, _capabilitiesRegistry common.Address) (common.Address, *types.Transaction, *DonIDClaimer, error) {
	parsed, err := DonIDClaimerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DonIDClaimerBin), backend, _capabilitiesRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DonIDClaimer{address: address, abi: *parsed, DonIDClaimerCaller: DonIDClaimerCaller{contract: contract}, DonIDClaimerTransactor: DonIDClaimerTransactor{contract: contract}, DonIDClaimerFilterer: DonIDClaimerFilterer{contract: contract}}, nil
}

type DonIDClaimer struct {
	address common.Address
	abi     abi.ABI
	DonIDClaimerCaller
	DonIDClaimerTransactor
	DonIDClaimerFilterer
}

type DonIDClaimerCaller struct {
	contract *bind.BoundContract
}

type DonIDClaimerTransactor struct {
	contract *bind.BoundContract
}

type DonIDClaimerFilterer struct {
	contract *bind.BoundContract
}

type DonIDClaimerSession struct {
	Contract     *DonIDClaimer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DonIDClaimerCallerSession struct {
	Contract *DonIDClaimerCaller
	CallOpts bind.CallOpts
}

type DonIDClaimerTransactorSession struct {
	Contract     *DonIDClaimerTransactor
	TransactOpts bind.TransactOpts
}

type DonIDClaimerRaw struct {
	Contract *DonIDClaimer
}

type DonIDClaimerCallerRaw struct {
	Contract *DonIDClaimerCaller
}

type DonIDClaimerTransactorRaw struct {
	Contract *DonIDClaimerTransactor
}

func NewDonIDClaimer(address common.Address, backend bind.ContractBackend) (*DonIDClaimer, error) {
	abi, err := abi.JSON(strings.NewReader(DonIDClaimerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindDonIDClaimer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimer{address: address, abi: abi, DonIDClaimerCaller: DonIDClaimerCaller{contract: contract}, DonIDClaimerTransactor: DonIDClaimerTransactor{contract: contract}, DonIDClaimerFilterer: DonIDClaimerFilterer{contract: contract}}, nil
}

func NewDonIDClaimerCaller(address common.Address, caller bind.ContractCaller) (*DonIDClaimerCaller, error) {
	contract, err := bindDonIDClaimer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerCaller{contract: contract}, nil
}

func NewDonIDClaimerTransactor(address common.Address, transactor bind.ContractTransactor) (*DonIDClaimerTransactor, error) {
	contract, err := bindDonIDClaimer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerTransactor{contract: contract}, nil
}

func NewDonIDClaimerFilterer(address common.Address, filterer bind.ContractFilterer) (*DonIDClaimerFilterer, error) {
	contract, err := bindDonIDClaimer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerFilterer{contract: contract}, nil
}

func bindDonIDClaimer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DonIDClaimerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_DonIDClaimer *DonIDClaimerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonIDClaimer.Contract.DonIDClaimerCaller.contract.Call(opts, result, method, params...)
}

func (_DonIDClaimer *DonIDClaimerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.DonIDClaimerTransactor.contract.Transfer(opts)
}

func (_DonIDClaimer *DonIDClaimerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.DonIDClaimerTransactor.contract.Transact(opts, method, params...)
}

func (_DonIDClaimer *DonIDClaimerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonIDClaimer.Contract.contract.Call(opts, result, method, params...)
}

func (_DonIDClaimer *DonIDClaimerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.contract.Transfer(opts)
}

func (_DonIDClaimer *DonIDClaimerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.contract.Transact(opts, method, params...)
}

func (_DonIDClaimer *DonIDClaimerCaller) GetNextDONId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DonIDClaimer.contract.Call(opts, &out, "getNextDONId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_DonIDClaimer *DonIDClaimerSession) GetNextDONId() (uint32, error) {
	return _DonIDClaimer.Contract.GetNextDONId(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerCallerSession) GetNextDONId() (uint32, error) {
	return _DonIDClaimer.Contract.GetNextDONId(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerCaller) IsAuthorizedDeployer(opts *bind.CallOpts, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _DonIDClaimer.contract.Call(opts, &out, "isAuthorizedDeployer", senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_DonIDClaimer *DonIDClaimerSession) IsAuthorizedDeployer(senderAddress common.Address) (bool, error) {
	return _DonIDClaimer.Contract.IsAuthorizedDeployer(&_DonIDClaimer.CallOpts, senderAddress)
}

func (_DonIDClaimer *DonIDClaimerCallerSession) IsAuthorizedDeployer(senderAddress common.Address) (bool, error) {
	return _DonIDClaimer.Contract.IsAuthorizedDeployer(&_DonIDClaimer.CallOpts, senderAddress)
}

func (_DonIDClaimer *DonIDClaimerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DonIDClaimer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DonIDClaimer *DonIDClaimerSession) Owner() (common.Address, error) {
	return _DonIDClaimer.Contract.Owner(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerCallerSession) Owner() (common.Address, error) {
	return _DonIDClaimer.Contract.Owner(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DonIDClaimer.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_DonIDClaimer *DonIDClaimerSession) TypeAndVersion() (string, error) {
	return _DonIDClaimer.Contract.TypeAndVersion(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerCallerSession) TypeAndVersion() (string, error) {
	return _DonIDClaimer.Contract.TypeAndVersion(&_DonIDClaimer.CallOpts)
}

func (_DonIDClaimer *DonIDClaimerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonIDClaimer.contract.Transact(opts, "acceptOwnership")
}

func (_DonIDClaimer *DonIDClaimerSession) AcceptOwnership() (*types.Transaction, error) {
	return _DonIDClaimer.Contract.AcceptOwnership(&_DonIDClaimer.TransactOpts)
}

func (_DonIDClaimer *DonIDClaimerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DonIDClaimer.Contract.AcceptOwnership(&_DonIDClaimer.TransactOpts)
}

func (_DonIDClaimer *DonIDClaimerTransactor) ClaimNextDONId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonIDClaimer.contract.Transact(opts, "claimNextDONId")
}

func (_DonIDClaimer *DonIDClaimerSession) ClaimNextDONId() (*types.Transaction, error) {
	return _DonIDClaimer.Contract.ClaimNextDONId(&_DonIDClaimer.TransactOpts)
}

func (_DonIDClaimer *DonIDClaimerTransactorSession) ClaimNextDONId() (*types.Transaction, error) {
	return _DonIDClaimer.Contract.ClaimNextDONId(&_DonIDClaimer.TransactOpts)
}

func (_DonIDClaimer *DonIDClaimerTransactor) SetAuthorizedDeployer(opts *bind.TransactOpts, senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _DonIDClaimer.contract.Transact(opts, "setAuthorizedDeployer", senderAddress, allowed)
}

func (_DonIDClaimer *DonIDClaimerSession) SetAuthorizedDeployer(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.SetAuthorizedDeployer(&_DonIDClaimer.TransactOpts, senderAddress, allowed)
}

func (_DonIDClaimer *DonIDClaimerTransactorSession) SetAuthorizedDeployer(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.SetAuthorizedDeployer(&_DonIDClaimer.TransactOpts, senderAddress, allowed)
}

func (_DonIDClaimer *DonIDClaimerTransactor) SyncNextDONIdWithOffset(opts *bind.TransactOpts, offset uint32) (*types.Transaction, error) {
	return _DonIDClaimer.contract.Transact(opts, "syncNextDONIdWithOffset", offset)
}

func (_DonIDClaimer *DonIDClaimerSession) SyncNextDONIdWithOffset(offset uint32) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.SyncNextDONIdWithOffset(&_DonIDClaimer.TransactOpts, offset)
}

func (_DonIDClaimer *DonIDClaimerTransactorSession) SyncNextDONIdWithOffset(offset uint32) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.SyncNextDONIdWithOffset(&_DonIDClaimer.TransactOpts, offset)
}

func (_DonIDClaimer *DonIDClaimerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DonIDClaimer.contract.Transact(opts, "transferOwnership", to)
}

func (_DonIDClaimer *DonIDClaimerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.TransferOwnership(&_DonIDClaimer.TransactOpts, to)
}

func (_DonIDClaimer *DonIDClaimerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DonIDClaimer.Contract.TransferOwnership(&_DonIDClaimer.TransactOpts, to)
}

type DonIDClaimerAuthorizedDeployerSetIterator struct {
	Event *DonIDClaimerAuthorizedDeployerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DonIDClaimerAuthorizedDeployerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonIDClaimerAuthorizedDeployerSet)
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
		it.Event = new(DonIDClaimerAuthorizedDeployerSet)
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

func (it *DonIDClaimerAuthorizedDeployerSetIterator) Error() error {
	return it.fail
}

func (it *DonIDClaimerAuthorizedDeployerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DonIDClaimerAuthorizedDeployerSet struct {
	SenderAddress common.Address
	Allowed       bool
	Raw           types.Log
}

func (_DonIDClaimer *DonIDClaimerFilterer) FilterAuthorizedDeployerSet(opts *bind.FilterOpts, senderAddress []common.Address) (*DonIDClaimerAuthorizedDeployerSetIterator, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _DonIDClaimer.contract.FilterLogs(opts, "AuthorizedDeployerSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerAuthorizedDeployerSetIterator{contract: _DonIDClaimer.contract, event: "AuthorizedDeployerSet", logs: logs, sub: sub}, nil
}

func (_DonIDClaimer *DonIDClaimerFilterer) WatchAuthorizedDeployerSet(opts *bind.WatchOpts, sink chan<- *DonIDClaimerAuthorizedDeployerSet, senderAddress []common.Address) (event.Subscription, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _DonIDClaimer.contract.WatchLogs(opts, "AuthorizedDeployerSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DonIDClaimerAuthorizedDeployerSet)
				if err := _DonIDClaimer.contract.UnpackLog(event, "AuthorizedDeployerSet", log); err != nil {
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

func (_DonIDClaimer *DonIDClaimerFilterer) ParseAuthorizedDeployerSet(log types.Log) (*DonIDClaimerAuthorizedDeployerSet, error) {
	event := new(DonIDClaimerAuthorizedDeployerSet)
	if err := _DonIDClaimer.contract.UnpackLog(event, "AuthorizedDeployerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DonIDClaimerDonIDClaimedIterator struct {
	Event *DonIDClaimerDonIDClaimed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DonIDClaimerDonIDClaimedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonIDClaimerDonIDClaimed)
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
		it.Event = new(DonIDClaimerDonIDClaimed)
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

func (it *DonIDClaimerDonIDClaimedIterator) Error() error {
	return it.fail
}

func (it *DonIDClaimerDonIDClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DonIDClaimerDonIDClaimed struct {
	Claimer common.Address
	DonId   uint32
	Raw     types.Log
}

func (_DonIDClaimer *DonIDClaimerFilterer) FilterDonIDClaimed(opts *bind.FilterOpts, claimer []common.Address) (*DonIDClaimerDonIDClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DonIDClaimer.contract.FilterLogs(opts, "DonIDClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerDonIDClaimedIterator{contract: _DonIDClaimer.contract, event: "DonIDClaimed", logs: logs, sub: sub}, nil
}

func (_DonIDClaimer *DonIDClaimerFilterer) WatchDonIDClaimed(opts *bind.WatchOpts, sink chan<- *DonIDClaimerDonIDClaimed, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DonIDClaimer.contract.WatchLogs(opts, "DonIDClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DonIDClaimerDonIDClaimed)
				if err := _DonIDClaimer.contract.UnpackLog(event, "DonIDClaimed", log); err != nil {
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

func (_DonIDClaimer *DonIDClaimerFilterer) ParseDonIDClaimed(log types.Log) (*DonIDClaimerDonIDClaimed, error) {
	event := new(DonIDClaimerDonIDClaimed)
	if err := _DonIDClaimer.contract.UnpackLog(event, "DonIDClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DonIDClaimerDonIDSyncedIterator struct {
	Event *DonIDClaimerDonIDSynced

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DonIDClaimerDonIDSyncedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonIDClaimerDonIDSynced)
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
		it.Event = new(DonIDClaimerDonIDSynced)
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

func (it *DonIDClaimerDonIDSyncedIterator) Error() error {
	return it.fail
}

func (it *DonIDClaimerDonIDSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DonIDClaimerDonIDSynced struct {
	NewDONId uint32
	Raw      types.Log
}

func (_DonIDClaimer *DonIDClaimerFilterer) FilterDonIDSynced(opts *bind.FilterOpts) (*DonIDClaimerDonIDSyncedIterator, error) {

	logs, sub, err := _DonIDClaimer.contract.FilterLogs(opts, "DonIDSynced")
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerDonIDSyncedIterator{contract: _DonIDClaimer.contract, event: "DonIDSynced", logs: logs, sub: sub}, nil
}

func (_DonIDClaimer *DonIDClaimerFilterer) WatchDonIDSynced(opts *bind.WatchOpts, sink chan<- *DonIDClaimerDonIDSynced) (event.Subscription, error) {

	logs, sub, err := _DonIDClaimer.contract.WatchLogs(opts, "DonIDSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DonIDClaimerDonIDSynced)
				if err := _DonIDClaimer.contract.UnpackLog(event, "DonIDSynced", log); err != nil {
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

func (_DonIDClaimer *DonIDClaimerFilterer) ParseDonIDSynced(log types.Log) (*DonIDClaimerDonIDSynced, error) {
	event := new(DonIDClaimerDonIDSynced)
	if err := _DonIDClaimer.contract.UnpackLog(event, "DonIDSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DonIDClaimerOwnershipTransferRequestedIterator struct {
	Event *DonIDClaimerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DonIDClaimerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonIDClaimerOwnershipTransferRequested)
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
		it.Event = new(DonIDClaimerOwnershipTransferRequested)
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

func (it *DonIDClaimerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *DonIDClaimerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DonIDClaimerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DonIDClaimer *DonIDClaimerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DonIDClaimerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DonIDClaimer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerOwnershipTransferRequestedIterator{contract: _DonIDClaimer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_DonIDClaimer *DonIDClaimerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DonIDClaimerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DonIDClaimer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DonIDClaimerOwnershipTransferRequested)
				if err := _DonIDClaimer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_DonIDClaimer *DonIDClaimerFilterer) ParseOwnershipTransferRequested(log types.Log) (*DonIDClaimerOwnershipTransferRequested, error) {
	event := new(DonIDClaimerOwnershipTransferRequested)
	if err := _DonIDClaimer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DonIDClaimerOwnershipTransferredIterator struct {
	Event *DonIDClaimerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DonIDClaimerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DonIDClaimerOwnershipTransferred)
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
		it.Event = new(DonIDClaimerOwnershipTransferred)
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

func (it *DonIDClaimerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *DonIDClaimerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DonIDClaimerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DonIDClaimer *DonIDClaimerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DonIDClaimerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DonIDClaimer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DonIDClaimerOwnershipTransferredIterator{contract: _DonIDClaimer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_DonIDClaimer *DonIDClaimerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DonIDClaimerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DonIDClaimer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DonIDClaimerOwnershipTransferred)
				if err := _DonIDClaimer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_DonIDClaimer *DonIDClaimerFilterer) ParseOwnershipTransferred(log types.Log) (*DonIDClaimerOwnershipTransferred, error) {
	event := new(DonIDClaimerOwnershipTransferred)
	if err := _DonIDClaimer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_DonIDClaimer *DonIDClaimer) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _DonIDClaimer.abi.Events["AuthorizedDeployerSet"].ID:
		return _DonIDClaimer.ParseAuthorizedDeployerSet(log)
	case _DonIDClaimer.abi.Events["DonIDClaimed"].ID:
		return _DonIDClaimer.ParseDonIDClaimed(log)
	case _DonIDClaimer.abi.Events["DonIDSynced"].ID:
		return _DonIDClaimer.ParseDonIDSynced(log)
	case _DonIDClaimer.abi.Events["OwnershipTransferRequested"].ID:
		return _DonIDClaimer.ParseOwnershipTransferRequested(log)
	case _DonIDClaimer.abi.Events["OwnershipTransferred"].ID:
		return _DonIDClaimer.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (DonIDClaimerAuthorizedDeployerSet) Topic() common.Hash {
	return common.HexToHash("0x016cd377780c9e2bbe411bc110907cdac96e95bd4bdc86f007cb3668ded01f42")
}

func (DonIDClaimerDonIDClaimed) Topic() common.Hash {
	return common.HexToHash("0x5c1797088dea65fb15359febc2804939866c2a62e82aaecd8cf0032beb889969")
}

func (DonIDClaimerDonIDSynced) Topic() common.Hash {
	return common.HexToHash("0x8bc6bcd971d85963c2ab42056c4c2ff253723ce5a6ba5358d54e62ce7cf70b1d")
}

func (DonIDClaimerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (DonIDClaimerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_DonIDClaimer *DonIDClaimer) Address() common.Address {
	return _DonIDClaimer.address
}

type DonIDClaimerInterface interface {
	GetNextDONId(opts *bind.CallOpts) (uint32, error)

	IsAuthorizedDeployer(opts *bind.CallOpts, senderAddress common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ClaimNextDONId(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAuthorizedDeployer(opts *bind.TransactOpts, senderAddress common.Address, allowed bool) (*types.Transaction, error)

	SyncNextDONIdWithOffset(opts *bind.TransactOpts, offset uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedDeployerSet(opts *bind.FilterOpts, senderAddress []common.Address) (*DonIDClaimerAuthorizedDeployerSetIterator, error)

	WatchAuthorizedDeployerSet(opts *bind.WatchOpts, sink chan<- *DonIDClaimerAuthorizedDeployerSet, senderAddress []common.Address) (event.Subscription, error)

	ParseAuthorizedDeployerSet(log types.Log) (*DonIDClaimerAuthorizedDeployerSet, error)

	FilterDonIDClaimed(opts *bind.FilterOpts, claimer []common.Address) (*DonIDClaimerDonIDClaimedIterator, error)

	WatchDonIDClaimed(opts *bind.WatchOpts, sink chan<- *DonIDClaimerDonIDClaimed, claimer []common.Address) (event.Subscription, error)

	ParseDonIDClaimed(log types.Log) (*DonIDClaimerDonIDClaimed, error)

	FilterDonIDSynced(opts *bind.FilterOpts) (*DonIDClaimerDonIDSyncedIterator, error)

	WatchDonIDSynced(opts *bind.WatchOpts, sink chan<- *DonIDClaimerDonIDSynced) (event.Subscription, error)

	ParseDonIDSynced(log types.Log) (*DonIDClaimerDonIDSynced, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DonIDClaimerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DonIDClaimerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*DonIDClaimerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DonIDClaimerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DonIDClaimerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*DonIDClaimerOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
