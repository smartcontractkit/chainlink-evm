// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_pausable_freezable_uups

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

var BurnMintERC20PausableFreezableUUPSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FREEZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultUpgrader\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unfreeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AccountFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AccountUnfrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableUUPS__AccountFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableUUPS__AccountNotFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableUUPS__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20UUPS__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613fb06100fd6000396000818161210d01528181612136015261233e0152613fb06000f3fe6080604052600436106103295760003560e01c80638456cb59116101a5578063ad3cb1cc116100ec578063d547741f11610095578063dd62ed3e1161006f578063dd62ed3e14610af2578063e583983614610b64578063e63ab1e914610bc9578063f72c0d8b14610bfd57600080fd5b8063d547741f14610a89578063d5abeb0114610aa9578063d602b9fd14610add57600080fd5b8063cefc1429116100c6578063cefc1429146109c8578063cf6eefb7146109dd578063d539139314610a5557600080fd5b8063ad3cb1cc1461094a578063c630948d14610993578063cc8463c8146109b357600080fd5b806395d89b411161014e578063a217fddf11610128578063a217fddf146108f5578063a8fa343c1461090a578063a9059cbb1461092a57600080fd5b806395d89b411461088c5780639dc29fac146108a1578063a1eda53c146108c157600080fd5b80638da5cb5b1161017f5780638da5cb5b146107f05780638fd6a6ac1461080557806391d148541461081a57600080fd5b80638456cb591461075057806384ef8ffc146107655780638d1fdf2f146107d057600080fd5b806336568abe1161027457806352d1902d1161021d578063634e93da116101f7578063634e93da1461068e578063649a5ec7146106ae57806370a08231146106ce57806379cc67901461073057600080fd5b806352d1902d14610622578063561cf2ab146106375780635c975abb1461065757600080fd5b806342966c681161024e57806342966c68146105cf57806345c8b1a6146105ef5780634f1ef2861461060f57600080fd5b806336568abe1461057a5780633f4ba83a1461059a57806340c10f19146105af57600080fd5b806318160ddd116102d6578063282c51f3116102b0578063282c51f3146104ca5780632f2ff15d146104fe578063313ce5671461051e57600080fd5b806318160ddd1461042757806323b872dd1461045b578063248a9ca31461047b57600080fd5b806306fdde031161030757806306fdde03146103ce578063095ea7b3146103f05780630aa6220b1461041057600080fd5b806301ffc9a71461032e578063022d63fb1461036357806306a85f0f1461038c575b600080fd5b34801561033a57600080fd5b5061034e610349366004613945565b610c31565b60405190151581526020015b60405180910390f35b34801561036f57600080fd5b50620697805b60405165ffffffffffff909116815260200161035a565b34801561039857600080fd5b506103c07f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d36381565b60405190815260200161035a565b3480156103da57600080fd5b506103e3610dfa565b60405161035a91906139ab565b3480156103fc57600080fd5b5061034e61040b366004613a25565b610ecf565b34801561041c57600080fd5b50610425610ee7565b005b34801561043357600080fd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02546103c0565b34801561046757600080fd5b5061034e610476366004613a4f565b610efd565b34801561048757600080fd5b506103c0610496366004613a8c565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156104d657600080fd5b506103c07f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b34801561050a57600080fd5b50610425610519366004613aa5565b610f23565b34801561052a57600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005474010000000000000000000000000000000000000000900460ff1660405160ff909116815260200161035a565b34801561058657600080fd5b50610425610595366004613aa5565b610f68565b3480156105a657600080fd5b506104256110d4565b3480156105bb57600080fd5b506104256105ca366004613a25565b61111e565b3480156105db57600080fd5b506104256105ea366004613a8c565b611202565b3480156105fb57600080fd5b5061042561060a366004613ad1565b611235565b61042561061d366004613bb3565b61136e565b34801561062e57600080fd5b506103c0611389565b34801561064357600080fd5b50610425610652366004613c35565b6113b8565b34801561066357600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661034e565b34801561069a57600080fd5b506104256106a9366004613ad1565b611684565b3480156106ba57600080fd5b506104256106c9366004613ceb565b611698565b3480156106da57600080fd5b506103c06106e9366004613ad1565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b34801561073c57600080fd5b5061042561074b366004613a25565b6116ac565b34801561075c57600080fd5b506104256116e0565b34801561077157600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161035a565b3480156107dc57600080fd5b506104256107eb366004613ad1565b611742565b3480156107fc57600080fd5b506107ab61194b565b34801561081157600080fd5b506107ab611990565b34801561082657600080fd5b5061034e610835366004613aa5565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561089857600080fd5b506103e36119d0565b3480156108ad57600080fd5b506104256108bc366004613a25565b611a21565b3480156108cd57600080fd5b506108d6611a2b565b6040805165ffffffffffff93841681529290911660208301520161035a565b34801561090157600080fd5b506103c0600081565b34801561091657600080fd5b50610425610925366004613ad1565b611aea565b34801561093657600080fd5b5061034e610945366004613a25565b611b8c565b34801561095657600080fd5b506103e36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b34801561099f57600080fd5b506104256109ae366004613ad1565b611b9a565b3480156109bf57600080fd5b50610375611bee565b3480156109d457600080fd5b50610425611ccf565b3480156109e957600080fd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161035a565b348015610a6157600080fd5b506103c07f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b348015610a9557600080fd5b50610425610aa4366004613aa5565b611d4a565b348015610ab557600080fd5b507fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff01546103c0565b348015610ae957600080fd5b50610425611d8b565b348015610afe57600080fd5b506103c0610b0d366004613d13565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b348015610b7057600080fd5b5061034e610b7f366004613ad1565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f36a30f686feb055c8d90421e230dafb8f47433e358189345608518a408badc00602052604090205460ff1690565b348015610bd557600080fd5b506103c07f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610c0957600080fd5b506103c07f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b07000000000000000000000000000000000000000000000000000000001480610cc457507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b80610d1057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610d5c57507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b80610da857507fffffffff0000000000000000000000000000000000000000000000000000000082167f52d1902d00000000000000000000000000000000000000000000000000000000145b80610df457507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610e4b90613d3d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e7790613d3d565b8015610ec45780601f10610e9957610100808354040283529160200191610ec4565b820191906000526020600020905b815481529060010190602001808311610ea757829003601f168201915b505050505091505090565b600033610edd818585611d9e565b5060019392505050565b6000610ef281611dab565b610efa611db5565b50565b600033610f0b858285611dc2565b610f16858585611eb0565b60019150505b9392505050565b81610f5a576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f648282611f5b565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610fd057507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b156110c5577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611043575065ffffffffffff8116155b8061105657504265ffffffffffff821610155b1561109c576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b6110cf8383611f9f565b505050565b60006110df81611dab565b6110e7611ff8565b6040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a150565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661114881611dab565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff015460006111947f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081158015906111ad5750816111ab8583613dbf565b115b156111f1576111bc8482613dbf565b6040517f193e245300000000000000000000000000000000000000000000000000000000815260040161109391815260200190565b6111fb858561208f565b5050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861122c81611dab565b610f64826120eb565b7f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d36361125f81611dab565b73ffffffffffffffffffffffffffffffffffffffff821660009081527f36a30f686feb055c8d90421e230dafb8f47433e358189345608518a408badc00602081905260409091205460ff166112f8576040517fffe16b4a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff831660008181526020839052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517ff915cd9fe234de6e8d3afe7bf2388d35b2b6d48e8c629a24602019bde79c213a9190a2505050565b6113766120f5565b61137f826121c3565b610f6482826121ed565b6000611393612326565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156114035750825b905060008267ffffffffffffffff1660011480156114205750303b155b90508115801561142e575080155b15611465576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156114c65784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6114d08c8c612395565b6114d86123a7565b6114e06123a7565b6114e86123a7565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff008054600182018c90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8e16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff8a16178155905088156115dd57898911156115d3576040517f193e2453000000000000000000000000000000000000000000000000000000008152600481018a9052602401611093565b6115dd888a61208f565b6115e86000896123af565b506116137f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3886123af565b505083156116765784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b600061168f81611dab565b610f64826124b9565b60006116a381611dab565b610f6482612539565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486116d681611dab565b6110cf83836125a9565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61170a81611dab565b6117126125be565b6040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001611113565b7f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d36361176c81611dab565b73ffffffffffffffffffffffffffffffffffffffff82166117d1576040517f175469d300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611093565b3073ffffffffffffffffffffffffffffffffffffffff831603611838576040517f175469d300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff821660009081527f36a30f686feb055c8d90421e230dafb8f47433e358189345608518a408badc00602081905260409091205460ff16156118d2576040517fe9c5a2e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff831660008181526020839052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f4f2a367e694e71282f29ab5eaa04c4c0be45ac5bf2ca74fb67068b98bdc2887d9190a2505050565b600061198b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b60007fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff005b5473ffffffffffffffffffffffffffffffffffffffff16919050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610e4b90613d3d565b610f6482826116ac565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008115801590611aae57504265ffffffffffff831610155b611aba57600080611ae1565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6000611af581611dab565b7fc5cdc2af358d9a68c0b2c9c0cc0618d81d3e3f32ffbcf23d38fc5724f437ff0080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610edd818585611eb0565b611bc47f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610f23565b610efa7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610f23565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015801590611c7157504265ffffffffffff8216105b611ca25781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611cc8565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff16338114611d42576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401611093565b610efa612637565b81611d81576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f648282612768565b6000611d9681611dab565b610efa6127ac565b6110cf83838360016127b7565b610efa81336128d5565b611dc060008061297c565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611eaa5781811015611e9b576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401611093565b611eaa848484840360006127b7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611f00576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff8216611f50576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b6110cf838383612b15565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611f9581611dab565b611eaa83836123af565b73ffffffffffffffffffffffffffffffffffffffff81163314611fee576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110cf8282612c32565b612000612cd6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001611113565b73ffffffffffffffffffffffffffffffffffffffff82166120df576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b610f6460008383612b15565b610efa3382612d31565b3073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148061218c57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612173612d8d565b73ffffffffffffffffffffffffffffffffffffffff1614155b15611dc0576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610f6481611dab565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612272575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261226f91810190613dd2565b60015b6122c0576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611093565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461231c576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611093565b6110cf8383612db5565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611dc0576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61239d612e18565b610f648282612e7f565b611dc0612e18565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400836124a75760006124177feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612464576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b6124b18484612ee2565b949350505050565b60006124c3611bee565b6124cc42613003565b6124d69190613deb565b90506124e28282613053565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b60006125448261310e565b61254d42613003565b6125579190613deb565b9050612563828261297c565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b6125b4823383611dc2565b610f648282612d31565b6125c6613156565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361206a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806126a757504265ffffffffffff821610155b156126e8576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401611093565b612730600061272b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b612c32565b5061273c6000836123af565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546127a281611dab565b611eaa8383612c32565b611dc0600080613053565b73ffffffffffffffffffffffffffffffffffffffff841660009081527f36a30f686feb055c8d90421e230dafb8f47433e358189345608518a408badc00602081905260409091205460ff1615612851576040517fe9c5a2e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205460ff16156128c9576040517fe9c5a2e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611093565b6111fb858585856131b2565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610f64576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401611093565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612a8f574265ffffffffffff82161015612a65576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612a8f565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff831660009081527f36a30f686feb055c8d90421e230dafb8f47433e358189345608518a408badc00602081905260409091205460ff1615612baf576040517fe9c5a2e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090205460ff1615612c27576040517fe9c5a2e000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611093565b611eaa8484846131c6565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083158015612c9c57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b15612ccc576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6124b184846131d9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611dc0576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216612d81576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b610f6482600083612b15565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6119b4565b612dbe826132b7565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612e10576110cf8282613386565b610f64613409565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611dc0576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e87612e18565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612ed38482613e50565b5060048101611eaa8382613e50565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16612ff95760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612f953390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610df4565b6000915050610df4565b600065ffffffffffff82111561304f576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401611093565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff8816171784559104168015611eaa576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b600080613119611bee565b90508065ffffffffffff168365ffffffffffff16116131415761313c8382613f69565b610f1c565b610f1c65ffffffffffff841662069780613441565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611dc0576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131ba613156565b611eaa84848484613457565b6131ce613156565b6110cf8383836134ca565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615612ff95760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610df4565b8073ffffffffffffffffffffffffffffffffffffffff163b600003613320576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401611093565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60606000808473ffffffffffffffffffffffffffffffffffffffff16846040516133b09190613f87565b600060405180830381855af49150503d80600081146133eb576040519150601f19603f3d011682016040523d82523d6000602084013e6133f0565b606091505b509150915061340085838361353c565b95945050505050565b3415611dc0576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008183106134505781610f1c565b5090919050565b3073ffffffffffffffffffffffffffffffffffffffff8416036134be576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611093565b611eaa848484846135c6565b3073ffffffffffffffffffffffffffffffffffffffff831603613531576040517f99817ca200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611093565b6110cf838383613732565b60608261354c5761313c82613903565b8151158015613570575073ffffffffffffffffffffffffffffffffffffffff84163b155b156135bf576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401611093565b5080610f1c565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516613637576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff8416613687576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401611093565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600183016020908152604080832093881683529290522083905581156111fb578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161372391815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff841661378d57818160020160008282546137829190613dbf565b9091555061383f9050565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205482811015613813576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401611093565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff831661386a576002810180548390039055613896565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516138f591815260200190565b60405180910390a350505050565b8051156139135780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561395757600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f1c57600080fd5b60005b838110156139a257818101518382015260200161398a565b50506000910152565b60208152600082518060208401526139ca816040850160208701613987565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613a2057600080fd5b919050565b60008060408385031215613a3857600080fd5b613a41836139fc565b946020939093013593505050565b600080600060608486031215613a6457600080fd5b613a6d846139fc565b9250613a7b602085016139fc565b929592945050506040919091013590565b600060208284031215613a9e57600080fd5b5035919050565b60008060408385031215613ab857600080fd5b82359150613ac8602084016139fc565b90509250929050565b600060208284031215613ae357600080fd5b610f1c826139fc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff841115613b3657613b36613aec565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715613b8357613b83613aec565b604052838152905080828401851015613b9b57600080fd5b83836020830137600060208583010152509392505050565b60008060408385031215613bc657600080fd5b613bcf836139fc565b9150602083013567ffffffffffffffff811115613beb57600080fd5b8301601f81018513613bfc57600080fd5b613c0b85823560208401613b1b565b9150509250929050565b600082601f830112613c2657600080fd5b610f1c83833560208501613b1b565b600080600080600080600060e0888a031215613c5057600080fd5b873567ffffffffffffffff811115613c6757600080fd5b613c738a828b01613c15565b975050602088013567ffffffffffffffff811115613c9057600080fd5b613c9c8a828b01613c15565b965050604088013560ff81168114613cb357600080fd5b94506060880135935060808801359250613ccf60a089016139fc565b9150613cdd60c089016139fc565b905092959891949750929550565b600060208284031215613cfd57600080fd5b813565ffffffffffff81168114610f1c57600080fd5b60008060408385031215613d2657600080fd5b613d2f836139fc565b9150613ac8602084016139fc565b600181811c90821680613d5157607f821691505b602082108103613d8a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610df457610df4613d90565b600060208284031215613de457600080fd5b5051919050565b65ffffffffffff8181168382160190811115610df457610df4613d90565b601f8211156110cf57806000526020600020601f840160051c81016020851015613e305750805b601f840160051c820191505b818110156111fb5760008155600101613e3c565b815167ffffffffffffffff811115613e6a57613e6a613aec565b613e7e81613e788454613d3d565b84613e09565b6020601f821160018114613ed05760008315613e9a5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556111fb565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015613f1e5787850151825560209485019460019092019101613efe565b5084821015613f5a57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff8281168282160390811115610df457610df4613d90565b60008251613f99818460208701613987565b919091019291505056fea164736f6c634300081a000a",
}

