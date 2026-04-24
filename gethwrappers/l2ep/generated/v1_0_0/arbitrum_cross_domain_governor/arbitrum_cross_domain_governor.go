// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_cross_domain_governor

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

var ArbitrumCrossDomainGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardDelegate\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161110b38038061110b83398101604081905261002e916101d2565b808033805f816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b03848116919091179091558116156100b4576100b4816100ce565b5050506100c68161017660201b60201c565b5050506101ff565b336001600160a01b038216036101265760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f602082840312156101e2575f80fd5b81516001600160a01b03811681146101f8575f80fd5b9392505050565b610eff8061020c5f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610164578063f2fde38b14610182578063f43b361314610195575f80fd5b806396b8d7c014610149578063b2ec07871461015c575f80fd5b80636fadcf72116100a25780636fadcf72146100f057806379ba5097146101035780638da5cb5b1461010b575f80fd5b8063181f5a77146100bd57806326929eb6146100db575b5f80fd5b6100c561019d565b6040516100d29190610d1e565b60405180910390f35b6100ee6100e9366004610dc3565b6101bd565b005b6100ee6100fe366004610dc3565b6102ea565b6100ee61040d565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d2565b6100ee610157366004610e9d565b610509565b6100ee6105b1565b60025473ffffffffffffffffffffffffffffffffffffffff16610124565b6100ee610190366004610e9d565b61069f565b6101246106b0565b6060604051806060016040528060218152602001610ed260219139905090565b6101c56106b0565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061021457505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6102a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102e582826040518060400160405280601e81526020017f476f7665726e6f722064656c656761746563616c6c20726576657274656400008152506106f1565b505050565b6102f26106b0565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061034157505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6103cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e657200000000000000000000000000000000000000000000000000606482015260840161029c565b6102e582826040518060400160405280601681526020017f476f7665726e6f722063616c6c20726576657274656400000000000000000000815250610817565b60015473ffffffffffffffffffffffffffffffffffffffff16331461048e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161029c565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105116106b0565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e6765720000604482015260640161029c565b6105ae8161082d565b50565b6003546105e79073ffffffffffffffffffffffffffffffffffffffff167311110000000000000000000000000000000011110190565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461067b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e657200000000000000604482015260640161029c565b60035461069d9073ffffffffffffffffffffffffffffffffffffffff16610922565b565b6106a76109a3565b6105ae81610a23565b5f6106ec6106d360025473ffffffffffffffffffffffffffffffffffffffff1690565b7311110000000000000000000000000000000011110190565b905090565b606073ffffffffffffffffffffffffffffffffffffffff84163b610797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e74726163740000000000000000000000000000000000000000000000000000606482015260840161029c565b5f808573ffffffffffffffffffffffffffffffffffffffff16856040516107be9190610eb6565b5f60405180830381855af49150503d805f81146107f6576040519150601f19603f3d011682016040523d82523d5f602084013e6107fb565b606091505b509150915061080b828286610b17565b925050505b9392505050565b606061082584845f85610b6a565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036108ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029c565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461069d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161029c565b3373ffffffffffffffffffffffffffffffffffffffff821603610aa2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60608315610b26575081610810565b825115610b365782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029c9190610d1e565b606082471015610bfc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161029c565b73ffffffffffffffffffffffffffffffffffffffff85163b610c7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161029c565b5f808673ffffffffffffffffffffffffffffffffffffffff168587604051610ca29190610eb6565b5f6040518083038185875af1925050503d805f8114610cdc576040519150601f19603f3d011682016040523d82523d5f602084013e610ce1565b606091505b5091509150610cf1828286610b17565b979650505050505050565b5f5b83811015610d16578181015183820152602001610cfe565b50505f910152565b602081525f8251806020840152610d3c816040850160208701610cfc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d91575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610dd4575f80fd5b610ddd83610d6e565b9150602083013567ffffffffffffffff80821115610df9575f80fd5b818501915085601f830112610e0c575f80fd5b813581811115610e1e57610e1e610d96565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610e6457610e64610d96565b81604052828152886020848701011115610e7c575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610ead575f80fd5b61081082610d6e565b5f8251610ec7818460208701610cfc565b919091019291505056fe417262697472756d43726f7373446f6d61696e476f7665726e6f7220312e302e30a164736f6c6343000818000a",
}

