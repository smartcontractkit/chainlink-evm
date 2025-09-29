// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package workflow_registry_wrapper_v2_dev

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

type WorkflowRegistryConfig struct {
	MaxNameLen   uint8
	MaxTagLen    uint8
	MaxUrlLen    uint8
	MaxAttrLen   uint16
	MaxExpiryLen uint32
}

type WorkflowRegistryDonConfigView struct {
	DonHash          [32]byte
	Family           string
	DonLimit         uint32
	DefaultUserLimit uint32
}

type WorkflowRegistryEventRecord struct {
	EventType uint8
	Timestamp uint32
	Payload   []byte
}

type WorkflowRegistryOwnerAllowlistedRequest struct {
	RequestDigest   [32]byte
	Owner           common.Address
	ExpiryTimestamp uint32
}

type WorkflowRegistryUserOverrideView struct {
	User  common.Address
	Limit uint32
}

type WorkflowRegistryWorkflowMetadataView struct {
	WorkflowId   [32]byte
	Owner        common.Address
	CreatedAt    uint64
	Status       uint8
	WorkflowName string
	BinaryUrl    string
	ConfigUrl    string
	Tag          string
	Attributes   []byte
	DonFamily    string
}

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistRequest\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedSigners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowlistedRequests\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allowlistedRequests\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.OwnerAllowlistedRequest[]\",\"components\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilitiesRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.Config\",\"components\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonConfigs\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.DonConfigView[]\",\"components\":[{\"name\":\"donHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"family\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDONOverrides\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.UserOverrideView[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowById\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapabilitiesRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"nameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"urlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"attrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"expiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"userLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsOnDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowedSigners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowlistedRequests\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEvents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONFamily\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilitiesRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BinaryURLRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateDONFamilyForPausedWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidExpiryTimestamp\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxAllowed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONDefaultLimitExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWorkflowIDNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60803460a8573315609757600180546001600160a01b0319163317905560a081016001600160401b03811182821017608157604090815280825260208083015260c882820152610400606083015262093a80608090920191909152600280546001600160481b03191667093a800400c82040179055516154aa90816100ae8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fdfe610140604052600436101561001357600080fd5b60003560e01c806301b7690514612ec05780630987294c14612e8657806317e0edfc14612ddf578063181f5a7714612dc35780631c08b00a14612cf45780631c71682c14612b05578063245b8e4e14612ac7578063289bd10814612a0a5780632afc4130146129b25780632c50a9551461295557806339d68c6a1461280357806339e43234146127815780633bb70578146126d05780633c17181b1461267d5780633c54b50b1461262a5780633d90a1081461260c5780634b6d2e5b1461246a578063530979d6146123db578063610431931461233e578063695e1340146122ed5780636caffc4c1461226b57806370ae26401461223557806379ba50971461219c578063865ec9e0146121635780638b42a96d146120155780638c42ffc514611f615780638da5cb5b14611f3a57806394ea0da614611d18578063952bb98414611ccd578063a0b8a4fe14611caf578063a408901614611c0c578063a6008f2014611aea578063a7d0185814611a43578063afbb240114611a2c578063b377bfc514610bc6578063b668435f146109b7578063ba87068614610999578063bae5c29a14610972578063bdf6b4ff1461090d578063be674333146108e6578063bf2a0d9314610893578063c3f909d4146107e2578063c59a655a14610708578063cabb9e7a146106c4578063d8b80738146105e2578063d8e4a7241461046d578063dc101969146103cf578063de49b95f146103b1578063e690f3321461032d578063ea32308b146102f35763f2fde38b1461024e57600080fd5b346102ee5760203660031901126102ee576001600160a01b0361026f6130ce565b61027761472f565b163381146102c457806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102ee5760203660031901126102ee576001600160a01b036103146130ce565b16600052600c6020526020604060002054604051908152f35b346102ee5760203660031901126102ee57610355336000526006602052604060002054151590565b1561039c57600435600052600b602052604060002054610375813361484e565b600160ff8183015460e01c1661038a8161323f565b0361039157005b61039a91614d16565b005b63c2dda3f960e01b6000523360045260246000fd5b346102ee5760003660031901126102ee576020601554604051908152f35b346102ee5760603660031901126102ee5760443560243567ffffffffffffffff82116102ee576104136104086001933690600401613211565b908360043533613d00565b3360005260076020528060406000205561042c336150bd565b5080600052600860205260406000208260ff19825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346102ee5760403660031901126102ee5760043567ffffffffffffffff81116102ee5761049e9036906004016133bb565b60243591821515928381036102ee576104b561472f565b60005b83811061054257505060405191806040840160408552526060830191906000905b80821061050f577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b9091928335906001600160a01b03821682036102ee57602080916001600160a01b036001941681520194019201906104d9565b6001600160a01b0361055d61055883878761384d565b614573565b16156105b85760019082156105935761058b6001600160a01b0361058561055884898961384d565b1661504c565b505b016104b8565b6105b26001600160a01b036105ac61055884898961384d565b16615297565b5061058d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ee5760203660031901126102ee5760043567ffffffffffffffff81116102ee576106139036906004016133bb565b801561069a57610630336000526006602052604060002054151590565b1561039c5760005b81811061064157005b8061064f600192848661384d565b35600052600b602052604060002054610668813361484e565b8360ff8183015460e01c1661067c8161323f565b0361068a575b505001610638565b61069391614d16565b8480610682565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ee5760203660031901126102ee5760206106fe6001600160a01b036106ea6130ce565b166000526004602052604060002054151590565b6040519015158152f35b346102ee5760403660031901126102ee5760043567ffffffffffffffff81116102ee5761073c610751913690600401613211565b91906024359261074a61472f565b36916135f7565b60208151910120600052600d60205260406000209060005b825490811515806107d1575b1561039a5760001982019182116107bb576107b16107966107b69386614eb6565b90549060031b1c806000526009602052604060002090614d16565b613b2d565b610769565b634e487b7160e01b600052601160045260246000fd5b508215806107755750828110610775565b346102ee5760003660031901126102ee57600060806040516108038161317b565b828152826020820152826040820152826060820152015260a06040516108288161317b565b63ffffffff60025461ffff60ff82169384815260ff60208201818560081c168152816040840191818760101c1683528760806060870196888a60181c168852019760281c16875260405198895251166020880152511660408601525116606084015251166080820152f35b346102ee5760603660031901126102ee5760043567ffffffffffffffff81116102ee576108c761039a913690600401613211565b6108cf613493565b906108d86134a6565b926108e161472f565b6142d6565b346102ee5760206108ff6108f936613453565b916140f4565b63ffffffff60405191168152f35b346102ee5760203660031901126102ee5760043567ffffffffffffffff81116102ee5761074a610941913690600401613211565b6020815191012060005260176020526040600181600020015463ffffffff825191818116835260201c166020820152f35b346102ee5760403660031901126102ee5760206106fe6109906130ce565b60243590614085565b346102ee5760003660031901126102ee576020601c54604051908152f35b346102ee5760803660031901126102ee576109d06130ce565b60243567ffffffffffffffff81116102ee576109f0903690600401613211565b906109f96134a6565b916064359283151584036102ee57610a0f61472f565b610a1a3683856135f7565b602081519101208060005260176020526040600020600181019063ffffffff82541615610b9f57600201604060006001600160a01b038a168152826020522096600014610b2357505463ffffffff92831692168211610af9577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c823249464010000000064ff00000000198254161781558263ffffffff19825416179055600052601a602052610ad66001600160a01b03604060002096168096615171565b50610aee604051938493604085526040850191613800565b9060208301520390a2005b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b919250507f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d84994506001600160a01b03861660005260205260006040812055600052601a602052610b826001600160a01b0360406000209516809561534b565b50610b9a604051928392602084526020840191613800565b0390a2005b60405163393f328760e11b81526020600482015280610bc260248201888a613800565b0390fd5b346102ee576101203660031901126102ee5760043567ffffffffffffffff81116102ee57610bf8903690600401613211565b610100526101205260243567ffffffffffffffff81116102ee57610c20903690600401613211565b90600260643510156102ee5760843567ffffffffffffffff81116102ee57610c4c903690600401613211565b9060e0529060a43567ffffffffffffffff81116102ee57610c71903690600401613211565b9160c43567ffffffffffffffff81116102ee57610c92903690600401613211565b9560e43567ffffffffffffffff81116102ee57610cb3903690600401613211565b929093610104359586151587036102ee57610cdb336000526006602052604060002054151590565b1561039c5760443515611a0257604435600052600b6020526040600020546119d25760025460ff8160101c1689156119a8578061196b575b5061ffff8160181c1680151580611962575b611931575084156119075760ff8160081c16801515806118fe575b6118cd575061010051156118a35760ff1680151580611897575b61186357506040516020810190610d8d81610d7f888761010051610120513389613f0e565b03601f1981018352826131b4565b5190209788600052600960205260406000206001600160a01b0360018201541680156000146114e1575050610dd2610dcc3661010051610120516135f7565b336145e2565b9a610de0368c60e0516135f7565b602081519101209815611486575b610df960643561323f565b6064351561145e575b604051978861012081011067ffffffffffffffff6101208b01111761128757610eaf899695610ea08897610ebe956101208a016040526044358a5260a0610e9460208c019a338c5260408d0160c05267ffffffffffffffff421660c0515260608d019a610e7060643561323f565b6064358c526080610e883661010051610120516135f7565b9e019d8e5236916135f7565b9b019a8b5236916135f7565b9860c08c01998a5236916135f7565b9760e08a0198895236916135f7565b610100880160a05260a051528860005260096020526040600020965187556001600160a01b038060018901935116166001600160a01b031983541617825560c051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b8085549360a01b1616911617825551610f458161323f565b610f4e8161323f565b60ff60e01b197cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff821161128757610faa82610fa485546134de565b85613cb9565b602090601f83116001146113f757610fdb9291600091836113ec575b50508160011b916000199060031b1c19161790565b90555b518051600385019167ffffffffffffffff82116112875761100382610fa485546134de565b602090601f83116001146113855761103392916000918361137a5750508160011b916000199060031b1c19161790565b90555b518051600484019167ffffffffffffffff82116112875761105b82610fa485546134de565b602090601f83116001146113135761108b92916000918361129d5750508160011b916000199060031b1c19161790565b90555b518051600583019167ffffffffffffffff8211611287576110b382610fa485546134de565b602090601f83116001146112a85791806110e892600696959460009261129d5750508160011b916000199060031b1c19161790565b90555b019360a0515194855167ffffffffffffffff8111611287576111178161111184546134de565b84613cb9565b6020601f821160011461121d57908061114b926111a198996000926112125750508160011b916000199060031b1c19161790565b90555b600052600a602052611164826040600020615171565b50604435600052600b60205281604060002055600052600e60205261118d816040600020615171565b5033600052600f6020526040600020615171565b506111ba6040519160608352606083019060e051613800565b6111c560643561323f565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a0933928061120d604435946101005161012051613800565b0390a3005b015190508980610fc6565b601f1982169783600052816000209860005b81811061126f5750916111a1989991846001959410611256575b505050811b01905561114e565b015160001960f88460031b161c19169055888080611249565b838301518b556001909a01996020938401930161122f565b634e487b7160e01b600052604160045260246000fd5b015190508a80610fc6565b90601f1983169184600052816000209260005b8181106112fb5750916001939185600698979694106112e2575b505050811b0190556110eb565b015160001960f88460031b161c191690558980806112d5565b929360206001819287860151815501950193016112bb565b90601f1983169184600052816000209260005b8181106113625750908460019594939210611349575b505050811b01905561108e565b015160001960f88460031b161c1916905589808061133c565b92936020600181928786015181550195019301611326565b015190508b80610fc6565b90601f1983169184600052816000209260005b8181106113d457509084600195949392106113bb575b505050811b019055611036565b015160001960f88460031b161c191690558a80806113ae565b92936020600181928786015181550195019301611398565b015190508c80610fc6565b90601f1983169184600052816000209260005b818110611446575090846001959493921061142d575b505050811b019055610fde565b015160001960f88460031b161c191690558b8080611420565b9293602060018192878601518155019501930161140a565b61147561146e368d60e0516135f7565b8a336148f6565b6114818c8a338d614df0565b610e02565b9a999897969594939291908a600052601060205260406000206080525b6080515480156114d057806000198101116107bb576107966114cb9160001901608051614eb6565b6114a3565b50909192939495969798999a610dee565b9294969850969a9450979150330361184e578454600052600b60205260006040812055604435600052600b6020526040600020558354956044358555600385019160405161153a816115338187613518565b03826131b4565b6020815191012061154c3684846135f7565b6020815191012003611787575b5050506004830191604051611572816115338187613518565b602081519101206115843684846135f7565b60208151910120036116b7575b5050506006019067ffffffffffffffff8111611287576115b58161111184546134de565b6000601f82116001146116535781906115e7939495966000926116485750508160011b916000199060031b1c19161790565b90555b6116026040519260408452604084019060e051613800565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e322339380611643604435956101005161012051613800565b0390a4005b013590508680610fc6565b601f198216958382526020822091805b88811061169f57508360019596979810611685575b505050811b0190556115ea565b0135600019600384901b60f8161c19169055858080611678565b90926020600181928686013581550194019101611663565b67ffffffffffffffff8211611287576116d482610fa485546134de565b600090601f831160011461171f5791806117099260069695946000926117145750508160011b916000199060031b1c19161790565b90555b908680611591565b013590508a80610fc6565b8382526020822091601f198416815b81811061176f575091600193918560069897969410611755575b505050811b01905561170c565b0135600019600384901b60f8161c19169055898080611748565b9193602060018192878701358155019501920161172e565b67ffffffffffffffff8211611287576117a482610fa485546134de565b600090601f83116001146117ea576117d49291600091836117df5750508160011b916000199060031b1c19161790565b90555b878080611559565b013590508b80610fc6565b8382526020822091601f198416815b818110611836575090846001959493921061181c575b505050811b0190556117d7565b0135600019600384901b60f8161c191690558a808061180f565b919360206001819287870135815501950192016117f9565b6331ee6dc760e01b6000523360045260246000fd5b7f36a7c503000000000000000000000000000000000000000000000000000000006000526101005160045260245260446000fd5b50806101005111610d5a565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b857f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808611610d40565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b867f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808711610d25565b808a1161199157808c1115610d13578b6219aad560e31b60005260045260245260446000fd5b896219aad560e31b60005260045260245260446000fd5b7f9cd963cf0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b7f315de7450000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ee5761039a611a3d36613423565b91613f43565b346102ee5760203660031901126102ee5760043567ffffffffffffffff81116102ee57611a749036906004016133bb565b801561069a5760005b818110611a8657005b80611a94600192848661384d565b35611a9d61472f565b600052600b602052604060002054806000526009602052604060002060ff8482015460e01c16611acc8161323f565b15611ada575b505001611a7d565b611ae391614d16565b8480611ad2565b346102ee5760a03660031901126102ee5760043560ff81168091036102ee576024359060ff8216908183036102ee576044359160ff8316908184036102ee576064359161ffff8316918284036102ee576084359663ffffffff8816948589036102ee5764ffff00000060a09862ff00007f9c1a161a4cdd9b19a46f9660eee21b6394dc5aa70fc9e093dbb36d2c1786d7739b611b8461472f565b896080604051611b938161317b565b8d81528960208201528a60408201528b6060820152015268ffffffff00000000006002549160281b169561ff0068ffffffff000000000019928d87199162ffffff191617169160081b1617169160101b16179160181b1617176002556040519485526020850152604084015260608301526080820152a1005b346102ee5760603660031901126102ee57611c256130ce565b60243567ffffffffffffffff81116102ee57611c45903690600401613211565b916044359267ffffffffffffffff84116102ee57611c8f611c9793610d7f611c74611cab973690600401613211565b90611c7d61363e565b50604051958694602086019889613f0e565b519020614601565b604051918291602083526020830190613249565b0390f35b346102ee5760003660031901126102ee576020600554604051908152f35b346102ee5760803660031901126102ee57611ce66130ce565b6064359067ffffffffffffffff82116102ee57611d0a61039a923690600401613211565b916044359060243590613d00565b346102ee5760403660031901126102ee57600435611d34613493565b63ffffffff60025460281c169063ffffffff811691428311801590611f1e575b611ee95750611d70336000526006602052604060002054151590565b1561039c57604080513360208201908152918101859052611d948160608101610d7f565b51902060008181526014602052604090205491908215611e19575060001982019182116107bb576001611dc9611deb936136dc565b50019063ffffffff60a01b1963ffffffff60a01b83549260a01b169116179055565b6040519081527ff69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d160203392a3005b91505060155490600182018092116107bb576000526014602052604060002055604051611e4581613143565b82815260208101338152604082018381526015546801000000000000000081101561128757806001611e7a92016015556136dc565b919091611ed3576001600160a01b0380600163ffffffff94611ece9751815501945116166001600160a01b0319845416178355511663ffffffff60a01b1963ffffffff60a01b83549260a01b169116179055565b611deb565b634e487b7160e01b600052600060045260246000fd5b82847f7ffd3b8f0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b508015158015611d54575080611f344285613abb565b11611d54565b346102ee5760003660031901126102ee5760206001600160a01b0360015416604051908152f35b346102ee5760603660031901126102ee57611f7a6130ce565b6001600160a01b0360243591169081600052600f602052611fa360443582604060002054614587565b90611fad8261368c565b9260005b838110611fc65760405180611cab8782613314565b60019082600052600f602052611ff9611fed6040600020611fe784886134d1565b90614eb6565b90549060031b1c614601565b612003828861359b565b5261200e818761359b565b5001611fb1565b346102ee576001600160a01b0361202b36613453565b92909161203661472f565b169081600052601660205260406000209267ffffffffffffffff81116112875761206a8161206486546134de565b86613cb9565b600093601f82116001146120e1576120bc82807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f1396976000916120d6575b508160011b916000199060031b1c19161790565b90555b610b9a604051928392602084526020840191613800565b9050850135886120a8565b80855260208520601f19831695805b87811061214b5750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13969710612131575b5050600182811b0190556120bf565b840135600019600385901b60f8161c191690558580612122565b909160206001819285890135815501930191016120f0565b346102ee5760003660031901126102ee576040601b5467ffffffffffffffff8251916001600160a01b038116835260a01c166020820152f35b346102ee5760003660031901126102ee576000546001600160a01b038116330361220b576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ee5760203660031901126102ee5761224e61363e565b50600435600052600b602052611cab611c97604060002054614601565b346102ee5760403660031901126102ee576122846130ce565b6001600160a01b036024359161229861472f565b16600052600c60205260406000209060005b825490811515806122dc575b1561039a5760001982019182116107bb576107b16107966122d79386614eb6565b6122aa565b508215806122b657508281106122b6565b346102ee5760203660031901126102ee57612315336000526006602052604060002054151590565b1561039c57600435600052600b60205261039a604060002054612338813361484e565b90614b7f565b346102ee5761237b61235e612352366133ec565b949290939136916135f7565b602081519101209283600052600e60205282604060002054614587565b906123858261368c565b9260005b83811061239e5760405180611cab8782613314565b60019082600052600e6020526123bf611fed6040600020611fe784886134d1565b6123c9828861359b565b526123d4818761359b565b5001612389565b346102ee576123e936613423565b9091612402336000526006602052604060002054151590565b1561039c57600052600b60205260406000205461241f813361484e565b9060ff600183015460e01c166124348161323f565b61243a57005b61246561245e61039a9561244f3687836135f7565b602081519101209536916135f7565b84336148f6565b614a5d565b346102ee5761248661247b366133ec565b9290939136916135f7565b6020815191012080600052601a6020526124a7604060002092848454614587565b916124b1836134b9565b936124bf60405195866131b4565b838552601f196124ce856134b9565b0160005b8181106125e757505060005b848110612546578560405180916020820160208352815180915260206040840192019060005b818110612512575050500390f35b825180516001600160a01b0316855260209081015163ffffffff168186015286955060409094019390920191600101612504565b806001600160a01b0361256461255e600194866134d1565b86614eb6565b90549060031b1c168560005260176020526040600281600020016000906001600160a01b0384168252602052206040519061259e8261315f565b5490602060ff63ffffffff841693848452821c161515910152604051916125c48361315f565b825260208201526125d5828961359b565b526125e0818861359b565b50016124de565b6020906040516125f68161315f565b6000815260008382015282828a010152016124d2565b346102ee5760003660031901126102ee576020600354604051908152f35b346102ee5760203660031901126102ee5760043567ffffffffffffffff81116102ee5761074a61265e913690600401613211565b60208151910120600052600d6020526020604060002054604051908152f35b346102ee5760203660031901126102ee5761269661472f565b600435600052600b602052604060002054806000526009602052604060002060ff600182015460e01c166126c98161323f565b1561039157005b346102ee576126ed6126e136613077565b81600393929354614587565b6126f6816135c5565b90600092600354935b8281106127145760405180611cab86826130e4565b61271e81836134d1565b60008682101561276d57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600191906001600160a01b0316612766828761359b565b52016126ff565b80634e487b7160e01b602492526032600452fd5b346102ee5761279861279236613077565b90613b3c565b60405180916020820160208352815180915260206040840192019060005b8181106127c4575050500390f35b919350916020606060019263ffffffff60408851805184526001600160a01b0386820151168685015201511660408201520194019101918493926127b6565b346102ee5761281136613374565b928291924211612930576001600160a01b0381169361283d856000526006602052604060002054151590565b1561291b5761284b85614ece565b936128646001600160a01b036106ea848489898961476d565b156128fe57858581600052600f6020526040600020905b815480156128bc5760001981019081116107bb5761289c6128b79184614eb6565b90549060031b1c806000526009602052604060002090614b7f565b61287b565b6000828580835260076020528260408120556128d7816151cd565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b60405163335d4ce160e01b8152948594610bc29460048701613821565b8463c2dda3f960e01b60005260045260246000fd5b6001600160a01b0390631ec5288b60e11b600052166004524260245260445260646000fd5b346102ee5760403660031901126102ee5760043567ffffffffffffffff81116102ee576129869036906004016133bb565b6024359167ffffffffffffffff83116102ee576129aa61039a933690600401613211565b92909161385d565b346102ee5760203660031901126102ee576001600160a01b036129d36130ce565b166000526016602052611cab6115336129f6604060002060405192838092613518565b60405191829160208352602083019061308d565b346102ee57612a1836613374565b8293924211612aa1576001600160a01b038316612a42816000526006602052604060002054151590565b15612a8d57612a5090614ece565b91612a696001600160a01b036106ea8484878a8a61476d565b15612a7057005b610bc29260405195869563335d4ce160e01b875260048701613821565b63c2dda3f960e01b60005260045260246000fd5b836001600160a01b0384631ec5288b60e11b600052166004524260245260445260646000fd5b346102ee5760403660031901126102ee57612ae06130ce565b60243567ffffffffffffffff811681036102ee5761039a91612b0061472f565b6136fb565b346102ee57612b22612b1636613077565b81601c93929354614587565b90612b2c826134b9565b91612b3a60405193846131b4565b808352601f19612b49826134b9565b0160005b818110612cc8575050601c54909160005b838110612c0a578460405160208101916020825280518093526040820192602060408260051b85010192019060005b818110612b9a5784840385f35b909192603f198582030186528351908151916003831015612bf457612be582606060406020959460019787965263ffffffff868201511686850152015191816040820152019061308d565b95019601910194919094612b8d565b634e487b7160e01b600052602160045260246000fd5b612c1481836134d1565b60008482101561276d57601c90526040519060011b7f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a211016000612c5683613143565b81549060ff8216906003821015612cb45750835260081c63ffffffff1660208301526040516001939291612c939082906115339082908801613518565b6040820152612ca2828861359b565b52612cad818761359b565b5001612b5e565b80634e487b7160e01b602492526021600452fd5b602090604051612cd781613143565b600081526000838201526060604082015282828801015201612b4d565b346102ee5760803660031901126102ee57612d0d6130ce565b60243567ffffffffffffffff81116102ee57612d4191612d34612d47923690600401613211565b93906044359436916135f7565b906145e2565b9081600052600a602052612d6360643582604060002054614587565b90612d6d8261368c565b9260005b838110612d865760405180611cab8782613314565b60019082600052600a602052612da7611fed6040600020611fe784886134d1565b612db1828861359b565b52612dbc818761359b565b5001612d71565b346102ee5760003660031901126102ee57611cab6129f66131d6565b346102ee57612dfc612df036613077565b81600593929354614587565b612e05816135c5565b90600092600554935b828110612e235760405180611cab86826130e4565b612e2d81836134d1565b60008682101561276d57600590527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0015460008190526007602052600191906001600160a01b0316612e7f828761359b565b5201612e0e565b346102ee5760203660031901126102ee5760206106fe6001600160a01b03612eac6130ce565b166000526006602052604060002054151590565b346102ee57612edd612ed136613077565b81601893929354614587565b612ee6816134b9565b91612ef460405193846131b4565b818352601f19612f03836134b9565b0160005b818110613044575050601854919060005b828110612fb657846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210612f5857505050500390f35b919360019193955060208091603f1989820301855287519081518152606063ffffffff81612f9386860151608088870152608086019061308d565b948260408201511660408601520151169101529601920192018594939192612f49565b612fc081836134d1565b60008582101561276d579060208260186001959452200163ffffffff6040600092549283815260176020522061153361301286830154926040519561300487613127565b865260405192838092613518565b6020840152818116604084015260201c166060820152613032828861359b565b5261303d818761359b565b5001612f18565b60209060405161305381613127565b60008152606083820152600060408201526000606082015282828801015201612f07565b60409060031901126102ee576004359060243590565b919082519283825260005b8481106130b9575050826000602080949584010152601f8019910116010190565b80602080928401015182828601015201613098565b600435906001600160a01b03821682036102ee57565b602060408183019282815284518094520192019060005b8181106131085750505090565b82516001600160a01b03168452602093840193909201916001016130fb565b6080810190811067ffffffffffffffff82111761128757604052565b6060810190811067ffffffffffffffff82111761128757604052565b6040810190811067ffffffffffffffff82111761128757604052565b60a0810190811067ffffffffffffffff82111761128757604052565b610140810190811067ffffffffffffffff82111761128757604052565b90601f8019910116810190811067ffffffffffffffff82111761128757604052565b604051906131e56040836131b4565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152565b9181601f840112156102ee5782359167ffffffffffffffff83116102ee57602083818601950101116102ee57565b60021115612bf457565b61331191815181526001600160a01b03602083015116602082015267ffffffffffffffff604083015116604082015260608201516132868161323f565b60608201526101206132ff6132eb6132d96132c76132b56080880151610140608089015261014088019061308d565b60a088015187820360a089015261308d565b60c087015186820360c088015261308d565b60e086015185820360e087015261308d565b61010085015184820361010086015261308d565b9201519061012081840391015261308d565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061334757505050505090565b9091929394602080613365600193603f198682030187528951613249565b97019301930191939290613338565b60606003198201126102ee576004356001600160a01b03811681036102ee5791602435916044359067ffffffffffffffff82116102ee576133b791600401613211565b9091565b9181601f840112156102ee5782359167ffffffffffffffff83116102ee576020808501948460051b0101116102ee57565b60606003198201126102ee576004359067ffffffffffffffff82116102ee5761341791600401613211565b90916024359060443590565b9060406003198301126102ee57600435916024359067ffffffffffffffff82116102ee576133b791600401613211565b9060406003198301126102ee576004356001600160a01b03811681036102ee57916024359067ffffffffffffffff82116102ee576133b791600401613211565b6024359063ffffffff821682036102ee57565b6044359063ffffffff821682036102ee57565b67ffffffffffffffff81116112875760051b60200190565b919082018092116107bb57565b90600182811c9216801561350e575b60208310146134f857565b634e487b7160e01b600052602260045260246000fd5b91607f16916134ed565b60009291815491613528836134de565b808352926001811690811561357e575060011461354457505050565b60009081526020812093945091925b838310613564575060209250010190565b600181602092949394548385870101520191019190613553565b915050602093945060ff929192191683830152151560051b010190565b80518210156135af5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906135cf826134b9565b6135dc60405191826131b4565b82815280926135ed601f19916134b9565b0190602036910137565b92919267ffffffffffffffff82116112875760405191613621601f8201601f1916602001846131b4565b8294818452818301116102ee578281602093846000960137010152565b6040519061364b82613197565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b90613696826134b9565b6136a360405191826131b4565b82815280926136b4601f19916134b9565b019060005b8281106136c557505050565b6020906136d061363e565b828285010152016136b9565b6015548110156135af57601560005260206000209060011b0190600090565b90601b54906001600160a01b038216906001600160a01b0367ffffffffffffffff8460a01c16941693828514948580966137ed575b6137e557806080957fc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b97156137d0575b505067ffffffffffffffff83169282840361378f575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b601b549260a01b16911617601b5538613776565b6001600160a01b03191617601b558038613760565b505050505050565b508167ffffffffffffffff841614613730565b908060209392818452848401376000828201840152601f01601f1916010190565b90926080926001600160a01b036133119795168352602083015260408201528160608201520191613800565b91908110156135af5760051b0190565b9291811561069a5761387c336000526006602052604060002054151590565b1561039c5761388c3684836135f7565b6020815191012092600091825b848110613a58575063ffffffff8316156137e5576138b89136916135f7565b90836000526017602052604060002090600182015463ffffffff81168015613a365786600052601260205263ffffffff6138f98482604060002054166148dc565b1611613a145763ffffffff60029160201c1692016001600160a01b0333166000528060205260ff60406000205460201c166139f4575b5063ffffffff61395c819233600052601160205260406000208860005260205282604060002054166148dc565b92169116116139d6575060005b8181106139765750505050565b80613984600192848761384d565b35600052600b6020526040600020548060005260096020528460406000208460ff8183015460e01c166139b68161323f565b146139c5575b50505001613969565b6139ce92614a5d565b3884816139bc565b610bc29060405191829163038857ff60e01b835233600484016145c2565b336000908152602091909152604090205463ffffffff908116925061392f565b60405163b993868760e01b81526020600482015280610bc2602482018761308d565b60405163393f328760e11b81526020600482015280610bc2602482018861308d565b92613a6484868961384d565b35600052600b60205260ff6001613a806040600020543361484e565b015460e01c16613a8f8161323f565b15613ab25763ffffffff1663ffffffff81146107bb576001809101935b01613899565b92600190613aac565b919082039182116107bb57565b90613ad2826134b9565b613adf60405191826131b4565b8281528092613af0601f19916134b9565b019060005b828110613b0157505050565b602090604051613b1081613143565b600081526000838201526000604082015282828501015201613af5565b60001981146107bb5760010190565b90613b4a9082601554614587565b908115613c3e57613b5a82613ac8565b91600091825b828110613bb75750508110613b73575090565b613b7c81613ac8565b9160005b828110613b8d5750505090565b80613b9a6001928461359b565b51613ba5828761359b565b52613bb0818661359b565b5001613b80565b613bc9613bc482846134d1565b6136dc565b50600181015463ffffffff8160a01c1690428211613bed575b505050600101613b60565b9160019396916001600160a01b03613c359460405193613c0c85613143565b5484521660208301526040820152613c24828961359b565b52613c2f818861359b565b50613b2d565b93903880613be2565b50506040516000613c506020836131b4565b81526000805b818110613c6257505090565b602090604051613c7181613143565b600081526000838201526000604082015282828601015201613c56565b80634e487b7160e01b602492526041600452fd5b818110613cad575050565b60008155600101613ca2565b9190601f8111613cc857505050565b613cf4926000526020600020906020601f840160051c83019310613cf6575b601f0160051c0190613ca2565b565b9091508190613ce7565b92909391844211613ecf576001600160a01b038416613d2c816000526006602052604060002054151590565b613ea25781600052600860205260ff60406000205416613e7257600090613d88613da2613d576131d6565b6040519283916020830195878752604084015246606084015230608084015260e060a084015261010083019061308d565b8a60c08301528660e083015203601f1981018352826131b4565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52613de3603c8220613ddd3687876135f7565b90614f31565b9091926004831015612cb45782613e3a575050506001600160a01b031660009081526004602052604090205415613e1b575050505050565b90610bc2929160405195869563335d4ce160e01b875260048701613821565b5060405163d36ab6b960e01b81526060600482015291829160ff613e62606485018a8a613800565b9216602484015260448301520390fd5b7f77a338580000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd9a5f5ca0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b846001600160a01b03857f502d038700000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b939161331195936001600160a01b03613f3593168652606060208701526060860191613800565b926040818503910152613800565b9091613f5c336000526006602052604060002054151590565b1561039c5781600052600b60205260406000205490613f7b823361484e565b60ff600182015460e01c16613f8f8161323f565b61405b57613fc49083600052601360205260406000205492836000526017602052613fcb604060002060405194858092613518565b03846131b4565b613fd63682896135f7565b602081519101208094146140525761402961403d94838761401b61404d967f9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d9a614d16565b61246561245e36878e6135f7565b60405193849360408552604085019061308d565b9083820360208501523397613800565b0390a3565b50505050505050565b7fd74915a80000000000000000000000000000000000000000000000000000000060005260046000fd5b604080516001600160a01b0390921660208301908152908201929092526140af8160608101610d7f565b519020600052601460205260406000205480156140ee5760001981019081116107bb5760016140e263ffffffff926136dc565b50015460a01c16421090565b50600090565b91906141019136916135f7565b6020815191012060005260176020526001600160a01b036040600020911660005260028101602052604060002060206040519161413d8361315f565b549160ff63ffffffff841693848352831c161515918291015261416b57506001015460201c63ffffffff1690565b905090565b90601c54680100000000000000008110156112875760018101601c556000601c5482101561276d57601c90526000929060011b7f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a21101815160038110156142c25760409160019160ff825491168060ff19831617835564ffffffff00602087015160081b169164ffffffffff191617178155019101519283519067ffffffffffffffff8211613c8e5761422682610fa485546134de565b602090601f83116001146142645790614255939495836142595750508160011b916000199060031b1c19161790565b9055565b015190503880610fc6565b90601f198316848352818320925b8181106142aa57509583600195969710614291575b505050811b019055565b015160001960f88460031b161c19169055388080614287565b9192602060018192868b015181550194019201614272565b602485634e487b7160e01b81526021600452fd5b919392936142e53683856135f7565b60208151910120948560005260176020526040600020956001870163ffffffff81549416938463ffffffff821614908161455d575b506145535763ffffffff831692848411614529578854614339816134de565b15614402575b507fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d979850906143db92918563ffffffff1983541617825567ffffffff0000000082549160201b169067ffffffff000000001916179055604051906020820152836040820152604081526143b46060826131b4565b604051906143c182613143565b6000825263ffffffff421660208301526040820152614170565b6143f2604051948594606086526060860191613800565b91602084015260408301520390a1565b67ffffffffffffffff8711611287578661441e614424926134de565b8b613cb9565b600098601f871160011461449f57906143db93929161447c88807fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9c9d60009161449457508160011b916000199060031b1c19161790565b90555b61448883615117565b5090919289985061433f565b90508b0135386120a8565b808a5260208a20601f1988168b5b8181106145115750907fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9a9b896143db9796959493106144f7575b5050600188811b01905561447f565b8a013560001960038b901b60f8161c1916905538806144e8565b8a8d0135835560209c8d019c600190930192016144ad565b7fec623c5f0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050505050509050565b905063ffffffff8085169160201c16143861431a565b356001600160a01b03811681036102ee5790565b808210156145ba57828161459e61331195856134d1565b11156145aa5750613abb565b6145b59150826134d1565b613abb565b505050600090565b6040906001600160a01b036133119493168152816020820152019061308d565b906145fb610d7f916040519283916020830195866145c2565b51902090565b61460961363e565b81600052600960205260406000206001810154926001600160a01b038416928315614727575091600661470c836115339567ffffffffffffffff61471e96549860ff8160e01c1692600052601360205260406000205460005260176020526040600020966040519a61467a8c613197565b8b5260208b015260a01c1660408901526146938161323f565b60608801526040516146ac816115338160028601613518565b60808801526040516146c5816115338160038601613518565b60a08801526040516146de816115338160048601613518565b60c08801526040516146f7816115338160058601613518565b60e08801526115336040518094819301613518565b61010085015260405192838092613518565b61012082015290565b935050505090565b6001600160a01b0360015416330361474357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b916147d3906147ba9261477e6131d6565b916040519485936001600160a01b03602086019860018a5216604086015246606086015230608086015260e060a086015261010085019061308d565b9160c084015260e083015203601f1981018352826131b4565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52614810603c600020613ddd3685856135f7565b6004829593951015612bf45781614828575050505090565b60ff613e6260405195869563d36ab6b960e01b8752606060048801526064870191613800565b9060005260096020526040600020906001600160a01b0360018301541690811561489a576001600160a01b0316809103614886575090565b6331ee6dc760e01b60005260045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff60019116019063ffffffff82116107bb57565b9063ffffffff8091169116019063ffffffff82116107bb57565b90806000526017602052604060002090600182015463ffffffff81168015614a165782600052601260205263ffffffff614938600182604060002054166148dc565b16116149f45763ffffffff60029160201c1692016001600160a01b0384166000528060205260ff60406000205460201c166149ce575b506001600160a01b038316600052601160205260406000209060005260205263ffffffff806149a5600182604060002054166148dc565b92169116116149b2575050565b610bc260405192839263038857ff60e01b8452600484016145c2565b9091506001600160a01b03831660005260205263ffffffff60406000205416903861496e565b60405163b993868760e01b81526020600482015280610bc2602482018861308d565b60405163393f328760e11b81526020600482015280610bc2602482018961308d565b9091614a4f61331193604084526040840190613518565b916020818403910152613518565b7f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a906001600160a01b036001840193614aba8286541693876002840195614ab4604051614aae81611533818c613518565b826145e2565b92614df0565b60ff60e01b198554168555614b0f815460405190886020830152604082015260408152614ae86060826131b4565b60405190614af582613143565b6001825263ffffffff421660208301526040820152614170565b54935416936000526017602052604060002061404d60405192839283614a38565b614b3a81546134de565b9081614b44575050565b81601f60009311600114614b56575055565b81835260208320614b7291601f0160051c810190600101613ca2565b8082528160208120915555565b90614aae907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94614c9e614c0d614cac6001600160a01b036001860195865460ff8a848316614bdb611533600287019d8e60405192838092613518565b938491836000526013602052604060002054948592836000526017602052614c1460406000206040519d8e8092613518565b038d6131b4565b60e01c16614c218161323f565b15614d04575b50505050600052600e602052614c418a604060002061534b565b5082885416600052600f602052614c5c8a604060002061534b565b50600052600a602052614c7389604060002061534b565b508054600052600b60205260006040812055549554169560405193849360408552604085019061308d565b908382036020850152613518565b0390a36000526009602052600660406000206000815560006001820155614cd560028201614b30565b614ce160038201614b30565b614ced60048201614b30565b614cf960058201614b30565b01614b3a81546134de565b614d0d93614f86565b8a828238614c27565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e48868960018301916001600160a01b038354947c010000000000000000000000000000000000000000000000000000000060ff60e01b198716178555826000526013602052614da682604060002054971693876002840195614da0604051614aae81611533818c613518565b92614f86565b614b0f815460405190886020830152604082015260408152614dc96060826131b4565b60405190614dd682613143565b6002825263ffffffff421660208301526040820152614170565b9192906001600160a01b0316806000526011602052604060002084600052602052604060002063ffffffff614e27818354166148c4565b1663ffffffff19825416179055836000526012602052604060002063ffffffff614e53818354166148c4565b1663ffffffff1982541617905583600052600d602052614e77836040600020615171565b50600052600c602052614e8e826040600020615171565b506000526010602052614ea5816040600020615171565b506000526013602052604060002055565b80548210156135af5760005260206000200190600090565b80600052600760205260406000205490811580614f1b575b614eee575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5080600052600660205260406000205415614ee6565b8151919060418303614f6257614f5b92506020820151906060604084015193015160001a90615408565b9192909190565b505060009160029190565b63ffffffff6000199116019063ffffffff82116107bb57565b9291906001600160a01b031680600052600c602052614fa984604060002061534b565b5081600052600d602052614fc184604060002061534b565b506000526011602052604060002081600052602052604060002063ffffffff614fec81835416614f6d565b1663ffffffff198254161790556000526012602052604060002063ffffffff61501781835416614f6d565b1663ffffffff19825416179055600052601060205261503a81604060002061534b565b50600052601360205260006040812055565b806000526004602052604060002054156000146140ee5760035468010000000000000000811015611287576150a461508d8260018594016003556003614eb6565b819391549060031b91821b91600019901b19161790565b9055600354906000526004602052604060002055600190565b806000526006602052604060002054156000146140ee5760055468010000000000000000811015611287576150fe61508d8260018594016005556005614eb6565b9055600554906000526006602052604060002055600190565b806000526019602052604060002054156000146140ee57601854680100000000000000008110156112875761515861508d8260018594016018556018614eb6565b9055601854906000526019602052604060002055600190565b60008281526001820160205260409020546151c6578054906801000000000000000082101561128757826151af61508d846001809601855584614eb6565b905580549260005201602052604060002055600190565b5050600090565b60008181526006602052604090205480156151c65760001981018181116107bb576005546000198101919082116107bb5781810361525d575b50505060055480156152475760001901615221816005614eb6565b8154906000199060031b1b19169055600555600052600660205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61527f61526e61508d936005614eb6565b90549060031b1c9283926005614eb6565b90556000526006602052604060002055388080615206565b60008181526004602052604090205480156151c65760001981018181116107bb576003546000198101919082116107bb57818103615311575b505050600354801561524757600019016152eb816003614eb6565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b61533361532261508d936003614eb6565b90549060031b1c9283926003614eb6565b905560005260046020526040600020553880806152d0565b90600182019181600052826020526040600020548015156000146153ff5760001981018181116107bb5782546000198101919082116107bb578181036153c8575b505050805480156152475760001901906153a68282614eb6565b8154906000199060031b1b191690555560005260205260006040812055600190565b6153e86153d861508d9386614eb6565b90549060031b1c92839286614eb6565b90556000528360205260406000205538808061538c565b50505050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411615491579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15615485576000516001600160a01b038116156154795790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) CanLinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canLinkOwner", owner, validityTimestamp, proof, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanLinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanLinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canUnlinkOwner", owner, validityTimestamp, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetAllowedSigners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getAllowedSigners", start, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetAllowedSigners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetAllowedSigners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetAllowedSigners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetAllowedSigners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetAllowlistedRequests(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getAllowlistedRequests", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryOwnerAllowlistedRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryOwnerAllowlistedRequest)).(*[]WorkflowRegistryOwnerAllowlistedRequest)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetAllowlistedRequests(start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	return _WorkflowRegistry.Contract.GetAllowlistedRequests(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetAllowlistedRequests(start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error) {
	return _WorkflowRegistry.Contract.GetAllowlistedRequests(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetCapabilitiesRegistry(opts *bind.CallOpts) (common.Address, uint64, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getCapabilitiesRegistry")

	if err != nil {
		return *new(common.Address), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetCapabilitiesRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetCapabilitiesRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetCapabilitiesRegistry() (common.Address, uint64, error) {
	return _WorkflowRegistry.Contract.GetCapabilitiesRegistry(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetConfig(opts *bind.CallOpts) (WorkflowRegistryConfig, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(WorkflowRegistryConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryConfig)).(*WorkflowRegistryConfig)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetConfig() (WorkflowRegistryConfig, error) {
	return _WorkflowRegistry.Contract.GetConfig(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetConfig() (WorkflowRegistryConfig, error) {
	return _WorkflowRegistry.Contract.GetConfig(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetDonConfigs(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getDonConfigs", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryDonConfigView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryDonConfigView)).(*[]WorkflowRegistryDonConfigView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetDonConfigs(start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	return _WorkflowRegistry.Contract.GetDonConfigs(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetDonConfigs(start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error) {
	return _WorkflowRegistry.Contract.GetDonConfigs(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getEvents", start, limit)

	if err != nil {
		return *new([]WorkflowRegistryEventRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryEventRecord)).(*[]WorkflowRegistryEventRecord)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetEvents(start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	return _WorkflowRegistry.Contract.GetEvents(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetEvents(start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error) {
	return _WorkflowRegistry.Contract.GetEvents(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getLinkedOwners", start, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetLinkedOwners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetLinkedOwners(start *big.Int, limit *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (GetMaxWorkflowsPerDON,

	error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donFamily)

	outstruct := new(GetMaxWorkflowsPerDON)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxWorkflows = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DefaultUserLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donFamily string) (GetMaxWorkflowsPerDON,

	error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donFamily string) (GetMaxWorkflowsPerDON,

	error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donFamily)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donFamily string) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetUserDONOverrides(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getUserDONOverrides", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryUserOverrideView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryUserOverrideView)).(*[]WorkflowRegistryUserOverrideView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetUserDONOverrides(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	return _WorkflowRegistry.Contract.GetUserDONOverrides(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetUserDONOverrides(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error) {
	return _WorkflowRegistry.Contract.GetUserDONOverrides(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflow(opts *bind.CallOpts, owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflow", owner, workflowName, tag)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadataView)).(*WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflow(owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflow(&_WorkflowRegistry.CallOpts, owner, workflowName, tag)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflow(owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflow(&_WorkflowRegistry.CallOpts, owner, workflowName, tag)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowById(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowById", workflowId)

	if err != nil {
		return *new(WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new(WorkflowRegistryWorkflowMetadataView)).(*WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowById(workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowById(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowById(workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowById(&_WorkflowRegistry.CallOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByDON", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByOwner", owner, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByOwner(owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwner(&_WorkflowRegistry.CallOpts, owner, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowListByOwnerAndName", owner, workflowName, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowListByOwnerAndName(owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetWorkflowListByOwnerAndName(&_WorkflowRegistry.CallOpts, owner, workflowName, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getWorkflowOwnerConfig", owner)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetWorkflowOwnerConfig(owner common.Address) ([]byte, error) {
	return _WorkflowRegistry.Contract.GetWorkflowOwnerConfig(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetWorkflowOwnerConfig(owner common.Address) ([]byte, error) {
	return _WorkflowRegistry.Contract.GetWorkflowOwnerConfig(&_WorkflowRegistry.CallOpts, owner)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) IsRequestAllowlisted(opts *bind.CallOpts, owner common.Address, requestDigest [32]byte) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isRequestAllowlisted", owner, requestDigest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsRequestAllowlisted(owner common.Address, requestDigest [32]byte) (bool, error) {
	return _WorkflowRegistry.Contract.IsRequestAllowlisted(&_WorkflowRegistry.CallOpts, owner, requestDigest)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsRequestAllowlisted(owner common.Address, requestDigest [32]byte) (bool, error) {
	return _WorkflowRegistry.Contract.IsRequestAllowlisted(&_WorkflowRegistry.CallOpts, owner, requestDigest)
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

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalActiveWorkflowsByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalActiveWorkflowsByOwner", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalActiveWorkflowsByOwner(owner common.Address) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsByOwner(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalActiveWorkflowsByOwner(owner common.Address) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsByOwner(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalActiveWorkflowsOnDON(opts *bind.CallOpts, donFamily string) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalActiveWorkflowsOnDON", donFamily)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalActiveWorkflowsOnDON(donFamily string) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsOnDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalActiveWorkflowsOnDON(donFamily string) (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalActiveWorkflowsOnDON(&_WorkflowRegistry.CallOpts, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalAllowedSigners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalAllowedSigners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalAllowedSigners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowedSigners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalAllowedSigners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowedSigners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalAllowlistedRequests(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalAllowlistedRequests")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalAllowlistedRequests() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowlistedRequests(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalAllowlistedRequests() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalAllowlistedRequests(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalEvents(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalEvents")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalEvents() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalEvents(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalEvents() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalEvents(&_WorkflowRegistry.CallOpts)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "activateWorkflow", workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) ActivateWorkflow(workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) ActivateWorkflow(workflowId [32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.ActivateWorkflow(&_WorkflowRegistry.TransactOpts, workflowId, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminBatchPauseWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminBatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminBatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminBatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminBatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByDON(opts *bind.TransactOpts, donFamily string, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByDON", donFamily, limit)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByDON(donFamily string, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donFamily, limit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByDON(donFamily string, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByDON(&_WorkflowRegistry.TransactOpts, donFamily, limit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseAllByOwner", owner, limit)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseAllByOwner(owner common.Address, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByOwner(&_WorkflowRegistry.TransactOpts, owner, limit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseAllByOwner(owner common.Address, limit *big.Int) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseAllByOwner(&_WorkflowRegistry.TransactOpts, owner, limit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "adminPauseWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AdminPauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AdminPauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AdminPauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "allowlistRequest", requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistrySession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AllowlistRequest(&_WorkflowRegistry.TransactOpts, requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AllowlistRequest(requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AllowlistRequest(&_WorkflowRegistry.TransactOpts, requestDigest, expiryTimestamp)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchActivateWorkflows", workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchActivateWorkflows(workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchActivateWorkflows(workflowIds [][32]byte, donFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchActivateWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds, donFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "batchPauseWorkflows", workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistrySession) BatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) BatchPauseWorkflows(workflowIds [][32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.BatchPauseWorkflows(&_WorkflowRegistry.TransactOpts, workflowIds)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "deleteWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) DeleteWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) DeleteWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.DeleteWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "pauseWorkflow", workflowId)
}

func (_WorkflowRegistry *WorkflowRegistrySession) PauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) PauseWorkflow(workflowId [32]byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.PauseWorkflow(&_WorkflowRegistry.TransactOpts, workflowId)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetCapabilitiesRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setCapabilitiesRegistry", registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetCapabilitiesRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetCapabilitiesRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetCapabilitiesRegistry(registry common.Address, chainSelector uint64) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetCapabilitiesRegistry(&_WorkflowRegistry.TransactOpts, registry, chainSelector)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setConfig", nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetConfig(nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetConfig(&_WorkflowRegistry.TransactOpts, nameLen, tagLen, urlLen, attrLen, expiryLen)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONLimit(opts *bind.TransactOpts, donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONLimit", donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONLimit(donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONLimit(donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONLimit(&_WorkflowRegistry.TransactOpts, donFamily, donLimit, userDefaultLimit)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donFamily, userLimit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, userLimit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donFamily, userLimit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setWorkflowOwnerConfig", owner, config)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetWorkflowOwnerConfig(owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetWorkflowOwnerConfig(&_WorkflowRegistry.TransactOpts, owner, config)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetWorkflowOwnerConfig(owner common.Address, config []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetWorkflowOwnerConfig(&_WorkflowRegistry.TransactOpts, owner, config)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "unlinkOwner", owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, signature)
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

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateWorkflowDONFamily(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateWorkflowDONFamily", workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateWorkflowDONFamily(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONFamily(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateWorkflowDONFamily(workflowId [32]byte, newDonFamily string) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateWorkflowDONFamily(&_WorkflowRegistry.TransactOpts, workflowId, newDonFamily)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "upsertWorkflow", workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpsertWorkflow(workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpsertWorkflow(&_WorkflowRegistry.TransactOpts, workflowName, tag, workflowId, status, donFamily, binaryUrl, configUrl, attributes, keepAlive)
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

type WorkflowRegistryCapabilitiesRegistryUpdatedIterator struct {
	Event *WorkflowRegistryCapabilitiesRegistryUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryCapabilitiesRegistryUpdated)
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
		it.Event = new(WorkflowRegistryCapabilitiesRegistryUpdated)
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

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryCapabilitiesRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryCapabilitiesRegistryUpdated struct {
	OldAddr          common.Address
	NewAddr          common.Address
	OldChainSelector uint64
	NewChainSelector uint64
	Raw              types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterCapabilitiesRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryCapabilitiesRegistryUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "CapabilitiesRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryCapabilitiesRegistryUpdatedIterator{contract: _WorkflowRegistry.contract, event: "CapabilitiesRegistryUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchCapabilitiesRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryCapabilitiesRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "CapabilitiesRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryCapabilitiesRegistryUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "CapabilitiesRegistryUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseCapabilitiesRegistryUpdated(log types.Log) (*WorkflowRegistryCapabilitiesRegistryUpdated, error) {
	event := new(WorkflowRegistryCapabilitiesRegistryUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "CapabilitiesRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryConfigUpdatedIterator struct {
	Event *WorkflowRegistryConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryConfigUpdated)
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
		it.Event = new(WorkflowRegistryConfigUpdated)
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

func (it *WorkflowRegistryConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryConfigUpdated struct {
	MaxNameLen   uint8
	MaxTagLen    uint8
	MaxUrlLen    uint8
	MaxAttrLen   uint16
	MaxExpiryLen uint32
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryConfigUpdatedIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryConfigUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseConfigUpdated(log types.Log) (*WorkflowRegistryConfigUpdated, error) {
	event := new(WorkflowRegistryConfigUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryDONLimitSetIterator struct {
	Event *WorkflowRegistryDONLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryDONLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryDONLimitSet)
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
		it.Event = new(WorkflowRegistryDONLimitSet)
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

func (it *WorkflowRegistryDONLimitSetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryDONLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryDONLimitSet struct {
	DonFamily        string
	DonLimit         uint32
	UserDefaultLimit uint32
	Raw              types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "DONLimitSet")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "DONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "DONLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryDONLimitSet)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "DONLimitSet", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error) {
	event := new(WorkflowRegistryDONLimitSet)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "DONLimitSet", log); err != nil {
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

type WorkflowRegistryRequestAllowlistedIterator struct {
	Event *WorkflowRegistryRequestAllowlisted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryRequestAllowlistedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryRequestAllowlisted)
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
		it.Event = new(WorkflowRegistryRequestAllowlisted)
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

func (it *WorkflowRegistryRequestAllowlistedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryRequestAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryRequestAllowlisted struct {
	Owner           common.Address
	RequestDigest   [32]byte
	ExpiryTimestamp uint32
	Raw             types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterRequestAllowlisted(opts *bind.FilterOpts, owner []common.Address, requestDigest [][32]byte) (*WorkflowRegistryRequestAllowlistedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestDigestRule []interface{}
	for _, requestDigestItem := range requestDigest {
		requestDigestRule = append(requestDigestRule, requestDigestItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "RequestAllowlisted", ownerRule, requestDigestRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryRequestAllowlistedIterator{contract: _WorkflowRegistry.contract, event: "RequestAllowlisted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchRequestAllowlisted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryRequestAllowlisted, owner []common.Address, requestDigest [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var requestDigestRule []interface{}
	for _, requestDigestItem := range requestDigest {
		requestDigestRule = append(requestDigestRule, requestDigestItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "RequestAllowlisted", ownerRule, requestDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryRequestAllowlisted)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "RequestAllowlisted", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseRequestAllowlisted(log types.Log) (*WorkflowRegistryRequestAllowlisted, error) {
	event := new(WorkflowRegistryRequestAllowlisted)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "RequestAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryUserDONLimitSetIterator struct {
	Event *WorkflowRegistryUserDONLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryUserDONLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryUserDONLimitSet)
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
		it.Event = new(WorkflowRegistryUserDONLimitSet)
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

func (it *WorkflowRegistryUserDONLimitSetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryUserDONLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryUserDONLimitSet struct {
	User      common.Address
	DonFamily string
	Limit     uint32
	Raw       types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitSetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitSet", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryUserDONLimitSet)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitSet", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseUserDONLimitSet(log types.Log) (*WorkflowRegistryUserDONLimitSet, error) {
	event := new(WorkflowRegistryUserDONLimitSet)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryUserDONLimitUnsetIterator struct {
	Event *WorkflowRegistryUserDONLimitUnset

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryUserDONLimitUnset)
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
		it.Event = new(WorkflowRegistryUserDONLimitUnset)
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

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryUserDONLimitUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryUserDONLimitUnset struct {
	User      common.Address
	DonFamily string
	Raw       types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitUnsetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "UserDONLimitUnset", userRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryUserDONLimitUnsetIterator{contract: _WorkflowRegistry.contract, event: "UserDONLimitUnset", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "UserDONLimitUnset", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryUserDONLimitUnset)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitUnset", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseUserDONLimitUnset(log types.Log) (*WorkflowRegistryUserDONLimitUnset, error) {
	event := new(WorkflowRegistryUserDONLimitUnset)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "UserDONLimitUnset", log); err != nil {
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
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowActivatedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowActivatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowActivated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowActivated", workflowIdRule, ownerRule)
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
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDeletedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDeletedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDeleted", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDeleted", workflowIdRule, ownerRule)
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

type WorkflowRegistryWorkflowDonFamilyUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowDonFamilyUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowDonFamilyUpdated)
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
		it.Event = new(WorkflowRegistryWorkflowDonFamilyUpdated)
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

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowDonFamilyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowDonFamilyUpdated struct {
	WorkflowId   [32]byte
	Owner        common.Address
	OldDonFamily string
	NewDonFamily string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowDonFamilyUpdated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDonFamilyUpdatedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowDonFamilyUpdated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowDonFamilyUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowDonFamilyUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowDonFamilyUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDonFamilyUpdated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowDonFamilyUpdated", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowDonFamilyUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDonFamilyUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowDonFamilyUpdated(log types.Log) (*WorkflowRegistryWorkflowDonFamilyUpdated, error) {
	event := new(WorkflowRegistryWorkflowDonFamilyUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowDonFamilyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryWorkflowOwnerConfigUpdatedIterator struct {
	Event *WorkflowRegistryWorkflowOwnerConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryWorkflowOwnerConfigUpdated)
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
		it.Event = new(WorkflowRegistryWorkflowOwnerConfigUpdated)
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

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryWorkflowOwnerConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryWorkflowOwnerConfigUpdated struct {
	Owner  common.Address
	Config []byte
	Raw    types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowOwnerConfigUpdated(opts *bind.FilterOpts, owner []common.Address) (*WorkflowRegistryWorkflowOwnerConfigUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowOwnerConfigUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowOwnerConfigUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowOwnerConfigUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowOwnerConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowOwnerConfigUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowOwnerConfigUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryWorkflowOwnerConfigUpdated)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowOwnerConfigUpdated", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseWorkflowOwnerConfigUpdated(log types.Log) (*WorkflowRegistryWorkflowOwnerConfigUpdated, error) {
	event := new(WorkflowRegistryWorkflowOwnerConfigUpdated)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "WorkflowOwnerConfigUpdated", log); err != nil {
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
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowPausedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowPausedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowPaused", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowPaused", workflowIdRule, ownerRule)
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
	WorkflowId   [32]byte
	Owner        common.Address
	DonFamily    string
	Status       uint8
	WorkflowName string
	Raw          types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowRegistered(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowRegistered", workflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowRegisteredIterator{contract: _WorkflowRegistry.contract, event: "WorkflowRegistered", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowRegistered", workflowIdRule, ownerRule)
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
	OldWorkflowId [32]byte
	NewWorkflowId [32]byte
	Owner         common.Address
	DonFamily     string
	WorkflowName  string
	Raw           types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error) {

	var oldWorkflowIdRule []interface{}
	for _, oldWorkflowIdItem := range oldWorkflowId {
		oldWorkflowIdRule = append(oldWorkflowIdRule, oldWorkflowIdItem)
	}
	var newWorkflowIdRule []interface{}
	for _, newWorkflowIdItem := range newWorkflowId {
		newWorkflowIdRule = append(newWorkflowIdRule, newWorkflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "WorkflowUpdated", oldWorkflowIdRule, newWorkflowIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryWorkflowUpdatedIterator{contract: _WorkflowRegistry.contract, event: "WorkflowUpdated", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var oldWorkflowIdRule []interface{}
	for _, oldWorkflowIdItem := range oldWorkflowId {
		oldWorkflowIdRule = append(oldWorkflowIdRule, oldWorkflowIdItem)
	}
	var newWorkflowIdRule []interface{}
	for _, newWorkflowIdItem := range newWorkflowId {
		newWorkflowIdRule = append(newWorkflowIdRule, newWorkflowIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "WorkflowUpdated", oldWorkflowIdRule, newWorkflowIdRule, ownerRule)
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

type GetMaxWorkflowsPerDON struct {
	MaxWorkflows     uint32
	DefaultUserLimit uint32
}

func (_WorkflowRegistry *WorkflowRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WorkflowRegistry.abi.Events["AllowedSignersUpdated"].ID:
		return _WorkflowRegistry.ParseAllowedSignersUpdated(log)
	case _WorkflowRegistry.abi.Events["CapabilitiesRegistryUpdated"].ID:
		return _WorkflowRegistry.ParseCapabilitiesRegistryUpdated(log)
	case _WorkflowRegistry.abi.Events["ConfigUpdated"].ID:
		return _WorkflowRegistry.ParseConfigUpdated(log)
	case _WorkflowRegistry.abi.Events["DONLimitSet"].ID:
		return _WorkflowRegistry.ParseDONLimitSet(log)
	case _WorkflowRegistry.abi.Events["OwnershipLinkUpdated"].ID:
		return _WorkflowRegistry.ParseOwnershipLinkUpdated(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferRequested(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferred"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferred(log)
	case _WorkflowRegistry.abi.Events["RequestAllowlisted"].ID:
		return _WorkflowRegistry.ParseRequestAllowlisted(log)
	case _WorkflowRegistry.abi.Events["UserDONLimitSet"].ID:
		return _WorkflowRegistry.ParseUserDONLimitSet(log)
	case _WorkflowRegistry.abi.Events["UserDONLimitUnset"].ID:
		return _WorkflowRegistry.ParseUserDONLimitUnset(log)
	case _WorkflowRegistry.abi.Events["WorkflowActivated"].ID:
		return _WorkflowRegistry.ParseWorkflowActivated(log)
	case _WorkflowRegistry.abi.Events["WorkflowDeleted"].ID:
		return _WorkflowRegistry.ParseWorkflowDeleted(log)
	case _WorkflowRegistry.abi.Events["WorkflowDonFamilyUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowDonFamilyUpdated(log)
	case _WorkflowRegistry.abi.Events["WorkflowOwnerConfigUpdated"].ID:
		return _WorkflowRegistry.ParseWorkflowOwnerConfigUpdated(log)
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

func (WorkflowRegistryCapabilitiesRegistryUpdated) Topic() common.Hash {
	return common.HexToHash("0xc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b")
}

func (WorkflowRegistryConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x9c1a161a4cdd9b19a46f9660eee21b6394dc5aa70fc9e093dbb36d2c1786d773")
}

func (WorkflowRegistryDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0xe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d")
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

func (WorkflowRegistryRequestAllowlisted) Topic() common.Hash {
	return common.HexToHash("0xf69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d1")
}

func (WorkflowRegistryUserDONLimitSet) Topic() common.Hash {
	return common.HexToHash("0x945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c82324")
}

func (WorkflowRegistryUserDONLimitUnset) Topic() common.Hash {
	return common.HexToHash("0x6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d849")
}

func (WorkflowRegistryWorkflowActivated) Topic() common.Hash {
	return common.HexToHash("0x7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a")
}

func (WorkflowRegistryWorkflowDeleted) Topic() common.Hash {
	return common.HexToHash("0x48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd94")
}

func (WorkflowRegistryWorkflowDonFamilyUpdated) Topic() common.Hash {
	return common.HexToHash("0x9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d")
}

func (WorkflowRegistryWorkflowOwnerConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f13")
}

func (WorkflowRegistryWorkflowPaused) Topic() common.Hash {
	return common.HexToHash("0xf764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e488689")
}

func (WorkflowRegistryWorkflowRegistered) Topic() common.Hash {
	return common.HexToHash("0x74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09")
}

func (WorkflowRegistryWorkflowUpdated) Topic() common.Hash {
	return common.HexToHash("0x03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e322")
}

func (_WorkflowRegistry *WorkflowRegistry) Address() common.Address {
	return _WorkflowRegistry.address
}

type WorkflowRegistryInterface interface {
	CanLinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) error

	GetAllowedSigners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetAllowlistedRequests(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryOwnerAllowlistedRequest, error)

	GetCapabilitiesRegistry(opts *bind.CallOpts) (common.Address, uint64, error)

	GetConfig(opts *bind.CallOpts) (WorkflowRegistryConfig, error)

	GetDonConfigs(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryDonConfigView, error)

	GetEvents(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]WorkflowRegistryEventRecord, error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, limit *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donFamily string) (GetMaxWorkflowsPerDON,

		error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donFamily string) (uint32, error)

	GetUserDONOverrides(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryUserOverrideView, error)

	GetWorkflow(opts *bind.CallOpts, owner common.Address, workflowName string, tag string) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowById(opts *bind.CallOpts, workflowId [32]byte) (WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwner(opts *bind.CallOpts, owner common.Address, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowListByOwnerAndName(opts *bind.CallOpts, owner common.Address, workflowName string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

	GetWorkflowOwnerConfig(opts *bind.CallOpts, owner common.Address) ([]byte, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	IsRequestAllowlisted(opts *bind.CallOpts, owner common.Address, requestDigest [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalActiveWorkflowsByOwner(opts *bind.CallOpts, owner common.Address) (*big.Int, error)

	TotalActiveWorkflowsOnDON(opts *bind.CallOpts, donFamily string) (*big.Int, error)

	TotalAllowedSigners(opts *bind.CallOpts) (*big.Int, error)

	TotalAllowlistedRequests(opts *bind.CallOpts) (*big.Int, error)

	TotalEvents(opts *bind.CallOpts) (*big.Int, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateWorkflow(opts *bind.TransactOpts, workflowId [32]byte, donFamily string) (*types.Transaction, error)

	AdminBatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	AdminPauseAllByDON(opts *bind.TransactOpts, donFamily string, limit *big.Int) (*types.Transaction, error)

	AdminPauseAllByOwner(opts *bind.TransactOpts, owner common.Address, limit *big.Int) (*types.Transaction, error)

	AdminPauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	AllowlistRequest(opts *bind.TransactOpts, requestDigest [32]byte, expiryTimestamp uint32) (*types.Transaction, error)

	BatchActivateWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte, donFamily string) (*types.Transaction, error)

	BatchPauseWorkflows(opts *bind.TransactOpts, workflowIds [][32]byte) (*types.Transaction, error)

	DeleteWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	PauseWorkflow(opts *bind.TransactOpts, workflowId [32]byte) (*types.Transaction, error)

	SetCapabilitiesRegistry(opts *bind.TransactOpts, registry common.Address, chainSelector uint64) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, nameLen uint8, tagLen uint8, urlLen uint8, attrLen uint16, expiryLen uint32) (*types.Transaction, error)

	SetDONLimit(opts *bind.TransactOpts, donFamily string, donLimit uint32, userDefaultLimit uint32) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donFamily string, userLimit uint32, enabled bool) (*types.Transaction, error)

	SetWorkflowOwnerConfig(opts *bind.TransactOpts, owner common.Address, config []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, signature []byte) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	UpdateWorkflowDONFamily(opts *bind.TransactOpts, workflowId [32]byte, newDonFamily string) (*types.Transaction, error)

	UpsertWorkflow(opts *bind.TransactOpts, workflowName string, tag string, workflowId [32]byte, status uint8, donFamily string, binaryUrl string, configUrl string, attributes []byte, keepAlive bool) (*types.Transaction, error)

	FilterAllowedSignersUpdated(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedIterator, error)

	WatchAllowedSignersUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdated) (event.Subscription, error)

	ParseAllowedSignersUpdated(log types.Log) (*WorkflowRegistryAllowedSignersUpdated, error)

	FilterCapabilitiesRegistryUpdated(opts *bind.FilterOpts) (*WorkflowRegistryCapabilitiesRegistryUpdatedIterator, error)

	WatchCapabilitiesRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryCapabilitiesRegistryUpdated) (event.Subscription, error)

	ParseCapabilitiesRegistryUpdated(log types.Log) (*WorkflowRegistryCapabilitiesRegistryUpdated, error)

	FilterConfigUpdated(opts *bind.FilterOpts) (*WorkflowRegistryConfigUpdatedIterator, error)

	WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryConfigUpdated) (event.Subscription, error)

	ParseConfigUpdated(log types.Log) (*WorkflowRegistryConfigUpdated, error)

	FilterDONLimitSet(opts *bind.FilterOpts) (*WorkflowRegistryDONLimitSetIterator, error)

	WatchDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryDONLimitSet) (event.Subscription, error)

	ParseDONLimitSet(log types.Log) (*WorkflowRegistryDONLimitSet, error)

	FilterOwnershipLinkUpdated(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedIterator, error)

	WatchOwnershipLinkUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdated, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error)

	ParseOwnershipLinkUpdated(log types.Log) (*WorkflowRegistryOwnershipLinkUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error)

	FilterRequestAllowlisted(opts *bind.FilterOpts, owner []common.Address, requestDigest [][32]byte) (*WorkflowRegistryRequestAllowlistedIterator, error)

	WatchRequestAllowlisted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryRequestAllowlisted, owner []common.Address, requestDigest [][32]byte) (event.Subscription, error)

	ParseRequestAllowlisted(log types.Log) (*WorkflowRegistryRequestAllowlisted, error)

	FilterUserDONLimitSet(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitSetIterator, error)

	WatchUserDONLimitSet(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitSet, user []common.Address) (event.Subscription, error)

	ParseUserDONLimitSet(log types.Log) (*WorkflowRegistryUserDONLimitSet, error)

	FilterUserDONLimitUnset(opts *bind.FilterOpts, user []common.Address) (*WorkflowRegistryUserDONLimitUnsetIterator, error)

	WatchUserDONLimitUnset(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryUserDONLimitUnset, user []common.Address) (event.Subscription, error)

	ParseUserDONLimitUnset(log types.Log) (*WorkflowRegistryUserDONLimitUnset, error)

	FilterWorkflowActivated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowActivatedIterator, error)

	WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowActivated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowActivated(log types.Log) (*WorkflowRegistryWorkflowActivated, error)

	FilterWorkflowDeleted(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDeletedIterator, error)

	WatchWorkflowDeleted(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDeleted, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowDeleted(log types.Log) (*WorkflowRegistryWorkflowDeleted, error)

	FilterWorkflowDonFamilyUpdated(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowDonFamilyUpdatedIterator, error)

	WatchWorkflowDonFamilyUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowDonFamilyUpdated, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowDonFamilyUpdated(log types.Log) (*WorkflowRegistryWorkflowDonFamilyUpdated, error)

	FilterWorkflowOwnerConfigUpdated(opts *bind.FilterOpts, owner []common.Address) (*WorkflowRegistryWorkflowOwnerConfigUpdatedIterator, error)

	WatchWorkflowOwnerConfigUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowOwnerConfigUpdated, owner []common.Address) (event.Subscription, error)

	ParseWorkflowOwnerConfigUpdated(log types.Log) (*WorkflowRegistryWorkflowOwnerConfigUpdated, error)

	FilterWorkflowPaused(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowPausedIterator, error)

	WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowPaused, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowPaused(log types.Log) (*WorkflowRegistryWorkflowPaused, error)

	FilterWorkflowRegistered(opts *bind.FilterOpts, workflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowRegisteredIterator, error)

	WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowRegistered, workflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowRegistered(log types.Log) (*WorkflowRegistryWorkflowRegistered, error)

	FilterWorkflowUpdated(opts *bind.FilterOpts, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (*WorkflowRegistryWorkflowUpdatedIterator, error)

	WatchWorkflowUpdated(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryWorkflowUpdated, oldWorkflowId [][32]byte, newWorkflowId [][32]byte, owner []common.Address) (event.Subscription, error)

	ParseWorkflowUpdated(log types.Log) (*WorkflowRegistryWorkflowUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
