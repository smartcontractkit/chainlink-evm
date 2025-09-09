// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_pausable_transparent

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

var BurnMintERC20PausableTransparentMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BurnMintERC20Transparent__MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b612f79806100d96000396000f3fe608060405234801561001057600080fd5b50600436106102de5760003560e01c80638456cb5911610186578063a9059cbb116100e3578063d539139311610097578063d602b9fd11610071578063d602b9fd146107e3578063dd62ed3e146107eb578063e63ab1e91461085057600080fd5b8063d539139314610782578063d547741f146107a9578063d5abeb01146107bc57600080fd5b8063cc8463c8116100c8578063cc8463c814610707578063cefc14291461070f578063cf6eefb71461071757600080fd5b8063a9059cbb146106e1578063c630948d146106f457600080fd5b806395d89b411161013a578063a1eda53c1161011f578063a1eda53c1461069f578063a217fddf146106c6578063a8fa343c146106ce57600080fd5b806395d89b41146106845780639dc29fac1461068c57600080fd5b80638da5cb5b1161016b5780638da5cb5b146105da5780638fd6a6ac146105e257806391d148541461061f57600080fd5b80638456cb591461057457806384ef8ffc1461057c57600080fd5b80632f2ff15d1161023f57806342966c68116101f3578063649a5ec7116101cd578063649a5ec7146104f957806370a082311461050c57806379cc67901461056157600080fd5b806342966c68146104a95780635c975abb146104bc578063634e93da146104e657600080fd5b806336568abe1161022457806336568abe1461047b5780633f4ba83a1461048e57806340c10f191461049657600080fd5b80632f2ff15d14610419578063313ce5671461042c57600080fd5b806318160ddd11610296578063248a9ca31161027b578063248a9ca31461039d578063282c51f3146103df5780632cd77a5a1461040657600080fd5b806318160ddd1461035957806323b872dd1461038a57600080fd5b806306fdde03116102c757806306fdde0314610327578063095ea7b31461033c5780630aa6220b1461034f57600080fd5b806301ffc9a7146102e3578063022d63fb1461030b575b600080fd5b6102f66102f13660046129ac565b610877565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff9091168152602001610302565b61032f6109f4565b60405161030291906129ee565b6102f661034a366004612a83565b610ac9565b610357610ae1565b005b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b604051908152602001610302565b6102f6610398366004612aad565b610af7565b61037c6103ab366004612aea565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b61037c7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b610357610414366004612bfe565b610b1b565b610357610427366004612ca2565b610db8565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005474010000000000000000000000000000000000000000900460ff1660405160ff9091168152602001610302565b610357610489366004612ca2565b610dfd565b610357610f64565b6103576104a4366004612a83565b610fae565b6103576104b7366004612aea565b6110b5565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102f6565b6103576104f4366004612cce565b6110e8565b610357610507366004612ce9565b6110fc565b61037c61051a366004612cce565b73ffffffffffffffffffffffffffffffffffffffff1660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052604090205490565b61035761056f366004612a83565b611110565b610357611144565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610302565b6105b56111a6565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489005473ffffffffffffffffffffffffffffffffffffffff166105b5565b6102f661062d366004612ca2565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61032f6111eb565b61035761069a366004612a83565b61123c565b6106a7611246565b6040805165ffffffffffff938416815292909116602083015201610302565b61037c600081565b6103576106dc366004612cce565b611305565b6102f66106ef366004612a83565b6113a7565b610357610702366004612cce565b6113b5565b610310611409565b6103576114ea565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610302565b61037c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6103576107b7366004612ca2565b611565565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489015461037c565b6103576115a6565b61037c6107f9366004612d11565b73ffffffffffffffffffffffffffffffffffffffff91821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b61037c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000148061090a57507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b8061095657507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b806109a257507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b806109ee57507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610a4590612d3b565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7190612d3b565b8015610abe5780601f10610a9357610100808354040283529160200191610abe565b820191906000526020600020905b815481529060010190602001808311610aa157829003601f168201915b505050505091505090565b600033610ad78185856115b9565b5060019392505050565b6000610aec816115c6565b610af46115d0565b50565b600033610b058582856115dd565b610b108585856116cb565b506001949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610b665750825b905060008267ffffffffffffffff166001148015610b835750303b155b905081158015610b91575080155b15610bc8576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610c295784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610c338b8b611776565b610c3b611788565b610c43611788565b60007fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489008054600182018b90557fffffffffffffffffffffff000000000000000000000000000000000000000000167401000000000000000000000000000000000000000060ff8d16027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff891617815590508715610d3d5788881115610d33576040517f25cc7152000000000000000000000000000000000000000000000000000000008152600481018990526024015b60405180910390fd5b610d3d8789611790565b610d486000886117ec565b50508315610dab5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b81610def576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610df982826118f6565b5050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840082158015610e6557507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8381169116145b15610f55577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610ed8575065ffffffffffff8116155b80610eeb57504265ffffffffffff821610155b15610f2c576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610d2a565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b610f5f838361193a565b505050565b6000610f6f816115c6565b610f77611993565b6040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a150565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6610fd8816115c6565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff079048901547fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff0790489009060006110467f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace025490565b9050811580159061105f57508161105d8683612dbd565b115b156110a35761106e8582612dbd565b6040517f25cc7152000000000000000000000000000000000000000000000000000000008152600401610d2a91815260200190565b6110ad8686611790565b505050505050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486110df816115c6565b610df982611a2a565b60006110f3816115c6565b610df982611a34565b6000611107816115c6565b610df982611ab4565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861113a816115c6565b610f5f8383611b24565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61116e816115c6565b611176611b39565b6040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610fa3565b60006111e67feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060917f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0091610a4590612d3b565b610df98282611110565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906112c957504265ffffffffffff831610155b6112d5576000806112fc565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6000611310816115c6565b7fc5ce4c6194754ec56151469c4af5ff17dd2a95dab96bf61ba95b3ff07904890080547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff858116918217845560405192169182907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a350505050565b600033610ad78185856116cb565b6113df7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610db8565b610af47f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610db8565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546000907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801580159061148c57504265ffffffffffff8216105b6114bd5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff166114e3565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984005473ffffffffffffffffffffffffffffffffffffffff1633811461155d576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610d2a565b610af4611bb2565b8161159c576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610df98282611ce3565b60006115b1816115c6565b610af4611d27565b610f5f8383836001611d32565b610af48133611d46565b6115db600080611ded565b565b73ffffffffffffffffffffffffffffffffffffffff83811660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146116c557818110156116b6576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610d2a565b6116c584848484036000611d32565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661171b576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b73ffffffffffffffffffffffffffffffffffffffff821661176b576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b610f5f838383611f86565b61177e611f99565b610df98282612000565b6115db611f99565b73ffffffffffffffffffffffffffffffffffffffff82166117e0576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b610df960008383611f86565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400836118e45760006118547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16146118a1576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85161790555b6118ee8484612063565b949350505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611930816115c6565b6116c583836117ec565b73ffffffffffffffffffffffffffffffffffffffff81163314611989576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f5f8282612184565b61199b612228565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610fa3565b610af43382612283565b6000611a3e611409565b611a47426122df565b611a519190612dd0565b9050611a5d828261232f565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b6000611abf826123ea565b611ac8426122df565b611ad29190612dd0565b9050611ade8282611ded565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611b2f8233836115dd565b610df98282612283565b611b41612439565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611a05565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400805473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611c2257504265ffffffffffff821610155b15611c63576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610d2a565b611cab6000611ca67feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff1690565b612184565b50611cb76000836117ec565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611d1d816115c6565b6116c58383612184565b6115db60008061232f565b611d3a612439565b6116c584848484612495565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610df9576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610d2a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611f00574265ffffffffffff82161015611ed6576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255611f00565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600101805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b611f8e612439565b610f5f838383612508565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166115db576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612008611f99565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036120548482612e35565b50600481016116c58382612e35565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1661217a5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556121163390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506109ee565b60009150506109ee565b60007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400831580156121ee57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984015473ffffffffffffffffffffffffffffffffffffffff8481169116145b1561221e576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6118ee848461257a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166115db576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166122d3576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b610df982600083611f86565b600065ffffffffffff82111561232b576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610d2a565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff88161717845591041680156116c5576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a150505050565b6000806123f5611409565b90508065ffffffffffff168365ffffffffffff161161241d576124188382612f4e565b612432565b61243265ffffffffffff841662069780612658565b9392505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156115db576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff8416036124fc576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610d2a565b6116c58484848461266e565b3073ffffffffffffffffffffffffffffffffffffffff83160361256f576040517f54e16bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610d2a565b610f5f8383836127db565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff161561217a5760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506109ee565b60008183106126675781612432565b5090919050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff85166126df576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b73ffffffffffffffffffffffffffffffffffffffff841661272f576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610d2a565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600183016020908152604080832093881683529290522083905581156127d4578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516127cb91815260200190565b60405180910390a35b5050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0073ffffffffffffffffffffffffffffffffffffffff8416612836578181600201600082825461282b9190612dbd565b909155506128e89050565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260208290526040902054828110156128bc576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810182905260448101849052606401610d2a565b73ffffffffffffffffffffffffffffffffffffffff851660009081526020839052604090209083900390555b73ffffffffffffffffffffffffffffffffffffffff831661291357600281018054839003905561293f565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020829052604090208054830190555b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161299e91815260200190565b60405180910390a350505050565b6000602082840312156129be57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461243257600080fd5b602081526000825180602084015260005b81811015612a1c57602081860181015160408684010152016129ff565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612a7e57600080fd5b919050565b60008060408385031215612a9657600080fd5b612a9f83612a5a565b946020939093013593505050565b600080600060608486031215612ac257600080fd5b612acb84612a5a565b9250612ad960208501612a5a565b929592945050506040919091013590565b600060208284031215612afc57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612b4357600080fd5b813567ffffffffffffffff811115612b5d57612b5d612b03565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612bc957612bc9612b03565b604052818152838201602001851015612be157600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215612c1757600080fd5b863567ffffffffffffffff811115612c2e57600080fd5b612c3a89828a01612b32565b965050602087013567ffffffffffffffff811115612c5757600080fd5b612c6389828a01612b32565b955050604087013560ff81168114612c7a57600080fd5b93506060870135925060808701359150612c9660a08801612a5a565b90509295509295509295565b60008060408385031215612cb557600080fd5b82359150612cc560208401612a5a565b90509250929050565b600060208284031215612ce057600080fd5b61243282612a5a565b600060208284031215612cfb57600080fd5b813565ffffffffffff8116811461243257600080fd5b60008060408385031215612d2457600080fd5b612d2d83612a5a565b9150612cc560208401612a5a565b600181811c90821680612d4f57607f821691505b602082108103612d88577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156109ee576109ee612d8e565b65ffffffffffff81811683821601908111156109ee576109ee612d8e565b601f821115610f5f57806000526020600020601f840160051c81016020851015612e155750805b601f840160051c820191505b818110156127d45760008155600101612e21565b815167ffffffffffffffff811115612e4f57612e4f612b03565b612e6381612e5d8454612d3b565b84612dee565b6020601f821160018114612eb55760008315612e7f5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556127d4565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015612f035787850151825560209485019460019092019101612ee3565b5084821015612f3f57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b65ffffffffffff82811682821603908111156109ee576109ee612d8e56fea164736f6c634300081a000a",
}

