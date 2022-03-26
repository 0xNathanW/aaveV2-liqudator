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

// LendingPoolAddressesProviderMetaData contains all meta data concerning the LendingPoolAddressesProvider contract.
var LendingPoolAddressesProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ConfigurationAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EmergencyAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCollateralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCollateralManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolCollateralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingPoolAddressesProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingPoolAddressesProviderMetaData.ABI instead.
var LendingPoolAddressesProviderABI = LendingPoolAddressesProviderMetaData.ABI

// LendingPoolAddressesProvider is an auto generated Go binding around an Ethereum contract.
type LendingPoolAddressesProvider struct {
	LendingPoolAddressesProviderCaller     // Read-only binding to the contract
	LendingPoolAddressesProviderTransactor // Write-only binding to the contract
	LendingPoolAddressesProviderFilterer   // Log filterer for contract events
}

// LendingPoolAddressesProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolAddressesProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolAddressesProviderSession struct {
	Contract     *LendingPoolAddressesProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LendingPoolAddressesProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolAddressesProviderCallerSession struct {
	Contract *LendingPoolAddressesProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// LendingPoolAddressesProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolAddressesProviderTransactorSession struct {
	Contract     *LendingPoolAddressesProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// LendingPoolAddressesProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolAddressesProviderRaw struct {
	Contract *LendingPoolAddressesProvider // Generic contract binding to access the raw methods on
}

// LendingPoolAddressesProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderCallerRaw struct {
	Contract *LendingPoolAddressesProviderCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolAddressesProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderTransactorRaw struct {
	Contract *LendingPoolAddressesProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolAddressesProvider creates a new instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProvider(address common.Address, backend bind.ContractBackend) (*LendingPoolAddressesProvider, error) {
	contract, err := bindLendingPoolAddressesProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProvider{LendingPoolAddressesProviderCaller: LendingPoolAddressesProviderCaller{contract: contract}, LendingPoolAddressesProviderTransactor: LendingPoolAddressesProviderTransactor{contract: contract}, LendingPoolAddressesProviderFilterer: LendingPoolAddressesProviderFilterer{contract: contract}}, nil
}

// NewLendingPoolAddressesProviderCaller creates a new read-only instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolAddressesProviderCaller, error) {
	contract, err := bindLendingPoolAddressesProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderCaller{contract: contract}, nil
}

// NewLendingPoolAddressesProviderTransactor creates a new write-only instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolAddressesProviderTransactor, error) {
	contract, err := bindLendingPoolAddressesProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderTransactor{contract: contract}, nil
}

// NewLendingPoolAddressesProviderFilterer creates a new log filterer instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolAddressesProviderFilterer, error) {
	contract, err := bindLendingPoolAddressesProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderFilterer{contract: contract}, nil
}

