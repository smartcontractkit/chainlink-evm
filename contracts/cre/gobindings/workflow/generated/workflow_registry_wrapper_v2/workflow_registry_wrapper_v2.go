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
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminBatchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseAllByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminPauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowlistRequest\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchActivateWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPauseWorkflows\",\"inputs\":[{\"name\":\"workflowIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveAllowlistedRequestsReverse\",\"inputs\":[{\"name\":\"endIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allowlistedRequests\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.OwnerAllowlistedRequest[]\",\"components\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"searchComplete\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedSigners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowlistedRequests\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"allowlistedRequests\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.OwnerAllowlistedRequest[]\",\"components\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilitiesRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.Config\",\"components\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonConfigs\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.DonConfigView[]\",\"components\":[{\"name\":\"donHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"family\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEvents\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.EventRecord[]\",\"components\":[{\"name\":\"eventType\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.EventType\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"maxWorkflows\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultUserLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserDONOverrides\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.UserOverrideView[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowById\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"workflow\",\"type\":\"tuple\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowListByOwnerAndName\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"list\",\"type\":\"tuple[]\",\"internalType\":\"structWorkflowRegistry.WorkflowMetadataView[]\",\"components\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseWorkflow\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCapabilitiesRegistry\",\"inputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"nameLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tagLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"urlLen\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"attrLen\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"expiryLen\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONLimit\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"userLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwnerConfig\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsByOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveWorkflowsOnDON\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowedSigners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAllowlistedRequests\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEvents\",\"inputs\":[],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWorkflowDONFamily\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upsertWorkflow\",\"inputs\":[{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"binaryUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attributes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keepAlive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdated\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilitiesRegistryUpdated\",\"inputs\":[{\"name\":\"oldAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"maxNameLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxTagLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxUrlLen\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxAttrLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxExpiryLen\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONLimitSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"donLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"userDefaultLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAllowlisted\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserDONLimitUnset\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowActivated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDeleted\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowDonFamilyUpdated\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newDonFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerConfigUpdated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowPaused\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowRegistered\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowUpdated\",\"inputs\":[{\"name\":\"oldWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newWorkflowId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"workflowName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AttributesTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BinaryURLRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotWorkflowOwner\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotChangeDONFamilyOnUpdate\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attemptedDonFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CannotChangeStatusOnUpdate\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tag\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"attemptedStatus\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.WorkflowStatus\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateDONFamilyForPausedWorkflows\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DonLimitNotSet\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"EmptyUpdateBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidExpiryTimestamp\",\"inputs\":[{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxAllowed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerDONExceeded\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MaxWorkflowsPerUserDONExceeded\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipProofAlreadyUsed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"PreviousAllowlistedRequestStillValid\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"URLTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UserDONDefaultLimitExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserDONOverrideExceedsDONLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowIDAlreadyExists\",\"inputs\":[{\"name\":\"workflowId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WorkflowNameRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowNameTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"WorkflowTagRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WorkflowTagTooLong\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAllowed\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWorkflowIDNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60803460a8573315609757600180546001600160a01b0319163317905560a081016001600160401b03811182821017608157604090815280825260208083015260c882820152610400606083015262093a80608090920191909152600280546001600160481b03191667093a800400c82040179055516158f190816100ae8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fdfe610140604052600436101561001357600080fd5b60003560e01c806301b76905146130f95780630987294c146130bf57806317e0edfc14613018578063181f5a7714612ffc5780631c08b00a14612f2d5780631c71682c14612d3e578063245b8e4e14612d00578063289bd10814612c435780632afc413014612beb5780632c50a95514612b8e57806339d68c6a14612a3c57806339e4323414612a0e5780633bb705781461295d5780633c17181b1461290a5780633c54b50b146128b75780633d90a108146128995780634b6d2e5b146126f7578063530979d61461266857806358c9d589146125d7578063610431931461253a578063695e1340146124e95780636caffc4c1461246757806370ae26401461243157806376c2ed86146123f757806379ba50971461235e578063865ec9e0146123255780638b42a96d146121d75780638c42ffc5146121235780638da5cb5b146120fc57806394ea0da614611ecf578063952bb98414611e84578063a0b8a4fe14611e66578063a408901614611dc3578063a6008f2014611ca1578063a7d0185814611bfa578063afbb240114611be3578063b377bfc514610c2e578063b668435f14610a1f578063ba87068614610a01578063bae5c29a14610988578063bdf6b4ff14610923578063be674333146108fc578063bf2a0d93146108a9578063c3f909d4146107f8578063c59a655a1461071e578063cabb9e7a146106da578063d8b80738146105f8578063d8e4a72414610483578063dc101969146103e5578063de49b95f146103c7578063e690f33214610343578063ea32308b146103095763f2fde38b1461026457600080fd5b34610304576020366003190112610304576001600160a01b03610285613307565b61028d614ad7565b163381146102da57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b34610304576020366003190112610304576001600160a01b0361032a613307565b16600052600c6020526020604060002054604051908152f35b346103045760203660031901126103045761036b336000526006602052604060002054151590565b156103b257600435600052600b60205260406000205461038b8133614bf6565b600160ff8183015460e01c166103a081613478565b036103a757005b6103b091615179565b005b63c2dda3f960e01b6000523360045260246000fd5b34610304576000366003190112610304576020601554604051908152f35b346103045760603660031901126103045760443560243567ffffffffffffffff82116103045761042961041e600193369060040161344a565b908360043533614117565b3360005260076020528060406000205561044233615504565b5080600052600860205260406000208260ff19825416179055337f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b35600080a4005b346103045760403660031901126103045760043567ffffffffffffffff8111610304576104b49036906004016135f4565b6024359182151592838103610304576104cb614ad7565b60005b83811061055857505060405191806040840160408552526060830191906000905b808210610525577f861d38caf3055a11344d9f540d5ab4e5c38d751dfcbd1156aed92b71805e13168580868960208301520390a1005b9091928335906001600160a01b038216820361030457602080916001600160a01b036001941681520194019201906104ef565b6001600160a01b0361057361056e838787613ae1565b61491b565b16156105ce5760019082156105a9576105a16001600160a01b0361059b61056e848989613ae1565b1661548d565b505b016104ce565b6105c86001600160a01b036105c261056e848989613ae1565b166156de565b506105a3565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346103045760203660031901126103045760043567ffffffffffffffff8111610304576106299036906004016135f4565b80156106b057610646336000526006602052604060002054151590565b156103b25760005b81811061065757005b806106656001928486613ae1565b35600052600b60205260406000205461067e8133614bf6565b8360ff8183015460e01c1661069281613478565b036106a0575b50500161064e565b6106a991615179565b8480610698565b7faea36d000000000000000000000000000000000000000000000000000000000060005260046000fd5b346103045760203660031901126103045760206107146001600160a01b03610700613307565b166000526004602052604060002054151590565b6040519015158152f35b346103045760403660031901126103045760043567ffffffffffffffff81116103045761075261076791369060040161344a565b919060243592610760614ad7565b369161388b565b60208151910120600052600d60205260406000209060005b825490811515806107e7575b156103b05760001982019182116107d1576107c76107ac6107cc938661530a565b90549060031b1c806000526009602052604060002090615179565b613e5a565b61077f565b634e487b7160e01b600052601160045260246000fd5b5082158061078b575082811061078b565b346103045760003660031901126103045760006080604051610819816133b4565b828152826020820152826040820152826060820152015260a060405161083e816133b4565b63ffffffff60025461ffff60ff82169384815260ff60208201818560081c168152816040840191818760101c1683528760806060870196888a60181c168852019760281c16875260405198895251166020880152511660408601525116606084015251166080820152f35b346103045760603660031901126103045760043567ffffffffffffffff8111610304576108dd6103b091369060040161344a565b6108e5613727565b906108ee61373a565b926108f7614ad7565b61467e565b3461030457602061091561090f366136e7565b9161449c565b63ffffffff60405191168152f35b346103045760203660031901126103045760043567ffffffffffffffff81116103045761076061095791369060040161344a565b6020815191012060005260176020526040600181600020015463ffffffff825191818116835260201c166020820152f35b34610304576040366003190112610304576109a1613307565b6040516109de816109d060208201946024359086602090939291936001600160a01b0360408201951681520152565b03601f1981018352826133ed565b5190206000526014602052602063ffffffff604060002054166040519042108152f35b34610304576000366003190112610304576020601c54604051908152f35b3461030457608036600319011261030457610a38613307565b60243567ffffffffffffffff811161030457610a5890369060040161344a565b90610a6161373a565b9160643592831515840361030457610a77614ad7565b610a8236838561388b565b602081519101208060005260176020526040600020600181019063ffffffff82541615610c0757600201604060006001600160a01b038a168152826020522096600014610b8b57505463ffffffff92831692168211610b61577f945de380da34dd2a3da003b018b92eb5714e63dbcc911e3caef8307407c823249464010000000064ff00000000198254161781558263ffffffff19825416179055600052601a602052610b3e6001600160a01b036040600020961680966155b8565b50610b56604051938493604085526040850191613a94565b9060208301520390a2005b7feabc4fd90000000000000000000000000000000000000000000000000000000060005260046000fd5b919250507f6b349f5a70df2e3faf5fb1a615930e6816698828af7279d4d231f0edc013d84994506001600160a01b03861660005260205260006040812055600052601a602052610bea6001600160a01b03604060002095168095615792565b50610c02604051928392602084526020840191613a94565b0390a2005b60405163393f328760e11b81526020600482015280610c2a60248201888a613a94565b0390fd5b34610304576101203660031901126103045760043567ffffffffffffffff811161030457610c6090369060040161344a565b610100526101205260243567ffffffffffffffff811161030457610c8890369060040161344a565b90600260643510156103045760843567ffffffffffffffff811161030457610cb490369060040161344a565b9060e0529060a43567ffffffffffffffff811161030457610cd990369060040161344a565b9160c43567ffffffffffffffff811161030457610cfa90369060040161344a565b9560e43567ffffffffffffffff811161030457610d1b90369060040161344a565b9290936101043595861515870361030457610d43336000526006602052604060002054151590565b156103b25760443515611bb957604435600052600b602052604060002054611b895760025460ff8160101c168915611b5f5780611b22575b5061ffff8160181c1680151580611b19575b611ae857508415611abe5760ff8160081c1680151580611ab5575b611a8457506101005115611a5a5760ff1680151580611a4e575b611a1a57506040516020810190610de7816109d0888761010051610120513389614325565b51902097886000526009602052604060002060018101546001600160a01b038116801560001461154d57505050610e2e610e2836610100516101205161388b565b3361498a565b9a610e3c368c60e05161388b565b6020815191012098156114f2575b610e55606435613478565b606435156114ca575b604051978861012081011067ffffffffffffffff6101208b0111176112f357610f0b899695610efc8897610f1a956101208a016040526044358a5260a0610ef060208c019a338c5260408d0160c05267ffffffffffffffff421660c0515260608d019a610ecc606435613478565b6064358c526080610ee436610100516101205161388b565b9e019d8e52369161388b565b9b019a8b52369161388b565b9860c08c01998a52369161388b565b9760e08a01988952369161388b565b610100880160a05260a051528860005260096020526040600020965187556001600160a01b038060018901935116166001600160a01b031983541617825560c051517fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b8085549360a01b1616911617825551610fa181613478565b610faa81613478565b60ff60e01b197cff0000000000000000000000000000000000000000000000000000000083549260e01b16911617905560028501905180519067ffffffffffffffff82116112f357611006826110008554613772565b856140d0565b602090601f831160011461146357611037929160009183611458575b50508160011b916000199060031b1c19161790565b90555b518051600385019167ffffffffffffffff82116112f35761105f826110008554613772565b602090601f83116001146113f15761108f9291600091836113e65750508160011b916000199060031b1c19161790565b90555b518051600484019167ffffffffffffffff82116112f3576110b7826110008554613772565b602090601f831160011461137f576110e79291600091836113095750508160011b916000199060031b1c19161790565b90555b518051600583019167ffffffffffffffff82116112f35761110f826110008554613772565b602090601f83116001146113145791806111449260069695946000926113095750508160011b916000199060031b1c19161790565b90555b019360a0515194855167ffffffffffffffff81116112f3576111738161116d8454613772565b846140d0565b6020601f82116001146112895790806111a79261120d989960009261127e5750508160011b916000199060031b1c19161790565b90555b600052600a6020526111c08260406000206155b8565b50604435600052600b6020528160406000205581600052601360205280604060002055600052600e6020526111f98160406000206155b8565b5033600052600f60205260406000206155b8565b506112266040519160608352606083019060e051613a94565b611231606435613478565b606435602083015281810360408301527f74dc2e5bdab0a48c5e7d33c1eaad00066fd19c8d9f29d4c3a251711c0a0e9a09339280611279604435946101005161012051613a94565b0390a3005b015190508980611022565b601f1982169783600052816000209860005b8181106112db57509161120d9899918460019594106112c2575b505050811b0190556111aa565b015160001960f88460031b161c191690558880806112b5565b838301518b556001909a01996020938401930161129b565b634e487b7160e01b600052604160045260246000fd5b015190508a80611022565b90601f1983169184600052816000209260005b81811061136757509160019391856006989796941061134e575b505050811b019055611147565b015160001960f88460031b161c19169055898080611341565b92936020600181928786015181550195019301611327565b90601f1983169184600052816000209260005b8181106113ce57509084600195949392106113b5575b505050811b0190556110ea565b015160001960f88460031b161c191690558980806113a8565b92936020600181928786015181550195019301611392565b015190508b80611022565b90601f1983169184600052816000209260005b8181106114405750908460019594939210611427575b505050811b019055611092565b015160001960f88460031b161c191690558a808061141a565b92936020600181928786015181550195019301611404565b015190508c80611022565b90601f1983169184600052816000209260005b8181106114b25750908460019594939210611499575b505050811b01905561103a565b015160001960f88460031b161c191690558b808061148c565b92936020600181928786015181550195019301611476565b6114e16114da368d60e05161388b565b8a33614c9e565b6114ed8c8a338d615253565b610e5e565b9a999897969594939291908a600052601060205260406000206080525b60805154801561153c57806000198101116107d1576107ac611537916000190160805161530a565b61150f565b50909192939495969798999a610e4a565b9091989c96929a9493959799503303611a055760e01c60ff16611571606435613478565b61157a81613478565b606435036119875782600052601360205260406000205461159e368c60e05161388b565b60208151910120036119075750508454600052600b60205260006040812055604435600052600b602052604060002055835495604435855560038501916040516115f3816115ec81876137ac565b03826133ed565b6020815191012061160536848461388b565b6020815191012003611840575b505050600483019160405161162b816115ec81876137ac565b6020815191012061163d36848461388b565b6020815191012003611770575b5050506006019067ffffffffffffffff81116112f35761166e8161116d8454613772565b6000601f821160011461170c5781906116a0939495966000926117015750508160011b916000199060031b1c19161790565b90555b6116bb6040519260408452604084019060e051613a94565b9082820360208401527f03d454e4bcb8ae5031ab165ca5f4161ebf48cfaf66d96cc490ba500a59a1e3223393806116fc604435956101005161012051613a94565b0390a4005b013590508680611022565b601f198216958382526020822091805b8881106117585750836001959697981061173e575b505050811b0190556116a3565b0135600019600384901b60f8161c19169055858080611731565b9092602060018192868601358155019401910161171c565b67ffffffffffffffff82116112f35761178d826110008554613772565b600090601f83116001146117d85791806117c29260069695946000926117cd5750508160011b916000199060031b1c19161790565b90555b90868061164a565b013590508a80611022565b8382526020822091601f198416815b81811061182857509160019391856006989796941061180e575b505050811b0190556117c5565b0135600019600384901b60f8161c19169055898080611801565b919360206001819287870135815501950192016117e7565b67ffffffffffffffff82116112f35761185d826110008554613772565b600090601f83116001146118a35761188d9291600091836118985750508160011b916000199060031b1c19161790565b90555b878080611612565b013590508b80611022565b8382526020822091601f198416815b8181106118ef57509084600195949392106118d5575b505050811b019055611890565b0135600019600384901b60f8161c191690558a80806118c8565b919360206001819287870135815501950192016118b2565b610c2a8a6119726040519485947f08f08b2c000000000000000000000000000000000000000000000000000000008652604435600487015233602487015260a0604487015261196060a487016101005161012051613a94565b86810360031901606488015291613a94565b9060031984830301608485015260e051613a94565b6119ee6040519283927f754824c9000000000000000000000000000000000000000000000000000000008452604435600485015233602485015260a060448501526119dc60a485016101005161012051613a94565b84810360031901606486015291613a94565b6119f9606435613478565b60643560848301520390fd5b6331ee6dc760e01b6000523360045260246000fd5b7f36a7c503000000000000000000000000000000000000000000000000000000006000526101005160045260245260446000fd5b50806101005111610dc2565b7f485b8ed40000000000000000000000000000000000000000000000000000000060005260046000fd5b857f436f97540000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808611610da8565b7f65cf28770000000000000000000000000000000000000000000000000000000060005260046000fd5b867f354f25140000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50808711610d8d565b808a11611b4857808c1115610d7b578b6219aad560e31b60005260045260245260446000fd5b896219aad560e31b60005260045260245260446000fd5b7f9cd963cf0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0d5354a40000000000000000000000000000000000000000000000000000000060005260443560045260246000fd5b7f315de7450000000000000000000000000000000000000000000000000000000060005260046000fd5b34610304576103b0611bf4366136b7565b9161435a565b346103045760203660031901126103045760043567ffffffffffffffff811161030457611c2b9036906004016135f4565b80156106b05760005b818110611c3d57005b80611c4b6001928486613ae1565b35611c54614ad7565b600052600b602052604060002054806000526009602052604060002060ff8482015460e01c16611c8381613478565b15611c91575b505001611c34565b611c9a91615179565b8480611c89565b346103045760a03660031901126103045760043560ff8116809103610304576024359060ff821690818303610304576044359160ff831690818403610304576064359161ffff831691828403610304576084359663ffffffff8816948589036103045764ffff00000060a09862ff00007f9c1a161a4cdd9b19a46f9660eee21b6394dc5aa70fc9e093dbb36d2c1786d7739b611d3b614ad7565b896080604051611d4a816133b4565b8d81528960208201528a60408201528b6060820152015268ffffffff00000000006002549160281b169561ff0068ffffffff000000000019928d87199162ffffff191617169160081b1617169160101b16179160181b1617176002556040519485526020850152604084015260608301526080820152a1005b3461030457606036600319011261030457611ddc613307565b60243567ffffffffffffffff811161030457611dfc90369060040161344a565b916044359267ffffffffffffffff841161030457611e46611e4e936109d0611e2b611e6297369060040161344a565b90611e346138d2565b50604051958694602086019889614325565b5190206149a9565b604051918291602083526020830190613482565b0390f35b34610304576000366003190112610304576020600554604051908152f35b3461030457608036600319011261030457611e9d613307565b6064359067ffffffffffffffff821161030457611ec16103b092369060040161344a565b916044359060243590614117565b3461030457604036600319011261030457600435611eeb613727565b63ffffffff8060025460281c169116904282118015906120e0575b6120ac5750611f22336000526006602052604060002054151590565b156103b257604080513360208201908152918101849052611f4681606081016109d0565b51902080600052601460205263ffffffff604060002054164281116120775750600052601460205260406000208163ffffffff19825416179055604051611f8c8161337c565b82815260208101903382526040810190838252601554680100000000000000008110156112f357806001611fc39201601555613970565b6120615760016001600160a01b039291839251815501935116166001600160a01b0319835416178255517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff000000000000000000000000000000000000000083549260a01b1691161790556040519081527ff69135e4f80a25991d2f877c365c191c51ec3c0063ecb9299d314cd9da4880d160203392a3005b634e487b7160e01b600052600060045260246000fd5b837f51aa42c0000000000000000000000000000000000000000000000000000000006000523360045260245260445260646000fd5b917f7ffd3b8f0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b508015158015611f065750806120f64284613d4f565b11611f06565b346103045760003660031901126103045760206001600160a01b0360015416604051908152f35b346103045760603660031901126103045761213c613307565b6001600160a01b0360243591169081600052600f6020526121656044358260406000205461492f565b9061216f82613920565b9260005b8381106121885760405180611e62878261354d565b60019082600052600f6020526121bb6121af60406000206121a98488613765565b9061530a565b90549060031b1c6149a9565b6121c5828861382f565b526121d0818761382f565b5001612173565b34610304576001600160a01b036121ed366136e7565b9290916121f8614ad7565b169081600052601660205260406000209267ffffffffffffffff81116112f35761222c816122268654613772565b866140d0565b600093601f82116001146122a35761227e82807f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f139697600091612298575b508160011b916000199060031b1c19161790565b90555b610c02604051928392602084526020840191613a94565b90508501358861226a565b80855260208520601f19831695805b87811061230d5750837f0d8eb32301e2fa82bb02c4905860f05470c9b7771dcb418fffde59818a053f139697106122f3575b5050600182811b019055612281565b840135600019600385901b60f8161c1916905585806122e4565b909160206001819285890135815501930191016122b2565b34610304576000366003190112610304576040601b5467ffffffffffffffff8251916001600160a01b038116835260a01c166020820152f35b34610304576000366003190112610304576000546001600160a01b03811633036123cd576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346103045761242561241161240b366132b0565b90613f51565b604051928392604084526040840190613625565b90151560208301520390f35b346103045760203660031901126103045761244a6138d2565b50600435600052600b602052611e62611e4e6040600020546149a9565b3461030457604036600319011261030457612480613307565b6001600160a01b0360243591612494614ad7565b16600052600c60205260406000209060005b825490811515806124d8575b156103b05760001982019182116107d1576107c76107ac6124d3938661530a565b6124a6565b508215806124b257508281106124b2565b3461030457602036600319011261030457612511336000526006602052604060002054151590565b156103b257600435600052600b6020526103b06040600020546125348133614bf6565b90614f84565b346103045761257761255a61254e36613680565b9492909391369161388b565b602081519101209283600052600e6020528260406000205461492f565b9061258182613920565b9260005b83811061259a5760405180611e62878261354d565b60019082600052600e6020526125bb6121af60406000206121a98488613765565b6125c5828861382f565b526125d0818761382f565b5001612585565b34610304576126086125eb61254e36613680565b602081519101209283600052600d6020528260406000205461492f565b9061261282613920565b9260005b83811061262b5760405180611e62878261354d565b60019082600052600d60205261264c6121af60406000206121a98488613765565b612656828861382f565b52612661818761382f565b5001612616565b3461030457612676366136b7565b909161268f336000526006602052604060002054151590565b156103b257600052600b6020526040600020546126ac8133614bf6565b9060ff600183015460e01c166126c181613478565b6126c757005b6126f26126eb6103b0956126dc36878361388b565b6020815191012095369161388b565b8433614c9e565b614e05565b346103045761271361270836613680565b92909391369161388b565b6020815191012080600052601a60205261273460406000209284845461492f565b9161273e8361374d565b9361274c60405195866133ed565b838552601f1961275b8561374d565b0160005b81811061287457505060005b8481106127d3578560405180916020820160208352815180915260206040840192019060005b81811061279f575050500390f35b825180516001600160a01b0316855260209081015163ffffffff168186015286955060409094019390920191600101612791565b806001600160a01b036127f16127eb60019486613765565b8661530a565b90549060031b1c168560005260176020526040600281600020016000906001600160a01b0384168252602052206040519061282b82613398565b5490602060ff63ffffffff841693848452821c1615159101526040519161285183613398565b82526020820152612862828961382f565b5261286d818861382f565b500161276b565b60209060405161288381613398565b6000815260008382015282828a0101520161275f565b34610304576000366003190112610304576020600354604051908152f35b346103045760203660031901126103045760043567ffffffffffffffff8111610304576107606128eb91369060040161344a565b60208151910120600052600d6020526020604060002054604051908152f35b3461030457602036600319011261030457612923614ad7565b600435600052600b602052604060002054806000526009602052604060002060ff600182015460e01c1661295681613478565b156103a757005b346103045761297a61296e366132b0565b8160039392935461492f565b61298381613859565b90600092600354935b8281106129a15760405180611e62868261331d565b6129ab8183613765565b6000868210156129fa57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0154600191906001600160a01b03166129f3828761382f565b520161298c565b80634e487b7160e01b602492526032600452fd5b3461030457611e62612a28612a22366132b0565b90613e69565b604051918291602083526020830190613625565b3461030457612a4a366135ad565b928291924211612b69576001600160a01b03811693612a76856000526006602052604060002054151590565b15612b5457612a8485615322565b93612a9d6001600160a01b036107008484898989614b15565b15612b3757858581600052600f6020526040600020905b81548015612af55760001981019081116107d157612ad5612af0918461530a565b90549060031b1c806000526009602052604060002090614f84565b612ab4565b600082858083526007602052826040812055612b1081615614565b507f07756706c87366f7add7b5c7df5dd4f570e02667b54e60b75e1fd1a2ac294b358380a4005b60405163335d4ce160e01b8152948594610c2a9460048701613ab5565b8463c2dda3f960e01b60005260045260246000fd5b6001600160a01b0390631ec5288b60e11b600052166004524260245260445260646000fd5b346103045760403660031901126103045760043567ffffffffffffffff811161030457612bbf9036906004016135f4565b6024359167ffffffffffffffff831161030457612be36103b093369060040161344a565b929091613af1565b34610304576020366003190112610304576001600160a01b03612c0c613307565b166000526016602052611e626115ec612c2f6040600020604051928380926137ac565b6040519182916020835260208301906132c6565b3461030457612c51366135ad565b8293924211612cda576001600160a01b038316612c7b816000526006602052604060002054151590565b15612cc657612c8990615322565b91612ca26001600160a01b036107008484878a8a614b15565b15612ca957005b610c2a9260405195869563335d4ce160e01b875260048701613ab5565b63c2dda3f960e01b60005260045260246000fd5b836001600160a01b0384631ec5288b60e11b600052166004524260245260445260646000fd5b3461030457604036600319011261030457612d19613307565b60243567ffffffffffffffff81168103610304576103b091612d39614ad7565b61398f565b3461030457612d5b612d4f366132b0565b81601c9392935461492f565b90612d658261374d565b91612d7360405193846133ed565b808352601f19612d828261374d565b0160005b818110612f01575050601c54909160005b838110612e43578460405160208101916020825280518093526040820192602060408260051b85010192019060005b818110612dd35784840385f35b909192603f198582030186528351908151916003831015612e2d57612e1e82606060406020959460019787965263ffffffff86820151168685015201519181604082015201906132c6565b95019601910194919094612dc6565b634e487b7160e01b600052602160045260246000fd5b612e4d8183613765565b6000848210156129fa57601c90526040519060011b7f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a211016000612e8f8361337c565b81549060ff8216906003821015612eed5750835260081c63ffffffff1660208301526040516001939291612ecc9082906115ec90829088016137ac565b6040820152612edb828861382f565b52612ee6818761382f565b5001612d97565b80634e487b7160e01b602492526021600452fd5b602090604051612f108161337c565b600081526000838201526060604082015282828801015201612d86565b3461030457608036600319011261030457612f46613307565b60243567ffffffffffffffff811161030457612f7a91612f6d612f8092369060040161344a565b939060443594369161388b565b9061498a565b9081600052600a602052612f9c6064358260406000205461492f565b90612fa682613920565b9260005b838110612fbf5760405180611e62878261354d565b60019082600052600a602052612fe06121af60406000206121a98488613765565b612fea828861382f565b52612ff5818761382f565b5001612faa565b3461030457600036600319011261030457611e62612c2f61340f565b3461030457613035613029366132b0565b8160059392935461492f565b61303e81613859565b90600092600554935b82811061305c5760405180611e62868261331d565b6130668183613765565b6000868210156129fa57600590527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0015460008190526007602052600191906001600160a01b03166130b8828761382f565b5201613047565b346103045760203660031901126103045760206107146001600160a01b036130e5613307565b166000526006602052604060002054151590565b346103045761311661310a366132b0565b8160189392935461492f565b61311f8161374d565b9161312d60405193846133ed565b818352601f1961313c8361374d565b0160005b81811061327d575050601854919060005b8281106131ef57846040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061319157505050500390f35b919360019193955060208091603f1989820301855287519081518152606063ffffffff816131cc8686015160808887015260808601906132c6565b948260408201511660408601520151169101529601920192018594939192613182565b6131f98183613765565b6000858210156129fa579060208260186001959452200163ffffffff604060009254928381526017602052206115ec61324b86830154926040519561323d87613360565b8652604051928380926137ac565b6020840152818116604084015260201c16606082015261326b828861382f565b52613276818761382f565b5001613151565b60209060405161328c81613360565b60008152606083820152600060408201526000606082015282828801015201613140565b6040906003190112610304576004359060243590565b919082519283825260005b8481106132f2575050826000602080949584010152601f8019910116010190565b806020809284010151828286010152016132d1565b600435906001600160a01b038216820361030457565b602060408183019282815284518094520192019060005b8181106133415750505090565b82516001600160a01b0316845260209384019390920191600101613334565b6080810190811067ffffffffffffffff8211176112f357604052565b6060810190811067ffffffffffffffff8211176112f357604052565b6040810190811067ffffffffffffffff8211176112f357604052565b60a0810190811067ffffffffffffffff8211176112f357604052565b610140810190811067ffffffffffffffff8211176112f357604052565b90601f8019910116810190811067ffffffffffffffff8211176112f357604052565b6040519061341e6040836133ed565b601682527f576f726b666c6f77526567697374727920322e302e30000000000000000000006020830152565b9181601f840112156103045782359167ffffffffffffffff8311610304576020838186019501011161030457565b60021115612e2d57565b61354a91815181526001600160a01b03602083015116602082015267ffffffffffffffff604083015116604082015260608201516134bf81613478565b60608201526101206135386135246135126135006134ee608088015161014060808901526101408801906132c6565b60a088015187820360a08901526132c6565b60c087015186820360c08801526132c6565b60e086015185820360e08701526132c6565b6101008501518482036101008601526132c6565b920151906101208184039101526132c6565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061358057505050505090565b909192939460208061359e600193603f198682030187528951613482565b97019301930191939290613571565b6060600319820112610304576004356001600160a01b03811681036103045791602435916044359067ffffffffffffffff8211610304576135f09160040161344a565b9091565b9181601f840112156103045782359167ffffffffffffffff8311610304576020808501948460051b01011161030457565b906020808351928381520192019060005b8181106136435750505090565b9091926020606060019263ffffffff60408851805184526001600160a01b0386820151168685015201511660408201520194019101919091613636565b6060600319820112610304576004359067ffffffffffffffff8211610304576136ab9160040161344a565b90916024359060443590565b90604060031983011261030457600435916024359067ffffffffffffffff8211610304576135f09160040161344a565b906040600319830112610304576004356001600160a01b038116810361030457916024359067ffffffffffffffff8211610304576135f09160040161344a565b6024359063ffffffff8216820361030457565b6044359063ffffffff8216820361030457565b67ffffffffffffffff81116112f35760051b60200190565b919082018092116107d157565b90600182811c921680156137a2575b602083101461378c57565b634e487b7160e01b600052602260045260246000fd5b91607f1691613781565b600092918154916137bc83613772565b808352926001811690811561381257506001146137d857505050565b60009081526020812093945091925b8383106137f8575060209250010190565b6001816020929493945483858701015201910191906137e7565b915050602093945060ff929192191683830152151560051b010190565b80518210156138435760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906138638261374d565b61387060405191826133ed565b8281528092613881601f199161374d565b0190602036910137565b92919267ffffffffffffffff82116112f357604051916138b5601f8201601f1916602001846133ed565b829481845281830111610304578281602093846000960137010152565b604051906138df826133d0565b6060610120836000815260006020820152600060408201526000838201528260808201528260a08201528260c08201528260e0820152826101008201520152565b9061392a8261374d565b61393760405191826133ed565b8281528092613948601f199161374d565b019060005b82811061395957505050565b6020906139646138d2565b8282850101520161394d565b60155481101561384357601560005260206000209060011b0190600090565b90601b54906001600160a01b038216906001600160a01b0367ffffffffffffffff8460a01c1694169382851494858096613a81575b613a7957806080957fc0c3ee74e6d6070ee9c493e8b4f0477d2e66600f22997a4e073288d38d65933b9715613a64575b505067ffffffffffffffff831692828403613a23575b50604051938452602084015260408301526060820152a1565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff60a01b601b549260a01b16911617601b5538613a0a565b6001600160a01b03191617601b5580386139f4565b505050505050565b508167ffffffffffffffff8416146139c4565b908060209392818452848401376000828201840152601f01601f1916010190565b90926080926001600160a01b0361354a9795168352602083015260408201528160608201520191613a94565b91908110156138435760051b0190565b929181156106b057613b10336000526006602052604060002054151590565b156103b257613b2036848361388b565b6020815191012092600091825b848110613cec575063ffffffff831615613a7957613b4c91369161388b565b90836000526017602052604060002090600182015463ffffffff81168015613cca5786600052601260205263ffffffff613b8d848260406000205416614c84565b1611613ca85763ffffffff60029160201c1692016001600160a01b0333166000528060205260ff60406000205460201c16613c88575b5063ffffffff613bf081923360005260116020526040600020886000526020528260406000205416614c84565b9216911611613c6a575060005b818110613c0a5750505050565b80613c186001928487613ae1565b35600052600b6020526040600020548060005260096020528460406000208460ff8183015460e01c16613c4a81613478565b14613c59575b50505001613bfd565b613c6292614e05565b388481613c50565b610c2a9060405191829163038857ff60e01b8352336004840161496a565b336000908152602091909152604090205463ffffffff9081169250613bc3565b60405163b993868760e01b81526020600482015280610c2a60248201876132c6565b60405163393f328760e11b81526020600482015280610c2a60248201886132c6565b92613cf8848689613ae1565b35600052600b60205260ff6001613d1460406000205433614bf6565b015460e01c16613d2381613478565b15613d465763ffffffff1663ffffffff81146107d1576001809101935b01613b2d565b92600190613d40565b919082039182116107d157565b60405190613d6b6020836133ed565b600080835282815b828110613d7f57505050565b602090604051613d8e8161337c565b600081526000838201526000604082015282828501015201613d73565b80634e487b7160e01b602492526041600452fd5b90613dc98261374d565b613dd660405191826133ed565b8281528092613de7601f199161374d565b019060005b828110613df857505050565b602090604051613e078161337c565b600081526000838201526000604082015282828501015201613dec565b90604051613e318161337c565b604063ffffffff600183958054855201546001600160a01b038116602085015260a01c16910152565b60001981146107d15760010190565b90613e77908260155461492f565b908115613f4757613e8782613dbf565b91600091825b828110613ee45750508110613ea0575090565b613ea981613dbf565b9160005b828110613eba5750505090565b80613ec76001928461382f565b51613ed2828761382f565b52613edd818661382f565b5001613ead565b613ef6613ef18284613765565b613970565b504263ffffffff600183015460a01c1611613f15575b50600101613e8d565b60019194613f25613f4092613e24565b613f2f828961382f565b52613f3a818861382f565b50613e5a565b9390613f0c565b505061354a613d5c565b919060009260155480159081156140a4575b50801561409b575b61408957613f798282613d4f565b600181018091116107d157613f8d81613dbf565b9263ffffffff60025460281c16801560001461407a575060005b6000935b82811015614007575b5050508110613fc257509190565b613fcb81613dbf565b9160005b828110613fdd575050509190565b80613fea6001928461382f565b51613ff5828761382f565b52614000818661382f565b5001613fcf565b61401081613970565b5063ffffffff600182015460a01c16838110614068574210614053575b5080156140435780156107d15760001901613fab565b5050509350600193388080613fb4565b94613f256140619296613e24565b933861402d565b50505050509350600193388080613fb4565b6140849042613d4f565b613fa7565b50509050614095613d5c565b90600190565b50808211613f6b565b6000198101915081116107d157811138613f63565b8181106140c4575050565b600081556001016140b9565b9190601f81116140df57505050565b61410b926000526020600020906020601f840160051c8301931061410d575b601f0160051c01906140b9565b565b90915081906140fe565b929093918442116142e6576001600160a01b038416614143816000526006602052604060002054151590565b6142b95781600052600860205260ff604060002054166142895760009061419f6141b961416e61340f565b6040519283916020830195878752604084015246606084015230608084015260e060a08401526101008301906132c6565b8a60c08301528660e083015203601f1981018352826133ed565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c526141fa603c82206141f436878761388b565b90615385565b9091926004831015612eed5782614251575050506001600160a01b031660009081526004602052604090205415614232575050505050565b90610c2a929160405195869563335d4ce160e01b875260048701613ab5565b5060405163d36ab6b960e01b81526060600482015291829160ff614279606485018a8a613a94565b9216602484015260448301520390fd5b7f77a338580000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7fd9a5f5ca0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b846001600160a01b03857f502d038700000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b939161354a95936001600160a01b0361434c93168652606060208701526060860191613a94565b926040818503910152613a94565b9091614373336000526006602052604060002054151590565b156103b25781600052600b602052604060002054906143928233614bf6565b60ff600182015460e01c166143a681613478565b614472576143db90836000526013602052604060002054928360005260176020526143e26040600020604051948580926137ac565b03846133ed565b6143ed36828961388b565b6020815191012080941461446957614440614454948387614432614464967f9b5361a5258ef6ac8039fd2d2ac276734695219cfd870711e7922c236e5db16d9a615179565b6126f26126eb36878e61388b565b6040519384936040855260408501906132c6565b9083820360208501523397613a94565b0390a3565b50505050505050565b7fd74915a80000000000000000000000000000000000000000000000000000000060005260046000fd5b91906144a991369161388b565b6020815191012060005260176020526001600160a01b03604060002091166000526002810160205260406000206020604051916144e583613398565b549160ff63ffffffff841693848352831c161515918291015261451357506001015460201c63ffffffff1690565b905090565b90601c54680100000000000000008110156112f35760018101601c556000601c548210156129fa57601c90526000929060011b7f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a211018151600381101561466a5760409160019160ff825491168060ff19831617835564ffffffff00602087015160081b169164ffffffffff191617178155019101519283519067ffffffffffffffff8211613dab576145ce826110008554613772565b602090601f831160011461460c57906145fd939495836146015750508160011b916000199060031b1c19161790565b9055565b015190503880611022565b90601f198316848352818320925b81811061465257509583600195969710614639575b505050811b019055565b015160001960f88460031b161c1916905538808061462f565b9192602060018192868b01518155019401920161461a565b602485634e487b7160e01b81526021600452fd5b9193929361468d36838561388b565b60208151910120948560005260176020526040600020956001870163ffffffff81549416938463ffffffff8216149081614905575b506148fb5763ffffffff8316928484116148d15788546146e181613772565b156147aa575b507fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9798509061478392918563ffffffff1983541617825567ffffffff0000000082549160201b169067ffffffff0000000019161790556040519060208201528360408201526040815261475c6060826133ed565b604051906147698261337c565b6000825263ffffffff421660208301526040820152614518565b61479a604051948594606086526060860191613a94565b91602084015260408301520390a1565b67ffffffffffffffff87116112f357866147c66147cc92613772565b8b6140d0565b600098601f8711600114614847579061478393929161482488807fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9c9d60009161483c57508160011b916000199060031b1c19161790565b90555b6148308361555e565b509091928998506146e7565b90508b01353861226a565b808a5260208a20601f1988168b5b8181106148b95750907fe6a468e36669d9496095c02796a8a8dcda8ec8d551f6c7454948ecc68aac162d9a9b8961478397969594931061489f575b5050600188811b019055614827565b8a013560001960038b901b60f8161c191690553880614890565b8a8d0135835560209c8d019c60019093019201614855565b7fec623c5f0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050505050509050565b905063ffffffff8085169160201c1614386146c2565b356001600160a01b03811681036103045790565b8082101561496257828161494661354a9585613765565b11156149525750613d4f565b61495d915082613765565b613d4f565b505050600090565b6040906001600160a01b0361354a949316815281602082015201906132c6565b906149a36109d09160405192839160208301958661496a565b51902090565b6149b16138d2565b81600052600960205260406000206001810154926001600160a01b038416928315614acf5750916006614ab4836115ec9567ffffffffffffffff614ac696549860ff8160e01c1692600052601360205260406000205460005260176020526040600020966040519a614a228c6133d0565b8b5260208b015260a01c166040890152614a3b81613478565b6060880152604051614a54816115ec81600286016137ac565b6080880152604051614a6d816115ec81600386016137ac565b60a0880152604051614a86816115ec81600486016137ac565b60c0880152604051614a9f816115ec81600586016137ac565b60e08801526115ec60405180948193016137ac565b610100850152604051928380926137ac565b61012082015290565b935050505090565b6001600160a01b03600154163303614aeb57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91614b7b90614b6292614b2661340f565b916040519485936001600160a01b03602086019860018a5216604086015246606086015230608086015260e060a08601526101008501906132c6565b9160c084015260e083015203601f1981018352826133ed565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52614bb8603c6000206141f436858561388b565b6004829593951015612e2d5781614bd0575050505090565b60ff61427960405195869563d36ab6b960e01b8752606060048801526064870191613a94565b9060005260096020526040600020906001600160a01b03600183015416908115614c42576001600160a01b0316809103614c2e575090565b6331ee6dc760e01b60005260045260246000fd5b7f871e01b20000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff60019116019063ffffffff82116107d157565b9063ffffffff8091169116019063ffffffff82116107d157565b90806000526017602052604060002090600182015463ffffffff81168015614dbe5782600052601260205263ffffffff614ce060018260406000205416614c84565b1611614d9c5763ffffffff60029160201c1692016001600160a01b0384166000528060205260ff60406000205460201c16614d76575b506001600160a01b038316600052601160205260406000209060005260205263ffffffff80614d4d60018260406000205416614c84565b9216911611614d5a575050565b610c2a60405192839263038857ff60e01b84526004840161496a565b9091506001600160a01b03831660005260205263ffffffff604060002054169038614d16565b60405163b993868760e01b81526020600482015280610c2a60248201886132c6565b60405163393f328760e11b81526020600482015280610c2a60248201896132c6565b9091614df761354a936040845260408401906137ac565b9160208184039101526137ac565b7f7fdfd5efa814f8c90ef997a11cfbb29660e5af33312debb9c4898169a73f824a90806000526013602052604060002054848103614ef0575b506001600160a01b036001840193614e7a8286541693876002840195614e74604051614e6e816115ec818c6137ac565b8261498a565b92615253565b60ff60e01b198554168555614ecf815460405190886020830152604082015260408152614ea86060826133ed565b60405190614eb58261337c565b6001825263ffffffff421660208301526040820152614518565b54935416936000526017602052604060002061446460405192839283614de0565b600052600e602052614f06816040600020615792565b5083600052600e602052614f1e8160406000206155b8565b508060005260136020528360406000205538614e3e565b614f3f8154613772565b9081614f49575050565b81601f60009311600114614f5b575055565b81835260208320614f7791601f0160051c8101906001016140b9565b8082528160208120915555565b90614e6e907f48f05f657e3e9a02bfe546c4d3961f08b1f9a0f9798b13fcbc2231173b1ecd946150a26150826150b06001600160a01b03600186019560ff87548a848216614fe06115ec600287019d8e604051928380926137ac565b92839183600052601360205260406000205495869260e01c1661500281613478565b15615167575b50915050600052600e6020526150228b6040600020615792565b5083895416600052600f60205261503d8b6040600020615792565b50600052600a6020526150548a6040600020615792565b508154600052600b6020526000604081205560005260176020526150896040600020604051958680926137ac565b03856133ed565b54955416956040519384936040855260408501906132c6565b9083820360208501526137ac565b0390a38060005260096020526006604060002060008155600060018201556150da60028201614f35565b6150e660038201614f35565b6150f260048201614f35565b6150fe60058201614f35565b016151098154613772565b9081615123575b5050600052601360205260006040812055565b81601f6000931160011461513b5750555b3880615110565b8183526020832061515791601f0160051c8101906001016140b9565b8082528160208120915555615134565b615170936153da565b8a818338615008565b7ff764e70143a953f513d351195b60c30d5fdaaca38edb60b262997f551e48868960018301916001600160a01b038354947c010000000000000000000000000000000000000000000000000000000060ff60e01b19871617855582600052601360205261520982604060002054971693876002840195615203604051614e6e816115ec818c6137ac565b926153da565b614ecf81546040519088602083015260408201526040815261522c6060826133ed565b604051906152398261337c565b6002825263ffffffff421660208301526040820152614518565b916001600160a01b0361530794921690816000526011602052604060002081600052602052604060002063ffffffff61528e81835416614c6c565b1663ffffffff19825416179055806000526012602052604060002063ffffffff6152ba81835416614c6c565b1663ffffffff19825416179055600052600d6020526152dd8360406000206155b8565b50600052600c6020526152f48260406000206155b8565b50600052601060205260406000206155b8565b50565b80548210156138435760005260206000200190600090565b8060005260076020526040600020549081158061536f575b615342575090565b7f02b566860000000000000000000000000000000000000000000000000000000060005260045260246000fd5b508060005260066020526040600020541561533a565b81519190604183036153b6576153af92506020820151906060604084015193015160001a9061584f565b9192909190565b505060009160029190565b63ffffffff6000199116019063ffffffff82116107d157565b916001600160a01b0361530794921680600052600c602052615400846040600020615792565b5081600052600d602052615418846040600020615792565b506000526011602052604060002081600052602052604060002063ffffffff615443818354166153c1565b1663ffffffff198254161790556000526012602052604060002063ffffffff61546e818354166153c1565b1663ffffffff1982541617905560005260106020526040600020615792565b806000526004602052604060002054156000146154fe57600354680100000000000000008110156112f3576154e56154ce826001859401600355600361530a565b819391549060031b91821b91600019901b19161790565b9055600354906000526004602052604060002055600190565b50600090565b806000526006602052604060002054156000146154fe57600554680100000000000000008110156112f3576155456154ce826001859401600555600561530a565b9055600554906000526006602052604060002055600190565b806000526019602052604060002054156000146154fe57601854680100000000000000008110156112f35761559f6154ce826001859401601855601861530a565b9055601854906000526019602052604060002055600190565b600082815260018201602052604090205461560d57805490680100000000000000008210156112f357826155f66154ce84600180960185558461530a565b905580549260005201602052604060002055600190565b5050600090565b600081815260066020526040902054801561560d5760001981018181116107d1576005546000198101919082116107d1578181036156a4575b505050600554801561568e576000190161566881600561530a565b8154906000199060031b1b19169055600555600052600660205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6156c66156b56154ce93600561530a565b90549060031b1c928392600561530a565b9055600052600660205260406000205538808061564d565b600081815260046020526040902054801561560d5760001981018181116107d1576003546000198101919082116107d157818103615758575b505050600354801561568e576000190161573281600361530a565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b61577a6157696154ce93600361530a565b90549060031b1c928392600361530a565b90556000526004602052604060002055388080615717565b90600182019181600052826020526040600020548015156000146158465760001981018181116107d15782546000198101919082116107d15781810361580f575b5050508054801561568e5760001901906157ed828261530a565b8154906000199060031b1b191690555560005260205260006040812055600190565b61582f61581f6154ce938661530a565b90549060031b1c9283928661530a565b9055600052836020526040600020553880806157d3565b50505050600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116158d8579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa156158cc576000516001600160a01b038116156158c05790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
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

func (_WorkflowRegistry *WorkflowRegistryCaller) GetActiveAllowlistedRequestsReverse(opts *bind.CallOpts, endIndex *big.Int, startIndex *big.Int) (GetActiveAllowlistedRequestsReverse,

	error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getActiveAllowlistedRequestsReverse", endIndex, startIndex)

	outstruct := new(GetActiveAllowlistedRequestsReverse)
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowlistedRequests = *abi.ConvertType(out[0], new([]WorkflowRegistryOwnerAllowlistedRequest)).(*[]WorkflowRegistryOwnerAllowlistedRequest)
	outstruct.SearchComplete = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetActiveAllowlistedRequestsReverse(endIndex *big.Int, startIndex *big.Int) (GetActiveAllowlistedRequestsReverse,

	error) {
	return _WorkflowRegistry.Contract.GetActiveAllowlistedRequestsReverse(&_WorkflowRegistry.CallOpts, endIndex, startIndex)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetActiveAllowlistedRequestsReverse(endIndex *big.Int, startIndex *big.Int) (GetActiveAllowlistedRequestsReverse,

	error) {
	return _WorkflowRegistry.Contract.GetActiveAllowlistedRequestsReverse(&_WorkflowRegistry.CallOpts, endIndex, startIndex)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetActiveWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getActiveWorkflowListByDON", donFamily, start, limit)

	if err != nil {
		return *new([]WorkflowRegistryWorkflowMetadataView), err
	}

	out0 := *abi.ConvertType(out[0], new([]WorkflowRegistryWorkflowMetadataView)).(*[]WorkflowRegistryWorkflowMetadataView)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetActiveWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetActiveWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetActiveWorkflowListByDON(donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error) {
	return _WorkflowRegistry.Contract.GetActiveWorkflowListByDON(&_WorkflowRegistry.CallOpts, donFamily, start, limit)
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

type GetActiveAllowlistedRequestsReverse struct {
	AllowlistedRequests []WorkflowRegistryOwnerAllowlistedRequest
	SearchComplete      bool
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

	GetActiveAllowlistedRequestsReverse(opts *bind.CallOpts, endIndex *big.Int, startIndex *big.Int) (GetActiveAllowlistedRequestsReverse,

		error)

	GetActiveWorkflowListByDON(opts *bind.CallOpts, donFamily string, start *big.Int, limit *big.Int) ([]WorkflowRegistryWorkflowMetadataView, error)

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
