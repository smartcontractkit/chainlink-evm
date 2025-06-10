// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package capabilities_registry_wrapper_v2

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

type CapabilitiesRegistryAdditionalDONParams struct {
	Name   string
	Config []byte
}

type CapabilitiesRegistryCapability struct {
	LabelledName          string
	Version               string
	CapabilityType        uint8
	ResponseType          uint8
	ConfigurationContract common.Address
}

type CapabilitiesRegistryCapabilityConfiguration struct {
	CapabilityId [32]byte
	Config       []byte
}

type CapabilitiesRegistryCapabilityInfo struct {
	HashedId              [32]byte
	LabelledName          string
	Version               string
	CapabilityType        uint8
	ResponseType          uint8
	ConfigurationContract common.Address
	IsDeprecated          bool
}

type CapabilitiesRegistryDONInfo struct {
	Id                       uint32
	ConfigCount              uint32
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
	NodeP2PIds               [][32]byte
	DonFamily                string
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
}

type CapabilitiesRegistryNodeOperator struct {
	Admin common.Address
	Name  string
}

type CapabilitiesRegistryNodeParams struct {
	NodeOperatorId      uint32
	Signer              [32]byte
	P2pId               [32]byte
	EncryptionPublicKey [32]byte
	HashedCapabilityIds [][32]byte
}

type INodeInfoProviderNodeInfo struct {
	NodeOperatorId      uint32
	ConfigCount         uint32
	WorkflowDONId       uint32
	Signer              [32]byte
	P2pId               [32]byte
	EncryptionPublicKey [32]byte
	HashedCapabilityIds [][32]byte
	CapabilitiesDONIds  []*big.Int
}

var CapabilitiesRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCapabilities\",\"inputs\":[{\"name\":\"capabilities\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.Capability[]\",\"components\":[{\"name\":\"labelledName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"capabilityType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityType\"},{\"name\":\"responseType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityResponseType\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addDON\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"additionalParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.AdditionalDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodeOperators\",\"inputs\":[{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecateCapabilities\",\"inputs\":[{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCapabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo[]\",\"components\":[{\"name\":\"hashedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"labelledName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"capabilityType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityType\"},{\"name\":\"responseType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityResponseType\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapability\",\"inputs\":[{\"name\":\"hashedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo\",\"components\":[{\"name\":\"hashedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"labelledName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"capabilityType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityType\"},{\"name\":\"responseType\",\"type\":\"uint8\",\"internalType\":\"enumCapabilitiesRegistry.CapabilityResponseType\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityConfigs\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.DONInfo[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONsInFamily\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHashedCapabilityId\",\"inputs\":[{\"name\":\"labelledName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nodeInfo\",\"type\":\"tuple\",\"internalType\":\"structINodeInfoProvider.NodeInfo\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesByP2PIds\",\"inputs\":[{\"name\":\"p2pIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCapabilityDeprecated\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDONNameTaken\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDONs\",\"inputs\":[{\"name\":\"donIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodes\",\"inputs\":[{\"name\":\"removedNodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"additionalParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.AdditionalDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CapabilityConfigured\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilityDeprecated\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONFamilySet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeAdded\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorUpdated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRemoved\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CapabilityAlreadyExists\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CapabilityDoesNotExist\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CapabilityIsDeprecated\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CapabilityRequiredByDON\",\"inputs\":[{\"name\":\"hashedCapabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONNameAlreadyTaken\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONCapability\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONNode\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCapabilityConfigurationContractInterface\",\"inputs\":[{\"name\":\"proposedConfigurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFaultTolerance\",\"inputs\":[{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCapabilities\",\"inputs\":[{\"name\":\"hashedCapabilityIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeEncryptionPublicKey\",\"inputs\":[{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeOperatorAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNodeP2PId\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"lengthOne\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lengthTwo\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyExists\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotExist\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotSupportCapability\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfCapabilitiesDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfWorkflowDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x608080604052346088573315604657600080546001600160a01b03191633179055601180546001600160401b031916640100000001179055604051615a9d908161008e8239f35b62461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806305a51966146102075780630fe5800a1461020257806312570011146101fd578063181f5a77146101f85780631d05394c146101f357806321450243146101ee57806322bdbcbc146101e957806323537405146101e4578063275459f2146101df5780632c01a1e8146101da578063358039f4146101d5578063398f3773146101d05780633f2a13c9146101cb578063461cbb72146101c657806350c946fe146101c1578063543f4025146101bc5780635d83d967146101b757806366acaa33146101b2578063715f5295146101ad57806379ba5097146101a85780637fd0c64e146101a357806384f5ed8a1461019e57806386fa4246146101995780638da5cb5b146101945780639cb7c5f41461018f578063b2246c491461018a578063c931517914610185578063ddbe4f8214610180578063e29581aa1461017b578063f2fde38b146101765763fcdc8efe1461017157600080fd5b612fc3565b612e83565b612d6f565b612c82565b612bcd565b612b60565b612b0f565b6129de565b61267f565b6122ab565b6121f7565b61209c565b611f80565b611e73565b611d13565b611caf565b611be7565b611b2d565b611933565b611743565b611184565b610f33565b610e9e565b610e3b565b610d8d565b610c36565b6107f1565b610774565b61069e565b61061b565b6103c5565b9181601f8401121561023d5782359167ffffffffffffffff831161023d576020808501948460051b01011161023d57565b600080fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261023d576004359067ffffffffffffffff821161023d5761028b9160040161020c565b9091565b906020808351928381520192019060005b8181106102ad5750505090565b82518452602093840193909201916001016102a0565b6103449163ffffffff82511681526102e86020830151602083019063ffffffff169052565b60408281015163ffffffff1690820152606082015160608201526080820151608082015260a082015160a082015260e061033360c084015161010060c085015261010084019061028f565b9201519060e081840391015261028f565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061037a57505050505090565b90919293946020806103b6837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516102c3565b9701930193019193929061036b565b3461023d576103d336610242565b6103dc81613073565b9060005b8181106103f957604051806103f58582610347565b0390f35b61040d61040782848761310f565b35613d66565b6104178285613124565b526104228184613124565b50608061042f8285613124565b5101511561043f576001016103e0565b610449918461310f565b357fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176104c257604052565b610477565b60a0810190811067ffffffffffffffff8211176104c257604052565b60c0810190811067ffffffffffffffff8211176104c257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176104c257604052565b6040519061054f6040836104ff565b565b6040519061054f60c0836104ff565b6040519061054f610100836104ff565b6040519061054f60e0836104ff565b6040519061054f610140836104ff565b67ffffffffffffffff81116104c257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926105d58261058f565b916105e360405193846104ff565b82948184528183011161023d578281602093846000960137010152565b9080601f8301121561023d57816020610344933591016105c9565b3461023d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760043567ffffffffffffffff811161023d5761066a903690600401610600565b60243567ffffffffffffffff811161023d57602091610690610696923690600401610600565b90613138565b604051908152f35b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760206106e86004356000526007602052604060002054151590565b6040519015158152f35b600091031261023d57565b60005b8381106107105750506000910152565b8181015183820152602001610700565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361075c815180928187528780880191016106fd565b0116010190565b906020610344928181520190610720565b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d576103f560408051906107b581836104ff565b601a82527f4361706162696c6974696573526567697374727920312e312e30000000000000602083015251918291602083526020830190610720565b3461023d576107ff36610242565b906108086143ef565b6000905b82821061081557005b61082861082383858461310f565b613192565b926108438463ffffffff16600052600e602052604060002090565b9283549261088161087761085e8663ffffffff9060201c1690565b600288019063ffffffff16600052602052604060002090565b9460501c60ff1690565b9260005b855481101561090b5760019085156108e5576108df6108b76108a7838a6153e3565b600052600d602052604060002090565b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff8154169055565b01610885565b61090560056108f76108a7848b6153e3565b0163ffffffff8b169061562b565b506108df565b5095925093909250610931610928825463ffffffff9060201c1690565b63ffffffff1690565b15610a4457906109706109486001809594016131ef565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008154169055565b61099261098d8263ffffffff16600052600f602052604060002090565b61333d565b8051610a02575b506109bc6109b78263ffffffff16600052600e602052604060002090565b6134a5565b7ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365163ffffffff604051921691806109f881906000602083019252565b0390a2019061080c565b610a0e610a1b91613358565b63ffffffff83169061562b565b50610a3e610a398263ffffffff16600052600f602052604060002090565b613426565b38610999565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff821660045260246000fd5b6000fd5b9080602083519182815201916020808360051b8301019401926000915b838310610aa857505050505090565b9091929394602080610af5837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b518051845201519181858201520190610720565b97019301930191939290610a99565b805163ffffffff1682526103449160208281015163ffffffff169082015260408281015160ff1690820152606082810151151590820152608082810151151590820152610120610ba6610b92610b80610b6e60a087015161014060a088015261014087019061028f565b60c087015186820360c0880152610720565b60e086015185820360e0870152610720565b610100850151848203610100860152610720565b92015190610120818403910152610a7c565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610beb57505050505090565b9091929394602080610c27837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610b04565b97019301930191939290610bdc565b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5763ffffffff60115460201c16610c85610c80610928836134b5565b613552565b60009160015b8163ffffffff821610610cc5575091610ca96109286103f5946134b5565b8103610cbd575b5060405191829182610bb8565b815238610cb0565b610cef610928610ce58363ffffffff16600052600e602052604060002090565b5463ffffffff1690565b610d02575b60010163ffffffff16610c8b565b926001610d3263ffffffff92610d178761446e565b610d218288613124565b52610d2c8187613124565b506135bf565b94915050610cf4565b63ffffffff81160361023d57565b90604060206103449373ffffffffffffffffffffffffffffffffffffffff81511684520151918160208201520190610720565b906020610344928181520190610d49565b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5763ffffffff600435610dcd81610d3b565b610dd56135ec565b5016600052600c6020526103f560406000206001610e2a60405192610df9846104a6565b73ffffffffffffffffffffffffffffffffffffffff8154168452610e23604051809481930161329b565b03826104ff565b602082015260405191829182610d7c565b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d576103f5610e8a600435610e7c81610d3b565b610e846134f1565b5061446e565b604051918291602083526020830190610b04565b3461023d57610eac36610242565b610eb46143ef565b60005b63ffffffff81169082821015610f315763ffffffff610eda610f2c93858761310f565b35610ee481610d3b565b1680600052600c602052610f02600160406000206000815501613426565b7fa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a600080a2613651565b610eb7565b005b3461023d57610f4136610242565b90610f7d610f6473ffffffffffffffffffffffffffffffffffffffff6000541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b3314159160005b818110610f8d57005b610f9881838561310f565b3590610fae82600052600d602052604060002090565b600181015480156111555760058201805461110f5750815463ffffffff604082901c16806110d75750879081611086575b506110585760019361102260027f5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb97532059461101a61104f9561548a565b500154615564565b5061103f61103a82600052600d602052604060002090565b613668565b6040519081529081906020820190565b0390a101610f84565b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6110ce91506110b461109e610f649263ffffffff1690565b63ffffffff16600052600c602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b33141538610fdf565b7f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602485905260446000fd5b8461112061092886610a78946153e3565b7f60a6d8980000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b7fd82f6adb00000000000000000000000000000000000000000000000000000000600052600484905260246000fd5b3461023d5761119236610242565b6111b4610f6473ffffffffffffffffffffffffffffffffffffffff6000541690565b33149060009115925b8183106111c657005b6111d96111d48484846136b9565b6136f9565b9260408401936111f48551600052600d602052604060002090565b9161120e61120961109e855463ffffffff1690565b613606565b90600184018054928315611714578990816116ea575b506110585760208401928351156116c0578351808203611684575b505050606083018051801561165757506080840151988951156116235795611277611272875463ffffffff9060201c1690565b613651565b86547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff16602082901b67ffffffff000000001617875597600487019760005b8c5181101561135f576112e86112e46112d08f8490613124565b516000526005602052604060002054151590565b1590565b61132757806113208e6113196001948f8f90611313919063ffffffff16600052602052604060002090565b92613124565b51906157e6565b50016112b6565b6040517f3748d4c60000000000000000000000000000000000000000000000000000000081528061135b8f600483016137b4565b0390fd5b508754919b509299989796949593919060401c63ffffffff1663ffffffff811661158f575b5097919861139460058601613839565b9260009a5b84519c63ffffffff8d169d8e10156114e55761142f60036114299e9f6113c26113cc918a613124565b5163ffffffff1690565b9e8f61141561140860026113f08463ffffffff16600052600e602052604060002090565b019263ffffffff16600052600e602052604060002090565b5460201c63ffffffff1690565b63ffffffff16600052602052604060002090565b01613839565b9b60005b8d518110156114cb5761147a6112e48f8f8f6113136114639287929063ffffffff16600052602052604060002090565b519060019160005201602052604060002054151590565b61148657600101611433565b8e6114958f92610a7893613124565b517f03dcd8620000000000000000000000000000000000000000000000000000000060005260045263ffffffff16602452604490565b50939c509c6114dc91929b50613651565b9a919990611399565b9b509b50925094600197506115849196507f4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b94600361152d63ffffffff965163ffffffff1690565b82547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8216178355945197886002840155519101555160405193849316958360209093929193604081019481520152565b0390a20191906111bd565b999192989a6115c360036114298d61141561140860026113f0849e9c9d9e63ffffffff16600052600e602052604060002090565b9960005b8b51811015611611576115f76112e48d611463848f8f90611313919063ffffffff16600052602052604060002090565b611603576001016115c7565b8c611495610a78928e613124565b5093929a509a98509493929438611384565b6040517f3748d4c60000000000000000000000000000000000000000000000000000000081528061135b8c600483016137b4565b7f37d897650000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61169b906000526009602052604060002054151590565b6116c0576116ac918451905561548a565b506116b7825161570c565b5038808061123f565b7f837731460000000000000000000000000000000000000000000000000000000060005260046000fd5b5161170b915073ffffffffffffffffffffffffffffffffffffffff16610f64565b33141538611224565b88517fd82f6adb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461023d5761175136610242565b9061175a6143ef565b60005b82811061176657005b611779611774828585613884565b6138c4565b9061179b610f64835173ffffffffffffffffffffffffffffffffffffffff1690565b156118e4576001916117b260115463ffffffff1690565b907f78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e73ffffffffffffffffffffffffffffffffffffffff6118c061180a845173ffffffffffffffffffffffffffffffffffffffff1690565b936118606020820195865161183c611820610540565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015261185b8863ffffffff16600052600c602052604060002090565b613a7b565b6118a661187561127260115463ffffffff1690565b63ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000006011541617601155565b5173ffffffffffffffffffffffffffffffffffffffff1690565b9251926118db63ffffffff6040519384931696169482610763565b0390a30161175d565b7feeacd9390000000000000000000000000000000000000000000000000000000060005260046000fd5b909161192561034493604084526040840190610720565b916020818403910152610720565b3461023d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760043561196e81610d3b565b602435906119dd61098d8360046119cd61199b6114088763ffffffff16600052600e602052604060002090565b60026119b78863ffffffff16600052600e602052604060002090565b019063ffffffff16600052602052604060002090565b0190600052602052604060002090565b9060609273ffffffffffffffffffffffffffffffffffffffff611a2d6002611a0f846000526003602052604060002090565b015460101c73ffffffffffffffffffffffffffffffffffffffff1690565b16611a44575b50506103f56040519283928361190e565b611aad92935090611a6c610f64610f646002611a0f6000966000526003602052604060002090565b60405180809581947f8318ed5d0000000000000000000000000000000000000000000000000000000083526004830191909163ffffffff6020820193169052565b03915afa908115611aec57600091611ac9575b50903880611a33565b611ae691503d806000833e611ade81836104ff565b810190613bbe565b38611ac0565b613c1d565b60643590811515820361023d57565b60443590811515820361023d57565b6084359060ff8216820361023d57565b9081604091031261023d5790565b3461023d5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57600435611b6881610d3b565b60243567ffffffffffffffff811161023d57611b8890369060040161020c565b60449291923567ffffffffffffffff811161023d57611bab90369060040161020c565b611bb3611af1565b91611bbc611b0f565b9360a4359667ffffffffffffffff881161023d57611be1610f31983690600401611b1f565b96613c29565b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d576103f5611c24600435613d66565b6040519182916020835260208301906102c3565b9181601f8401121561023d5782359167ffffffffffffffff831161023d576020838186019501011161023d57565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261023d576004359067ffffffffffffffff821161023d5761028b91600401611c38565b3461023d57611cce611cc9611cc336611c66565b90613eaf565b613839565b60405180916020820160208352815180915260206040840192019060005b818110611cfa575050500390f35b8251845285945060209384019390920191600101611cec565b3461023d57611d2136610242565b611d296143ef565b60005b818110611d3557005b611d4081838561310f565b35611d58816000526005602052604060002054151590565b15611dc857611d6681615747565b15611d9a57906001917fdcea1b78b6ddc31592a94607d537543fcaafda6cc52d6d5cc7bbfca1422baf21600080a201611d2c565b7ff7d7a294000000000000000000000000000000000000000000000000000000006000526024906004526000fd5b7fe181733f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611e2857505050505090565b9091929394602080611e64837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610d49565b97019301930191939290611e19565b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760115463ffffffff16611ebf611eba610928836134b5565b613ec8565b60009163ffffffff811660015b8163ffffffff821610611f06576103f58486611eea610928876134b5565b8103611efe575b5060405191829182611df5565b815282611ef1565b611f26610f646110b48363ffffffff16600052600c602052604060002090565b611f39575b60010163ffffffff16611ecc565b936001611f7763ffffffff92611f626112098963ffffffff16600052600c602052604060002090565b611f6c8289613124565b52610d2c8188613124565b95915050611f2b565b3461023d57611f8e36610242565b611f966143ef565b60005b818110611fa257005b611fad8183856136b9565b60a08136031261023d5760405190611fc4826104c7565b803567ffffffffffffffff811161023d57611fe29036908301610600565b8252602081013567ffffffffffffffff811161023d576120059036908301610600565b806020840152604082013591600483101561023d5760806120409161204c94604087015261203560608201613f35565b606087015201612e62565b60808401528251613138565b6120586112e48261577c565b61206f576001929161206991615268565b01611f99565b7febf525510000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5773ffffffffffffffffffffffffffffffffffffffff6001541633036121995760005473ffffffffffffffffffffffffffffffffffffffff16600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163317905561215a7fffffffffffffffffffffffff000000000000000000000000000000000000000060015416600155565b73ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b3461023d5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760043567ffffffffffffffff811161023d5761224690369060040161020c565b9060243567ffffffffffffffff811161023d5761226790369060040161020c565b61226f611b00565b612277611af1565b91612280611b0f565b9360a4359667ffffffffffffffff881161023d576122a5610f31983690600401611b1f565b96613f42565b3461023d576122b936610242565b6122db610f6473ffffffffffffffffffffffffffffffffffffffff6000541690565b33149060009115925b8183106122ed57005b6122fb6111d48484846136b9565b9261231061120961109e865163ffffffff1690565b612331610f64825173ffffffffffffffffffffffffffffffffffffffff1690565b1561263c57859081612612575b5061105857604084019361235d8551600052600d602052604060002090565b94600186019081546125e457805180156125b75750602083019182518015908115612599575b506116c05760608401978851801561165757506080850151998a511561256557966123f56123bc611272845463ffffffff9060201c1690565b83547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1660209190911b67ffffffff0000000016178355565b815460201c63ffffffff1698600483019860005b8d5181101561248b576124236112e48f836112d091613124565b612455578061244e8f6113198f948f600196611313919063ffffffff16600052602052604060002090565b5001612409565b61135b8e6040519182917f3748d4c6000000000000000000000000000000000000000000000000000000008352600483016137b4565b509a63ffffffff95919c5061253c929a50600199507f74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f05969461255994612524925160038201556125136124e28b5163ffffffff1690565b829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60028451910155855180915561570c565b5061252f81516157b1565b5051955163ffffffff1690565b915160405193849316958360209093929193604081019481520152565b0390a2019192906122e4565b6040517f3748d4c60000000000000000000000000000000000000000000000000000000081528061135b8d600483016137b4565b6125b191506000526009602052604060002054151590565b38612383565b7f64e2ee920000000000000000000000000000000000000000000000000000000060005260045260246000fd5b517f546184830000000000000000000000000000000000000000000000000000000060005260045260246000fd5b51612633915073ffffffffffffffffffffffffffffffffffffffff16610f64565b3314153861233e565b610a7861264d865163ffffffff1690565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b3461023d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5760043567ffffffffffffffff811161023d576126ce90369060040161020c565b60243567ffffffffffffffff811161023d576126ee90369060040161020c565b928383036129aa5760005473ffffffffffffffffffffffffffffffffffffffff169060005b84811061271c57005b61272a61082382878561310f565b906127458263ffffffff16600052600c602052604060002090565b91612764835473ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff8416801561297657612790611774858c8b613884565b906127b2610f64835173ffffffffffffffffffffffffffffffffffffffff1690565b156118e45733141580612956575b6110585760019473ffffffffffffffffffffffffffffffffffffffff6127fd610f64845173ffffffffffffffffffffffffffffffffffffffff1690565b9116148015906128ea575b612816575b50505001612713565b73ffffffffffffffffffffffffffffffffffffffff6128c4826128b26128717f86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28955173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6118a660208201958987519101613954565b9251926128df63ffffffff6040519384931696169482610763565b0390a338808061280d565b50604051602081019061292f8161290389870185614004565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826104ff565b519020602082015160405161294c81612903602082019485610763565b5190201415612808565b5073ffffffffffffffffffffffffffffffffffffffff87163314156127c0565b7fadd9ae1e0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b7fab8b67c6000000000000000000000000000000000000000000000000000000006000526004839052602484905260446000fd5b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60041115612a6957565b612a30565b60021115612a6957565b906002821015612a695752565b80518252612ab7612aa5602083015160e0602086015260e0850190610720565b60408301518482036040860152610720565b916060820151916004831015612a695760c08091610344946060850152612ae660808201516080860190612a78565b60a08181015173ffffffffffffffffffffffffffffffffffffffff169085015201511515910152565b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d576103f5612b4c600435614077565b604051918291602083526020830190612a85565b3461023d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57600435612b9b81610d3b565b6024359067ffffffffffffffff821161023d57612bbf610f31923690600401611c38565b91612bc86143ef565b614292565b3461023d57602063ffffffff81612be336611c66565b91908260405193849283378101600281520301902054161515604051908152f35b602081016020825282518091526040820191602060408360051b8301019401926000915b838310612c3757505050505090565b9091929394602080612c73837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951612a85565b97019301930191939290612c28565b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57612cb96137ee565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612cff612ce984613007565b93612cf760405195866104ff565b808552613007565b0160005b818110612d5857505060005b8151811015612d4a5780612d2e612d2860019385613124565b51614077565b612d388286613124565b52612d438185613124565b5001612d0f565b604051806103f58582612c04565b602090612d63614015565b82828701015201612d03565b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57604051600a548082528160208101600a60005260206000209260005b818110612e26575050612dcf925003826104ff565b612dd98151613073565b9060005b8151811015612e185780612dfc612df660019385613124565b51613d66565b612e068286613124565b52612e118185613124565b5001612ddd565b604051806103f58582610347565b8454835260019485019486945060209093019201612dba565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361023d57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361023d57565b3461023d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d5773ffffffffffffffffffffffffffffffffffffffff612ecf612e3f565b612ed76143ef565b16338114612f6557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff612f3d60005473ffffffffffffffffffffffffffffffffffffffff1690565b167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b3461023d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261023d57602063ffffffff601154821c16604051908152f35b67ffffffffffffffff81116104c25760051b60200190565b60405190610100820182811067ffffffffffffffff8211176104c257604052606060e08360008152600060208201526000604082015260008382015260006080820152600060a08201528260c08201520152565b9061307d82613007565b61308a60405191826104ff565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06130b88294613007565b019060005b8281106130c957505050565b6020906130d461301f565b828285010152016130bd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561311f5760051b0190565b6130e0565b805182101561311f5760209160051b010190565b9061318c6129039160405192839161315c6020840196604088526060850190610720565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152610720565b51902090565b3561034481610d3b565b90600182811c921680156131e5575b60208310146131b657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916131ab565b604051908160008254926132028461319c565b93600181169081156132635750600114613227575b5060209250600281520301902090565b9150506000528160206000206000905b83821061324b575050602091810138613217565b60209192508060019154848701520191018391613237565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168452506020938015150283019150389050613217565b600092918154916132ab8361319c565b808352926001811690811561330157506001146132c757505050565b60009081526020812093945091925b8383106132e7575060209250010190565b6001816020929493945483858701015201910191906132d6565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b9061054f613351926040519384809261329b565b03836104ff565b60206133719181604051938285809451938492016106fd565b8101601081520301902090565b60206133979181604051938285809451938492016106fd565b8101600281520301902090565b916133dc918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8181106133eb575050565b600081556001016133e0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b613430815461319c565b908161343a575050565b81601f6000931160011461344c575055565b8183526020832061346891601f0160051c8101906001016133e0565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555565b600161054f916000815501613426565b63ffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9116019063ffffffff82116134ec57565b6133f7565b60405190610140820182811067ffffffffffffffff8211176104c257604052606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b9061355c82613007565b61356960405191826104ff565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06135978294613007565b019060005b8281106135a857505050565b6020906135b36134f1565b8282850101520161359c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146134ec5760010190565b604051906135f9826104a6565b6060602083600081520152565b9060016020604051613617816104a6565b61364d819573ffffffffffffffffffffffffffffffffffffffff8154168352613646604051809681930161329b565b03846104ff565b0152565b63ffffffff1663ffffffff81146134ec5760010190565b60059060008155600060018201556000600282015560006003820155018054906000815581613695575050565b6000526020600020908101905b8181106136ad575050565b600081556001016136a2565b919081101561311f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561023d570190565b60a08136031261023d5760405190613710826104c7565b803561371b81610d3b565b825260208101356020830152604081013560408301526060810135606083015260808101359067ffffffffffffffff821161023d570136601f8201121561023d5780359061376882613007565b9161377660405193846104ff565b80835260208084019160051b8301019136831161023d57602001905b8282106137a457505050608082015290565b8135815260209182019101613792565b602060408183019282815284518094520192019060005b8181106137d85750505090565b82518452602093840193909201916001016137cb565b604051906004548083528260208101600460005260206000209260005b81811061382057505061054f925003836104ff565b845483526001948501948794506020909301920161380b565b906040519182815491828252602082019060005260206000209260005b81811061386b57505061054f925003836104ff565b8454835260019485019487945060209093019201613856565b919081101561311f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18136030182121561023d570190565b60408136031261023d57604051906138db826104a6565b6138e481612e62565b825260208101359067ffffffffffffffff821161023d5761390791369101610600565b602082015290565b9190601f811161391e57505050565b61054f926000526020600020906020601f840160051c8301931061394a575b601f0160051c01906133e0565b909150819061393d565b919091825167ffffffffffffffff81116104c25761397c81613976845461319c565b8461390f565b6020601f82116001146139d65781906133dc9394956000926139cb575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880613999565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690613a0984600052602060002090565b9160005b818110613a6357509583600195969710613a2c575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613a22565b9192602060018192868b015181550194019201613a0d565b6001602091939293613adf73ffffffffffffffffffffffffffffffffffffffff865116829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b0192015191825167ffffffffffffffff81116104c257613b0381613976845461319c565b6020601f8211600114613b515781906133dc9394956000926139cb5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690613b8484600052602060002090565b9160005b818110613ba657509583600195969710613a2c57505050811b019055565b9192602060018192868b015181550194019201613b88565b60208183031261023d5780519067ffffffffffffffff821161023d570181601f8201121561023d578051613bf18161058f565b92613bff60405194856104ff565b8184526020828401011161023d5761034491602080850191016106fd565b6040513d6000823e3d90fd5b95909691929394613c386143ef565b63ffffffff8716600052600e6020526040600020549063ffffffff8260201c1663ffffffff811615613cc157613c6d90613651565b9160501c60ff1691613c7d610551565b63ffffffff909916895263ffffffff1660208901525b151560408801521515606087015260ff16608086015236613cb391613cf5565b60a085015261054f9461475d565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff891660045260246000fd5b919060408382031261023d5760405190613d0e826104a6565b8193803567ffffffffffffffff811161023d5782613d2d918301610600565b835260208101359067ffffffffffffffff821161023d570181601f8201121561023d576020918183613d61933591016105c9565b910152565b613d6e61301f565b50613d86610ce582600052600d602052604060002090565b906002613d9d82600052600d602052604060002090565b01546001613db583600052600d602052604060002090565b0154906003613dce84600052600d602052604060002090565b015490613e05611cc96004613ded87600052600d602052604060002090565b0161141561140888600052600d602052604060002090565b92613e93613e2061140887600052600d602052604060002090565b95613e86613e616005611429613e50613e4386600052600d602052604060002090565b5460401c63ffffffff1690565b94600052600d602052604060002090565b97613e79613e6d610560565b63ffffffff909b168b52565b63ffffffff1660208a0152565b63ffffffff166040880152565b6060860152608085015260a084015260c083015260e082015290565b6020908260405193849283378101601081520301902090565b90613ed282613007565b613edf60405191826104ff565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613f0d8294613007565b019060005b828110613f1e57505050565b602090613f296135ec565b82828501015201613f12565b3590600282101561023d57565b96919395909294613f516143ef565b60115460201c63ffffffff1696613f6788613651565b6011549060201b67ffffffff0000000016907fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff161760115587600052600e60205287604060002090613fe3919063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60405197613ff0896104e3565b63ffffffff16885260016020890152613c93565b90602061034492818152019061329b565b6040519060e0820182811067ffffffffffffffff8211176104c257604052600060c08382815260606020820152606060408201528260608201528260808201528260a08201520152565b6004821015612a695752565b6002821015612a695752565b61407f614015565b506000818152600360205260409020610344906141586001820161413b6140aa600285015460ff1690565b6141326140d160026140c68a6000526003602052604060002090565b015460081c60ff1690565b916141246140ee6002611a0f8c6000526003602052604060002090565b9561411a6141098c6000526007602052604060002054151590565b99614112610570565b9c8d5261333d565b60208c015261333d565b60408a01526060890161405f565b6080870161406b565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b151560c0830152565b90929167ffffffffffffffff81116104c25761418181613976845461319c565b6000601f82116001146141da5781906133dc9394956000926141cf5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b013590503880613999565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461420d84600052602060002090565b91805b87811061426557508360019596971061422d57505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080613a22565b90926020600181928686013581550194019101614210565b81604051928392833781016000815203902090565b9163ffffffff6142b56114088563ffffffff16600052600e602052604060002090565b16156143bb576142d861098d8463ffffffff16600052600f602052604060002090565b90815160208301206142eb3683866105c9565b60208151910120146143b55763ffffffff92826143309351614399575b50816143595761432b610a398663ffffffff16600052600f602052604060002090565b61427d565b91167f3143778fe5911daca9ced3c73dde73fdb95e7edae557614c40588f3f1a1b44b8600080a3565b61437d82826143788863ffffffff16600052600f602052604060002090565b614161565b61439361438a8383613eaf565b858716906157e6565b5061427d565b6143a56143ae91613358565b8587169061562b565b5038614308565b50505050565b7f2b62be9b0000000000000000000000000000000000000000000000000000000060005263ffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60005416330361441057565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fd5b6144766134f1565b506144946114088263ffffffff16600052600e602052604060002090565b6144b48160026119b78563ffffffff16600052600e602052604060002090565b6144c060038201613839565b936144cb8551613ec8565b92600483019260005b855181101561453d57806144ea6001928a613124565b5161451c61450c6144fb848d613124565b518990600052602052604060002090565b614514610540565b92835261333d565b602082015261452b8289613124565b526145368188613124565b50016144d4565b5092509390945061455e8163ffffffff16600052600e602052604060002090565b5463ffffffff16936145808263ffffffff16600052600e602052604060002090565b6001019161459e8163ffffffff16600052600e602052604060002090565b5460401c60ff166145bf8263ffffffff16600052600e602052604060002090565b5460481c60ff166145e08363ffffffff16600052600e602052604060002090565b5460501c60ff16916145f187613839565b9361460c9063ffffffff16600052600f602052604060002090565b9461461561057f565b63ffffffff909a168a5263ffffffff1660208a015260ff166040890152151560608801521515608087015260a086015261464e9061333d565b60c085015261465c9061333d565b60e084015260020161466d9061333d565b61010083015261012082015290565b60ff60019116019060ff82116134ec57565b610344905461319c565b805482101561311f5760005260206000200190600090565b805490680100000000000000008210156104c257816146d79160016133dc94018155614698565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561023d570180359067ffffffffffffffff821161023d5760200191813603831361023d57565b91939290614788614772855163ffffffff1690565b63ffffffff16600052600e602052604060002090565b92600261479c614772875163ffffffff1690565b01956147b46020870197611415895163ffffffff1690565b9060808701956147c5875160ff1690565b8560ff821615918215615007575b5050614fc6576001016147e58161333d565b602081519101209060a08901918251516020815191012003614f10575b50600163ffffffff6148188b5163ffffffff1690565b1611614e77575b60009691965b858110614cba575060009160048401966002600386019501935b8181106148a9575050505050505050505063ffffffff6148926148877ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651935163ffffffff1690565b935163ffffffff1690565b60405163ffffffff919091168152921691602090a2565b6148b481838c613884565b8035906148d16112e4836000526005602052604060002054151590565b614c8b576148ec826000526007602052604060002054151590565b614c5c5761490c614907838d90600052602052604060002090565b61468e565b614c155789898f938f938e948c60005b868110614b6257508361492e916146b0565b60208201958361493e888561470c565b91614953919390600052602052604060002090565b9161495d92614161565b8a516020015161496d908d613954565b60408101511515815163ffffffff166149969063ffffffff16600052600e602052604060002090565b80547fffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffff1691151560481b69ff0000000000000000001691909117905560608101511515815163ffffffff166149fb9063ffffffff16600052600e602052604060002090565b80547fffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffff1691151560501b6aff0000000000000000000016919091179055895160ff16815163ffffffff16614a5f9063ffffffff16600052600e602052604060002090565b90614a9e91907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff68ff000000000000000083549260401b169116179055565b865163ffffffff16815163ffffffff16614ac89063ffffffff16600052600e602052604060002090565b90614b0691907fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff67ffffffff0000000083549260201b169116179055565b8a5151815163ffffffff166000908152600e6020526040902060010190614b2c91613954565b5163ffffffff16955163ffffffff1694614b459161470c565b3690614b50926105c9565b93614b5a9561588d565b60010161483f565b958498506112e494939750614bab9250859150611408614b9388836004614ba4614b938461141598614bc09e61310f565b35600052600d602052604060002090565b019561310f565b60019160005201602052604060002054151590565b614bd657600101898f938f938e948c8f9561491c565b82614be5610a78928d8d61310f565b7fa7e792500000000000000000000000000000000000000000000000000000000060005235600452602452604490565b610a7882614c278f5163ffffffff1690565b7f3927d0800000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b7ff7d7a29400000000000000000000000000000000000000000000000000000000600052600482905260246000fd5b7fe181733f00000000000000000000000000000000000000000000000000000000600052600482905260246000fd5b614cd56112e4614cce8389899c979c61310f565b35866157e6565b614e2857606089015115614df457614cf4613e43614b9383898961310f565b63ffffffff614d0a6109288c5163ffffffff1690565b9116141580614dd5575b614d8057600190614d77614d2c8b5163ffffffff1690565b614d3a614b93848b8b61310f565b907fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff6bffffffff000000000000000083549260401b169116179055565b01969196614825565b84614d9d610a789288614d978d5163ffffffff1690565b9361310f565b7f60b9df730000000000000000000000000000000000000000000000000000000060005263ffffffff90911660045235602452604490565b5063ffffffff614dec613e43614b93848a8a61310f565b161515614d14565b80614e226005614e0a614b936001958b8b61310f565b01614e1c6109288d5163ffffffff1690565b906157e6565b50614d77565b84614e3f610a789288614d978d5163ffffffff1690565b7f636e40570000000000000000000000000000000000000000000000000000000060005263ffffffff90911660045235602452604490565b9794929098959391614eac6002614e956147728b5163ffffffff1690565b01611415614ea7895163ffffffff1690565b6134b5565b9960005b8b54811015614f005780614efa6108b76108a78f94858f614ef490614eee6109286005614ee26108a78860019e6153e3565b01925163ffffffff1690565b9061562b565b506153e3565b01614eb0565b509193959890929496995061481f565b610948614f1c916131ef565b80515180511561480257614f37610ce563ffffffff9261337e565b16614f8e57614f88614f4d895163ffffffff1690565b614f5883515161337e565b9063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b38614802565b61135b9051516040519182917f07bf02d600000000000000000000000000000000000000000000000000000000835260048301610763565b610a7885614fd5895160ff1690565b7f25b4d6180000000000000000000000000000000000000000000000000000000060005260ff16600452602452604490565b60ff9192506150159061467c565b161185386147d3565b906004811015612a695760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b906002811015612a69577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff61ff0083549260081b169116179055565b9190805192835167ffffffffffffffff81116104c2576150b581613976845461319c565b6020601f82116001146151b35761516d9261510e8360809460029461054f999a6000926139cb5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b81555b615122602086015160018301613954565b019261513b604082015161513581612a5f565b8561501e565b615152606082015161514c81612a6e565b85615055565b015173ffffffffffffffffffffffffffffffffffffffff1690565b7fffffffffffffffffffff0000000000000000000000000000000000000000ffff75ffffffffffffffffffffffffffffffffffffffff000083549260101b169116179055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216956151e684600052602060002090565b9660005b81811061525057508361054f979860029461516d979460809760019510615219575b505050811b018155615111565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061520c565b838301518955600190980197602093840193016151ea565b906080810173ffffffffffffffffffffffffffffffffffffffff815116806152d1575b50506152aa906152a5836000526003602052604060002090565b615091565b7f04f0a9bcf3f3a3b42a4d7ca081119755f82ebe43e0d30c8f7292c4fe0dc4a2ae600080a2565b6153639060206000604051828101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526153386044826104ff565b519084617530fa903d60005190836153d7575b50826153cd575b50816153bb575b816153b157501590565b61536d578061528b565b517fabb5e3fd0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6112e49150615a2c565b90506153c68161599b565b1590615359565b1515915038615352565b6020111592503861534b565b906153ed91614698565b90549060031b1c90565b8054801561545b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061542c8282614698565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526009602052604090205490811561555d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134ec57600854927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134ec578383615526946000960361552c575b50505061551560086153f7565b600990600052602052604060002090565b55600190565b61551561554e916155446153ed615554956008614698565b9283916008614698565b906133a4565b55388080615508565b5050600090565b6000818152600b602052604090205490811561555d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134ec57600a54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134ec5783836155269460009603615600575b5050506155ef600a6153f7565b600b90600052602052604060002090565b6155ef61554e916156186153ed61562295600a614698565b928391600a614698565b553880806155e2565b6001810191806000528260205260406000205492831515600014615703577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116134ec578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116134ec5760009585836156bb9461552698036156ca575b5050506153f7565b90600052602052604060002090565b6156ea61554e916156e16153ed6156fa9588614698565b92839187614698565b8590600052602052604060002090565b553880806156b3565b50505050600090565b6000818152600960205260409020546157415761572a8160086146b0565b600854906000526009602052604060002055600190565b50600090565b600081815260076020526040902054615741576157658160066146b0565b600654906000526007602052604060002055600190565b6000818152600560205260409020546157415761579a8160046146b0565b600454906000526005602052604060002055600190565b6000818152600b6020526040902054615741576157cf81600a6146b0565b600a5490600052600b602052604060002055600190565b600082815260018201602052604090205461555d5780615808836001936146b0565b80549260005201602052604060002055600190565b939594909291608085528060808601527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811161023d5761588063ffffffff9360a0606096869460051b8091838b01378801818982030160208a01520190610720565b9616604085015216910152565b91949293909473ffffffffffffffffffffffffffffffffffffffff60026158be836000526003602052604060002090565b015460101c166158d1575b505050505050565b610f6473ffffffffffffffffffffffffffffffffffffffff6002615902615921946000526003602052604060002090565b015460101c1673ffffffffffffffffffffffffffffffffffffffff1690565b91823b1561023d576000946159678692604051988997889687957ffba64a7c0000000000000000000000000000000000000000000000000000000087526004870161581d565b03925af18015611aec57615980575b80808080806158c9565b8061598f6000615995936104ff565b806106f2565b38615976565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526159ff6044826104ff565b5191617530fa6000513d82615a20575b5081615a19575090565b9050151590565b60201115915038615a0f565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f78bea721000000000000000000000000000000000000000000000000000000006024820152602481526159ff6044826104ff56fea164736f6c634300081a000a",
}

