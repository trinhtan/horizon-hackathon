// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeBank

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeBankABI is the input ABI used to generate the binding from.
const BridgeBankABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_harmonyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountHarmonyToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"LogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogUnlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeDenominator\",\"type\":\"uint256\"}],\"name\":\"UpdateFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newHarmonyBridge\",\"type\":\"address\"}],\"name\":\"UpdateHarmonyBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"UpdateOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONEAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SAFE_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"activateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_harmonyToken\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bandOracleInterface\",\"outputs\":[{\"internalType\":\"contractBandOracleInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"deactivateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getLockedFund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getTokenMappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"getTotalERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalETHBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"harmonyBridge\",\"outputs\":[{\"internalType\":\"contractHarmonyBridge\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_harmonyBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bandOracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lendingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_harmonyWETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"}],\"name\":\"isActiveToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"swapETHForONE\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destTokenSymbol\",\"type\":\"string\"}],\"name\":\"swapETHForToken\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"swapETHForWETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"}],\"name\":\"swapTokenForONE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destTokenSymbol\",\"type\":\"string\"}],\"name\":\"swapTokenForToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"}],\"name\":\"swapTokenForWETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"}],\"name\":\"swapToken_1_1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"harmonyMappedToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethereumTokenAmount\",\"type\":\"uint256\"}],\"name\":\"unlockERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"unlockETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeDenominator\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_harmonyBridge\",\"type\":\"address\"}],\"name\":\"updateHmyBridge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"name\":\"updateOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"contractIWETHGateway\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethereumToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountEthereumToken\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBank is an auto generated Go binding around an Ethereum contract.
type BridgeBank struct {
	BridgeBankCaller     // Read-only binding to the contract
	BridgeBankTransactor // Write-only binding to the contract
	BridgeBankFilterer   // Log filterer for contract events
}

// BridgeBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBankSession struct {
	Contract     *BridgeBank       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBankCallerSession struct {
	Contract *BridgeBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBankTransactorSession struct {
	Contract     *BridgeBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBankRaw struct {
	Contract *BridgeBank // Generic contract binding to access the raw methods on
}

// BridgeBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBankCallerRaw struct {
	Contract *BridgeBankCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBankTransactorRaw struct {
	Contract *BridgeBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBank creates a new instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBank(address common.Address, backend bind.ContractBackend) (*BridgeBank, error) {
	contract, err := bindBridgeBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// NewBridgeBankCaller creates a new read-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankCaller(address common.Address, caller bind.ContractCaller) (*BridgeBankCaller, error) {
	contract, err := bindBridgeBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankCaller{contract: contract}, nil
}

// NewBridgeBankTransactor creates a new write-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBankTransactor, error) {
	contract, err := bindBridgeBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankTransactor{contract: contract}, nil
}

// NewBridgeBankFilterer creates a new log filterer instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBankFilterer, error) {
	contract, err := bindBridgeBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBankFilterer{contract: contract}, nil
}

// bindBridgeBank binds a generic wrapper to an already deployed contract.
func bindBridgeBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.BridgeBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transact(opts, method, params...)
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankCaller) ETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "ETHAddress")
	return *ret0, err
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankSession) ETHAddress() (common.Address, error) {
	return _BridgeBank.Contract.ETHAddress(&_BridgeBank.CallOpts)
}

// ETHAddress is a free data retrieval call binding the contract method 0xda228a9b.
//
// Solidity: function ETHAddress() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) ETHAddress() (common.Address, error) {
	return _BridgeBank.Contract.ETHAddress(&_BridgeBank.CallOpts)
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankCaller) ONEAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "ONEAddress")
	return *ret0, err
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankSession) ONEAddress() (common.Address, error) {
	return _BridgeBank.Contract.ONEAddress(&_BridgeBank.CallOpts)
}

// ONEAddress is a free data retrieval call binding the contract method 0x9d05974b.
//
// Solidity: function ONEAddress() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) ONEAddress() (common.Address, error) {
	return _BridgeBank.Contract.ONEAddress(&_BridgeBank.CallOpts)
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) SAFENUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "SAFE_NUMBER")
	return *ret0, err
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankSession) SAFENUMBER() (*big.Int, error) {
	return _BridgeBank.Contract.SAFENUMBER(&_BridgeBank.CallOpts)
}