var BurnMintERC20PausableFreezableUUPSABI = BurnMintERC20PausableFreezableUUPSMetaData.ABI

var BurnMintERC20PausableFreezableUUPSBin = BurnMintERC20PausableFreezableUUPSMetaData.Bin

func DeployBurnMintERC20PausableFreezableUUPS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20PausableFreezableUUPS, error) {
	parsed, err := BurnMintERC20PausableFreezableUUPSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20PausableFreezableUUPSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20PausableFreezableUUPS{address: address, abi: *parsed, BurnMintERC20PausableFreezableUUPSCaller: BurnMintERC20PausableFreezableUUPSCaller{contract: contract}, BurnMintERC20PausableFreezableUUPSTransactor: BurnMintERC20PausableFreezableUUPSTransactor{contract: contract}, BurnMintERC20PausableFreezableUUPSFilterer: BurnMintERC20PausableFreezableUUPSFilterer{contract: contract}}, nil
}

type BurnMintERC20PausableFreezableUUPS struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20PausableFreezableUUPSCaller
	BurnMintERC20PausableFreezableUUPSTransactor
	BurnMintERC20PausableFreezableUUPSFilterer
}

type BurnMintERC20PausableFreezableUUPSCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableUUPSTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableUUPSFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableUUPSSession struct {
	Contract     *BurnMintERC20PausableFreezableUUPS
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableFreezableUUPSCallerSession struct {
	Contract *BurnMintERC20PausableFreezableUUPSCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20PausableFreezableUUPSTransactorSession struct {
	Contract     *BurnMintERC20PausableFreezableUUPSTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableFreezableUUPSRaw struct {
	Contract *BurnMintERC20PausableFreezableUUPS
}

type BurnMintERC20PausableFreezableUUPSCallerRaw struct {
	Contract *BurnMintERC20PausableFreezableUUPSCaller
}

type BurnMintERC20PausableFreezableUUPSTransactorRaw struct {
	Contract *BurnMintERC20PausableFreezableUUPSTransactor
}

func NewBurnMintERC20PausableFreezableUUPS(address common.Address, backend bind.ContractBackend) (*BurnMintERC20PausableFreezableUUPS, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20PausableFreezableUUPSABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20PausableFreezableUUPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPS{address: address, abi: abi, BurnMintERC20PausableFreezableUUPSCaller: BurnMintERC20PausableFreezableUUPSCaller{contract: contract}, BurnMintERC20PausableFreezableUUPSTransactor: BurnMintERC20PausableFreezableUUPSTransactor{contract: contract}, BurnMintERC20PausableFreezableUUPSFilterer: BurnMintERC20PausableFreezableUUPSFilterer{contract: contract}}, nil
}

func NewBurnMintERC20PausableFreezableUUPSCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20PausableFreezableUUPSCaller, error) {
	contract, err := bindBurnMintERC20PausableFreezableUUPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSCaller{contract: contract}, nil
}

func NewBurnMintERC20PausableFreezableUUPSTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20PausableFreezableUUPSTransactor, error) {
	contract, err := bindBurnMintERC20PausableFreezableUUPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSTransactor{contract: contract}, nil
}

func NewBurnMintERC20PausableFreezableUUPSFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20PausableFreezableUUPSFilterer, error) {
	contract, err := bindBurnMintERC20PausableFreezableUUPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSFilterer{contract: contract}, nil
}

