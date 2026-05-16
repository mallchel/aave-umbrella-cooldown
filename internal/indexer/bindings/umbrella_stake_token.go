// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
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

// IERC4626StakeTokenCooldownSnapshot is an auto generated low-level Go binding around an user-defined struct.
type IERC4626StakeTokenCooldownSnapshot struct {
	Amount           *big.Int
	EndOfCooldown    uint32
	WithdrawalWindow uint32
}

// IERC4626StakeTokenSignatureParams is an auto generated low-level Go binding around an user-defined struct.
type IERC4626StakeTokenSignatureParams struct {
	V uint8
	R [32]byte
	S [32]byte
}

// UmbrellaStakeTokenMetaData contains all meta data concerning the UmbrellaStakeToken contract.
var UmbrellaStakeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRewardsController\",\"name\":\"rewardsController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"NotApprovedForCooldown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRescueGuardian\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmountSlashing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBalanceInStaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFundsAvailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCooldown\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCooldown\",\"type\":\"uint256\"}],\"name\":\"CooldownChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"CooldownOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Rescued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokensRescued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endOfCooldown\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeWindow\",\"type\":\"uint256\"}],\"name\":\"StakerCooldownUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldUnstakeWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newUnstakeWindow\",\"type\":\"uint256\"}],\"name\":\"UnstakeWindowChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_ASSETS_REMAINING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_CONTROLLER\",\"outputs\":[{\"internalType\":\"contractIRewardsController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"cooldownNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"cooldownOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIERC4626StakeToken.SignatureParams\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"cooldownWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIERC4626StakeToken.SignatureParams\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"depositWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyEtherTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSlashableAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStakerCooldown\",\"outputs\":[{\"components\":[{\"internalType\":\"uint192\",\"name\":\"amount\",\"type\":\"uint192\"},{\"internalType\":\"uint32\",\"name\":\"endOfCooldown\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalWindow\",\"type\":\"uint32\"}],\"internalType\":\"structIERC4626StakeToken.CooldownSnapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnstakeWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakedToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cooldown_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeWindow_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isCooldownOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Token\",\"type\":\"address\"}],\"name\":\"maxRescue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCooldown\",\"type\":\"uint256\"}],\"name\":\"setCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setCooldownOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnstakeWindow\",\"type\":\"uint256\"}],\"name\":\"setUnstakeWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whoCanRescue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UmbrellaStakeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use UmbrellaStakeTokenMetaData.ABI instead.
var UmbrellaStakeTokenABI = UmbrellaStakeTokenMetaData.ABI

// UmbrellaStakeToken is an auto generated Go binding around an Ethereum contract.
type UmbrellaStakeToken struct {
	UmbrellaStakeTokenCaller     // Read-only binding to the contract
	UmbrellaStakeTokenTransactor // Write-only binding to the contract
	UmbrellaStakeTokenFilterer   // Log filterer for contract events
}

// UmbrellaStakeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type UmbrellaStakeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmbrellaStakeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UmbrellaStakeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmbrellaStakeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UmbrellaStakeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UmbrellaStakeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UmbrellaStakeTokenSession struct {
	Contract     *UmbrellaStakeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UmbrellaStakeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UmbrellaStakeTokenCallerSession struct {
	Contract *UmbrellaStakeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UmbrellaStakeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UmbrellaStakeTokenTransactorSession struct {
	Contract     *UmbrellaStakeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UmbrellaStakeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type UmbrellaStakeTokenRaw struct {
	Contract *UmbrellaStakeToken // Generic contract binding to access the raw methods on
}

// UmbrellaStakeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UmbrellaStakeTokenCallerRaw struct {
	Contract *UmbrellaStakeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// UmbrellaStakeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UmbrellaStakeTokenTransactorRaw struct {
	Contract *UmbrellaStakeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUmbrellaStakeToken creates a new instance of UmbrellaStakeToken, bound to a specific deployed contract.
func NewUmbrellaStakeToken(address common.Address, backend bind.ContractBackend) (*UmbrellaStakeToken, error) {
	contract, err := bindUmbrellaStakeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeToken{UmbrellaStakeTokenCaller: UmbrellaStakeTokenCaller{contract: contract}, UmbrellaStakeTokenTransactor: UmbrellaStakeTokenTransactor{contract: contract}, UmbrellaStakeTokenFilterer: UmbrellaStakeTokenFilterer{contract: contract}}, nil
}

// NewUmbrellaStakeTokenCaller creates a new read-only instance of UmbrellaStakeToken, bound to a specific deployed contract.
func NewUmbrellaStakeTokenCaller(address common.Address, caller bind.ContractCaller) (*UmbrellaStakeTokenCaller, error) {
	contract, err := bindUmbrellaStakeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenCaller{contract: contract}, nil
}

// NewUmbrellaStakeTokenTransactor creates a new write-only instance of UmbrellaStakeToken, bound to a specific deployed contract.
func NewUmbrellaStakeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*UmbrellaStakeTokenTransactor, error) {
	contract, err := bindUmbrellaStakeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenTransactor{contract: contract}, nil
}

// NewUmbrellaStakeTokenFilterer creates a new log filterer instance of UmbrellaStakeToken, bound to a specific deployed contract.
func NewUmbrellaStakeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*UmbrellaStakeTokenFilterer, error) {
	contract, err := bindUmbrellaStakeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenFilterer{contract: contract}, nil
}

