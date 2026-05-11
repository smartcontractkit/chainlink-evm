// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_cross_domain_governor

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

var OptimismCrossDomainGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"crossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"contractiOVM_CrossDomainMessenger\"},{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardDelegate\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a060405234801562000010575f80fd5b506040516200158138038062001581833981016040819052620000339162000269565b81818033805f816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be816200014b565b505050620000d281620001f560201b60201c565b506001600160a01b038216620001355760405162461bcd60e51b815260206004820152602160248201527f496e76616c69642078446f6d61696e204d657373656e676572206164647265736044820152607360f81b606482015260840162000083565b506001600160a01b031660805250620002a69050565b336001600160a01b03821603620001a55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b6001600160a01b038116811462000266575f80fd5b50565b5f80604083850312156200027b575f80fd5b8251620002888162000251565b60208401519092506200029b8162000251565b809150509250929050565b60805161129f620002e25f395f8181610197015281816101dd01528181610435015281816107940152818161084c0152610988015261129f5ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610164578063f2fde38b14610182578063f43b361314610195575f80fd5b806396b8d7c014610149578063b2ec07871461015c575f80fd5b80636fadcf72116100a25780636fadcf72146100f057806379ba5097146101035780638da5cb5b1461010b575f80fd5b8063181f5a77146100bd57806326929eb6146100db575b5f80fd5b6100c56101bb565b6040516100d291906110a2565b60405180910390f35b6100ee6100e9366004611140565b6101db565b005b6100ee6100fe366004611140565b610433565b6100ee610680565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d2565b6100ee61015736600461121c565b61077c565b6100ee610986565b60025473ffffffffffffffffffffffffffffffffffffffff16610124565b6100ee61019036600461121c565b610b59565b7f0000000000000000000000000000000000000000000000000000000000000000610124565b606060405180606001604052806025815260200161126e60259139905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff811633148061023657505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6102c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811633036104235760025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103809190611237565b73ffffffffffffffffffffffffffffffffffffffff1614610423576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102be565b61042d8383610b6a565b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff811633148061048e57505f5473ffffffffffffffffffffffffffffffffffffffff1633145b61051a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084016102be565b73ffffffffffffffffffffffffffffffffffffffff811633036106765760025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105af573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105d39190611237565b73ffffffffffffffffffffffffffffffffffffffff1614610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102be565b61042d8383610be9565b60015473ffffffffffffffffffffffffffffffffffffffff163314610701576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102be565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461081b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064016102be565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108b3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108d79190611237565b73ffffffffffffffffffffffffffffffffffffffff161461097a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f657200000000000000000000000000000000000000000000000000000000000060648201526084016102be565b61098381610bfd565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff81163314610a26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064016102be565b600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff92831692841691636e296e459160048083019260209291908290030181865afa158015610a96573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aba9190611237565b73ffffffffffffffffffffffffffffffffffffffff1614610b37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e65720000000000000060448201526064016102be565b6003546109839073ffffffffffffffffffffffffffffffffffffffff16610cf2565b610b61610d73565b61098381610df5565b60605f808473ffffffffffffffffffffffffffffffffffffffff1684604051610b939190611252565b5f60405180830381855af49150503d805f8114610bcb576040519150601f19603f3d011682016040523d82523d5f602084013e610bd0565b606091505b5091509150610be0858383610ee9565b95945050505050565b6060610bf683835f610f78565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610c7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102be565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610df3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102be565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610e74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102be565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606082610efe57610ef98261103e565b610bf6565b8151158015610f22575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610f71576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016102be565b5092915050565b606081471015610fbd576040517fcf479181000000000000000000000000000000000000000000000000000000008152476004820152602481018390526044016102be565b5f808573ffffffffffffffffffffffffffffffffffffffff168486604051610fe59190611252565b5f6040518083038185875af1925050503d805f811461101f576040519150601f19603f3d011682016040523d82523d5f602084013e611024565b606091505b5091509150611034868383610ee9565b9695505050505050565b80511561104e5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561109a578181015183820152602001611082565b50505f910152565b602081525f82518060208401526110c0816040850160208701611080565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610983575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215611151575f80fd5b823561115c816110f2565b9150602083013567ffffffffffffffff80821115611178575f80fd5b818501915085601f83011261118b575f80fd5b81358181111561119d5761119d611113565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156111e3576111e3611113565b816040528281528860208487010111156111fb575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f6020828403121561122c575f80fd5b8135610bf6816110f2565b5f60208284031215611247575f80fd5b8151610bf6816110f2565b5f8251611263818460208701611080565b919091019291505056fe4f7074696d69736d43726f7373446f6d61696e476f7665726e6f7220312e312e302d646576a164736f6c6343000818000a",
}

