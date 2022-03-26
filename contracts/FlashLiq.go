// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// FlashLiqMetaData contains all meta data concerning the FlashLiq contract.
var FlashLiqMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"ADDRESSES_PROVIDER\",\"inputs\":[],\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"constant\":null,\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LENDING_POOL\",\"inputs\":[],\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"constant\":null,\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_owner\",\"inputs\":[],\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"constant\":null,\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOperation\",\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"constant\":null,\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapRouter\",\"inputs\":[],\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"constant\":null,\"stateMutability\":\"view\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// FlashLiqABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashLiqMetaData.ABI instead.
var FlashLiqABI = FlashLiqMetaData.ABI

// FlashLiq is an auto generated Go binding around an Ethereum contract.
type FlashLiq struct {
	FlashLiqCaller     // Read-only binding to the contract
	FlashLiqTransactor // Write-only binding to the contract
	FlashLiqFilterer   // Log filterer for contract events
}

// FlashLiqCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLiqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLiqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLiqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLiqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLiqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLiqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLiqSession struct {
	Contract     *FlashLiq         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashLiqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLiqCallerSession struct {
	Contract *FlashLiqCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FlashLiqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLiqTransactorSession struct {
	Contract     *FlashLiqTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FlashLiqRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLiqRaw struct {
	Contract *FlashLiq // Generic contract binding to access the raw methods on
}

// FlashLiqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLiqCallerRaw struct {
	Contract *FlashLiqCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLiqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLiqTransactorRaw struct {
	Contract *FlashLiqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLiq creates a new instance of FlashLiq, bound to a specific deployed contract.
func NewFlashLiq(address common.Address, backend bind.ContractBackend) (*FlashLiq, error) {
	contract, err := bindFlashLiq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLiq{FlashLiqCaller: FlashLiqCaller{contract: contract}, FlashLiqTransactor: FlashLiqTransactor{contract: contract}, FlashLiqFilterer: FlashLiqFilterer{contract: contract}}, nil
}

// NewFlashLiqCaller creates a new read-only instance of FlashLiq, bound to a specific deployed contract.
func NewFlashLiqCaller(address common.Address, caller bind.ContractCaller) (*FlashLiqCaller, error) {
	contract, err := bindFlashLiq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLiqCaller{contract: contract}, nil
}

// NewFlashLiqTransactor creates a new write-only instance of FlashLiq, bound to a specific deployed contract.
func NewFlashLiqTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLiqTransactor, error) {
	contract, err := bindFlashLiq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLiqTransactor{contract: contract}, nil
}

// NewFlashLiqFilterer creates a new log filterer instance of FlashLiq, bound to a specific deployed contract.
func NewFlashLiqFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLiqFilterer, error) {
	contract, err := bindFlashLiq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLiqFilterer{contract: contract}, nil
}

// bindFlashLiq binds a generic wrapper to an already deployed contract.
func bindFlashLiq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLiqABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLiq *FlashLiqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLiq.Contract.FlashLiqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLiq *FlashLiqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLiq.Contract.FlashLiqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLiq *FlashLiqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLiq.Contract.FlashLiqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLiq *FlashLiqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLiq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLiq *FlashLiqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLiq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLiq *FlashLiqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLiq.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLiq *FlashLiqCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLiq.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLiq *FlashLiqSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _FlashLiq.Contract.ADDRESSESPROVIDER(&_FlashLiq.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLiq *FlashLiqCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _FlashLiq.Contract.ADDRESSESPROVIDER(&_FlashLiq.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLiq *FlashLiqCaller) LENDINGPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLiq.contract.Call(opts, &out, "LENDING_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLiq *FlashLiqSession) LENDINGPOOL() (common.Address, error) {
	return _FlashLiq.Contract.LENDINGPOOL(&_FlashLiq.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLiq *FlashLiqCallerSession) LENDINGPOOL() (common.Address, error) {
	return _FlashLiq.Contract.LENDINGPOOL(&_FlashLiq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FlashLiq *FlashLiqCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLiq.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FlashLiq *FlashLiqSession) Owner() (common.Address, error) {
	return _FlashLiq.Contract.Owner(&_FlashLiq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_FlashLiq *FlashLiqCallerSession) Owner() (common.Address, error) {
	return _FlashLiq.Contract.Owner(&_FlashLiq.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_FlashLiq *FlashLiqCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLiq.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_FlashLiq *FlashLiqSession) SwapRouter() (common.Address, error) {
	return _FlashLiq.Contract.SwapRouter(&_FlashLiq.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_FlashLiq *FlashLiqCallerSession) SwapRouter() (common.Address, error) {
	return _FlashLiq.Contract.SwapRouter(&_FlashLiq.CallOpts)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLiq *FlashLiqTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLiq.contract.Transact(opts, "executeOperation", assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLiq *FlashLiqSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLiq.Contract.ExecuteOperation(&_FlashLiq.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLiq *FlashLiqTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLiq.Contract.ExecuteOperation(&_FlashLiq.TransactOpts, assets, amounts, premiums, initiator, params)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLiq *FlashLiqTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLiq.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLiq *FlashLiqSession) Receive() (*types.Transaction, error) {
	return _FlashLiq.Contract.Receive(&_FlashLiq.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashLiq *FlashLiqTransactorSession) Receive() (*types.Transaction, error) {
	return _FlashLiq.Contract.Receive(&_FlashLiq.TransactOpts)
}