// SAFENUMBER is a free data retrieval call binding the contract method 0x204897ff.
//
// Solidity: function SAFE_NUMBER() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) SAFENUMBER() (*big.Int, error) {
	return _BridgeBank.Contract.SAFENUMBER(&_BridgeBank.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "WETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankSession) WETH() (common.Address, error) {
	return _BridgeBank.Contract.WETH(&_BridgeBank.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) WETH() (common.Address, error) {
	return _BridgeBank.Contract.WETH(&_BridgeBank.CallOpts)
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankCaller) BandOracleInterface(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "bandOracleInterface")
	return *ret0, err
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankSession) BandOracleInterface() (common.Address, error) {
	return _BridgeBank.Contract.BandOracleInterface(&_BridgeBank.CallOpts)
}

// BandOracleInterface is a free data retrieval call binding the contract method 0xa1c1d3e2.
//
// Solidity: function bandOracleInterface() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) BandOracleInterface() (common.Address, error) {
	return _BridgeBank.Contract.BandOracleInterface(&_BridgeBank.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) FeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "feeDenominator")
	return *ret0, err
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankSession) FeeDenominator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeDenominator(&_BridgeBank.CallOpts)
}

// FeeDenominator is a free data retrieval call binding the contract method 0x180b0d7e.
//
// Solidity: function feeDenominator() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) FeeDenominator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeDenominator(&_BridgeBank.CallOpts)
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) FeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "feeNumerator")
	return *ret0, err
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankSession) FeeNumerator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeNumerator(&_BridgeBank.CallOpts)
}

