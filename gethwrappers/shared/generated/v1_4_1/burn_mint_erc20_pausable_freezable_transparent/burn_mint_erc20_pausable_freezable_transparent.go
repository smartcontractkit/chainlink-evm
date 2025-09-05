// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_pausable_freezable_transparent

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

var BurnMintERC20PausableFreezableTransparentMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FREEZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unfreeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AccountFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AccountUnfrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableTransparent__AccountFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableTransparent__AccountNotFrozen\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20PausableFreezableTransparent__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6135d1806100d96000396000f3fe608060405234801561001057600080fd5b506004361061030a5760003560e01c80638456cb591161019c578063a9059cbb116100ee578063d547741f11610097578063dd62ed3e11610071578063dd62ed3e14610868578063e5839836146108cd578063e63ab1e91461092557600080fd5b8063d547741f14610826578063d5abeb0114610839578063d602b9fd1461086057600080fd5b8063cefc1429116100c8578063cefc14291461078c578063cf6eefb714610794578063d5391393146107ff57600080fd5b8063a9059cbb1461075e578063c630948d14610771578063cc8463c81461078457600080fd5b806391d1485411610150578063a1eda53c1161012a578063a1eda53c1461071c578063a217fddf14610743578063a8fa343c1461074b57600080fd5b806391d148541461069c57806395d89b41146107015780639dc29fac1461070957600080fd5b80638d1fdf2f116101815780638d1fdf2f146106445780638da5cb5b146106575780638fd6a6ac1461065f57600080fd5b80638456cb59146105de57806384ef8ffc146105e657600080fd5b80632f2ff15d1161026057806345c8b1a611610209578063649a5ec7116101e3578063649a5ec71461056357806370a082311461057657806379cc6790146105cb57600080fd5b806345c8b1a6146105135780635c975abb14610526578063634e93da1461055057600080fd5b80633f4ba83a1161023a5780633f4ba83a146104e557806340c10f19146104ed57806342966c681461050057600080fd5b80632f2ff15d14610470578063313ce5671461048357806336568abe146104d257600080fd5b80630aa6220b116102c2578063248a9ca31161029c578063248a9ca3146103f4578063282c51f3146104365780632cd77a5a1461045d57600080fd5b80630aa6220b146103b057806318160ddd146103ba57806323b872dd146103e157600080fd5b806306a85f0f116102f357806306a85f0f1461035357806306fdde0314610388578063095ea7b31461039d57600080fd5b806301ffc9a71461030f578063022d63fb14610337575b600080fd5b61032261031d366004613004565b61094c565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff909116815260200161032e565b61037a7f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d36381565b60405190815260200161032e565b610390610ac9565b60405161032e9190613046565b6103226103ab3660046130db565b610b9e565b6103b8610bb6565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025461037a565b6103226103ef366004613105565b610bcc565b61037a610402366004613142565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b61037a7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6103b861046b366004613256565b610bf0565b6103b861047e3660046132fa565b610e8d565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005474010000000000000000000000000000000000000000900460ff1660405160ff909116815260200161032e565b6103b86104e03660046132fa565b610ed2565b6103b8611039565b6103b86104fb3660046130db565b611083565b6103b861050e366004613142565b61118a565b6103b8610521366004613326565b6111bd565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610322565b6103b861055e366004613326565b6112f6565b6103b8610571366004613341565b61130a565b61037a610584366004613326565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b6103b86105d93660046130db565b61131e565b6103b8611352565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161032e565b6103b8610652366004613326565b6113b4565b61061f6115bd565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005473ffffffffffffffffffffffffffffffffffffffff1661061f565b6103226106aa3660046132fa565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610390611602565b6103b86107173660046130db565b611653565b61072461165d565b6040805165ffffffffffff93841681529290911660208301520161032e565b61037a600081565b6103b8610759366004613326565b61171c565b61032261076c3660046130db565b6117be565b6103b861077f366004613326565b6117cc565b61033c611820565b6103b8611901565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161032e565b61037a7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6103b86108343660046132fa565b61197c565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489015461037a565b6103b86119bd565b61037a610876366004613369565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b6103226108db366004613326565b73ffffffffffffffffffffffffffffffffffffffff1660009081527fe4a0d511ce93f7d3bf378a3a2c82dfeda12e9faf72c0533ddcd2be06e2d60f00602052604090205460ff1690565b61037a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b070000000000000000000000000000000000000000000000000000000014806109df57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b80610a2b57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610a7757507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b80610ac357507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610b1a90613393565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4690613393565b8015610b935780601f10610b6857610100808354040283529160200191610b93565b820191906000526020600020905b815481529060010190602001808311610b7657829003601f168201915b505050505091505090565b600033610bac8185856119d0565b5060019392505050565b6000610bc1816119dd565b610bc96119e7565b50565b600033610bda8582856119f4565b610be5858585611ae2565b506001949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610c3b5750825b905060008267ffffffffffffffff166001148015610c585750303b155b905081158015610c66575080155b15610c9d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610cfe5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610d088b8b611b8d565b610d10611b9f565b610d18611b9f565b60007fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489008054600182018b90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8d16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff891617815590508715610e125788881115610e08576040517f25cc7152000000000000000000000000000000000000000000000000000000008152600481018990526024015b60405180910390fd5b610e128789611ba7565b610e1d600088611c03565b50508315610e805784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b81610ec4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ece8282611d0d565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610f3a57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561102a577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610fad575065ffffffffffff8116155b80610fc057504265ffffffffffff821610155b15611001576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610dff565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b6110348383611d51565b505050565b6000611044816119dd565b61104c611daa565b6040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a150565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66110ad816119dd565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff079048901547fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff07904890090600061111b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b905081158015906111345750816111328683613415565b115b15611178576111438582613415565b6040517f25cc7152000000000000000000000000000000000000000000000000000000008152600401610dff91815260200190565b6111828686611ba7565b505050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486111b4816119dd565b610ece82611e41565b7f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d3636111e7816119dd565b73ffffffffffffffffffffffffffffffffffffffff821660009081527fe4a0d511ce93f7d3bf378a3a2c82dfeda12e9faf72c0533ddcd2be06e2d60f00602081905260409091205460ff16611280576040517f069c4fbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff831660008181526020839052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517ff915cd9fe234de6e8d3afe7bf2388d35b2b6d48e8c629a24602019bde79c213a9190a2505050565b6000611301816119dd565b610ece82611e4b565b6000611315816119dd565b610ece82611ecb565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848611348816119dd565b6110348383611f3b565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61137c816119dd565b611384611f50565b6040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001611078565b7f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d3636113de816119dd565b73ffffffffffffffffffffffffffffffffffffffff8216611443576040517f030b079800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610dff565b3073ffffffffffffffffffffffffffffffffffffffff8316036114aa576040517f030b079800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff821660009081527fe4a0d511ce93f7d3bf378a3a2c82dfeda12e9faf72c0533ddcd2be06e2d60f00602081905260409091205460ff1615611544576040517f4494d7ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff831660008181526020839052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f4f2a367e694e71282f29ab5eaa04c4c0be45ac5bf2ca74fb67068b98bdc2887d9190a2505050565b60006115fd7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610b1a90613393565b610ece828261131e565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906116e057504265ffffffffffff831610155b6116ec57600080611713565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6000611727816119dd565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff07904890080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610bac818585611ae2565b6117f67f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610e8d565b610bc97f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610e8d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906118a357504265ffffffffffff8216105b6118d45781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff166118fa565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff16338114611974576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610dff565b610bc9611fc9565b816119b3576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ece82826120fa565b60006119c8816119dd565b610bc961213e565b6110348383836001612149565b610bc9813361226e565b6119f2600080612315565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611adc5781811015611acd576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610dff565b611adc84848484036000612149565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611b32576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff8216611b82576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b6110348383836124ae565b611b956125cb565b610ece8282612632565b6119f26125cb565b73ffffffffffffffffffffffffffffffffffffffff8216611bf7576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b610ece600083836124ae565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083611cfb576000611c6b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614611cb8576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b611d058484612695565b949350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611d47816119dd565b611adc8383611c03565b73ffffffffffffffffffffffffffffffffffffffff81163314611da0576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61103482826127b6565b611db261285a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001611078565b610bc933826128b5565b6000611e55611820565b611e5e42612911565b611e689190613428565b9050611e748282612961565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b6000611ed682612a1c565b611edf42612911565b611ee99190613428565b9050611ef58282612315565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611f468233836119f4565b610ece82826128b5565b611f58612a6b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611e1c565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061203957504265ffffffffffff821610155b1561207a576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610dff565b6120c260006120bd7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b6127b6565b506120ce600083611c03565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154612134816119dd565b611adc83836127b6565b6119f2600080612961565b73ffffffffffffffffffffffffffffffffffffffff841660009081527fe4a0d511ce93f7d3bf378a3a2c82dfeda12e9faf72c0533ddcd2be06e2d60f00602081905260409091205460ff16156121e3576040517f4494d7ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86166004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205460ff161561225b576040517f4494d7ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610dff565b61226785858585612ac7565b5050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610ece576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610dff565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612428574265ffffffffffff821610156123fe576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612428565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff831660009081527fe4a0d511ce93f7d3bf378a3a2c82dfeda12e9faf72c0533ddcd2be06e2d60f00602081905260409091205460ff1615612548576040517f4494d7ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090205460ff16156125c0576040517f4494d7ea00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610dff565b611adc848484612adb565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166119f2576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61263a6125cb565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612686848261348d565b5060048101611adc838261348d565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166127ac5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556127483390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610ac3565b6000915050610ac3565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561282057507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b15612850576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b611d058484612aee565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166119f2576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216612905576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b610ece826000836124ae565b600065ffffffffffff82111561295d576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610dff565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff8816171784559104168015611adc576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b600080612a27611820565b90508065ffffffffffff168365ffffffffffff1611612a4f57612a4a83826135a6565b612a64565b612a6465ffffffffffff841662069780612bcc565b9392505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156119f2576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612acf612a6b565b611adc84848484612be2565b612ae3612a6b565b611034838383612c55565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156127ac5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610ac3565b6000818310612bdb5781612a64565b5090919050565b3073ffffffffffffffffffffffffffffffffffffffff841603612c49576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610dff565b611adc84848484612cc7565b3073ffffffffffffffffffffffffffffffffffffffff831603612cbc576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610dff565b611034838383612e33565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8516612d38576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff8416612d88576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610dff565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260018301602090815260408083209388168352929052208390558115612267578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051612e2491815260200190565b60405180910390a35050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416612e8e5781816002016000828254612e839190613415565b90915550612f409050565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020829052604090205482811015612f14576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610dff565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff8316612f6b576002810180548390039055612f97565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612ff691815260200190565b60405180910390a350505050565b60006020828403121561301657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114612a6457600080fd5b602081526000825180602084015260005b818110156130745760208186018101516040868401015201613057565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146130d657600080fd5b919050565b600080604083850312156130ee57600080fd5b6130f7836130b2565b946020939093013593505050565b60008060006060848603121561311a57600080fd5b613123846130b2565b9250613131602085016130b2565b929592945050506040919091013590565b60006020828403121561315457600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261319b57600080fd5b813567ffffffffffffffff8111156131b5576131b561315b565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156132215761322161315b565b60405281815283820160200185101561323957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561326f57600080fd5b863567ffffffffffffffff81111561328657600080fd5b61329289828a0161318a565b965050602087013567ffffffffffffffff8111156132af57600080fd5b6132bb89828a0161318a565b955050604087013560ff811681146132d257600080fd5b935060608701359250608087013591506132ee60a088016130b2565b90509295509295509295565b6000806040838503121561330d57600080fd5b8235915061331d602084016130b2565b90509250929050565b60006020828403121561333857600080fd5b612a64826130b2565b60006020828403121561335357600080fd5b813565ffffffffffff81168114612a6457600080fd5b6000806040838503121561337c57600080fd5b613385836130b2565b915061331d602084016130b2565b600181811c908216806133a757607f821691505b6020821081036133e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610ac357610ac36133e6565b65ffffffffffff8181168382160190811115610ac357610ac36133e6565b601f82111561103457806000526020600020601f840160051c8101602085101561346d5750805b601f840160051c820191505b818110156122675760008155600101613479565b815167ffffffffffffffff8111156134a7576134a761315b565b6134bb816134b58454613393565b84613446565b6020601f82116001811461350d57600083156134d75750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455612267565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561355b578785015182556020948501946001909201910161353b565b508482101561359757868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff8281168282160390811115610ac357610ac36133e656fea164736f6c634300081a000a",
}

