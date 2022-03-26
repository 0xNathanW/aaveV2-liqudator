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

// LendingRateOracleMetaData contains all meta data concerning the LendingRateOracle contract.
var LendingRateOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"MarketBorrowRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getMarketBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setMarketBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingRateOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingRateOracleMetaData.ABI instead.
var LendingRateOracleABI = LendingRateOracleMetaData.ABI

// LendingRateOracle is an auto generated Go binding around an Ethereum contract.
type LendingRateOracle struct {
	LendingRateOracleCaller     // Read-only binding to the contract
	LendingRateOracleTransactor // Write-only binding to the contract
	LendingRateOracleFilterer   // Log filterer for contract events
}

// LendingRateOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingRateOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingRateOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingRateOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingRateOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingRateOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingRateOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingRateOracleSession struct {
	Contract     *LendingRateOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LendingRateOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingRateOracleCallerSession struct {
	Contract *LendingRateOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LendingRateOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingRateOracleTransactorSession struct {
	Contract     *LendingRateOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LendingRateOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingRateOracleRaw struct {
	Contract *LendingRateOracle // Generic contract binding to access the raw methods on
}

// LendingRateOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingRateOracleCallerRaw struct {
	Contract *LendingRateOracleCaller // Generic read-only contract binding to access the raw methods on
}

// LendingRateOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingRateOracleTransactorRaw struct {
	Contract *LendingRateOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingRateOracle creates a new instance of LendingRateOracle, bound to a specific deployed contract.
func NewLendingRateOracle(address common.Address, backend bind.ContractBackend) (*LendingRateOracle, error) {
	contract, err := bindLendingRateOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracle{LendingRateOracleCaller: LendingRateOracleCaller{contract: contract}, LendingRateOracleTransactor: LendingRateOracleTransactor{contract: contract}, LendingRateOracleFilterer: LendingRateOracleFilterer{contract: contract}}, nil
}

// NewLendingRateOracleCaller creates a new read-only instance of LendingRateOracle, bound to a specific deployed contract.
func NewLendingRateOracleCaller(address common.Address, caller bind.ContractCaller) (*LendingRateOracleCaller, error) {
	contract, err := bindLendingRateOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracleCaller{contract: contract}, nil
}

// NewLendingRateOracleTransactor creates a new write-only instance of LendingRateOracle, bound to a specific deployed contract.
func NewLendingRateOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingRateOracleTransactor, error) {
	contract, err := bindLendingRateOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracleTransactor{contract: contract}, nil
}

// NewLendingRateOracleFilterer creates a new log filterer instance of LendingRateOracle, bound to a specific deployed contract.
func NewLendingRateOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingRateOracleFilterer, error) {
	contract, err := bindLendingRateOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracleFilterer{contract: contract}, nil
}

// bindLendingRateOracle binds a generic wrapper to an already deployed contract.
func bindLendingRateOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingRateOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingRateOracle *LendingRateOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingRateOracle.Contract.LendingRateOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingRateOracle *LendingRateOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.LendingRateOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingRateOracle *LendingRateOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.LendingRateOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingRateOracle *LendingRateOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingRateOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingRateOracle *LendingRateOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingRateOracle *LendingRateOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.contract.Transact(opts, method, params...)
}