var OptimismCrossDomainGovernorABI = OptimismCrossDomainGovernorMetaData.ABI

var OptimismCrossDomainGovernorBin = OptimismCrossDomainGovernorMetaData.Bin

func DeployOptimismCrossDomainGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, crossDomainMessengerAddr common.Address, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *OptimismCrossDomainGovernor, error) {
	parsed, err := OptimismCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismCrossDomainGovernorBin), backend, crossDomainMessengerAddr, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismCrossDomainGovernor{address: address, abi: *parsed, OptimismCrossDomainGovernorCaller: OptimismCrossDomainGovernorCaller{contract: contract}, OptimismCrossDomainGovernorTransactor: OptimismCrossDomainGovernorTransactor{contract: contract}, OptimismCrossDomainGovernorFilterer: OptimismCrossDomainGovernorFilterer{contract: contract}}, nil
}

type OptimismCrossDomainGovernor struct {
	address common.Address
	abi     abi.ABI
	OptimismCrossDomainGovernorCaller
	OptimismCrossDomainGovernorTransactor
	OptimismCrossDomainGovernorFilterer
}

type OptimismCrossDomainGovernorCaller struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainGovernorTransactor struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainGovernorFilterer struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainGovernorSession struct {
	Contract     *OptimismCrossDomainGovernor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismCrossDomainGovernorCallerSession struct {
	Contract *OptimismCrossDomainGovernorCaller
	CallOpts bind.CallOpts
}

type OptimismCrossDomainGovernorTransactorSession struct {
	Contract     *OptimismCrossDomainGovernorTransactor
	TransactOpts bind.TransactOpts
}

type OptimismCrossDomainGovernorRaw struct {
	Contract *OptimismCrossDomainGovernor
}

type OptimismCrossDomainGovernorCallerRaw struct {
	Contract *OptimismCrossDomainGovernorCaller
}

type OptimismCrossDomainGovernorTransactorRaw struct {
	Contract *OptimismCrossDomainGovernorTransactor
}

func NewOptimismCrossDomainGovernor(address common.Address, backend bind.ContractBackend) (*OptimismCrossDomainGovernor, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismCrossDomainGovernorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismCrossDomainGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernor{address: address, abi: abi, OptimismCrossDomainGovernorCaller: OptimismCrossDomainGovernorCaller{contract: contract}, OptimismCrossDomainGovernorTransactor: OptimismCrossDomainGovernorTransactor{contract: contract}, OptimismCrossDomainGovernorFilterer: OptimismCrossDomainGovernorFilterer{contract: contract}}, nil
}

func NewOptimismCrossDomainGovernorCaller(address common.Address, caller bind.ContractCaller) (*OptimismCrossDomainGovernorCaller, error) {
	contract, err := bindOptimismCrossDomainGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorCaller{contract: contract}, nil
}

func NewOptimismCrossDomainGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismCrossDomainGovernorTransactor, error) {
	contract, err := bindOptimismCrossDomainGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorTransactor{contract: contract}, nil
}

func NewOptimismCrossDomainGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismCrossDomainGovernorFilterer, error) {
	contract, err := bindOptimismCrossDomainGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorFilterer{contract: contract}, nil
}

func bindOptimismCrossDomainGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismCrossDomainGovernor.Contract.OptimismCrossDomainGovernorCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.OptimismCrossDomainGovernorTransactor.contract.Transfer(opts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.OptimismCrossDomainGovernorTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismCrossDomainGovernor.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.contract.Transfer(opts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainGovernor.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) CrossDomainMessenger() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.CrossDomainMessenger(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.CrossDomainMessenger(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainGovernor.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) L1Owner() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.L1Owner(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCallerSession) L1Owner() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.L1Owner(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) Owner() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.Owner(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCallerSession) Owner() (common.Address, error) {
	return _OptimismCrossDomainGovernor.Contract.Owner(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismCrossDomainGovernor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) TypeAndVersion() (string, error) {
	return _OptimismCrossDomainGovernor.Contract.TypeAndVersion(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorCallerSession) TypeAndVersion() (string, error) {
	return _OptimismCrossDomainGovernor.Contract.TypeAndVersion(&_OptimismCrossDomainGovernor.CallOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "acceptL1Ownership")
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.AcceptL1Ownership(&_OptimismCrossDomainGovernor.TransactOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.AcceptL1Ownership(&_OptimismCrossDomainGovernor.TransactOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "acceptOwnership")
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.AcceptOwnership(&_OptimismCrossDomainGovernor.TransactOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.AcceptOwnership(&_OptimismCrossDomainGovernor.TransactOpts)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "forward", target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.Forward(&_OptimismCrossDomainGovernor.TransactOpts, target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.Forward(&_OptimismCrossDomainGovernor.TransactOpts, target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "forwardDelegate", target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.ForwardDelegate(&_OptimismCrossDomainGovernor.TransactOpts, target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.ForwardDelegate(&_OptimismCrossDomainGovernor.TransactOpts, target, data)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "transferL1Ownership", to)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.TransferL1Ownership(&_OptimismCrossDomainGovernor.TransactOpts, to)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.TransferL1Ownership(&_OptimismCrossDomainGovernor.TransactOpts, to)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.contract.Transact(opts, "transferOwnership", to)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.TransferOwnership(&_OptimismCrossDomainGovernor.TransactOpts, to)
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainGovernor.Contract.TransferOwnership(&_OptimismCrossDomainGovernor.TransactOpts, to)
}

type OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator struct {
	Event *OptimismCrossDomainGovernorL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainGovernorL1OwnershipTransferRequested)
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
		it.Event = new(OptimismCrossDomainGovernorL1OwnershipTransferRequested)
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

func (it *OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainGovernorL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator{contract: _OptimismCrossDomainGovernor.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainGovernorL1OwnershipTransferRequested)
				if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*OptimismCrossDomainGovernorL1OwnershipTransferRequested, error) {
	event := new(OptimismCrossDomainGovernorL1OwnershipTransferRequested)
	if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainGovernorL1OwnershipTransferredIterator struct {
	Event *OptimismCrossDomainGovernorL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainGovernorL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainGovernorL1OwnershipTransferred)
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
		it.Event = new(OptimismCrossDomainGovernorL1OwnershipTransferred)
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

func (it *OptimismCrossDomainGovernorL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainGovernorL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainGovernorL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorL1OwnershipTransferredIterator{contract: _OptimismCrossDomainGovernor.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainGovernorL1OwnershipTransferred)
				if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) ParseL1OwnershipTransferred(log types.Log) (*OptimismCrossDomainGovernorL1OwnershipTransferred, error) {
	event := new(OptimismCrossDomainGovernorL1OwnershipTransferred)
	if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainGovernorOwnershipTransferRequestedIterator struct {
	Event *OptimismCrossDomainGovernorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainGovernorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainGovernorOwnershipTransferRequested)
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
		it.Event = new(OptimismCrossDomainGovernorOwnershipTransferRequested)
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

func (it *OptimismCrossDomainGovernorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainGovernorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainGovernorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorOwnershipTransferRequestedIterator{contract: _OptimismCrossDomainGovernor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainGovernorOwnershipTransferRequested)
				if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OptimismCrossDomainGovernorOwnershipTransferRequested, error) {
	event := new(OptimismCrossDomainGovernorOwnershipTransferRequested)
	if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainGovernorOwnershipTransferredIterator struct {
	Event *OptimismCrossDomainGovernorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainGovernorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainGovernorOwnershipTransferred)
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
		it.Event = new(OptimismCrossDomainGovernorOwnershipTransferred)
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

func (it *OptimismCrossDomainGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainGovernorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainGovernorOwnershipTransferredIterator{contract: _OptimismCrossDomainGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainGovernorOwnershipTransferred)
				if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*OptimismCrossDomainGovernorOwnershipTransferred, error) {
	event := new(OptimismCrossDomainGovernorOwnershipTransferred)
	if err := _OptimismCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OptimismCrossDomainGovernorL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (OptimismCrossDomainGovernorL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (OptimismCrossDomainGovernorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OptimismCrossDomainGovernorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OptimismCrossDomainGovernor *OptimismCrossDomainGovernor) Address() common.Address {
	return _OptimismCrossDomainGovernor.address
}

type OptimismCrossDomainGovernorInterface interface {
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

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*OptimismCrossDomainGovernorL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*OptimismCrossDomainGovernorL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OptimismCrossDomainGovernorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainGovernorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OptimismCrossDomainGovernorOwnershipTransferred, error)

	Address() common.Address
}