func bindBurnMintERC20PausableFreezableUUPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20PausableFreezableUUPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BurnMintERC20PausableFreezableUUPSCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BurnMintERC20PausableFreezableUUPSTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BurnMintERC20PausableFreezableUUPSTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableFreezableUUPS.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BURNERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BURNERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) FREEZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "FREEZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) FREEZERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.FREEZERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) FREEZERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.FREEZERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.MINTERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.MINTERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PAUSERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PAUSERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UPGRADERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UPGRADERROLE(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UPGRADEINTERFACEVERSION(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Allowance(&_BurnMintERC20PausableFreezableUUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Allowance(&_BurnMintERC20PausableFreezableUUPS.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BalanceOf(&_BurnMintERC20PausableFreezableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BalanceOf(&_BurnMintERC20PausableFreezableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Decimals(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Decimals(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GetCCIPAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GetCCIPAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GetRoleAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts, role)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GetRoleAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts, role)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.HasRole(&_BurnMintERC20PausableFreezableUUPS.CallOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.HasRole(&_BurnMintERC20PausableFreezableUUPS.CallOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) IsFrozen(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "isFrozen", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) IsFrozen(account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.IsFrozen(&_BurnMintERC20PausableFreezableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) IsFrozen(account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.IsFrozen(&_BurnMintERC20PausableFreezableUUPS.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.MaxSupply(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.MaxSupply(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Name() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Name(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Name() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Name(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Owner(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Owner(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Paused() (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Paused(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Paused() (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Paused(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.ProxiableUUID(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.ProxiableUUID(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.SupportsInterface(&_BurnMintERC20PausableFreezableUUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.SupportsInterface(&_BurnMintERC20PausableFreezableUUPS.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Symbol() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Symbol(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) Symbol() (string, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Symbol(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableUUPS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.TotalSupply(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.TotalSupply(&_BurnMintERC20PausableFreezableUUPS.CallOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Approve(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Approve(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Burn(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Burn(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Burn0(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Burn0(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BurnFrom(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.BurnFrom(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Freeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "freeze", account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Freeze(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Freeze(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GrantRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.GrantRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Initialize(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Initialize(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin, defaultUpgrader)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Mint(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Mint(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "pause")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Pause(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Pause(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RenounceRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RenounceRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RevokeRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RevokeRole(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.SetCCIPAdmin(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.SetCCIPAdmin(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Transfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Transfer(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.TransferFrom(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.TransferFrom(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Unfreeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "unfreeze", account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Unfreeze(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Unfreeze(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "unpause")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Unpause(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.Unpause(&_BurnMintERC20PausableFreezableUUPS.TransactOpts)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UpgradeToAndCall(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newImplementation, data)
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableUUPS.Contract.UpgradeToAndCall(&_BurnMintERC20PausableFreezableUUPS.TransactOpts, newImplementation, data)
}

