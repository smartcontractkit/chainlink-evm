// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scroll_cross_domain_governor

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

var ScrollCrossDomainGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"crossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"contract IScrollMessenger\"},{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardDelegate\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a060405234801562000010575f80fd5b506040516200163538038062001635833981016040819052620000339162000264565b8033805f816200008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000bc57620000bc8162000146565b505050620000d081620001f060201b60201c565b506001600160a01b038216620001335760405162461bcd60e51b815260206004820152602160248201527f496e76616c69642078446f6d61696e204d657373656e676572206164647265736044820152607360f81b606482015260840162000081565b506001600160a01b0316608052620002a1565b336001600160a01b03821603620001a05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000081565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b6001600160a01b038116811462000261575f80fd5b50565b5f806040838503121562000276575f80fd5b825162000283816200024c565b602084015190925062000296816200024c565b809150509250929050565b608051611335620003005f395f8181610197015281816101ef015281816102d901528181610330015281816104850152818161056a015281816105c10152818161080d015281816108c501528181610a170152610ae501526113355ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610164578063f2fde38b14610182578063f43b361314610195575f80fd5b806396b8d7c014610149578063b2ec07871461015c575f80fd5b80636fadcf72116100a25780636fadcf72146100f057806379ba5097146101035780638da5cb5b1461010b575f80fd5b8063181f5a77146100bd57806326929eb6146100db575b5f80fd5b6100c56101bb565b6040516100d2919061113a565b60405180910390f35b6100ee6100e93660046111d8565b6101d7565b005b6100ee6100fe3660046111d8565b61046d565b6100ee6106f9565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d2565b6100ee6101573660046112b4565b6107f5565b6100ee6109ff565b60025473ffffffffffffffffffffffffffffffffffffffff16610124565b6100ee6101903660046112b4565b610bf3565b7f0000000000000000000000000000000000000000000000000000000000000000610124565b6040518060600160405280602381526020016113066023913981565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061023157505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6102c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016330361045e5760025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610397573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103bb91906112cf565b73ffffffffffffffffffffffffffffffffffffffff161461045e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102b9565b6104688282610c04565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614806104c757505f5473ffffffffffffffffffffffffffffffffffffffff1633145b610553576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084016102b9565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036106ef5760025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610628573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064c91906112cf565b73ffffffffffffffffffffffffffffffffffffffff16146106ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102b9565b6104688282610c83565b60015473ffffffffffffffffffffffffffffffffffffffff16331461077a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102b9565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610894576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064016102b9565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561092c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095091906112cf565b73ffffffffffffffffffffffffffffffffffffffff16146109f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102b9565b6109fc81610c97565b50565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064016102b9565b600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff928316927f00000000000000000000000000000000000000000000000000000000000000001691636e296e459160048083019260209291908290030181865afa158015610b2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5291906112cf565b73ffffffffffffffffffffffffffffffffffffffff1614610bcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e65720000000000000060448201526064016102b9565b600354610bf19073ffffffffffffffffffffffffffffffffffffffff16610d8c565b565b610bfb610e0d565b6109fc81610e8d565b60605f808473ffffffffffffffffffffffffffffffffffffffff1684604051610c2d91906112ea565b5f60405180830381855af49150503d805f8114610c65576040519150601f19603f3d011682016040523d82523d5f602084013e610c6a565b606091505b5091509150610c7a858383610f81565b95945050505050565b6060610c9083835f611010565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610d16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b9565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610bf1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102b9565b3373ffffffffffffffffffffffffffffffffffffffff821603610f0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b9565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606082610f9657610f91826110d6565b610c90565b8151158015610fba575073ffffffffffffffffffffffffffffffffffffffff84163b155b15611009576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016102b9565b5092915050565b606081471015611055576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018390526044016102b9565b5f808573ffffffffffffffffffffffffffffffffffffffff16848660405161107d91906112ea565b5f6040518083038185875af1925050503d805f81146110b7576040519150601f19603f3d011682016040523d82523d5f602084013e6110bc565b606091505b50915091506110cc868383610f81565b9695505050505050565b8051156110e65780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561113257818101518382015260200161111a565b50505f910152565b602081525f8251806020840152611158816040850160208701611118565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146109fc575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80604083850312156111e9575f80fd5b82356111f48161118a565b9150602083013567ffffffffffffffff80821115611210575f80fd5b818501915085601f830112611223575f80fd5b813581811115611235576112356111ab565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561127b5761127b6111ab565b81604052828152886020848701011115611293575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f602082840312156112c4575f80fd5b8135610c908161118a565b5f602082840312156112df575f80fd5b8151610c908161118a565b5f82516112fb818460208701611118565b919091019291505056fe5363726f6c6c43726f7373446f6d61696e476f7665726e6f7220312e312e302d646576a164736f6c6343000818000a",
}