// FeeNumerator is a free data retrieval call binding the contract method 0xe86dea4a.
//
// Solidity: function feeNumerator() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) FeeNumerator() (*big.Int, error) {
	return _BridgeBank.Contract.FeeNumerator(&_BridgeBank.CallOpts)
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetLockedFund(opts *bind.CallOpts, _ethereumToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getLockedFund", _ethereumToken)
	return *ret0, err
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetLockedFund(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetLockedFund(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetLockedFund is a free data retrieval call binding the contract method 0xfd930c7d.
//
// Solidity: function getLockedFund(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetLockedFund(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetLockedFund(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankCaller) GetTokenMappedAddress(opts *bind.CallOpts, _ethereumToken common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTokenMappedAddress", _ethereumToken)
	return *ret0, err
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankSession) GetTokenMappedAddress(_ethereumToken common.Address) (common.Address, error) {
	return _BridgeBank.Contract.GetTokenMappedAddress(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTokenMappedAddress is a free data retrieval call binding the contract method 0x722b1d23.
//
// Solidity: function getTokenMappedAddress(address _ethereumToken) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) GetTokenMappedAddress(_ethereumToken common.Address) (common.Address, error) {
	return _BridgeBank.Contract.GetTokenMappedAddress(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetTotalERC20Balance(opts *bind.CallOpts, _ethereumToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTotalERC20Balance", _ethereumToken)
	return *ret0, err
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetTotalERC20Balance(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalERC20Balance(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalERC20Balance is a free data retrieval call binding the contract method 0xac4a875c.
//
// Solidity: function getTotalERC20Balance(address _ethereumToken) view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetTotalERC20Balance(_ethereumToken common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalERC20Balance(&_BridgeBank.CallOpts, _ethereumToken)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) GetTotalETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "getTotalETHBalance")
	return *ret0, err
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankSession) GetTotalETHBalance() (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalETHBalance(&_BridgeBank.CallOpts)
}

// GetTotalETHBalance is a free data retrieval call binding the contract method 0x964d042c.
//
// Solidity: function getTotalETHBalance() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) GetTotalETHBalance() (*big.Int, error) {
	return _BridgeBank.Contract.GetTotalETHBalance(&_BridgeBank.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankCaller) HarmonyBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "harmonyBridge")
	return *ret0, err
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankSession) HarmonyBridge() (common.Address, error) {
	return _BridgeBank.Contract.HarmonyBridge(&_BridgeBank.CallOpts)
}

// HarmonyBridge is a free data retrieval call binding the contract method 0x972ded09.
//
// Solidity: function harmonyBridge() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) HarmonyBridge() (common.Address, error) {
	return _BridgeBank.Contract.HarmonyBridge(&_BridgeBank.CallOpts)
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankCaller) IsActiveToken(opts *bind.CallOpts, _ethereumToken common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "isActiveToken", _ethereumToken)
	return *ret0, err
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankSession) IsActiveToken(_ethereumToken common.Address) (bool, error) {
	return _BridgeBank.Contract.IsActiveToken(&_BridgeBank.CallOpts, _ethereumToken)
}

// IsActiveToken is a free data retrieval call binding the contract method 0x575c51c2.
//
// Solidity: function isActiveToken(address _ethereumToken) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) IsActiveToken(_ethereumToken common.Address) (bool, error) {
	return _BridgeBank.Contract.IsActiveToken(&_BridgeBank.CallOpts, _ethereumToken)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankCaller) LendingPool(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lendingPool")
	return *ret0, err
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankSession) LendingPool() (common.Address, error) {
	return _BridgeBank.Contract.LendingPool(&_BridgeBank.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) LendingPool() (common.Address, error) {
	return _BridgeBank.Contract.LendingPool(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "lockNonce")
	return *ret0, err
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_BridgeBank *BridgeBankCaller) SymbolToToken(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "symbolToToken", arg0)
	return *ret0, err
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_BridgeBank *BridgeBankSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _BridgeBank.Contract.SymbolToToken(&_BridgeBank.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _BridgeBank.Contract.SymbolToToken(&_BridgeBank.CallOpts, arg0)
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankCaller) TokensData(opts *bind.CallOpts, arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	ret := new(struct {
		LockedFund         *big.Int
		HarmonyMappedToken common.Address
		IsActive           bool
	})
	out := ret
	err := _BridgeBank.contract.Call(opts, out, "tokensData", arg0)
	return *ret, err
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankSession) TokensData(arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	return _BridgeBank.Contract.TokensData(&_BridgeBank.CallOpts, arg0)
}

// TokensData is a free data retrieval call binding the contract method 0xac66c6e0.
//
// Solidity: function tokensData(address ) view returns(uint256 lockedFund, address harmonyMappedToken, bool isActive)
func (_BridgeBank *BridgeBankCallerSession) TokensData(arg0 common.Address) (struct {
	LockedFund         *big.Int
	HarmonyMappedToken common.Address
	IsActive           bool
}, error) {
	return _BridgeBank.Contract.TokensData(&_BridgeBank.CallOpts, arg0)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankCaller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeBank.contract.Call(opts, out, "wethGateway")
	return *ret0, err
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankSession) WethGateway() (common.Address, error) {
	return _BridgeBank.Contract.WethGateway(&_BridgeBank.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) WethGateway() (common.Address, error) {
	return _BridgeBank.Contract.WethGateway(&_BridgeBank.CallOpts)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) ActivateToken(opts *bind.TransactOpts, _ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "activateToken", _ethereumToken)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankSession) ActivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ActivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) ActivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ActivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _ethereumToken, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankTransactor) AddToken(opts *bind.TransactOpts, _ethereumToken common.Address, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "addToken", _ethereumToken, _harmonyToken)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _ethereumToken, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankSession) AddToken(_ethereumToken common.Address, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken(&_BridgeBank.TransactOpts, _ethereumToken, _harmonyToken)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address _ethereumToken, address _harmonyToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) AddToken(_ethereumToken common.Address, _harmonyToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken(&_BridgeBank.TransactOpts, _ethereumToken, _harmonyToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) DeactivateToken(opts *bind.TransactOpts, _ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "deactivateToken", _ethereumToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankSession) DeactivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.DeactivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address _ethereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) DeactivateToken(_ethereumToken common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.DeactivateToken(&_BridgeBank.TransactOpts, _ethereumToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _operatorAddress, address _oracleAddress, address _harmonyBridgeAddress, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyWETH) returns()
func (_BridgeBank *BridgeBankTransactor) Initialize(opts *bind.TransactOpts, _operatorAddress common.Address, _oracleAddress common.Address, _harmonyBridgeAddress common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyWETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "initialize", _operatorAddress, _oracleAddress, _harmonyBridgeAddress, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyWETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _operatorAddress, address _oracleAddress, address _harmonyBridgeAddress, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyWETH) returns()
func (_BridgeBank *BridgeBankSession) Initialize(_operatorAddress common.Address, _oracleAddress common.Address, _harmonyBridgeAddress common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyWETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _operatorAddress, _oracleAddress, _harmonyBridgeAddress, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyWETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _operatorAddress, address _oracleAddress, address _harmonyBridgeAddress, address _bandOracleAddress, address _lendingPool, address _wethGateway, address _weth, address _harmonyWETH) returns()
func (_BridgeBank *BridgeBankTransactorSession) Initialize(_operatorAddress common.Address, _oracleAddress common.Address, _harmonyBridgeAddress common.Address, _bandOracleAddress common.Address, _lendingPool common.Address, _wethGateway common.Address, _weth common.Address, _harmonyWETH common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.Initialize(&_BridgeBank.TransactOpts, _operatorAddress, _oracleAddress, _harmonyBridgeAddress, _bandOracleAddress, _lendingPool, _wethGateway, _weth, _harmonyWETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForONE(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForONE", _harmonyReceiver, _amountETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForONE(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForONE is a paid mutator transaction binding the contract method 0x5e01297d.
//
// Solidity: function swapETHForONE(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForONE(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0xdf1c9337.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, string _destTokenSymbol) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForToken(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForToken", _harmonyReceiver, _amountETH, _destTokenSymbol)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0xdf1c9337.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, string _destTokenSymbol) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForToken(_harmonyReceiver common.Address, _amountETH *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH, _destTokenSymbol)
}

// SwapETHForToken is a paid mutator transaction binding the contract method 0xdf1c9337.
//
// Solidity: function swapETHForToken(address _harmonyReceiver, uint256 _amountETH, string _destTokenSymbol) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForToken(_harmonyReceiver common.Address, _amountETH *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH, _destTokenSymbol)
}

// SwapETHForWETH is a paid mutator transaction binding the contract method 0xda1a7bde.
//
// Solidity: function swapETHForWETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactor) SwapETHForWETH(opts *bind.TransactOpts, _harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapETHForWETH", _harmonyReceiver, _amountETH)
}

// SwapETHForWETH is a paid mutator transaction binding the contract method 0xda1a7bde.
//
// Solidity: function swapETHForWETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankSession) SwapETHForWETH(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForWETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapETHForWETH is a paid mutator transaction binding the contract method 0xda1a7bde.
//
// Solidity: function swapETHForWETH(address _harmonyReceiver, uint256 _amountETH) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapETHForWETH(_harmonyReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapETHForWETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _amountETH)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForONE(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForONE", _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForONE(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapTokenForONE is a paid mutator transaction binding the contract method 0xba1adba4.
//
// Solidity: function swapTokenForONE(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForONE(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForONE(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0xc3677410.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken, string _destTokenSymbol) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForToken(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForToken", _harmonyReceiver, _ethereumToken, _amountEthereumToken, _destTokenSymbol)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0xc3677410.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken, string _destTokenSymbol) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForToken(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken, _destTokenSymbol)
}