// bindUmbrellaStakeToken binds a generic wrapper to an already deployed contract.
func bindUmbrellaStakeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UmbrellaStakeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UmbrellaStakeToken *UmbrellaStakeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UmbrellaStakeToken.Contract.UmbrellaStakeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UmbrellaStakeToken *UmbrellaStakeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.UmbrellaStakeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UmbrellaStakeToken *UmbrellaStakeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.UmbrellaStakeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UmbrellaStakeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UmbrellaStakeToken.Contract.DOMAINSEPARATOR(&_UmbrellaStakeToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UmbrellaStakeToken.Contract.DOMAINSEPARATOR(&_UmbrellaStakeToken.CallOpts)
}

// MINASSETSREMAINING is a free data retrieval call binding the contract method 0x13187000.
//
// Solidity: function MIN_ASSETS_REMAINING() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MINASSETSREMAINING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "MIN_ASSETS_REMAINING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINASSETSREMAINING is a free data retrieval call binding the contract method 0x13187000.
//
// Solidity: function MIN_ASSETS_REMAINING() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MINASSETSREMAINING() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MINASSETSREMAINING(&_UmbrellaStakeToken.CallOpts)
}

// MINASSETSREMAINING is a free data retrieval call binding the contract method 0x13187000.
//
// Solidity: function MIN_ASSETS_REMAINING() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MINASSETSREMAINING() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MINASSETSREMAINING(&_UmbrellaStakeToken.CallOpts)
}

// REWARDSCONTROLLER is a free data retrieval call binding the contract method 0xcd086d45.
//
// Solidity: function REWARDS_CONTROLLER() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) REWARDSCONTROLLER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "REWARDS_CONTROLLER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSCONTROLLER is a free data retrieval call binding the contract method 0xcd086d45.
//
// Solidity: function REWARDS_CONTROLLER() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) REWARDSCONTROLLER() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.REWARDSCONTROLLER(&_UmbrellaStakeToken.CallOpts)
}

// REWARDSCONTROLLER is a free data retrieval call binding the contract method 0xcd086d45.
//
// Solidity: function REWARDS_CONTROLLER() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) REWARDSCONTROLLER() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.REWARDSCONTROLLER(&_UmbrellaStakeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.Allowance(&_UmbrellaStakeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.Allowance(&_UmbrellaStakeToken.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Asset() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.Asset(&_UmbrellaStakeToken.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Asset() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.Asset(&_UmbrellaStakeToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.BalanceOf(&_UmbrellaStakeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.BalanceOf(&_UmbrellaStakeToken.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.ConvertToAssets(&_UmbrellaStakeToken.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.ConvertToAssets(&_UmbrellaStakeToken.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.ConvertToShares(&_UmbrellaStakeToken.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.ConvertToShares(&_UmbrellaStakeToken.CallOpts, assets)
}

// CooldownNonces is a free data retrieval call binding the contract method 0xbc8fcdca.
//
// Solidity: function cooldownNonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) CooldownNonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "cooldownNonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownNonces is a free data retrieval call binding the contract method 0xbc8fcdca.
//
// Solidity: function cooldownNonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) CooldownNonces(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.CooldownNonces(&_UmbrellaStakeToken.CallOpts, owner)
}

// CooldownNonces is a free data retrieval call binding the contract method 0xbc8fcdca.
//
// Solidity: function cooldownNonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) CooldownNonces(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.CooldownNonces(&_UmbrellaStakeToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Decimals() (uint8, error) {
	return _UmbrellaStakeToken.Contract.Decimals(&_UmbrellaStakeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Decimals() (uint8, error) {
	return _UmbrellaStakeToken.Contract.Decimals(&_UmbrellaStakeToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _UmbrellaStakeToken.Contract.Eip712Domain(&_UmbrellaStakeToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _UmbrellaStakeToken.Contract.Eip712Domain(&_UmbrellaStakeToken.CallOpts)
}

// GetCooldown is a free data retrieval call binding the contract method 0x218e4a15.
//
// Solidity: function getCooldown() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) GetCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "getCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCooldown is a free data retrieval call binding the contract method 0x218e4a15.
//
// Solidity: function getCooldown() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) GetCooldown() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetCooldown(&_UmbrellaStakeToken.CallOpts)
}

// GetCooldown is a free data retrieval call binding the contract method 0x218e4a15.
//
// Solidity: function getCooldown() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) GetCooldown() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetCooldown(&_UmbrellaStakeToken.CallOpts)
}

// GetMaxSlashableAssets is a free data retrieval call binding the contract method 0x299f3966.
//
// Solidity: function getMaxSlashableAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) GetMaxSlashableAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "getMaxSlashableAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSlashableAssets is a free data retrieval call binding the contract method 0x299f3966.
//
// Solidity: function getMaxSlashableAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) GetMaxSlashableAssets() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetMaxSlashableAssets(&_UmbrellaStakeToken.CallOpts)
}

// GetMaxSlashableAssets is a free data retrieval call binding the contract method 0x299f3966.
//
// Solidity: function getMaxSlashableAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) GetMaxSlashableAssets() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetMaxSlashableAssets(&_UmbrellaStakeToken.CallOpts)
}

// GetStakerCooldown is a free data retrieval call binding the contract method 0x2f8a8f42.
//
// Solidity: function getStakerCooldown(address user) view returns((uint192,uint32,uint32))
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) GetStakerCooldown(opts *bind.CallOpts, user common.Address) (IERC4626StakeTokenCooldownSnapshot, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "getStakerCooldown", user)

	if err != nil {
		return *new(IERC4626StakeTokenCooldownSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC4626StakeTokenCooldownSnapshot)).(*IERC4626StakeTokenCooldownSnapshot)

	return out0, err

}

