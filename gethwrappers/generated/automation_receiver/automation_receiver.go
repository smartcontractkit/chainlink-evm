// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package automation_receiver

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

var AutomationReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getExpectedAuthor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getForwarderAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedAuthor\",\"inputs\":[{\"name\":\"_author\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowId\",\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAllowedSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedAuthorUpdated\",\"inputs\":[{\"name\":\"previousAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowIdUpdated\",\"inputs\":[{\"name\":\"previousId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowNameUpdated\",\"inputs\":[{\"name\":\"previousName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"},{\"name\":\"newName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwarderAddressUpdated\",\"inputs\":[{\"name\":\"previousForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityWarning\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallNotAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidAuthor\",\"inputs\":[{\"name\":\"received\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidForwarderAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTargetAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWorkflowId\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowName\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"expected\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}]},{\"type\":\"error\",\"name\":\"MissingSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameRequiresAuthorValidation\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162001a2738038062001a2783398101604081905261003191610101565b8061003b336100b1565b6001600160a01b0381166100615760405162e0775560e61b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040516000907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e908290a35050610131565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561011357600080fd5b81516001600160a01b038116811461012a57600080fd5b9392505050565b6118e680620001416000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80639c1c77ca11610097578063d60c884b11610066578063d60c884b146102ac578063d777cc6d146102bf578063f2fde38b146102d2578063f5c793ef146102e557600080fd5b80639c1c77ca14610224578063a619d81814610237578063bc1fc27a14610286578063c3c44ac21461029957600080fd5b8063715018a6116100d3578063715018a61461017f578063797c8d6914610189578063805f2132146101f35780638da5cb5b1461020657600080fd5b806301ffc9a7146100fa5780633397cf67146101225780633441856f14610161575b600080fd5b61010d61010836600461140a565b6102f6565b60405190151581526020015b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610119565b60015473ffffffffffffffffffffffffffffffffffffffff1661013c565b61018761038f565b005b61010d61019736600461144e565b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290205460ff1692915050565b6101876102013660046114cc565b6103a3565b60005473ffffffffffffffffffffffffffffffffffffffff1661013c565b610187610232366004611538565b610765565b60025474010000000000000000000000000000000000000000900460b01b6040517fffffffffffffffffffff000000000000000000000000000000000000000000009091168152602001610119565b610187610294366004611586565b610874565b6101876102a73660046115c8565b610af6565b6101876102ba3660046115e1565b610b37565b6101876102cd3660046115e1565b610bb6565b6101876102e03660046115e1565b610ce4565b600354604051908152602001610119565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f805f213200000000000000000000000000000000000000000000000000000000148061038957507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b610397610d9b565b6103a16000610e1c565b565b60015473ffffffffffffffffffffffffffffffffffffffff16158015906103e2575060015473ffffffffffffffffffffffffffffffffffffffff163314155b15610440576001546040517fe1130dba00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b600354151580610467575060025473ffffffffffffffffffffffffffffffffffffffff1615155b806104b0575060025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff000000000000000000000000000000000000000000001615155b156107555760008060006104f987878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e9192505050565b60035492955090935091501580159061051457506003548314155b15610559576003546040517f9bfa39ba000000000000000000000000000000000000000000000000000000008152610437918591600401918252602082015260400190565b60025473ffffffffffffffffffffffffffffffffffffffff161580159061059b575060025473ffffffffffffffffffffffffffffffffffffffff828116911614155b156105f6576002546040517fb8a98af800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301529091166024820152604401610437565b60025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff0000000000000000000000000000000000000000000016156107515760025473ffffffffffffffffffffffffffffffffffffffff1661068a576040517f4847901100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547fffffffffffffffffffff000000000000000000000000000000000000000000008381167401000000000000000000000000000000000000000090920460b01b1614610751576002546040517f6c4609a60000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffff0000000000000000000000000000000000000000000084811660048301527401000000000000000000000000000000000000000090920460b01b9091166024820152604401610437565b5050505b61075f8282610eaa565b50505050565b61076d610d9b565b73ffffffffffffffffffffffffffffffffffffffff83166107ba576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831660008181526004602090815260408083207fffffffff0000000000000000000000000000000000000000000000000000000087168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f0925d576b7c865d78d7fe746ae46d080d64b9e6b04db5f034f71a79c41dda2e7910160405180910390a3505050565b61087c610d9b565b60025474010000000000000000000000000000000000000000900460b01b600082900361091f57600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff1690556040516000907fffffffffffffffffffff000000000000000000000000000000000000000000008316907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5908390a3505050565b6000600284846040516109339291906115fe565b602060405180830381855afa158015610950573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610973919061160e565b905060006109a18260405160200161098d91815260200190565b604051602081830303815290604052611178565b60408051600a8082528183019092529192506000919060208201818036833701905050905060005b600a811015610a42578281815181106109e4576109e4611656565b602001015160f81c60f81b828281518110610a0157610a01611656565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080610a3a816116b4565b9150506109c9565b50610a4c816116ec565b600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060b093841c81029190911791829055604051910490911b7fffffffffffffffffffff0000000000000000000000000000000000000000000090811691908616907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa590600090a3505050505050565b610afe610d9b565b6003805490829055604051829082907f0dbedcdf21925e053b4c574eae180d7f2883235ab4976ecc0873598a2a999b0390600090a35050565b610b3f610d9b565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f3321cda85c145617e47418aa14255e9dcbec53a753778e57591703b89a3cad3190600090a35050565b610bbe610d9b565b60015473ffffffffffffffffffffffffffffffffffffffff908116908216610c6e577f704da7db165c79c1e33d542c079333bbde970a733032d2f95fec8fb7d770cbf7604051610c659060208082526038908201527f466f7277617264657220616464726573732073657420746f207a65726f202d2060408201527f636f6e7472616374206973206e6f7720494e5345435552450000000000000000606082015260800190565b60405180910390a15b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e90600090a35050565b610cec610d9b565b73ffffffffffffffffffffffffffffffffffffffff8116610d8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610437565b610d9881610e1c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610437565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60208101516040820151604a83015160601c9193909250565b600080610eb98385018561173c565b909250905073ffffffffffffffffffffffffffffffffffffffff8216610f0b576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600481511015610f47576040517f47d7741900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208181015173ffffffffffffffffffffffffffffffffffffffff841660009081526004835260408082207fffffffff0000000000000000000000000000000000000000000000000000000084168352909352919091205460ff16611018576040517f805043f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201527fffffffff0000000000000000000000000000000000000000000000000000000082166024820152604401610437565b6000808473ffffffffffffffffffffffffffffffffffffffff16846040516110409190611842565b6000604051808303816000865af19150503d806000811461107d576040519150601f19603f3d011682016040523d82523d6000602084013e611082565b606091505b5091509150811561110057827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168573ffffffffffffffffffffffffffffffffffffffff167fbe82131bb3404498c769b0511da41a4ad409fa7152562c2b6669241cbe3bb884836040516110f3919061185e565b60405180910390a361116f565b827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168573ffffffffffffffffffffffffffffffffffffffff167fefa88af289a36b936ccacf9bd9eaaa185775cd54ae263973d3579c01111593b683604051611166919061185e565b60405180910390a35b50505050505050565b606060008251600261118a91906118af565b67ffffffffffffffff8111156111a2576111a2611627565b6040519080825280601f01601f1916602001820160405280156111cc576020820181803683370190505b50905060005b83518110156113ce576040518060400160405280601081526020017f3031323334353637383961626364656600000000000000000000000000000000815250600485838151811061122557611225611656565b016020015182517fff0000000000000000000000000000000000000000000000000000000000000090911690911c60f81c90811061126557611265611656565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016826112988360026118af565b815181106112a8576112a8611656565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506040518060400160405280601081526020017f303132333435363738396162636465660000000000000000000000000000000081525084828151811061131f5761131f611656565b602091010151815160f89190911c600f1690811061133f5761133f611656565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016826113728360026118af565b61137d9060016118c6565b8151811061138d5761138d611656565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806113c6816116b4565b9150506111d2565b5092915050565b80357fffffffff000000000000000000000000000000000000000000000000000000008116811461140557600080fd5b919050565b60006020828403121561141c57600080fd5b611425826113d5565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610d9857600080fd5b6000806040838503121561146157600080fd5b823561146c8161142c565b915061147a602084016113d5565b90509250929050565b60008083601f84011261149557600080fd5b50813567ffffffffffffffff8111156114ad57600080fd5b6020830191508360208285010111156114c557600080fd5b9250929050565b600080600080604085870312156114e257600080fd5b843567ffffffffffffffff808211156114fa57600080fd5b61150688838901611483565b9096509450602087013591508082111561151f57600080fd5b5061152c87828801611483565b95989497509550505050565b60008060006060848603121561154d57600080fd5b83356115588161142c565b9250611566602085016113d5565b91506040840135801515811461157b57600080fd5b809150509250925092565b6000806020838503121561159957600080fd5b823567ffffffffffffffff8111156115b057600080fd5b6115bc85828601611483565b90969095509350505050565b6000602082840312156115da57600080fd5b5035919050565b6000602082840312156115f357600080fd5b81356114258161142c565b8183823760009101908152919050565b60006020828403121561162057600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116e5576116e5611685565b5060010190565b6000815160208301517fffffffffffffffffffff000000000000000000000000000000000000000000008082169350600a83101561173457808184600a0360031b1b83161693505b505050919050565b6000806040838503121561174f57600080fd5b823561175a8161142c565b9150602083013567ffffffffffffffff8082111561177757600080fd5b818501915085601f83011261178b57600080fd5b81358181111561179d5761179d611627565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156117e3576117e3611627565b816040528281528860208487010111156117fc57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b83811015611839578181015183820152602001611821565b50506000910152565b6000825161185481846020870161181e565b9190910192915050565b602081526000825180602084015261187d81604085016020870161181e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b808202811582820484141761038957610389611685565b808201808211156103895761038961168556fea164736f6c6343000813000a",
}

