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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBlockNumberCheck\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lastReportBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsumerGasLimit\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedAuthor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedWorkflowName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getForwarderAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"retryable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retryableWhilePaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBlockNumberCheck\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"initialBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCallAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConsumerGasLimit\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedAuthor\",\"inputs\":[{\"name\":\"_author\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowId\",\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpectedWorkflowName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BlockNumberCheckSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"initialBlockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallAllowedSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsumerGasLimitSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"previousLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedAuthorUpdated\",\"inputs\":[{\"name\":\"previousAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowIdUpdated\",\"inputs\":[{\"name\":\"previousId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpectedWorkflowNameUpdated\",\"inputs\":[{\"name\":\"previousName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"},{\"name\":\"newName\",\"type\":\"bytes10\",\"indexed\":true,\"internalType\":\"bytes10\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwarderAddressUpdated\",\"inputs\":[{\"name\":\"previousForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newForwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportSkippedWhilePaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SecurityWarning\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleReportSkipped\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"reportBlockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lastReportBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallNotAllowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGas\",\"inputs\":[{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAuthor\",\"inputs\":[{\"name\":\"received\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidForwarderAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTargetAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWorkflowId\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowName\",\"inputs\":[{\"name\":\"received\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"},{\"name\":\"expected\",\"type\":\"bytes10\",\"internalType\":\"bytes10\"}]},{\"type\":\"error\",\"name\":\"MissingSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TargetHasNoCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WorkflowIdentityNotConfigured\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameRequiresAuthorValidation\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516124d03803806124d083398101604081905261002e9161012c565b80338061005457604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005d816100dd565b506001600160a01b0381166100845760405162e0775560e61b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040515f907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e908290a350506004805460ff19169055610159565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020828403121561013c575f80fd5b81516001600160a01b0381168114610152575f80fd5b9392505050565b61236a806101665f395ff3fe608060405234801561000f575f80fd5b5060043610610184575f3560e01c80638da5cb5b116100dd578063c3c44ac211610088578063d8fb859411610063578063d8fb859414610429578063f2fde38b14610453578063f5c793ef14610466575f80fd5b8063c3c44ac2146103f0578063d60c884b14610403578063d777cc6d14610416575f80fd5b8063b25de11e116100b8578063b25de11e146103bf578063bc1fc27a146103ca578063c2f7510d146103dd575f80fd5b80638da5cb5b146103405780639c1c77ca1461035d578063a619d81814610370575f80fd5b80635c975abb1161013d578063797c8d6911610118578063797c8d6914610250578063805f2132146102b9578063851ca9c1146102cc575f80fd5b80635c975abb1461022a578063715018a61461023557806375483caf1461023d575f80fd5b80633397cf671161016d5780633397cf67146101c55780633441856f146102045780633f4ba83a14610222575f80fd5b806301ffc9a71461018857806302329a29146101b0575b5f80fd5b61019b610196366004611def565b61046e565b60405190151581526020015b60405180910390f35b6101c36101be366004611e1e565b610506565b005b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b60015473ffffffffffffffffffffffffffffffffffffffff166101df565b6101c3610545565b60045460ff1661019b565b6101c3610557565b6101c361024b366004611e58565b610568565b61019b61025e366004611ea2565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526005602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290205460ff1692915050565b6101c36102c7366004611f13565b6106ec565b6103326102da366004611ea2565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290205492915050565b6040519081526020016101a7565b5f5473ffffffffffffffffffffffffffffffffffffffff166101df565b6101c361036b366004611f7f565b610aab565b60025474010000000000000000000000000000000000000000900460b01b6040517fffffffffffffffffffff0000000000000000000000000000000000000000000090911681526020016101a7565b60095460ff1661019b565b6101c36103d8366004611fc1565b610c2a565b6101c36103eb366004612000565b610e98565b6101c36103fe36600461203c565b610f85565b6101c3610411366004612053565b610fc5565b6101c3610424366004612053565b611043565b61043c610437366004611ea2565b611170565b6040805192151583526020830191909152016101a7565b6101c3610461366004612053565b6111e5565b600354610332565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f805f213200000000000000000000000000000000000000000000000000000000148061050057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b61050e611245565b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016821515179055610542611297565b50565b61054d611245565b61055561131c565b565b61055f611245565b6105555f611373565b610570611245565b73ffffffffffffffffffffffffffffffffffffffff84166105bd576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081526007602090815260408083207fffffffff0000000000000000000000000000000000000000000000000000000087168452909152812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001684151517905582610643575f610651565b811561064f5781610651565b435b73ffffffffffffffffffffffffffffffffffffffff86165f8181526008602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008a16808552908352928190208590558051881515815291820185905293945090927ff19a6052943cc4c32e1644d475a7b4e6f517bcae63159220028f5de68d6d0364910160405180910390a35050505050565b60015473ffffffffffffffffffffffffffffffffffffffff161580159061072b575060015473ffffffffffffffffffffffffffffffffffffffff163314155b15610789576001546040517fe1130dba00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b6003541515806107b0575060025473ffffffffffffffffffffffffffffffffffffffff1615155b806107f9575060025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff000000000000000000000000000000000000000000001615155b15610a9b575f805f61083f87878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506113e792505050565b60035492955090935091501580159061085a57506003548314155b1561089f576003546040517f9bfa39ba000000000000000000000000000000000000000000000000000000008152610780918591600401918252602082015260400190565b60025473ffffffffffffffffffffffffffffffffffffffff16158015906108e1575060025473ffffffffffffffffffffffffffffffffffffffff828116911614155b1561093c576002546040517fb8a98af800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301529091166024820152604401610780565b60025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff000000000000000000000000000000000000000000001615610a975760025473ffffffffffffffffffffffffffffffffffffffff166109d0576040517f4847901100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547fffffffffffffffffffff000000000000000000000000000000000000000000008381167401000000000000000000000000000000000000000090920460b01b1614610a97576002546040517f6c4609a60000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffff0000000000000000000000000000000000000000000084811660048301527401000000000000000000000000000000000000000090920460b01b9091166024820152604401610780565b5050505b610aa58282611400565b50505050565b610ab3611245565b73ffffffffffffffffffffffffffffffffffffffff8316610b00576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808015610b22575073ffffffffffffffffffffffffffffffffffffffff83163b155b15610b71576040517ffff2336100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610780565b73ffffffffffffffffffffffffffffffffffffffff83165f8181526005602090815260408083207fffffffff0000000000000000000000000000000000000000000000000000000087168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f0925d576b7c865d78d7fe746ae46d080d64b9e6b04db5f034f71a79c41dda2e7910160405180910390a3505050565b610c32611245565b60025474010000000000000000000000000000000000000000900460b01b5f829003610cd357600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff1690556040515f907fffffffffffffffffffff000000000000000000000000000000000000000000008316907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5908390a3505050565b5f60028484604051610ce692919061206e565b602060405180830381855afa158015610d01573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d24919061207d565b90505f610d5182604051602001610d3d91815260200190565b604051602081830303815290604052611af3565b60408051600a8082528183019092529192505f91906020820181803683370190505090505f5b600a811015610de557828181518110610d9257610d926120c1565b602001015160f81c60f81b828281518110610daf57610daf6120c1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610d77565b50610def816120ee565b600280547fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000060b093841c81029190911791829055604051910490911b7fffffffffffffffffffff0000000000000000000000000000000000000000000090811691908616907f1e7ddd09d504c82dcfc784a464b167469f5aad967606ec4822d848ef9141dfa5905f90a3505050505050565b610ea0611245565b73ffffffffffffffffffffffffffffffffffffffff8316610eed576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f8181526006602090815260408083207fffffffff00000000000000000000000000000000000000000000000000000000871680855290835292819020805490869055815181815292830186905293917ff3f16b36d6fb2a97e5e66d28e758a4d457e80197e1d455f692a96be32091fa3a910160405180910390a350505050565b610f8d611245565b6003805490829055604051829082907f0dbedcdf21925e053b4c574eae180d7f2883235ab4976ecc0873598a2a999b03905f90a35050565b610fcd611245565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f3321cda85c145617e47418aa14255e9dcbec53a753778e57591703b89a3cad31905f90a35050565b61104b611245565b60015473ffffffffffffffffffffffffffffffffffffffff9081169082166110fb577f704da7db165c79c1e33d542c079333bbde970a733032d2f95fec8fb7d770cbf76040516110f29060208082526038908201527f466f7277617264657220616464726573732073657420746f207a65726f202d2060408201527f636f6e7472616374206973206e6f7720494e5345435552450000000000000000606082015260800190565b60405180910390a15b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f039ad854736757070884dd787ef1a7f58db33546639d1f3efddcf4a33fb8997e905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff82165f8181526007602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008616808552908352818420549484526008835281842090845290915290205460ff909116905b9250929050565b6111ed611245565b73ffffffffffffffffffffffffffffffffffffffff811661123c576040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081525f6004820152602401610780565b61054281611373565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610555576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610780565b61129f611d42565b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586112f23390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b611324611d7f565b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336112f2565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60208101516040820151604a83015160601c9193909250565b60045460ff16156114745760095460ff1615611448576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517ff37c389ae9f5e13dbbef9f62b1ad0893e89c865ee834b2e221d11ebcf3cacd51905f90a15050565b60015473ffffffffffffffffffffffffffffffffffffffff166114c3576040517f381dd54000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354158015611530575060025473ffffffffffffffffffffffffffffffffffffffff161580611530575060025474010000000000000000000000000000000000000000900460b01b7fffffffffffffffffffff0000000000000000000000000000000000000000000016155b15611567576040517f8ec26f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80806115768486018661215a565b9194509250905073ffffffffffffffffffffffffffffffffffffffff83166115ca576040517ff1a492cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600481511015611606576040517f47d7741900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208181015173ffffffffffffffffffffffffffffffffffffffff85165f9081526005835260408082207fffffffff0000000000000000000000000000000000000000000000000000000084168352909352919091205460ff166116d6576040517f805043f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201527fffffffff0000000000000000000000000000000000000000000000000000000082166024820152604401610780565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526007602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290205460ff161561185e5773ffffffffffffffffffffffffffffffffffffffff84165f9081526008602090815260408083207fffffffff0000000000000000000000000000000000000000000000000000000085168452909152902054808410156118085760408051858152602081018390527fffffffff0000000000000000000000000000000000000000000000000000000084169173ffffffffffffffffffffffffffffffffffffffff8816917f96e22437f0b9d1dfe909221668f6ad9813c1128b60a7b4b3dee74e4248df834a910160405180910390a350505050505050565b5073ffffffffffffffffffffffffffffffffffffffff84165f9081526008602090815260408083207fffffffff000000000000000000000000000000000000000000000000000000008516845290915290208390555b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602090815260408083207fffffffff00000000000000000000000000000000000000000000000000000000851684529091528120549060608215611996575f611b586118c8603f86612292565b6118d290866122ca565b6118dc91906122ca565b9050805a1015611924575a6040517f23e228cb000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610780565b8773ffffffffffffffffffffffffffffffffffffffff16848760405161194a91906122dd565b5f604051808303815f8787f1925050503d805f8114611984576040519150601f19603f3d011682016040523d82523d5f602084013e611989565b606091505b509093509150611a009050565b8673ffffffffffffffffffffffffffffffffffffffff16856040516119bb91906122dd565b5f604051808303815f865af19150503d805f81146119f4576040519150601f19603f3d011682016040523d82523d5f602084013e6119f9565b606091505b5090925090505b8115611a7957837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168773ffffffffffffffffffffffffffffffffffffffff167fbe82131bb3404498c769b0511da41a4ad409fa7152562c2b6669241cbe3bb88483604051611a6c91906122f3565b60405180910390a3611ae8565b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168773ffffffffffffffffffffffffffffffffffffffff167fefa88af289a36b936ccacf9bd9eaaa185775cd54ae263973d3579c01111593b683604051611adf91906122f3565b60405180910390a35b505050505050505050565b60605f82516002611b049190612346565b67ffffffffffffffff811115611b1c57611b1c612094565b6040519080825280601f01601f191660200182016040528015611b46576020820181803683370190505b5090505f5b8351811015611d3b576040518060400160405280601081526020017f30313233343536373839616263646566000000000000000000000000000000008152506004858381518110611b9e57611b9e6120c1565b016020015182517fff0000000000000000000000000000000000000000000000000000000000000090911690911c60f81c908110611bde57611bde6120c1565b01602001517fff000000000000000000000000000000000000000000000000000000000000001682611c11836002612346565b81518110611c2157611c216120c1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053506040518060400160405280601081526020017f3031323334353637383961626364656600000000000000000000000000000000815250848281518110611c9757611c976120c1565b602091010151815160f89190911c600f16908110611cb757611cb76120c1565b01602001517fff000000000000000000000000000000000000000000000000000000000000001682611cea836002612346565b611cf59060016122ca565b81518110611d0557611d056120c1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101611b4b565b5092915050565b60045460ff1615610555576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460ff16610555576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114611dea575f80fd5b919050565b5f60208284031215611dff575f80fd5b611e0882611dbb565b9392505050565b80358015158114611dea575f80fd5b5f60208284031215611e2e575f80fd5b611e0882611e0f565b73ffffffffffffffffffffffffffffffffffffffff81168114610542575f80fd5b5f805f8060808587031215611e6b575f80fd5b8435611e7681611e37565b9350611e8460208601611dbb565b9250611e9260408601611e0f565b9396929550929360600135925050565b5f8060408385031215611eb3575f80fd5b8235611ebe81611e37565b9150611ecc60208401611dbb565b90509250929050565b5f8083601f840112611ee5575f80fd5b50813567ffffffffffffffff811115611efc575f80fd5b6020830191508360208285010111156111de575f80fd5b5f805f8060408587031215611f26575f80fd5b843567ffffffffffffffff811115611f3c575f80fd5b611f4887828801611ed5565b909550935050602085013567ffffffffffffffff811115611f67575f80fd5b611f7387828801611ed5565b95989497509550505050565b5f805f60608486031215611f91575f80fd5b8335611f9c81611e37565b9250611faa60208501611dbb565b9150611fb860408501611e0f565b90509250925092565b5f8060208385031215611fd2575f80fd5b823567ffffffffffffffff811115611fe8575f80fd5b611ff485828601611ed5565b90969095509350505050565b5f805f60608486031215612012575f80fd5b833561201d81611e37565b925061202b60208501611dbb565b929592945050506040919091013590565b5f6020828403121561204c575f80fd5b5035919050565b5f60208284031215612063575f80fd5b8135611e0881611e37565b818382375f9101908152919050565b5f6020828403121561208d575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b805160208201517fffffffffffffffffffff0000000000000000000000000000000000000000000081169190600a821015612153577fffffffffffffffffffff000000000000000000000000000000000000000000008083600a0360031b1b82161692505b5050919050565b5f805f6060848603121561216c575f80fd5b833561217781611e37565b925060208401359150604084013567ffffffffffffffff811115612199575f80fd5b8401601f810186136121a9575f80fd5b803567ffffffffffffffff8111156121c3576121c3612094565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561222f5761222f612094565b604052818152828201602001881015612246575f80fd5b816020840160208301375f602083830101528093505050509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f826122c5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8082018082111561050057610500612265565b5f82518060208501845e5f920191825250919050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b80820281158282048414176105005761050061226556fea164736f6c634300081a000a",
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

func (_AutomationReceiver *AutomationReceiverCaller) GetBlockNumberCheck(opts *bind.CallOpts, target common.Address, selector [4]byte) (GetBlockNumberCheck,

	error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getBlockNumberCheck", target, selector)

	outstruct := new(GetBlockNumberCheck)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.LastReportBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetBlockNumberCheck(target common.Address, selector [4]byte) (GetBlockNumberCheck,

	error) {
	return _AutomationReceiver.Contract.GetBlockNumberCheck(&_AutomationReceiver.CallOpts, target, selector)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetBlockNumberCheck(target common.Address, selector [4]byte) (GetBlockNumberCheck,

	error) {
	return _AutomationReceiver.Contract.GetBlockNumberCheck(&_AutomationReceiver.CallOpts, target, selector)
}

func (_AutomationReceiver *AutomationReceiverCaller) GetConsumerGasLimit(opts *bind.CallOpts, target common.Address, selector [4]byte) (*big.Int, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "getConsumerGasLimit", target, selector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) GetConsumerGasLimit(target common.Address, selector [4]byte) (*big.Int, error) {
	return _AutomationReceiver.Contract.GetConsumerGasLimit(&_AutomationReceiver.CallOpts, target, selector)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) GetConsumerGasLimit(target common.Address, selector [4]byte) (*big.Int, error) {
	return _AutomationReceiver.Contract.GetConsumerGasLimit(&_AutomationReceiver.CallOpts, target, selector)
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

func (_AutomationReceiver *AutomationReceiverCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) Paused() (bool, error) {
	return _AutomationReceiver.Contract.Paused(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) Paused() (bool, error) {
	return _AutomationReceiver.Contract.Paused(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCaller) RetryableWhilePaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AutomationReceiver.contract.Call(opts, &out, "retryableWhilePaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AutomationReceiver *AutomationReceiverSession) RetryableWhilePaused() (bool, error) {
	return _AutomationReceiver.Contract.RetryableWhilePaused(&_AutomationReceiver.CallOpts)
}

func (_AutomationReceiver *AutomationReceiverCallerSession) RetryableWhilePaused() (bool, error) {
	return _AutomationReceiver.Contract.RetryableWhilePaused(&_AutomationReceiver.CallOpts)
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

func (_AutomationReceiver *AutomationReceiverTransactor) Pause(opts *bind.TransactOpts, retryable bool) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "pause", retryable)
}

func (_AutomationReceiver *AutomationReceiverSession) Pause(retryable bool) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.Pause(&_AutomationReceiver.TransactOpts, retryable)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) Pause(retryable bool) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.Pause(&_AutomationReceiver.TransactOpts, retryable)
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

func (_AutomationReceiver *AutomationReceiverTransactor) SetBlockNumberCheck(opts *bind.TransactOpts, target common.Address, selector [4]byte, enabled bool, initialBlockNumber *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setBlockNumberCheck", target, selector, enabled, initialBlockNumber)
}

func (_AutomationReceiver *AutomationReceiverSession) SetBlockNumberCheck(target common.Address, selector [4]byte, enabled bool, initialBlockNumber *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetBlockNumberCheck(&_AutomationReceiver.TransactOpts, target, selector, enabled, initialBlockNumber)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetBlockNumberCheck(target common.Address, selector [4]byte, enabled bool, initialBlockNumber *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetBlockNumberCheck(&_AutomationReceiver.TransactOpts, target, selector, enabled, initialBlockNumber)
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

func (_AutomationReceiver *AutomationReceiverTransactor) SetConsumerGasLimit(opts *bind.TransactOpts, target common.Address, selector [4]byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "setConsumerGasLimit", target, selector, gasLimit)
}

func (_AutomationReceiver *AutomationReceiverSession) SetConsumerGasLimit(target common.Address, selector [4]byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetConsumerGasLimit(&_AutomationReceiver.TransactOpts, target, selector, gasLimit)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) SetConsumerGasLimit(target common.Address, selector [4]byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _AutomationReceiver.Contract.SetConsumerGasLimit(&_AutomationReceiver.TransactOpts, target, selector, gasLimit)
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

func (_AutomationReceiver *AutomationReceiverTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationReceiver.contract.Transact(opts, "unpause")
}

func (_AutomationReceiver *AutomationReceiverSession) Unpause() (*types.Transaction, error) {
	return _AutomationReceiver.Contract.Unpause(&_AutomationReceiver.TransactOpts)
}

func (_AutomationReceiver *AutomationReceiverTransactorSession) Unpause() (*types.Transaction, error) {
	return _AutomationReceiver.Contract.Unpause(&_AutomationReceiver.TransactOpts)
}

type AutomationReceiverBlockNumberCheckSetIterator struct {
	Event *AutomationReceiverBlockNumberCheckSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverBlockNumberCheckSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverBlockNumberCheckSet)
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
		it.Event = new(AutomationReceiverBlockNumberCheckSet)
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

func (it *AutomationReceiverBlockNumberCheckSetIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverBlockNumberCheckSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverBlockNumberCheckSet struct {
	Target             common.Address
	Selector           [4]byte
	Enabled            bool
	InitialBlockNumber *big.Int
	Raw                types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterBlockNumberCheckSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverBlockNumberCheckSetIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "BlockNumberCheckSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverBlockNumberCheckSetIterator{contract: _AutomationReceiver.contract, event: "BlockNumberCheckSet", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchBlockNumberCheckSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverBlockNumberCheckSet, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "BlockNumberCheckSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverBlockNumberCheckSet)
				if err := _AutomationReceiver.contract.UnpackLog(event, "BlockNumberCheckSet", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseBlockNumberCheckSet(log types.Log) (*AutomationReceiverBlockNumberCheckSet, error) {
	event := new(AutomationReceiverBlockNumberCheckSet)
	if err := _AutomationReceiver.contract.UnpackLog(event, "BlockNumberCheckSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

type AutomationReceiverConsumerGasLimitSetIterator struct {
	Event *AutomationReceiverConsumerGasLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverConsumerGasLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverConsumerGasLimitSet)
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
		it.Event = new(AutomationReceiverConsumerGasLimitSet)
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

func (it *AutomationReceiverConsumerGasLimitSetIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverConsumerGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverConsumerGasLimitSet struct {
	Target        common.Address
	Selector      [4]byte
	PreviousLimit *big.Int
	NewLimit      *big.Int
	Raw           types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterConsumerGasLimitSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverConsumerGasLimitSetIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ConsumerGasLimitSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverConsumerGasLimitSetIterator{contract: _AutomationReceiver.contract, event: "ConsumerGasLimitSet", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchConsumerGasLimitSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverConsumerGasLimitSet, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ConsumerGasLimitSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverConsumerGasLimitSet)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ConsumerGasLimitSet", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseConsumerGasLimitSet(log types.Log) (*AutomationReceiverConsumerGasLimitSet, error) {
	event := new(AutomationReceiverConsumerGasLimitSet)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ConsumerGasLimitSet", log); err != nil {
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

type AutomationReceiverPausedIterator struct {
	Event *AutomationReceiverPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverPaused)
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
		it.Event = new(AutomationReceiverPaused)
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

func (it *AutomationReceiverPausedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterPaused(opts *bind.FilterOpts) (*AutomationReceiverPausedIterator, error) {

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverPausedIterator{contract: _AutomationReceiver.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverPaused) (event.Subscription, error) {

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverPaused)
				if err := _AutomationReceiver.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParsePaused(log types.Log) (*AutomationReceiverPaused, error) {
	event := new(AutomationReceiverPaused)
	if err := _AutomationReceiver.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverReportSkippedWhilePausedIterator struct {
	Event *AutomationReceiverReportSkippedWhilePaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverReportSkippedWhilePausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverReportSkippedWhilePaused)
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
		it.Event = new(AutomationReceiverReportSkippedWhilePaused)
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

func (it *AutomationReceiverReportSkippedWhilePausedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverReportSkippedWhilePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverReportSkippedWhilePaused struct {
	Raw types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterReportSkippedWhilePaused(opts *bind.FilterOpts) (*AutomationReceiverReportSkippedWhilePausedIterator, error) {

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "ReportSkippedWhilePaused")
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverReportSkippedWhilePausedIterator{contract: _AutomationReceiver.contract, event: "ReportSkippedWhilePaused", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchReportSkippedWhilePaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverReportSkippedWhilePaused) (event.Subscription, error) {

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "ReportSkippedWhilePaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverReportSkippedWhilePaused)
				if err := _AutomationReceiver.contract.UnpackLog(event, "ReportSkippedWhilePaused", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseReportSkippedWhilePaused(log types.Log) (*AutomationReceiverReportSkippedWhilePaused, error) {
	event := new(AutomationReceiverReportSkippedWhilePaused)
	if err := _AutomationReceiver.contract.UnpackLog(event, "ReportSkippedWhilePaused", log); err != nil {
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

type AutomationReceiverStaleReportSkippedIterator struct {
	Event *AutomationReceiverStaleReportSkipped

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverStaleReportSkippedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverStaleReportSkipped)
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
		it.Event = new(AutomationReceiverStaleReportSkipped)
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

func (it *AutomationReceiverStaleReportSkippedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverStaleReportSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverStaleReportSkipped struct {
	Target            common.Address
	Selector          [4]byte
	ReportBlockNumber *big.Int
	LastReportBlock   *big.Int
	Raw               types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterStaleReportSkipped(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverStaleReportSkippedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "StaleReportSkipped", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverStaleReportSkippedIterator{contract: _AutomationReceiver.contract, event: "StaleReportSkipped", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchStaleReportSkipped(opts *bind.WatchOpts, sink chan<- *AutomationReceiverStaleReportSkipped, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "StaleReportSkipped", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverStaleReportSkipped)
				if err := _AutomationReceiver.contract.UnpackLog(event, "StaleReportSkipped", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseStaleReportSkipped(log types.Log) (*AutomationReceiverStaleReportSkipped, error) {
	event := new(AutomationReceiverStaleReportSkipped)
	if err := _AutomationReceiver.contract.UnpackLog(event, "StaleReportSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationReceiverUnpausedIterator struct {
	Event *AutomationReceiverUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationReceiverUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationReceiverUnpaused)
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
		it.Event = new(AutomationReceiverUnpaused)
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

func (it *AutomationReceiverUnpausedIterator) Error() error {
	return it.fail
}

func (it *AutomationReceiverUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationReceiverUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_AutomationReceiver *AutomationReceiverFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AutomationReceiverUnpausedIterator, error) {

	logs, sub, err := _AutomationReceiver.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AutomationReceiverUnpausedIterator{contract: _AutomationReceiver.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_AutomationReceiver *AutomationReceiverFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverUnpaused) (event.Subscription, error) {

	logs, sub, err := _AutomationReceiver.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationReceiverUnpaused)
				if err := _AutomationReceiver.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_AutomationReceiver *AutomationReceiverFilterer) ParseUnpaused(log types.Log) (*AutomationReceiverUnpaused, error) {
	event := new(AutomationReceiverUnpaused)
	if err := _AutomationReceiver.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetBlockNumberCheck struct {
	Enabled         bool
	LastReportBlock *big.Int
}

func (_AutomationReceiver *AutomationReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AutomationReceiver.abi.Events["BlockNumberCheckSet"].ID:
		return _AutomationReceiver.ParseBlockNumberCheckSet(log)
	case _AutomationReceiver.abi.Events["CallAllowedSet"].ID:
		return _AutomationReceiver.ParseCallAllowedSet(log)
	case _AutomationReceiver.abi.Events["CallExecuted"].ID:
		return _AutomationReceiver.ParseCallExecuted(log)
	case _AutomationReceiver.abi.Events["CallFailed"].ID:
		return _AutomationReceiver.ParseCallFailed(log)
	case _AutomationReceiver.abi.Events["ConsumerGasLimitSet"].ID:
		return _AutomationReceiver.ParseConsumerGasLimitSet(log)
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
	case _AutomationReceiver.abi.Events["Paused"].ID:
		return _AutomationReceiver.ParsePaused(log)
	case _AutomationReceiver.abi.Events["ReportSkippedWhilePaused"].ID:
		return _AutomationReceiver.ParseReportSkippedWhilePaused(log)
	case _AutomationReceiver.abi.Events["SecurityWarning"].ID:
		return _AutomationReceiver.ParseSecurityWarning(log)
	case _AutomationReceiver.abi.Events["StaleReportSkipped"].ID:
		return _AutomationReceiver.ParseStaleReportSkipped(log)
	case _AutomationReceiver.abi.Events["Unpaused"].ID:
		return _AutomationReceiver.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AutomationReceiverBlockNumberCheckSet) Topic() common.Hash {
	return common.HexToHash("0xf19a6052943cc4c32e1644d475a7b4e6f517bcae63159220028f5de68d6d0364")
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

func (AutomationReceiverConsumerGasLimitSet) Topic() common.Hash {
	return common.HexToHash("0xf3f16b36d6fb2a97e5e66d28e758a4d457e80197e1d455f692a96be32091fa3a")
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

func (AutomationReceiverPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (AutomationReceiverReportSkippedWhilePaused) Topic() common.Hash {
	return common.HexToHash("0xf37c389ae9f5e13dbbef9f62b1ad0893e89c865ee834b2e221d11ebcf3cacd51")
}

func (AutomationReceiverSecurityWarning) Topic() common.Hash {
	return common.HexToHash("0x704da7db165c79c1e33d542c079333bbde970a733032d2f95fec8fb7d770cbf7")
}

func (AutomationReceiverStaleReportSkipped) Topic() common.Hash {
	return common.HexToHash("0x96e22437f0b9d1dfe909221668f6ad9813c1128b60a7b4b3dee74e4248df834a")
}

func (AutomationReceiverUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_AutomationReceiver *AutomationReceiver) Address() common.Address {
	return _AutomationReceiver.address
}

type AutomationReceiverInterface interface {
	GetBlockNumberCheck(opts *bind.CallOpts, target common.Address, selector [4]byte) (GetBlockNumberCheck,

		error)

	GetConsumerGasLimit(opts *bind.CallOpts, target common.Address, selector [4]byte) (*big.Int, error)

	GetExpectedAuthor(opts *bind.CallOpts) (common.Address, error)

	GetExpectedWorkflowId(opts *bind.CallOpts) ([32]byte, error)

	GetExpectedWorkflowName(opts *bind.CallOpts) ([10]byte, error)

	GetForwarderAddress(opts *bind.CallOpts) (common.Address, error)

	IsCallAllowed(opts *bind.CallOpts, target common.Address, selector [4]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	RetryableWhilePaused(opts *bind.CallOpts) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, retryable bool) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetBlockNumberCheck(opts *bind.TransactOpts, target common.Address, selector [4]byte, enabled bool, initialBlockNumber *big.Int) (*types.Transaction, error)

	SetCallAllowed(opts *bind.TransactOpts, target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error)

	SetConsumerGasLimit(opts *bind.TransactOpts, target common.Address, selector [4]byte, gasLimit *big.Int) (*types.Transaction, error)

	SetExpectedAuthor(opts *bind.TransactOpts, _author common.Address) (*types.Transaction, error)

	SetExpectedWorkflowId(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error)

	SetExpectedWorkflowName(opts *bind.TransactOpts, _name string) (*types.Transaction, error)

	SetForwarderAddress(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBlockNumberCheckSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverBlockNumberCheckSetIterator, error)

	WatchBlockNumberCheckSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverBlockNumberCheckSet, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseBlockNumberCheckSet(log types.Log) (*AutomationReceiverBlockNumberCheckSet, error)

	FilterCallAllowedSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallAllowedSetIterator, error)

	WatchCallAllowedSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallAllowedSet, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallAllowedSet(log types.Log) (*AutomationReceiverCallAllowedSet, error)

	FilterCallExecuted(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallExecutedIterator, error)

	WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallExecuted, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallExecuted(log types.Log) (*AutomationReceiverCallExecuted, error)

	FilterCallFailed(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverCallFailedIterator, error)

	WatchCallFailed(opts *bind.WatchOpts, sink chan<- *AutomationReceiverCallFailed, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseCallFailed(log types.Log) (*AutomationReceiverCallFailed, error)

	FilterConsumerGasLimitSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverConsumerGasLimitSetIterator, error)

	WatchConsumerGasLimitSet(opts *bind.WatchOpts, sink chan<- *AutomationReceiverConsumerGasLimitSet, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseConsumerGasLimitSet(log types.Log) (*AutomationReceiverConsumerGasLimitSet, error)

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

	FilterPaused(opts *bind.FilterOpts) (*AutomationReceiverPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*AutomationReceiverPaused, error)

	FilterReportSkippedWhilePaused(opts *bind.FilterOpts) (*AutomationReceiverReportSkippedWhilePausedIterator, error)

	WatchReportSkippedWhilePaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverReportSkippedWhilePaused) (event.Subscription, error)

	ParseReportSkippedWhilePaused(log types.Log) (*AutomationReceiverReportSkippedWhilePaused, error)

	FilterSecurityWarning(opts *bind.FilterOpts) (*AutomationReceiverSecurityWarningIterator, error)

	WatchSecurityWarning(opts *bind.WatchOpts, sink chan<- *AutomationReceiverSecurityWarning) (event.Subscription, error)

	ParseSecurityWarning(log types.Log) (*AutomationReceiverSecurityWarning, error)

	FilterStaleReportSkipped(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*AutomationReceiverStaleReportSkippedIterator, error)

	WatchStaleReportSkipped(opts *bind.WatchOpts, sink chan<- *AutomationReceiverStaleReportSkipped, target []common.Address, selector [][4]byte) (event.Subscription, error)

	ParseStaleReportSkipped(log types.Log) (*AutomationReceiverStaleReportSkipped, error)

	FilterUnpaused(opts *bind.FilterOpts) (*AutomationReceiverUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationReceiverUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*AutomationReceiverUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
