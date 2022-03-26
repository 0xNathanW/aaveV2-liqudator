// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveV2

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

// AddressProviderV2MetaData contains all meta data concerning the AddressProviderV2 contract.
var AddressProviderV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ConfigurationAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EmergencyAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCollateralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCollateralManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolCollateralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AddressProviderV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressProviderV2MetaData.ABI instead.
var AddressProviderV2ABI = AddressProviderV2MetaData.ABI

// AddressProviderV2 is an auto generated Go binding around an Ethereum contract.
type AddressProviderV2 struct {
	AddressProviderV2Caller     // Read-only binding to the contract
	AddressProviderV2Transactor // Write-only binding to the contract
	AddressProviderV2Filterer   // Log filterer for contract events
}

// AddressProviderV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AddressProviderV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressProviderV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressProviderV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressProviderV2Session struct {
	Contract     *AddressProviderV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressProviderV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressProviderV2CallerSession struct {
	Contract *AddressProviderV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AddressProviderV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressProviderV2TransactorSession struct {
	Contract     *AddressProviderV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AddressProviderV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AddressProviderV2Raw struct {
	Contract *AddressProviderV2 // Generic contract binding to access the raw methods on
}

// AddressProviderV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressProviderV2CallerRaw struct {
	Contract *AddressProviderV2Caller // Generic read-only contract binding to access the raw methods on
}

// AddressProviderV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressProviderV2TransactorRaw struct {
	Contract *AddressProviderV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressProviderV2 creates a new instance of AddressProviderV2, bound to a specific deployed contract.
func NewAddressProviderV2(address common.Address, backend bind.ContractBackend) (*AddressProviderV2, error) {
	contract, err := bindAddressProviderV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2{AddressProviderV2Caller: AddressProviderV2Caller{contract: contract}, AddressProviderV2Transactor: AddressProviderV2Transactor{contract: contract}, AddressProviderV2Filterer: AddressProviderV2Filterer{contract: contract}}, nil
}

// NewAddressProviderV2Caller creates a new read-only instance of AddressProviderV2, bound to a specific deployed contract.
func NewAddressProviderV2Caller(address common.Address, caller bind.ContractCaller) (*AddressProviderV2Caller, error) {
	contract, err := bindAddressProviderV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2Caller{contract: contract}, nil
}

// NewAddressProviderV2Transactor creates a new write-only instance of AddressProviderV2, bound to a specific deployed contract.
func NewAddressProviderV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AddressProviderV2Transactor, error) {
	contract, err := bindAddressProviderV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2Transactor{contract: contract}, nil
}

// NewAddressProviderV2Filterer creates a new log filterer instance of AddressProviderV2, bound to a specific deployed contract.
func NewAddressProviderV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AddressProviderV2Filterer, error) {
	contract, err := bindAddressProviderV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2Filterer{contract: contract}, nil
}