var BurnMintERC20PausableFreezableTransparentABI = BurnMintERC20PausableFreezableTransparentMetaData.ABI

var BurnMintERC20PausableFreezableTransparentBin = BurnMintERC20PausableFreezableTransparentMetaData.Bin

func DeployBurnMintERC20PausableFreezableTransparent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20PausableFreezableTransparent, error) {
	parsed, err := BurnMintERC20PausableFreezableTransparentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20PausableFreezableTransparentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20PausableFreezableTransparent{address: address, abi: *parsed, BurnMintERC20PausableFreezableTransparentCaller: BurnMintERC20PausableFreezableTransparentCaller{contract: contract}, BurnMintERC20PausableFreezableTransparentTransactor: BurnMintERC20PausableFreezableTransparentTransactor{contract: contract}, BurnMintERC20PausableFreezableTransparentFilterer: BurnMintERC20PausableFreezableTransparentFilterer{contract: contract}}, nil
}

type BurnMintERC20PausableFreezableTransparent struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20PausableFreezableTransparentCaller
	BurnMintERC20PausableFreezableTransparentTransactor
	BurnMintERC20PausableFreezableTransparentFilterer
}

type BurnMintERC20PausableFreezableTransparentCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableTransparentTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableTransparentFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableFreezableTransparentSession struct {
	Contract     *BurnMintERC20PausableFreezableTransparent
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableFreezableTransparentCallerSession struct {
	Contract *BurnMintERC20PausableFreezableTransparentCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20PausableFreezableTransparentTransactorSession struct {
	Contract     *BurnMintERC20PausableFreezableTransparentTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableFreezableTransparentRaw struct {
	Contract *BurnMintERC20PausableFreezableTransparent
}

type BurnMintERC20PausableFreezableTransparentCallerRaw struct {
	Contract *BurnMintERC20PausableFreezableTransparentCaller
}

type BurnMintERC20PausableFreezableTransparentTransactorRaw struct {
	Contract *BurnMintERC20PausableFreezableTransparentTransactor
}

func NewBurnMintERC20PausableFreezableTransparent(address common.Address, backend bind.ContractBackend) (*BurnMintERC20PausableFreezableTransparent, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20PausableFreezableTransparentABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20PausableFreezableTransparent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparent{address: address, abi: abi, BurnMintERC20PausableFreezableTransparentCaller: BurnMintERC20PausableFreezableTransparentCaller{contract: contract}, BurnMintERC20PausableFreezableTransparentTransactor: BurnMintERC20PausableFreezableTransparentTransactor{contract: contract}, BurnMintERC20PausableFreezableTransparentFilterer: BurnMintERC20PausableFreezableTransparentFilterer{contract: contract}}, nil
}

func NewBurnMintERC20PausableFreezableTransparentCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20PausableFreezableTransparentCaller, error) {
	contract, err := bindBurnMintERC20PausableFreezableTransparent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentCaller{contract: contract}, nil
}

func NewBurnMintERC20PausableFreezableTransparentTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20PausableFreezableTransparentTransactor, error) {
	contract, err := bindBurnMintERC20PausableFreezableTransparent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentTransactor{contract: contract}, nil
}

func NewBurnMintERC20PausableFreezableTransparentFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20PausableFreezableTransparentFilterer, error) {
	contract, err := bindBurnMintERC20PausableFreezableTransparent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentFilterer{contract: contract}, nil
}