// GetStakerCooldown is a free data retrieval call binding the contract method 0x2f8a8f42.
//
// Solidity: function getStakerCooldown(address user) view returns((uint192,uint32,uint32))
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) GetStakerCooldown(user common.Address) (IERC4626StakeTokenCooldownSnapshot, error) {
	return _UmbrellaStakeToken.Contract.GetStakerCooldown(&_UmbrellaStakeToken.CallOpts, user)
}

// GetStakerCooldown is a free data retrieval call binding the contract method 0x2f8a8f42.
//
// Solidity: function getStakerCooldown(address user) view returns((uint192,uint32,uint32))
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) GetStakerCooldown(user common.Address) (IERC4626StakeTokenCooldownSnapshot, error) {
	return _UmbrellaStakeToken.Contract.GetStakerCooldown(&_UmbrellaStakeToken.CallOpts, user)
}

// GetUnstakeWindow is a free data retrieval call binding the contract method 0x90b9f9e4.
//
// Solidity: function getUnstakeWindow() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) GetUnstakeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "getUnstakeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnstakeWindow is a free data retrieval call binding the contract method 0x90b9f9e4.
//
// Solidity: function getUnstakeWindow() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) GetUnstakeWindow() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetUnstakeWindow(&_UmbrellaStakeToken.CallOpts)
}

// GetUnstakeWindow is a free data retrieval call binding the contract method 0x90b9f9e4.
//
// Solidity: function getUnstakeWindow() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) GetUnstakeWindow() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.GetUnstakeWindow(&_UmbrellaStakeToken.CallOpts)
}

// IsCooldownOperator is a free data retrieval call binding the contract method 0x2279c0c2.
//
// Solidity: function isCooldownOperator(address user, address operator) view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) IsCooldownOperator(opts *bind.CallOpts, user common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "isCooldownOperator", user, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCooldownOperator is a free data retrieval call binding the contract method 0x2279c0c2.
//
// Solidity: function isCooldownOperator(address user, address operator) view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) IsCooldownOperator(user common.Address, operator common.Address) (bool, error) {
	return _UmbrellaStakeToken.Contract.IsCooldownOperator(&_UmbrellaStakeToken.CallOpts, user, operator)
}