var CapabilitiesRegistryABI = CapabilitiesRegistryMetaData.ABI

var CapabilitiesRegistryBin = CapabilitiesRegistryMetaData.Bin

func DeployCapabilitiesRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CapabilitiesRegistry, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CapabilitiesRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CapabilitiesRegistry{address: address, abi: *parsed, CapabilitiesRegistryCaller: CapabilitiesRegistryCaller{contract: contract}, CapabilitiesRegistryTransactor: CapabilitiesRegistryTransactor{contract: contract}, CapabilitiesRegistryFilterer: CapabilitiesRegistryFilterer{contract: contract}}, nil
}

type CapabilitiesRegistry struct {
	address common.Address
	abi     abi.ABI
	CapabilitiesRegistryCaller
	CapabilitiesRegistryTransactor
	CapabilitiesRegistryFilterer
}

type CapabilitiesRegistryCaller struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistryTransactor struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistryFilterer struct {
	contract *bind.BoundContract
}

type CapabilitiesRegistrySession struct {
	Contract     *CapabilitiesRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CapabilitiesRegistryCallerSession struct {
	Contract *CapabilitiesRegistryCaller
	CallOpts bind.CallOpts
}

type CapabilitiesRegistryTransactorSession struct {
	Contract     *CapabilitiesRegistryTransactor
	TransactOpts bind.TransactOpts
}

type CapabilitiesRegistryRaw struct {
	Contract *CapabilitiesRegistry
}

type CapabilitiesRegistryCallerRaw struct {
	Contract *CapabilitiesRegistryCaller
}

type CapabilitiesRegistryTransactorRaw struct {
	Contract *CapabilitiesRegistryTransactor
}

func NewCapabilitiesRegistry(address common.Address, backend bind.ContractBackend) (*CapabilitiesRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(CapabilitiesRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCapabilitiesRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistry{address: address, abi: abi, CapabilitiesRegistryCaller: CapabilitiesRegistryCaller{contract: contract}, CapabilitiesRegistryTransactor: CapabilitiesRegistryTransactor{contract: contract}, CapabilitiesRegistryFilterer: CapabilitiesRegistryFilterer{contract: contract}}, nil
}

func NewCapabilitiesRegistryCaller(address common.Address, caller bind.ContractCaller) (*CapabilitiesRegistryCaller, error) {
	contract, err := bindCapabilitiesRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCaller{contract: contract}, nil
}

func NewCapabilitiesRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CapabilitiesRegistryTransactor, error) {
	contract, err := bindCapabilitiesRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryTransactor{contract: contract}, nil
}

func NewCapabilitiesRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CapabilitiesRegistryFilterer, error) {
	contract, err := bindCapabilitiesRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryFilterer{contract: contract}, nil
}

func bindCapabilitiesRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryTransactor.contract.Transfer(opts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.CapabilitiesRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CapabilitiesRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.contract.Transfer(opts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapabilities(opts *bind.CallOpts) ([]CapabilitiesRegistryCapabilityInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapabilities")

	if err != nil {
		return *new([]CapabilitiesRegistryCapabilityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryCapabilityInfo)).(*[]CapabilitiesRegistryCapabilityInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapabilities() ([]CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilities(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapabilities() ([]CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilities(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapability(opts *bind.CallOpts, hashedId [32]byte) (CapabilitiesRegistryCapabilityInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapability", hashedId)

	if err != nil {
		return *new(CapabilitiesRegistryCapabilityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryCapabilityInfo)).(*CapabilitiesRegistryCapabilityInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapability(hashedId [32]byte) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, hashedId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapability(hashedId [32]byte) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, hashedId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId [32]byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapabilityConfigs", donId, capabilityId)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapabilityConfigs(donId uint32, capabilityId [32]byte) ([]byte, []byte, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilityConfigs(&_CapabilitiesRegistry.CallOpts, donId, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapabilityConfigs(donId uint32, capabilityId [32]byte) ([]byte, []byte, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilityConfigs(&_CapabilitiesRegistry.CallOpts, donId, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDON(opts *bind.CallOpts, donId uint32) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDON", donId)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDON(donId uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDON(&_CapabilitiesRegistry.CallOpts, donId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDON(donId uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDON(&_CapabilitiesRegistry.CallOpts, donId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONs(opts *bind.CallOpts) ([]CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONs")

	if err != nil {
		return *new([]CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryDONInfo)).(*[]CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONs() ([]CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONs(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONs() ([]CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONs(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONsInFamily(opts *bind.CallOpts, donFamily string) ([]*big.Int, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONsInFamily", donFamily)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONsInFamily(donFamily string) ([]*big.Int, error) {
	return _CapabilitiesRegistry.Contract.GetDONsInFamily(&_CapabilitiesRegistry.CallOpts, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONsInFamily(donFamily string) ([]*big.Int, error) {
	return _CapabilitiesRegistry.Contract.GetDONsInFamily(&_CapabilitiesRegistry.CallOpts, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetHashedCapabilityId(opts *bind.CallOpts, labelledName string, version string) ([32]byte, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getHashedCapabilityId", labelledName, version)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetHashedCapabilityId(labelledName string, version string) ([32]byte, error) {
	return _CapabilitiesRegistry.Contract.GetHashedCapabilityId(&_CapabilitiesRegistry.CallOpts, labelledName, version)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetHashedCapabilityId(labelledName string, version string) ([32]byte, error) {
	return _CapabilitiesRegistry.Contract.GetHashedCapabilityId(&_CapabilitiesRegistry.CallOpts, labelledName, version)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNextDONId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNextDONId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNextDONId() (uint32, error) {
	return _CapabilitiesRegistry.Contract.GetNextDONId(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNextDONId() (uint32, error) {
	return _CapabilitiesRegistry.Contract.GetNextDONId(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNode(opts *bind.CallOpts, p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNode", p2pId)

	if err != nil {
		return *new(INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(INodeInfoProviderNodeInfo)).(*INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNode(p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNode(&_CapabilitiesRegistry.CallOpts, p2pId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNode(p2pId [32]byte) (INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNode(&_CapabilitiesRegistry.CallOpts, p2pId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(CapabilitiesRegistryNodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryNodeOperator)).(*CapabilitiesRegistryNodeOperator)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodeOperator(nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperator(&_CapabilitiesRegistry.CallOpts, nodeOperatorId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodeOperator(nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperator(&_CapabilitiesRegistry.CallOpts, nodeOperatorId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodeOperators(opts *bind.CallOpts) ([]CapabilitiesRegistryNodeOperator, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodeOperators")

	if err != nil {
		return *new([]CapabilitiesRegistryNodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([]CapabilitiesRegistryNodeOperator)).(*[]CapabilitiesRegistryNodeOperator)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodeOperators() ([]CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperators(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodeOperators() ([]CapabilitiesRegistryNodeOperator, error) {
	return _CapabilitiesRegistry.Contract.GetNodeOperators(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodes(opts *bind.CallOpts) ([]INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodes")

	if err != nil {
		return *new([]INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]INodeInfoProviderNodeInfo)).(*[]INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodes() ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodes(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodes() ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodes(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetNodesByP2PIds(opts *bind.CallOpts, p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getNodesByP2PIds", p2pIds)

	if err != nil {
		return *new([]INodeInfoProviderNodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]INodeInfoProviderNodeInfo)).(*[]INodeInfoProviderNodeInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetNodesByP2PIds(p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodesByP2PIds(&_CapabilitiesRegistry.CallOpts, p2pIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetNodesByP2PIds(p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error) {
	return _CapabilitiesRegistry.Contract.GetNodesByP2PIds(&_CapabilitiesRegistry.CallOpts, p2pIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) IsCapabilityDeprecated(opts *bind.CallOpts, hashedCapabilityId [32]byte) (bool, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "isCapabilityDeprecated", hashedCapabilityId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) IsCapabilityDeprecated(hashedCapabilityId [32]byte) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, hashedCapabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) IsCapabilityDeprecated(hashedCapabilityId [32]byte) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, hashedCapabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) IsDONNameTaken(opts *bind.CallOpts, donName string) (bool, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "isDONNameTaken", donName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) IsDONNameTaken(donName string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsDONNameTaken(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) IsDONNameTaken(donName string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsDONNameTaken(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) Owner() (common.Address, error) {
	return _CapabilitiesRegistry.Contract.Owner(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) Owner() (common.Address, error) {
	return _CapabilitiesRegistry.Contract.Owner(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) TypeAndVersion() (string, error) {
	return _CapabilitiesRegistry.Contract.TypeAndVersion(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) TypeAndVersion() (string, error) {
	return _CapabilitiesRegistry.Contract.TypeAndVersion(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AcceptOwnership(&_CapabilitiesRegistry.TransactOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AcceptOwnership(&_CapabilitiesRegistry.TransactOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddCapabilities(opts *bind.TransactOpts, capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addCapabilities", capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddCapabilities(capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddCapabilities(capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilities)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddDON(opts *bind.TransactOpts, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, acceptsWorkflows bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addDON", nodes, capabilityConfigurations, isPublic, acceptsWorkflows, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddDON(nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, acceptsWorkflows bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDON(&_CapabilitiesRegistry.TransactOpts, nodes, capabilityConfigurations, isPublic, acceptsWorkflows, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddDON(nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, acceptsWorkflows bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDON(&_CapabilitiesRegistry.TransactOpts, nodes, capabilityConfigurations, isPublic, acceptsWorkflows, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddNodeOperators(opts *bind.TransactOpts, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addNodeOperators", nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddNodeOperators(nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddNodeOperators(nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addNodes", nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) DeprecateCapabilities(opts *bind.TransactOpts, hashedCapabilityIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "deprecateCapabilities", hashedCapabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) DeprecateCapabilities(hashedCapabilityIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, hashedCapabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) DeprecateCapabilities(hashedCapabilityIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, hashedCapabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveDONs(opts *bind.TransactOpts, donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeDONs", donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveDONs(donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONs(&_CapabilitiesRegistry.TransactOpts, donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveDONs(donIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONs(&_CapabilitiesRegistry.TransactOpts, donIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeNodeOperators", nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveNodeOperators(nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveNodeOperators(nodeOperatorIds []uint32) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveNodes(opts *bind.TransactOpts, removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeNodes", removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveNodes(removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodes(&_CapabilitiesRegistry.TransactOpts, removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveNodes(removedNodeP2PIds [][32]byte) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveNodes(&_CapabilitiesRegistry.TransactOpts, removedNodeP2PIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) SetDONFamily(opts *bind.TransactOpts, donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "setDONFamily", donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) SetDONFamily(donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamily(&_CapabilitiesRegistry.TransactOpts, donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) SetDONFamily(donId uint32, donFamily string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamily(&_CapabilitiesRegistry.TransactOpts, donId, donFamily)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.TransferOwnership(&_CapabilitiesRegistry.TransactOpts, to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.TransferOwnership(&_CapabilitiesRegistry.TransactOpts, to)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDON(opts *bind.TransactOpts, donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDON", donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDON(donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDON(donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, nodes, capabilityConfigurations, isPublic, f, additionalParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateNodeOperators", nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateNodeOperators(nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateNodeOperators(nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodeOperators(&_CapabilitiesRegistry.TransactOpts, nodeOperatorIds, nodeOperators)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateNodes", nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateNodes(nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateNodes(&_CapabilitiesRegistry.TransactOpts, nodes)
}

type CapabilitiesRegistryCapabilityConfiguredIterator struct {
	Event *CapabilitiesRegistryCapabilityConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryCapabilityConfigured)
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
		it.Event = new(CapabilitiesRegistryCapabilityConfigured)
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

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryCapabilityConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryCapabilityConfigured struct {
	HashedCapabilityId [32]byte
	Raw                types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityConfigured(opts *bind.FilterOpts, hashedCapabilityId [][32]byte) (*CapabilitiesRegistryCapabilityConfiguredIterator, error) {

	var hashedCapabilityIdRule []interface{}
	for _, hashedCapabilityIdItem := range hashedCapabilityId {
		hashedCapabilityIdRule = append(hashedCapabilityIdRule, hashedCapabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityConfigured", hashedCapabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityConfiguredIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityConfigured", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, hashedCapabilityId [][32]byte) (event.Subscription, error) {

	var hashedCapabilityIdRule []interface{}
	for _, hashedCapabilityIdItem := range hashedCapabilityId {
		hashedCapabilityIdRule = append(hashedCapabilityIdRule, hashedCapabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityConfigured", hashedCapabilityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryCapabilityConfigured)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityConfigured", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseCapabilityConfigured(log types.Log) (*CapabilitiesRegistryCapabilityConfigured, error) {
	event := new(CapabilitiesRegistryCapabilityConfigured)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryCapabilityDeprecatedIterator struct {
	Event *CapabilitiesRegistryCapabilityDeprecated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryCapabilityDeprecated)
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
		it.Event = new(CapabilitiesRegistryCapabilityDeprecated)
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

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryCapabilityDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryCapabilityDeprecated struct {
	HashedCapabilityId [32]byte
	Raw                types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityDeprecated(opts *bind.FilterOpts, hashedCapabilityId [][32]byte) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error) {

	var hashedCapabilityIdRule []interface{}
	for _, hashedCapabilityIdItem := range hashedCapabilityId {
		hashedCapabilityIdRule = append(hashedCapabilityIdRule, hashedCapabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityDeprecated", hashedCapabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityDeprecatedIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityDeprecated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, hashedCapabilityId [][32]byte) (event.Subscription, error) {

	var hashedCapabilityIdRule []interface{}
	for _, hashedCapabilityIdItem := range hashedCapabilityId {
		hashedCapabilityIdRule = append(hashedCapabilityIdRule, hashedCapabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityDeprecated", hashedCapabilityIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryCapabilityDeprecated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityDeprecated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseCapabilityDeprecated(log types.Log) (*CapabilitiesRegistryCapabilityDeprecated, error) {
	event := new(CapabilitiesRegistryCapabilityDeprecated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "CapabilityDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryConfigSetIterator struct {
	Event *CapabilitiesRegistryConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryConfigSet)
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
		it.Event = new(CapabilitiesRegistryConfigSet)
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

func (it *CapabilitiesRegistryConfigSetIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryConfigSet struct {
	DonId       uint32
	ConfigCount uint32
	Raw         types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterConfigSet(opts *bind.FilterOpts, donId []uint32) (*CapabilitiesRegistryConfigSetIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "ConfigSet", donIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryConfigSetIterator{contract: _CapabilitiesRegistry.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryConfigSet, donId []uint32) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "ConfigSet", donIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryConfigSet)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseConfigSet(log types.Log) (*CapabilitiesRegistryConfigSet, error) {
	event := new(CapabilitiesRegistryConfigSet)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryDONFamilySetIterator struct {
	Event *CapabilitiesRegistryDONFamilySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONFamilySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONFamilySet)
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
		it.Event = new(CapabilitiesRegistryDONFamilySet)
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

func (it *CapabilitiesRegistryDONFamilySetIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONFamilySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONFamilySet struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONFamilySet(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONFamilySetIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONFamilySet", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONFamilySetIterator{contract: _CapabilitiesRegistry.contract, event: "DONFamilySet", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONFamilySet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONFamilySet, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONFamilySet", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONFamilySet)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONFamilySet", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONFamilySet(log types.Log) (*CapabilitiesRegistryDONFamilySet, error) {
	event := new(CapabilitiesRegistryDONFamilySet)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONFamilySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeAddedIterator struct {
	Event *CapabilitiesRegistryNodeAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeAdded)
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
		it.Event = new(CapabilitiesRegistryNodeAdded)
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

func (it *CapabilitiesRegistryNodeAddedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeAdded struct {
	P2pId          [32]byte
	NodeOperatorId uint32
	Signer         [32]byte
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeAdded(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeAddedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeAdded, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeAdded)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeAdded(log types.Log) (*CapabilitiesRegistryNodeAdded, error) {
	event := new(CapabilitiesRegistryNodeAdded)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorAddedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorAdded)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorAdded)
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

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorAdded struct {
	NodeOperatorId uint32
	Admin          common.Address
	Name           string
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorAddedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorAdded, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorAdded)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorAdded(log types.Log) (*CapabilitiesRegistryNodeOperatorAdded, error) {
	event := new(CapabilitiesRegistryNodeOperatorAdded)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorRemovedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorRemoved)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorRemoved)
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

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorRemoved struct {
	NodeOperatorId uint32
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorRemoved(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeOperatorRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorRemovedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorRemoved", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorRemoved, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorRemoved)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorRemoved", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorRemoved(log types.Log) (*CapabilitiesRegistryNodeOperatorRemoved, error) {
	event := new(CapabilitiesRegistryNodeOperatorRemoved)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeOperatorUpdatedIterator struct {
	Event *CapabilitiesRegistryNodeOperatorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeOperatorUpdated)
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
		it.Event = new(CapabilitiesRegistryNodeOperatorUpdated)
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

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeOperatorUpdated struct {
	NodeOperatorId uint32
	Admin          common.Address
	Name           string
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeOperatorUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorUpdatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeOperatorUpdated", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeOperatorUpdatedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeOperatorUpdated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorUpdated, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeOperatorUpdated", nodeOperatorIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeOperatorUpdated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeOperatorUpdated(log types.Log) (*CapabilitiesRegistryNodeOperatorUpdated, error) {
	event := new(CapabilitiesRegistryNodeOperatorUpdated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeRemovedIterator struct {
	Event *CapabilitiesRegistryNodeRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeRemoved)
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
		it.Event = new(CapabilitiesRegistryNodeRemoved)
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

func (it *CapabilitiesRegistryNodeRemovedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeRemoved struct {
	P2pId [32]byte
	Raw   types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeRemoved(opts *bind.FilterOpts) (*CapabilitiesRegistryNodeRemovedIterator, error) {

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeRemovedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeRemoved) (event.Subscription, error) {

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeRemoved)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeRemoved(log types.Log) (*CapabilitiesRegistryNodeRemoved, error) {
	event := new(CapabilitiesRegistryNodeRemoved)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryNodeUpdatedIterator struct {
	Event *CapabilitiesRegistryNodeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryNodeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryNodeUpdated)
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
		it.Event = new(CapabilitiesRegistryNodeUpdated)
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

func (it *CapabilitiesRegistryNodeUpdatedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryNodeUpdated struct {
	P2pId          [32]byte
	NodeOperatorId uint32
	Signer         [32]byte
	Raw            types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterNodeUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeUpdatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "NodeUpdated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryNodeUpdatedIterator{contract: _CapabilitiesRegistry.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeUpdated, nodeOperatorId []uint32) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "NodeUpdated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryNodeUpdated)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseNodeUpdated(log types.Log) (*CapabilitiesRegistryNodeUpdated, error) {
	event := new(CapabilitiesRegistryNodeUpdated)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryOwnershipTransferRequestedIterator struct {
	Event *CapabilitiesRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryOwnershipTransferRequested)
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
		it.Event = new(CapabilitiesRegistryOwnershipTransferRequested)
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

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryOwnershipTransferRequestedIterator{contract: _CapabilitiesRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryOwnershipTransferRequested)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*CapabilitiesRegistryOwnershipTransferRequested, error) {
	event := new(CapabilitiesRegistryOwnershipTransferRequested)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryOwnershipTransferredIterator struct {
	Event *CapabilitiesRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryOwnershipTransferred)
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
		it.Event = new(CapabilitiesRegistryOwnershipTransferred)
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

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryOwnershipTransferredIterator{contract: _CapabilitiesRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryOwnershipTransferred)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*CapabilitiesRegistryOwnershipTransferred, error) {
	event := new(CapabilitiesRegistryOwnershipTransferred)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CapabilitiesRegistry.abi.Events["CapabilityConfigured"].ID:
		return _CapabilitiesRegistry.ParseCapabilityConfigured(log)
	case _CapabilitiesRegistry.abi.Events["CapabilityDeprecated"].ID:
		return _CapabilitiesRegistry.ParseCapabilityDeprecated(log)
	case _CapabilitiesRegistry.abi.Events["ConfigSet"].ID:
		return _CapabilitiesRegistry.ParseConfigSet(log)
	case _CapabilitiesRegistry.abi.Events["DONFamilySet"].ID:
		return _CapabilitiesRegistry.ParseDONFamilySet(log)
	case _CapabilitiesRegistry.abi.Events["NodeAdded"].ID:
		return _CapabilitiesRegistry.ParseNodeAdded(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorAdded"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorAdded(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorRemoved"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorRemoved(log)
	case _CapabilitiesRegistry.abi.Events["NodeOperatorUpdated"].ID:
		return _CapabilitiesRegistry.ParseNodeOperatorUpdated(log)
	case _CapabilitiesRegistry.abi.Events["NodeRemoved"].ID:
		return _CapabilitiesRegistry.ParseNodeRemoved(log)
	case _CapabilitiesRegistry.abi.Events["NodeUpdated"].ID:
		return _CapabilitiesRegistry.ParseNodeUpdated(log)
	case _CapabilitiesRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _CapabilitiesRegistry.ParseOwnershipTransferRequested(log)
	case _CapabilitiesRegistry.abi.Events["OwnershipTransferred"].ID:
		return _CapabilitiesRegistry.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CapabilitiesRegistryCapabilityConfigured) Topic() common.Hash {
	return common.HexToHash("0x04f0a9bcf3f3a3b42a4d7ca081119755f82ebe43e0d30c8f7292c4fe0dc4a2ae")
}

func (CapabilitiesRegistryCapabilityDeprecated) Topic() common.Hash {
	return common.HexToHash("0xdcea1b78b6ddc31592a94607d537543fcaafda6cc52d6d5cc7bbfca1422baf21")
}

func (CapabilitiesRegistryConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651")
}

func (CapabilitiesRegistryDONFamilySet) Topic() common.Hash {
	return common.HexToHash("0x3143778fe5911daca9ced3c73dde73fdb95e7edae557614c40588f3f1a1b44b8")
}

func (CapabilitiesRegistryNodeAdded) Topic() common.Hash {
	return common.HexToHash("0x74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f05")
}

func (CapabilitiesRegistryNodeOperatorAdded) Topic() common.Hash {
	return common.HexToHash("0x78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e")
}

func (CapabilitiesRegistryNodeOperatorRemoved) Topic() common.Hash {
	return common.HexToHash("0xa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a")
}

func (CapabilitiesRegistryNodeOperatorUpdated) Topic() common.Hash {
	return common.HexToHash("0x86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a28")
}

func (CapabilitiesRegistryNodeRemoved) Topic() common.Hash {
	return common.HexToHash("0x5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb9753205")
}

func (CapabilitiesRegistryNodeUpdated) Topic() common.Hash {
	return common.HexToHash("0x4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b")
}

func (CapabilitiesRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CapabilitiesRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CapabilitiesRegistry *CapabilitiesRegistry) Address() common.Address {
	return _CapabilitiesRegistry.address
}

type CapabilitiesRegistryInterface interface {
	GetCapabilities(opts *bind.CallOpts) ([]CapabilitiesRegistryCapabilityInfo, error)

	GetCapability(opts *bind.CallOpts, hashedId [32]byte) (CapabilitiesRegistryCapabilityInfo, error)

	GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId [32]byte) ([]byte, []byte, error)

	GetDON(opts *bind.CallOpts, donId uint32) (CapabilitiesRegistryDONInfo, error)

	GetDONs(opts *bind.CallOpts) ([]CapabilitiesRegistryDONInfo, error)

	GetDONsInFamily(opts *bind.CallOpts, donFamily string) ([]*big.Int, error)

	GetHashedCapabilityId(opts *bind.CallOpts, labelledName string, version string) ([32]byte, error)

	GetNextDONId(opts *bind.CallOpts) (uint32, error)

	GetNode(opts *bind.CallOpts, p2pId [32]byte) (INodeInfoProviderNodeInfo, error)

	GetNodeOperator(opts *bind.CallOpts, nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error)

	GetNodeOperators(opts *bind.CallOpts) ([]CapabilitiesRegistryNodeOperator, error)

	GetNodes(opts *bind.CallOpts) ([]INodeInfoProviderNodeInfo, error)

	GetNodesByP2PIds(opts *bind.CallOpts, p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error)

	IsCapabilityDeprecated(opts *bind.CallOpts, hashedCapabilityId [32]byte) (bool, error)

	IsDONNameTaken(opts *bind.CallOpts, donName string) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddCapabilities(opts *bind.TransactOpts, capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error)

	AddDON(opts *bind.TransactOpts, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, acceptsWorkflows bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error)

	AddNodeOperators(opts *bind.TransactOpts, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	AddNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	DeprecateCapabilities(opts *bind.TransactOpts, hashedCapabilityIds [][32]byte) (*types.Transaction, error)

	RemoveDONs(opts *bind.TransactOpts, donIds []uint32) (*types.Transaction, error)

	RemoveNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32) (*types.Transaction, error)

	RemoveNodes(opts *bind.TransactOpts, removedNodeP2PIds [][32]byte) (*types.Transaction, error)

	SetDONFamily(opts *bind.TransactOpts, donId uint32, donFamily string) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDON(opts *bind.TransactOpts, donId uint32, nodes [][32]byte, capabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration, isPublic bool, f uint8, additionalParams CapabilitiesRegistryAdditionalDONParams) (*types.Transaction, error)

	UpdateNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	UpdateNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	FilterCapabilityConfigured(opts *bind.FilterOpts, hashedCapabilityId [][32]byte) (*CapabilitiesRegistryCapabilityConfiguredIterator, error)

	WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, hashedCapabilityId [][32]byte) (event.Subscription, error)

	ParseCapabilityConfigured(log types.Log) (*CapabilitiesRegistryCapabilityConfigured, error)

	FilterCapabilityDeprecated(opts *bind.FilterOpts, hashedCapabilityId [][32]byte) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error)

	WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, hashedCapabilityId [][32]byte) (event.Subscription, error)

	ParseCapabilityDeprecated(log types.Log) (*CapabilitiesRegistryCapabilityDeprecated, error)

	FilterConfigSet(opts *bind.FilterOpts, donId []uint32) (*CapabilitiesRegistryConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryConfigSet, donId []uint32) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CapabilitiesRegistryConfigSet, error)

	FilterDONFamilySet(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONFamilySetIterator, error)

	WatchDONFamilySet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONFamilySet, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONFamilySet(log types.Log) (*CapabilitiesRegistryDONFamilySet, error)

	FilterNodeAdded(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeAddedIterator, error)

	WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeAdded, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeAdded(log types.Log) (*CapabilitiesRegistryNodeAdded, error)

	FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorAddedIterator, error)

	WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorAdded, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error)

	ParseNodeOperatorAdded(log types.Log) (*CapabilitiesRegistryNodeOperatorAdded, error)

	FilterNodeOperatorRemoved(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeOperatorRemovedIterator, error)

	WatchNodeOperatorRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorRemoved, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeOperatorRemoved(log types.Log) (*CapabilitiesRegistryNodeOperatorRemoved, error)

	FilterNodeOperatorUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32, admin []common.Address) (*CapabilitiesRegistryNodeOperatorUpdatedIterator, error)

	WatchNodeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeOperatorUpdated, nodeOperatorId []uint32, admin []common.Address) (event.Subscription, error)

	ParseNodeOperatorUpdated(log types.Log) (*CapabilitiesRegistryNodeOperatorUpdated, error)

	FilterNodeRemoved(opts *bind.FilterOpts) (*CapabilitiesRegistryNodeRemovedIterator, error)

	WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeRemoved) (event.Subscription, error)

	ParseNodeRemoved(log types.Log) (*CapabilitiesRegistryNodeRemoved, error)

	FilterNodeUpdated(opts *bind.FilterOpts, nodeOperatorId []uint32) (*CapabilitiesRegistryNodeUpdatedIterator, error)

	WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryNodeUpdated, nodeOperatorId []uint32) (event.Subscription, error)

	ParseNodeUpdated(log types.Log) (*CapabilitiesRegistryNodeUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CapabilitiesRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CapabilitiesRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CapabilitiesRegistryOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