var BurnMintERC20PausableTransparentABI = BurnMintERC20PausableTransparentMetaData.ABI

var BurnMintERC20PausableTransparentBin = BurnMintERC20PausableTransparentMetaData.Bin

func DeployBurnMintERC20PausableTransparent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMintERC20PausableTransparent, error) {
	parsed, err := BurnMintERC20PausableTransparentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20PausableTransparentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20PausableTransparent{address: address, abi: *parsed, BurnMintERC20PausableTransparentCaller: BurnMintERC20PausableTransparentCaller{contract: contract}, BurnMintERC20PausableTransparentTransactor: BurnMintERC20PausableTransparentTransactor{contract: contract}, BurnMintERC20PausableTransparentFilterer: BurnMintERC20PausableTransparentFilterer{contract: contract}}, nil
}

type BurnMintERC20PausableTransparent struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20PausableTransparentCaller
	BurnMintERC20PausableTransparentTransactor
	BurnMintERC20PausableTransparentFilterer
}

type BurnMintERC20PausableTransparentCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableTransparentTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableTransparentFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20PausableTransparentSession struct {
	Contract     *BurnMintERC20PausableTransparent
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableTransparentCallerSession struct {
	Contract *BurnMintERC20PausableTransparentCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20PausableTransparentTransactorSession struct {
	Contract     *BurnMintERC20PausableTransparentTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20PausableTransparentRaw struct {
	Contract *BurnMintERC20PausableTransparent
}

type BurnMintERC20PausableTransparentCallerRaw struct {
	Contract *BurnMintERC20PausableTransparentCaller
}

type BurnMintERC20PausableTransparentTransactorRaw struct {
	Contract *BurnMintERC20PausableTransparentTransactor
}

func NewBurnMintERC20PausableTransparent(address common.Address, backend bind.ContractBackend) (*BurnMintERC20PausableTransparent, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20PausableTransparentABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20PausableTransparent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparent{address: address, abi: abi, BurnMintERC20PausableTransparentCaller: BurnMintERC20PausableTransparentCaller{contract: contract}, BurnMintERC20PausableTransparentTransactor: BurnMintERC20PausableTransparentTransactor{contract: contract}, BurnMintERC20PausableTransparentFilterer: BurnMintERC20PausableTransparentFilterer{contract: contract}}, nil
}

func NewBurnMintERC20PausableTransparentCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20PausableTransparentCaller, error) {
	contract, err := bindBurnMintERC20PausableTransparent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentCaller{contract: contract}, nil
}

func NewBurnMintERC20PausableTransparentTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20PausableTransparentTransactor, error) {
	contract, err := bindBurnMintERC20PausableTransparent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentTransactor{contract: contract}, nil
}

func NewBurnMintERC20PausableTransparentFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20PausableTransparentFilterer, error) {
	contract, err := bindBurnMintERC20PausableTransparent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentFilterer{contract: contract}, nil
}