// IsCooldownOperator is a free data retrieval call binding the contract method 0x2279c0c2.
//
// Solidity: function isCooldownOperator(address user, address operator) view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) IsCooldownOperator(user common.Address, operator common.Address) (bool, error) {
	return _UmbrellaStakeToken.Contract.IsCooldownOperator(&_UmbrellaStakeToken.CallOpts, user, operator)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) LatestAnswer() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.LatestAnswer(&_UmbrellaStakeToken.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) LatestAnswer() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.LatestAnswer(&_UmbrellaStakeToken.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxDeposit(&_UmbrellaStakeToken.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxDeposit(&_UmbrellaStakeToken.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxMint(&_UmbrellaStakeToken.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxMint(&_UmbrellaStakeToken.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxRedeem(&_UmbrellaStakeToken.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxRedeem(&_UmbrellaStakeToken.CallOpts, owner)
}

// MaxRescue is a free data retrieval call binding the contract method 0xd7408715.
//
// Solidity: function maxRescue(address erc20Token) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MaxRescue(opts *bind.CallOpts, erc20Token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "maxRescue", erc20Token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRescue is a free data retrieval call binding the contract method 0xd7408715.
//
// Solidity: function maxRescue(address erc20Token) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MaxRescue(erc20Token common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxRescue(&_UmbrellaStakeToken.CallOpts, erc20Token)
}

// MaxRescue is a free data retrieval call binding the contract method 0xd7408715.
//
// Solidity: function maxRescue(address erc20Token) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MaxRescue(erc20Token common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxRescue(&_UmbrellaStakeToken.CallOpts, erc20Token)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxWithdraw(&_UmbrellaStakeToken.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.MaxWithdraw(&_UmbrellaStakeToken.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Name() (string, error) {
	return _UmbrellaStakeToken.Contract.Name(&_UmbrellaStakeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Name() (string, error) {
	return _UmbrellaStakeToken.Contract.Name(&_UmbrellaStakeToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.Nonces(&_UmbrellaStakeToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.Nonces(&_UmbrellaStakeToken.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Owner() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.Owner(&_UmbrellaStakeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Owner() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.Owner(&_UmbrellaStakeToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Paused() (bool, error) {
	return _UmbrellaStakeToken.Contract.Paused(&_UmbrellaStakeToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Paused() (bool, error) {
	return _UmbrellaStakeToken.Contract.Paused(&_UmbrellaStakeToken.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewDeposit(&_UmbrellaStakeToken.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewDeposit(&_UmbrellaStakeToken.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewMint(&_UmbrellaStakeToken.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewMint(&_UmbrellaStakeToken.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewRedeem(&_UmbrellaStakeToken.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewRedeem(&_UmbrellaStakeToken.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewWithdraw(&_UmbrellaStakeToken.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.PreviewWithdraw(&_UmbrellaStakeToken.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Symbol() (string, error) {
	return _UmbrellaStakeToken.Contract.Symbol(&_UmbrellaStakeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) Symbol() (string, error) {
	return _UmbrellaStakeToken.Contract.Symbol(&_UmbrellaStakeToken.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) TotalAssets() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.TotalAssets(&_UmbrellaStakeToken.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) TotalAssets() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.TotalAssets(&_UmbrellaStakeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) TotalSupply() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.TotalSupply(&_UmbrellaStakeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _UmbrellaStakeToken.Contract.TotalSupply(&_UmbrellaStakeToken.CallOpts)
}

// WhoCanRescue is a free data retrieval call binding the contract method 0xa4757b0f.
//
// Solidity: function whoCanRescue() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCaller) WhoCanRescue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UmbrellaStakeToken.contract.Call(opts, &out, "whoCanRescue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhoCanRescue is a free data retrieval call binding the contract method 0xa4757b0f.
//
// Solidity: function whoCanRescue() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) WhoCanRescue() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.WhoCanRescue(&_UmbrellaStakeToken.CallOpts)
}

// WhoCanRescue is a free data retrieval call binding the contract method 0xa4757b0f.
//
// Solidity: function whoCanRescue() view returns(address)
func (_UmbrellaStakeToken *UmbrellaStakeTokenCallerSession) WhoCanRescue() (common.Address, error) {
	return _UmbrellaStakeToken.Contract.WhoCanRescue(&_UmbrellaStakeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Approve(&_UmbrellaStakeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Approve(&_UmbrellaStakeToken.TransactOpts, spender, value)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Cooldown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "cooldown")
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Cooldown() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Cooldown(&_UmbrellaStakeToken.TransactOpts)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Cooldown() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Cooldown(&_UmbrellaStakeToken.TransactOpts)
}

// CooldownOnBehalfOf is a paid mutator transaction binding the contract method 0x250201db.
//
// Solidity: function cooldownOnBehalfOf(address owner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) CooldownOnBehalfOf(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "cooldownOnBehalfOf", owner)
}

// CooldownOnBehalfOf is a paid mutator transaction binding the contract method 0x250201db.
//
// Solidity: function cooldownOnBehalfOf(address owner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) CooldownOnBehalfOf(owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.CooldownOnBehalfOf(&_UmbrellaStakeToken.TransactOpts, owner)
}

// CooldownOnBehalfOf is a paid mutator transaction binding the contract method 0x250201db.
//
// Solidity: function cooldownOnBehalfOf(address owner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) CooldownOnBehalfOf(owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.CooldownOnBehalfOf(&_UmbrellaStakeToken.TransactOpts, owner)
}

// CooldownWithPermit is a paid mutator transaction binding the contract method 0xfae9fed0.
//
// Solidity: function cooldownWithPermit(address user, uint256 deadline, (uint8,bytes32,bytes32) sig) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) CooldownWithPermit(opts *bind.TransactOpts, user common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "cooldownWithPermit", user, deadline, sig)
}

// CooldownWithPermit is a paid mutator transaction binding the contract method 0xfae9fed0.
//
// Solidity: function cooldownWithPermit(address user, uint256 deadline, (uint8,bytes32,bytes32) sig) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) CooldownWithPermit(user common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.CooldownWithPermit(&_UmbrellaStakeToken.TransactOpts, user, deadline, sig)
}

// CooldownWithPermit is a paid mutator transaction binding the contract method 0xfae9fed0.
//
// Solidity: function cooldownWithPermit(address user, uint256 deadline, (uint8,bytes32,bytes32) sig) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) CooldownWithPermit(user common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.CooldownWithPermit(&_UmbrellaStakeToken.TransactOpts, user, deadline, sig)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Deposit(&_UmbrellaStakeToken.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Deposit(&_UmbrellaStakeToken.TransactOpts, assets, receiver)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x03fc9747.
//
// Solidity: function depositWithPermit(uint256 assets, address receiver, uint256 deadline, (uint8,bytes32,bytes32) sig) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) DepositWithPermit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "depositWithPermit", assets, receiver, deadline, sig)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x03fc9747.
//
// Solidity: function depositWithPermit(uint256 assets, address receiver, uint256 deadline, (uint8,bytes32,bytes32) sig) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) DepositWithPermit(assets *big.Int, receiver common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.DepositWithPermit(&_UmbrellaStakeToken.TransactOpts, assets, receiver, deadline, sig)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x03fc9747.
//
// Solidity: function depositWithPermit(uint256 assets, address receiver, uint256 deadline, (uint8,bytes32,bytes32) sig) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) DepositWithPermit(assets *big.Int, receiver common.Address, deadline *big.Int, sig IERC4626StakeTokenSignatureParams) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.DepositWithPermit(&_UmbrellaStakeToken.TransactOpts, assets, receiver, deadline, sig)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) EmergencyEtherTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "emergencyEtherTransfer", to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.EmergencyEtherTransfer(&_UmbrellaStakeToken.TransactOpts, to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.EmergencyEtherTransfer(&_UmbrellaStakeToken.TransactOpts, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address erc20Token, address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) EmergencyTokenTransfer(opts *bind.TransactOpts, erc20Token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "emergencyTokenTransfer", erc20Token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address erc20Token, address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) EmergencyTokenTransfer(erc20Token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.EmergencyTokenTransfer(&_UmbrellaStakeToken.TransactOpts, erc20Token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address erc20Token, address to, uint256 amount) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) EmergencyTokenTransfer(erc20Token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.EmergencyTokenTransfer(&_UmbrellaStakeToken.TransactOpts, erc20Token, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0bc95a7.
//
// Solidity: function initialize(address stakedToken, string name, string symbol, address owner, uint256 cooldown_, uint256 unstakeWindow_) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Initialize(opts *bind.TransactOpts, stakedToken common.Address, name string, symbol string, owner common.Address, cooldown_ *big.Int, unstakeWindow_ *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "initialize", stakedToken, name, symbol, owner, cooldown_, unstakeWindow_)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0bc95a7.
//
// Solidity: function initialize(address stakedToken, string name, string symbol, address owner, uint256 cooldown_, uint256 unstakeWindow_) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Initialize(stakedToken common.Address, name string, symbol string, owner common.Address, cooldown_ *big.Int, unstakeWindow_ *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Initialize(&_UmbrellaStakeToken.TransactOpts, stakedToken, name, symbol, owner, cooldown_, unstakeWindow_)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0bc95a7.
//
// Solidity: function initialize(address stakedToken, string name, string symbol, address owner, uint256 cooldown_, uint256 unstakeWindow_) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Initialize(stakedToken common.Address, name string, symbol string, owner common.Address, cooldown_ *big.Int, unstakeWindow_ *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Initialize(&_UmbrellaStakeToken.TransactOpts, stakedToken, name, symbol, owner, cooldown_, unstakeWindow_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Mint(&_UmbrellaStakeToken.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Mint(&_UmbrellaStakeToken.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Pause() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Pause(&_UmbrellaStakeToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Pause(&_UmbrellaStakeToken.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Permit(&_UmbrellaStakeToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Permit(&_UmbrellaStakeToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Redeem(&_UmbrellaStakeToken.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Redeem(&_UmbrellaStakeToken.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.RenounceOwnership(&_UmbrellaStakeToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.RenounceOwnership(&_UmbrellaStakeToken.TransactOpts)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) SetCooldown(opts *bind.TransactOpts, newCooldown *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "setCooldown", newCooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) SetCooldown(newCooldown *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetCooldown(&_UmbrellaStakeToken.TransactOpts, newCooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) SetCooldown(newCooldown *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetCooldown(&_UmbrellaStakeToken.TransactOpts, newCooldown)
}

// SetCooldownOperator is a paid mutator transaction binding the contract method 0xda27578f.
//
// Solidity: function setCooldownOperator(address operator, bool flag) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) SetCooldownOperator(opts *bind.TransactOpts, operator common.Address, flag bool) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "setCooldownOperator", operator, flag)
}

// SetCooldownOperator is a paid mutator transaction binding the contract method 0xda27578f.
//
// Solidity: function setCooldownOperator(address operator, bool flag) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) SetCooldownOperator(operator common.Address, flag bool) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetCooldownOperator(&_UmbrellaStakeToken.TransactOpts, operator, flag)
}

// SetCooldownOperator is a paid mutator transaction binding the contract method 0xda27578f.
//
// Solidity: function setCooldownOperator(address operator, bool flag) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) SetCooldownOperator(operator common.Address, flag bool) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetCooldownOperator(&_UmbrellaStakeToken.TransactOpts, operator, flag)
}

// SetUnstakeWindow is a paid mutator transaction binding the contract method 0xf8f10dfc.
//
// Solidity: function setUnstakeWindow(uint256 newUnstakeWindow) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) SetUnstakeWindow(opts *bind.TransactOpts, newUnstakeWindow *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "setUnstakeWindow", newUnstakeWindow)
}

// SetUnstakeWindow is a paid mutator transaction binding the contract method 0xf8f10dfc.
//
// Solidity: function setUnstakeWindow(uint256 newUnstakeWindow) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) SetUnstakeWindow(newUnstakeWindow *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetUnstakeWindow(&_UmbrellaStakeToken.TransactOpts, newUnstakeWindow)
}

// SetUnstakeWindow is a paid mutator transaction binding the contract method 0xf8f10dfc.
//
// Solidity: function setUnstakeWindow(uint256 newUnstakeWindow) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) SetUnstakeWindow(newUnstakeWindow *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.SetUnstakeWindow(&_UmbrellaStakeToken.TransactOpts, newUnstakeWindow)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address destination, uint256 amount) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Slash(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "slash", destination, amount)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address destination, uint256 amount) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Slash(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Slash(&_UmbrellaStakeToken.TransactOpts, destination, amount)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address destination, uint256 amount) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Slash(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Slash(&_UmbrellaStakeToken.TransactOpts, destination, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Transfer(&_UmbrellaStakeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Transfer(&_UmbrellaStakeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.TransferFrom(&_UmbrellaStakeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.TransferFrom(&_UmbrellaStakeToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.TransferOwnership(&_UmbrellaStakeToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.TransferOwnership(&_UmbrellaStakeToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Unpause() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Unpause(&_UmbrellaStakeToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Unpause(&_UmbrellaStakeToken.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Withdraw(&_UmbrellaStakeToken.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UmbrellaStakeToken *UmbrellaStakeTokenTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UmbrellaStakeToken.Contract.Withdraw(&_UmbrellaStakeToken.TransactOpts, assets, receiver, owner)
}

// UmbrellaStakeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenApprovalIterator struct {
	Event *UmbrellaStakeTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenApproval)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenApproval)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenApproval represents a Approval event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UmbrellaStakeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenApprovalIterator{contract: _UmbrellaStakeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenApproval)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseApproval(log types.Log) (*UmbrellaStakeTokenApproval, error) {
	event := new(UmbrellaStakeTokenApproval)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenCooldownChangedIterator is returned from FilterCooldownChanged and is used to iterate over the raw logs and unpacked data for CooldownChanged events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenCooldownChangedIterator struct {
	Event *UmbrellaStakeTokenCooldownChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenCooldownChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenCooldownChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenCooldownChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenCooldownChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenCooldownChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenCooldownChanged represents a CooldownChanged event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenCooldownChanged struct {
	OldCooldown *big.Int
	NewCooldown *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCooldownChanged is a free log retrieval operation binding the contract event 0x0731af75921ee6c66096a5c95daa1adcf95ff01e0ce8063a2369cb218ee4bcc9.
//
// Solidity: event CooldownChanged(uint256 oldCooldown, uint256 newCooldown)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterCooldownChanged(opts *bind.FilterOpts) (*UmbrellaStakeTokenCooldownChangedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "CooldownChanged")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenCooldownChangedIterator{contract: _UmbrellaStakeToken.contract, event: "CooldownChanged", logs: logs, sub: sub}, nil
}

// WatchCooldownChanged is a free log subscription operation binding the contract event 0x0731af75921ee6c66096a5c95daa1adcf95ff01e0ce8063a2369cb218ee4bcc9.
//
// Solidity: event CooldownChanged(uint256 oldCooldown, uint256 newCooldown)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchCooldownChanged(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenCooldownChanged) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "CooldownChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenCooldownChanged)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "CooldownChanged", log); err != nil {
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

// ParseCooldownChanged is a log parse operation binding the contract event 0x0731af75921ee6c66096a5c95daa1adcf95ff01e0ce8063a2369cb218ee4bcc9.
//
// Solidity: event CooldownChanged(uint256 oldCooldown, uint256 newCooldown)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseCooldownChanged(log types.Log) (*UmbrellaStakeTokenCooldownChanged, error) {
	event := new(UmbrellaStakeTokenCooldownChanged)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "CooldownChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenCooldownOperatorSetIterator is returned from FilterCooldownOperatorSet and is used to iterate over the raw logs and unpacked data for CooldownOperatorSet events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenCooldownOperatorSetIterator struct {
	Event *UmbrellaStakeTokenCooldownOperatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenCooldownOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenCooldownOperatorSet)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenCooldownOperatorSet)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenCooldownOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenCooldownOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenCooldownOperatorSet represents a CooldownOperatorSet event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenCooldownOperatorSet struct {
	User     common.Address
	Operator common.Address
	Flag     bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCooldownOperatorSet is a free log retrieval operation binding the contract event 0xa38c17a90fc13db5a035877139d3649d9c7f05910b0a71aedd38e8fc97e8ec0b.
//
// Solidity: event CooldownOperatorSet(address indexed user, address indexed operator, bool flag)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterCooldownOperatorSet(opts *bind.FilterOpts, user []common.Address, operator []common.Address) (*UmbrellaStakeTokenCooldownOperatorSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "CooldownOperatorSet", userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenCooldownOperatorSetIterator{contract: _UmbrellaStakeToken.contract, event: "CooldownOperatorSet", logs: logs, sub: sub}, nil
}

// WatchCooldownOperatorSet is a free log subscription operation binding the contract event 0xa38c17a90fc13db5a035877139d3649d9c7f05910b0a71aedd38e8fc97e8ec0b.
//
// Solidity: event CooldownOperatorSet(address indexed user, address indexed operator, bool flag)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchCooldownOperatorSet(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenCooldownOperatorSet, user []common.Address, operator []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "CooldownOperatorSet", userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenCooldownOperatorSet)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "CooldownOperatorSet", log); err != nil {
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

// ParseCooldownOperatorSet is a log parse operation binding the contract event 0xa38c17a90fc13db5a035877139d3649d9c7f05910b0a71aedd38e8fc97e8ec0b.
//
// Solidity: event CooldownOperatorSet(address indexed user, address indexed operator, bool flag)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseCooldownOperatorSet(log types.Log) (*UmbrellaStakeTokenCooldownOperatorSet, error) {
	event := new(UmbrellaStakeTokenCooldownOperatorSet)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "CooldownOperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenDepositIterator struct {
	Event *UmbrellaStakeTokenDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenDeposit)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenDeposit)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenDeposit represents a Deposit event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*UmbrellaStakeTokenDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenDepositIterator{contract: _UmbrellaStakeToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenDeposit)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseDeposit(log types.Log) (*UmbrellaStakeTokenDeposit, error) {
	event := new(UmbrellaStakeTokenDeposit)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenEIP712DomainChangedIterator struct {
	Event *UmbrellaStakeTokenEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenEIP712DomainChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenEIP712DomainChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*UmbrellaStakeTokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenEIP712DomainChangedIterator{contract: _UmbrellaStakeToken.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenEIP712DomainChanged)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseEIP712DomainChanged(log types.Log) (*UmbrellaStakeTokenEIP712DomainChanged, error) {
	event := new(UmbrellaStakeTokenEIP712DomainChanged)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenERC20RescuedIterator is returned from FilterERC20Rescued and is used to iterate over the raw logs and unpacked data for ERC20Rescued events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenERC20RescuedIterator struct {
	Event *UmbrellaStakeTokenERC20Rescued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenERC20RescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenERC20Rescued)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenERC20Rescued)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenERC20RescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenERC20RescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenERC20Rescued represents a ERC20Rescued event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenERC20Rescued struct {
	Caller common.Address
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20Rescued is a free log retrieval operation binding the contract event 0xc7af665d489507e14ae25ac7ab0030fc7f570869610bdd32117ea56b60ae5c61.
//
// Solidity: event ERC20Rescued(address indexed caller, address indexed token, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterERC20Rescued(opts *bind.FilterOpts, caller []common.Address, token []common.Address, to []common.Address) (*UmbrellaStakeTokenERC20RescuedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "ERC20Rescued", callerRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenERC20RescuedIterator{contract: _UmbrellaStakeToken.contract, event: "ERC20Rescued", logs: logs, sub: sub}, nil
}

// WatchERC20Rescued is a free log subscription operation binding the contract event 0xc7af665d489507e14ae25ac7ab0030fc7f570869610bdd32117ea56b60ae5c61.
//
// Solidity: event ERC20Rescued(address indexed caller, address indexed token, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchERC20Rescued(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenERC20Rescued, caller []common.Address, token []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "ERC20Rescued", callerRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenERC20Rescued)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "ERC20Rescued", log); err != nil {
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

// ParseERC20Rescued is a log parse operation binding the contract event 0xc7af665d489507e14ae25ac7ab0030fc7f570869610bdd32117ea56b60ae5c61.
//
// Solidity: event ERC20Rescued(address indexed caller, address indexed token, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseERC20Rescued(log types.Log) (*UmbrellaStakeTokenERC20Rescued, error) {
	event := new(UmbrellaStakeTokenERC20Rescued)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "ERC20Rescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenInitializedIterator struct {
	Event *UmbrellaStakeTokenInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenInitialized)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenInitialized)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenInitialized represents a Initialized event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*UmbrellaStakeTokenInitializedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenInitializedIterator{contract: _UmbrellaStakeToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenInitialized)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseInitialized(log types.Log) (*UmbrellaStakeTokenInitialized, error) {
	event := new(UmbrellaStakeTokenInitialized)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenNativeTokensRescuedIterator is returned from FilterNativeTokensRescued and is used to iterate over the raw logs and unpacked data for NativeTokensRescued events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenNativeTokensRescuedIterator struct {
	Event *UmbrellaStakeTokenNativeTokensRescued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenNativeTokensRescuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenNativeTokensRescued)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenNativeTokensRescued)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenNativeTokensRescuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenNativeTokensRescuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenNativeTokensRescued represents a NativeTokensRescued event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenNativeTokensRescued struct {
	Caller common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeTokensRescued is a free log retrieval operation binding the contract event 0xb7c602059992183c7b767c08204223afc99f1895fd175adf9ece23ce9f5bb8b7.
//
// Solidity: event NativeTokensRescued(address indexed caller, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterNativeTokensRescued(opts *bind.FilterOpts, caller []common.Address, to []common.Address) (*UmbrellaStakeTokenNativeTokensRescuedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "NativeTokensRescued", callerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenNativeTokensRescuedIterator{contract: _UmbrellaStakeToken.contract, event: "NativeTokensRescued", logs: logs, sub: sub}, nil
}

// WatchNativeTokensRescued is a free log subscription operation binding the contract event 0xb7c602059992183c7b767c08204223afc99f1895fd175adf9ece23ce9f5bb8b7.
//
// Solidity: event NativeTokensRescued(address indexed caller, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchNativeTokensRescued(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenNativeTokensRescued, caller []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "NativeTokensRescued", callerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenNativeTokensRescued)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "NativeTokensRescued", log); err != nil {
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

// ParseNativeTokensRescued is a log parse operation binding the contract event 0xb7c602059992183c7b767c08204223afc99f1895fd175adf9ece23ce9f5bb8b7.
//
// Solidity: event NativeTokensRescued(address indexed caller, address indexed to, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseNativeTokensRescued(log types.Log) (*UmbrellaStakeTokenNativeTokensRescued, error) {
	event := new(UmbrellaStakeTokenNativeTokensRescued)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "NativeTokensRescued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenOwnershipTransferredIterator struct {
	Event *UmbrellaStakeTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenOwnershipTransferred)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenOwnershipTransferred)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenOwnershipTransferred represents a OwnershipTransferred event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UmbrellaStakeTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenOwnershipTransferredIterator{contract: _UmbrellaStakeToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenOwnershipTransferred)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseOwnershipTransferred(log types.Log) (*UmbrellaStakeTokenOwnershipTransferred, error) {
	event := new(UmbrellaStakeTokenOwnershipTransferred)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenPausedIterator struct {
	Event *UmbrellaStakeTokenPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenPaused)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenPaused)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenPaused represents a Paused event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*UmbrellaStakeTokenPausedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenPausedIterator{contract: _UmbrellaStakeToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenPaused) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenPaused)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParsePaused(log types.Log) (*UmbrellaStakeTokenPaused, error) {
	event := new(UmbrellaStakeTokenPaused)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenSlashedIterator struct {
	Event *UmbrellaStakeTokenSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenSlashed)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenSlashed)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenSlashed represents a Slashed event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenSlashed struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed destination, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterSlashed(opts *bind.FilterOpts, destination []common.Address) (*UmbrellaStakeTokenSlashedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Slashed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenSlashedIterator{contract: _UmbrellaStakeToken.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed destination, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenSlashed, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Slashed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenSlashed)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed destination, uint256 amount)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseSlashed(log types.Log) (*UmbrellaStakeTokenSlashed, error) {
	event := new(UmbrellaStakeTokenSlashed)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenStakerCooldownUpdatedIterator is returned from FilterStakerCooldownUpdated and is used to iterate over the raw logs and unpacked data for StakerCooldownUpdated events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenStakerCooldownUpdatedIterator struct {
	Event *UmbrellaStakeTokenStakerCooldownUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenStakerCooldownUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenStakerCooldownUpdated)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenStakerCooldownUpdated)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenStakerCooldownUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenStakerCooldownUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenStakerCooldownUpdated represents a StakerCooldownUpdated event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenStakerCooldownUpdated struct {
	User          common.Address
	Amount        *big.Int
	EndOfCooldown *big.Int
	UnstakeWindow *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakerCooldownUpdated is a free log retrieval operation binding the contract event 0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419.
//
// Solidity: event StakerCooldownUpdated(address indexed user, uint256 amount, uint256 endOfCooldown, uint256 unstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterStakerCooldownUpdated(opts *bind.FilterOpts, user []common.Address) (*UmbrellaStakeTokenStakerCooldownUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "StakerCooldownUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenStakerCooldownUpdatedIterator{contract: _UmbrellaStakeToken.contract, event: "StakerCooldownUpdated", logs: logs, sub: sub}, nil
}

// WatchStakerCooldownUpdated is a free log subscription operation binding the contract event 0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419.
//
// Solidity: event StakerCooldownUpdated(address indexed user, uint256 amount, uint256 endOfCooldown, uint256 unstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchStakerCooldownUpdated(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenStakerCooldownUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "StakerCooldownUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenStakerCooldownUpdated)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "StakerCooldownUpdated", log); err != nil {
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

// ParseStakerCooldownUpdated is a log parse operation binding the contract event 0xddc8760931d97309f92a4266c6046f83387e6407bcd727e7dd2130bfc430c419.
//
// Solidity: event StakerCooldownUpdated(address indexed user, uint256 amount, uint256 endOfCooldown, uint256 unstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseStakerCooldownUpdated(log types.Log) (*UmbrellaStakeTokenStakerCooldownUpdated, error) {
	event := new(UmbrellaStakeTokenStakerCooldownUpdated)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "StakerCooldownUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenTransferIterator struct {
	Event *UmbrellaStakeTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenTransfer)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenTransfer)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenTransfer represents a Transfer event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UmbrellaStakeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenTransferIterator{contract: _UmbrellaStakeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenTransfer)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseTransfer(log types.Log) (*UmbrellaStakeTokenTransfer, error) {
	event := new(UmbrellaStakeTokenTransfer)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenUnpausedIterator struct {
	Event *UmbrellaStakeTokenUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenUnpaused)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenUnpaused)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenUnpaused represents a Unpaused event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UmbrellaStakeTokenUnpausedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenUnpausedIterator{contract: _UmbrellaStakeToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenUnpaused)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseUnpaused(log types.Log) (*UmbrellaStakeTokenUnpaused, error) {
	event := new(UmbrellaStakeTokenUnpaused)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenUnstakeWindowChangedIterator is returned from FilterUnstakeWindowChanged and is used to iterate over the raw logs and unpacked data for UnstakeWindowChanged events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenUnstakeWindowChangedIterator struct {
	Event *UmbrellaStakeTokenUnstakeWindowChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenUnstakeWindowChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenUnstakeWindowChanged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenUnstakeWindowChanged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenUnstakeWindowChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenUnstakeWindowChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenUnstakeWindowChanged represents a UnstakeWindowChanged event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenUnstakeWindowChanged struct {
	OldUnstakeWindow *big.Int
	NewUnstakeWindow *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnstakeWindowChanged is a free log retrieval operation binding the contract event 0x6fca801becb9707cbca62182fa0b26a34d43b1a631a501b6c1ac5ae2232a70e9.
//
// Solidity: event UnstakeWindowChanged(uint256 oldUnstakeWindow, uint256 newUnstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterUnstakeWindowChanged(opts *bind.FilterOpts) (*UmbrellaStakeTokenUnstakeWindowChangedIterator, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "UnstakeWindowChanged")
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenUnstakeWindowChangedIterator{contract: _UmbrellaStakeToken.contract, event: "UnstakeWindowChanged", logs: logs, sub: sub}, nil
}

// WatchUnstakeWindowChanged is a free log subscription operation binding the contract event 0x6fca801becb9707cbca62182fa0b26a34d43b1a631a501b6c1ac5ae2232a70e9.
//
// Solidity: event UnstakeWindowChanged(uint256 oldUnstakeWindow, uint256 newUnstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchUnstakeWindowChanged(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenUnstakeWindowChanged) (event.Subscription, error) {

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "UnstakeWindowChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenUnstakeWindowChanged)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "UnstakeWindowChanged", log); err != nil {
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

// ParseUnstakeWindowChanged is a log parse operation binding the contract event 0x6fca801becb9707cbca62182fa0b26a34d43b1a631a501b6c1ac5ae2232a70e9.
//
// Solidity: event UnstakeWindowChanged(uint256 oldUnstakeWindow, uint256 newUnstakeWindow)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseUnstakeWindowChanged(log types.Log) (*UmbrellaStakeTokenUnstakeWindowChanged, error) {
	event := new(UmbrellaStakeTokenUnstakeWindowChanged)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "UnstakeWindowChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UmbrellaStakeTokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenWithdrawIterator struct {
	Event *UmbrellaStakeTokenWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UmbrellaStakeTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UmbrellaStakeTokenWithdraw)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UmbrellaStakeTokenWithdraw)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UmbrellaStakeTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UmbrellaStakeTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UmbrellaStakeTokenWithdraw represents a Withdraw event raised by the UmbrellaStakeToken contract.
type UmbrellaStakeTokenWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*UmbrellaStakeTokenWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &UmbrellaStakeTokenWithdrawIterator{contract: _UmbrellaStakeToken.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *UmbrellaStakeTokenWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UmbrellaStakeToken.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UmbrellaStakeTokenWithdraw)
				if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UmbrellaStakeToken *UmbrellaStakeTokenFilterer) ParseWithdraw(log types.Log) (*UmbrellaStakeTokenWithdraw, error) {
	event := new(UmbrellaStakeTokenWithdraw)
	if err := _UmbrellaStakeToken.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