var ArbitrumCrossDomainGovernorABI = ArbitrumCrossDomainGovernorMetaData.ABI

var ArbitrumCrossDomainGovernorBin = ArbitrumCrossDomainGovernorMetaData.Bin

func DeployArbitrumCrossDomainGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *ArbitrumCrossDomainGovernor, error) {
	parsed, err := ArbitrumCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumCrossDomainGovernorBin), backend, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumCrossDomainGovernor{address: address, abi: *parsed, ArbitrumCrossDomainGovernorCaller: ArbitrumCrossDomainGovernorCaller{contract: contract}, ArbitrumCrossDomainGovernorTransactor: ArbitrumCrossDomainGovernorTransactor{contract: contract}, ArbitrumCrossDomainGovernorFilterer: ArbitrumCrossDomainGovernorFilterer{contract: contract}}, nil
}

type ArbitrumCrossDomainGovernor struct {
	address common.Address
	abi     abi.ABI
	ArbitrumCrossDomainGovernorCaller
	ArbitrumCrossDomainGovernorTransactor
	ArbitrumCrossDomainGovernorFilterer
}

type ArbitrumCrossDomainGovernorCaller struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorSession struct {
	Contract     *ArbitrumCrossDomainGovernor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainGovernorCallerSession struct {
	Contract *ArbitrumCrossDomainGovernorCaller
	CallOpts bind.CallOpts
}

type ArbitrumCrossDomainGovernorTransactorSession struct {
	Contract     *ArbitrumCrossDomainGovernorTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainGovernorRaw struct {
	Contract *ArbitrumCrossDomainGovernor
}

type ArbitrumCrossDomainGovernorCallerRaw struct {
	Contract *ArbitrumCrossDomainGovernorCaller
}

type ArbitrumCrossDomainGovernorTransactorRaw struct {
	Contract *ArbitrumCrossDomainGovernorTransactor
}

func NewArbitrumCrossDomainGovernor(address common.Address, backend bind.ContractBackend) (*ArbitrumCrossDomainGovernor, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumCrossDomainGovernorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumCrossDomainGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernor{address: address, abi: abi, ArbitrumCrossDomainGovernorCaller: ArbitrumCrossDomainGovernorCaller{contract: contract}, ArbitrumCrossDomainGovernorTransactor: ArbitrumCrossDomainGovernorTransactor{contract: contract}, ArbitrumCrossDomainGovernorFilterer: ArbitrumCrossDomainGovernorFilterer{contract: contract}}, nil
}

func NewArbitrumCrossDomainGovernorCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumCrossDomainGovernorCaller, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorCaller{contract: contract}, nil
}

func NewArbitrumCrossDomainGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumCrossDomainGovernorTransactor, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorTransactor{contract: contract}, nil
}

func NewArbitrumCrossDomainGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumCrossDomainGovernorFilterer, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorFilterer{contract: contract}, nil
}

func bindArbitrumCrossDomainGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorTransactor.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.L1Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.L1Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TypeAndVersion(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TypeAndVersion(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "acceptL1Ownership")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "acceptOwnership")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "forward", target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Forward(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Forward(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "forwardDelegate", target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ForwardDelegate(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ForwardDelegate(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "transferL1Ownership", to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "transferOwnership", to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainGovernorL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
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

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainGovernorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainGovernorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferred)
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

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorOwnershipTransferredIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorOwnershipTransferred)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainGovernorOwnershipTransferred)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ArbitrumCrossDomainGovernorL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (ArbitrumCrossDomainGovernorL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (ArbitrumCrossDomainGovernorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ArbitrumCrossDomainGovernorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernor) Address() common.Address {
	return _ArbitrumCrossDomainGovernor.address
}

type ArbitrumCrossDomainGovernorInterface interface {
	CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error)

	L1Owner(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferred, error)

	Address() common.Address
}