func bindBurnMintERC20PausableFreezableTransparent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20PausableFreezableTransparentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BurnMintERC20PausableFreezableTransparentCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BurnMintERC20PausableFreezableTransparentTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BurnMintERC20PausableFreezableTransparentTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableFreezableTransparent.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BURNERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BURNERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) FREEZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "FREEZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) FREEZERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.FREEZERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) FREEZERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.FREEZERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.MINTERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.MINTERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PAUSERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PAUSERROLE(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Allowance(&_BurnMintERC20PausableFreezableTransparent.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Allowance(&_BurnMintERC20PausableFreezableTransparent.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BalanceOf(&_BurnMintERC20PausableFreezableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BalanceOf(&_BurnMintERC20PausableFreezableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Decimals(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Decimals(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GetCCIPAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GetCCIPAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GetRoleAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts, role)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GetRoleAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts, role)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.HasRole(&_BurnMintERC20PausableFreezableTransparent.CallOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.HasRole(&_BurnMintERC20PausableFreezableTransparent.CallOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) IsFrozen(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "isFrozen", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) IsFrozen(account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.IsFrozen(&_BurnMintERC20PausableFreezableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) IsFrozen(account common.Address) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.IsFrozen(&_BurnMintERC20PausableFreezableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.MaxSupply(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.MaxSupply(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Name() (string, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Name(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Name() (string, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Name(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Owner(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Owner(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Paused() (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Paused(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Paused() (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Paused(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.SupportsInterface(&_BurnMintERC20PausableFreezableTransparent.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.SupportsInterface(&_BurnMintERC20PausableFreezableTransparent.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Symbol() (string, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Symbol(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) Symbol() (string, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Symbol(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableFreezableTransparent.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.TotalSupply(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.TotalSupply(&_BurnMintERC20PausableFreezableTransparent.CallOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Approve(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Approve(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Burn(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Burn(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Burn0(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Burn0(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BurnFrom(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.BurnFrom(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Freeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "freeze", account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Freeze(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Freeze(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GrantRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.GrantRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Initialize(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Initialize(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Mint(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Mint(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "pause")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Pause(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Pause(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RenounceRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RenounceRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RevokeRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RevokeRole(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.SetCCIPAdmin(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.SetCCIPAdmin(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Transfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Transfer(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.TransferFrom(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.TransferFrom(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Unfreeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "unfreeze", account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Unfreeze(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Unfreeze(&_BurnMintERC20PausableFreezableTransparent.TransactOpts, account)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.contract.Transact(opts, "unpause")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Unpause(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentTransactorSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableFreezableTransparent.Contract.Unpause(&_BurnMintERC20PausableFreezableTransparent.TransactOpts)
}

