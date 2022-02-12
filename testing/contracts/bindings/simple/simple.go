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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"NewString\",\"type\":\"string\"}],\"name\":\"StorageSetA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"NewString\",\"type\":\"string\"}],\"name\":\"StorageSetB\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getB\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"setA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"setB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d46300fd": "getA()",
		"a1c51915": "getB()",
		"b958abd5": "setA(string)",
		"b5e7bc60": "setB(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b506103f7806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063a1c5191514610051578063b5e7bc601461006f578063b958abd514610084578063d46300fd14610097575b600080fd5b61005961009f565b604051610066919061026a565b60405180910390f35b61008261007d3660046102d5565b610131565b005b6100826100923660046102d5565b61017f565b6100596101c2565b6060600180546100ae90610386565b80601f01602080910402602001604051908101604052809291908181526020018280546100da90610386565b80156101275780601f106100fc57610100808354040283529160200191610127565b820191906000526020600020905b81548152906001019060200180831161010a57829003601f168201915b5050505050905090565b80516101449060019060208401906101d1565b507fffad11ecb618b0f8f566637fe9991a6c3ec300a88293e7e9e68ddcee51b861b581604051610174919061026a565b60405180910390a150565b80516101929060009060208401906101d1565b507fe394429e5515af0c383c3d1fc989369cfcce852ba8bc8300b01b64bfdb3d18be81604051610174919061026a565b6060600080546100ae90610386565b8280546101dd90610386565b90600052602060002090601f0160209004810192826101ff5760008555610245565b82601f1061021857805160ff1916838001178555610245565b82800160010185558215610245579182015b8281111561024557825182559160200191906001019061022a565b50610251929150610255565b5090565b5b808211156102515760008155600101610256565b600060208083528351808285015260005b818110156102975785810183015185820160400152820161027b565b818111156102a9576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156102e757600080fd5b813567ffffffffffffffff808211156102ff57600080fd5b818401915084601f83011261031357600080fd5b813581811115610325576103256102bf565b604051601f8201601f19908116603f0116810190838211818310171561034d5761034d6102bf565b8160405282815287602084870101111561036657600080fd5b826020860160208301376000928101602001929092525095945050505050565b600181811c9082168061039a57607f821691505b602082108114156103bb57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220417a622ce55771beebb61181aeaaa5b5215b50e4382bea264d04acb8de65154c64736f6c634300080a0033",
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

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(string)
func (_SimpleStorage *SimpleStorageCaller) GetA(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(string)
func (_SimpleStorage *SimpleStorageSession) GetA() (string, error) {
	return _SimpleStorage.Contract.GetA(&_SimpleStorage.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(string)
func (_SimpleStorage *SimpleStorageCallerSession) GetA() (string, error) {
	return _SimpleStorage.Contract.GetA(&_SimpleStorage.CallOpts)
}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() view returns(string)
func (_SimpleStorage *SimpleStorageCaller) GetB(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "getB")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() view returns(string)
func (_SimpleStorage *SimpleStorageSession) GetB() (string, error) {
	return _SimpleStorage.Contract.GetB(&_SimpleStorage.CallOpts)
}

// GetB is a free data retrieval call binding the contract method 0xa1c51915.
//
// Solidity: function getB() view returns(string)
func (_SimpleStorage *SimpleStorageCallerSession) GetB() (string, error) {
	return _SimpleStorage.Contract.GetB(&_SimpleStorage.CallOpts)
}

// SetA is a paid mutator transaction binding the contract method 0xb958abd5.
//
// Solidity: function setA(string x) returns()
func (_SimpleStorage *SimpleStorageTransactor) SetA(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "setA", x)
}

// SetA is a paid mutator transaction binding the contract method 0xb958abd5.
//
// Solidity: function setA(string x) returns()
func (_SimpleStorage *SimpleStorageSession) SetA(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetA(&_SimpleStorage.TransactOpts, x)
}

// SetA is a paid mutator transaction binding the contract method 0xb958abd5.
//
// Solidity: function setA(string x) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) SetA(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetA(&_SimpleStorage.TransactOpts, x)
}

// SetB is a paid mutator transaction binding the contract method 0xb5e7bc60.
//
// Solidity: function setB(string x) returns()
func (_SimpleStorage *SimpleStorageTransactor) SetB(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "setB", x)
}

// SetB is a paid mutator transaction binding the contract method 0xb5e7bc60.
//
// Solidity: function setB(string x) returns()
func (_SimpleStorage *SimpleStorageSession) SetB(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetB(&_SimpleStorage.TransactOpts, x)
}

