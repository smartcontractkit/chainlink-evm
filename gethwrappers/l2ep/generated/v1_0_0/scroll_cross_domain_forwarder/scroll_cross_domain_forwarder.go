// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scroll_cross_domain_forwarder

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

var ScrollCrossDomainForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"crossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"contract IScrollMessenger\"},{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801562000010575f80fd5b50604051620012da380380620012da833981016040819052620000339162000264565b8033805f816200008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000bc57620000bc8162000146565b505050620000d081620001f060201b60201c565b506001600160a01b038216620001335760405162461bcd60e51b815260206004820152602160248201527f496e76616c69642078446f6d61696e204d657373656e676572206164647265736044820152607360f81b606482015260840162000081565b506001600160a01b0316608052620002a1565b336001600160a01b03821603620001a05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000081565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b6001600160a01b038116811462000261575f80fd5b50565b5f806040838503121562000276575f80fd5b825162000283816200024c565b602084015190925062000296816200024c565b809150509250929050565b608051610ff6620002e45f395f818161019e015281816101da015281816102970152818161051e015281816105d60152818161072801526107f60152610ff65ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db63721461016b578063f2fde38b14610189578063f43b36131461019c575f80fd5b806396b8d7c014610150578063b2ec078714610163575f80fd5b8063181f5a77146100a35780636fadcf72146100f557806379ba50971461010a5780638da5cb5b14610112575b5f80fd5b6100df6040518060400160405280602081526020017f5363726f6c6c43726f7373446f6d61696e466f7277617264657220312e302e3081525081565b6040516100ec9190610e1e565b60405180910390f35b610108610103366004610ebc565b6101c2565b005b61010861040a565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ec565b61010861015e366004610f98565b610506565b610108610710565b60025473ffffffffffffffffffffffffffffffffffffffff1661012b565b610108610197366004610f98565b610904565b7f000000000000000000000000000000000000000000000000000000000000000061012b565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064015b60405180910390fd5b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102fe573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103229190610fb3565b73ffffffffffffffffffffffffffffffffffffffff16146103c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f6572000000000000000000000000000000000000000000000000000000000000606482015260840161025d565b61040582826040518060400160405280601781526020017f466f727761726465722063616c6c207265766572746564000000000000000000815250610915565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461048b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161025d565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e6765720000604482015260640161025d565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561063d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106619190610fb3565b73ffffffffffffffffffffffffffffffffffffffff1614610704576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f6572000000000000000000000000000000000000000000000000000000000000606482015260840161025d565b61070d8161092d565b50565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e6765720000604482015260640161025d565b600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff928316927f00000000000000000000000000000000000000000000000000000000000000001691636e296e459160048083019260209291908290030181865afa15801561083f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108639190610fb3565b73ffffffffffffffffffffffffffffffffffffffff16146108e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e657200000000000000604482015260640161025d565b6003546109029073ffffffffffffffffffffffffffffffffffffffff16610a22565b565b61090c610aa3565b61070d81610b23565b606061092384845f85610c17565b90505b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036109ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161025d565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610902576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161025d565b3373ffffffffffffffffffffffffffffffffffffffff821603610ba2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161025d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606082471015610ca9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161025d565b73ffffffffffffffffffffffffffffffffffffffff85163b610d27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161025d565b5f808673ffffffffffffffffffffffffffffffffffffffff168587604051610d4f9190610fce565b5f6040518083038185875af1925050503d805f8114610d89576040519150601f19603f3d011682016040523d82523d5f602084013e610d8e565b606091505b5091509150610d9e828286610da9565b979650505050505050565b60608315610db8575081610926565b825115610dc85782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025d9190610e1e565b5f5b83811015610e16578181015183820152602001610dfe565b50505f910152565b602081525f8251806020840152610e3c816040850160208701610dfc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461070d575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610ecd575f80fd5b8235610ed881610e6e565b9150602083013567ffffffffffffffff80821115610ef4575f80fd5b818501915085601f830112610f07575f80fd5b813581811115610f1957610f19610e8f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610f5f57610f5f610e8f565b81604052828152886020848701011115610f77575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610fa8575f80fd5b813561092681610e6e565b5f60208284031215610fc3575f80fd5b815161092681610e6e565b5f8251610fdf818460208701610dfc565b919091019291505056fea164736f6c6343000818000a",
}