type BurnMintERC20PausableFreezableTransparentAccountFrozenIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentAccountFrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentAccountFrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentAccountFrozen)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentAccountFrozen)
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

func (it *BurnMintERC20PausableFreezableTransparentAccountFrozenIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentAccountFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentAccountFrozen struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterAccountFrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableTransparentAccountFrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "AccountFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentAccountFrozenIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "AccountFrozen", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentAccountFrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "AccountFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentAccountFrozen)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseAccountFrozen(log types.Log) (*BurnMintERC20PausableFreezableTransparentAccountFrozen, error) {
	event := new(BurnMintERC20PausableFreezableTransparentAccountFrozen)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentAccountUnfrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentAccountUnfrozen)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentAccountUnfrozen)
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

func (it *BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentAccountUnfrozen struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterAccountUnfrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "AccountUnfrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "AccountUnfrozen", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentAccountUnfrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "AccountUnfrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentAccountUnfrozen)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseAccountUnfrozen(log types.Log) (*BurnMintERC20PausableFreezableTransparentAccountUnfrozen, error) {
	event := new(BurnMintERC20PausableFreezableTransparentAccountUnfrozen)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentApprovalIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentApproval)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentApproval)
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

func (it *BurnMintERC20PausableFreezableTransparentApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableFreezableTransparentApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentApprovalIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentApproval)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseApproval(log types.Log) (*BurnMintERC20PausableFreezableTransparentApproval, error) {
	event := new(BurnMintERC20PausableFreezableTransparentApproval)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred)
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

