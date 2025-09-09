// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package capabilities_registry_wrapper_v2_dev

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

type CapabilitiesRegistryCapability struct {
	CapabilityId          string
	ConfigurationContract common.Address
	Metadata              []byte
}

type CapabilitiesRegistryCapabilityConfiguration struct {
	CapabilityId string
	Config       []byte
}

type CapabilitiesRegistryCapabilityInfo struct {
	CapabilityId          string
	ConfigurationContract common.Address
	IsDeprecated          bool
	Metadata              []byte
}

type CapabilitiesRegistryConstructorParams struct {
	CanAddOneNodeDONs bool
}

type CapabilitiesRegistryDONInfo struct {
	Id                       uint32
	ConfigCount              uint32
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
	NodeP2PIds               [][32]byte
	DonFamilies              []string
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
}

type CapabilitiesRegistryNewDONParams struct {
	Name                     string
	DonFamilies              []string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
	Nodes                    [][32]byte
	F                        uint8
	IsPublic                 bool
	AcceptsWorkflows         bool
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
	CsaKey              [32]byte
	CapabilityIds       []string
}

type CapabilitiesRegistryUpdateDONParams struct {
	Name                     string
	Config                   []byte
	CapabilityConfigurations []CapabilitiesRegistryCapabilityConfiguration
	Nodes                    [][32]byte
	F                        uint8
	IsPublic                 bool
}

type INodeInfoProviderNodeInfo struct {
	NodeOperatorId      uint32
	ConfigCount         uint32
	WorkflowDONId       uint32
	Signer              [32]byte
	P2pId               [32]byte
	EncryptionPublicKey [32]byte
	CsaKey              [32]byte
	CapabilityIds       []string
	CapabilitiesDONIds  []*big.Int
}

var CapabilitiesRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.ConstructorParams\",\"components\":[{\"name\":\"canAddOneNodeDONs\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCapabilities\",\"inputs\":[{\"name\":\"capabilities\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.Capability[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addDONs\",\"inputs\":[{\"name\":\"newDONs\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NewDONParams[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodeOperators\",\"inputs\":[{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deprecateCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCapabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapability\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.CapabilityInfo\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"configurationContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeprecated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCapabilityConfigs\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONFamilies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.DONInfo[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDONsInFamily\",\"inputs\":[{\"name\":\"donFamily\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHistoricalDONInfo\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.DONInfo\",\"components\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptsWorkflows\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"donFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextDONId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nodeInfo\",\"type\":\"tuple\",\"internalType\":\"structINodeInfoProvider.NodeInfo\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperator\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodesByP2PIds\",\"inputs\":[{\"name\":\"p2pIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structINodeInfoProvider.NodeInfo[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"workflowDONId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"capabilitiesDONIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDONNameTaken\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDONs\",\"inputs\":[{\"name\":\"donIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDONsByName\",\"inputs\":[{\"name\":\"donNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNodes\",\"inputs\":[{\"name\":\"removedNodeP2PIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDONFamilies\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addToFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"removeFromFamilies\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDONByName\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"updateDONParams\",\"type\":\"tuple\",\"internalType\":\"structCapabilitiesRegistry.UpdateDONParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"capabilityConfigurations\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.CapabilityConfiguration[]\",\"components\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"nodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isPublic\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodeOperators\",\"inputs\":[{\"name\":\"nodeOperatorIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nodeOperators\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeOperator[]\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateNodes\",\"inputs\":[{\"name\":\"nodes\",\"type\":\"tuple[]\",\"internalType\":\"structCapabilitiesRegistry.NodeParams[]\",\"components\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CapabilityConfigured\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapabilityDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"configCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONAddedToFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DONRemovedFromFamily\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"donFamily\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeAdded\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorAdded\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeOperatorUpdated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRemoved\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"signer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CapabilityAlreadyExists\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityDoesNotExist\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityIsDeprecated\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"CapabilityRequiredByDON\",\"inputs\":[{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONConfigDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"requestedConfigCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONDoesNotExist\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONNameAlreadyTaken\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DONNameCannotBeEmpty\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"DONWithNameDoesNotExist\",\"inputs\":[{\"name\":\"donName\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONCapability\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateDONNode\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCapabilityConfigurationContractInterface\",\"inputs\":[{\"name\":\"proposedConfigurationContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidFaultTolerance\",\"inputs\":[{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCSAKey\",\"inputs\":[{\"name\":\"csaKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeCapabilities\",\"inputs\":[{\"name\":\"capabilityIds\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeEncryptionPublicKey\",\"inputs\":[{\"name\":\"encryptionPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeOperatorAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNodeP2PId\",\"inputs\":[{\"name\":\"p2pId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"lengthOne\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lengthTwo\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeAlreadyExists\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotExist\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodeDoesNotSupportCapability\",\"inputs\":[{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"capabilityId\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorAlreadyExists\",\"inputs\":[{\"name\":\"existingNodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfCapabilitiesDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NodePartOfWorkflowDON\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nodeP2PId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60a0604052346100e557604051601f6156e338819003918201601f19168301916001600160401b038311848410176100ea578084926020946040528339810103126100e55760405190600090602083016001600160401b038111848210176100d1576040525180151581036100cd57825233156100be5750600180546001600160a01b03191633179055601680546001600160401b0319166401000000011790555115156080526040516155e290816101018239608051816147760152f35b639b15e16f60e01b8152600490fd5b5080fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80628375c61461024657806305a519661461024157806307e1959c1461023c578063181f5a77146102375780631d05394c14610232578063214502431461022d57806322bdbcbc146102285780632353740514610223578063275459f21461021e5780632af97674146102195780632c01a1e814610214578063398f37731461020f57806350c946fe1461020a57806353a25dd714610205578063543f40251461020057806359003602146101fb57806359110666146101f657806366acaa33146101f157806379ba5097146101ec57806386fa4246146101e757806388ea09ee146101e257806388eafafb146101dd5780638da5cb5b146101d857806394bbb012146101d357806396ef4fc9146101ce578063a04ab55e146101c9578063a9044eb5146101c4578063b8521761146101bf578063bfa8eef5146101ba578063c9315179146101b5578063cd71fd09146101b0578063ddbe4f82146101ab578063e29581aa146101a6578063f2fde38b146101a15763fcdc8efe1461019c57600080fd5b612c1d565b612b7a565b612ab4565b612a02565b612952565b6128d2565b612816565b6127ad565b612463565b6123f1565b612363565b6122e4565b6122bd565b6121ed565b612118565b611e13565b611d7a565b611ca0565b611bce565b611b95565b611b1c565b611927565b6118b8565b6116b2565b61150b565b611360565b611217565b611195565b61112c565b610fe4565b610e0e565b610daf565b610676565b6105dd565b6102b0565b9181601f8401121561027c5782359167ffffffffffffffff831161027c576020808501948460051b01011161027c57565b600080fd5b602060031982011261027c576004359067ffffffffffffffff821161027c576102ac9160040161024b565b9091565b3461027c576102be36610281565b906102c7614024565b60005b8281106102d357005b6102e66102e1828585612c59565b612c80565b6102f581516020815191012090565b61030561030182615022565b1590565b6103e4576103278251610322836000526015602052604060002090565b612db6565b6020820180516001600160a01b031680610393575b5050816103626001949361035d610368946000526003602052604060002090565b612e84565b51612fa0565b7fe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb600080a2016102ca565b61030161039f91614062565b6103a9578061033c565b517fabb5e3fd000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b6000fd5b61041b82516040519182917f8f51ece800000000000000000000000000000000000000000000000000000000835260048301610d9e565b0390fd5b60005b8381106104325750506000910152565b8181015183820152602001610422565b9060209161045b8151809281855285808601910161041f565b601f01601f1916010190565b9080602083519182815201916020808360051b8301019401926000915b83831061049357505050505090565b90919293946020806104b1600193601f198682030187528951610442565b97019301930191939290610484565b906020808351928381520192019060005b8181106104de5750505090565b82518452602093840193909201916001016104d1565b805163ffffffff16825261057a9160208281015163ffffffff169082015260408281015163ffffffff1690820152606082015160608201526080820151608082015260a082015160a082015260c082015160c082015261010061056860e084015161012060e0850152610120840190610467565b920151906101008184039101526104c0565b90565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106105b057505050505090565b90919293946020806105ce600193603f1986820301875289516104f4565b970193019301919392906105a1565b3461027c576105eb36610281565b6105f481613034565b9060005b818110610611576040518061060d858261057d565b0390f35b61062561061f828487613084565b356137d2565b61062f8285613094565b5261063a8184613094565b5060806106478285613094565b51015115610657576001016105f8565b6106619184613084565b3563d82f6adb60e01b60005260045260246000fd5b3461027c5761068436610281565b6106a56106996001600160a01b036001541690565b6001600160a01b031690565b33149060009115905b8083106106b757005b6106ca6106c58483876130a8565b613149565b60408101906106e482516000526010602052604060002090565b61071361070e6106f8835463ffffffff1690565b63ffffffff16600052600e602052604060002090565b61325c565b95600182018054978815610c8057879081610c63575b50610c35576020840197885115610c0b578851808203610bcf575b5050506060830180518015610ba25750608084019485518015610b75575060a0850151998a5115610b5a579661078a610785865463ffffffff9060201c1690565b6132b0565b855467ffffffff000000001916602082901b67ffffffff000000001617865598600586019860005b8d5181101561084b576107ed6103018f6107cf846107da92613094565b516020815191012090565b6000526005602052604060002054151590565b61082e57806108278f6108218f8f9561081b6107cf926001989063ffffffff16600052602052604060002090565b93613094565b90615131565b50016107b2565b61041b8e604051918291636db4786160e11b835260048301612303565b508654919c509a999897959694919391929060401c63ffffffff1663ffffffff8116610a81575b506108836006889c9b959c01613362565b9360009b5b855163ffffffff8e16908110156109e9576108a7909c9e919c87613094565b5163ffffffff169d8e6108ca8163ffffffff166000526011602052604060002090565b600101906108e89063ffffffff166000526011602052604060002090565b5460201c63ffffffff1661090b919063ffffffff16600052602052604060002090565b60030161091790613362565b9c60008e5b518110156109d1576109698f8f908f8461094c61030194610952939063ffffffff16600052602052604060002090565b92613094565b519060019160005201602052604060002054151590565b610976576001018e61091c565b90508f935061099c925061098b91508d613094565b516000526015602052604060002090565b61041b6040519283927f16c2b7c4000000000000000000000000000000000000000000000000000000008452600484016133ad565b50929c509c6109e1919e506132b0565b9b9a90610888565b509c9b50919690935060019850610a7592975060047f4b5b465e22eea0c3d40c30e936643245b80d19b2dcf75788c0699fe8d8db645b96610a3363ffffffff975163ffffffff1690565b835463ffffffff191663ffffffff8216178455955198896002850155516003840155519101555160405193849316958360209093929193604081019481520152565b0390a2019190926106ae565b9b610af86003610af28f9d9e9d610ade610ad16001610ab99d9e9c9d849c999a9b9c63ffffffff166000526011602052604060002090565b019263ffffffff166000526011602052604060002090565b5460201c63ffffffff1690565b63ffffffff16600052602052604060002090565b01613362565b9a60005b8c51811015610b4757610b2c6103018e6109528f8f869161094c919063ffffffff16600052602052604060002090565b610b3857600101610afc565b61099c8f9161098b908f613094565b509b9a509b509291909594939538610872565b604051636db4786160e11b81528061041b8d60048301612303565b7fd79735610000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f37d897650000000000000000000000000000000000000000000000000000000060005260045260246000fd5b610be6906000526009602052604060002054151590565b610c0b57610bf79189519055615228565b50610c02875161505d565b50388080610744565b7f837731460000000000000000000000000000000000000000000000000000000060005260046000fd5b7f9473075d000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b51610c7791506001600160a01b0316610699565b33141538610729565b855163d82f6adb60e01b60005260045260246000fd5b600091031261027c57565b634e487b7160e01b600052604160045260246000fd5b60c0810190811067ffffffffffffffff821117610cd357604052565b610ca1565b6040810190811067ffffffffffffffff821117610cd357604052565b6080810190811067ffffffffffffffff821117610cd357604052565b90601f8019910116810190811067ffffffffffffffff821117610cd357604052565b60405190610d41604083610d10565b565b60405190610d4160e083610d10565b60405190610d4161012083610d10565b60405190610d4161010083610d10565b60405190610d4161014083610d10565b67ffffffffffffffff8111610cd357601f01601f191660200190565b90602061057a928181520190610442565b3461027c57600036600319011261027c5761060d6040805190610dd28183610d10565b601e82527f4361706162696c6974696573526567697374727920322e302e302d6465760000602083015251918291602083526020830190610442565b3461027c57610e1c36610281565b90610e25614024565b60005b828110610e3157005b80610e51610e426001938686613084565b35610e4c816110f8565b61420c565b01610e28565b9080602083519182815201916020808360051b8301019401926000915b838310610e8357505050505090565b9091929394602080610ec1600193601f1986820301875289519083610eb18351604084526040840190610442565b9201519084818403910152610442565b97019301930191939290610e74565b805163ffffffff16825261057a9160208281015163ffffffff169082015260408281015160ff1690820152606082810151151590820152608082810151151590820152610120610f72610f5e610f4c610f3a60a087015161014060a08801526101408701906104c0565b60c087015186820360c0880152610467565b60e086015185820360e0870152610442565b610100850151848203610100860152610442565b92015190610120818403910152610e57565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310610fb757505050505090565b9091929394602080610fd5600193603f198682030187528951610ed0565b97019301930191939290610fa8565b3461027c57600036600319011261027c5760165460201c63ffffffff1661101e611019611010836133db565b63ffffffff1690565b613455565b60009163ffffffff811660015b8163ffffffff8216106110655761060d8486611049611010876133db565b810361105d575b5060405191829182610f84565b815282611050565b61108f6110106110858363ffffffff166000526011602052604060002090565b5463ffffffff1690565b6110a2575b60010163ffffffff1661102b565b9360016110ef63ffffffff926110d46110ce610ad18a63ffffffff166000526011602052604060002090565b89614437565b6110de8289613094565b526110e98188613094565b506134a5565b95915050611094565b63ffffffff81160361027c57565b906040602061057a936001600160a01b0381511684520151918160208201520190610442565b3461027c57602036600319011261027c5763ffffffff60043561114e816110f8565b6111566134b4565b5016600052600e60205261060d611170604060002061325c565b604051918291602083526020830190611106565b90602061057a928181520190610ed0565b3461027c57602036600319011261027c576004356111b2816110f8565b6111ba6133f4565b5063ffffffff81169081600052601160205263ffffffff60406000205460201c169182156112035761060d6111ef8484614437565b604051918291602083526020830190610ed0565b632b62be9b60e01b60005260045260246000fd5b3461027c5761122536610281565b9061122e614024565b60005b63ffffffff811690838210156112e75761124f6112e2928585613084565b35611259816110f8565b6112b86112b363ffffffff83169283600052600e602052611285611280604060002061325c565b61463b565b600052600f602052604060002063ffffffff19815416905563ffffffff16600052600e602052604060002090565b6134ce565b7fa59268ca81d40429e65ccea5385b59cf2d3fc6519371dee92f8eb1dae5107a7a600080a26132b0565b611231565b005b9291926112f582610d82565b916113036040519384610d10565b82948184528183011161027c578281602093846000960137010152565b9080601f8301121561027c5781602061057a933591016112e9565b909161135261057a93604084526040840190610442565b916020818403910152610442565b3461027c57604036600319011261027c5760043561137d816110f8565b60243567ffffffffffffffff811161027c5761139d903690600401611320565b9061141461140f6113c1610ad18463ffffffff166000526011602052604060002090565b9360066114086113d5836020815191012090565b9660016113f28863ffffffff166000526011602052604060002090565b019063ffffffff16600052602052604060002090565b019061355c565b613241565b906060926001600160a01b036114476001611439846000526003602052604060002090565b01546001600160a01b031690565b1661145e575b505061060d6040519283928361133b565b6114c79293509061148661069961069960016114396000966000526003602052604060002090565b60405180809581947f8318ed5d0000000000000000000000000000000000000000000000000000000083526004830191909163ffffffff6020820193169052565b03915afa908115611506576000916114e3575b5090388061144d565b61150091503d806000833e6114f88183610d10565b810190613582565b386114da565b6135e1565b3461027c5761151936610281565b9061152f6106996001600160a01b036001541690565b3314159160005b81811061153f57005b61154a818385613084565b3590611560826000526010602052604060002090565b6001810154801561169c576006820180546116575750815463ffffffff604082901c1680611638575087908161160a575b50610c35576001936115d460027f5254e609a97bab37b7cc79fe128f85c097bd6015c6e1624ae0ba392eb9753205946115cc61160195615228565b5001546152c9565b506115f16115ec826000526010602052604060002090565b6135ed565b6040519081529081906020820190565b0390a101611536565b61162f91506116226106f86106999263ffffffff1690565b546001600160a01b031690565b33141538611591565b6360b9df7360e01b60005263ffffffff16600452602485905260446000fd5b846116676110106103e093615484565b7f60a6d8980000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602452604490565b63d82f6adb60e01b600052600484905260246000fd5b3461027c576116c036610281565b6116c8614024565b6000915b8183106116d557005b6116e86116e3848484613645565b613667565b926116fd61069985516001600160a01b031690565b1561188e5760165463ffffffff1684516001600160a01b031661175e6020870191825161173a61172b610d32565b6001600160a01b039093168352565b60208201526117598463ffffffff16600052600e602052604060002090565b6136b2565b6117678661463b565b9561178261101061108589600052600f602052604060002090565b611842576001600160a01b0361181c7f78e94ca80be2c30abc061b99e7eb8583b1254781734b1e3ce339abb57da2fe8e926117e4866117cf60019a9b9c600052600f602052604060002090565b9063ffffffff1663ffffffff19825416179055565b61180f6117f961078560165463ffffffff1690565b63ffffffff1663ffffffff196016541617601655565b516001600160a01b031690565b92519261183763ffffffff6040519384931696169482610d9e565b0390a30191906116cc565b6103e061185c61108589600052600f602052604060002090565b7f8c0346380000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b7feeacd9390000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c57602036600319011261027c5761060d6118d76004356137d2565b6040519182916020835260208301906104f4565b9181601f8401121561027c5782359167ffffffffffffffff831161027c576020838186019501011161027c57565b908160c091031261027c5790565b3461027c57604036600319011261027c5760043567ffffffffffffffff811161027c576119589036906004016118eb565b9060243567ffffffffffffffff811161027c57611979903690600401611919565b91611982614024565b61198f6110858284613989565b9163ffffffff831615611abc5750506119b88163ffffffff166000526011602052604060002090565b6119c560608401846139ca565b6119d260408601866139ca565b8454909691949060201c63ffffffff166119eb906132b0565b815467ffffffff000000001916602082901b67ffffffff000000001617825591611a1760a08201613a0a565b915460401c60ff165b611a2c60808301613a1f565b611a368380613a29565b92909360208101611a4691613a29565b959096611a51610d43565b63ffffffff909c168c5263ffffffff1660208c0152151560408b0152151560608a015260ff1660808901523690611a87926112e9565b60a08701523690611a97926112e9565b60c08501523690611aa792613a5c565b923690611ab392613aa8565b6112e79261471a565b61041b6040519283927f4071db54000000000000000000000000000000000000000000000000000000008452600484016139a2565b602060031982011261027c576004359067ffffffffffffffff821161027c576102ac916004016118eb565b3461027c57611b34611b2d36611af1565b36916112e9565b602081519101206000526013602052611b506040600020613362565b60405180916020820160208352815180915260206040840192019060005b818110611b7c575050500390f35b8251845285945060209384019390920191600101611b6e565b3461027c576020611bc4611bab611b2d36611af1565b8281519101206000526007602052604060002054151590565b6040519015158152f35b3461027c57611bdc36611af1565b611be46133f4565b5063ffffffff6040518284823760208184810160028152030190205416918215611abc5761060d611c3484806000526011602052611c2e604060002063ffffffff905460201c1690565b90614437565b60405191829182611184565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611c7357505050505090565b9091929394602080611c91600193603f198682030187528951611106565b97019301930191939290611c64565b3461027c57600036600319011261027c5760165463ffffffff16611cce611cc9611010836133db565b613b65565b60009163ffffffff811660015b8163ffffffff821610611d155761060d8486611cf9611010876133db565b8103611d0d575b5060405191829182611c40565b815282611d00565b611d356106996116228363ffffffff16600052600e602052604060002090565b611d48575b60010163ffffffff16611cdb565b936001611d7163ffffffff926110d461070e8963ffffffff16600052600e602052604060002090565b95915050611d3a565b3461027c57600036600319011261027c576000546001600160a01b0381163303611de9576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c57604036600319011261027c5760043567ffffffffffffffff811161027c57611e4490369060040161024b565b60243567ffffffffffffffff811161027c57611e6490369060040161024b565b8083949294036120e657828491611e836001600160a01b036001541690565b906000945b838610611e9157005b611ea4611e9f878684613084565b6133d1565b95611ebf8763ffffffff16600052600e602052604060002090565b96611ed461069989546001600160a01b031690565b9788156120cb57611ee96116e384878b613645565b98611efe6106998b516001600160a01b031690565b1561188e57331415806120b8575b610c3557611f198961463b565b98611f346110106110858c600052600f602052604060002090565b61209e57611f76836117cf60019798999a9b9c611f59611f69611f596112808961325c565b600052600f602052604060002090565b805463ffffffff19169055565b81546001600160a01b03166001600160a01b03611f9d61069984516001600160a01b031690565b911614801590612050575b611fbb575b505050019493929190611e88565b6001600160a01b0361202a82612018611ffc7f86f41145bde5dd7f523305452e4aad3685508c181432ec733d5f345009358a2895516001600160a01b031690565b86906001600160a01b03166001600160a01b0319825416179055565b61180f60208201958987519101612db6565b92519261204563ffffffff6040519384931696169482610d9e565b0390a3878080611fad565b5060405160208101906120778161206989870185613bb5565b03601f198101835282610d10565b519020602082015160405161209481612069602082019485610d9e565b5190201415611fa8565b6103e061185c6110858c600052600f602052604060002090565b506001600160a01b038616331415611f0c565b6356ecd70f60e11b60005263ffffffff821660045260246000fd5b7fab8b67c600000000000000000000000000000000000000000000000000000000600052600483905260245260446000fd5b3461027c5761212636610281565b9061212f614024565b60005b82811061213b57005b612149611b2d828585613bc6565b8051602082012061216a610301826000526005602052604060002054151590565b6121d25761030161217a91615092565b6121b5579061218a600192612fa0565b7fb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83600080a201612132565b61041b906040519182916388c8a73760e01b835260048301610d9e565b6040516327fcf24560e11b81528061041b8460048301610d9e565b3461027c57604036600319011261027c5760043561220a816110f8565b60243567ffffffffffffffff811161027c5761222a903690600401611919565b90612233614024565b61224d8163ffffffff166000526011602052604060002090565b549163ffffffff602084901c169081156122a25761226e60608201826139ca565b909161227d60408201826139ca565b969094612289906132b0565b9161229660a08201613a0a565b9160401c60ff16611a20565b632b62be9b60e01b60005263ffffffff831660045260246000fd5b3461027c57600036600319011261027c5760206001600160a01b0360015416604051908152f35b3461027c576112e76122f536610281565b906122fe614024565b613d4d565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061233657505050505090565b9091929394602080612354600193603f198682030187528951610442565b97019301930191939290612327565b3461027c57600036600319011261027c5761237c6132cc565b6123868151613788565b9060005b81518110156123e357806123a060019284613094565b5160005260146020526123c06123c76040600020604051928380926131be565b0382610d10565b6123d18286613094565b526123dc8185613094565b500161238a565b6040518061060d8582612303565b3461027c57606036600319011261027c5760043561240e816110f8565b60243567ffffffffffffffff811161027c5761242e90369060040161024b565b916044359267ffffffffffffffff841161027c576124536112e794369060040161024b565b93909261245e614024565b613e8a565b3461027c5761247136610281565b906124876106996001546001600160a01b031690565b3314600090155b83821061249757005b6124a56106c58386866130a8565b916124ba61070e6106f8855163ffffffff1690565b6124ce61069982516001600160a01b031690565b1561278357829081612766575b50610c355760408301926124fa84516000526010602052604060002090565b9360018501908154612738578051801561270b57506020830191825180159081156126ed575b50610c0b57606084019687518015610ba257506080850180518015610b75575060a0860151998a5115610b5a5798999a8b9861258c61256a610785865463ffffffff9060201c1690565b855467ffffffff00000000191660209190911b67ffffffff0000000016178555565b835460201c63ffffffff169a6000600586019b5b51811015612629576125b68f826107cf91613094565b6125d0610301826000526005602052604060002054151590565b61260b578f949392916126008f928f6001946125fb919063ffffffff16600052602052604060002090565b615131565b5001909192936125a0565b5061041b8f604051918291636db4786160e11b835260048301612303565b509a509a63ffffffff95919c50600199507f74becb12a5e8fd0e98077d02dfba8f647c9670c9df177e42c2418cf17a636f0596946126e2946126c5946126ad9351600383015551600482015561269c6126868b5163ffffffff1690565b825463ffffffff191663ffffffff909116178255565b60028451910155855180915561505d565b506126b881516150c7565b5051955163ffffffff1690565b915160405193849316958360209093929193604081019481520152565b0390a201909161248e565b61270591506000526009602052604060002054151590565b38612520565b7f64e2ee920000000000000000000000000000000000000000000000000000000060005260045260246000fd5b517f546184830000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5161277a91506001600160a01b0316610699565b331415386124db565b6103e0612794855163ffffffff1690565b6356ecd70f60e11b60005263ffffffff16600452602490565b3461027c576127bb36610281565b6127c3614024565b60005b8181106127cf57005b8063ffffffff60206127e46001948688613bc6565b919082604051938492833781016002815203019020541680156128105761280a9061420c565b016127c6565b5061280a565b3461027c57604036600319011261027c57600435612833816110f8565b60243590612840826110f8565b6128486133f4565b5063ffffffff811680600052601160205263ffffffff60406000205460201c1680156128bd5763ffffffff84169181831161288a5761060d611c348686614437565b7ff3c16e2c0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50632b62be9b60e01b60005260045260246000fd5b3461027c57602063ffffffff816128e836611af1565b91908260405193849283378101600281520301902054161515604051908152f35b61057a9160606129228351608084526080840190610442565b926001600160a01b0360208201511660208401526040810151151560408401520151906060818403910152610442565b3461027c57602036600319011261027c5760043567ffffffffffffffff811161027c5761298e61298961060d923690600401611320565b613fa2565b604051918291602083526020830190612909565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106129d557505050505090565b90919293946020806129f3600193603f198682030187528951612909565b970193019301919392906129c6565b3461027c57600036600319011261027c57612a1b613317565b805190612a2782612fc0565b91612a356040519384610d10565b808352612a44601f1991612fc0565b0160005b818110612a9d57505060005b8151811015612a8f5780612a7361298961140f61098b60019587613094565b612a7d8286613094565b52612a888185613094565b5001612a54565b6040518061060d85826129a2565b602090612aa8613f7c565b82828701015201612a48565b3461027c57600036600319011261027c57604051600a548082528160208101600a60005260206000209260005b818110612b4d575050612af692500382610d10565b612b008151613034565b9060005b8151811015612b3f5780612b23612b1d60019385613094565b516137d2565b612b2d8286613094565b52612b388185613094565b5001612b04565b6040518061060d858261057d565b8454835260019485019486945060209093019201612ae1565b35906001600160a01b038216820361027c57565b3461027c57602036600319011261027c576004356001600160a01b03811680910361027c57612ba7614024565b338114612bf357806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461027c57600036600319011261027c57602063ffffffff601654821c16604051908152f35b634e487b7160e01b600052603260045260246000fd5b9190811015612c7b5760051b81013590605e198136030182121561027c570190565b612c43565b60608136031261027c57604051906060820182811067ffffffffffffffff821117610cd357604052803567ffffffffffffffff811161027c57612cc69036908301611320565b8252612cd460208201612b66565b602083015260408101359067ffffffffffffffff821161027c57612cfa91369101611320565b604082015290565b90600182811c92168015612d32575b6020831014612d1c57565b634e487b7160e01b600052602260045260246000fd5b91607f1691612d11565b91612d569183549060031b91821b91600019901b19161790565b9055565b818110612d65575050565b60008155600101612d5a565b9190601f8111612d8057505050565b610d41926000526020600020906020601f840160051c83019310612dac575b601f0160051c0190612d5a565b9091508190612d9f565b919091825167ffffffffffffffff8111610cd357612dde81612dd88454612d02565b84612d71565b6020601f8211600114612e1b578190612d56939495600092612e10575b50508160011b916000199060031b1c19161790565b015190503880612dfb565b601f19821690612e3084600052602060002090565b9160005b818110612e6c57509583600195969710612e53575b505050811b019055565b015160001960f88460031b161c19169055388080612e49565b9192602060018192868b015181550194019201612e34565b919091825192835167ffffffffffffffff8111610cd357612eaf81612ea98554612d02565b85612d71565b6020601f8211600114612f2a5791612ee882604093600295610d419899600092612e105750508160011b916000199060031b1c19161790565b84555b612f21612f0260208301516001600160a01b031690565b60018601906001600160a01b03166001600160a01b0319825416179055565b01519101612db6565b601f19821695612f3f85600052602060002090565b9660005b818110612f88575092610d419697600295936001938360409710612f6f575b505050811b018455612eeb565b015160001960f88460031b161c19169055388080612f62565b83830151895560019098019760209384019301612f43565b612fb89060206040519282848094519384920161041f565b810103902090565b67ffffffffffffffff8111610cd35760051b60200190565b60405190610120820182811067ffffffffffffffff821117610cd35760405260606101008360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e08201520152565b9061303e82612fc0565b61304b6040519182610d10565b828152809261305c601f1991612fc0565b019060005b82811061306d57505050565b602090613078612fd8565b82828501015201613061565b9190811015612c7b5760051b0190565b8051821015612c7b5760209160051b010190565b9190811015612c7b5760051b8101359060be198136030182121561027c570190565b9080601f8301121561027c5781356130e181612fc0565b926130ef6040519485610d10565b81845260208085019260051b8201019183831161027c5760208201905b83821061311b57505050505090565b813567ffffffffffffffff811161027c5760209161313e87848094880101611320565b81520191019061310c565b60c08136031261027c576040519061316082610cb7565b803561316b816110f8565b82526020810135602083015260408101356040830152606081013560608301526080810135608083015260a08101359067ffffffffffffffff821161027c576131b6913691016130ca565b60a082015290565b600092918154916131ce83612d02565b808352926001811690811561322457506001146131ea57505050565b60009081526020812093945091925b83831061320a575060209250010190565b6001816020929493945483858701015201910191906131f9565b915050602093945060ff929192191683830152151560051b010190565b90610d4161325592604051938480926131be565b0383610d10565b906001602060405161326d81610cd8565b61329681956001600160a01b03815416835261328f60405180968193016131be565b0384610d10565b0152565b634e487b7160e01b600052601160045260246000fd5b63ffffffff1663ffffffff81146132c75760010190565b61329a565b60405190600c548083528260208101600c60005260206000209260005b8181106132fe575050610d4192500383610d10565b84548352600194850194879450602090930192016132e9565b604051906004548083528260208101600460005260206000209260005b818110613349575050610d4192500383610d10565b8454835260019485019487945060209093019201613334565b906040519182815491828252602082019060005260206000209260005b818110613394575050610d4192500383610d10565b845483526001948501948794506020909301920161337f565b9063ffffffff6133ca6020929594956040855260408501906131be565b9416910152565b3561057a816110f8565b63ffffffff6000199116019063ffffffff82116132c757565b60405190610140820182811067ffffffffffffffff821117610cd357604052606061012083600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201528260e0820152826101008201520152565b9061345f82612fc0565b61346c6040519182610d10565b828152809261347d601f1991612fc0565b019060005b82811061348e57505050565b6020906134996133f4565b82828501015201613482565b60001981146132c75760010190565b604051906134c182610cd8565b6060602083600081520152565b60019060008155016134e08154612d02565b90816134ea575050565b81601f600093116001146134fc575055565b8183526020832061351891601f0160051c810190600101612d5a565b808252602082209081548360011b9084198560031b1c191617905555565b602061354f91816040519382858094519384920161041f565b8101600281520301902090565b60209061357692826040519483868095519384920161041f565b82019081520301902090565b60208183031261027c5780519067ffffffffffffffff821161027c570181601f8201121561027c5780516135b581610d82565b926135c36040519485610d10565b8184526020828401011161027c5761057a916020808501910161041f565b6040513d6000823e3d90fd5b6006906000815560006001820155600060028201556000600382015560006004820155018054906000815581613621575050565b6000526020600020908101905b818110613639575050565b6000815560010161362e565b9190811015612c7b5760051b81013590603e198136030182121561027c570190565b60408136031261027c576040519061367e82610cd8565b61368781612b66565b825260208101359067ffffffffffffffff821161027c576136aa91369101611320565b602082015290565b60016020919392936136e46001600160a01b0386511682906001600160a01b03166001600160a01b0319825416179055565b0192015191825167ffffffffffffffff8111610cd35761370881612dd88454612d02565b6020601f8211600114613739578190612d56939495600092612e105750508160011b916000199060031b1c19161790565b601f1982169061374e84600052602060002090565b9160005b81811061377057509583600195969710612e5357505050811b019055565b9192602060018192868b015181550194019201613752565b9061379282612fc0565b61379f6040519182610d10565b82815280926137b0601f1991612fc0565b019060005b8281106137c157505050565b8060606020809385010152016137b5565b906137db612fd8565b5061381561381060056137f8856000526010602052604060002090565b01610ade610ad1866000526010602052604060002090565b613362565b61381f8151613788565b9160005b825181101561385b578061383f61140f61098b60019487613094565b6138498287613094565b526138548186613094565b5001613823565b50929050613876611085826000526010602052604060002090565b91600261388d836000526010602052604060002090565b015460016138a5846000526010602052604060002090565b01549060036138be856000526010602052604060002090565b01549060046138d7866000526010602052604060002090565b0154926139676138f4610ad1886000526010602052604060002090565b9661395a6139356006610af2613924613917866000526010602052604060002090565b5460401c63ffffffff1690565b946000526010602052604060002090565b9861394d613941610d52565b63ffffffff909c168c52565b63ffffffff1660208b0152565b63ffffffff166040890152565b6060870152608086015260a085015260c084015260e083015261010082015290565b6020908260405193849283378101600281520301902090565b90918060409360208452816020850152848401376000828201840152601f01601f1916010190565b903590601e198136030182121561027c570180359067ffffffffffffffff821161027c57602001918160051b3603831361027c57565b8015150361027c57565b3561057a81613a00565b60ff81160361027c57565b3561057a81613a14565b903590601e198136030182121561027c570180359067ffffffffffffffff821161027c5760200191813603831361027c57565b929190613a6881612fc0565b93613a766040519586610d10565b602085838152019160051b810192831161027c57905b828210613a9857505050565b8135815260209182019101613a8c565b92919092613ab584612fc0565b93613ac36040519586610d10565b602085828152019060051b82019183831161027c5780915b838310613ae9575050505050565b823567ffffffffffffffff811161027c57820160408187031261027c5760405191613b1383610cd8565b813567ffffffffffffffff811161027c5787613b30918401611320565b835260208201359267ffffffffffffffff841161027c57613b5688602095869501611320565b83820152815201920191613adb565b90613b6f82612fc0565b613b7c6040519182610d10565b8281528092613b8d601f1991612fc0565b019060005b828110613b9e57505050565b602090613ba96134b4565b82828501015201613b92565b90602061057a9281815201906131be565b90821015612c7b576102ac9160051b810190613a29565b9190811015612c7b5760051b8101359060fe198136030182121561027c570190565b9080601f8301121561027c5781602061057a93359101613aa8565b9080601f8301121561027c5781602061057a93359101613a5c565b3590610d4182613a14565b3590610d4182613a00565b6101008136031261027c57613c5e610d62565b90803567ffffffffffffffff811161027c57613c7d9036908301611320565b8252602081013567ffffffffffffffff811161027c57613ca090369083016130ca565b6020830152604081013567ffffffffffffffff811161027c57613cc69036908301611320565b6040830152606081013567ffffffffffffffff811161027c57613cec9036908301613bff565b6060830152608081013567ffffffffffffffff811161027c57613d4591613d1860e09236908301613c1a565b6080850152613d2960a08201613c35565b60a0850152613d3a60c08201613c40565b60c085015201613c40565b60e082015290565b908015613e86579060005b828110613d6457505050565b613d79613d748285859795613bdd565b613c4b565b92613d8d60165463ffffffff9060201c1690565b92613dbd613d9a856132b0565b67ffffffff000000006016549160201b169067ffffffff00000000191617601655565b613e44846080870151606088015190613dd960c08a0151151590565b89613e35613dea60e0830151151590565b613e2b613dfb60a085015160ff1690565b91613e226040865196015196613e12613941610d43565b600160208c0152151560408b0152565b15156060890152565b60ff166080870152565b60a085015260c084015261471a565b602060009501945b85518051821015613e745790613e6e613e6782600194613094565b5187614eb6565b01613e4c565b50509493509150600101919091613d58565b5050565b92939163ffffffff613eaf610ad18663ffffffff166000526011602052604060002090565b1615613f615760005b818110613f415750505060005b818110613ed25750505050565b613eed613ee3611b2d838588613bc6565b6020815191012090565b90613f2461030183613f0f8763ffffffff166000526012602052604060002090565b60019160005201602052604060002054151590565b613f3a57613f3460019285614f80565b01613ec5565b5050505050565b80613f5b613f55611b2d6001948688613bc6565b87614eb6565b01613eb8565b632b62be9b60e01b60005263ffffffff841660045260246000fd5b60405190613f8982610cf4565b6060808381815260006020820152600060408201520152565b613faa613f7c565b5061401c815160208301208060005260036020526002604060002001908060005260036020526001600160a01b03613ffa8160016040600020015416926000526007602052604060002054151590565b916040519561400887610cf4565b865216602085015215156040840152613241565b606082015290565b6001600160a01b0360015416330361403857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60206000604051828101906301ffc9a760e01b82526301ffc9a760e01b602482015260248152614093604482610d10565b519084617530fa903d6000519083614175575b508261416b575b50816140e9575b816140bd575090565b61057a91507f78bea721000000000000000000000000000000000000000000000000000000009061516f565b905060206000604051828101906301ffc9a760e01b82527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152614135604482610d10565b519084617530fa6000513d8261415f575b5081614155575b5015906140b4565b905015153861414d565b60201115915038614146565b15159150386140ad565b602011159250386140a6565b6000929181549161419183612d02565b92600181169081156141dc57506001146141aa57505050565b909192935060005260206000206000905b8382106141c85750500190565b6001816020925484860152019101906141bb565b60ff191683525050811515909102019150565b60206142019160405192838092614181565b600281520301902090565b6142268163ffffffff166000526011602052604060002090565b9081549261423b8463ffffffff9060201c1690565b90600184019061426861425e84849063ffffffff16600052602052604060002090565b9660401c60ff1690565b9260005b87548110156142df5760019085156142b9576142b361429e61428e838c615497565b6000526010602052604060002090565b80546bffffffff000000000000000019169055565b0161426c565b6142d960066142cb61428e848d615497565b0163ffffffff8916906153df565b506142b3565b5094549195509293915060201c63ffffffff16156122a25760005b6143148463ffffffff166000526012602052604060002090565b54811015614351578061434b6143456001936143408863ffffffff166000526012602052604060002090565b615497565b86614f80565b016142fa565b506005614375611f699261437b94969063ffffffff16600052602052604060002090565b016141ef565b60006143978263ffffffff166000526011602052604060002090565b557ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c158170365163ffffffff604051921691806143d481906000602083019252565b0390a2565b906143e382612fc0565b6143f06040519182610d10565b8281528092614401601f1991612fc0565b019060005b82811061441257505050565b60209060405161442181610cd8565b6060815260608382015282828501015201614406565b90916144416133f4565b5061445c8263ffffffff166000526011602052604060002090565b61447984600183019063ffffffff16600052602052604060002090565b61448560038201613362565b9161449083516143d9565b94600683019460005b87518110156144f357806144b561140f61098b6001948a613094565b6144d26144c28a8361355c565b6144ca610d32565b928352613241565b60208201526144e1828b613094565b526144ec818a613094565b5001614499565b50935093509394909461451f6145198563ffffffff166000526012602052604060002090565b54613788565b9560005b61453d8663ffffffff166000526012602052604060002090565b54811015614598578061457c61140f61456c6001946143408b63ffffffff166000526012602052604060002090565b6000526014602052604060002090565b614586828b613094565b52614591818a613094565b5001614523565b509295919490935054936145af8563ffffffff1690565b9460401c60ff166004840154600881901c60ff169060ff16906145d186613362565b936145da610d72565b63ffffffff909916895263ffffffff16602089015260ff166040880152151560608701521515608086015260a085015260c084015261461b60058201613241565b60e084015260020161462c90613241565b61010083015261012082015290565b61206961466e60206001600160a01b03845116930151604051928391602083019586526040808401526060830190610442565b51902090565b60ff60019116019060ff82116132c757565b61057a9054612d02565b60409063ffffffff61057a94931681528160208201520190610442565b60409061057a939281528160208201520190610442565b8054821015612c7b5760005260206000200190600090565b80549068010000000000000000821015610cd35781614703916001612d56940181556146c4565b819391549060031b91821b91600019901b19161790565b9190600161474561472f845163ffffffff1690565b63ffffffff166000526011602052604060002090565b0190602083019461477361475d875163ffffffff1690565b849063ffffffff16600052602052604060002090565b917f00000000000000000000000000000000000000000000000000000000000000001580614e9e575b8015614e7b575b614e335760a085019384515115614df0576147ce90610ade6147c98a5163ffffffff1690565b6133db565b6147dd613ee360058301613241565b8551906147ee826020815191012090565b03614d73575b5060016148086110108a5163ffffffff1690565b11614d0f575b5061483a614820865163ffffffff1690565b6117cf8163ffffffff166000526011602052604060002090565b6148e5606086019461488661484f8751151590565b61486061472f8a5163ffffffff1690565b9068ff0000000000000000825491151560401b169068ff00000000000000001916179055565b6148ca6148978a5163ffffffff1690565b6148a861472f8a5163ffffffff1690565b9067ffffffff0000000082549160201b169067ffffffff000000001916179055565b6148db60c088015160028701612db6565b5160058501612db6565b61492a6148f56040870151151590565b61490f6004860191829060ff801983541691151516179055565b6080870151815461ff00191660089190911b61ff0016179055565b60005b8651811015614abc5761494d610301614946838a613094565b5186615131565b614a6e57845115614a4157614979613917614968838a613094565b516000526010602052604060002090565b63ffffffff61498f611010895163ffffffff1690565b9116141580614a23575b6149ee576001906149e86149b1885163ffffffff1690565b6149be614968848c613094565b906bffffffff000000000000000082549160401b16906bffffffff00000000000000001916179055565b0161492d565b86614a046103e09261094c895163ffffffff1690565b516360b9df7360e01b60005263ffffffff909116600452602452604490565b5063ffffffff614a39613917614968848b613094565b161515614999565b80614a686006614a566149686001958c613094565b016108216110108a5163ffffffff1690565b506149e8565b86614a846103e09261094c895163ffffffff1690565b517f636e40570000000000000000000000000000000000000000000000000000000060005263ffffffff909116600452602452604490565b50959092509392936000916003600682019101955b8751841015614ca857614ae48489613094565b5197614af589516020815191012090565b95614b10610301886000526005602052604060002054151590565b614c8a57614b2b876000526007602052604060002054151590565b614c6c57614b42614b3d858c5161355c565b614686565b614c255760005b8851811015614bcd57614b7d61030189613f0f8c610ade610ad1614968886005614b766149688388613094565b0194613094565b614b8957600101614b49565b614b948b918a613094565b5190519061041b6040519283927f4b5786e7000000000000000000000000000000000000000000000000000000008452600484016146ad565b5094909298614c1a90614be3600194988b6146dc565b8860208201614bf98d610322835191865161355c565b895163ffffffff1692614c10895163ffffffff1690565b9051915193615506565b019293969096614ad1565b89614c34845163ffffffff1690565b90519061041b6040519283927f368812ac00000000000000000000000000000000000000000000000000000000845260048401614690565b61041b8a516040519182916388c8a73760e01b835260048301610d9e565b61041b8a516040519182916327fcf24560e11b835260048301610d9e565b905063ffffffff939650614cea9195507ff264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c15817036519450614cf592505163ffffffff1690565b935163ffffffff1690565b60405163ffffffff909116815292169180602081016143d4565b9693916000969391965b8854811015614d665780614d4f6006614d3761428e6001958e615497565b01614d496110108b5163ffffffff1690565b906153df565b50614d6061429e61428e838d615497565b01614d19565b509193965091943861480e565b614d8461108563ffffffff92613536565b16614db957614d98611f69600586016141ef565b614db3614da9875163ffffffff1690565b6117cf8751613536565b386147f4565b61041b85516040519182917f07bf02d600000000000000000000000000000000000000000000000000000000835260048301610d9e565b6103e0614e01875163ffffffff1690565b7f1caf5f2f0000000000000000000000000000000000000000000000000000000060005263ffffffff16600452602490565b6103e086614e45608088015160ff1690565b90517f25b4d6180000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b50614e92614e8d608087015160ff1690565b614674565b60ff87519116116147a3565b5060ff614eaf608087015160ff1690565b161561479c565b90805160208201209063ffffffff831692836000526012602052614eee83604060002060019160005201602052604060002054151590565b614f7a57826125fb614f4c92614f52956000526014602052614f14856040600020612db6565b614f1d836150fc565b50826000526013602052614f35876040600020615131565b5063ffffffff166000526012602052604060002090565b50612fa0565b907fc00ca38a0d4dd24af204fcc9a39d94708b58426bcf57796b94c4b5437919ede2600080a3565b50505050565b63ffffffff1690816000526012602052614f9e8160406000206153df565b50806000526013602052614fb68260406000206153df565b5080600052601360205260406000205415615013575b6000526014602052614fe8604060002060405191828092614181565b039020907f257129637d1e1b80e89cae4f5e49de63c09628e1622724b24dd19b406627de30600080a3565b61501c81615354565b50614fcc565b600081815260056020526040902054615057576150408160046146dc565b600454906000526005602052604060002055600190565b50600090565b6000818152600960205260409020546150575761507b8160086146dc565b600854906000526009602052604060002055600190565b600081815260076020526040902054615057576150b08160066146dc565b600654906000526007602052604060002055600190565b6000818152600b6020526040902054615057576150e581600a6146dc565b600a5490600052600b602052604060002055600190565b6000818152600d60205260409020546150575761511a81600c6146dc565b600c5490600052600d602052604060002055600190565b60008281526001820160205260409020546151685780615153836001936146dc565b80549260005201602052604060002055600190565b5050600090565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201926301ffc9a760e01b8452166024820152602481526151bc604482610d10565b5191617530fa6000513d826151dd575b50816151d6575090565b9050151590565b602011159150386151cc565b8054801561521257600019019061520082826146c4565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b600081815260096020526040902054908115615168576000198201908282116132c7576008546000198101939084116132c7578383615288946000960361528e575b50505061527760086151e9565b600990600052602052604060002090565b55600190565b6152776152ba916152b06152a66152c09560086146c4565b90549060031b1c90565b92839160086146c4565b90612d3c565b5538808061526a565b6000818152600b6020526040902054908115615168576000198201908282116132c757600a546000198101939084116132c75783836152889460009603615329575b505050615318600a6151e9565b600b90600052602052604060002090565b6153186152ba916153416152a661534b95600a6146c4565b928391600a6146c4565b5538808061530b565b6000818152600d6020526040902054908115615168576000198201908282116132c757600c546000198101939084116132c757838361528894600096036153b4575b5050506153a3600c6151e9565b600d90600052602052604060002090565b6153a36152ba916153cc6152a66153d695600c6146c4565b928391600c6146c4565b55388080615396565b600181019180600052826020526040600020549283151560001461547b5760001984018481116132c75783546000198101949085116132c7576000958583615433946152889803615442575b5050506151e9565b90600052602052604060002090565b6154626152ba916154596152a661547295886146c4565b928391876146c4565b8590600052602052604060002090565b5538808061542b565b50505050600090565b805415612c7b5760005260206000205490565b906152a6916146c4565b9294939160808401608085528251809152602060a0860193019060005b8181106154f0575050509163ffffffff6154e383606095878496036020890152610442565b9616604085015216910152565b82518552602094850194909201916001016154be565b939091602081519101206001600160a01b0380600161552f846000526003602052604060002090565b0154161661553e575050505050565b610699610699600161143961555d946000526003602052604060002090565b90813b1561027c57600080946155a2604051978896879586947ffba64a7c000000000000000000000000000000000000000000000000000000008652600486016154a1565b03925af18015611506576155ba575b80808080613f3a565b806155c960006155cf93610d10565b80610c96565b386155b156fea164736f6c634300081a000a",
}

