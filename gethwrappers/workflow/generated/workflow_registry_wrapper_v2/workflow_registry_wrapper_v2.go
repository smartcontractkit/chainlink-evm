// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package workflow_registry_wrapper_v2

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

type WorkflowRegistryWorkflowMetadata struct {
	WorkflowID   [32]byte
	Owner        common.Address
	CreatedAt    uint64
	Status       uint8
	DonLabel     [32]byte
	WorkflowName string
	BinaryURL    string
	ConfigURL    string
	Tag          string
	Attributes   []byte
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIDs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"batchSizeOverride\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIDs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"batchSizeOverride\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"maxActive\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadata\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"meta\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata\",\"components\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowMetadataListByOwnerName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadata[]\",\"components\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"binaryURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BatchTooLarge\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUnlinkWithActiveWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWorkflowID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WorkflowAlreadyInDesiredStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b031916331790556148b590816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063034a6ed5146128335780630987294c146127c457806317e0edfc1461272c578063181f5a77146126de57806323f381bb1461268b5780633c17181b1461261157806347d1ed83146125e057806352453a44146125745780635af1ee771461221e578063695e1340146120fc5780636ee80b4414611f4b5780636f35177114611e9e57806379ba509714611db55780637e832c3514611ce5578063809b23cc14611c955780638109676714611b965780638da5cb5b14611b445780639c5b230414611b015780639f2b2cbd1461084a578063a0b8a4fe1461080e578063b87a0194146107b4578063cabb9e7a1461074a578063d8e4a72414610580578063dc101969146104ec578063dfcb0b31146104d2578063e690f3321461047e578063eec860b7146103a1578063f23e37ff146102505763f2fde38b1461015b57600080fd5b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5773ffffffffffffffffffffffffffffffffffffffff6101a7612af5565b6101af61390a565b1633811461022157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461024b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b57610287612af5565b602435906044359163ffffffff831680930361024b57606435801515810361024b576102b161390a565b1561036957604051926102c384612b35565b835273ffffffffffffffffffffffffffffffffffffffff60208401926001845216600052601160205260406000209060005260205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b169116179055600080f35b915073ffffffffffffffffffffffffffffffffffffffff16600052601160205260406000209060005260205260006040812055600080f35b3461024b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576103d8612af5565b602435906103e461390a565b73ffffffffffffffffffffffffffffffffffffffff6012541682602073ffffffffffffffffffffffffffffffffffffffff6040519461042286612b35565b16938481520152817fffffffffffffffffffffffff00000000000000000000000000000000000000006012541617601255826013557fc60b6bc6904a32fe864ae9ad697fd792f40e1d9e9fee5da0a10d0fb26b7380d6600080a4005b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576104b633613da3565b60043560005260096020526104d060406000205433613fde565b005b3461024b576104d06104e336612d84565b9291909161367c565b3461024b5760016105086104ff36612d84565b9180949361367c565b3360005260056020528060406000205561052133614751565b508060005260066020526040600020827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b3461024b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5760043567ffffffffffffffff811161024b576105cf903690600401612cf8565b6024359182151580930361024b576105e561390a565b60ff831660005b83811061069057505060405191806040840160408552526060830191906000905b808210610643577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff8216820361024b576020809173ffffffffffffffffffffffffffffffffffffffff60019416815201940192019061060d565b73ffffffffffffffffffffffffffffffffffffffff6106b86106b3838787613429565b6135e3565b1615610720578073ffffffffffffffffffffffffffffffffffffffff6106e46106b36001948888613429565b1660005260026020526040600020837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055016105ec565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5773ffffffffffffffffffffffffffffffffffffffff610796612af5565b166000526002602052602060ff604060002054166040519015158152f35b3461024b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5761080a6107fe6107f1612af5565b6044359060243590613513565b60405191829182612c7a565b0390f35b3461024b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576020600354604051908152f35b3461024b576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5760043567ffffffffffffffff811161024b5761089a903690600401612bcd565b60243567ffffffffffffffff811161024b576108ba903690600401612bcd565b906002606435101561024b5760a43567ffffffffffffffff811161024b576108e6903690600401612bcd565b909160c43567ffffffffffffffff811161024b57610908903690600401612bcd565b94909160e43567ffffffffffffffff811161024b5761092b903690600401612bcd565b929093610104359788159889150361024b57610948368b8d61315f565b6109ab6109d761095936878961315f565b60405192839161097b60208401963388526060604086015260808501906129c9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160608501526129c9565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612b51565b5190209788600052600760205273ffffffffffffffffffffffffffffffffffffffff600160406000200154161560001461167b57610a1433613da3565b8a156116515760408b1161161e5783156115f457602084116115c157610a3b604435614327565b610a45828961439a565b610a4e866143df565b610a62610a5c368d8f61315f565b33613d62565b99610a6e6064356129bf565b606435156115b0575b611385575b60405197610a8989612b18565b60443589523360208a015267ffffffffffffffff421660408a0152610aaf6064356129bf565b60643560608a015260843560808a0152610aca368d8f61315f565b60a08a01523690610ada9261315f565b60c08801523690610aea9261315f565b60e08601523690610afa9261315f565b6101008401523690610b0b9261315f565b61012082015281600052600760205260406000208151815560018101602083015173ffffffffffffffffffffffffffffffffffffffff1681547fffffffffffffffffffffffff000000000000000000000000000000000000000016178155604083015181549060a01b7bffffffffffffffff000000000000000000000000000000000000000016907fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff161781556060830151610bc6816129bf565b610bcf816129bf565b81549060e01b7cff0000000000000000000000000000000000000000000000000000000016907fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff16179055608082015160028201556003810160a083015180519067ffffffffffffffff82116110ad57610c5382610c4d8554612e28565b85614024565b602090601f83116001146112e257610ca1929160009183611191575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6004810160c083015180519067ffffffffffffffff82116110ad57610ccd82610c4d8554612e28565b602090601f831160011461123f57610d1a9291600091836111915750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6005810160e083015180519067ffffffffffffffff82116110ad57610d4682610c4d8554612e28565b602090601f831160011461119c57610d939291600091836111915750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6006810161010083015180519067ffffffffffffffff82116110ad57610dc082610c4d8554612e28565b602090601f83116001146110e75782610120959360079593610e15936000926110dc5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b0191015180519067ffffffffffffffff82116110ad57610e3c82610c4d8554612e28565b602090601f831160011461100a57610e89929160009183610fff5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b816000526008602052610ea38160406000206147b1565b50604435600052600960205280604060002055608435600052600c602052610ecf8160406000206147b1565b5033600052600d602052610ee78160406000206147b1565b50610ef36064356129bf565b60643515610f5a575b50507f03bdc9e1a7b65f76fa1e989e76c28e6ac9f556262cfed31a11a9461e77ad2ff56040516084358152610f326064356129bf565b60643560208201526060604082015280610f553395604435956060840191613604565b0390a3005b610ff791608435600052600b602052610f778260406000206147b1565b5033600052600a602052610f8f8260406000206147b1565b5033600052600f6020526040600020608435600052602052604060002063ffffffff610fbd81835416613439565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e60205260406000206147b1565b508280610efc565b015190508880610c6f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611095575090846001959493921061105e575b505050811b019055610e8c565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055878080611051565b9293602060018192878601518155019501930161103b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b015190508c80610c6f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b818110611179575092600192859261012098966007989610611142575b505050811b019055610e18565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558b8080611135565b92936020600181928786015181550195019301611118565b015190508a80610c6f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b81811061122757509084600195949392106111f0575b505050811b019055610d96565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558980806111e3565b929360206001819287860151815501950193016111cd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b8181106112ca5750908460019594939210611293575b505050811b019055610d1d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611286565b92936020600181928786015181550195019301611270565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b81811061136d5750908460019594939210611336575b505050811b019055610ca4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055898080611329565b92936020600181928786015181550195019301611313565b989a979593918b60009b989694929b52600e60205260406000209a5b8b54600181111561159e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161156f576113e0908d614069565b90549060031b1c8060005260076020528d60406000209161150e600184019173ffffffffffffffffffffffffffffffffffffffff83547c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff821617855516600052600f60205260406000209360028601948554600052602052604060002063ffffffff61148c81835416613955565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000082541617905573ffffffffffffffffffffffffffffffffffffffff845416600052600a6020526114e282604060002061461d565b508454600052600b6020526114fb82604060002061461d565b50600052600e602052604060002061461d565b507f52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a397861156773ffffffffffffffffffffffffffffffffffffffff8554935416935494604051918291602083526003602084019101612e7b565b0390a46113a1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50919395979a50919395979a98610a7c565b6115bc6084353361427e565b610a77565b837f436f975400000000000000000000000000000000000000000000000000000000600052600452602060245260446000fd5b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b8a7f36a7c50300000000000000000000000000000000000000000000000000000000600052600452604060245260446000fd5b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b92509250929394965061168d33613da3565b611698604435614327565b6116a2828861439a565b6116ab846143df565b6116b58633613e05565b958654600052600960205260006040812055604435600052600960205260406000205585549660443587556004870191611711604051611700816116f98188612e7b565b0382612b51565b61170b36858561315f565b9061441e565b156119e1575b5050506005850191611733604051611700816116f98188612e7b565b156118cc575b505050600783019167ffffffffffffffff82116110ad5761175e82610c4d8554612e28565b600090601f831160011461180757827f67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e9593600295936117d1936000926117fc5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b015460405190815260406020820152806117f73396604435966040840191613604565b0390a4005b013590508a80610c6f565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106118b457509260019285927f67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e9896600298961061187c575b505050811b0190556117d4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905589808061186f565b91936020600181928787013581550195019201611834565b67ffffffffffffffff82116110ad576118e982610c4d8554612e28565b600090601f8311600114611941576119369291600091836117fc5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b868080611739565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106119c95750908460019594939210611991575b505050811b019055611939565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055898080611984565b9193602060018192878701358155019501920161196e565b67ffffffffffffffff82116110ad576119fe82610c4d8554612e28565b600090601f8311600114611a6157611a4b929160009183611a565750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b888080611717565b013590508c80610c6f565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b818110611ae95750908460019594939210611ab1575b505050811b019055611a4e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558b8080611aa4565b91936020600181928787013581550195019201611a8e565b3461024b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5761080a6107fe604435602435600435613451565b3461024b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461024b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576004356024359063ffffffff821680920361024b57604435801515810361024b57611bee61390a565b15611c845760405191611c0083612b35565b8252602082019060018252600052601060205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b169116179055600080f35b600090815260106020526040812055005b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576004356000526010602052602063ffffffff60406000205416604051908152f35b3461024b57611cf336612d29565b8115611d8b5760ff811615611d815760ff905b16808211611d505750611d1833613da3565b60005b818110611d2457005b80611d326001928486613429565b356000526009602052611d4a60406000205433613fde565b01611d1b565b907fb451d9ef0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5060ff601e611d06565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5760005473ffffffffffffffffffffffffffffffffffffffff81163303611e74577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b57611ed633613da3565b6004356000526009602052604060002054611ef18133613e05565b60ff600182015460e01c16611f05816129bf565b15611f21576104d091611f1c60028301543361427e565b613e84565b7f6f861db10000000000000000000000000000000000000000000000000000000060005260046000fd5b3461024b57611f5936612bfb565b919060038310156120cd57611f7791831594856120bf575b86613b6a565b60009115611fe2575b5073ffffffffffffffffffffffffffffffffffffffff819216611fa28161421b565b908083526005602052826040812055611fba81614487565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a480f35b9073ffffffffffffffffffffffffffffffffffffffff83168152600a602052604081205b8054801561209f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111612072576120429082614069565b90549060031b1c6120538186613e05565b60028503612069576120649161398c565b612006565b612064916140d0565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5090915073ffffffffffffffffffffffffffffffffffffffff9050611f80565b6120c887613b16565b611f71565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5760043560005260096020526104d06040600020546121498133613e05565b9060ff600183015460e01c1661215e816129bf565b6140d057600282015461217e604051610a5c816116f98160038901612e7b565b9033600052600a60205261219683604060002061461d565b5080600052600b6020526121ae83604060002061461d565b5033600052600f602052604060002090600052602052604060002063ffffffff6121da81835416613955565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055600052600e60205261221881604060002061461d565b506140d0565b3461024b5761222c36612d29565b8115611d8b5760ff81161561256a5760ff905b16808211611d50575061225133613da3565b61225a81612f1d565b6122676040519182612b51565b8181527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061229483612f1d565b013660208301376122a482612f1d565b906122b26040519283612b51565b8282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06122df84612f1d565b01366020840137600090815b84811061245f575060005b82811061236457858560005b81811061230b57005b806123196001928486613429565b3560005260096020526040600020546123328133613e05565b8360ff8183015460e01c16612346816129bf565b14612354575b505001612302565b61235d91613e84565b848061234c565b61236e8183612f4f565b5163ffffffff61237e8387612f4f565b511681600052600b60205261239881604060002054612f35565b82600052601060205263ffffffff60406000205416106124315733600052600f60205260406000208260005260205263ffffffff604060002054160163ffffffff811161156f5763ffffffff806123ef84336130e7565b1691161161240057506001016122f6565b7f6721d3de000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b507fcad00a8d0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61246a818688613429565b35600052600960205261248260406000205433613e05565b60ff600182015460e01c16612496816129bf565b1561256157600201546000805b85811061250b575b50156124bd575b506001905b016122eb565b6124ca8484959395612f4f565b5260016124d78286612f4f565b527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461156f57600180910192906124b2565b826125168287612f4f565b5114612524576001016124a3565b905063ffffffff6125358288612f4f565b51169063ffffffff821461156f57600161255463ffffffff9289612f4f565b92011690526001886124ab565b506001906124b7565b5060ff601e61223f565b3461024b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b576125ab612af5565b60243567ffffffffffffffff811161024b5761080a916125d26107fe923690600401612bcd565b60643592604435929061326b565b3461024b576125ee36612bfb565b9360038594929410156120cd576104d094613b6a5761260c81613b16565b613b6a565b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5761264861390a565b6004356000526009602052604060002054806000526007602052604060002060ff600182015460e01c1661267b816129bf565b1561268257005b6104d09161398c565b3461024b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5760206126d06126c7612af5565b602435906130e7565b63ffffffff60405191168152f35b3461024b5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5761080a612718612b92565b6040519182916020835260208301906129c9565b3461024b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b57612769602435600435612f92565b60405180916020820160208352815180915260206040840192019060005b818110612795575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612787565b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b57602061282973ffffffffffffffffffffffffffffffffffffffff612815612af5565b166000526004602052604060002054151590565b6040519015158152f35b3461024b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024b5761286a612dd9565b5060043560005260096020526040600020548015612995576000526007602052604060002073ffffffffffffffffffffffffffffffffffffffff604051916128b183612b18565b80548352600761297060018301549260ff6020870194868116865267ffffffffffffffff8160a01c16604089015260e01c166128ec816129bf565b60608701526002810154608087015260405161290f816116f98160038601612e7b565b60a0870152604051612928816116f98160048601612e7b565b60c0870152604051612941816116f98160058601612e7b565b60e087015260405161295a816116f98160068601612e7b565b6101008701526116f96040518094819301612e7b565b6101208401525116156129955761080a90604051918291602083526020830190612a28565b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b600211156120cd57565b919082519283825260005b848110612a135750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016129d4565b612af2918151815273ffffffffffffffffffffffffffffffffffffffff602083015116602082015267ffffffffffffffff60408301511660408201526060820151612a72816129bf565b606082015260808201516080820152610120612ae0612acc612aba612aa860a087015161014060a08801526101408701906129c9565b60c087015186820360c08801526129c9565b60e086015185820360e08701526129c9565b6101008501518482036101008601526129c9565b920151906101208184039101526129c9565b90565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361024b57565b610140810190811067ffffffffffffffff8211176110ad57604052565b6040810190811067ffffffffffffffff8211176110ad57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176110ad57604052565b60405190612ba1604083612b51565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b9181601f8401121561024b5782359167ffffffffffffffff831161024b576020838186019501011161024b57565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261024b5760043573ffffffffffffffffffffffffffffffffffffffff8116810361024b5791602435916044359067ffffffffffffffff821161024b57612c6991600401612bcd565b9091606435600381101561024b5790565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310612cad57505050505090565b9091929394602080612ce9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951612a28565b97019301930191939290612c9e565b9181601f8401121561024b5782359167ffffffffffffffff831161024b576020808501948460051b01011161024b57565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261024b576004359067ffffffffffffffff821161024b57612d7291600401612cf8565b909160243560ff8116810361024b5790565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261024b5760043591602435916044359067ffffffffffffffff821161024b57612dd591600401612bcd565b9091565b60405190612de682612b18565b606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b90600182811c92168015612e71575b6020831014612e4257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612e37565b60009291815491612e8b83612e28565b8083529260018116908115612ee15750600114612ea757505050565b60009081526020812093945091925b838310612ec7575060209250010190565b600181602092949394548385870101520191019190612eb6565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b67ffffffffffffffff81116110ad5760051b60200190565b9190820180921161156f57565b9190820391821161156f57565b8051821015612f635760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060035490818310156130ca578181612fbc93612fb0869485612f35565b11156130ba5750612f42565b90612fc682612f1d565b91612fd46040519384612b51565b8083527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061300182612f1d565b013660208501376003549160005b82811061301d575050505090565b6130278183612f35565b60008582101561308d57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600081905260056020526001919073ffffffffffffffffffffffffffffffffffffffff166130868288612f4f565b520161300f565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b6130c5915082612f35565b612f42565b5050506040516130db602082612b51565b60008152600036813790565b73ffffffffffffffffffffffffffffffffffffffff166000526011602052604060002081600052602052604060002060206040519161312583612b35565b549160ff63ffffffff841693848352831c161515918291015261315a5750600052601060205263ffffffff6040600020541690565b905090565b92919267ffffffffffffffff82116110ad57604051916131a7601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184612b51565b82948184528183011161024b578281602093846000960137010152565b604051906131d3602083612b51565b600080835282815b8281106131e757505050565b6020906131f2612dd9565b828285010152016131db565b9061320882612f1d565b6132156040519182612b51565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06132438294612f1d565b019060005b82811061325457505050565b60209061325f612dd9565b82828501015201613248565b9061327b9061328193369161315f565b90613d62565b91826000526008602052604060002054908183101561341d576132bb918160648593118015613415575b61340d575b81612fb08285612f35565b906132c5826131fe565b9260005b8381106132d7575050505090565b6001908260005260086020526132fb60406000206132f58387612f35565b90614069565b90549060031b1c6000526007602052604060002060076133eb6040519261332184612b18565b8054845260ff8682015473ffffffffffffffffffffffffffffffffffffffff8116602087015267ffffffffffffffff8160a01c16604087015260e01c16613367816129bf565b60608501526002810154608085015260405161338a816116f98160038601612e7b565b60a08501526040516133a3816116f98160048601612e7b565b60c08501526040516133bc816116f98160058601612e7b565b60e08501526040516133d5816116f98160068601612e7b565b6101008501526116f96040518094819301612e7b565b6101208201526133fb8288612f4f565b526134068187612f4f565b50016132c9565b5060646132b0565b5080156132ab565b50505050612af26131c4565b9190811015612f635760051b0190565b63ffffffff60019116019063ffffffff821161156f57565b9182600052600c602052604060002054908183101561341d57613487918184921580156135095761340d5781612fb08285612f35565b90613491826131fe565b9260005b8381106134a3575050505090565b60019082600052600c6020526134c160406000206132f58387612f35565b90549060031b1c6000526007602052604060002060076134e76040519261332184612b18565b6101208201526134f78288612f4f565b526135028187612f4f565b5001613495565b50606481116132ab565b73ffffffffffffffffffffffffffffffffffffffff169182600052600d602052604060002054908183101561341d576135619181606485931180156134155761340d5781612fb08285612f35565b9061356b826131fe565b9260005b83811061357d575050505090565b60019082600052600d60205261359b60406000206132f58387612f35565b90549060031b1c6000526007602052604060002060076135c16040519261332184612b18565b6101208201526135d18288612f4f565b526135dc8187612f4f565b500161356f565b3573ffffffffffffffffffffffffffffffffffffffff8116810361024b5790565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b909260809273ffffffffffffffffffffffffffffffffffffffff612af29795168352602083015260408201528160608201520191613604565b919290928242116138d45761369e336000526004602052604060002054151590565b6138a65783600052600660205260ff604060002054166138745760006136c2612b92565b604051613730816136f8602082019486865233604084015246606084015230608084015260e060a08401526101008301906129c9565b8860c08301528960e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612b51565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613771603c822061376b36868661315f565b9061444b565b909192600483101561384757826137f45750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff16156137b85750505050565b906137f0916040519485947f335d4ce10000000000000000000000000000000000000000000000000000000086523360048701613643565b0390fd5b8593505060ff6138376040519586957fd36ab6b9000000000000000000000000000000000000000000000000000000008752606060048801526064870191613604565b9216602484015260448301520390fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b837f77a33858000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b7fd9a5f5ca000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b827f502d038700000000000000000000000000000000000000000000000000000000600052336004524260245260445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361392b57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff821161156f57565b600182019173ffffffffffffffffffffffffffffffffffffffff835416600052600f6020527f52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a397873ffffffffffffffffffffffffffffffffffffffff60406000209260028101938454600052602052604060002063ffffffff613a1081835416613955565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008254161790558354600052600b602052613a5085604060002061461d565b5081865416600052600a602052613a6b85604060002061461d565b50613aac8287541695613a9a600384019760405190613a9582613a8e818d612e7b565b0383612b51565b613d62565b600052600e602052604060002061461d565b508554957c01000000000000000000000000000000000000000000000000000000007fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff881617905554925494613b116040519283926020845216956020830190612e7b565b0390a4565b73ffffffffffffffffffffffffffffffffffffffff16600052600a602052604060002054613b4057565b7f61bc2e180000000000000000000000000000000000000000000000000000000060005260046000fd5b91929092834211613d1657613b7e83613da3565b73ffffffffffffffffffffffffffffffffffffffff8316613b9e8161421b565b90600090613be2613c1a613bb0612b92565b604051928391602083019560018752604084015246606084015230608084015260e060a08401526101008301906129c9565b8a60c08301528660e0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282612b51565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613c55603c822061376b36878761315f565b90919260048310156138475782613cd55750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1615613c9d575050505050565b906137f092916040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701613643565b506040517fd36ab6b90000000000000000000000000000000000000000000000000000000081526060600482015291829160ff613837606485018a8a613604565b8373ffffffffffffffffffffffffffffffffffffffff847f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b90613d9d6109ab9160405192839173ffffffffffffffffffffffffffffffffffffffff602084019616865260408084015260608301906129c9565b51902090565b73ffffffffffffffffffffffffffffffffffffffff16613dd0816000526004602052604060002054151590565b15613dd85750565b7fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff6001830154169081156129955773ffffffffffffffffffffffffffffffffffffffff16809103613e57575090565b7f31ee6dc70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600182019173ffffffffffffffffffffffffffffffffffffffff835416600052600f6020527f3131ee488b1b459e313b249599f9ebc23dfc5dbfcbd082dd503d171c7436fe8273ffffffffffffffffffffffffffffffffffffffff60406000209260028101938454600052602052604060002063ffffffff613f0881835416613439565b167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008254161790558354600052600b602052613f488560406000206147b1565b5081865416600052600a602052613f638560406000206147b1565b50613f988287541695613f86600384019760405190613a9582613a8e818d612e7b565b600052600e60205260406000206147b1565b508554957fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff8716905554925494613b116040519283926020845216956020830190612e7b565b81613fe891613e05565b600160ff8183015460e01c16613ffd816129bf565b14611f215761400b9161398c565b565b818110614018575050565b6000815560010161400d565b9190601f811161403357505050565b61400b926000526020600020906020601f840160051c8301931061405f575b601f0160051c019061400d565b9091508190614052565b8054821015612f635760005260206000200190600090565b61408b8154612e28565b9081614095575050565b81601f600093116001146140a7575055565b818352602083206140c391601f0160051c81019060010161400d565b8082528160208120915555565b90600281018054600052600c6020526140ed83604060002061461d565b506001820173ffffffffffffffffffffffffffffffffffffffff815416600052600d60205261412084604060002061461d565b507fbe9d8e7e7313f6e3020af5157a64045cca8bcadce6e6497ca12fe56dbb4fdeb96141bc73ffffffffffffffffffffffffffffffffffffffff8084541695614179600382019760405190613a9582613a8e818d612e7b565b600052600860205261418f88604060002061461d565b50805460005260096020526000604081205554935416935494604051918291602083526020830190612e7b565b0390a46000526007602052600760406000206000815560006001820155600060028201556141ec60038201614081565b6141f860048201614081565b61420460058201614081565b61421060068201614081565b0161408b8154612e28565b80600052600560205260406000205490811580614268575b61423b575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600460205260406000205415614233565b81600052600b60205260406000205482600052601060205263ffffffff6040600020541611156124315773ffffffffffffffffffffffffffffffffffffffff81169081600052600f60205260406000208360005260205263ffffffff6142ec848260406000205416936130e7565b1611156142f7575050565b7f6721d3de0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8015614370578060005260096020526040600020546143435750565b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f7dc2f4e10000000000000000000000000000000000000000000000000000000060005260046000fd5b60c881116143ae575060c881116143ae5750565b7ecd56a80000000000000000000000000000000000000000000000000000000060005260045260c860245260446000fd5b61020081116143eb5750565b7f354f25140000000000000000000000000000000000000000000000000000000060005260045261020060245260446000fd5b9081518151908181149384614435575b5050505090565b602092939450820120920120143880808061442e565b815191906041830361447c5761447592506020820151906060604084015193015160001a90614806565b9192909190565b505060009160029190565b6000818152600460205260409020548015614616577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161156f57600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161156f578181036145a7575b5050506003548015614578577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614535816003614069565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6145fe6145b86145c9936003614069565b90549060031b1c9283926003614069565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260046020526040600020553880806144fc565b5050600090565b9060018201918160005282602052604060002054801515600014614748577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161156f578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161156f57818103614711575b50505080548015614578577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906146d28282614069565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6147316147216145c99386614069565b90549060031b1c92839286614069565b90556000528360205260406000205538808061469a565b50505050600090565b806000526004602052604060002054156000146147ab57600354680100000000000000008110156110ad576147926145c98260018594016003556003614069565b9055600354906000526004602052604060002055600190565b50600090565b600082815260018201602052604090205461461657805490680100000000000000008210156110ad57826147ef6145c9846001809601855584614069565b905580549260005201602052604060002055600190565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161489c579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156148905760005173ffffffffffffffffffffffffffffffffffffffff8116156148845790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
}