func (it *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred, error) {
	event := new(BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentInitializedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentInitialized)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentInitialized)
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

func (it *BurnMintERC20PausableFreezableTransparentInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentInitializedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentInitialized)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseInitialized(log types.Log) (*BurnMintERC20PausableFreezableTransparentInitialized, error) {
	event := new(BurnMintERC20PausableFreezableTransparentInitialized)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentPausedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentPaused)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentPaused)
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

func (it *BurnMintERC20PausableFreezableTransparentPausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentPausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentPausedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentPaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentPaused)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParsePaused(log types.Log) (*BurnMintERC20PausableFreezableTransparentPaused, error) {
	event := new(BurnMintERC20PausableFreezableTransparentPaused)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentRoleAdminChanged)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentRoleAdminChanged)
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

func (it *BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentRoleAdminChanged)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleAdminChanged, error) {
	event := new(BurnMintERC20PausableFreezableTransparentRoleAdminChanged)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentRoleGrantedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentRoleGranted)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentRoleGranted)
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

func (it *BurnMintERC20PausableFreezableTransparentRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableTransparentRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentRoleGrantedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentRoleGranted)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleGranted, error) {
	event := new(BurnMintERC20PausableFreezableTransparentRoleGranted)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentRoleRevokedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentRoleRevoked)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentRoleRevoked)
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