func bindBurnMintERC20PausableTransparent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20PausableTransparentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableTransparent.Contract.BurnMintERC20PausableTransparentCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BurnMintERC20PausableTransparentTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BurnMintERC20PausableTransparentTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20PausableTransparent.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.BURNERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.BURNERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.DEFAULTADMINROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.MINTERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.MINTERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.PAUSERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.PAUSERROLE(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.Allowance(&_BurnMintERC20PausableTransparent.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.Allowance(&_BurnMintERC20PausableTransparent.CallOpts, owner, spender)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.BalanceOf(&_BurnMintERC20PausableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.BalanceOf(&_BurnMintERC20PausableTransparent.CallOpts, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableTransparent.Contract.Decimals(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20PausableTransparent.Contract.Decimals(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdminDelay(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdminDelay(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.GetCCIPAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.GetCCIPAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.GetRoleAdmin(&_BurnMintERC20PausableTransparent.CallOpts, role)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20PausableTransparent.Contract.GetRoleAdmin(&_BurnMintERC20PausableTransparent.CallOpts, role)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.HasRole(&_BurnMintERC20PausableTransparent.CallOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.HasRole(&_BurnMintERC20PausableTransparent.CallOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.MaxSupply(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.MaxSupply(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Name() (string, error) {
	return _BurnMintERC20PausableTransparent.Contract.Name(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Name() (string, error) {
	return _BurnMintERC20PausableTransparent.Contract.Name(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.Owner(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC20PausableTransparent.Contract.Owner(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Paused() (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.Paused(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Paused() (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.Paused(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableTransparent.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintERC20PausableTransparent.Contract.PendingDefaultAdmin(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableTransparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintERC20PausableTransparent.Contract.PendingDefaultAdminDelay(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.SupportsInterface(&_BurnMintERC20PausableTransparent.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20PausableTransparent.Contract.SupportsInterface(&_BurnMintERC20PausableTransparent.CallOpts, interfaceId)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Symbol() (string, error) {
	return _BurnMintERC20PausableTransparent.Contract.Symbol(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) Symbol() (string, error) {
	return _BurnMintERC20PausableTransparent.Contract.Symbol(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20PausableTransparent.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.TotalSupply(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20PausableTransparent.Contract.TotalSupply(&_BurnMintERC20PausableTransparent.CallOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.AcceptDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Approve(&_BurnMintERC20PausableTransparent.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Approve(&_BurnMintERC20PausableTransparent.TransactOpts, spender, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BeginDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Burn(&_BurnMintERC20PausableTransparent.TransactOpts, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Burn(&_BurnMintERC20PausableTransparent.TransactOpts, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Burn0(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Burn0(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BurnFrom(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.BurnFrom(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.CancelDefaultAdminTransfer(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableTransparent.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.ChangeDefaultAdminDelay(&_BurnMintERC20PausableTransparent.TransactOpts, newDelay)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableTransparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.GrantMintAndBurnRoles(&_BurnMintERC20PausableTransparent.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.GrantRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.GrantRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "initialize", name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Initialize(&_BurnMintERC20PausableTransparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Initialize(name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, defaultAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Initialize(&_BurnMintERC20PausableTransparent.TransactOpts, name, symbol, decimals_, maxSupply_, preMint, defaultAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Mint(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Mint(&_BurnMintERC20PausableTransparent.TransactOpts, account, amount)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "pause")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Pause(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Pause() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Pause(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RenounceRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RenounceRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RevokeRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RevokeRole(&_BurnMintERC20PausableTransparent.TransactOpts, role, account)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.RollbackDefaultAdminDelay(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.SetCCIPAdmin(&_BurnMintERC20PausableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.SetCCIPAdmin(&_BurnMintERC20PausableTransparent.TransactOpts, newAdmin)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Transfer(&_BurnMintERC20PausableTransparent.TransactOpts, to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Transfer(&_BurnMintERC20PausableTransparent.TransactOpts, to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.TransferFrom(&_BurnMintERC20PausableTransparent.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.TransferFrom(&_BurnMintERC20PausableTransparent.TransactOpts, from, to, value)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.contract.Transact(opts, "unpause")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Unpause(&_BurnMintERC20PausableTransparent.TransactOpts)
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentTransactorSession) Unpause() (*types.Transaction, error) {
	return _BurnMintERC20PausableTransparent.Contract.Unpause(&_BurnMintERC20PausableTransparent.TransactOpts)
}

type BurnMintERC20PausableTransparentApprovalIterator struct {
	Event *BurnMintERC20PausableTransparentApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentApproval)
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
		it.Event = new(BurnMintERC20PausableTransparentApproval)
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

func (it *BurnMintERC20PausableTransparentApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableTransparentApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentApprovalIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentApproval)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseApproval(log types.Log) (*BurnMintERC20PausableTransparentApproval, error) {
	event := new(BurnMintERC20PausableTransparentApproval)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20PausableTransparentCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20PausableTransparentCCIPAdminTransferred)
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

func (it *BurnMintERC20PausableTransparentCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableTransparentCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentCCIPAdminTransferredIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentCCIPAdminTransferred)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableTransparentCCIPAdminTransferred, error) {
	event := new(BurnMintERC20PausableTransparentCCIPAdminTransferred)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled)
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

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled)
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

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintERC20PausableTransparentDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentDefaultAdminTransferCanceled)
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
		it.Event = new(BurnMintERC20PausableTransparentDefaultAdminTransferCanceled)
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

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentDefaultAdminTransferCanceled)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminTransferCanceled, error) {
	event := new(BurnMintERC20PausableTransparentDefaultAdminTransferCanceled)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintERC20PausableTransparentDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentDefaultAdminTransferScheduled)
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
		it.Event = new(BurnMintERC20PausableTransparentDefaultAdminTransferScheduled)
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

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentDefaultAdminTransferScheduled)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminTransferScheduled, error) {
	event := new(BurnMintERC20PausableTransparentDefaultAdminTransferScheduled)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentInitializedIterator struct {
	Event *BurnMintERC20PausableTransparentInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentInitialized)
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
		it.Event = new(BurnMintERC20PausableTransparentInitialized)
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

func (it *BurnMintERC20PausableTransparentInitializedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentInitialized struct {
	Version uint64
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentInitializedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentInitializedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentInitialized)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseInitialized(log types.Log) (*BurnMintERC20PausableTransparentInitialized, error) {
	event := new(BurnMintERC20PausableTransparentInitialized)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentPausedIterator struct {
	Event *BurnMintERC20PausableTransparentPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentPaused)
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
		it.Event = new(BurnMintERC20PausableTransparentPaused)
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

func (it *BurnMintERC20PausableTransparentPausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentPausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentPausedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentPaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentPaused)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParsePaused(log types.Log) (*BurnMintERC20PausableTransparentPaused, error) {
	event := new(BurnMintERC20PausableTransparentPaused)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentRoleAdminChangedIterator struct {
	Event *BurnMintERC20PausableTransparentRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentRoleAdminChanged)
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
		it.Event = new(BurnMintERC20PausableTransparentRoleAdminChanged)
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

func (it *BurnMintERC20PausableTransparentRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableTransparentRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentRoleAdminChangedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentRoleAdminChanged)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableTransparentRoleAdminChanged, error) {
	event := new(BurnMintERC20PausableTransparentRoleAdminChanged)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentRoleGrantedIterator struct {
	Event *BurnMintERC20PausableTransparentRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentRoleGranted)
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
		it.Event = new(BurnMintERC20PausableTransparentRoleGranted)
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

func (it *BurnMintERC20PausableTransparentRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableTransparentRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentRoleGrantedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentRoleGranted)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20PausableTransparentRoleGranted, error) {
	event := new(BurnMintERC20PausableTransparentRoleGranted)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentRoleRevokedIterator struct {
	Event *BurnMintERC20PausableTransparentRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentRoleRevoked)
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
		it.Event = new(BurnMintERC20PausableTransparentRoleRevoked)
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

func (it *BurnMintERC20PausableTransparentRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableTransparentRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentRoleRevokedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentRoleRevoked)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableTransparentRoleRevoked, error) {
	event := new(BurnMintERC20PausableTransparentRoleRevoked)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentTransferIterator struct {
	Event *BurnMintERC20PausableTransparentTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentTransfer)
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
		it.Event = new(BurnMintERC20PausableTransparentTransfer)
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

func (it *BurnMintERC20PausableTransparentTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableTransparentTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentTransferIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentTransfer)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseTransfer(log types.Log) (*BurnMintERC20PausableTransparentTransfer, error) {
	event := new(BurnMintERC20PausableTransparentTransfer)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20PausableTransparentUnpausedIterator struct {
	Event *BurnMintERC20PausableTransparentUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20PausableTransparentUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20PausableTransparentUnpaused)
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
		it.Event = new(BurnMintERC20PausableTransparentUnpaused)
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

func (it *BurnMintERC20PausableTransparentUnpausedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20PausableTransparentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20PausableTransparentUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentUnpausedIterator, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20PausableTransparentUnpausedIterator{contract: _BurnMintERC20PausableTransparent.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentUnpaused) (event.Subscription, error) {

	logs, sub, err := _BurnMintERC20PausableTransparent.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20PausableTransparentUnpaused)
				if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparentFilterer) ParseUnpaused(log types.Log) (*BurnMintERC20PausableTransparentUnpaused, error) {
	event := new(BurnMintERC20PausableTransparentUnpaused)
	if err := _BurnMintERC20PausableTransparent.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparent) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20PausableTransparent.abi.Events["Approval"].ID:
		return _BurnMintERC20PausableTransparent.ParseApproval(log)
	case _BurnMintERC20PausableTransparent.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20PausableTransparent.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20PausableTransparent.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _BurnMintERC20PausableTransparent.ParseDefaultAdminDelayChangeCanceled(log)
	case _BurnMintERC20PausableTransparent.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _BurnMintERC20PausableTransparent.ParseDefaultAdminDelayChangeScheduled(log)
	case _BurnMintERC20PausableTransparent.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _BurnMintERC20PausableTransparent.ParseDefaultAdminTransferCanceled(log)
	case _BurnMintERC20PausableTransparent.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _BurnMintERC20PausableTransparent.ParseDefaultAdminTransferScheduled(log)
	case _BurnMintERC20PausableTransparent.abi.Events["Initialized"].ID:
		return _BurnMintERC20PausableTransparent.ParseInitialized(log)
	case _BurnMintERC20PausableTransparent.abi.Events["Paused"].ID:
		return _BurnMintERC20PausableTransparent.ParsePaused(log)
	case _BurnMintERC20PausableTransparent.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20PausableTransparent.ParseRoleAdminChanged(log)
	case _BurnMintERC20PausableTransparent.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20PausableTransparent.ParseRoleGranted(log)
	case _BurnMintERC20PausableTransparent.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20PausableTransparent.ParseRoleRevoked(log)
	case _BurnMintERC20PausableTransparent.abi.Events["Transfer"].ID:
		return _BurnMintERC20PausableTransparent.ParseTransfer(log)
	case _BurnMintERC20PausableTransparent.abi.Events["Unpaused"].ID:
		return _BurnMintERC20PausableTransparent.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC20PausableTransparentApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20PausableTransparentCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintERC20PausableTransparentDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintERC20PausableTransparentDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintERC20PausableTransparentInitialized) Topic() common.Hash {
	return common.HexToHash("0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2")
}

func (BurnMintERC20PausableTransparentPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (BurnMintERC20PausableTransparentRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20PausableTransparentRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20PausableTransparentRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20PausableTransparentTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (BurnMintERC20PausableTransparentUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_BurnMintERC20PausableTransparent *BurnMintERC20PausableTransparent) Address() common.Address {
	return _BurnMintERC20PausableTransparent.address
}

type BurnMintERC20PausableTransparentInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

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

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20PausableTransparentApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20PausableTransparentApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20PausableTransparentCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20PausableTransparentCCIPAdminTransferred, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintERC20PausableTransparentDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintERC20PausableTransparentDefaultAdminTransferScheduled, error)

	FilterInitialized(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*BurnMintERC20PausableTransparentInitialized, error)

	FilterPaused(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*BurnMintERC20PausableTransparentPaused, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20PausableTransparentRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20PausableTransparentRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableTransparentRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20PausableTransparentRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20PausableTransparentRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20PausableTransparentRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20PausableTransparentTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20PausableTransparentTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*BurnMintERC20PausableTransparentUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BurnMintERC20PausableTransparentUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*BurnMintERC20PausableTransparentUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