type BurnMintERC20PausableFreezableUUPSAccountFrozenIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSAccountFrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSAccountFrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSAccountFrozen)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSAccountFrozen)
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

func (it *BurnMintERC20PausableFreezableUUPSAccountFrozenIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSAccountFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSAccountFrozen struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterAccountFrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableUUPSAccountFrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "AccountFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSAccountFrozenIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "AccountFrozen", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSAccountFrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "AccountFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSAccountFrozen)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseAccountFrozen(log types.Log) (*BurnMintERC20PausableFreezableUUPSAccountFrozen, error) {
	event := new(BurnMintERC20PausableFreezableUUPSAccountFrozen)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSAccountUnfrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSAccountUnfrozen)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSAccountUnfrozen)
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

func (it *BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSAccountUnfrozen struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterAccountUnfrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "AccountUnfrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "AccountUnfrozen", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSAccountUnfrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "AccountUnfrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSAccountUnfrozen)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseAccountUnfrozen(log types.Log) (*BurnMintERC20PausableFreezableUUPSAccountUnfrozen, error) {
	event := new(BurnMintERC20PausableFreezableUUPSAccountUnfrozen)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSApprovalIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSApproval)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSApproval)
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

func (it *BurnMintERC20PausableFreezableUUPSApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableFreezableUUPSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSApprovalIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSApproval)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseApproval(log types.Log) (*BurnMintERC20PausableFreezableUUPSApproval, error) {
	event := new(BurnMintERC20PausableFreezableUUPSApproval)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred)
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