func (it *BurnMintERC20PausableFreezableTransparentRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableTransparentRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentRoleRevokedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentRoleRevoked)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleRevoked, error) {
	event := new(BurnMintERC20PausableFreezableTransparentRoleRevoked)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentTransferIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentTransfer)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentTransfer)
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

func (it *BurnMintERC20PausableFreezableTransparentTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableFreezableTransparentTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentTransferIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentTransfer)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseTransfer(log types.Log) (*BurnMintERC20PausableFreezableTransparentTransfer, error) {
	event := new(BurnMintERC20PausableFreezableTransparentTransfer)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableFreezableTransparentUnpausedIterator struct {
	Event *BurnMintERC20PausableFreezableTransparentUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableFreezableTransparentUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableFreezableTransparentUnpaused)
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
		it.Event = new(BurnMintERC20PausableFreezableTransparentUnpaused)
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

func (it *BurnMintERC20PausableFreezableTransparentUnpausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableFreezableTransparentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableFreezableTransparentUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentUnpausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableFreezableTransparentUnpausedIterator{contract: _BurnMintERC20PausableFreezableTransparent.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentUnpaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableFreezableTransparent.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableFreezableTransparentUnpaused)
				if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparentFilterer) ParseUnpaused(log types.Log) (*BurnMintERC20PausableFreezableTransparentUnpaused, error) {
	event := new(BurnMintERC20PausableFreezableTransparentUnpaused)
	if err := _BurnMintERC20PausableFreezableTransparent.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparent) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["AccountFrozen"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseAccountFrozen(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["AccountUnfrozen"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseAccountUnfrozen(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["Approval"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseApproval(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["Initialized"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseInitialized(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["Paused"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParsePaused(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseRoleAdminChanged(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseRoleGranted(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseRoleRevoked(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["Transfer"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseTransfer(log)
	case _BurnMintERC20PausableFreezableTransparent.abi.Events["Unpaused"].ID:
		return _BurnMintERC20PausableFreezableTransparent.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20PausableFreezableTransparentAccountFrozen) Topic() common.Hash {
	return common.HexToHash("0x4f2a367e694e71282f29ab5eaa04c4c0be45ac5bf2ca74fb67068b98bdc2887d")
}

func (BurnMintERC20PausableFreezableTransparentAccountUnfrozen) Topic() common.Hash {
	return common.HexToHash("0xf915cd9fe234de6e8d3afe7bf2388d35b2b6d48e8c629a24602019bde79c213a")
}

func (BurnMintERC20PausableFreezableTransparentApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20PausableFreezableTransparentInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20PausableFreezableTransparentPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BurnMintERC20PausableFreezableTransparentRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20PausableFreezableTransparentRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20PausableFreezableTransparentRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20PausableFreezableTransparentTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC20PausableFreezableTransparentUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_BurnMintERC20PausableFreezableTransparent *BurnMintERC20PausableFreezableTransparent) Address() common.Address {
	return _BurnMintERC20PausableFreezableTransparent.address
}

type BurnMintERC20PausableFreezableTransparentInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	FREEZERROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	PAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

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

	Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error)

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

	FilterAccountFrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableTransparentAccountFrozenIterator, error)

	WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentAccountFrozen, account []common.Address) (event.Subscription, error)

	ParseAccountFrozen(log types.Log) (*BurnMintERC20PausableFreezableTransparentAccountFrozen, error)

	FilterAccountUnfrozen(opts *bind.FilterOpts, account []common.Address) (*BurnMintERC20PausableFreezableTransparentAccountUnfrozenIterator, error)

	WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentAccountUnfrozen, account []common.Address) (event.Subscription, error)

	ParseAccountUnfrozen(log types.Log) (*BurnMintERC20PausableFreezableTransparentAccountUnfrozen, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableFreezableTransparentApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20PausableFreezableTransparentApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableFreezableTransparentCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableFreezableTransparentCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableFreezableTransparentDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20PausableFreezableTransparentInitialized, error)

	FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BurnMintERC20PausableFreezableTransparentPaused, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableFreezableTransparentRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableTransparentRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableFreezableTransparentRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableFreezableTransparentRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableFreezableTransparentTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20PausableFreezableTransparentTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableFreezableTransparentUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableFreezableTransparentUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BurnMintERC20PausableFreezableTransparentUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