var AutomationReceiverABI = AutomationReceiverMetaData.ABI

var AutomationReceiverBin = AutomationReceiverMetaData.Bin

func DeployAutomationReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, _forwarder common.Address) (common.Address, *types.Transaction, *AutomationReceiver, error) {
	parsed, err := AutomationReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AutomationReceiverBin), backend, _forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AutomationReceiver{address: address, abi: *parsed, AutomationReceiverCaller: AutomationReceiverCaller{contract: contract}, AutomationReceiverTransactor: AutomationReceiverTransactor{contract: contract}, AutomationReceiverFilterer: AutomationReceiverFilterer{contract: contract}}, nil
}

type AutomationReceiver struct {
	address common.Address
	abi     abi.ABI
	AutomationReceiverCaller
	AutomationReceiverTransactor
	AutomationReceiverFilterer
}

type AutomationReceiverCaller struct {
	contract *bind.BoundContract
}

type AutomationReceiverTransactor struct {
	contract *bind.BoundContract
}

type AutomationReceiverFilterer struct {
	contract *bind.BoundContract
}

type AutomationReceiverSession struct {
	Contract     *AutomationReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AutomationReceiverCallerSession struct {
	Contract *AutomationReceiverCaller
	CallOpts bind.CallOpts
}

type AutomationReceiverTransactorSession struct {
	Contract     *AutomationReceiverTransactor
	TransactOpts bind.TransactOpts
}

type AutomationReceiverRaw struct {
	Contract *AutomationReceiver
}

type AutomationReceiverCallerRaw struct {
	Contract *AutomationReceiverCaller
}

type AutomationReceiverTransactorRaw struct {
	Contract *AutomationReceiverTransactor
}

func NewAutomationReceiver(address common.Address, backend bind.ContractBackend) (*AutomationReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(AutomationReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAutomationReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiver{address: address, abi: abi, AutomationReceiverCaller: AutomationReceiverCaller{contract: contract}, AutomationReceiverTransactor: AutomationReceiverTransactor{contract: contract}, AutomationReceiverFilterer: AutomationReceiverFilterer{contract: contract}}, nil
}

func NewAutomationReceiverCaller(address common.Address, caller bind.ContractCaller) (*AutomationReceiverCaller, error) {
	contract, err := bindAutomationReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverCaller{contract: contract}, nil
}

func NewAutomationReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*AutomationReceiverTransactor, error) {
	contract, err := bindAutomationReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverTransactor{contract: contract}, nil
}

func NewAutomationReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*AutomationReceiverFilterer, error) {
	contract, err := bindAutomationReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverFilterer{contract: contract}, nil
}

func bindAutomationReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomationReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_AutomationReceiver *AutomationReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationReceiver.Contract.AutomationReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_AutomationReceiver *AutomationReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.AutomationReceiverTransactor.contract.Transfer(opts)
}

func (_AutomationReceiver *AutomationReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.AutomationReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_AutomationReceiver *AutomationReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_AutomationReceiver *AutomationReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.contract.Transfer(opts)
}

func (_AutomationReceiver *AutomationReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_AutomationReceiver *AutomationReceiverCaller) GetExpectedAuthor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getExpectedAuthor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetExpectedAuthor() (common.Address, error) {
	return _AutomationReceiver.Contract.GetExpectedAuthor(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetExpectedAuthor() (common.Address, error) {
	return _AutomationReceiver.Contract.GetExpectedAuthor(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) GetExpectedWorkflowId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getExpectedWorkflowId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetExpectedWorkflowId() ([32]byte, error) {
	return _AutomationReceiver.Contract.GetExpectedWorkflowId(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetExpectedWorkflowId() ([32]byte, error) {
	return _AutomationReceiver.Contract.GetExpectedWorkflowId(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) GetExpectedWorkflowName(opts *bind.CallOpts) ([10]byte, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getExpectedWorkflowName")

	if err != nil {
		return *new([10]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([10]byte)).(*[10]byte)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetExpectedWorkflowName() ([10]byte, error) {
	return _AutomationReceiver.Contract.GetExpectedWorkflowName(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetExpectedWorkflowName() ([10]byte, error) {
	return _AutomationReceiver.Contract.GetExpectedWorkflowName(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) GetForwarderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getForwarderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetForwarderAddress() (common.Address, error) {
	return _AutomationReceiver.Contract.GetForwarderAddress(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetForwarderAddress() (common.Address, error) {
	return _AutomationReceiver.Contract.GetForwarderAddress(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) IsCallAllowed(opts *bind.CallOpts, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "isCallAllowed", target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) IsCallAllowed(target common.Address, selector [4]byte) (bool, error) {
	return _AutomationReceiver.Contract.IsCallAllowed(&_AutomationReceiver.CallOpts, target, selector)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) IsCallAllowed(target common.Address, selector [4]byte) (bool, error) {
	return _AutomationReceiver.Contract.IsCallAllowed(&_AutomationReceiver.CallOpts, target, selector)
}

func (_AutomationReceiver *AutomationReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) Owner() (common.Address, error) {
	return _AutomationReceiver.Contract.Owner(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) Owner() (common.Address, error) {
	return _AutomationReceiver.Contract.Owner(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AutomationReceiver.Contract.SupportsInterface(&_AutomationReceiver.CallOpts, interfaceId)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AutomationReceiver.Contract.SupportsInterface(&_AutomationReceiver.CallOpts, interfaceId)
}

func (_AutomationReceiver *AutomationReceiverTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "onReport", metadata, report)
}

func (_AutomationReceiver *AutomationReceiverSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.OnReport(&_AutomationReceiver.TransactOpts, metadata, report)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.OnReport(&_AutomationReceiver.TransactOpts, metadata, report)
}

func (_AutomationReceiver *AutomationReceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "renounceOwnership")
}

func (_AutomationReceiver *AutomationReceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _AutomationReceiver.Contract.RenounceOwnership(&_AutomationReceiver.TransactOpts)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AutomationReceiver.Contract.RenounceOwnership(&_AutomationReceiver.TransactOpts)
}

func (_AutomationReceiver *AutomationReceiverTransactor) SetCallAllowed(opts *bind.TransactOpts, target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setCallAllowed", target, selector, allowed)
}

func (_AutomationReceiver *AutomationReceiverSession) SetCallAllowed(target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetCallAllowed(&_AutomationReceiver.TransactOpts, target, selector, allowed)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetCallAllowed(target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetCallAllowed(&_AutomationReceiver.TransactOpts, target, selector, allowed)
}

func (_AutomationReceiver *AutomationReceiverTransactor) SetExpectedAuthor(opts *bind.TransactOpts, _author common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setExpectedAuthor", _author)
}

func (_AutomationReceiver *AutomationReceiverSession) SetExpectedAuthor(_author common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedAuthor(&_AutomationReceiver.TransactOpts, _author)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetExpectedAuthor(_author common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedAuthor(&_AutomationReceiver.TransactOpts, _author)
}

func (_AutomationReceiver *AutomationReceiverTransactor) SetExpectedWorkflowId(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setExpectedWorkflowId", _id)
}

func (_AutomationReceiver *AutomationReceiverSession) SetExpectedWorkflowId(_id [32]byte) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedWorkflowId(&_AutomationReceiver.TransactOpts, _id)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetExpectedWorkflowId(_id [32]byte) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedWorkflowId(&_AutomationReceiver.TransactOpts, _id)
}

func (_AutomationReceiver *AutomationReceiverTransactor) SetExpectedWorkflowName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setExpectedWorkflowName", _name)
}

func (_AutomationReceiver *AutomationReceiverSession) SetExpectedWorkflowName(_name string) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedWorkflowName(&_AutomationReceiver.TransactOpts, _name)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetExpectedWorkflowName(_name string) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetExpectedWorkflowName(&_AutomationReceiver.TransactOpts, _name)
}

func (_AutomationReceiver *AutomationReceiverTransactor) SetForwarderAddress(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setForwarderAddress", _forwarder)
}

func (_AutomationReceiver *AutomationReceiverSession) SetForwarderAddress(_forwarder common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetForwarderAddress(&_AutomationReceiver.TransactOpts, _forwarder)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetForwarderAddress(_forwarder common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetForwarderAddress(&_AutomationReceiver.TransactOpts, _forwarder)
}

func (_AutomationReceiver *AutomationReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_AutomationReceiver *AutomationReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.TransferOwnership(&_AutomationReceiver.TransactOpts, newOwner)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.TransferOwnership(&_AutomationReceiver.TransactOpts, newOwner)
}

type AutomationReceiverCallAllowedSetIterator struct {
	Event *AutomationReceiverCallAllowedSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverCallAllowedSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverCallAllowedSet)
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
		it.Event = new(AutomationReceiverCallAllowedSet)
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

func (it *AutomationReceiverCallAllowedSetIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverCallAllowedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverCallAllowedSet struct {
	Target   common.Address
	Selector [4]byte
	Allowed  bool
	Raw      types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterCallAllowedSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallAllowedSetIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "CallAllowedSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverCallAllowedSetIterator{contract: _AutomationReceiver.contract, event: "CallAllowedSet", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchCallAllowedSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallAllowedSet, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "CallAllowedSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverCallAllowedSet)
				if err := _AutomationReceiver.contract.UnpackLog(event, "CallAllowedSet", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseCallAllowedSet(log types.Log) (*AutomationReceiverCallAllowedSet, error) {
	event := new(AutomationReceiverCallAllowedSet)
	if err := _AutomationReceiver.contract.UnpackLog(event, "CallAllowedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverCallExecutedIterator struct {
	Event *AutomationReceiverCallExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverCallExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverCallExecuted)
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
		it.Event = new(AutomationReceiverCallExecuted)
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

func (it *AutomationReceiverCallExecutedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverCallExecuted struct {
	Target     common.Address
	Selector   [4]byte
	ReturnData []byte
	Raw        types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterCallExecuted(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "CallExecuted", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverCallExecutedIterator{contract: _AutomationReceiver.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallExecuted, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "CallExecuted", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverCallExecuted)
				if err := _AutomationReceiver.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseCallExecuted(log types.Log) (*AutomationReceiverCallExecuted, error) {
	event := new(AutomationReceiverCallExecuted)
	if err := _AutomationReceiver.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverCallFailedIterator struct {
	Event *AutomationReceiverCallFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverCallFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverCallFailed)
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
		it.Event = new(AutomationReceiverCallFailed)
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

func (it *AutomationReceiverCallFailedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverCallFailed struct {
	Target   common.Address
	Selector [4]byte
	Reason   []byte
	Raw      types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterCallFailed(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallFailedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "CallFailed", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverCallFailedIterator{contract: _AutomationReceiver.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallFailed, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "CallFailed", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverCallFailed)
				if err := _AutomationReceiver.contract.UnpackLog(event, "CallFailed", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseCallFailed(log types.Log) (*AutomationReceiverCallFailed, error) {
	event := new(AutomationReceiverCallFailed)
	if err := _AutomationReceiver.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverExpectedAuthorUpdatedIterator struct {
	Event *AutomationReceiverExpectedAuthorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverExpectedAuthorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverExpectedAuthorUpdated)
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
		it.Event = new(AutomationReceiverExpectedAuthorUpdated)
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

func (it *AutomationReceiverExpectedAuthorUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverExpectedAuthorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverExpectedAuthorUpdated struct {
	PreviousAuthor common.Address
	NewAuthor      common.Address
	Raw            types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterExpectedAuthorUpdated(opts *bind.FilterOpts, previousAuthor []common.Address, newAuthor []common.Address) (*AutomationReceiverExpectedAuthorUpdatedIterator, error) {

	var previousAuthorRule []interface{}
	for _, previousAuthorItem := range previousAuthor {
		previousAuthorRule = append(previousAuthorRule, previousAuthorItem)
	}
	var newAuthorRule []interface{}
	for _, newAuthorItem := range newAuthor {
		newAuthorRule = append(newAuthorRule, newAuthorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ExpectedAuthorUpdated", previousAuthorRule, newAuthorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverExpectedAuthorUpdatedIterator{contract: _AutomationReceiver.contract, event: "ExpectedAuthorUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchExpectedAuthorUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedAuthorUpdated, previousAuthor []common.Address, newAuthor []common.Address) (event.Subscription, error) {

	var previousAuthorRule []interface{}
	for _, previousAuthorItem := range previousAuthor {
		previousAuthorRule = append(previousAuthorRule, previousAuthorItem)
	}
	var newAuthorRule []interface{}
	for _, newAuthorItem := range newAuthor {
		newAuthorRule = append(newAuthorRule, newAuthorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ExpectedAuthorUpdated", previousAuthorRule, newAuthorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverExpectedAuthorUpdated)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedAuthorUpdated", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseExpectedAuthorUpdated(log types.Log) (*AutomationReceiverExpectedAuthorUpdated, error) {
	event := new(AutomationReceiverExpectedAuthorUpdated)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedAuthorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverExpectedWorkflowIdUpdatedIterator struct {
	Event *AutomationReceiverExpectedWorkflowIdUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverExpectedWorkflowIdUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverExpectedWorkflowIdUpdated)
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
		it.Event = new(AutomationReceiverExpectedWorkflowIdUpdated)
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

func (it *AutomationReceiverExpectedWorkflowIdUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverExpectedWorkflowIdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverExpectedWorkflowIdUpdated struct {
	PreviousId [32]byte
	NewId      [32]byte
	Raw        types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterExpectedWorkflowIdUpdated(opts *bind.FilterOpts, previousId [][32]byte, newId [][32]byte) (*AutomationReceiverExpectedWorkflowIdUpdatedIterator, error) {

	var previousIdRule []interface{}
	for _, previousIdItem := range previousId {
		previousIdRule = append(previousIdRule, previousIdItem)
	}
	var newIdRule []interface{}
	for _, newIdItem := range newId {
		newIdRule = append(newIdRule, newIdItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ExpectedWorkflowIdUpdated", previousIdRule, newIdRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverExpectedWorkflowIdUpdatedIterator{contract: _AutomationReceiver.contract, event: "ExpectedWorkflowIdUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchExpectedWorkflowIdUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedWorkflowIdUpdated, previousId [][32]byte, newId [][32]byte) (event.Subscription, error) {

	var previousIdRule []interface{}
	for _, previousIdItem := range previousId {
		previousIdRule = append(previousIdRule, previousIdItem)
	}
	var newIdRule []interface{}
	for _, newIdItem := range newId {
		newIdRule = append(newIdRule, newIdItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ExpectedWorkflowIdUpdated", previousIdRule, newIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverExpectedWorkflowIdUpdated)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedWorkflowIdUpdated", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseExpectedWorkflowIdUpdated(log types.Log) (*AutomationReceiverExpectedWorkflowIdUpdated, error) {
	event := new(AutomationReceiverExpectedWorkflowIdUpdated)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedWorkflowIdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverExpectedWorkflowNameUpdatedIterator struct {
	Event *AutomationReceiverExpectedWorkflowNameUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverExpectedWorkflowNameUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverExpectedWorkflowNameUpdated)
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
		it.Event = new(AutomationReceiverExpectedWorkflowNameUpdated)
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

func (it *AutomationReceiverExpectedWorkflowNameUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverExpectedWorkflowNameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverExpectedWorkflowNameUpdated struct {
	PreviousName [10]byte
	NewName      [10]byte
	Raw          types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterExpectedWorkflowNameUpdated(opts *bind.FilterOpts, previousName [][10]byte, newName [][10]byte) (*AutomationReceiverExpectedWorkflowNameUpdatedIterator, error) {

	var previousNameRule []interface{}
	for _, previousNameItem := range previousName {
		previousNameRule = append(previousNameRule, previousNameItem)
	}
	var newNameRule []interface{}
	for _, newNameItem := range newName {
		newNameRule = append(newNameRule, newNameItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ExpectedWorkflowNameUpdated", previousNameRule, newNameRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverExpectedWorkflowNameUpdatedIterator{contract: _AutomationReceiver.contract, event: "ExpectedWorkflowNameUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchExpectedWorkflowNameUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedWorkflowNameUpdated, previousName [][10]byte, newName [][10]byte) (event.Subscription, error) {

	var previousNameRule []interface{}
	for _, previousNameItem := range previousName {
		previousNameRule = append(previousNameRule, previousNameItem)
	}
	var newNameRule []interface{}
	for _, newNameItem := range newName {
		newNameRule = append(newNameRule, newNameItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ExpectedWorkflowNameUpdated", previousNameRule, newNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverExpectedWorkflowNameUpdated)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedWorkflowNameUpdated", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseExpectedWorkflowNameUpdated(log types.Log) (*AutomationReceiverExpectedWorkflowNameUpdated, error) {
	event := new(AutomationReceiverExpectedWorkflowNameUpdated)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ExpectedWorkflowNameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverForwarderAddressUpdatedIterator struct {
	Event *AutomationReceiverForwarderAddressUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverForwarderAddressUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverForwarderAddressUpdated)
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
		it.Event = new(AutomationReceiverForwarderAddressUpdated)
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

func (it *AutomationReceiverForwarderAddressUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverForwarderAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverForwarderAddressUpdated struct {
	PreviousForwarder common.Address
	NewForwarder      common.Address
	Raw               types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterForwarderAddressUpdated(opts *bind.FilterOpts, previousForwarder []common.Address, newForwarder []common.Address) (*AutomationReceiverForwarderAddressUpdatedIterator, error) {

	var previousForwarderRule []interface{}
	for _, previousForwarderItem := range previousForwarder {
		previousForwarderRule = append(previousForwarderRule, previousForwarderItem)
	}
	var newForwarderRule []interface{}
	for _, newForwarderItem := range newForwarder {
		newForwarderRule = append(newForwarderRule, newForwarderItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ForwarderAddressUpdated", previousForwarderRule, newForwarderRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverForwarderAddressUpdatedIterator{contract: _AutomationReceiver.contract, event: "ForwarderAddressUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchForwarderAddressUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverForwarderAddressUpdated, previousForwarder []common.Address, newForwarder []common.Address) (event.Subscription, error) {

	var previousForwarderRule []interface{}
	for _, previousForwarderItem := range previousForwarder {
		previousForwarderRule = append(previousForwarderRule, previousForwarderItem)
	}
	var newForwarderRule []interface{}
	for _, newForwarderItem := range newForwarder {
		newForwarderRule = append(newForwarderRule, newForwarderItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ForwarderAddressUpdated", previousForwarderRule, newForwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverForwarderAddressUpdated)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ForwarderAddressUpdated", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseForwarderAddressUpdated(log types.Log) (*AutomationReceiverForwarderAddressUpdated, error) {
	event := new(AutomationReceiverForwarderAddressUpdated)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ForwarderAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverOwnershipTransferredIterator struct {
	Event *AutomationReceiverOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverOwnershipTransferred)
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
		it.Event = new(AutomationReceiverOwnershipTransferred)
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

func (it *AutomationReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AutomationReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverOwnershipTransferredIterator{contract: _AutomationReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverOwnershipTransferred)
				if err := _AutomationReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*AutomationReceiverOwnershipTransferred, error) {
	event := new(AutomationReceiverOwnershipTransferred)
	if err := _AutomationReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverSecurityWarningIterator struct {
	Event *AutomationReceiverSecurityWarning

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverSecurityWarningIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverSecurityWarning)
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
		it.Event = new(AutomationReceiverSecurityWarning)
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

func (it *AutomationReceiverSecurityWarningIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverSecurityWarningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverSecurityWarning struct {
	Message string
	Raw     types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterSecurityWarning(opts *bind.FilterOpts) (*AutomationReceiverSecurityWarningIterator, error) {

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "SecurityWarning")
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverSecurityWarningIterator{contract: _AutomationReceiver.contract, event: "SecurityWarning", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchSecurityWarning(opts *bind.WatchOpts, sink chan<- *AutomationReceiverSecurityWarning) (event.Subscription, error) {

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "SecurityWarning")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverSecurityWarning)
				if err := _AutomationReceiver.contract.UnpackLog(event, "SecurityWarning", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseSecurityWarning(log types.Log) (*AutomationReceiverSecurityWarning, error) {
	event := new(AutomationReceiverSecurityWarning)
	if err := _AutomationReceiver.contract.UnpackLog(event, "SecurityWarning", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_AutomationReceiver *AutomationReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AutomationReceiver.abi.Events["CallAllowedSet"].ID:
		return _AutomationReceiver.ParseCallAllowedSet(log)
	case _AutomationReceiver.abi.Events["CallExecuted"].ID:
		return _AutomationReceiver.ParseCallExecuted(log)
	case _AutomationReceiver.abi.Events["CallFailed"].ID:
		return _AutomationReceiver.ParseCallFailed(log)
	case _AutomationReceiver.abi.Events["ExpectedAuthorUpdated"].ID:
		return _AutomationReceiver.ParseExpectedAuthorUpdated(log)
	case _AutomationReceiver.abi.Events["ExpectedWorkflowIdUpdated"].ID:
		return _AutomationReceiver.ParseExpectedWorkflowIdUpdated(log)
	case _AutomationReceiver.abi.Events["ExpectedWorkflowNameUpdated"].ID:
		return _AutomationReceiver.ParseExpectedWorkflowNameUpdated(log)
	case _AutomationReceiver.abi.Events["ForwarderAddressUpdated"].ID:
		return _AutomationReceiver.ParseForwarderAddressUpdated(log)
	case _AutomationReceiver.abi.Events["OwnershipTransferred"].ID:
		return _AutomationReceiver.ParseOwnershipTransferred(log)
	case _AutomationReceiver.abi.Events["SecurityWarning"].ID:
		return _AutomationReceiver.ParseSecurityWarning(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AutomationReceiverCallAllowedSet) Topic() common.Hash {
	return common.HexToHash("0x0925d576b7c865d78d7fe746ae46d080d64b9e6b04db5f034f71a79c41dda2e7")
}

func (AutomationReceiverCallExecuted) Topic() common.Hash {
	return common.HexToHash("0xbe82131bb3404498c769b0511da41a4ad409fa7152562c2b6669241cbe3bb884")
}

func (AutomationReceiverCallFailed) Topic() common.Hash {
	return common.HexToHash("0xefa88af289a36b936ccacf9bd9eaaa185775cd54ae263973d3579c01111593b6")
}

func (AutomationReceiverExpectedAuthorUpdated) Topic() common.Hash {
	return common.HexToHash("0x3321cda85c145617e47418aa14255e9dcbec53a753778e57591703b89a3cad31")
}

func (AutomationReceiverExpectedWorkflowIdUpdated) Topic() common.Hash {
	return common.HexToHash("0x0dbedcdf21925e053b4c574eae180d7f2883235ab4976ecc0873598a2a999b03")
}

func (AutomationReceiverExpectedWorkflowNameUpdated) Topic() common.Hash {
	return common.HexToHash("0x1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5")
}

func (AutomationReceiverForwarderAddressUpdated) Topic() common.Hash {
	return common.HexToHash("0x039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e")
}

func (AutomationReceiverOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AutomationReceiverSecurityWarning) Topic() common.Hash {
	return common.HexToHash("0x704da7db165c79c1e33d542c079333bbde970a733032d2f95fec8fb7d770cbf7")
}

func (_AutomationReceiver *AutomationReceiver) Address() common.Address {
	return _AutomationReceiver.address
}

type AutomationReceiverInterface interface {
	GetExpectedAuthor(opts *bind.CallOpts) (common.Address, error)

	GetExpectedWorkflowId(opts *bind.CallOpts) ([32]byte, error)

	GetExpectedWorkflowName(opts *bind.CallOpts) ([10]byte, error)

	GetForwarderAddress(opts *bind.CallOpts) (common.Address, error)

	IsCallAllowed(opts *bind.CallOpts, target common.Address, selector [4]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCallAllowed(opts *bind.TransactOpts, target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error)

	SetExpectedAuthor(opts *bind.TransactOpts, _author common.Address) (*types.Transaction, error)

	SetExpectedWorkflowId(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error)

	SetExpectedWorkflowName(opts *bind.TransactOpts, _name string) (*types.Transaction, error)

	SetForwarderAddress(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	FilterCallAllowedSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallAllowedSetIterator, error)

	WatchCallAllowedSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallAllowedSet, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallAllowedSet(log types.Log) (*AutomationReceiverCallAllowedSet, error)

	FilterCallExecuted(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallExecutedIterator, error)

	WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallExecuted, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallExecuted(log types.Log) (*AutomationReceiverCallExecuted, error)

	FilterCallFailed(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallFailedIterator, error)

	WatchCallFailed(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallFailed, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallFailed(log types.Log) (*AutomationReceiverCallFailed, error)

	FilterExpectedAuthorUpdated(opts *bind.FilterOpts, previousAuthor []common.Address, newAuthor []common.Address) (*AutomationReceiverExpectedAuthorUpdatedIterator, error)

	WatchExpectedAuthorUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedAuthorUpdated, previousAuthor []common.Address, newAuthor []common.Address) (event.Subscription, error)

	ParseExpectedAuthorUpdated(log types.Log) (*AutomationReceiverExpectedAuthorUpdated, error)

	FilterExpectedWorkflowIdUpdated(opts *bind.FilterOpts, previousId [][32]byte, newId [][32]byte) (*AutomationReceiverExpectedWorkflowIdUpdatedIterator, error)

	WatchExpectedWorkflowIdUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedWorkflowIdUpdated, previousId [][32]byte, newId [][32]byte) (event.Subscription, error)

	ParseExpectedWorkflowIdUpdated(log types.Log) (*AutomationReceiverExpectedWorkflowIdUpdated, error)

	FilterExpectedWorkflowNameUpdated(opts *bind.FilterOpts, previousName [][10]byte, newName [][10]byte) (*AutomationReceiverExpectedWorkflowNameUpdatedIterator, error)

	WatchExpectedWorkflowNameUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverExpectedWorkflowNameUpdated, previousName [][10]byte, newName [][10]byte) (event.Subscription, error)

	ParseExpectedWorkflowNameUpdated(log types.Log) (*AutomationReceiverExpectedWorkflowNameUpdated, error)

	FilterForwarderAddressUpdated(opts *bind.FilterOpts, previousForwarder []common.Address, newForwarder []common.Address) (*AutomationReceiverForwarderAddressUpdatedIterator, error)

	WatchForwarderAddressUpdated(opts *bind.WatchOpts, sink chan<- *AutomationReceiverForwarderAddressUpdated, previousForwarder []common.Address, newForwarder []common.Address) (event.Subscription, error)

	ParseForwarderAddressUpdated(log types.Log) (*AutomationReceiverForwarderAddressUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AutomationReceiverOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AutomationReceiverOwnershipTransferred, error)

	FilterSecurityWarning(opts *bind.FilterOpts) (*AutomationReceiverSecurityWarningIterator, error)

	WatchSecurityWarning(opts *bind.WatchOpts, sink chan<- *AutomationReceiverSecurityWarning) (event.Subscription, error)

	ParseSecurityWarning(log types.Log) (*AutomationReceiverSecurityWarning, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