func (it *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred, error) {
	event := new(BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSInitializedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSInitialized)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSInitialized)
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

func (it *BurnMintERC20PausableFreezableUUPSInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSInitializedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSInitialized)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseInitialized(log types.Log) (*BurnMintERC20PausableFreezableUUPSInitialized, error) {
	event := new(BurnMintERC20PausableFreezableUUPSInitialized)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSPausedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSPaused)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSPaused)
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

func (it *BurnMintERC20PausableFreezableUUPSPausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSPausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSPausedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSPaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSPaused)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParsePaused(log types.Log) (*BurnMintERC20PausableFreezableUUPSPaused, error) {
	event := new(BurnMintERC20PausableFreezableUUPSPaused)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSRoleAdminChanged)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSRoleAdminChanged)
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

func (it *BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSRoleAdminChanged)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleAdminChanged, error) {
	event := new(BurnMintERC20PausableFreezableUUPSRoleAdminChanged)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSRoleGrantedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSRoleGranted)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSRoleGranted)
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

func (it *BurnMintERC20PausableFreezableUUPSRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableUUPSRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSRoleGrantedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSRoleGranted)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleGranted, error) {
	event := new(BurnMintERC20PausableFreezableUUPSRoleGranted)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSRoleRevokedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSRoleRevoked)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSRoleRevoked)
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

