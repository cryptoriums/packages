// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple

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

// SimpleStorageMetaData contains all meta data concerning the SimpleStorage contract.
var SimpleStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"StorageSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6d4ce63c": "get()",
		"4ed3885e": "set(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610374806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634ed3885e1461003b5780636d4ce63c14610050575b600080fd5b61004e6100493660046101fd565b61006e565b005b6100586100bc565b60405161006591906102ae565b60405180910390f35b805161008190600090602084019061014e565b507f60cdc157ca4737a5d39a0e107532616e7bf6946e615228812c7f9bc9f81797a6816040516100b191906102ae565b60405180910390a150565b6060600080546100cb90610303565b80601f01602080910402602001604051908101604052809291908181526020018280546100f790610303565b80156101445780601f1061011957610100808354040283529160200191610144565b820191906000526020600020905b81548152906001019060200180831161012757829003601f168201915b5050505050905090565b82805461015a90610303565b90600052602060002090601f01602090048101928261017c57600085556101c2565b82601f1061019557805160ff19168380011785556101c2565b828001600101855582156101c2579182015b828111156101c25782518255916020019190600101906101a7565b506101ce9291506101d2565b5090565b5b808211156101ce57600081556001016101d3565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561020f57600080fd5b813567ffffffffffffffff8082111561022757600080fd5b818401915084601f83011261023b57600080fd5b81358181111561024d5761024d6101e7565b604051601f8201601f19908116603f01168101908382118183101715610275576102756101e7565b8160405282815287602084870101111561028e57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156102db578581018301518582016040015282016102bf565b818111156102ed576000604083870101525b50601f01601f1916929092016040019392505050565b600181811c9082168061031757607f821691505b6020821081141561033857634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212206e2edd1110079e40519c07d414dbebe9516dd3ce44fa505d73d872316ffa151264736f6c634300080a0033",
}

// SimpleStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleStorageMetaData.ABI instead.
var SimpleStorageABI = SimpleStorageMetaData.ABI

// Deprecated: Use SimpleStorageMetaData.Sigs instead.
// SimpleStorageFuncSigs maps the 4-byte function signature to its string representation.
var SimpleStorageFuncSigs = SimpleStorageMetaData.Sigs

// SimpleStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleStorageMetaData.Bin instead.
var SimpleStorageBin = SimpleStorageMetaData.Bin

// DeploySimpleStorage deploys a new Ethereum contract, binding an instance of SimpleStorage to it.
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	parsed, err := SimpleStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_SimpleStorage *SimpleStorageCaller) Get(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_SimpleStorage *SimpleStorageSession) Get() (string, error) {
	return _SimpleStorage.Contract.Get(&_SimpleStorage.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(string)
func (_SimpleStorage *SimpleStorageCallerSession) Get() (string, error) {
	return _SimpleStorage.Contract.Get(&_SimpleStorage.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string x) returns()
func (_SimpleStorage *SimpleStorageTransactor) Set(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string x) returns()
func (_SimpleStorage *SimpleStorageSession) Set(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Set(&_SimpleStorage.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string x) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) Set(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Set(&_SimpleStorage.TransactOpts, x)
}

// SimpleStorageStorageSetIterator is returned from FilterStorageSet and is used to iterate over the raw logs and unpacked data for StorageSet events raised by the SimpleStorage contract.
type SimpleStorageStorageSetIterator struct {
	Event *SimpleStorageStorageSet // Event containing the contract specifics and raw log

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
func (it *SimpleStorageStorageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageStorageSet)
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
		it.Event = new(SimpleStorageStorageSet)
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
func (it *SimpleStorageStorageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageStorageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageStorageSet represents a StorageSet event raised by the SimpleStorage contract.
type SimpleStorageStorageSet struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStorageSet is a free log retrieval operation binding the contract event 0x60cdc157ca4737a5d39a0e107532616e7bf6946e615228812c7f9bc9f81797a6.
//
// Solidity: event StorageSet(string arg0)
func (_SimpleStorage *SimpleStorageFilterer) FilterStorageSet(opts *bind.FilterOpts) (*SimpleStorageStorageSetIterator, error) {

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "StorageSet")
	if err != nil {
		return nil, err
	}
	return &SimpleStorageStorageSetIterator{contract: _SimpleStorage.contract, event: "StorageSet", logs: logs, sub: sub}, nil
}

// WatchStorageSet is a free log subscription operation binding the contract event 0x60cdc157ca4737a5d39a0e107532616e7bf6946e615228812c7f9bc9f81797a6.
//
// Solidity: event StorageSet(string arg0)
func (_SimpleStorage *SimpleStorageFilterer) WatchStorageSet(opts *bind.WatchOpts, sink chan<- *SimpleStorageStorageSet) (event.Subscription, error) {

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "StorageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageStorageSet)
				if err := _SimpleStorage.contract.UnpackLog(event, "StorageSet", log); err != nil {
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

// ParseStorageSet is a log parse operation binding the contract event 0x60cdc157ca4737a5d39a0e107532616e7bf6946e615228812c7f9bc9f81797a6.
//
// Solidity: event StorageSet(string arg0)
func (_SimpleStorage *SimpleStorageFilterer) ParseStorageSet(log types.Log) (*SimpleStorageStorageSet, error) {
	event := new(SimpleStorageStorageSet)
	if err := _SimpleStorage.contract.UnpackLog(event, "StorageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