// GetMarketBorrowRate is a free data retrieval call binding the contract method 0xbb85c0bb.
//
// Solidity: function getMarketBorrowRate(address asset) view returns(uint256)
func (_LendingRateOracle *LendingRateOracleCaller) GetMarketBorrowRate(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingRateOracle.contract.Call(opts, &out, "getMarketBorrowRate", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketBorrowRate is a free data retrieval call binding the contract method 0xbb85c0bb.
//
// Solidity: function getMarketBorrowRate(address asset) view returns(uint256)
func (_LendingRateOracle *LendingRateOracleSession) GetMarketBorrowRate(asset common.Address) (*big.Int, error) {
	return _LendingRateOracle.Contract.GetMarketBorrowRate(&_LendingRateOracle.CallOpts, asset)
}

// GetMarketBorrowRate is a free data retrieval call binding the contract method 0xbb85c0bb.
//
// Solidity: function getMarketBorrowRate(address asset) view returns(uint256)
func (_LendingRateOracle *LendingRateOracleCallerSession) GetMarketBorrowRate(asset common.Address) (*big.Int, error) {
	return _LendingRateOracle.Contract.GetMarketBorrowRate(&_LendingRateOracle.CallOpts, asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingRateOracle *LendingRateOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingRateOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingRateOracle *LendingRateOracleSession) Owner() (common.Address, error) {
	return _LendingRateOracle.Contract.Owner(&_LendingRateOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingRateOracle *LendingRateOracleCallerSession) Owner() (common.Address, error) {
	return _LendingRateOracle.Contract.Owner(&_LendingRateOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingRateOracle *LendingRateOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingRateOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingRateOracle *LendingRateOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingRateOracle.Contract.RenounceOwnership(&_LendingRateOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingRateOracle *LendingRateOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingRateOracle.Contract.RenounceOwnership(&_LendingRateOracle.TransactOpts)
}

// SetMarketBorrowRate is a paid mutator transaction binding the contract method 0x72eb293d.
//
// Solidity: function setMarketBorrowRate(address asset, uint256 rate) returns()
func (_LendingRateOracle *LendingRateOracleTransactor) SetMarketBorrowRate(opts *bind.TransactOpts, asset common.Address, rate *big.Int) (*types.Transaction, error) {
	return _LendingRateOracle.contract.Transact(opts, "setMarketBorrowRate", asset, rate)
}

// SetMarketBorrowRate is a paid mutator transaction binding the contract method 0x72eb293d.
//
// Solidity: function setMarketBorrowRate(address asset, uint256 rate) returns()
func (_LendingRateOracle *LendingRateOracleSession) SetMarketBorrowRate(asset common.Address, rate *big.Int) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.SetMarketBorrowRate(&_LendingRateOracle.TransactOpts, asset, rate)
}

// SetMarketBorrowRate is a paid mutator transaction binding the contract method 0x72eb293d.
//
// Solidity: function setMarketBorrowRate(address asset, uint256 rate) returns()
func (_LendingRateOracle *LendingRateOracleTransactorSession) SetMarketBorrowRate(asset common.Address, rate *big.Int) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.SetMarketBorrowRate(&_LendingRateOracle.TransactOpts, asset, rate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingRateOracle *LendingRateOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingRateOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingRateOracle *LendingRateOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.TransferOwnership(&_LendingRateOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingRateOracle *LendingRateOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingRateOracle.Contract.TransferOwnership(&_LendingRateOracle.TransactOpts, newOwner)
}

// LendingRateOracleMarketBorrowRateSetIterator is returned from FilterMarketBorrowRateSet and is used to iterate over the raw logs and unpacked data for MarketBorrowRateSet events raised by the LendingRateOracle contract.
type LendingRateOracleMarketBorrowRateSetIterator struct {
	Event *LendingRateOracleMarketBorrowRateSet // Event containing the contract specifics and raw log

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
func (it *LendingRateOracleMarketBorrowRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRateOracleMarketBorrowRateSet)
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
		it.Event = new(LendingRateOracleMarketBorrowRateSet)
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
func (it *LendingRateOracleMarketBorrowRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRateOracleMarketBorrowRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRateOracleMarketBorrowRateSet represents a MarketBorrowRateSet event raised by the LendingRateOracle contract.
type LendingRateOracleMarketBorrowRateSet struct {
	Asset common.Address
	Rate  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMarketBorrowRateSet is a free log retrieval operation binding the contract event 0xeddf60a41f2e6c1a025372190672a8274d9c1fcb0eacda2aa792493d0af343d7.
//
// Solidity: event MarketBorrowRateSet(address indexed asset, uint256 rate)
func (_LendingRateOracle *LendingRateOracleFilterer) FilterMarketBorrowRateSet(opts *bind.FilterOpts, asset []common.Address) (*LendingRateOracleMarketBorrowRateSetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _LendingRateOracle.contract.FilterLogs(opts, "MarketBorrowRateSet", assetRule)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracleMarketBorrowRateSetIterator{contract: _LendingRateOracle.contract, event: "MarketBorrowRateSet", logs: logs, sub: sub}, nil
}

// WatchMarketBorrowRateSet is a free log subscription operation binding the contract event 0xeddf60a41f2e6c1a025372190672a8274d9c1fcb0eacda2aa792493d0af343d7.
//
// Solidity: event MarketBorrowRateSet(address indexed asset, uint256 rate)
func (_LendingRateOracle *LendingRateOracleFilterer) WatchMarketBorrowRateSet(opts *bind.WatchOpts, sink chan<- *LendingRateOracleMarketBorrowRateSet, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _LendingRateOracle.contract.WatchLogs(opts, "MarketBorrowRateSet", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRateOracleMarketBorrowRateSet)
				if err := _LendingRateOracle.contract.UnpackLog(event, "MarketBorrowRateSet", log); err != nil {
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

// ParseMarketBorrowRateSet is a log parse operation binding the contract event 0xeddf60a41f2e6c1a025372190672a8274d9c1fcb0eacda2aa792493d0af343d7.
//
// Solidity: event MarketBorrowRateSet(address indexed asset, uint256 rate)
func (_LendingRateOracle *LendingRateOracleFilterer) ParseMarketBorrowRateSet(log types.Log) (*LendingRateOracleMarketBorrowRateSet, error) {
	event := new(LendingRateOracleMarketBorrowRateSet)
	if err := _LendingRateOracle.contract.UnpackLog(event, "MarketBorrowRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRateOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingRateOracle contract.
type LendingRateOracleOwnershipTransferredIterator struct {
	Event *LendingRateOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingRateOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRateOracleOwnershipTransferred)
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
		it.Event = new(LendingRateOracleOwnershipTransferred)
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
func (it *LendingRateOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRateOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRateOracleOwnershipTransferred represents a OwnershipTransferred event raised by the LendingRateOracle contract.
type LendingRateOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingRateOracle *LendingRateOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingRateOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingRateOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingRateOracleOwnershipTransferredIterator{contract: _LendingRateOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingRateOracle *LendingRateOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingRateOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingRateOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRateOracleOwnershipTransferred)
				if err := _LendingRateOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LendingRateOracle *LendingRateOracleFilterer) ParseOwnershipTransferred(log types.Log) (*LendingRateOracleOwnershipTransferred, error) {
	event := new(LendingRateOracleOwnershipTransferred)
	if err := _LendingRateOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