var ScrollCrossDomainGovernorABI = ScrollCrossDomainGovernorMetaData.ABI

var ScrollCrossDomainGovernorBin = ScrollCrossDomainGovernorMetaData.Bin

func DeployScrollCrossDomainGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, crossDomainMessengerAddr common.Address, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *ScrollCrossDomainGovernor, error) {
	parsed, err := ScrollCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScrollCrossDomainGovernorBin), backend, crossDomainMessengerAddr, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScrollCrossDomainGovernor{address: address, abi: *parsed, ScrollCrossDomainGovernorCaller: ScrollCrossDomainGovernorCaller{contract: contract}, ScrollCrossDomainGovernorTransactor: ScrollCrossDomainGovernorTransactor{contract: contract}, ScrollCrossDomainGovernorFilterer: ScrollCrossDomainGovernorFilterer{contract: contract}}, nil
}

type ScrollCrossDomainGovernor struct {
	address common.Address
	abi     abi.ABI
	ScrollCrossDomainGovernorCaller
	ScrollCrossDomainGovernorTransactor
	ScrollCrossDomainGovernorFilterer
}

type ScrollCrossDomainGovernorCaller struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainGovernorTransactor struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainGovernorFilterer struct {
	contract *bind.BoundContract
}

type ScrollCrossDomainGovernorSession struct {
	Contract     *ScrollCrossDomainGovernor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ScrollCrossDomainGovernorCallerSession struct {
	Contract *ScrollCrossDomainGovernorCaller
	CallOpts bind.CallOpts
}

type ScrollCrossDomainGovernorTransactorSession struct {
	Contract     *ScrollCrossDomainGovernorTransactor
	TransactOpts bind.TransactOpts
}

type ScrollCrossDomainGovernorRaw struct {
	Contract *ScrollCrossDomainGovernor
}

type ScrollCrossDomainGovernorCallerRaw struct {
	Contract *ScrollCrossDomainGovernorCaller
}

type ScrollCrossDomainGovernorTransactorRaw struct {
	Contract *ScrollCrossDomainGovernorTransactor
}

func NewScrollCrossDomainGovernor(address common.Address, backend bind.ContractBackend) (*ScrollCrossDomainGovernor, error) {
	abi, err := abi.JSON(strings.NewReader(ScrollCrossDomainGovernorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindScrollCrossDomainGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernor{address: address, abi: abi, ScrollCrossDomainGovernorCaller: ScrollCrossDomainGovernorCaller{contract: contract}, ScrollCrossDomainGovernorTransactor: ScrollCrossDomainGovernorTransactor{contract: contract}, ScrollCrossDomainGovernorFilterer: ScrollCrossDomainGovernorFilterer{contract: contract}}, nil
}

func NewScrollCrossDomainGovernorCaller(address common.Address, caller bind.ContractCaller) (*ScrollCrossDomainGovernorCaller, error) {
	contract, err := bindScrollCrossDomainGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorCaller{contract: contract}, nil
}

func NewScrollCrossDomainGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollCrossDomainGovernorTransactor, error) {
	contract, err := bindScrollCrossDomainGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorTransactor{contract: contract}, nil
}

func NewScrollCrossDomainGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollCrossDomainGovernorFilterer, error) {
	contract, err := bindScrollCrossDomainGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorFilterer{contract: contract}, nil
}

func bindScrollCrossDomainGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScrollCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollCrossDomainGovernor.Contract.ScrollCrossDomainGovernorCaller.contract.Call(opts, result, method, params...)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.ScrollCrossDomainGovernorTransactor.contract.Transfer(opts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.ScrollCrossDomainGovernorTransactor.contract.Transact(opts, method, params...)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollCrossDomainGovernor.Contract.contract.Call(opts, result, method, params...)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.contract.Transfer(opts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.contract.Transact(opts, method, params...)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainGovernor.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) CrossDomainMessenger() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.CrossDomainMessenger(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.CrossDomainMessenger(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainGovernor.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) L1Owner() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.L1Owner(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCallerSession) L1Owner() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.L1Owner(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollCrossDomainGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) Owner() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.Owner(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCallerSession) Owner() (common.Address, error) {
	return _ScrollCrossDomainGovernor.Contract.Owner(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrollCrossDomainGovernor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) TypeAndVersion() (string, error) {
	return _ScrollCrossDomainGovernor.Contract.TypeAndVersion(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorCallerSession) TypeAndVersion() (string, error) {
	return _ScrollCrossDomainGovernor.Contract.TypeAndVersion(&_ScrollCrossDomainGovernor.CallOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "acceptL1Ownership")
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.AcceptL1Ownership(&_ScrollCrossDomainGovernor.TransactOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.AcceptL1Ownership(&_ScrollCrossDomainGovernor.TransactOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "acceptOwnership")
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.AcceptOwnership(&_ScrollCrossDomainGovernor.TransactOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.AcceptOwnership(&_ScrollCrossDomainGovernor.TransactOpts)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "forward", target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.Forward(&_ScrollCrossDomainGovernor.TransactOpts, target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.Forward(&_ScrollCrossDomainGovernor.TransactOpts, target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "forwardDelegate", target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.ForwardDelegate(&_ScrollCrossDomainGovernor.TransactOpts, target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.ForwardDelegate(&_ScrollCrossDomainGovernor.TransactOpts, target, data)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "transferL1Ownership", to)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.TransferL1Ownership(&_ScrollCrossDomainGovernor.TransactOpts, to)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.TransferL1Ownership(&_ScrollCrossDomainGovernor.TransactOpts, to)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.contract.Transact(opts, "transferOwnership", to)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.TransferOwnership(&_ScrollCrossDomainGovernor.TransactOpts, to)
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollCrossDomainGovernor.Contract.TransferOwnership(&_ScrollCrossDomainGovernor.TransactOpts, to)
}

type ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator struct {
	Event *ScrollCrossDomainGovernorL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainGovernorL1OwnershipTransferRequested)
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
		it.Event = new(ScrollCrossDomainGovernorL1OwnershipTransferRequested)
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

func (it *ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainGovernorL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator{contract: _ScrollCrossDomainGovernor.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainGovernorL1OwnershipTransferRequested)
				if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*ScrollCrossDomainGovernorL1OwnershipTransferRequested, error) {
	event := new(ScrollCrossDomainGovernorL1OwnershipTransferRequested)
	if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainGovernorL1OwnershipTransferredIterator struct {
	Event *ScrollCrossDomainGovernorL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainGovernorL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainGovernorL1OwnershipTransferred)
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
		it.Event = new(ScrollCrossDomainGovernorL1OwnershipTransferred)
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

func (it *ScrollCrossDomainGovernorL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainGovernorL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainGovernorL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorL1OwnershipTransferredIterator{contract: _ScrollCrossDomainGovernor.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainGovernorL1OwnershipTransferred)
				if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) ParseL1OwnershipTransferred(log types.Log) (*ScrollCrossDomainGovernorL1OwnershipTransferred, error) {
	event := new(ScrollCrossDomainGovernorL1OwnershipTransferred)
	if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainGovernorOwnershipTransferRequestedIterator struct {
	Event *ScrollCrossDomainGovernorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainGovernorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainGovernorOwnershipTransferRequested)
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
		it.Event = new(ScrollCrossDomainGovernorOwnershipTransferRequested)
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

func (it *ScrollCrossDomainGovernorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainGovernorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainGovernorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorOwnershipTransferRequestedIterator{contract: _ScrollCrossDomainGovernor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainGovernorOwnershipTransferRequested)
				if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ScrollCrossDomainGovernorOwnershipTransferRequested, error) {
	event := new(ScrollCrossDomainGovernorOwnershipTransferRequested)
	if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollCrossDomainGovernorOwnershipTransferredIterator struct {
	Event *ScrollCrossDomainGovernorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollCrossDomainGovernorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollCrossDomainGovernorOwnershipTransferred)
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
		it.Event = new(ScrollCrossDomainGovernorOwnershipTransferred)
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

func (it *ScrollCrossDomainGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ScrollCrossDomainGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollCrossDomainGovernorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollCrossDomainGovernorOwnershipTransferredIterator{contract: _ScrollCrossDomainGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollCrossDomainGovernorOwnershipTransferred)
				if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollCrossDomainGovernorOwnershipTransferred, error) {
	event := new(ScrollCrossDomainGovernorOwnershipTransferred)
	if err := _ScrollCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ScrollCrossDomainGovernorL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (ScrollCrossDomainGovernorL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (ScrollCrossDomainGovernorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ScrollCrossDomainGovernorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ScrollCrossDomainGovernor *ScrollCrossDomainGovernor) Address() common.Address {
	return _ScrollCrossDomainGovernor.address
}

type ScrollCrossDomainGovernorInterface interface {
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

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*ScrollCrossDomainGovernorL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*ScrollCrossDomainGovernorL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ScrollCrossDomainGovernorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollCrossDomainGovernorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ScrollCrossDomainGovernorOwnershipTransferred, error)

	Address() common.Address
}