// bindAddressProviderV2 binds a generic wrapper to an already deployed contract.
func bindAddressProviderV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressProviderV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProviderV2 *AddressProviderV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProviderV2.Contract.AddressProviderV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProviderV2 *AddressProviderV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.AddressProviderV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProviderV2 *AddressProviderV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.AddressProviderV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProviderV2 *AddressProviderV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProviderV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProviderV2 *AddressProviderV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProviderV2 *AddressProviderV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetAddress(id [32]byte) (common.Address, error) {
	return _AddressProviderV2.Contract.GetAddress(&_AddressProviderV2.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AddressProviderV2.Contract.GetAddress(&_AddressProviderV2.CallOpts, id)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getEmergencyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetEmergencyAdmin() (common.Address, error) {
	return _AddressProviderV2.Contract.GetEmergencyAdmin(&_AddressProviderV2.CallOpts)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetEmergencyAdmin() (common.Address, error) {
	return _AddressProviderV2.Contract.GetEmergencyAdmin(&_AddressProviderV2.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetLendingPool() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPool(&_AddressProviderV2.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetLendingPool() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPool(&_AddressProviderV2.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetLendingPoolCollateralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getLendingPoolCollateralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetLendingPoolCollateralManager() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPoolCollateralManager(&_AddressProviderV2.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPoolCollateralManager(&_AddressProviderV2.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetLendingPoolConfigurator() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPoolConfigurator(&_AddressProviderV2.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingPoolConfigurator(&_AddressProviderV2.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetLendingRateOracle() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingRateOracle(&_AddressProviderV2.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetLendingRateOracle() (common.Address, error) {
	return _AddressProviderV2.Contract.GetLendingRateOracle(&_AddressProviderV2.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AddressProviderV2 *AddressProviderV2Caller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AddressProviderV2 *AddressProviderV2Session) GetMarketId() (string, error) {
	return _AddressProviderV2.Contract.GetMarketId(&_AddressProviderV2.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetMarketId() (string, error) {
	return _AddressProviderV2.Contract.GetMarketId(&_AddressProviderV2.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetPoolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getPoolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetPoolAdmin() (common.Address, error) {
	return _AddressProviderV2.Contract.GetPoolAdmin(&_AddressProviderV2.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetPoolAdmin() (common.Address, error) {
	return _AddressProviderV2.Contract.GetPoolAdmin(&_AddressProviderV2.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) GetPriceOracle() (common.Address, error) {
	return _AddressProviderV2.Contract.GetPriceOracle(&_AddressProviderV2.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) GetPriceOracle() (common.Address, error) {
	return _AddressProviderV2.Contract.GetPriceOracle(&_AddressProviderV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProviderV2 *AddressProviderV2Session) Owner() (common.Address, error) {
	return _AddressProviderV2.Contract.Owner(&_AddressProviderV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressProviderV2 *AddressProviderV2CallerSession) Owner() (common.Address, error) {
	return _AddressProviderV2.Contract.Owner(&_AddressProviderV2.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProviderV2 *AddressProviderV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _AddressProviderV2.Contract.RenounceOwnership(&_AddressProviderV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressProviderV2.Contract.RenounceOwnership(&_AddressProviderV2.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetAddress(&_AddressProviderV2.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetAddress(&_AddressProviderV2.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setAddressAsProxy", id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetAddressAsProxy(&_AddressProviderV2.TransactOpts, id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetAddressAsProxy(&_AddressProviderV2.TransactOpts, id, implementationAddress)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetEmergencyAdmin(opts *bind.TransactOpts, emergencyAdmin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setEmergencyAdmin", emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetEmergencyAdmin(&_AddressProviderV2.TransactOpts, emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetEmergencyAdmin(&_AddressProviderV2.TransactOpts, emergencyAdmin)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetLendingPoolCollateralManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setLendingPoolCollateralManager", manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolCollateralManager(&_AddressProviderV2.TransactOpts, manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolCollateralManager(&_AddressProviderV2.TransactOpts, manager)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, configurator common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setLendingPoolConfiguratorImpl", configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolConfiguratorImpl(&_AddressProviderV2.TransactOpts, configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolConfiguratorImpl(&_AddressProviderV2.TransactOpts, configurator)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetLendingPoolImpl(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setLendingPoolImpl", pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolImpl(&_AddressProviderV2.TransactOpts, pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingPoolImpl(&_AddressProviderV2.TransactOpts, pool)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetLendingRateOracle(opts *bind.TransactOpts, lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setLendingRateOracle", lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingRateOracle(&_AddressProviderV2.TransactOpts, lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetLendingRateOracle(&_AddressProviderV2.TransactOpts, lendingRateOracle)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetMarketId(opts *bind.TransactOpts, marketId string) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setMarketId", marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetMarketId(marketId string) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetMarketId(&_AddressProviderV2.TransactOpts, marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetMarketId(&_AddressProviderV2.TransactOpts, marketId)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetPoolAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setPoolAdmin", admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetPoolAdmin(&_AddressProviderV2.TransactOpts, admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetPoolAdmin(&_AddressProviderV2.TransactOpts, admin)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_AddressProviderV2 *AddressProviderV2Session) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetPriceOracle(&_AddressProviderV2.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.SetPriceOracle(&_AddressProviderV2.TransactOpts, priceOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProviderV2 *AddressProviderV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProviderV2 *AddressProviderV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.TransferOwnership(&_AddressProviderV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressProviderV2 *AddressProviderV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressProviderV2.Contract.TransferOwnership(&_AddressProviderV2.TransactOpts, newOwner)
}

// AddressProviderV2AddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AddressProviderV2 contract.
type AddressProviderV2AddressSetIterator struct {
	Event *AddressProviderV2AddressSet // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2AddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2AddressSet)
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
		it.Event = new(AddressProviderV2AddressSet)
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
func (it *AddressProviderV2AddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2AddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2AddressSet represents a AddressSet event raised by the AddressProviderV2 contract.
type AddressProviderV2AddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2AddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2AddressSetIterator{contract: _AddressProviderV2.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AddressProviderV2AddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2AddressSet)
				if err := _AddressProviderV2.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseAddressSet(log types.Log) (*AddressProviderV2AddressSet, error) {
	event := new(AddressProviderV2AddressSet)
	if err := _AddressProviderV2.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2ConfigurationAdminUpdatedIterator is returned from FilterConfigurationAdminUpdated and is used to iterate over the raw logs and unpacked data for ConfigurationAdminUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2ConfigurationAdminUpdatedIterator struct {
	Event *AddressProviderV2ConfigurationAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2ConfigurationAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2ConfigurationAdminUpdated)
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
		it.Event = new(AddressProviderV2ConfigurationAdminUpdated)
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
func (it *AddressProviderV2ConfigurationAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2ConfigurationAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2ConfigurationAdminUpdated represents a ConfigurationAdminUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2ConfigurationAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigurationAdminUpdated is a free log retrieval operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterConfigurationAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2ConfigurationAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2ConfigurationAdminUpdatedIterator{contract: _AddressProviderV2.contract, event: "ConfigurationAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigurationAdminUpdated is a free log subscription operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchConfigurationAdminUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2ConfigurationAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2ConfigurationAdminUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
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

// ParseConfigurationAdminUpdated is a log parse operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseConfigurationAdminUpdated(log types.Log) (*AddressProviderV2ConfigurationAdminUpdated, error) {
	event := new(AddressProviderV2ConfigurationAdminUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2EmergencyAdminUpdatedIterator is returned from FilterEmergencyAdminUpdated and is used to iterate over the raw logs and unpacked data for EmergencyAdminUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2EmergencyAdminUpdatedIterator struct {
	Event *AddressProviderV2EmergencyAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2EmergencyAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2EmergencyAdminUpdated)
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
		it.Event = new(AddressProviderV2EmergencyAdminUpdated)
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
func (it *AddressProviderV2EmergencyAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2EmergencyAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2EmergencyAdminUpdated represents a EmergencyAdminUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2EmergencyAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminUpdated is a free log retrieval operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterEmergencyAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2EmergencyAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2EmergencyAdminUpdatedIterator{contract: _AddressProviderV2.contract, event: "EmergencyAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminUpdated is a free log subscription operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchEmergencyAdminUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2EmergencyAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2EmergencyAdminUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
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

// ParseEmergencyAdminUpdated is a log parse operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseEmergencyAdminUpdated(log types.Log) (*AddressProviderV2EmergencyAdminUpdated, error) {
	event := new(AddressProviderV2EmergencyAdminUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2LendingPoolCollateralManagerUpdatedIterator is returned from FilterLendingPoolCollateralManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCollateralManagerUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolCollateralManagerUpdatedIterator struct {
	Event *AddressProviderV2LendingPoolCollateralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2LendingPoolCollateralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2LendingPoolCollateralManagerUpdated)
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
		it.Event = new(AddressProviderV2LendingPoolCollateralManagerUpdated)
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
func (it *AddressProviderV2LendingPoolCollateralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2LendingPoolCollateralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2LendingPoolCollateralManagerUpdated represents a LendingPoolCollateralManagerUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolCollateralManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCollateralManagerUpdated is a free log retrieval operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterLendingPoolCollateralManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2LendingPoolCollateralManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2LendingPoolCollateralManagerUpdatedIterator{contract: _AddressProviderV2.contract, event: "LendingPoolCollateralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCollateralManagerUpdated is a free log subscription operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchLendingPoolCollateralManagerUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2LendingPoolCollateralManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2LendingPoolCollateralManagerUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
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

// ParseLendingPoolCollateralManagerUpdated is a log parse operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseLendingPoolCollateralManagerUpdated(log types.Log) (*AddressProviderV2LendingPoolCollateralManagerUpdated, error) {
	event := new(AddressProviderV2LendingPoolCollateralManagerUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2LendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolConfiguratorUpdatedIterator struct {
	Event *AddressProviderV2LendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2LendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2LendingPoolConfiguratorUpdated)
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
		it.Event = new(AddressProviderV2LendingPoolConfiguratorUpdated)
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
func (it *AddressProviderV2LendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2LendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2LendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2LendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2LendingPoolConfiguratorUpdatedIterator{contract: _AddressProviderV2.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2LendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2LendingPoolConfiguratorUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*AddressProviderV2LendingPoolConfiguratorUpdated, error) {
	event := new(AddressProviderV2LendingPoolConfiguratorUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2LendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolUpdatedIterator struct {
	Event *AddressProviderV2LendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2LendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2LendingPoolUpdated)
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
		it.Event = new(AddressProviderV2LendingPoolUpdated)
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
func (it *AddressProviderV2LendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2LendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2LendingPoolUpdated represents a LendingPoolUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2LendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2LendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2LendingPoolUpdatedIterator{contract: _AddressProviderV2.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2LendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2LendingPoolUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseLendingPoolUpdated(log types.Log) (*AddressProviderV2LendingPoolUpdated, error) {
	event := new(AddressProviderV2LendingPoolUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2LendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2LendingRateOracleUpdatedIterator struct {
	Event *AddressProviderV2LendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2LendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2LendingRateOracleUpdated)
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
		it.Event = new(AddressProviderV2LendingRateOracleUpdated)
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
func (it *AddressProviderV2LendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2LendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2LendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2LendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2LendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2LendingRateOracleUpdatedIterator{contract: _AddressProviderV2.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2LendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2LendingRateOracleUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseLendingRateOracleUpdated(log types.Log) (*AddressProviderV2LendingRateOracleUpdated, error) {
	event := new(AddressProviderV2LendingRateOracleUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2MarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the AddressProviderV2 contract.
type AddressProviderV2MarketIdSetIterator struct {
	Event *AddressProviderV2MarketIdSet // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2MarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2MarketIdSet)
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
		it.Event = new(AddressProviderV2MarketIdSet)
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
func (it *AddressProviderV2MarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2MarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2MarketIdSet represents a MarketIdSet event raised by the AddressProviderV2 contract.
type AddressProviderV2MarketIdSet struct {
	NewMarketId string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterMarketIdSet(opts *bind.FilterOpts) (*AddressProviderV2MarketIdSetIterator, error) {

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2MarketIdSetIterator{contract: _AddressProviderV2.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *AddressProviderV2MarketIdSet) (event.Subscription, error) {

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2MarketIdSet)
				if err := _AddressProviderV2.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseMarketIdSet(log types.Log) (*AddressProviderV2MarketIdSet, error) {
	event := new(AddressProviderV2MarketIdSet)
	if err := _AddressProviderV2.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressProviderV2 contract.
type AddressProviderV2OwnershipTransferredIterator struct {
	Event *AddressProviderV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2OwnershipTransferred)
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
		it.Event = new(AddressProviderV2OwnershipTransferred)
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
func (it *AddressProviderV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2OwnershipTransferred represents a OwnershipTransferred event raised by the AddressProviderV2 contract.
type AddressProviderV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressProviderV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2OwnershipTransferredIterator{contract: _AddressProviderV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressProviderV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2OwnershipTransferred)
				if err := _AddressProviderV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseOwnershipTransferred(log types.Log) (*AddressProviderV2OwnershipTransferred, error) {
	event := new(AddressProviderV2OwnershipTransferred)
	if err := _AddressProviderV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2PriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the AddressProviderV2 contract.
type AddressProviderV2PriceOracleUpdatedIterator struct {
	Event *AddressProviderV2PriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2PriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2PriceOracleUpdated)
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
		it.Event = new(AddressProviderV2PriceOracleUpdated)
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
func (it *AddressProviderV2PriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2PriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2PriceOracleUpdated represents a PriceOracleUpdated event raised by the AddressProviderV2 contract.
type AddressProviderV2PriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2PriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2PriceOracleUpdatedIterator{contract: _AddressProviderV2.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2PriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2PriceOracleUpdated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParsePriceOracleUpdated(log types.Log) (*AddressProviderV2PriceOracleUpdated, error) {
	event := new(AddressProviderV2PriceOracleUpdated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressProviderV2ProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the AddressProviderV2 contract.
type AddressProviderV2ProxyCreatedIterator struct {
	Event *AddressProviderV2ProxyCreated // Event containing the contract specifics and raw log

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
func (it *AddressProviderV2ProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderV2ProxyCreated)
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
		it.Event = new(AddressProviderV2ProxyCreated)
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
func (it *AddressProviderV2ProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderV2ProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderV2ProxyCreated represents a ProxyCreated event raised by the AddressProviderV2 contract.
type AddressProviderV2ProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*AddressProviderV2ProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderV2ProxyCreatedIterator{contract: _AddressProviderV2.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *AddressProviderV2ProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AddressProviderV2.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderV2ProxyCreated)
				if err := _AddressProviderV2.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AddressProviderV2 *AddressProviderV2Filterer) ParseProxyCreated(log types.Log) (*AddressProviderV2ProxyCreated, error) {
	event := new(AddressProviderV2ProxyCreated)
	if err := _AddressProviderV2.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