var CapabilitiesRegistryABI = CapabilitiesRegistryMetaData.ABI

var CapabilitiesRegistryBin = CapabilitiesRegistryMetaData.Bin

func DeployCapabilitiesRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, params CapabilitiesRegistryConstructorParams) (common.Address, *types.Transaction, *CapabilitiesRegistry, error) {
	parsed, err := CapabilitiesRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CapabilitiesRegistryBin), backend, params)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapability(opts *bind.CallOpts, capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapability", capabilityId)

	if err != nil {
		return *new(CapabilitiesRegistryCapabilityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryCapabilityInfo)).(*CapabilitiesRegistryCapabilityInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapability(capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapability(capabilityId string) (CapabilitiesRegistryCapabilityInfo, error) {
	return _CapabilitiesRegistry.Contract.GetCapability(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId string) ([]byte, []byte, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getCapabilityConfigs", donId, capabilityId)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetCapabilityConfigs(donId uint32, capabilityId string) ([]byte, []byte, error) {
	return _CapabilitiesRegistry.Contract.GetCapabilityConfigs(&_CapabilitiesRegistry.CallOpts, donId, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetCapabilityConfigs(donId uint32, capabilityId string) ([]byte, []byte, error) {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONByName(opts *bind.CallOpts, donName string) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONByName", donName)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONByName(donName string) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONByName(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONByName(donName string) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetDONByName(&_CapabilitiesRegistry.CallOpts, donName)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetDONFamilies(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getDONFamilies")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetDONFamilies() ([]string, error) {
	return _CapabilitiesRegistry.Contract.GetDONFamilies(&_CapabilitiesRegistry.CallOpts)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetDONFamilies() ([]string, error) {
	return _CapabilitiesRegistry.Contract.GetDONFamilies(&_CapabilitiesRegistry.CallOpts)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) GetHistoricalDONInfo(opts *bind.CallOpts, donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "getHistoricalDONInfo", donId, configCount)

	if err != nil {
		return *new(CapabilitiesRegistryDONInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CapabilitiesRegistryDONInfo)).(*CapabilitiesRegistryDONInfo)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) GetHistoricalDONInfo(donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetHistoricalDONInfo(&_CapabilitiesRegistry.CallOpts, donId, configCount)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) GetHistoricalDONInfo(donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error) {
	return _CapabilitiesRegistry.Contract.GetHistoricalDONInfo(&_CapabilitiesRegistry.CallOpts, donId, configCount)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryCaller) IsCapabilityDeprecated(opts *bind.CallOpts, capabilityId string) (bool, error) {
	var out []interface{}
	err := _CapabilitiesRegistry.contract.Call(opts, &out, "isCapabilityDeprecated", capabilityId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) IsCapabilityDeprecated(capabilityId string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, capabilityId)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryCallerSession) IsCapabilityDeprecated(capabilityId string) (bool, error) {
	return _CapabilitiesRegistry.Contract.IsCapabilityDeprecated(&_CapabilitiesRegistry.CallOpts, capabilityId)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) AddDONs(opts *bind.TransactOpts, newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "addDONs", newDONs)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) AddDONs(newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDONs(&_CapabilitiesRegistry.TransactOpts, newDONs)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) AddDONs(newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.AddDONs(&_CapabilitiesRegistry.TransactOpts, newDONs)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) DeprecateCapabilities(opts *bind.TransactOpts, capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "deprecateCapabilities", capabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) DeprecateCapabilities(capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilityIds)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) DeprecateCapabilities(capabilityIds []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.DeprecateCapabilities(&_CapabilitiesRegistry.TransactOpts, capabilityIds)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) RemoveDONsByName(opts *bind.TransactOpts, donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "removeDONsByName", donNames)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) RemoveDONsByName(donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONsByName(&_CapabilitiesRegistry.TransactOpts, donNames)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) RemoveDONsByName(donNames []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.RemoveDONsByName(&_CapabilitiesRegistry.TransactOpts, donNames)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) SetDONFamilies(opts *bind.TransactOpts, donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "setDONFamilies", donId, addToFamilies, removeFromFamilies)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) SetDONFamilies(donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamilies(&_CapabilitiesRegistry.TransactOpts, donId, addToFamilies, removeFromFamilies)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) SetDONFamilies(donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.SetDONFamilies(&_CapabilitiesRegistry.TransactOpts, donId, addToFamilies, removeFromFamilies)
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

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDON(opts *bind.TransactOpts, donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDON", donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDON(donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDON(donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDON(&_CapabilitiesRegistry.TransactOpts, donId, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactor) UpdateDONByName(opts *bind.TransactOpts, donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.contract.Transact(opts, "updateDONByName", donName, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistrySession) UpdateDONByName(donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, updateDONParams)
}

func (_CapabilitiesRegistry *CapabilitiesRegistryTransactorSession) UpdateDONByName(donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error) {
	return _CapabilitiesRegistry.Contract.UpdateDONByName(&_CapabilitiesRegistry.TransactOpts, donName, updateDONParams)
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
	CapabilityId common.Hash
	Raw          types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityConfigured(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityConfiguredIterator, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityConfigured", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityConfiguredIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityConfigured", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, capabilityId []string) (event.Subscription, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityConfigured", capabilityIdRule)
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
	CapabilityId common.Hash
	Raw          types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterCapabilityDeprecated(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "CapabilityDeprecated", capabilityIdRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryCapabilityDeprecatedIterator{contract: _CapabilitiesRegistry.contract, event: "CapabilityDeprecated", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, capabilityId []string) (event.Subscription, error) {

	var capabilityIdRule []interface{}
	for _, capabilityIdItem := range capabilityId {
		capabilityIdRule = append(capabilityIdRule, capabilityIdItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "CapabilityDeprecated", capabilityIdRule)
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

type CapabilitiesRegistryDONAddedToFamilyIterator struct {
	Event *CapabilitiesRegistryDONAddedToFamily

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONAddedToFamily)
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
		it.Event = new(CapabilitiesRegistryDONAddedToFamily)
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

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONAddedToFamilyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONAddedToFamily struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONAddedToFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONAddedToFamilyIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONAddedToFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONAddedToFamilyIterator{contract: _CapabilitiesRegistry.contract, event: "DONAddedToFamily", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONAddedToFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONAddedToFamily, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONAddedToFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONAddedToFamily)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONAddedToFamily", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONAddedToFamily(log types.Log) (*CapabilitiesRegistryDONAddedToFamily, error) {
	event := new(CapabilitiesRegistryDONAddedToFamily)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONAddedToFamily", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CapabilitiesRegistryDONRemovedFromFamilyIterator struct {
	Event *CapabilitiesRegistryDONRemovedFromFamily

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CapabilitiesRegistryDONRemovedFromFamily)
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
		it.Event = new(CapabilitiesRegistryDONRemovedFromFamily)
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

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Error() error {
	return it.fail
}

func (it *CapabilitiesRegistryDONRemovedFromFamilyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CapabilitiesRegistryDONRemovedFromFamily struct {
	DonId     uint32
	DonFamily common.Hash
	Raw       types.Log
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) FilterDONRemovedFromFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONRemovedFromFamilyIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.FilterLogs(opts, "DONRemovedFromFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return &CapabilitiesRegistryDONRemovedFromFamilyIterator{contract: _CapabilitiesRegistry.contract, event: "DONRemovedFromFamily", logs: logs, sub: sub}, nil
}

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) WatchDONRemovedFromFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONRemovedFromFamily, donId []uint32, donFamily []string) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var donFamilyRule []interface{}
	for _, donFamilyItem := range donFamily {
		donFamilyRule = append(donFamilyRule, donFamilyItem)
	}

	logs, sub, err := _CapabilitiesRegistry.contract.WatchLogs(opts, "DONRemovedFromFamily", donIdRule, donFamilyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CapabilitiesRegistryDONRemovedFromFamily)
				if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONRemovedFromFamily", log); err != nil {
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

func (_CapabilitiesRegistry *CapabilitiesRegistryFilterer) ParseDONRemovedFromFamily(log types.Log) (*CapabilitiesRegistryDONRemovedFromFamily, error) {
	event := new(CapabilitiesRegistryDONRemovedFromFamily)
	if err := _CapabilitiesRegistry.contract.UnpackLog(event, "DONRemovedFromFamily", log); err != nil {
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
	case _CapabilitiesRegistry.abi.Events["DONAddedToFamily"].ID:
		return _CapabilitiesRegistry.ParseDONAddedToFamily(log)
	case _CapabilitiesRegistry.abi.Events["DONRemovedFromFamily"].ID:
		return _CapabilitiesRegistry.ParseDONRemovedFromFamily(log)
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
	return common.HexToHash("0xe671cf109707667795a875c19f031bdbc7ed40a130f6dc18a55615a0e0099fbb")
}

func (CapabilitiesRegistryCapabilityDeprecated) Topic() common.Hash {
	return common.HexToHash("0xb2553249d353abf34f62139c85f44b5bdeab968ec0ab296a9bf735b75200ed83")
}

func (CapabilitiesRegistryConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf264aae70bf6a9d90e68e0f9b393f4e7fbea67b063b0f336e0b36c1581703651")
}

func (CapabilitiesRegistryDONAddedToFamily) Topic() common.Hash {
	return common.HexToHash("0xc00ca38a0d4dd24af204fcc9a39d94708b58426bcf57796b94c4b5437919ede2")
}

func (CapabilitiesRegistryDONRemovedFromFamily) Topic() common.Hash {
	return common.HexToHash("0x257129637d1e1b80e89cae4f5e49de63c09628e1622724b24dd19b406627de30")
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

	GetCapability(opts *bind.CallOpts, capabilityId string) (CapabilitiesRegistryCapabilityInfo, error)

	GetCapabilityConfigs(opts *bind.CallOpts, donId uint32, capabilityId string) ([]byte, []byte, error)

	GetDON(opts *bind.CallOpts, donId uint32) (CapabilitiesRegistryDONInfo, error)

	GetDONByName(opts *bind.CallOpts, donName string) (CapabilitiesRegistryDONInfo, error)

	GetDONFamilies(opts *bind.CallOpts) ([]string, error)

	GetDONs(opts *bind.CallOpts) ([]CapabilitiesRegistryDONInfo, error)

	GetDONsInFamily(opts *bind.CallOpts, donFamily string) ([]*big.Int, error)

	GetHistoricalDONInfo(opts *bind.CallOpts, donId uint32, configCount uint32) (CapabilitiesRegistryDONInfo, error)

	GetNextDONId(opts *bind.CallOpts) (uint32, error)

	GetNode(opts *bind.CallOpts, p2pId [32]byte) (INodeInfoProviderNodeInfo, error)

	GetNodeOperator(opts *bind.CallOpts, nodeOperatorId uint32) (CapabilitiesRegistryNodeOperator, error)

	GetNodeOperators(opts *bind.CallOpts) ([]CapabilitiesRegistryNodeOperator, error)

	GetNodes(opts *bind.CallOpts) ([]INodeInfoProviderNodeInfo, error)

	GetNodesByP2PIds(opts *bind.CallOpts, p2pIds [][32]byte) ([]INodeInfoProviderNodeInfo, error)

	IsCapabilityDeprecated(opts *bind.CallOpts, capabilityId string) (bool, error)

	IsDONNameTaken(opts *bind.CallOpts, donName string) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddCapabilities(opts *bind.TransactOpts, capabilities []CapabilitiesRegistryCapability) (*types.Transaction, error)

	AddDONs(opts *bind.TransactOpts, newDONs []CapabilitiesRegistryNewDONParams) (*types.Transaction, error)

	AddNodeOperators(opts *bind.TransactOpts, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	AddNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	DeprecateCapabilities(opts *bind.TransactOpts, capabilityIds []string) (*types.Transaction, error)

	RemoveDONs(opts *bind.TransactOpts, donIds []uint32) (*types.Transaction, error)

	RemoveDONsByName(opts *bind.TransactOpts, donNames []string) (*types.Transaction, error)

	RemoveNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32) (*types.Transaction, error)

	RemoveNodes(opts *bind.TransactOpts, removedNodeP2PIds [][32]byte) (*types.Transaction, error)

	SetDONFamilies(opts *bind.TransactOpts, donId uint32, addToFamilies []string, removeFromFamilies []string) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDON(opts *bind.TransactOpts, donId uint32, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error)

	UpdateDONByName(opts *bind.TransactOpts, donName string, updateDONParams CapabilitiesRegistryUpdateDONParams) (*types.Transaction, error)

	UpdateNodeOperators(opts *bind.TransactOpts, nodeOperatorIds []uint32, nodeOperators []CapabilitiesRegistryNodeOperator) (*types.Transaction, error)

	UpdateNodes(opts *bind.TransactOpts, nodes []CapabilitiesRegistryNodeParams) (*types.Transaction, error)

	FilterCapabilityConfigured(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityConfiguredIterator, error)

	WatchCapabilityConfigured(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityConfigured, capabilityId []string) (event.Subscription, error)

	ParseCapabilityConfigured(log types.Log) (*CapabilitiesRegistryCapabilityConfigured, error)

	FilterCapabilityDeprecated(opts *bind.FilterOpts, capabilityId []string) (*CapabilitiesRegistryCapabilityDeprecatedIterator, error)

	WatchCapabilityDeprecated(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryCapabilityDeprecated, capabilityId []string) (event.Subscription, error)

	ParseCapabilityDeprecated(log types.Log) (*CapabilitiesRegistryCapabilityDeprecated, error)

	FilterConfigSet(opts *bind.FilterOpts, donId []uint32) (*CapabilitiesRegistryConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryConfigSet, donId []uint32) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CapabilitiesRegistryConfigSet, error)

	FilterDONAddedToFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONAddedToFamilyIterator, error)

	WatchDONAddedToFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONAddedToFamily, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONAddedToFamily(log types.Log) (*CapabilitiesRegistryDONAddedToFamily, error)

	FilterDONRemovedFromFamily(opts *bind.FilterOpts, donId []uint32, donFamily []string) (*CapabilitiesRegistryDONRemovedFromFamilyIterator, error)

	WatchDONRemovedFromFamily(opts *bind.WatchOpts, sink chan<- *CapabilitiesRegistryDONRemovedFromFamily, donId []uint32, donFamily []string) (event.Subscription, error)

	ParseDONRemovedFromFamily(log types.Log) (*CapabilitiesRegistryDONRemovedFromFamily, error)

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