func (it *BurnMintERC20PausableFreezableUUPSRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableUUPSRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSRoleRevokedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSRoleRevoked)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleRevoked, error) {
	event := new(BurnMintERC20PausableFreezableUUPSRoleRevoked)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSTransferIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSTransfer)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSTransfer)
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

func (it *BurnMintERC20PausableFreezableUUPSTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableFreezableUUPSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSTransferIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSTransfer)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseTransfer(log types.Log) (*BurnMintERC20PausableFreezableUUPSTransfer, error) {
	event := new(BurnMintERC20PausableFreezableUUPSTransfer)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSUnpausedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSUnpaused)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSUnpaused)
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

func (it *BurnMintERC20PausableFreezableUUPSUnpausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSUnpausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSUnpausedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSUnpaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSUnpaused)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseUnpaused(log types.Log) (*BurnMintERC20PausableFreezableUUPSUnpaused, error) {
	event := new(BurnMintERC20PausableFreezableUUPSUnpaused)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableUUPSUpgradedIterator struct {
	Event *BurnMintERC20PausableFreezableUUPSUpgraded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableUUPSUpgradedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableUUPSUpgraded)
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
		it.Event = new(BurnMintERC20PausableFreezableUUPSUpgraded)
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

func (it *BurnMintERC20PausableFreezableUUPSUpgradedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableUUPSUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableUUPSUpgraded struct {
	Implementation common.Address
	Raw            types.Log
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20PausableFreezableUUPSUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableUUPSUpgradedIterator{contract: _BurnMintERC20PausableFreezableUUPS.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableUUPS.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableUUPSUpgraded)
				if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPSFilterer) ParseUpgraded(log types.Log) (*BurnMintERC20PausableFreezableUUPSUpgraded, error) {
	event := new(BurnMintERC20PausableFreezableUUPSUpgraded)
	if err := _BurnMintERC20PausableFreezableUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPS) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["AccountFrozen"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseAccountFrozen(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["AccountUnfrozen"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseAccountUnfrozen(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Approval"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseApproval(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Initialized"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseInitialized(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Paused"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParsePaused(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseRoleAdminChanged(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseRoleGranted(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseRoleRevoked(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Transfer"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseTransfer(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Unpaused"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseUnpaused(log)
	case _BurnMintERC20PausableFreezableUUPS.abi.Events["Upgraded"].ID:
		return _BurnMintERC20PausableFreezableUUPS.ParseUpgraded(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20PausableFreezableUUPSAccountFrozen) Topic() common.Hash {
	return common.HexToHash("0x4f2a367e694e71282f29ab5eaa04c4c0be45ac5bf2ca74fb67068b98bdc2887d")
}

func (BurnMintERC20PausableFreezableUUPSAccountUnfrozen) Topic() common.Hash {
	return common.HexToHash("0xf915cd9fe234de6e8d3afe7bf2388d35b2b6d48e8c629a24602019bde79c213a")
}

func (BurnMintERC20PausableFreezableUUPSApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20PausableFreezableUUPSInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20PausableFreezableUUPSPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BurnMintERC20PausableFreezableUUPSRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20PausableFreezableUUPSRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20PausableFreezableUUPSRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20PausableFreezableUUPSTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC20PausableFreezableUUPSUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (BurnMintERC20PausableFreezableUUPSUpgraded) Topic() common.Hash {
	return common.HexToHash("0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b")
}

func (_BurnMintERC20PausableFreezableUUPS *BurnMintERC20PausableFreezableUUPS) Address() common.Address {
	return _BurnMintERC20PausableFreezableUUPS.address
}

type BurnMintERC20PausableFreezableUUPSInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	FREEZERROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	PAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

	UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error)

	UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error)

	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsFrozen(opts *bind.CallOpts, account common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	ProxiableUUID(opts *bind.CallOpts) ([32]byte, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	Freeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address, defaultUpgrader common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	Unfreeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error)

	FilterAccountFrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableUUPSAccountFrozenIterator, error)

	WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSAccountFrozen, account []common.Address) (event.Subscription, error)

	ParseAccountFrozen(log types.Log) (*BurnMintERC20PausableFreezableUUPSAccountFrozen, error)

	FilterAccountUnfrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableUUPSAccountUnfrozenIterator, error)

	WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSAccountUnfrozen, account []common.Address) (event.Subscription, error)

	ParseAccountUnfrozen(log types.Log) (*BurnMintERC20PausableFreezableUUPSAccountUnfrozen, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableFreezableUUPSApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20PausableFreezableUUPSApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableFreezableUUPSCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableFreezableUUPSCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableFreezableUUPSDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20PausableFreezableUUPSInitialized, error)

	FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BurnMintERC20PausableFreezableUUPSPaused, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableFreezableUUPSRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableUUPSRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableUUPSRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableFreezableUUPSRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableFreezableUUPSTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20PausableFreezableUUPSTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableUUPSUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BurnMintERC20PausableFreezableUUPSUnpaused, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BurnMintERC20PausableFreezableUUPSUpgradedIterator, error)

	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableUUPSUpgraded, implementation []common.Address) (event.Subscription, error)

	ParseUpgraded(log types.Log) (*BurnMintERC20PausableFreezableUUPSUpgraded, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