// SwapTokenForToken is a paid mutator transaction binding the contract method 0xc3677410.
//
// Solidity: function swapTokenForToken(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken, string _destTokenSymbol) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForToken(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int, _destTokenSymbol string) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForToken(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken, _destTokenSymbol)
}

// SwapTokenForWETH is a paid mutator transaction binding the contract method 0x80753a39.
//
// Solidity: function swapTokenForWETH(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) SwapTokenForWETH(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapTokenForWETH", _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapTokenForWETH is a paid mutator transaction binding the contract method 0x80753a39.
//
// Solidity: function swapTokenForWETH(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankSession) SwapTokenForWETH(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForWETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapTokenForWETH is a paid mutator transaction binding the contract method 0x80753a39.
//
// Solidity: function swapTokenForWETH(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapTokenForWETH(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapTokenForWETH(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) SwapToken11(opts *bind.TransactOpts, _harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "swapToken_1_1", _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankSession) SwapToken11(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapToken11(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// SwapToken11 is a paid mutator transaction binding the contract method 0xbb0a64db.
//
// Solidity: function swapToken_1_1(address _harmonyReceiver, address _ethereumToken, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) SwapToken11(_harmonyReceiver common.Address, _ethereumToken common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SwapToken11(&_BridgeBank.TransactOpts, _harmonyReceiver, _ethereumToken, _amountEthereumToken)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactor) UnlockERC20(opts *bind.TransactOpts, _ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlockERC20", _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankSession) UnlockERC20(_ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockERC20(&_BridgeBank.TransactOpts, _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockERC20 is a paid mutator transaction binding the contract method 0x3c298e78.
//
// Solidity: function unlockERC20(address _ethereumReceiver, address _ethereumToken, uint256 _ethereumTokenAmount) returns()
func (_BridgeBank *BridgeBankTransactorSession) UnlockERC20(_ethereumReceiver common.Address, _ethereumToken common.Address, _ethereumTokenAmount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockERC20(&_BridgeBank.TransactOpts, _ethereumReceiver, _ethereumToken, _ethereumTokenAmount)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactor) UnlockETH(opts *bind.TransactOpts, _ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlockETH", _ethereumReceiver, _amountETH)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankSession) UnlockETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// UnlockETH is a paid mutator transaction binding the contract method 0xed581fd3.
//
// Solidity: function unlockETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactorSession) UnlockETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UnlockETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateFee(opts *bind.TransactOpts, _feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateFee", _feeNumerator, _feeDenominator)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankSession) UpdateFee(_feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateFee(&_BridgeBank.TransactOpts, _feeNumerator, _feeDenominator)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x2740c197.
//
// Solidity: function updateFee(uint256 _feeNumerator, uint256 _feeDenominator) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateFee(_feeNumerator *big.Int, _feeDenominator *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateFee(&_BridgeBank.TransactOpts, _feeNumerator, _feeDenominator)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateHmyBridge(opts *bind.TransactOpts, _harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateHmyBridge", _harmonyBridge)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankSession) UpdateHmyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateHmyBridge(&_BridgeBank.TransactOpts, _harmonyBridge)
}