// bindLendingPoolAddressesProvider binds a generic wrapper to an already deployed contract.
func bindLendingPoolAddressesProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolAddressesProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolAddressesProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetAddress(&_LendingPoolAddressesProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetAddress(&_LendingPoolAddressesProvider.CallOpts, id)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getEmergencyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetEmergencyAdmin() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetEmergencyAdmin(&_LendingPoolAddressesProvider.CallOpts)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetEmergencyAdmin() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetEmergencyAdmin(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPool() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPool(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPool() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPool(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolCollateralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolCollateralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolCollateralManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolCollateralManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetMarketId() (string, error) {
	return _LendingPoolAddressesProvider.Contract.GetMarketId(&_LendingPoolAddressesProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetMarketId() (string, error) {
	return _LendingPoolAddressesProvider.Contract.GetMarketId(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetPoolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getPoolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetPoolAdmin() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPoolAdmin(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetPoolAdmin() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPoolAdmin(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetPriceOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPriceOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPriceOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) Owner() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.Owner(&_LendingPoolAddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) Owner() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.Owner(&_LendingPoolAddressesProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.RenounceOwnership(&_LendingPoolAddressesProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.RenounceOwnership(&_LendingPoolAddressesProvider.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetAddress(&_LendingPoolAddressesProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetAddress(&_LendingPoolAddressesProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setAddressAsProxy", id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetAddressAsProxy(&_LendingPoolAddressesProvider.TransactOpts, id, implementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address implementationAddress) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetAddressAsProxy(id [32]byte, implementationAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetAddressAsProxy(&_LendingPoolAddressesProvider.TransactOpts, id, implementationAddress)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setEmergencyAdmin", emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetEmergencyAdmin(&_LendingPoolAddressesProvider.TransactOpts, emergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address emergencyAdmin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetEmergencyAdmin(emergencyAdmin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetEmergencyAdmin(&_LendingPoolAddressesProvider.TransactOpts, emergencyAdmin)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolCollateralManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolCollateralManager", manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolCollateralManager(&_LendingPoolAddressesProvider.TransactOpts, manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolCollateralManager(&_LendingPoolAddressesProvider.TransactOpts, manager)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolConfiguratorImpl", configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingPoolAddressesProvider.TransactOpts, configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingPoolAddressesProvider.TransactOpts, configurator)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolImpl", pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_LendingPoolAddressesProvider.TransactOpts, pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_LendingPoolAddressesProvider.TransactOpts, pool)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingRateOracle(opts *bind.TransactOpts, lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingRateOracle", lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_LendingPoolAddressesProvider.TransactOpts, lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_LendingPoolAddressesProvider.TransactOpts, lendingRateOracle)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetMarketId(opts *bind.TransactOpts, marketId string) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setMarketId", marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetMarketId(&_LendingPoolAddressesProvider.TransactOpts, marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetMarketId(&_LendingPoolAddressesProvider.TransactOpts, marketId)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetPoolAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setPoolAdmin", admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPoolAdmin(&_LendingPoolAddressesProvider.TransactOpts, admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPoolAdmin(&_LendingPoolAddressesProvider.TransactOpts, admin)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPriceOracle(&_LendingPoolAddressesProvider.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPriceOracle(&_LendingPoolAddressesProvider.TransactOpts, priceOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.TransferOwnership(&_LendingPoolAddressesProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.TransferOwnership(&_LendingPoolAddressesProvider.TransactOpts, newOwner)
}

// LendingPoolAddressesProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderAddressSetIterator struct {
	Event *LendingPoolAddressesProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderAddressSet)
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
		it.Event = new(LendingPoolAddressesProviderAddressSet)
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
func (it *LendingPoolAddressesProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderAddressSet represents a AddressSet event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderAddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderAddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderAddressSetIterator{contract: _LendingPoolAddressesProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderAddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderAddressSet)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseAddressSet(log types.Log) (*LendingPoolAddressesProviderAddressSet, error) {
	event := new(LendingPoolAddressesProviderAddressSet)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderConfigurationAdminUpdatedIterator is returned from FilterConfigurationAdminUpdated and is used to iterate over the raw logs and unpacked data for ConfigurationAdminUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderConfigurationAdminUpdatedIterator struct {
	Event *LendingPoolAddressesProviderConfigurationAdminUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderConfigurationAdminUpdated)
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
		it.Event = new(LendingPoolAddressesProviderConfigurationAdminUpdated)
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
func (it *LendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderConfigurationAdminUpdated represents a ConfigurationAdminUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderConfigurationAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigurationAdminUpdated is a free log retrieval operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterConfigurationAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderConfigurationAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderConfigurationAdminUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "ConfigurationAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigurationAdminUpdated is a free log subscription operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchConfigurationAdminUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderConfigurationAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderConfigurationAdminUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseConfigurationAdminUpdated(log types.Log) (*LendingPoolAddressesProviderConfigurationAdminUpdated, error) {
	event := new(LendingPoolAddressesProviderConfigurationAdminUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderEmergencyAdminUpdatedIterator is returned from FilterEmergencyAdminUpdated and is used to iterate over the raw logs and unpacked data for EmergencyAdminUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderEmergencyAdminUpdatedIterator struct {
	Event *LendingPoolAddressesProviderEmergencyAdminUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderEmergencyAdminUpdated)
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
		it.Event = new(LendingPoolAddressesProviderEmergencyAdminUpdated)
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
func (it *LendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderEmergencyAdminUpdated represents a EmergencyAdminUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderEmergencyAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminUpdated is a free log retrieval operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterEmergencyAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderEmergencyAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderEmergencyAdminUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "EmergencyAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminUpdated is a free log subscription operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchEmergencyAdminUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderEmergencyAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderEmergencyAdminUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseEmergencyAdminUpdated(log types.Log) (*LendingPoolAddressesProviderEmergencyAdminUpdated, error) {
	event := new(LendingPoolAddressesProviderEmergencyAdminUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator is returned from FilterLendingPoolCollateralManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCollateralManagerUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated represents a LendingPoolCollateralManagerUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCollateralManagerUpdated is a free log retrieval operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolCollateralManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolCollateralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCollateralManagerUpdated is a free log subscription operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolCollateralManagerUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolCollateralManagerUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolConfiguratorUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolUpdated represents a LendingPoolUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingRateOracleUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingRateOracleUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingRateOracleUpdated)
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
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingRateOracleUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingRateOracleUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingRateOracleUpdated(log types.Log) (*LendingPoolAddressesProviderLendingRateOracleUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingRateOracleUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderMarketIdSetIterator struct {
	Event *LendingPoolAddressesProviderMarketIdSet // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderMarketIdSet)
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
		it.Event = new(LendingPoolAddressesProviderMarketIdSet)
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
func (it *LendingPoolAddressesProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderMarketIdSet represents a MarketIdSet event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderMarketIdSet struct {
	NewMarketId string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts) (*LendingPoolAddressesProviderMarketIdSetIterator, error) {

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderMarketIdSetIterator{contract: _LendingPoolAddressesProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderMarketIdSet) (event.Subscription, error) {

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderMarketIdSet)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseMarketIdSet(log types.Log) (*LendingPoolAddressesProviderMarketIdSet, error) {
	event := new(LendingPoolAddressesProviderMarketIdSet)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderOwnershipTransferredIterator struct {
	Event *LendingPoolAddressesProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderOwnershipTransferred)
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
		it.Event = new(LendingPoolAddressesProviderOwnershipTransferred)
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
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderOwnershipTransferred represents a OwnershipTransferred event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingPoolAddressesProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderOwnershipTransferredIterator{contract: _LendingPoolAddressesProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderOwnershipTransferred)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseOwnershipTransferred(log types.Log) (*LendingPoolAddressesProviderOwnershipTransferred, error) {
	event := new(LendingPoolAddressesProviderOwnershipTransferred)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderPriceOracleUpdatedIterator struct {
	Event *LendingPoolAddressesProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderPriceOracleUpdated)
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
		it.Event = new(LendingPoolAddressesProviderPriceOracleUpdated)
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
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderPriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderPriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderPriceOracleUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderPriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderPriceOracleUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*LendingPoolAddressesProviderPriceOracleUpdated, error) {
	event := new(LendingPoolAddressesProviderPriceOracleUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolAddressesProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderProxyCreatedIterator struct {
	Event *LendingPoolAddressesProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderProxyCreated)
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
		it.Event = new(LendingPoolAddressesProviderProxyCreated)
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
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderProxyCreated represents a ProxyCreated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderProxyCreatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderProxyCreated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseProxyCreated(log types.Log) (*LendingPoolAddressesProviderProxyCreated, error) {
	event := new(LendingPoolAddressesProviderProxyCreated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