var ScrollCrossDomainForwarderABI = ScrollCrossDomainForwarderMetaData.ABI

var ScrollCrossDomainForwarderBin = ScrollCrossDomainForwarderMetaData.Bin

func DeployScrollCrossDomainForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, crossDomainMessengerAddr common.Address, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *ScrollCrossDomainForwarder, error) {
	parsed, err := ScrollCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScrollCrossDomainForwarderBin), backend, crossDomainMessengerAddr, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScrollCrossDomainForwarder{address: address, abi: *parsed, ScrollCrossDomainForwarderCaller: ScrollCrossDomainForwarderCaller{contract: contract}, ScrollCrossDomainForwarderTransactor: ScrollCrossDomainForwarderTransactor{contract: contract}, ScrollCrossDomainForwarderFilterer: ScrollCrossDomainForwarderFilterer{contract: contract}}, nil
}

type ScrollCrossDomainForwarder struct {
	address common.Address
	abi     abi.ABI
	ScrollCrossDomainForwarderCaller
	ScrollCrossDomainForwarderTransactor
	ScrollCrossDomainForwarderFilterer
}

type ScrollCrossDomainForwarderCaller struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainForwarderTransactor struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainForwarderFilterer struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainForwarderSession struct {
	Contract     *ScrollCrossDomainForwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ScrollCrossDomainForwarderCallerSession struct {
	Contract *ScrollCrossDomainForwarderCaller
	CallOpts bind.CallOpts
}

type ScrollCrossDomainForwarderTransactorSession struct {
	Contract     *ScrollCrossDomainForwarderTransactor
	TransactOpts bind.TransactOpts
}

type ScrollCrossDomainForwarderRaw struct {
	Contract *ScrollCrossDomainForwarder
}

type ScrollCrossDomainForwarderCallerRaw struct {
	Contract *ScrollCrossDomainForwarderCaller
}

type ScrollCrossDomainForwarderTransactorRaw struct {
	Contract *ScrollCrossDomainForwarderTransactor
}

func NewScrollCrossDomainForwarder(address common.Address, backend bind.ContractBackend) (*ScrollCrossDomainForwarder, error) {
	abi, err := abi.JSON(strings.NewReader(ScrollCrossDomainForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindScrollCrossDomainForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarder{address: address, abi: abi, ScrollCrossDomainForwarderCaller: ScrollCrossDomainForwarderCaller{contract: contract}, ScrollCrossDomainForwarderTransactor: ScrollCrossDomainForwarderTransactor{contract: contract}, ScrollCrossDomainForwarderFilterer: ScrollCrossDomainForwarderFilterer{contract: contract}}, nil
}

func NewScrollCrossDomainForwarderCaller(address common.Address, caller bind.ContractCaller) (*ScrollCrossDomainForwarderCaller, error) {
	contract, err := bindScrollCrossDomainForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderCaller{contract: contract}, nil
}

func NewScrollCrossDomainForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollCrossDomainForwarderTransactor, error) {
	contract, err := bindScrollCrossDomainForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderTransactor{contract: contract}, nil
}

func NewScrollCrossDomainForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollCrossDomainForwarderFilterer, error) {
	contract, err := bindScrollCrossDomainForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderFilterer{contract: contract}, nil
}

func bindScrollCrossDomainForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScrollCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollCrossDomainForwarder.Contract.ScrollCrossDomainForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.ScrollCrossDomainForwarderTransactor.contract.Transfer(opts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.ScrollCrossDomainForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollCrossDomainForwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.contract.Transfer(opts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.contract.Transact(opts, method, params...)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainForwarder.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) CrossDomainMessenger() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.CrossDomainMessenger(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.CrossDomainMessenger(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainForwarder.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) L1Owner() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.L1Owner(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCallerSession) L1Owner() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.L1Owner(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainForwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) Owner() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.Owner(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCallerSession) Owner() (common.Address, error) {
	return _ScrollCrossDomainForwarder.Contract.Owner(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrollCrossDomainForwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) TypeAndVersion() (string, error) {
	return _ScrollCrossDomainForwarder.Contract.TypeAndVersion(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderCallerSession) TypeAndVersion() (string, error) {
	return _ScrollCrossDomainForwarder.Contract.TypeAndVersion(&_ScrollCrossDomainForwarder.CallOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.contract.Transact(opts, "acceptL1Ownership")
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.AcceptL1Ownership(&_ScrollCrossDomainForwarder.TransactOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.AcceptL1Ownership(&_ScrollCrossDomainForwarder.TransactOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.contract.Transact(opts, "acceptOwnership")
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.AcceptOwnership(&_ScrollCrossDomainForwarder.TransactOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.AcceptOwnership(&_ScrollCrossDomainForwarder.TransactOpts)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.contract.Transact(opts, "forward", target, data)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.Forward(&_ScrollCrossDomainForwarder.TransactOpts, target, data)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.Forward(&_ScrollCrossDomainForwarder.TransactOpts, target, data)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.contract.Transact(opts, "transferL1Ownership", to)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.TransferL1Ownership(&_ScrollCrossDomainForwarder.TransactOpts, to)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.TransferL1Ownership(&_ScrollCrossDomainForwarder.TransactOpts, to)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.TransferOwnership(&_ScrollCrossDomainForwarder.TransactOpts, to)
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainForwarder.Contract.TransferOwnership(&_ScrollCrossDomainForwarder.TransactOpts, to)
}

type ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator struct {
	Event *ScrollCrossDomainForwarderL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainForwarderL1OwnershipTransferRequested)
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
		it.Event = new(ScrollCrossDomainForwarderL1OwnershipTransferRequested)
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

func (it *ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainForwarderL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator{contract: _ScrollCrossDomainForwarder.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainForwarderL1OwnershipTransferRequested)
				if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*ScrollCrossDomainForwarderL1OwnershipTransferRequested, error) {
	event := new(ScrollCrossDomainForwarderL1OwnershipTransferRequested)
	if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainForwarderL1OwnershipTransferredIterator struct {
	Event *ScrollCrossDomainForwarderL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainForwarderL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainForwarderL1OwnershipTransferred)
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
		it.Event = new(ScrollCrossDomainForwarderL1OwnershipTransferred)
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

func (it *ScrollCrossDomainForwarderL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainForwarderL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainForwarderL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderL1OwnershipTransferredIterator{contract: _ScrollCrossDomainForwarder.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainForwarderL1OwnershipTransferred)
				if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) ParseL1OwnershipTransferred(log types.Log) (*ScrollCrossDomainForwarderL1OwnershipTransferred, error) {
	event := new(ScrollCrossDomainForwarderL1OwnershipTransferred)
	if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainForwarderOwnershipTransferRequestedIterator struct {
	Event *ScrollCrossDomainForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainForwarderOwnershipTransferRequested)
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
		it.Event = new(ScrollCrossDomainForwarderOwnershipTransferRequested)
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

func (it *ScrollCrossDomainForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderOwnershipTransferRequestedIterator{contract: _ScrollCrossDomainForwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainForwarderOwnershipTransferRequested)
				if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*ScrollCrossDomainForwarderOwnershipTransferRequested, error) {
	event := new(ScrollCrossDomainForwarderOwnershipTransferRequested)
	if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainForwarderOwnershipTransferredIterator struct {
	Event *ScrollCrossDomainForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainForwarderOwnershipTransferred)
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
		it.Event = new(ScrollCrossDomainForwarderOwnershipTransferred)
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

func (it *ScrollCrossDomainForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainForwarderOwnershipTransferredIterator{contract: _ScrollCrossDomainForwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainForwarderOwnershipTransferred)
				if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollCrossDomainForwarderOwnershipTransferred, error) {
	event := new(ScrollCrossDomainForwarderOwnershipTransferred)
	if err := _ScrollCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ScrollCrossDomainForwarderL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (ScrollCrossDomainForwarderL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (ScrollCrossDomainForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ScrollCrossDomainForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ScrollCrossDomainForwarder *ScrollCrossDomainForwarder) Address() common.Address {
	return _ScrollCrossDomainForwarder.address
}

type ScrollCrossDomainForwarderInterface interface {
	CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error)

	L1Owner(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*ScrollCrossDomainForwarderL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*ScrollCrossDomainForwarderL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ScrollCrossDomainForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ScrollCrossDomainForwarderOwnershipTransferred, error)

	Address() common.Address
}