// UpdateHmyBridge is a paid mutator transaction binding the contract method 0x17877820.
//
// Solidity: function updateHmyBridge(address _harmonyBridge) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateHmyBridge(_harmonyBridge common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateHmyBridge(&_BridgeBank.TransactOpts, _harmonyBridge)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankTransactor) UpdateOracle(opts *bind.TransactOpts, _oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "updateOracle", _oracleAddress)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankSession) UpdateOracle(_oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateOracle(&_BridgeBank.TransactOpts, _oracleAddress)
}

// UpdateOracle is a paid mutator transaction binding the contract method 0x1cb44dfc.
//
// Solidity: function updateOracle(address _oracleAddress) returns()
func (_BridgeBank *BridgeBankTransactorSession) UpdateOracle(_oracleAddress common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.UpdateOracle(&_BridgeBank.TransactOpts, _oracleAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactor) WithdrawERC20(opts *bind.TransactOpts, _ethereumToken common.Address, _ethereumReceiver common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "withdrawERC20", _ethereumToken, _ethereumReceiver, _amountEthereumToken)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankSession) WithdrawERC20(_ethereumToken common.Address, _ethereumReceiver common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawERC20(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumReceiver, _amountEthereumToken)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address _ethereumToken, address _ethereumReceiver, uint256 _amountEthereumToken) returns()
func (_BridgeBank *BridgeBankTransactorSession) WithdrawERC20(_ethereumToken common.Address, _ethereumReceiver common.Address, _amountEthereumToken *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawERC20(&_BridgeBank.TransactOpts, _ethereumToken, _ethereumReceiver, _amountEthereumToken)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactor) WithdrawETH(opts *bind.TransactOpts, _ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "withdrawETH", _ethereumReceiver, _amountETH)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankSession) WithdrawETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address _ethereumReceiver, uint256 _amountETH) returns()
func (_BridgeBank *BridgeBankTransactorSession) WithdrawETH(_ethereumReceiver common.Address, _amountETH *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.WithdrawETH(&_BridgeBank.TransactOpts, _ethereumReceiver, _amountETH)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// BridgeBankLogLockIterator is returned from FilterLogLock and is used to iterate over the raw logs and unpacked data for LogLock events raised by the BridgeBank contract.
type BridgeBankLogLockIterator struct {
	Event *BridgeBankLogLock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogLock)
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
		it.Event = new(BridgeBankLogLock)
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
func (it *BridgeBankLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogLock represents a LogLock event raised by the BridgeBank contract.
type BridgeBankLogLock struct {
	EthereumSender      common.Address
	HarmonyReceiver     common.Address
	EthereumToken       common.Address
	HarmonyToken        common.Address
	AmountEthereumToken *big.Int
	AmountHarmonyToken  *big.Int
	Nonce               *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogLock is a free log retrieval operation binding the contract event 0x798ff61b60fe9fc186a882759a2cc2a91e6b9b17841f5b0abc965cf86da4d0b1.
//
// Solidity: event LogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _amountEthereumToken, uint256 _amountHarmonyToken, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) FilterLogLock(opts *bind.FilterOpts) (*BridgeBankLogLockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogLockIterator{contract: _BridgeBank.contract, event: "LogLock", logs: logs, sub: sub}, nil
}

// WatchLogLock is a free log subscription operation binding the contract event 0x798ff61b60fe9fc186a882759a2cc2a91e6b9b17841f5b0abc965cf86da4d0b1.
//
// Solidity: event LogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _amountEthereumToken, uint256 _amountHarmonyToken, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) WatchLogLock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogLock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogLock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
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

// ParseLogLock is a log parse operation binding the contract event 0x798ff61b60fe9fc186a882759a2cc2a91e6b9b17841f5b0abc965cf86da4d0b1.
//
// Solidity: event LogLock(address _ethereumSender, address _harmonyReceiver, address _ethereumToken, address _harmonyToken, uint256 _amountEthereumToken, uint256 _amountHarmonyToken, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) ParseLogLock(log types.Log) (*BridgeBankLogLock, error) {
	event := new(BridgeBankLogLock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankLogUnlockIterator is returned from FilterLogUnlock and is used to iterate over the raw logs and unpacked data for LogUnlock events raised by the BridgeBank contract.
type BridgeBankLogUnlockIterator struct {
	Event *BridgeBankLogUnlock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogUnlock)
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
		it.Event = new(BridgeBankLogUnlock)
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
func (it *BridgeBankLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogUnlock represents a LogUnlock event raised by the BridgeBank contract.
type BridgeBankLogUnlock struct {
	To    common.Address
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogUnlock is a free log retrieval operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) FilterLogUnlock(opts *bind.FilterOpts) (*BridgeBankLogUnlockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogUnlockIterator{contract: _BridgeBank.contract, event: "LogUnlock", logs: logs, sub: sub}, nil
}

// WatchLogUnlock is a free log subscription operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) WatchLogUnlock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogUnlock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogUnlock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
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

// ParseLogUnlock is a log parse operation binding the contract event 0xc2c64ff0cfc4d626042b306aa9b2f79227fcf39aeb429429a4d98d573fd009a4.
//
// Solidity: event LogUnlock(address _to, address _token, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) ParseLogUnlock(log types.Log) (*BridgeBankLogUnlock, error) {
	event := new(BridgeBankLogUnlock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankUpdateFeeIterator is returned from FilterUpdateFee and is used to iterate over the raw logs and unpacked data for UpdateFee events raised by the BridgeBank contract.
type BridgeBankUpdateFeeIterator struct {
	Event *BridgeBankUpdateFee // Event containing the contract specifics and raw log

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
func (it *BridgeBankUpdateFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankUpdateFee)
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
		it.Event = new(BridgeBankUpdateFee)
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
func (it *BridgeBankUpdateFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankUpdateFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankUpdateFee represents a UpdateFee event raised by the BridgeBank contract.
type BridgeBankUpdateFee struct {
	FeeNumerator   *big.Int
	FeeDenominator *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateFee is a free log retrieval operation binding the contract event 0x8987e6f43a6c6bf408c8c427dceb2f98377f859348939ef4ab7b770b510a395a.
//
// Solidity: event UpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) FilterUpdateFee(opts *bind.FilterOpts) (*BridgeBankUpdateFeeIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "UpdateFee")
	if err != nil {
		return nil, err
	}
	return &BridgeBankUpdateFeeIterator{contract: _BridgeBank.contract, event: "UpdateFee", logs: logs, sub: sub}, nil
}

// WatchUpdateFee is a free log subscription operation binding the contract event 0x8987e6f43a6c6bf408c8c427dceb2f98377f859348939ef4ab7b770b510a395a.
//
// Solidity: event UpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) WatchUpdateFee(opts *bind.WatchOpts, sink chan<- *BridgeBankUpdateFee) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "UpdateFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankUpdateFee)
				if err := _BridgeBank.contract.UnpackLog(event, "UpdateFee", log); err != nil {
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

// ParseUpdateFee is a log parse operation binding the contract event 0x8987e6f43a6c6bf408c8c427dceb2f98377f859348939ef4ab7b770b510a395a.
//
// Solidity: event UpdateFee(uint256 _feeNumerator, uint256 _feeDenominator)
func (_BridgeBank *BridgeBankFilterer) ParseUpdateFee(log types.Log) (*BridgeBankUpdateFee, error) {
	event := new(BridgeBankUpdateFee)
	if err := _BridgeBank.contract.UnpackLog(event, "UpdateFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankUpdateHarmonyBridgeIterator is returned from FilterUpdateHarmonyBridge and is used to iterate over the raw logs and unpacked data for UpdateHarmonyBridge events raised by the BridgeBank contract.
type BridgeBankUpdateHarmonyBridgeIterator struct {
	Event *BridgeBankUpdateHarmonyBridge // Event containing the contract specifics and raw log

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
func (it *BridgeBankUpdateHarmonyBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankUpdateHarmonyBridge)
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
		it.Event = new(BridgeBankUpdateHarmonyBridge)
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
func (it *BridgeBankUpdateHarmonyBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankUpdateHarmonyBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankUpdateHarmonyBridge represents a UpdateHarmonyBridge event raised by the BridgeBank contract.
type BridgeBankUpdateHarmonyBridge struct {
	NewHarmonyBridge common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateHarmonyBridge is a free log retrieval operation binding the contract event 0x586528e6ed06db137d25b0e2820571479b742a3ad01671a6614356db20dddd72.
//
// Solidity: event UpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) FilterUpdateHarmonyBridge(opts *bind.FilterOpts) (*BridgeBankUpdateHarmonyBridgeIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "UpdateHarmonyBridge")
	if err != nil {
		return nil, err
	}
	return &BridgeBankUpdateHarmonyBridgeIterator{contract: _BridgeBank.contract, event: "UpdateHarmonyBridge", logs: logs, sub: sub}, nil
}

// WatchUpdateHarmonyBridge is a free log subscription operation binding the contract event 0x586528e6ed06db137d25b0e2820571479b742a3ad01671a6614356db20dddd72.
//
// Solidity: event UpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) WatchUpdateHarmonyBridge(opts *bind.WatchOpts, sink chan<- *BridgeBankUpdateHarmonyBridge) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "UpdateHarmonyBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankUpdateHarmonyBridge)
				if err := _BridgeBank.contract.UnpackLog(event, "UpdateHarmonyBridge", log); err != nil {
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

// ParseUpdateHarmonyBridge is a log parse operation binding the contract event 0x586528e6ed06db137d25b0e2820571479b742a3ad01671a6614356db20dddd72.
//
// Solidity: event UpdateHarmonyBridge(address _newHarmonyBridge)
func (_BridgeBank *BridgeBankFilterer) ParseUpdateHarmonyBridge(log types.Log) (*BridgeBankUpdateHarmonyBridge, error) {
	event := new(BridgeBankUpdateHarmonyBridge)
	if err := _BridgeBank.contract.UnpackLog(event, "UpdateHarmonyBridge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankUpdateOracleIterator is returned from FilterUpdateOracle and is used to iterate over the raw logs and unpacked data for UpdateOracle events raised by the BridgeBank contract.
type BridgeBankUpdateOracleIterator struct {
	Event *BridgeBankUpdateOracle // Event containing the contract specifics and raw log

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
func (it *BridgeBankUpdateOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankUpdateOracle)
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
		it.Event = new(BridgeBankUpdateOracle)
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
func (it *BridgeBankUpdateOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankUpdateOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankUpdateOracle represents a UpdateOracle event raised by the BridgeBank contract.
type BridgeBankUpdateOracle struct {
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateOracle is a free log retrieval operation binding the contract event 0x09ad0a3595604db9b7aef0dbd4918cea3642b96bc65ad7c9fb501a1529becd79.
//
// Solidity: event UpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) FilterUpdateOracle(opts *bind.FilterOpts) (*BridgeBankUpdateOracleIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "UpdateOracle")
	if err != nil {
		return nil, err
	}
	return &BridgeBankUpdateOracleIterator{contract: _BridgeBank.contract, event: "UpdateOracle", logs: logs, sub: sub}, nil
}

// WatchUpdateOracle is a free log subscription operation binding the contract event 0x09ad0a3595604db9b7aef0dbd4918cea3642b96bc65ad7c9fb501a1529becd79.
//
// Solidity: event UpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) WatchUpdateOracle(opts *bind.WatchOpts, sink chan<- *BridgeBankUpdateOracle) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "UpdateOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankUpdateOracle)
				if err := _BridgeBank.contract.UnpackLog(event, "UpdateOracle", log); err != nil {
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

// ParseUpdateOracle is a log parse operation binding the contract event 0x09ad0a3595604db9b7aef0dbd4918cea3642b96bc65ad7c9fb501a1529becd79.
//
// Solidity: event UpdateOracle(address _newOracle)
func (_BridgeBank *BridgeBankFilterer) ParseUpdateOracle(log types.Log) (*BridgeBankUpdateOracle, error) {
	event := new(BridgeBankUpdateOracle)
	if err := _BridgeBank.contract.UnpackLog(event, "UpdateOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the BridgeBank contract.
type BridgeBankWithdrawERC20Iterator struct {
	Event *BridgeBankWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeBankWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankWithdrawERC20)
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
		it.Event = new(BridgeBankWithdrawERC20)
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
func (it *BridgeBankWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankWithdrawERC20 represents a WithdrawERC20 event raised by the BridgeBank contract.
type BridgeBankWithdrawERC20 struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0x33c35f9541201e342d5e7467016e65a0a06182eb12a5f17103f71cec95b6cb29.
//
// Solidity: event WithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) FilterWithdrawERC20(opts *bind.FilterOpts) (*BridgeBankWithdrawERC20Iterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeBankWithdrawERC20Iterator{contract: _BridgeBank.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0x33c35f9541201e342d5e7467016e65a0a06182eb12a5f17103f71cec95b6cb29.
//
// Solidity: event WithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *BridgeBankWithdrawERC20) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankWithdrawERC20)
				if err := _BridgeBank.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0x33c35f9541201e342d5e7467016e65a0a06182eb12a5f17103f71cec95b6cb29.
//
// Solidity: event WithdrawERC20(address _token, address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) ParseWithdrawERC20(log types.Log) (*BridgeBankWithdrawERC20, error) {
	event := new(BridgeBankWithdrawERC20)
	if err := _BridgeBank.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeBankWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the BridgeBank contract.
type BridgeBankWithdrawETHIterator struct {
	Event *BridgeBankWithdrawETH // Event containing the contract specifics and raw log

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
func (it *BridgeBankWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankWithdrawETH)
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
		it.Event = new(BridgeBankWithdrawETH)
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
func (it *BridgeBankWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankWithdrawETH represents a WithdrawETH event raised by the BridgeBank contract.
type BridgeBankWithdrawETH struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) FilterWithdrawETH(opts *bind.FilterOpts) (*BridgeBankWithdrawETHIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "WithdrawETH")
	if err != nil {
		return nil, err
	}
	return &BridgeBankWithdrawETHIterator{contract: _BridgeBank.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *BridgeBankWithdrawETH) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "WithdrawETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankWithdrawETH)
				if err := _BridgeBank.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0x566e45b1c8057e725bf62796a7f1d37ae294393cab069725a09daddd1af98b79.
//
// Solidity: event WithdrawETH(address _receiver, uint256 _amount)
func (_BridgeBank *BridgeBankFilterer) ParseWithdrawETH(log types.Log) (*BridgeBankWithdrawETH, error) {
	event := new(BridgeBankWithdrawETH)
	if err := _BridgeBank.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	return event, nil
}