// SetB is a paid mutator transaction binding the contract method 0xb5e7bc60.
//
// Solidity: function setB(string x) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) SetB(x string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetB(&_SimpleStorage.TransactOpts, x)
}

// SimpleStorageStorageSetAIterator is returned from FilterStorageSetA and is used to iterate over the raw logs and unpacked data for StorageSetA events raised by the SimpleStorage contract.
type SimpleStorageStorageSetAIterator struct {
	Event *SimpleStorageStorageSetA // Event containing the contract specifics and raw log

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
func (it *SimpleStorageStorageSetAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageStorageSetA)
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
		it.Event = new(SimpleStorageStorageSetA)
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
func (it *SimpleStorageStorageSetAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageStorageSetAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageStorageSetA represents a StorageSetA event raised by the SimpleStorage contract.
type SimpleStorageStorageSetA struct {
	NewString string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStorageSetA is a free log retrieval operation binding the contract event 0xe394429e5515af0c383c3d1fc989369cfcce852ba8bc8300b01b64bfdb3d18be.
//
// Solidity: event StorageSetA(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) FilterStorageSetA(opts *bind.FilterOpts) (*SimpleStorageStorageSetAIterator, error) {

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "StorageSetA")
	if err != nil {
		return nil, err
	}
	return &SimpleStorageStorageSetAIterator{contract: _SimpleStorage.contract, event: "StorageSetA", logs: logs, sub: sub}, nil
}

// WatchStorageSetA is a free log subscription operation binding the contract event 0xe394429e5515af0c383c3d1fc989369cfcce852ba8bc8300b01b64bfdb3d18be.
//
// Solidity: event StorageSetA(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) WatchStorageSetA(opts *bind.WatchOpts, sink chan<- *SimpleStorageStorageSetA) (event.Subscription, error) {

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "StorageSetA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageStorageSetA)
				if err := _SimpleStorage.contract.UnpackLog(event, "StorageSetA", log); err != nil {
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

// ParseStorageSetA is a log parse operation binding the contract event 0xe394429e5515af0c383c3d1fc989369cfcce852ba8bc8300b01b64bfdb3d18be.
//
// Solidity: event StorageSetA(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) ParseStorageSetA(log types.Log) (*SimpleStorageStorageSetA, error) {
	event := new(SimpleStorageStorageSetA)
	if err := _SimpleStorage.contract.UnpackLog(event, "StorageSetA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleStorageStorageSetBIterator is returned from FilterStorageSetB and is used to iterate over the raw logs and unpacked data for StorageSetB events raised by the SimpleStorage contract.
type SimpleStorageStorageSetBIterator struct {
	Event *SimpleStorageStorageSetB // Event containing the contract specifics and raw log

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
func (it *SimpleStorageStorageSetBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageStorageSetB)
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
		it.Event = new(SimpleStorageStorageSetB)
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
func (it *SimpleStorageStorageSetBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageStorageSetBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageStorageSetB represents a StorageSetB event raised by the SimpleStorage contract.
type SimpleStorageStorageSetB struct {
	NewString string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStorageSetB is a free log retrieval operation binding the contract event 0xffad11ecb618b0f8f566637fe9991a6c3ec300a88293e7e9e68ddcee51b861b5.
//
// Solidity: event StorageSetB(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) FilterStorageSetB(opts *bind.FilterOpts) (*SimpleStorageStorageSetBIterator, error) {

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "StorageSetB")
	if err != nil {
		return nil, err
	}
	return &SimpleStorageStorageSetBIterator{contract: _SimpleStorage.contract, event: "StorageSetB", logs: logs, sub: sub}, nil
}

// WatchStorageSetB is a free log subscription operation binding the contract event 0xffad11ecb618b0f8f566637fe9991a6c3ec300a88293e7e9e68ddcee51b861b5.
//
// Solidity: event StorageSetB(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) WatchStorageSetB(opts *bind.WatchOpts, sink chan<- *SimpleStorageStorageSetB) (event.Subscription, error) {

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "StorageSetB")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageStorageSetB)
				if err := _SimpleStorage.contract.UnpackLog(event, "StorageSetB", log); err != nil {
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

// ParseStorageSetB is a log parse operation binding the contract event 0xffad11ecb618b0f8f566637fe9991a6c3ec300a88293e7e9e68ddcee51b861b5.
//
// Solidity: event StorageSetB(string NewString)
func (_SimpleStorage *SimpleStorageFilterer) ParseStorageSetB(log types.Log) (*SimpleStorageStorageSetB, error) {
	event := new(SimpleStorageStorageSetB)
	if err := _SimpleStorage.contract.UnpackLog(event, "StorageSetB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