var WorkflowRegistryABI = WorkflowRegistryMetaData.ABI

var WorkflowRegistryBin = WorkflowRegistryMetaData.Bin

func DeployWorkflowRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WorkflowRegistry, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WorkflowRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WorkflowRegistry{address: address, abi: *parsed, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

type WorkflowRegistry struct {
	address common.Address
	abi     abi.ABI
	WorkflowRegistryCaller
	WorkflowRegistryTransactor
	WorkflowRegistryFilterer
}

type WorkflowRegistryCaller struct {
	contract *bind.BoundContract
}

type WorkflowRegistryTransactor struct {
	contract *bind.BoundContract
}

type WorkflowRegistryFilterer struct {
	contract *bind.BoundContract
}

type WorkflowRegistrySession struct {
	Contract     *WorkflowRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryCallerSession struct {
	Contract *WorkflowRegistryCaller
	CallOpts bind.CallOpts
}

type WorkflowRegistryTransactorSession struct {
	Contract     *WorkflowRegistryTransactor
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryRaw struct {
	Contract *WorkflowRegistry
}

type WorkflowRegistryCallerRaw struct {
	Contract *WorkflowRegistryCaller
}

type WorkflowRegistryTransactorRaw struct {
	Contract *WorkflowRegistryTransactor
}

func NewWorkflowRegistry(address common.Address, backend bind.ContractBackend) (*WorkflowRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(WorkflowRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWorkflowRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistry{address: address, abi: abi, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

func NewWorkflowRegistryCaller(address common.Address, caller bind.ContractCaller) (*WorkflowRegistryCaller, error) {
	contract, err := bindWorkflowRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryCaller{contract: contract}, nil
}

func NewWorkflowRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkflowRegistryTransactor, error) {
	contract, err := bindWorkflowRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryTransactor{contract: contract}, nil
}

func NewWorkflowRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkflowRegistryFilterer, error) {
	contract, err := bindWorkflowRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryFilterer{contract: contract}, nil
}

func bindWorkflowRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.WorkflowRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canLinkOwner", validityTimestamp, proof, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canUnlinkOwner", owner, validityTimestamp, signature, action)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetLinkedOwners(opts *bind.CallOpts, start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getLinkedOwners", start, batchSize)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetLinkedOwners(start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, batchSize)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetLinkedOwners(start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, batchSize)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadata(opts *bind.CallOpts, workflowID [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadata", workflowID)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadata)).(*WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadata(workflowID [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadata(&_WorkflowRegistry.CallOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadata(workflowID [32]byte) (WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadata(&_WorkflowRegistry.CallOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByDON(opts *bind.CallOpts, donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByDON", donLabel, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByDON(donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donLabel, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByDON(donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByDON(&_WorkflowRegistry.CallOpts, donLabel, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByOwner", owner, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowMetadataListByOwnerName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowMetadataListByOwnerName", owner, workflowName, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadata)).(*[]WorkflowRegistryWorkflowMetadata)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowMetadataListByOwnerName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowMetadataListByOwnerName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error) {
	return _WorkflowRegistry.Contract.GetWorkflowMetadataListByOwnerName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isAllowedSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isOwnerLinked", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalLinkedOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_WorkflowRegistry *WorkflowRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) ActivateWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "activateWorkflow", workflowID)
}

func (_WorkflowRegistry *WorkflowRegistrySession) ActivateWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) ActivateWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseWorkflow", workflowID)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchActivateWorkflows(opts *bind.TransactOpts, workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchActivateWorkflows", workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchActivateWorkflows(workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchActivateWorkflows(workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchPauseWorkflows(opts *bind.TransactOpts, workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchPauseWorkflows", workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchPauseWorkflows(workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchPauseWorkflows(workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIDs, batchSizeOverride)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) DeleteWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "deleteWorkflow", workflowID)
}

func (_WorkflowRegistry *WorkflowRegistrySession) DeleteWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) DeleteWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "linkOwner", validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistrySession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) PauseWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "pauseWorkflow", workflowID)
}

func (_WorkflowRegistry *WorkflowRegistrySession) PauseWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) PauseWorkflow(workflowID [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowID)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONRegistry", registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONRegistry(registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONRegistry(registry common.Address, chainSelector *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_WorkflowRegistry *WorkflowRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "unlinkOwner", owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature, action)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateAllowedSigners", signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowID [32]byte, status uint8, donLabel [32]byte, binaryURL string, configURL string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "upsertWorkflow", workflowName, tag, workflowID, status, donLabel, binaryURL, configURL, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpsertWorkflow(workflowName string, tag string, workflowID [32]byte, status uint8, donLabel [32]byte, binaryURL string, configURL string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowID, status, donLabel, binaryURL, configURL, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpsertWorkflow(workflowName string, tag string, workflowID [32]byte, status uint8, donLabel [32]byte, binaryURL string, configURL string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowID, status, donLabel, binaryURL, configURL, attributes, keepAlive)
}

type WorkflowRegistryAllowedSignersUpdatedIterator struct {
	Event *WorkflowRegistryAllowedSignersUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryAllowedSignersUpdated)
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
		it.Event = new(WorkflowRegistryAllowedSignersUpdated)
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

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryAllowedSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryAllowedSignersUpdated struct {
	Signers []common.Address
	Allowed bool
	Raw     types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "AllowedSignersUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryAllowedSignersUpdatedIterator{contract: _WorkflowRegistry.contract, event: "AllowedSignersUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "AllowedSignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryAllowedSignersUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error) {
	event := new(WorkflowRegistryAllowedSignersUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryDONRegistryUpdatedIterator struct {
	Event *WorkflowRegistryDONRegistryUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryDONRegistryUpdated)
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
		it.Event = new(WorkflowRegistryDONRegistryUpdated)
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

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryDONRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryDONRegistryUpdated struct {
	OldAddr       common.Address
	NewAddr       common.Address
	ChainSelector *big.Int
	Raw           types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONRegistryUpdated(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (*WorkflowRegistryDONRegistryUpdatedIterator, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONRegistryUpdated", oldAddrRule, newAddrRule, chainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONRegistryUpdatedIterator{contract: _WorkflowRegistry.contract, event: "DONRegistryUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (event.Subscription, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}
	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONRegistryUpdated", oldAddrRule, newAddrRule, chainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryDONRegistryUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "DONRegistryUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseDONRegistryUpdated(log types.Log) (*WorkflowRegistryDONRegistryUpdated, error) {
	event := new(WorkflowRegistryDONRegistryUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "DONRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipLinkUpdatedIterator struct {
	Event *WorkflowRegistryOwnershipLinkUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipLinkUpdated)
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
		it.Event = new(WorkflowRegistryOwnershipLinkUpdated)
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

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipLinkUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipLinkUpdated struct {
	Owner common.Address
	Proof [32]byte
	Added bool
	Raw   types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipLinkUpdated", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipLinkUpdatedIterator{contract: _WorkflowRegistry.contract, event: "OwnershipLinkUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipLinkUpdated", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipLinkUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error) {
	event := new(WorkflowRegistryOwnershipLinkUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferRequestedIterator struct {
	Event *WorkflowRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferRequested)
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
		it.Event = new(WorkflowRegistryOwnershipTransferRequested)
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

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferRequestedIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferRequested)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error) {
	event := new(WorkflowRegistryOwnershipTransferRequested)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferredIterator struct {
	Event *WorkflowRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferred)
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
		it.Event = new(WorkflowRegistryOwnershipTransferred)
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

func (it *WorkflowRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferredIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferred)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error) {
	event := new(WorkflowRegistryOwnershipTransferred)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowActivatedIterator struct {
	Event *WorkflowRegistryWorkflowActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowActivated)
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
		it.Event = new(WorkflowRegistryWorkflowActivated)
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

func (it *WorkflowRegistryWorkflowActivatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowActivated struct {
	WorkflowID   [32]byte
	Owner        common.Address
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowActivated(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowActivatedIterator, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowActivated", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowActivatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowActivated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowActivated", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowActivated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error) {
	event := new(WorkflowRegistryWorkflowActivated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowDeletedIterator struct {
	Event *WorkflowRegistryWorkflowDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowDeleted)
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
		it.Event = new(WorkflowRegistryWorkflowDeleted)
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

func (it *WorkflowRegistryWorkflowDeletedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowDeleted struct {
	WorkflowID   [32]byte
	Owner        common.Address
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDeleted(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowDeletedIterator, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDeleted", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDeletedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDeleted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDeleted", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowDeleted)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDeleted", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error) {
	event := new(WorkflowRegistryWorkflowDeleted)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowPausedIterator struct {
	Event *WorkflowRegistryWorkflowPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowPaused)
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
		it.Event = new(WorkflowRegistryWorkflowPaused)
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

func (it *WorkflowRegistryWorkflowPausedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowPaused struct {
	WorkflowID   [32]byte
	Owner        common.Address
	DonLabel     [32]byte
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowPaused(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowPausedIterator, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowPaused", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowPausedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowPaused", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var donLabelRule []interface{}
	for _, donLabelItem := range donLabel {
		donLabelRule = append(donLabelRule, donLabelItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowPaused", workflowIDRule, ownerRule, donLabelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowPaused)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowPaused(log types.Log) (*WorkflowRegistryWorkflowPaused, error) {
	event := new(WorkflowRegistryWorkflowPaused)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowRegisteredIterator struct {
	Event *WorkflowRegistryWorkflowRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowRegistered)
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
		it.Event = new(WorkflowRegistryWorkflowRegistered)
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

func (it *WorkflowRegistryWorkflowRegisteredIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowRegistered struct {
	WorkflowID   [32]byte
	Owner        common.Address
	DonLabel     [32]byte
	Status       uint8
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowRegistered(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowRegistered", workflowIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowRegisteredIterator{contract: _WorkflowRegistry.contract, event: "WorkflowRegistered", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIDRule []interface{}
	for _, workflowIDItem := range workflowID {
		workflowIDRule = append(workflowIDRule, workflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowRegistered", workflowIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowRegistered)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowRegistered(log types.Log) (*WorkflowRegistryWorkflowRegistered, error) {
	event := new(WorkflowRegistryWorkflowRegistered)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowUpdated)
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
		it.Event = new(WorkflowRegistryWorkflowUpdated)
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

func (it *WorkflowRegistryWorkflowUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowUpdated struct {
	OldWorkflowID [32]byte
	NewWorkflowID [32]byte
	Owner         common.Address
	DonLabel      [32]byte
	WorkflowName  string
	Raw           types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowID [][32]byte, newWorkflowID [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error) {

	var oldWorkflowIDRule []interface{}
	for _, oldWorkflowIDItem := range oldWorkflowID {
		oldWorkflowIDRule = append(oldWorkflowIDRule, oldWorkflowIDItem)
	}
	var newWorkflowIDRule []interface{}
	for _, newWorkflowIDItem := range newWorkflowID {
		newWorkflowIDRule = append(newWorkflowIDRule, newWorkflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowUpdated", oldWorkflowIDRule, newWorkflowIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowID [][32]byte, newWorkflowID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var oldWorkflowIDRule []interface{}
	for _, oldWorkflowIDItem := range oldWorkflowID {
		oldWorkflowIDRule = append(oldWorkflowIDRule, oldWorkflowIDItem)
	}
	var newWorkflowIDRule []interface{}
	for _, newWorkflowIDItem := range newWorkflowID {
		newWorkflowIDRule = append(newWorkflowIDRule, newWorkflowIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowUpdated", oldWorkflowIDRule, newWorkflowIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowUpdated(log types.Log) (*WorkflowRegistryWorkflowUpdated, error) {
	event := new(WorkflowRegistryWorkflowUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_WorkflowRegistry *WorkflowRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WorkflowRegistry.abi.Events["AllowedSignersUpdated"].ID:
		return _WorkflowRegistry.ParseAllowedSignersUpdated(log)
	case _WorkflowRegistry.abi.Events["DONRegistryUpdated"].ID:
		return _WorkflowRegistry.ParseDONRegistryUpdated(log)
	case _WorkflowRegistry.abi.Events["OwnershipLinkUpdated"].ID:
		return _WorkflowRegistry.ParseOwnershipLinkUpdated(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferRequested(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferred"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferred(log)
	case _WorkflowRegistry.abi.Events["WorkflowActivated"].ID:
		return _WorkflowRegistry.ParseWorkflowActivated(log)
	case _WorkflowRegistry.abi.Events["WorkflowDeleted"].ID:
		return _WorkflowRegistry.ParseWorkflowDeleted(log)
	case _WorkflowRegistry.abi.Events["WorkflowPaused"].ID:
		return _WorkflowRegistry.ParseWorkflowPaused(log)
	case _WorkflowRegistry.abi.Events["WorkflowRegistered"].ID:
		return _WorkflowRegistry.ParseWorkflowRegistered(log)
	case _WorkflowRegistry.abi.Events["WorkflowUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WorkflowRegistryAllowedSignersUpdated) Topic() common.Hash {
	return common.HexToHash("0x861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e1316")
}

func (WorkflowRegistryDONRegistryUpdated) Topic() common.Hash {
	return common.HexToHash("0xc60b6bc6904a32fe864ae9ad697fd792f40e1d9e9fee5da0a10d0fb26b7380d6")
}

func (WorkflowRegistryOwnershipLinkUpdated) Topic() common.Hash {
	return common.HexToHash("0x07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35")
}

func (WorkflowRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (WorkflowRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (WorkflowRegistryWorkflowActivated) Topic() common.Hash {
	return common.HexToHash("0x3131ee488b1b459e313b249599f9ebc23dfc5dbfcbd082dd503d171c7436fe82")
}

func (WorkflowRegistryWorkflowDeleted) Topic() common.Hash {
	return common.HexToHash("0xbe9d8e7e7313f6e3020af5157a64045cca8bcadce6e6497ca12fe56dbb4fdeb9")
}

func (WorkflowRegistryWorkflowPaused) Topic() common.Hash {
	return common.HexToHash("0x52613f4f9f39f66629197514b47e67cab5e5c40b83fad8fccec907d71a8a3978")
}

func (WorkflowRegistryWorkflowRegistered) Topic() common.Hash {
	return common.HexToHash("0x03bdc9e1a7b65f76fa1e989e76c28e6ac9f556262cfed31a11a9461e77ad2ff5")
}

func (WorkflowRegistryWorkflowUpdated) Topic() common.Hash {
	return common.HexToHash("0x67057592e7b546efc47d8093ece5f7d8837b769a512710d66a892555875bc05e")
}

func (_WorkflowRegistry *WorkflowRegistry) Address() common.Address {
	return _WorkflowRegistry.address
}

type WorkflowRegistryInterface interface {
	CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) error

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, batchSize *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error)

	GetWorkflowMetadata(opts *bind.CallOpts, workflowID [32]byte) (WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByDON(opts *bind.CallOpts, donLabel [32]byte, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	GetWorkflowMetadataListByOwnerName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadata, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error)

	BatchPauseWorkflows(opts *bind.TransactOpts, workflowIDs [][32]byte, batchSizeOverride uint8) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowID [32]byte) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	SetDONRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector *big.Int) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte, action uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowID [32]byte, status uint8, donLabel [32]byte, binaryURL string, configURL string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterDONRegistryUpdated(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (*WorkflowRegistryDONRegistryUpdatedIterator, error)

	WatchDONRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONRegistryUpdated, oldAddr []common.Address, newAddr []common.Address, chainSelector []*big.Int) (event.Subscription, error)

	ParseDONRegistryUpdated(log types.Log) (*WorkflowRegistryDONRegistryUpdated, error)

	FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error)

	WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error)

	ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error)

	FilterWorkflowActivated(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowActivatedIterator, error)

	WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error)

	FilterWorkflowDeleted(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowDeletedIterator, error)

	WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error)

	FilterWorkflowPaused(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (*WorkflowRegistryWorkflowPausedIterator, error)

	WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowID [][32]byte, owner []common.Address, donLabel [][32]byte) (event.Subscription, error)

	ParseWorkflowPaused(log types.Log) (*WorkflowRegistryWorkflowPaused, error)

	FilterWorkflowRegistered(opts *bind.FilterOpts, workflowID [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error)

	WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowID [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowRegistered(log types.Log) (*WorkflowRegistryWorkflowRegistered, error)

	FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowID [][32]byte, newWorkflowID [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error)

	WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowID [][32]byte, newWorkflowID [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowUpdated(log types.Log) (*WorkflowRegistryWorkflowUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
