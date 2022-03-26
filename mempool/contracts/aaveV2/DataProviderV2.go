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

// AaveProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type AaveProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// DataProviderV2MetaData contains all meta data concerning the DataProviderV2 contract.
var DataProviderV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllATokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveTokensAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataProviderV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use DataProviderV2MetaData.ABI instead.
var DataProviderV2ABI = DataProviderV2MetaData.ABI

// DataProviderV2 is an auto generated Go binding around an Ethereum contract.
type DataProviderV2 struct {
	DataProviderV2Caller     // Read-only binding to the contract
	DataProviderV2Transactor // Write-only binding to the contract
	DataProviderV2Filterer   // Log filterer for contract events
}

// DataProviderV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type DataProviderV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataProviderV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DataProviderV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataProviderV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataProviderV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataProviderV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataProviderV2Session struct {
	Contract     *DataProviderV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataProviderV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataProviderV2CallerSession struct {
	Contract *DataProviderV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DataProviderV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataProviderV2TransactorSession struct {
	Contract     *DataProviderV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DataProviderV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type DataProviderV2Raw struct {
	Contract *DataProviderV2 // Generic contract binding to access the raw methods on
}

// DataProviderV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataProviderV2CallerRaw struct {
	Contract *DataProviderV2Caller // Generic read-only contract binding to access the raw methods on
}

// DataProviderV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataProviderV2TransactorRaw struct {
	Contract *DataProviderV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDataProviderV2 creates a new instance of DataProviderV2, bound to a specific deployed contract.
func NewDataProviderV2(address common.Address, backend bind.ContractBackend) (*DataProviderV2, error) {
	contract, err := bindDataProviderV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataProviderV2{DataProviderV2Caller: DataProviderV2Caller{contract: contract}, DataProviderV2Transactor: DataProviderV2Transactor{contract: contract}, DataProviderV2Filterer: DataProviderV2Filterer{contract: contract}}, nil
}

// NewDataProviderV2Caller creates a new read-only instance of DataProviderV2, bound to a specific deployed contract.
func NewDataProviderV2Caller(address common.Address, caller bind.ContractCaller) (*DataProviderV2Caller, error) {
	contract, err := bindDataProviderV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataProviderV2Caller{contract: contract}, nil
}

// NewDataProviderV2Transactor creates a new write-only instance of DataProviderV2, bound to a specific deployed contract.
func NewDataProviderV2Transactor(address common.Address, transactor bind.ContractTransactor) (*DataProviderV2Transactor, error) {
	contract, err := bindDataProviderV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataProviderV2Transactor{contract: contract}, nil
}

// NewDataProviderV2Filterer creates a new log filterer instance of DataProviderV2, bound to a specific deployed contract.
func NewDataProviderV2Filterer(address common.Address, filterer bind.ContractFilterer) (*DataProviderV2Filterer, error) {
	contract, err := bindDataProviderV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataProviderV2Filterer{contract: contract}, nil
}

// bindDataProviderV2 binds a generic wrapper to an already deployed contract.
func bindDataProviderV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataProviderV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataProviderV2 *DataProviderV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataProviderV2.Contract.DataProviderV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataProviderV2 *DataProviderV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataProviderV2.Contract.DataProviderV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataProviderV2 *DataProviderV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataProviderV2.Contract.DataProviderV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataProviderV2 *DataProviderV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataProviderV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataProviderV2 *DataProviderV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataProviderV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataProviderV2 *DataProviderV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataProviderV2.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_DataProviderV2 *DataProviderV2Caller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_DataProviderV2 *DataProviderV2Session) ADDRESSESPROVIDER() (common.Address, error) {
	return _DataProviderV2.Contract.ADDRESSESPROVIDER(&_DataProviderV2.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_DataProviderV2 *DataProviderV2CallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _DataProviderV2.Contract.ADDRESSESPROVIDER(&_DataProviderV2.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2Caller) GetAllATokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getAllATokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2Session) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _DataProviderV2.Contract.GetAllATokens(&_DataProviderV2.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2CallerSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _DataProviderV2.Contract.GetAllATokens(&_DataProviderV2.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2Caller) GetAllReservesTokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getAllReservesTokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2Session) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _DataProviderV2.Contract.GetAllReservesTokens(&_DataProviderV2.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_DataProviderV2 *DataProviderV2CallerSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _DataProviderV2.Contract.GetAllReservesTokens(&_DataProviderV2.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_DataProviderV2 *DataProviderV2Caller) GetReserveConfigurationData(opts *bind.CallOpts, asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getReserveConfigurationData", asset)

	outstruct := new(struct {
		Decimals                 *big.Int
		Ltv                      *big.Int
		LiquidationThreshold     *big.Int
		LiquidationBonus         *big.Int
		ReserveFactor            *big.Int
		UsageAsCollateralEnabled bool
		BorrowingEnabled         bool
		StableBorrowRateEnabled  bool
		IsActive                 bool
		IsFrozen                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimals = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveFactor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsFrozen = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_DataProviderV2 *DataProviderV2Session) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _DataProviderV2.Contract.GetReserveConfigurationData(&_DataProviderV2.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_DataProviderV2 *DataProviderV2CallerSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _DataProviderV2.Contract.GetReserveConfigurationData(&_DataProviderV2.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_DataProviderV2 *DataProviderV2Caller) GetReserveData(opts *bind.CallOpts, asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getReserveData", asset)

	outstruct := new(struct {
		AvailableLiquidity      *big.Int
		TotalStableDebt         *big.Int
		TotalVariableDebt       *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailableLiquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_DataProviderV2 *DataProviderV2Session) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _DataProviderV2.Contract.GetReserveData(&_DataProviderV2.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_DataProviderV2 *DataProviderV2CallerSession) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _DataProviderV2.Contract.GetReserveData(&_DataProviderV2.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_DataProviderV2 *DataProviderV2Caller) GetReserveTokensAddresses(opts *bind.CallOpts, asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getReserveTokensAddresses", asset)

	outstruct := new(struct {
		ATokenAddress            common.Address
		StableDebtTokenAddress   common.Address
		VariableDebtTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ATokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StableDebtTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VariableDebtTokenAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_DataProviderV2 *DataProviderV2Session) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _DataProviderV2.Contract.GetReserveTokensAddresses(&_DataProviderV2.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_DataProviderV2 *DataProviderV2CallerSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _DataProviderV2.Contract.GetReserveTokensAddresses(&_DataProviderV2.CallOpts, asset)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_DataProviderV2 *DataProviderV2Caller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _DataProviderV2.contract.Call(opts, &out, "getUserReserveData", asset, user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrincipalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScaledVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableRateLastUpdated = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_DataProviderV2 *DataProviderV2Session) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _DataProviderV2.Contract.GetUserReserveData(&_DataProviderV2.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_DataProviderV2 *DataProviderV2CallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _DataProviderV2.Contract.GetUserReserveData(&_DataProviderV2.CallOpts, asset, user)
}
