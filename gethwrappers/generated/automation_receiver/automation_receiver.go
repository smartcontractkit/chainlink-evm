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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getExpectedAuthor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getForwarderAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedAuthor\",\"inputs\":[{\"name\":\"_author\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowId\",\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CallAllowedSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedAuthorUpdated\",\"inputs\":[{\"name\":\"previousAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowIdUpdated\",\"inputs\":[{\"name\":\"previousId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowNameUpdated\",\"inputs\":[{\"name\":\"previousName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"},{\"name\":\"newName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwarderAddressUpdated\",\"inputs\":[{\"name\":\"previousForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityWarning\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallNotAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidAuthor\",\"inputs\":[{\"name\":\"received\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidForwarderAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTargetAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWorkflowId\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowName\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"expected\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}]},{\"type\":\"error\",\"name\":\"MissingSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequiresAuthorValidation\",\"inputs\":[]}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200198a3803806200198a83398101604081905262000033916200012c565b8033806200005a57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6200006581620000dd565b506001600160a01b0381166200008d5760405162e0775560e61b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040515f907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e908290a350506200015b565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156200013d575f80fd5b81516001600160a01b038116811462000154575f80fd5b9392505050565b61182180620001695f395ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c80639c1c77ca11610093578063d60c884b11610063578063d60c884b146102f8578063d777cc6d1461030b578063f2fde38b1461031e578063f5c793ef14610331575f80fd5b80639c1c77ca14610270578063a619d81814610283578063bc1fc27a146102d2578063c3c44ac2146102e5575f80fd5b8063715018a6116100ce578063715018a6146101cd578063797c8d69146101d7578063805f2132146102405780638da5cb5b14610253575f80fd5b806301ffc9a7146100ff578063181f5a77146101275780633397cf67146101705780633441856f146101af575b5f80fd5b61011261010d36600461139f565b610342565b60405190151581526020015b60405180910390f35b6101636040518060400160405280601881526020017f4175746f6d6174696f6e526563656976657220312e302e30000000000000000081525081565b60405161011e919061142a565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161011e565b60015473ffffffffffffffffffffffffffffffffffffffff1661018a565b6101d56103da565b005b6101126101e536600461145d565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526004602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290205460ff1692915050565b6101d561024e3660046114d5565b6103ed565b5f5473ffffffffffffffffffffffffffffffffffffffff1661018a565b6101d561027e36600461153c565b6107ac565b60025474010000000000000000000000000000000000000000900460b01b6040517fffffffffffffffffffff00000000000000000000000000000000000000000000909116815260200161011e565b6101d56102e0366004611586565b6108ba565b6101d56102f33660046115c5565b610b28565b6101d56103063660046115dc565b610b68565b6101d56103193660046115dc565b610be6565b6101d561032c3660046115dc565b610d13565b60035460405190815260200161011e565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f805f21320000000000000000000000000000000000000000000000000000000014806103d457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6103e2610d76565b6103eb5f610dc8565b565b60015473ffffffffffffffffffffffffffffffffffffffff161580159061042c575060015473ffffffffffffffffffffffffffffffffffffffff163314155b1561048a576001546040517fe1130dba00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b6003541515806104b1575060025473ffffffffffffffffffffffffffffffffffffffff1615155b806104fa575060025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff000000000000000000000000000000000000000000001615155b1561079c575f805f61054087878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610e3c92505050565b60035492955090935091501580159061055b57506003548314155b156105a0576003546040517f9bfa39ba000000000000000000000000000000000000000000000000000000008152610481918591600401918252602082015260400190565b60025473ffffffffffffffffffffffffffffffffffffffff16158015906105e2575060025473ffffffffffffffffffffffffffffffffffffffff828116911614155b1561063d576002546040517fb8a98af800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301529091166024820152604401610481565b60025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff0000000000000000000000000000000000000000000016156107985760025473ffffffffffffffffffffffffffffffffffffffff166106d1576040517f4847901100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547fffffffffffffffffffff000000000000000000000000000000000000000000008381167401000000000000000000000000000000000000000090920460b01b1614610798576002546040517f6c4609a60000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffff0000000000000000000000000000000000000000000084811660048301527401000000000000000000000000000000000000000090920460b01b9091166024820152604401610481565b5050505b6107a68282610e55565b50505050565b6107b4610d76565b73ffffffffffffffffffffffffffffffffffffffff8316610801576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f8181526004602090815260408083207fffffffff0000000000000000000000000000000000000000000000000000000087168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f0925d576b7c865d78d7fe746ae46d080d64b9e6b04db5f034f71a79c41dda2e7910160405180910390a3505050565b6108c2610d76565b60025474010000000000000000000000000000000000000000900460b01b5f82900361096357600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff1690556040515f907fffffffffffffffffffff000000000000000000000000000000000000000000008316907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5908390a3505050565b5f600284846040516109769291906115f7565b602060405180830381855afa158015610991573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906109b49190611606565b90505f6109e1826040516020016109cd91815260200190565b60405160208183030381529060405261111c565b60408051600a8082528183019092529192505f91906020820181803683370190505090505f5b600a811015610a7557828181518110610a2257610a2261164a565b602001015160f81c60f81b828281518110610a3f57610a3f61164a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610a07565b50610a7f81611677565b600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060b093841c81029190911791829055604051910490911b7fffffffffffffffffffff0000000000000000000000000000000000000000000090811691908616907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5905f90a3505050505050565b610b30610d76565b6003805490829055604051829082907f0dbedcdf21925e053b4c574eae180d7f2883235ab4976ecc0873598a2a999b03905f90a35050565b610b70610d76565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f3321cda85c145617e47418aa14255e9dcbec53a753778e57591703b89a3cad31905f90a35050565b610bee610d76565b60015473ffffffffffffffffffffffffffffffffffffffff908116908216610c9e577f704da7db165c79c1e33d542c079333bbde970a733032d2f95fec8fb7d770cbf7604051610c959060208082526038908201527f466f7277617264657220616464726573732073657420746f207a65726f202d2060408201527f636f6e7472616374206973206e6f7720494e5345435552450000000000000000606082015260800190565b60405180910390a15b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e905f90a35050565b610d1b610d76565b73ffffffffffffffffffffffffffffffffffffffff8116610d6a576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610481565b610d7381610dc8565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146103eb576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610481565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60208101516040820151604a83015160601c9193909250565b5f80610e63838501856116c6565b909250905073ffffffffffffffffffffffffffffffffffffffff8216610eb5576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600481511015610ef1576040517f47d7741900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208181015173ffffffffffffffffffffffffffffffffffffffff84165f9081526004835260408082207fffffffff0000000000000000000000000000000000000000000000000000000084168352909352919091205460ff16610fc1576040517f805043f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201527fffffffff0000000000000000000000000000000000000000000000000000000082166024820152604401610481565b5f808473ffffffffffffffffffffffffffffffffffffffff1684604051610fe891906117a2565b5f604051808303815f865af19150503d805f8114611021576040519150601f19603f3d011682016040523d82523d5f602084013e611026565b606091505b509150915081156110a457827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168573ffffffffffffffffffffffffffffffffffffffff167fbe82131bb3404498c769b0511da41a4ad409fa7152562c2b6669241cbe3bb88483604051611097919061142a565b60405180910390a3611113565b827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168573ffffffffffffffffffffffffffffffffffffffff167fefa88af289a36b936ccacf9bd9eaaa185775cd54ae263973d3579c01111593b68360405161110a919061142a565b60405180910390a35b50505050505050565b60605f8251600261112d91906117ea565b67ffffffffffffffff8111156111455761114561161d565b6040519080825280601f01601f19166020018201604052801561116f576020820181803683370190505b5090505f5b8351811015611364576040518060400160405280601081526020017f303132333435363738396162636465660000000000000000000000000000000081525060048583815181106111c7576111c761164a565b016020015182517fff0000000000000000000000000000000000000000000000000000000000000090911690911c60f81c9081106112075761120761164a565b01602001517fff00000000000000000000000000000000000000000000000000000000000000168261123a8360026117ea565b8151811061124a5761124a61164a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053506040518060400160405280601081526020017f30313233343536373839616263646566000000000000000000000000000000008152508482815181106112c0576112c061164a565b602091010151815160f89190911c600f169081106112e0576112e061164a565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016826113138360026117ea565b61131e906001611801565b8151811061132e5761132e61164a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101611174565b5092915050565b80357fffffffff000000000000000000000000000000000000000000000000000000008116811461139a575f80fd5b919050565b5f602082840312156113af575f80fd5b6113b88261136b565b9392505050565b5f5b838110156113d95781810151838201526020016113c1565b50505f910152565b5f81518084526113f88160208601602086016113bf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6113b860208301846113e1565b73ffffffffffffffffffffffffffffffffffffffff81168114610d73575f80fd5b5f806040838503121561146e575f80fd5b82356114798161143c565b91506114876020840161136b565b90509250929050565b5f8083601f8401126114a0575f80fd5b50813567ffffffffffffffff8111156114b7575f80fd5b6020830191508360208285010111156114ce575f80fd5b9250929050565b5f805f80604085870312156114e8575f80fd5b843567ffffffffffffffff808211156114ff575f80fd5b61150b88838901611490565b90965094506020870135915080821115611523575f80fd5b5061153087828801611490565b95989497509550505050565b5f805f6060848603121561154e575f80fd5b83356115598161143c565b92506115676020850161136b565b91506040840135801515811461157b575f80fd5b809150509250925092565b5f8060208385031215611597575f80fd5b823567ffffffffffffffff8111156115ad575f80fd5b6115b985828601611490565b90969095509350505050565b5f602082840312156115d5575f80fd5b5035919050565b5f602082840312156115ec575f80fd5b81356113b88161143c565b818382375f9101908152919050565b5f60208284031215611616575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f815160208301517fffffffffffffffffffff000000000000000000000000000000000000000000008082169350600a8310156116be57808184600a0360031b1b83161693505b505050919050565b5f80604083850312156116d7575f80fd5b82356116e28161143c565b9150602083013567ffffffffffffffff808211156116fe575f80fd5b818501915085601f830112611711575f80fd5b8135818111156117235761172361161d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156117695761176961161d565b81604052828152886020848701011115611781575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f82516117b38184602087016113bf565b9190910192915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176103d4576103d46117bd565b808201808211156103d4576103d46117bd56fea164736f6c6343000818000a",
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

func (_AutomationReceiver *AutomationReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) TypeAndVersion() (string, error) {
	return _AutomationReceiver.Contract.TypeAndVersion(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) TypeAndVersion() (string, error) {
	return _AutomationReceiver.Contract.TypeAndVersion(&_AutomationReceiver.CallOpts)
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

	TypeAndVersion(opts *bind.CallOpts) (string, error)

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
