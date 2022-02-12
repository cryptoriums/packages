// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package booster

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122009f4f922fb7452696e9f3fbc93b0a6c25c876029e4f095334664147d52b6cc3f64736f6c634300060c0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BoosterMetaData contains all meta data concerning the Booster contract.
var BoosterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaxFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stashVersion\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"depositAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionAddressId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earmarkFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earmarkIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"earmarkRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistro\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crvRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stash\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardArbitrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arb\",\"type\":\"address\"}],\"name\":\"setArbitrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tfactory\",\"type\":\"address\"}],\"name\":\"setFactories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeM\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platform\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"setGaugeRedirect\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolM\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakerRewards\",\"type\":\"address\"}],\"name\":\"setRewardContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voteDelegate\",\"type\":\"address\"}],\"name\":\"setVoteDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"shutdownPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stashFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_votingAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_support\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauge\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weight\",\"type\":\"uint256[]\"}],\"name\":\"voteGaugeWeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteOwnership\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteParameter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d73792a9": "FEE_DENOMINATOR()",
		"7303df9a": "MaxFees()",
		"7e29d6c2": "addPool(address,address,uint256)",
		"6c7b69cb": "claimRewards(uint256,address)",
		"6a4874a1": "crv()",
		"43a0d066": "deposit(uint256,uint256,bool)",
		"60759fce": "depositAll(uint256,bool)",
		"93e846a0": "distributionAddressId()",
		"22230b96": "earmarkFees()",
		"3a088cd2": "earmarkIncentive()",
		"cc956f3f": "earmarkRewards(uint256)",
		"d6a0f530": "feeDistro()",
		"d0fb0203": "feeManager()",
		"647846a5": "feeToken()",
		"cb0d5b52": "gaugeMap(address)",
		"bf86d690": "isShutdown()",
		"ab366292": "lockFees()",
		"50940618": "lockIncentive()",
		"376d771a": "lockRewards()",
		"07546172": "minter()",
		"8da5cb5b": "owner()",
		"26232a2e": "platformFee()",
		"1526fe27": "poolInfo(uint256)",
		"081e3eda": "poolLength()",
		"dc4c90d3": "poolManager()",
		"7b103999": "registry()",
		"043b684a": "rewardArbitrator()",
		"71192b17": "rewardClaimed(uint256,address,uint256)",
		"245e4bf0": "rewardFactory()",
		"b0eefabe": "setArbitrator(address)",
		"7bd3b995": "setFactories(address,address,address)",
		"5a4ae5ca": "setFeeInfo()",
		"472d35b9": "setFeeManager(address)",
		"6fcba377": "setFees(uint256,uint256,uint256,uint256)",
		"9123d404": "setGaugeRedirect(uint256)",
		"13af4035": "setOwner(address)",
		"7aef6715": "setPoolManager(address)",
		"95539a1d": "setRewardContracts(address,address)",
		"f0f44260": "setTreasury(address)",
		"74874323": "setVoteDelegate(address)",
		"60cafe84": "shutdownPool(uint256)",
		"354af919": "shutdownSystem()",
		"5ebaf1db": "staker()",
		"62d28ac7": "stakerIncentive()",
		"cfb9cfba": "stakerRewards()",
		"068eb19e": "stashFactory()",
		"e77772fe": "tokenFactory()",
		"61d027b3": "treasury()",
		"e2cdd42a": "vote(uint256,address,bool)",
		"9f00332b": "voteDelegate()",
		"bfad96ba": "voteGaugeWeight(address[],uint256[])",
		"a386a080": "voteOwnership()",
		"b42eda71": "voteParameter()",
		"441a3e70": "withdraw(uint256,uint256)",
		"958e2d31": "withdrawAll(uint256)",
		"14cd70e4": "withdrawTo(uint256,uint256,address)",
	},
	Bin: "0x60c06040526103e86000556101c26001556032600255600060035534801561002657600080fd5b506040516139733803806139738339818101604052604081101561004957600080fd5b508051602090910151601180546001600160601b0319606094851b8116608052600480546001600160a01b031990811633908117909255600b80548216831790556005805482168317905560068054821690921790915560108054821690556001600160a81b0319909216909255600c80549091169055911b1660a05260805160601c60a05160601c61383e6101356000398061098452806119bb525080610b295280610d7e5280610eef5280610f6752806114e5528061164f528061178f5280611dd8528061205b528061224a528061264b52806128965280612a3c5280612ea9525061383e6000f3fe608060405234801561001057600080fd5b50600436106103785760003560e01c80637303df9a116101d3578063b0eefabe11610104578063d0fb0203116100a2578063dc4c90d31161007c578063dc4c90d3146108fa578063e2cdd42a14610902578063e77772fe14610936578063f0f442601461093e57610378565b8063d0fb0203146108e2578063d6a0f530146108ea578063d73792a9146108f257610378565b8063bfad96ba116100de578063bfad96ba146107d5578063cb0d5b5214610897578063cc956f3f146108bd578063cfb9cfba146108da57610378565b8063b0eefabe1461079f578063b42eda71146107c5578063bf86d690146107cd57610378565b80639123d40411610171578063958e2d311161014b578063958e2d311461076a5780639f00332b14610787578063a386a0801461078f578063ab3662921461079757610378565b80639123d4041461071757806393e846a01461073457806395539a1d1461073c57610378565b80637b103999116101ad5780637b103999146106995780637bd3b995146106a15780637e29d6c2146106d95780638da5cb5b1461070f57610378565b80637303df9a14610645578063748743231461064d5780637aef67151461067357610378565b8063441a3e70116102ad57806361d027b31161024b5780636a4874a1116102255780636a4874a1146105b05780636c7b69cb146105b85780636fcba377146105e457806371192b171461061357610378565b806361d027b31461059857806362d28ac7146105a0578063647846a5146105a857610378565b80635a4ae5ca116102875780635a4ae5ca146105465780635ebaf1db1461054e57806360759fce1461055657806360cafe841461057b57610378565b8063441a3e70146104f5578063472d35b914610518578063509406181461053e57610378565b806322230b961161031a578063354af919116102f4578063354af919146104b2578063376d771a146104ba5780633a088cd2146104c257806343a0d066146104ca57610378565b806322230b961461049a578063245e4bf0146104a257806326232a2e146104aa57610378565b8063081e3eda11610356578063081e3eda146103b157806313af4035146103cb57806314cd70e4146103f35780631526fe271461043957610378565b8063043b684a1461037d578063068eb19e146103a157806307546172146103a9575b600080fd5b610385610964565b604080516001600160a01b039092168252519081900360200190f35b610385610973565b610385610982565b6103b96109a6565b60408051918252519081900360200190f35b6103f1600480360360208110156103e157600080fd5b50356001600160a01b03166109ac565b005b6104256004803603606081101561040957600080fd5b50803590602081013590604001356001600160a01b0316610a15565b604080519115158252519081900360200190f35b6104566004803603602081101561044f57600080fd5b5035610a99565b604080516001600160a01b0397881681529587166020870152938616858501529185166060850152909316608083015291151560a082015290519081900360c00190f35b610425610af7565b610385610c8d565b6103b9610c9c565b6103f1610ca2565b610385610e07565b6103b9610e16565b610425600480360360608110156104e057600080fd5b50803590602081013590604001351515610e1c565b6104256004803603604081101561050b57600080fd5b508035906020013561123f565b6103f16004803603602081101561052e57600080fd5b50356001600160a01b0316611257565b6103b96112c0565b6103f16112c6565b6103856114e3565b6104256004803603604081101561056c57600080fd5b50803590602001351515611507565b6104256004803603602081101561059157600080fd5b50356115b4565b6103856116e6565b6103b96116f5565b6103856116fb565b61038561170a565b610425600480360360408110156105ce57600080fd5b50803590602001356001600160a01b0316611722565b6103f1600480360360808110156105fa57600080fd5b508035906020810135906040810135906060013561181e565b6104256004803603606081101561062957600080fd5b508035906001600160a01b036020820135169060400135611939565b6103b9611a53565b6103f16004803603602081101561066357600080fd5b50356001600160a01b0316611a59565b6103f16004803603602081101561068957600080fd5b50356001600160a01b0316611ac2565b610385611b2b565b6103f1600480360360608110156106b757600080fd5b506001600160a01b038135811691602081013582169160409091013516611b3f565b610425600480360360608110156106ef57600080fd5b506001600160a01b03813581169160208101359091169060400135611beb565b610385612135565b6104256004803603602081101561072d57600080fd5b5035612144565b6103b96123ed565b6103f16004803603604081101561075257600080fd5b506001600160a01b03813581169160200135166123f2565b6104256004803603602081101561078057600080fd5b503561247e565b61038561252d565b61038561253c565b610385612554565b6103f1600480360360208110156107b557600080fd5b50356001600160a01b0316612563565b6103856125cc565b6104256125e4565b610425600480360360408110156107eb57600080fd5b81019060208101813564010000000081111561080657600080fd5b82018360208201111561081857600080fd5b8035906020019184602083028401116401000000008311171561083a57600080fd5b91939092909160208101903564010000000081111561085857600080fd5b82018360208201111561086a57600080fd5b8035906020019184602083028401116401000000008311171561088c57600080fd5b5090925090506125f4565b610425600480360360208110156108ad57600080fd5b50356001600160a01b0316612711565b610425600480360360208110156108d357600080fd5b5035612726565b610385612784565b610385612793565b6103856127a2565b6103b96127b1565b6103856127b7565b6104256004803603606081101561091857600080fd5b508035906001600160a01b03602082013516906040013515156127c6565b610385612915565b6103f16004803603602081101561095457600080fd5b50356001600160a01b0316612924565b600a546001600160a01b031681565b6008546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60125490565b6004546001600160a01b031633146109f3576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055565b60008060128581548110610a2557fe5b60009182526020909120600360059092020101546001600160a01b03169050338114610a80576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b610a8c8585338661298d565b60019150505b9392505050565b60128181548110610aa657fe5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001600160a01b03938416955091831693908316929081169190811690600160a01b900460ff1686565b60105460115460408051632dbfa73560e01b81526001600160a01b0393841660048201529183166024830152516000927f00000000000000000000000000000000000000000000000000000000000000001691632dbfa735916044808301928692919082900301818387803b158015610b6f57600080fd5b505af1158015610b83573d6000803e3d6000fd5b5050601154604080516370a0823160e01b81523060048201529051600094506001600160a01b0390921692506370a08231916024808301926020929190829003018186803b158015610bd457600080fd5b505afa158015610be8573d6000803e3d6000fd5b505050506040513d6020811015610bfe57600080fd5b5051600f54601154919250610c20916001600160a01b03908116911683612be6565b600f546040805163590a41f560e01b81526004810184905290516001600160a01b039092169163590a41f59160248082019260009290919082900301818387803b158015610c6d57600080fd5b505af1158015610c81573d6000803e3d6000fd5b50505050600191505090565b6007546001600160a01b031681565b60035481565b6004546001600160a01b03163314610ce9576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b6011805460ff60a01b1916600160a01b17905560005b601254811015610e0457600060128281548110610d1857fe5b906000526020600020906005020190508060040160149054906101000a900460ff1615610d455750610dfc565b80546002820154604080516301395c5960e31b81526001600160a01b0393841660048201819052928416602482018190529151929391927f0000000000000000000000000000000000000000000000000000000000000000909216916309cae2c89160448082019260009290919082900301818387803b158015610dc857600080fd5b505af1925050508015610dd9575060015b610de257610df8565b60048301805460ff60a01b1916600160a01b1790555b5050505b600101610cff565b50565b600e546001600160a01b031681565b60025481565b601154600090600160a01b900460ff1615610e69576040805162461bcd60e51b815260206004820152600860248201526739b43aba3237bbb760c11b604482015290519081900360640190fd5b600060128581548110610e7857fe5b600091825260209091206005909102016004810154909150600160a01b900460ff1615610edd576040805162461bcd60e51b815260206004820152600e60248201526d1c1bdbdb081a5cc818db1bdcd95960921b604482015290519081900360640190fd5b80546001600160a01b0316610f1481337f000000000000000000000000000000000000000000000000000000000000000088612c3d565b60028201546001600160a01b031680610f65576040805162461bcd60e51b815260206004820152600e60248201526d2167617567652073657474696e6760901b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f9609f0883836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b0316815260200192505050600060405180830381600087803b158015610fe557600080fd5b505af1158015610ff9573d6000803e3d6000fd5b5050505060048301546001600160a01b0316801561107957806001600160a01b031663b87bd4816040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561104c57600080fd5b505af1158015611060573d6000803e3d6000fd5b505050506040513d602081101561107657600080fd5b50505b60018401546001600160a01b0316861561119257604080516340c10f1960e01b8152306004820152602481018a905290516001600160a01b038316916340c10f1991604480830192600092919082900301818387803b1580156110db57600080fd5b505af11580156110ef573d6000803e3d6000fd5b50505060038601546001600160a01b039081169150611112908316826000612c9d565b6111266001600160a01b038316828b612c9d565b604080516305dc812160e31b8152336004820152602481018b905290516001600160a01b03831691632ee4090891604480830192600092919082900301818387803b15801561117457600080fd5b505af1158015611188573d6000803e3d6000fd5b50505050506111f9565b604080516340c10f1960e01b8152336004820152602481018a905290516001600160a01b038316916340c10f1991604480830192600092919082900301818387803b1580156111e057600080fd5b505af11580156111f4573d6000803e3d6000fd5b505050505b6040805189815290518a9133917f73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca9181900360200190a350600198975050505050505050565b600061124d8383333361298d565b5060015b92915050565b6005546001600160a01b0316331461129e576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b60005481565b6005546001600160a01b0316331461130d576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b6040805163124fd3dd60e21b815260048181015290516f22d53366457f9d5e68ec105046fc43839163493f4f74916024808301926020929190829003018186803b15801561135a57600080fd5b505afa15801561136e573d6000803e3d6000fd5b505050506040513d602081101561138457600080fd5b5051601080546001600160a01b0319166001600160a01b03928316179081905560408051637e062a3560e11b81529051600093929092169163fc0c546a91600480820192602092909190829003018186803b1580156113e257600080fd5b505afa1580156113f6573d6000803e3d6000fd5b505050506040513d602081101561140c57600080fd5b50516011549091506001600160a01b03808316911614610e0457600754600e5460408051637c6b091760e11b81526001600160a01b03858116600483015292831660248201523060448201529051919092169163f8d6122e9160648083019260209291908290030181600087803b15801561148657600080fd5b505af115801561149a573d6000803e3d6000fd5b505050506040513d60208110156114b057600080fd5b5051600f80546001600160a01b039283166001600160a01b03199182161790915560118054939092169216919091179055565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000806012848154811061151757fe5b60009182526020808320600590920290910154604080516370a0823160e01b815233600482015290516001600160a01b03909216945084926370a0823192602480840193829003018186803b15801561156f57600080fd5b505afa158015611583573d6000803e3d6000fd5b505050506040513d602081101561159957600080fd5b505190506115a8858286610e1c565b50600195945050505050565b6006546000906001600160a01b031633146115fe576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b60006012838154811061160d57fe5b60009182526020822060059091020180546002820154604080516301395c5960e31b81526001600160a01b0393841660048201529183166024830152519294507f000000000000000000000000000000000000000000000000000000000000000091909116926309cae2c89260448084019382900301818387803b15801561169457600080fd5b505af19250505080156116a5575060015b5060048101805460ff60a01b1916600160a01b179055600201546001600160a01b03166000908152601360205260409020805460ff19169055506001919050565b600c546001600160a01b031681565b60015481565b6011546001600160a01b031681565b73d533a949740bb3306d119cc777fa900ba034cd5281565b6000806012848154811061173257fe5b60009182526020909120600460059092020101546001600160a01b0316905033811461178d576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ef5cfb8c846040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156117fc57600080fd5b505af1158015611810573d6000803e3d6000fd5b506001979650505050505050565b6005546001600160a01b03163314611865576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600061187d8261187785818989612db0565b90612db0565b90506107d08111156118c1576040805162461bcd60e51b81526020600482015260086024820152673e4d61784665657360c01b604482015290519081900360640190fd5b6103e885101580156118d557506105dc8511155b80156118e3575061012c8410155b80156118f157506102588411155b80156118fe5750600a8310155b801561190b575060648311155b8015611918575060c88211155b156119325760008590556001849055600283905560038290555b5050505050565b6000806012858154811061194957fe5b60009182526020909120600360059092020101546001600160a01b03169050338114806119805750600e546001600160a01b031633145b6119b9576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166340c10f1985856040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015611a3057600080fd5b505af1158015611a44573d6000803e3d6000fd5b50600198975050505050505050565b6107d081565b600b546001600160a01b03163314611aa0576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6006546001600160a01b03163314611b09576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6f22d53366457f9d5e68ec105046fc438381565b6004546001600160a01b03163314611b86576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b6007546001600160a01b0316611bc757600780546001600160a01b038086166001600160a01b03199283161790925560098054928416929091169190911790555b50600880546001600160a01b0319166001600160a01b039290921691909117905550565b6006546000906001600160a01b031633148015611c125750601154600160a01b900460ff16155b611c4c576040805162461bcd60e51b815260206004808301919091526024820152630858591960e21b604482015290519081900360640190fd5b6001600160a01b03831615801590611c6c57506001600160a01b03841615155b611ca6576040805162461bcd60e51b815260206004820152600660248201526521706172616d60d01b604482015290519081900360640190fd5b60125460095460408051630452a26760e21b81526001600160a01b0388811660048301529151600093929092169163114a899c9160248082019260209290919082900301818787803b158015611cfb57600080fd5b505af1158015611d0f573d6000803e3d6000fd5b505050506040513d6020811015611d2557600080fd5b5051600754604080516358cbfd4560e01b8152600481018690526001600160a01b038085166024830152915193945060009391909216916358cbfd4591604480830192602092919082900301818787803b158015611d8257600080fd5b505af1158015611d96573d6000803e3d6000fd5b505050506040513d6020811015611dac57600080fd5b505160085460408051634ce5896f60e11b8152600481018790526001600160a01b038a811660248301527f000000000000000000000000000000000000000000000000000000000000000081166044830152606482018a9052915193945060009391909216916399cb12de91608480830192602092919082900301818787803b158015611e3857600080fd5b505af1158015611e4c573d6000803e3d6000fd5b505050506040513d6020811015611e6257600080fd5b50516040805160c0810182526001600160a01b03808c16825286811660208381019182528c83168486018181528985166060870190815285891660808801818152600060a08a0181815260128054600181810183559184529b516005909c027fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec3444810180549d8d166001600160a01b03199e8f1617905599517fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec34458b018054918d16918e1691909117905595517fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec34468a018054918c16918d1691909117905593517fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec344789018054918b16918c1691909117905590517fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec3448909701805493511515600160a01b0260ff60a01b199890991693909916929092179590951695909517909555835260139052929020805460ff191690911790559091501561181057806012858154811061200a57fe5b6000918252602082206004600590920201810180546001600160a01b039485166001600160a01b031990911617905560408051637d1cb25960e11b81528585169281019290925260016024830152517f00000000000000000000000000000000000000000000000000000000000000009093169263fa3964b29260448084019391929182900301818387803b1580156120a257600080fd5b505af11580156120b6573d6000803e3d6000fd5b50506007546040805163b84614a560e01b81526001600160a01b03868116600483015260016024830152915191909216935063b84614a59250604480830192600092919082900301818387803b15801561210f57600080fd5b505af1158015612123573d6000803e3d6000fd5b50505050506001979650505050505050565b6004546001600160a01b031681565b6000806012838154811061215457fe5b60009182526020909120600460059092020101546001600160a01b031690503381146121af576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b6000601284815481106121be57fe5b6000918252602080832060026005909302019190910154604080516001600160a01b038781166024808401919091528351808403820181526044938401855295860180516001600160e01b0316635efcc08b60e11b1781529351635b0e93fb60e11b815294821660048601818152918601889052606093860193845286516064870152865190985095967f00000000000000000000000000000000000000000000000000000000000000009092169563b61d27f6958995939489949092608490920191808383895b8381101561229e578181015183820152602001612286565b50505050905090810190601f1680156122cb5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156122ec57600080fd5b505af1158015612300573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561232957600080fd5b81516020830180516040519294929383019291908464010000000082111561235057600080fd5b90830190602082018581111561236557600080fd5b825164010000000081118282018810171561237f57600080fd5b82525081516020918201929091019080838360005b838110156123ac578181015183820152602001612394565b50505050905090810190601f1680156123d95780820380516001836020036101000a031916815260200191505b506040525060019998505050505050505050565b600481565b6004546001600160a01b03163314612439576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600e546001600160a01b031661247a57600e80546001600160a01b038085166001600160a01b031992831617909255600d8054928416929091169190911790555b5050565b6000806012838154811061248e57fe5b6000918252602080832060016005909302019190910154604080516370a0823160e01b815233600482015290516001600160a01b03909216945084926370a0823192602480840193829003018186803b1580156124ea57600080fd5b505afa1580156124fe573d6000803e3d6000fd5b505050506040513d602081101561251457600080fd5b50519050612522848261123f565b506001949350505050565b600b546001600160a01b031681565b73e478de485ad2fe566d49342cbd03e49ed7db335681565b600f546001600160a01b031681565b6004546001600160a01b031633146125aa576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b73bcff8b0b9419b9a88c44546519b1e909cf33039981565b601154600160a01b900460ff1681565b600b546000906001600160a01b0316331461263e576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b60005b848110156115a8577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635d7e9bcb87878481811061268457fe5b905060200201356001600160a01b03168686858181106126a057fe5b905060200201356040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156126ed57600080fd5b505af1158015612701573d6000803e3d6000fd5b5050600190920191506126419050565b60136020526000908152604090205460ff1681565b601154600090600160a01b900460ff1615612773576040805162461bcd60e51b815260206004820152600860248201526739b43aba3237bbb760c11b604482015290519081900360640190fd5b61277c82612e0a565b506001919050565b600d546001600160a01b031681565b6005546001600160a01b031681565b6010546001600160a01b031681565b61271081565b6006546001600160a01b031681565b600b546000906001600160a01b03163314612810576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b6001600160a01b03831673e478de485ad2fe566d49342cbd03e49ed7db3356148061285757506001600160a01b03831673bcff8b0b9419b9a88c44546519b1e909cf330399145b612894576040805162461bcd60e51b815260206004820152600960248201526810bb37ba32a0b2323960b91b604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e2cdd42a8585856040518463ffffffff1660e01b815260040180848152602001836001600160a01b0316815260200182151581526020019350505050600060405180830381600087803b1580156117fc57600080fd5b6009546001600160a01b031681565b6005546001600160a01b0316331461296b576040805162461bcd60e51b8152602060048201526005602482015264042c2eae8d60db1b604482015290519081900360640190fd5b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b60006012858154811061299c57fe5b60009182526020822060059091020180546002820154600183015460408051632770a7eb60e21b81526001600160a01b038a81166004830152602482018c9052915195975093811695928116949116928392639dc29fac9260448084019382900301818387803b158015612a0f57600080fd5b505af1158015612a23573d6000803e3d6000fd5b505050506004840154600160a01b900460ff16612adb577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d9caed1284848a6040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b031681526020018281526020019350505050600060405180830381600087803b158015612ac257600080fd5b505af1158015612ad6573d6000803e3d6000fd5b505050505b60048401546001600160a01b03168015801590612b025750601154600160a01b900460ff16155b8015612b1a57506004850154600160a01b900460ff16155b15612b8757806001600160a01b031663b87bd4816040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612b5a57600080fd5b505af1158015612b6e573d6000803e3d6000fd5b505050506040513d6020811015612b8457600080fd5b50505b612b9b6001600160a01b038516878a612be6565b6040805189815290518a916001600160a01b038916917f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc69181900360200190a3505050505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052612c38908490613376565b505050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052612c97908590613376565b50505050565b801580612d23575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b158015612cf557600080fd5b505afa158015612d09573d6000803e3d6000fd5b505050506040513d6020811015612d1f57600080fd5b5051155b612d5e5760405162461bcd60e51b81526004018080602001828103825260368152602001806137d36036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052612c38908490613376565b600082820183811015610a92576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600060128281548110612e1957fe5b600091825260209091206005909102016004810154909150600160a01b900460ff1615612e7e576040805162461bcd60e51b815260206004820152600e60248201526d1c1bdbdb081a5cc818db1bdcd95960921b604482015290519081900360640190fd5b600281015460408051631ff4de0360e11b81526001600160a01b0392831660048201819052915191927f00000000000000000000000000000000000000000000000000000000000000001691633fe9bc06916024808201926020929091908290030181600087803b158015612ef257600080fd5b505af1158015612f06573d6000803e3d6000fd5b505050506040513d6020811015612f1c57600080fd5b505060048201546001600160a01b0316801561300857806001600160a01b031663372500ab6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015612f6d57600080fd5b505af1158015612f81573d6000803e3d6000fd5b505050506040513d6020811015612f9757600080fd5b50506040805163654580bb60e11b815290516001600160a01b0383169163ca8b01769160048083019260209291908290030181600087803b158015612fdb57600080fd5b505af1158015612fef573d6000803e3d6000fd5b505050506040513d602081101561300557600080fd5b50505b604080516370a0823160e01b8152306004820152905160009173d533a949740bb3306d119cc777fa900ba034cd52916370a0823191602480820192602092909190829003018186803b15801561305d57600080fd5b505afa158015613071573d6000803e3d6000fd5b505050506040513d602081101561308757600080fd5b5051905080156119325760006130b46127106130ae6000548561342790919063ffffffff16565b90613480565b905060006130d36127106130ae6001548661342790919063ffffffff16565b905060006130f26127106130ae6002548761342790919063ffffffff16565b600c549091506001600160a01b03161580159061311a5750600c546001600160a01b03163014155b801561312857506000600354115b1561318757600061314a6127106130ae6003548861342790919063ffffffff16565b905061315685826134e7565b600c549095506131859073d533a949740bb3306d119cc777fa900ba034cd52906001600160a01b031683612be6565b505b61319d82613197838188886134e7565b906134e7565b93506131be73d533a949740bb3306d119cc777fa900ba034cd523383612be6565b60038701546001600160a01b03166131eb73d533a949740bb3306d119cc777fa900ba034cd528287612be6565b806001600160a01b031663590a41f5866040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561323157600080fd5b505af1158015613245573d6000803e3d6000fd5b5050600e54613275925073d533a949740bb3306d119cc777fa900ba034cd5291506001600160a01b031686612be6565b600e546040805163590a41f560e01b81526004810187905290516001600160a01b039092169163590a41f59160248082019260009290919082900301818387803b1580156132c257600080fd5b505af11580156132d6573d6000803e3d6000fd5b5050600d54613306925073d533a949740bb3306d119cc777fa900ba034cd5291506001600160a01b031685612be6565b600d546040805163590a41f560e01b81526004810186905290516001600160a01b039092169163590a41f59160248082019260009290919082900301818387803b15801561335357600080fd5b505af1158015613367573d6000803e3d6000fd5b50505050505050505050505050565b60606133cb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166135449092919063ffffffff16565b805190915015612c38578080602001905160208110156133ea57600080fd5b5051612c385760405162461bcd60e51b815260040180806020018281038252602a8152602001806137a9602a913960400191505060405180910390fd5b60008261343657506000611251565b8282028284828161344357fe5b0414610a925760405162461bcd60e51b81526004018080602001828103825260218152602001806137886021913960400191505060405180910390fd5b60008082116134d6576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b8183816134df57fe5b049392505050565b60008282111561353e576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6060613553848460008561355b565b949350505050565b60608247101561359c5760405162461bcd60e51b81526004018080602001828103825260268152602001806137626026913960400191505060405180910390fd5b6135a5856136b7565b6135f6576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106136355780518252601f199092019160209182019101613616565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613697576040519150601f19603f3d011682016040523d82523d6000602084013e61369c565b606091505b50915091506136ac8282866136bd565b979650505050505050565b3b151590565b606083156136cc575081610a92565b8251156136dc5782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561372657818101518382015260200161370e565b50505050905090810190601f1680156137535780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565645361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a264697066735822122070205b64b16cdd1bc42412596470a0860eec23572972d8129388c0fb2febe39c64736f6c634300060c0033",
}

// BoosterABI is the input ABI used to generate the binding from.
// Deprecated: Use BoosterMetaData.ABI instead.
var BoosterABI = BoosterMetaData.ABI

// Deprecated: Use BoosterMetaData.Sigs instead.
// BoosterFuncSigs maps the 4-byte function signature to its string representation.
var BoosterFuncSigs = BoosterMetaData.Sigs

// BoosterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BoosterMetaData.Bin instead.
var BoosterBin = BoosterMetaData.Bin

// DeployBooster deploys a new Ethereum contract, binding an instance of Booster to it.
func DeployBooster(auth *bind.TransactOpts, backend bind.ContractBackend, _staker common.Address, _minter common.Address) (common.Address, *types.Transaction, *Booster, error) {
	parsed, err := BoosterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BoosterBin), backend, _staker, _minter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Booster{BoosterCaller: BoosterCaller{contract: contract}, BoosterTransactor: BoosterTransactor{contract: contract}, BoosterFilterer: BoosterFilterer{contract: contract}}, nil
}

// Booster is an auto generated Go binding around an Ethereum contract.
type Booster struct {
	BoosterCaller     // Read-only binding to the contract
	BoosterTransactor // Write-only binding to the contract
	BoosterFilterer   // Log filterer for contract events
}

// BoosterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoosterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoosterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoosterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoosterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoosterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoosterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoosterSession struct {
	Contract     *Booster          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoosterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoosterCallerSession struct {
	Contract *BoosterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BoosterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoosterTransactorSession struct {
	Contract     *BoosterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BoosterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoosterRaw struct {
	Contract *Booster // Generic contract binding to access the raw methods on
}

// BoosterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoosterCallerRaw struct {
	Contract *BoosterCaller // Generic read-only contract binding to access the raw methods on
}

// BoosterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoosterTransactorRaw struct {
	Contract *BoosterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBooster creates a new instance of Booster, bound to a specific deployed contract.
func NewBooster(address common.Address, backend bind.ContractBackend) (*Booster, error) {
	contract, err := bindBooster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Booster{BoosterCaller: BoosterCaller{contract: contract}, BoosterTransactor: BoosterTransactor{contract: contract}, BoosterFilterer: BoosterFilterer{contract: contract}}, nil
}

// NewBoosterCaller creates a new read-only instance of Booster, bound to a specific deployed contract.
func NewBoosterCaller(address common.Address, caller bind.ContractCaller) (*BoosterCaller, error) {
	contract, err := bindBooster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoosterCaller{contract: contract}, nil
}

// NewBoosterTransactor creates a new write-only instance of Booster, bound to a specific deployed contract.
func NewBoosterTransactor(address common.Address, transactor bind.ContractTransactor) (*BoosterTransactor, error) {
	contract, err := bindBooster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoosterTransactor{contract: contract}, nil
}

// NewBoosterFilterer creates a new log filterer instance of Booster, bound to a specific deployed contract.
func NewBoosterFilterer(address common.Address, filterer bind.ContractFilterer) (*BoosterFilterer, error) {
	contract, err := bindBooster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoosterFilterer{contract: contract}, nil
}

// bindBooster binds a generic wrapper to an already deployed contract.
func bindBooster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoosterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Booster *BoosterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Booster.Contract.BoosterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Booster *BoosterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Booster.Contract.BoosterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Booster *BoosterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Booster.Contract.BoosterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Booster *BoosterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Booster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Booster *BoosterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Booster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Booster *BoosterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Booster.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Booster *BoosterCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Booster *BoosterSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Booster.Contract.FEEDENOMINATOR(&_Booster.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Booster *BoosterCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Booster.Contract.FEEDENOMINATOR(&_Booster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_Booster *BoosterCaller) MaxFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "MaxFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_Booster *BoosterSession) MaxFees() (*big.Int, error) {
	return _Booster.Contract.MaxFees(&_Booster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_Booster *BoosterCallerSession) MaxFees() (*big.Int, error) {
	return _Booster.Contract.MaxFees(&_Booster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_Booster *BoosterCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_Booster *BoosterSession) Crv() (common.Address, error) {
	return _Booster.Contract.Crv(&_Booster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_Booster *BoosterCallerSession) Crv() (common.Address, error) {
	return _Booster.Contract.Crv(&_Booster.CallOpts)
}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_Booster *BoosterCaller) DistributionAddressId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "distributionAddressId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_Booster *BoosterSession) DistributionAddressId() (*big.Int, error) {
	return _Booster.Contract.DistributionAddressId(&_Booster.CallOpts)
}

// DistributionAddressId is a free data retrieval call binding the contract method 0x93e846a0.
//
// Solidity: function distributionAddressId() view returns(uint256)
func (_Booster *BoosterCallerSession) DistributionAddressId() (*big.Int, error) {
	return _Booster.Contract.DistributionAddressId(&_Booster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_Booster *BoosterCaller) EarmarkIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "earmarkIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_Booster *BoosterSession) EarmarkIncentive() (*big.Int, error) {
	return _Booster.Contract.EarmarkIncentive(&_Booster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_Booster *BoosterCallerSession) EarmarkIncentive() (*big.Int, error) {
	return _Booster.Contract.EarmarkIncentive(&_Booster.CallOpts)
}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_Booster *BoosterCaller) FeeDistro(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "feeDistro")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_Booster *BoosterSession) FeeDistro() (common.Address, error) {
	return _Booster.Contract.FeeDistro(&_Booster.CallOpts)
}

// FeeDistro is a free data retrieval call binding the contract method 0xd6a0f530.
//
// Solidity: function feeDistro() view returns(address)
func (_Booster *BoosterCallerSession) FeeDistro() (common.Address, error) {
	return _Booster.Contract.FeeDistro(&_Booster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Booster *BoosterCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Booster *BoosterSession) FeeManager() (common.Address, error) {
	return _Booster.Contract.FeeManager(&_Booster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Booster *BoosterCallerSession) FeeManager() (common.Address, error) {
	return _Booster.Contract.FeeManager(&_Booster.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Booster *BoosterCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Booster *BoosterSession) FeeToken() (common.Address, error) {
	return _Booster.Contract.FeeToken(&_Booster.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Booster *BoosterCallerSession) FeeToken() (common.Address, error) {
	return _Booster.Contract.FeeToken(&_Booster.CallOpts)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_Booster *BoosterCaller) GaugeMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "gaugeMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_Booster *BoosterSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _Booster.Contract.GaugeMap(&_Booster.CallOpts, arg0)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_Booster *BoosterCallerSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _Booster.Contract.GaugeMap(&_Booster.CallOpts, arg0)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Booster *BoosterCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Booster *BoosterSession) IsShutdown() (bool, error) {
	return _Booster.Contract.IsShutdown(&_Booster.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Booster *BoosterCallerSession) IsShutdown() (bool, error) {
	return _Booster.Contract.IsShutdown(&_Booster.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_Booster *BoosterCaller) LockFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "lockFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_Booster *BoosterSession) LockFees() (common.Address, error) {
	return _Booster.Contract.LockFees(&_Booster.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_Booster *BoosterCallerSession) LockFees() (common.Address, error) {
	return _Booster.Contract.LockFees(&_Booster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_Booster *BoosterCaller) LockIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "lockIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_Booster *BoosterSession) LockIncentive() (*big.Int, error) {
	return _Booster.Contract.LockIncentive(&_Booster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_Booster *BoosterCallerSession) LockIncentive() (*big.Int, error) {
	return _Booster.Contract.LockIncentive(&_Booster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_Booster *BoosterCaller) LockRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "lockRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_Booster *BoosterSession) LockRewards() (common.Address, error) {
	return _Booster.Contract.LockRewards(&_Booster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_Booster *BoosterCallerSession) LockRewards() (common.Address, error) {
	return _Booster.Contract.LockRewards(&_Booster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Booster *BoosterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Booster *BoosterSession) Minter() (common.Address, error) {
	return _Booster.Contract.Minter(&_Booster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Booster *BoosterCallerSession) Minter() (common.Address, error) {
	return _Booster.Contract.Minter(&_Booster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Booster *BoosterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Booster *BoosterSession) Owner() (common.Address, error) {
	return _Booster.Contract.Owner(&_Booster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Booster *BoosterCallerSession) Owner() (common.Address, error) {
	return _Booster.Contract.Owner(&_Booster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Booster *BoosterCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Booster *BoosterSession) PlatformFee() (*big.Int, error) {
	return _Booster.Contract.PlatformFee(&_Booster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Booster *BoosterCallerSession) PlatformFee() (*big.Int, error) {
	return _Booster.Contract.PlatformFee(&_Booster.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_Booster *BoosterCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Lptoken    common.Address
		Token      common.Address
		Gauge      common.Address
		CrvRewards common.Address
		Stash      common.Address
		Shutdown   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lptoken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gauge = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CrvRewards = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Stash = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Shutdown = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_Booster *BoosterSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _Booster.Contract.PoolInfo(&_Booster.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_Booster *BoosterCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _Booster.Contract.PoolInfo(&_Booster.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Booster *BoosterCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Booster *BoosterSession) PoolLength() (*big.Int, error) {
	return _Booster.Contract.PoolLength(&_Booster.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Booster *BoosterCallerSession) PoolLength() (*big.Int, error) {
	return _Booster.Contract.PoolLength(&_Booster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Booster *BoosterCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Booster *BoosterSession) PoolManager() (common.Address, error) {
	return _Booster.Contract.PoolManager(&_Booster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Booster *BoosterCallerSession) PoolManager() (common.Address, error) {
	return _Booster.Contract.PoolManager(&_Booster.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Booster *BoosterCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Booster *BoosterSession) Registry() (common.Address, error) {
	return _Booster.Contract.Registry(&_Booster.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Booster *BoosterCallerSession) Registry() (common.Address, error) {
	return _Booster.Contract.Registry(&_Booster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_Booster *BoosterCaller) RewardArbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "rewardArbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_Booster *BoosterSession) RewardArbitrator() (common.Address, error) {
	return _Booster.Contract.RewardArbitrator(&_Booster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_Booster *BoosterCallerSession) RewardArbitrator() (common.Address, error) {
	return _Booster.Contract.RewardArbitrator(&_Booster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_Booster *BoosterCaller) RewardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "rewardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_Booster *BoosterSession) RewardFactory() (common.Address, error) {
	return _Booster.Contract.RewardFactory(&_Booster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_Booster *BoosterCallerSession) RewardFactory() (common.Address, error) {
	return _Booster.Contract.RewardFactory(&_Booster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Booster *BoosterCaller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "staker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Booster *BoosterSession) Staker() (common.Address, error) {
	return _Booster.Contract.Staker(&_Booster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Booster *BoosterCallerSession) Staker() (common.Address, error) {
	return _Booster.Contract.Staker(&_Booster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_Booster *BoosterCaller) StakerIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "stakerIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_Booster *BoosterSession) StakerIncentive() (*big.Int, error) {
	return _Booster.Contract.StakerIncentive(&_Booster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_Booster *BoosterCallerSession) StakerIncentive() (*big.Int, error) {
	return _Booster.Contract.StakerIncentive(&_Booster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_Booster *BoosterCaller) StakerRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "stakerRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_Booster *BoosterSession) StakerRewards() (common.Address, error) {
	return _Booster.Contract.StakerRewards(&_Booster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_Booster *BoosterCallerSession) StakerRewards() (common.Address, error) {
	return _Booster.Contract.StakerRewards(&_Booster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_Booster *BoosterCaller) StashFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "stashFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_Booster *BoosterSession) StashFactory() (common.Address, error) {
	return _Booster.Contract.StashFactory(&_Booster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_Booster *BoosterCallerSession) StashFactory() (common.Address, error) {
	return _Booster.Contract.StashFactory(&_Booster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Booster *BoosterCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Booster *BoosterSession) TokenFactory() (common.Address, error) {
	return _Booster.Contract.TokenFactory(&_Booster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Booster *BoosterCallerSession) TokenFactory() (common.Address, error) {
	return _Booster.Contract.TokenFactory(&_Booster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Booster *BoosterCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Booster *BoosterSession) Treasury() (common.Address, error) {
	return _Booster.Contract.Treasury(&_Booster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Booster *BoosterCallerSession) Treasury() (common.Address, error) {
	return _Booster.Contract.Treasury(&_Booster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_Booster *BoosterCaller) VoteDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "voteDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_Booster *BoosterSession) VoteDelegate() (common.Address, error) {
	return _Booster.Contract.VoteDelegate(&_Booster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_Booster *BoosterCallerSession) VoteDelegate() (common.Address, error) {
	return _Booster.Contract.VoteDelegate(&_Booster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_Booster *BoosterCaller) VoteOwnership(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "voteOwnership")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_Booster *BoosterSession) VoteOwnership() (common.Address, error) {
	return _Booster.Contract.VoteOwnership(&_Booster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_Booster *BoosterCallerSession) VoteOwnership() (common.Address, error) {
	return _Booster.Contract.VoteOwnership(&_Booster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_Booster *BoosterCaller) VoteParameter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Booster.contract.Call(opts, &out, "voteParameter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_Booster *BoosterSession) VoteParameter() (common.Address, error) {
	return _Booster.Contract.VoteParameter(&_Booster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_Booster *BoosterCallerSession) VoteParameter() (common.Address, error) {
	return _Booster.Contract.VoteParameter(&_Booster.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_Booster *BoosterTransactor) AddPool(opts *bind.TransactOpts, _lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "addPool", _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_Booster *BoosterSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.AddPool(&_Booster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_Booster *BoosterTransactorSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.AddPool(&_Booster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_Booster *BoosterTransactor) ClaimRewards(opts *bind.TransactOpts, _pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "claimRewards", _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_Booster *BoosterSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _Booster.Contract.ClaimRewards(&_Booster.TransactOpts, _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_Booster *BoosterTransactorSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _Booster.Contract.ClaimRewards(&_Booster.TransactOpts, _pid, _gauge)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_Booster *BoosterTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "deposit", _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_Booster *BoosterSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.Contract.Deposit(&_Booster.TransactOpts, _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_Booster *BoosterTransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.Contract.Deposit(&_Booster.TransactOpts, _pid, _amount, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_Booster *BoosterTransactor) DepositAll(opts *bind.TransactOpts, _pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "depositAll", _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_Booster *BoosterSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.Contract.DepositAll(&_Booster.TransactOpts, _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_Booster *BoosterTransactorSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _Booster.Contract.DepositAll(&_Booster.TransactOpts, _pid, _stake)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_Booster *BoosterTransactor) EarmarkFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "earmarkFees")
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_Booster *BoosterSession) EarmarkFees() (*types.Transaction, error) {
	return _Booster.Contract.EarmarkFees(&_Booster.TransactOpts)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0x22230b96.
//
// Solidity: function earmarkFees() returns(bool)
func (_Booster *BoosterTransactorSession) EarmarkFees() (*types.Transaction, error) {
	return _Booster.Contract.EarmarkFees(&_Booster.TransactOpts)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactor) EarmarkRewards(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "earmarkRewards", _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_Booster *BoosterSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.EarmarkRewards(&_Booster.TransactOpts, _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactorSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.EarmarkRewards(&_Booster.TransactOpts, _pid)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_Booster *BoosterTransactor) RewardClaimed(opts *bind.TransactOpts, _pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "rewardClaimed", _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_Booster *BoosterSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.RewardClaimed(&_Booster.TransactOpts, _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_Booster *BoosterTransactorSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.RewardClaimed(&_Booster.TransactOpts, _pid, _address, _amount)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_Booster *BoosterTransactor) SetArbitrator(opts *bind.TransactOpts, _arb common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setArbitrator", _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_Booster *BoosterSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetArbitrator(&_Booster.TransactOpts, _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_Booster *BoosterTransactorSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetArbitrator(&_Booster.TransactOpts, _arb)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_Booster *BoosterTransactor) SetFactories(opts *bind.TransactOpts, _rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setFactories", _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_Booster *BoosterSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetFactories(&_Booster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_Booster *BoosterTransactorSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetFactories(&_Booster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_Booster *BoosterTransactor) SetFeeInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setFeeInfo")
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_Booster *BoosterSession) SetFeeInfo() (*types.Transaction, error) {
	return _Booster.Contract.SetFeeInfo(&_Booster.TransactOpts)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x5a4ae5ca.
//
// Solidity: function setFeeInfo() returns()
func (_Booster *BoosterTransactorSession) SetFeeInfo() (*types.Transaction, error) {
	return _Booster.Contract.SetFeeInfo(&_Booster.TransactOpts)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_Booster *BoosterTransactor) SetFeeManager(opts *bind.TransactOpts, _feeM common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setFeeManager", _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_Booster *BoosterSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetFeeManager(&_Booster.TransactOpts, _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_Booster *BoosterTransactorSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetFeeManager(&_Booster.TransactOpts, _feeM)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_Booster *BoosterTransactor) SetFees(opts *bind.TransactOpts, _lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setFees", _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_Booster *BoosterSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.SetFees(&_Booster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_Booster *BoosterTransactorSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.SetFees(&_Booster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactor) SetGaugeRedirect(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setGaugeRedirect", _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_Booster *BoosterSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.SetGaugeRedirect(&_Booster.TransactOpts, _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactorSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.SetGaugeRedirect(&_Booster.TransactOpts, _pid)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Booster *BoosterTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Booster *BoosterSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetOwner(&_Booster.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Booster *BoosterTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetOwner(&_Booster.TransactOpts, _owner)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_Booster *BoosterTransactor) SetPoolManager(opts *bind.TransactOpts, _poolM common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setPoolManager", _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_Booster *BoosterSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetPoolManager(&_Booster.TransactOpts, _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_Booster *BoosterTransactorSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetPoolManager(&_Booster.TransactOpts, _poolM)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_Booster *BoosterTransactor) SetRewardContracts(opts *bind.TransactOpts, _rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setRewardContracts", _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_Booster *BoosterSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetRewardContracts(&_Booster.TransactOpts, _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_Booster *BoosterTransactorSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetRewardContracts(&_Booster.TransactOpts, _rewards, _stakerRewards)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Booster *BoosterTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Booster *BoosterSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetTreasury(&_Booster.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_Booster *BoosterTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetTreasury(&_Booster.TransactOpts, _treasury)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_Booster *BoosterTransactor) SetVoteDelegate(opts *bind.TransactOpts, _voteDelegate common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "setVoteDelegate", _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_Booster *BoosterSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetVoteDelegate(&_Booster.TransactOpts, _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_Booster *BoosterTransactorSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _Booster.Contract.SetVoteDelegate(&_Booster.TransactOpts, _voteDelegate)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactor) ShutdownPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "shutdownPool", _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_Booster *BoosterSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.ShutdownPool(&_Booster.TransactOpts, _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactorSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.ShutdownPool(&_Booster.TransactOpts, _pid)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_Booster *BoosterTransactor) ShutdownSystem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "shutdownSystem")
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_Booster *BoosterSession) ShutdownSystem() (*types.Transaction, error) {
	return _Booster.Contract.ShutdownSystem(&_Booster.TransactOpts)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_Booster *BoosterTransactorSession) ShutdownSystem() (*types.Transaction, error) {
	return _Booster.Contract.ShutdownSystem(&_Booster.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_Booster *BoosterTransactor) Vote(opts *bind.TransactOpts, _voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "vote", _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_Booster *BoosterSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _Booster.Contract.Vote(&_Booster.TransactOpts, _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_Booster *BoosterTransactorSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _Booster.Contract.Vote(&_Booster.TransactOpts, _voteId, _votingAddress, _support)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_Booster *BoosterTransactor) VoteGaugeWeight(opts *bind.TransactOpts, _gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "voteGaugeWeight", _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_Booster *BoosterSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _Booster.Contract.VoteGaugeWeight(&_Booster.TransactOpts, _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_Booster *BoosterTransactorSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _Booster.Contract.VoteGaugeWeight(&_Booster.TransactOpts, _gauge, _weight)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_Booster *BoosterTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_Booster *BoosterSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.Withdraw(&_Booster.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_Booster *BoosterTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.Withdraw(&_Booster.TransactOpts, _pid, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactor) WithdrawAll(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "withdrawAll", _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_Booster *BoosterSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.WithdrawAll(&_Booster.TransactOpts, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_Booster *BoosterTransactorSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _Booster.Contract.WithdrawAll(&_Booster.TransactOpts, _pid)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_Booster *BoosterTransactor) WithdrawTo(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Booster.contract.Transact(opts, "withdrawTo", _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_Booster *BoosterSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Booster.Contract.WithdrawTo(&_Booster.TransactOpts, _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_Booster *BoosterTransactorSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Booster.Contract.WithdrawTo(&_Booster.TransactOpts, _pid, _amount, _to)
}

// BoosterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Booster contract.
type BoosterDepositedIterator struct {
	Event *BoosterDeposited // Event containing the contract specifics and raw log

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
func (it *BoosterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoosterDeposited)
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
		it.Event = new(BoosterDeposited)
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
func (it *BoosterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoosterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoosterDeposited represents a Deposited event raised by the Booster contract.
type BoosterDeposited struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*BoosterDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Booster.contract.FilterLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &BoosterDepositedIterator{contract: _Booster.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *BoosterDeposited, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Booster.contract.WatchLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoosterDeposited)
				if err := _Booster.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) ParseDeposited(log types.Log) (*BoosterDeposited, error) {
	event := new(BoosterDeposited)
	if err := _Booster.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoosterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Booster contract.
type BoosterWithdrawnIterator struct {
	Event *BoosterWithdrawn // Event containing the contract specifics and raw log

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
func (it *BoosterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoosterWithdrawn)
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
		it.Event = new(BoosterWithdrawn)
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
func (it *BoosterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoosterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoosterWithdrawn represents a Withdrawn event raised by the Booster contract.
type BoosterWithdrawn struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*BoosterWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Booster.contract.FilterLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &BoosterWithdrawnIterator{contract: _Booster.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BoosterWithdrawn, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Booster.contract.WatchLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoosterWithdrawn)
				if err := _Booster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_Booster *BoosterFilterer) ParseWithdrawn(log types.Log) (*BoosterWithdrawn, error) {
	event := new(BoosterWithdrawn)
	if err := _Booster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICrvDepositMetaData contains all meta data concerning the ICrvDeposit contract.
var ICrvDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9a408321": "deposit(uint256,bool)",
		"50940618": "lockIncentive()",
	},
}

// ICrvDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use ICrvDepositMetaData.ABI instead.
var ICrvDepositABI = ICrvDepositMetaData.ABI

// Deprecated: Use ICrvDepositMetaData.Sigs instead.
// ICrvDepositFuncSigs maps the 4-byte function signature to its string representation.
var ICrvDepositFuncSigs = ICrvDepositMetaData.Sigs

// ICrvDeposit is an auto generated Go binding around an Ethereum contract.
type ICrvDeposit struct {
	ICrvDepositCaller     // Read-only binding to the contract
	ICrvDepositTransactor // Write-only binding to the contract
	ICrvDepositFilterer   // Log filterer for contract events
}

// ICrvDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrvDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrvDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrvDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrvDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrvDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrvDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrvDepositSession struct {
	Contract     *ICrvDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICrvDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrvDepositCallerSession struct {
	Contract *ICrvDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICrvDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrvDepositTransactorSession struct {
	Contract     *ICrvDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICrvDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrvDepositRaw struct {
	Contract *ICrvDeposit // Generic contract binding to access the raw methods on
}

// ICrvDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrvDepositCallerRaw struct {
	Contract *ICrvDepositCaller // Generic read-only contract binding to access the raw methods on
}

// ICrvDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrvDepositTransactorRaw struct {
	Contract *ICrvDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrvDeposit creates a new instance of ICrvDeposit, bound to a specific deployed contract.
func NewICrvDeposit(address common.Address, backend bind.ContractBackend) (*ICrvDeposit, error) {
	contract, err := bindICrvDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrvDeposit{ICrvDepositCaller: ICrvDepositCaller{contract: contract}, ICrvDepositTransactor: ICrvDepositTransactor{contract: contract}, ICrvDepositFilterer: ICrvDepositFilterer{contract: contract}}, nil
}

// NewICrvDepositCaller creates a new read-only instance of ICrvDeposit, bound to a specific deployed contract.
func NewICrvDepositCaller(address common.Address, caller bind.ContractCaller) (*ICrvDepositCaller, error) {
	contract, err := bindICrvDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrvDepositCaller{contract: contract}, nil
}

// NewICrvDepositTransactor creates a new write-only instance of ICrvDeposit, bound to a specific deployed contract.
func NewICrvDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrvDepositTransactor, error) {
	contract, err := bindICrvDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrvDepositTransactor{contract: contract}, nil
}

// NewICrvDepositFilterer creates a new log filterer instance of ICrvDeposit, bound to a specific deployed contract.
func NewICrvDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrvDepositFilterer, error) {
	contract, err := bindICrvDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrvDepositFilterer{contract: contract}, nil
}

// bindICrvDeposit binds a generic wrapper to an already deployed contract.
func bindICrvDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICrvDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrvDeposit *ICrvDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrvDeposit.Contract.ICrvDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrvDeposit *ICrvDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.ICrvDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrvDeposit *ICrvDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.ICrvDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrvDeposit *ICrvDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrvDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrvDeposit *ICrvDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrvDeposit *ICrvDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.contract.Transact(opts, method, params...)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_ICrvDeposit *ICrvDepositCaller) LockIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICrvDeposit.contract.Call(opts, &out, "lockIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_ICrvDeposit *ICrvDepositSession) LockIncentive() (*big.Int, error) {
	return _ICrvDeposit.Contract.LockIncentive(&_ICrvDeposit.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_ICrvDeposit *ICrvDepositCallerSession) LockIncentive() (*big.Int, error) {
	return _ICrvDeposit.Contract.LockIncentive(&_ICrvDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 , bool ) returns()
func (_ICrvDeposit *ICrvDepositTransactor) Deposit(opts *bind.TransactOpts, arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _ICrvDeposit.contract.Transact(opts, "deposit", arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 , bool ) returns()
func (_ICrvDeposit *ICrvDepositSession) Deposit(arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.Deposit(&_ICrvDeposit.TransactOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 , bool ) returns()
func (_ICrvDeposit *ICrvDepositTransactorSession) Deposit(arg0 *big.Int, arg1 bool) (*types.Transaction, error) {
	return _ICrvDeposit.Contract.Deposit(&_ICrvDeposit.TransactOpts, arg0, arg1)
}

// ICurveGaugeMetaData contains all meta data concerning the ICurveGauge contract.
var ICurveGaugeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim_rewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reward_tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewarded_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"e6f1daf2": "claim_rewards()",
		"b6b55f25": "deposit(uint256)",
		"54c49fe9": "reward_tokens(uint256)",
		"16fa50b1": "rewarded_token()",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// ICurveGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurveGaugeMetaData.ABI instead.
var ICurveGaugeABI = ICurveGaugeMetaData.ABI

// Deprecated: Use ICurveGaugeMetaData.Sigs instead.
// ICurveGaugeFuncSigs maps the 4-byte function signature to its string representation.
var ICurveGaugeFuncSigs = ICurveGaugeMetaData.Sigs

// ICurveGauge is an auto generated Go binding around an Ethereum contract.
type ICurveGauge struct {
	ICurveGaugeCaller     // Read-only binding to the contract
	ICurveGaugeTransactor // Write-only binding to the contract
	ICurveGaugeFilterer   // Log filterer for contract events
}

// ICurveGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurveGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurveGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurveGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurveGaugeSession struct {
	Contract     *ICurveGauge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurveGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurveGaugeCallerSession struct {
	Contract *ICurveGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICurveGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurveGaugeTransactorSession struct {
	Contract     *ICurveGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICurveGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurveGaugeRaw struct {
	Contract *ICurveGauge // Generic contract binding to access the raw methods on
}

// ICurveGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurveGaugeCallerRaw struct {
	Contract *ICurveGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// ICurveGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurveGaugeTransactorRaw struct {
	Contract *ICurveGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurveGauge creates a new instance of ICurveGauge, bound to a specific deployed contract.
func NewICurveGauge(address common.Address, backend bind.ContractBackend) (*ICurveGauge, error) {
	contract, err := bindICurveGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurveGauge{ICurveGaugeCaller: ICurveGaugeCaller{contract: contract}, ICurveGaugeTransactor: ICurveGaugeTransactor{contract: contract}, ICurveGaugeFilterer: ICurveGaugeFilterer{contract: contract}}, nil
}

// NewICurveGaugeCaller creates a new read-only instance of ICurveGauge, bound to a specific deployed contract.
func NewICurveGaugeCaller(address common.Address, caller bind.ContractCaller) (*ICurveGaugeCaller, error) {
	contract, err := bindICurveGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveGaugeCaller{contract: contract}, nil
}

// NewICurveGaugeTransactor creates a new write-only instance of ICurveGauge, bound to a specific deployed contract.
func NewICurveGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurveGaugeTransactor, error) {
	contract, err := bindICurveGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveGaugeTransactor{contract: contract}, nil
}

// NewICurveGaugeFilterer creates a new log filterer instance of ICurveGauge, bound to a specific deployed contract.
func NewICurveGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurveGaugeFilterer, error) {
	contract, err := bindICurveGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurveGaugeFilterer{contract: contract}, nil
}

// bindICurveGauge binds a generic wrapper to an already deployed contract.
func bindICurveGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICurveGaugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveGauge *ICurveGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveGauge.Contract.ICurveGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveGauge *ICurveGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveGauge.Contract.ICurveGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveGauge *ICurveGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveGauge.Contract.ICurveGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveGauge *ICurveGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveGauge *ICurveGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveGauge *ICurveGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveGauge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ICurveGauge *ICurveGaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICurveGauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ICurveGauge *ICurveGaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ICurveGauge.Contract.BalanceOf(&_ICurveGauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ICurveGauge *ICurveGaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ICurveGauge.Contract.BalanceOf(&_ICurveGauge.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 ) view returns(address)
func (_ICurveGauge *ICurveGaugeCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICurveGauge.contract.Call(opts, &out, "reward_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 ) view returns(address)
func (_ICurveGauge *ICurveGaugeSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _ICurveGauge.Contract.RewardTokens(&_ICurveGauge.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 ) view returns(address)
func (_ICurveGauge *ICurveGaugeCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _ICurveGauge.Contract.RewardTokens(&_ICurveGauge.CallOpts, arg0)
}

// RewardedToken is a free data retrieval call binding the contract method 0x16fa50b1.
//
// Solidity: function rewarded_token() view returns(address)
func (_ICurveGauge *ICurveGaugeCaller) RewardedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICurveGauge.contract.Call(opts, &out, "rewarded_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardedToken is a free data retrieval call binding the contract method 0x16fa50b1.
//
// Solidity: function rewarded_token() view returns(address)
func (_ICurveGauge *ICurveGaugeSession) RewardedToken() (common.Address, error) {
	return _ICurveGauge.Contract.RewardedToken(&_ICurveGauge.CallOpts)
}

// RewardedToken is a free data retrieval call binding the contract method 0x16fa50b1.
//
// Solidity: function rewarded_token() view returns(address)
func (_ICurveGauge *ICurveGaugeCallerSession) RewardedToken() (common.Address, error) {
	return _ICurveGauge.Contract.RewardedToken(&_ICurveGauge.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_ICurveGauge *ICurveGaugeTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveGauge.contract.Transact(opts, "claim_rewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_ICurveGauge *ICurveGaugeSession) ClaimRewards() (*types.Transaction, error) {
	return _ICurveGauge.Contract.ClaimRewards(&_ICurveGauge.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_ICurveGauge *ICurveGaugeTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _ICurveGauge.Contract.ClaimRewards(&_ICurveGauge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeTransactor) Deposit(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.contract.Transact(opts, "deposit", arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeSession) Deposit(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.Contract.Deposit(&_ICurveGauge.TransactOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeTransactorSession) Deposit(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.Contract.Deposit(&_ICurveGauge.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.Contract.Withdraw(&_ICurveGauge.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_ICurveGauge *ICurveGaugeTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveGauge.Contract.Withdraw(&_ICurveGauge.TransactOpts, arg0)
}

// ICurveVoteEscrowMetaData contains all meta data concerning the ICurveVoteEscrow contract.
var ICurveVoteEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"create_lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increase_amount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increase_unlock_time\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smart_wallet_checker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"65fc3873": "create_lock(uint256,uint256)",
		"4957677c": "increase_amount(uint256)",
		"eff7a612": "increase_unlock_time(uint256)",
		"7175d4f7": "smart_wallet_checker()",
		"3ccfd60b": "withdraw()",
	},
}

// ICurveVoteEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use ICurveVoteEscrowMetaData.ABI instead.
var ICurveVoteEscrowABI = ICurveVoteEscrowMetaData.ABI

// Deprecated: Use ICurveVoteEscrowMetaData.Sigs instead.
// ICurveVoteEscrowFuncSigs maps the 4-byte function signature to its string representation.
var ICurveVoteEscrowFuncSigs = ICurveVoteEscrowMetaData.Sigs

// ICurveVoteEscrow is an auto generated Go binding around an Ethereum contract.
type ICurveVoteEscrow struct {
	ICurveVoteEscrowCaller     // Read-only binding to the contract
	ICurveVoteEscrowTransactor // Write-only binding to the contract
	ICurveVoteEscrowFilterer   // Log filterer for contract events
}

// ICurveVoteEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICurveVoteEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveVoteEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICurveVoteEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveVoteEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICurveVoteEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICurveVoteEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICurveVoteEscrowSession struct {
	Contract     *ICurveVoteEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICurveVoteEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICurveVoteEscrowCallerSession struct {
	Contract *ICurveVoteEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ICurveVoteEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICurveVoteEscrowTransactorSession struct {
	Contract     *ICurveVoteEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ICurveVoteEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICurveVoteEscrowRaw struct {
	Contract *ICurveVoteEscrow // Generic contract binding to access the raw methods on
}

// ICurveVoteEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICurveVoteEscrowCallerRaw struct {
	Contract *ICurveVoteEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// ICurveVoteEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICurveVoteEscrowTransactorRaw struct {
	Contract *ICurveVoteEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICurveVoteEscrow creates a new instance of ICurveVoteEscrow, bound to a specific deployed contract.
func NewICurveVoteEscrow(address common.Address, backend bind.ContractBackend) (*ICurveVoteEscrow, error) {
	contract, err := bindICurveVoteEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICurveVoteEscrow{ICurveVoteEscrowCaller: ICurveVoteEscrowCaller{contract: contract}, ICurveVoteEscrowTransactor: ICurveVoteEscrowTransactor{contract: contract}, ICurveVoteEscrowFilterer: ICurveVoteEscrowFilterer{contract: contract}}, nil
}

// NewICurveVoteEscrowCaller creates a new read-only instance of ICurveVoteEscrow, bound to a specific deployed contract.
func NewICurveVoteEscrowCaller(address common.Address, caller bind.ContractCaller) (*ICurveVoteEscrowCaller, error) {
	contract, err := bindICurveVoteEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveVoteEscrowCaller{contract: contract}, nil
}

// NewICurveVoteEscrowTransactor creates a new write-only instance of ICurveVoteEscrow, bound to a specific deployed contract.
func NewICurveVoteEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ICurveVoteEscrowTransactor, error) {
	contract, err := bindICurveVoteEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICurveVoteEscrowTransactor{contract: contract}, nil
}

// NewICurveVoteEscrowFilterer creates a new log filterer instance of ICurveVoteEscrow, bound to a specific deployed contract.
func NewICurveVoteEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ICurveVoteEscrowFilterer, error) {
	contract, err := bindICurveVoteEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICurveVoteEscrowFilterer{contract: contract}, nil
}

// bindICurveVoteEscrow binds a generic wrapper to an already deployed contract.
func bindICurveVoteEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICurveVoteEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveVoteEscrow *ICurveVoteEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveVoteEscrow.Contract.ICurveVoteEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveVoteEscrow *ICurveVoteEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.ICurveVoteEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveVoteEscrow *ICurveVoteEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.ICurveVoteEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICurveVoteEscrow *ICurveVoteEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICurveVoteEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.contract.Transact(opts, method, params...)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_ICurveVoteEscrow *ICurveVoteEscrowCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICurveVoteEscrow.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_ICurveVoteEscrow *ICurveVoteEscrowSession) SmartWalletChecker() (common.Address, error) {
	return _ICurveVoteEscrow.Contract.SmartWalletChecker(&_ICurveVoteEscrow.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_ICurveVoteEscrow *ICurveVoteEscrowCallerSession) SmartWalletChecker() (common.Address, error) {
	return _ICurveVoteEscrow.Contract.SmartWalletChecker(&_ICurveVoteEscrow.CallOpts)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 , uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactor) CreateLock(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.contract.Transact(opts, "create_lock", arg0, arg1)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 , uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowSession) CreateLock(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.CreateLock(&_ICurveVoteEscrow.TransactOpts, arg0, arg1)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 , uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorSession) CreateLock(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.CreateLock(&_ICurveVoteEscrow.TransactOpts, arg0, arg1)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactor) IncreaseAmount(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.contract.Transact(opts, "increase_amount", arg0)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowSession) IncreaseAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.IncreaseAmount(&_ICurveVoteEscrow.TransactOpts, arg0)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorSession) IncreaseAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.IncreaseAmount(&_ICurveVoteEscrow.TransactOpts, arg0)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.contract.Transact(opts, "increase_unlock_time", arg0)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowSession) IncreaseUnlockTime(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.IncreaseUnlockTime(&_ICurveVoteEscrow.TransactOpts, arg0)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 ) returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorSession) IncreaseUnlockTime(arg0 *big.Int) (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.IncreaseUnlockTime(&_ICurveVoteEscrow.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICurveVoteEscrow.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowSession) Withdraw() (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.Withdraw(&_ICurveVoteEscrow.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ICurveVoteEscrow *ICurveVoteEscrowTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ICurveVoteEscrow.Contract.Withdraw(&_ICurveVoteEscrow.TransactOpts)
}

// IDepositMetaData contains all meta data concerning the IDeposit contract.
var IDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardArbitrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardClaimed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"6c7b69cb": "claimRewards(uint256,address)",
		"bf86d690": "isShutdown()",
		"1526fe27": "poolInfo(uint256)",
		"043b684a": "rewardArbitrator()",
		"71192b17": "rewardClaimed(uint256,address,uint256)",
		"18160ddd": "totalSupply()",
		"14cd70e4": "withdrawTo(uint256,uint256,address)",
	},
}

// IDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use IDepositMetaData.ABI instead.
var IDepositABI = IDepositMetaData.ABI

// Deprecated: Use IDepositMetaData.Sigs instead.
// IDepositFuncSigs maps the 4-byte function signature to its string representation.
var IDepositFuncSigs = IDepositMetaData.Sigs

// IDeposit is an auto generated Go binding around an Ethereum contract.
type IDeposit struct {
	IDepositCaller     // Read-only binding to the contract
	IDepositTransactor // Write-only binding to the contract
	IDepositFilterer   // Log filterer for contract events
}

// IDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositSession struct {
	Contract     *IDeposit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositCallerSession struct {
	Contract *IDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositTransactorSession struct {
	Contract     *IDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositRaw struct {
	Contract *IDeposit // Generic contract binding to access the raw methods on
}

// IDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositCallerRaw struct {
	Contract *IDepositCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositTransactorRaw struct {
	Contract *IDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDeposit creates a new instance of IDeposit, bound to a specific deployed contract.
func NewIDeposit(address common.Address, backend bind.ContractBackend) (*IDeposit, error) {
	contract, err := bindIDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDeposit{IDepositCaller: IDepositCaller{contract: contract}, IDepositTransactor: IDepositTransactor{contract: contract}, IDepositFilterer: IDepositFilterer{contract: contract}}, nil
}

// NewIDepositCaller creates a new read-only instance of IDeposit, bound to a specific deployed contract.
func NewIDepositCaller(address common.Address, caller bind.ContractCaller) (*IDepositCaller, error) {
	contract, err := bindIDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositCaller{contract: contract}, nil
}

// NewIDepositTransactor creates a new write-only instance of IDeposit, bound to a specific deployed contract.
func NewIDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositTransactor, error) {
	contract, err := bindIDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositTransactor{contract: contract}, nil
}

// NewIDepositFilterer creates a new log filterer instance of IDeposit, bound to a specific deployed contract.
func NewIDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositFilterer, error) {
	contract, err := bindIDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositFilterer{contract: contract}, nil
}

// bindIDeposit binds a generic wrapper to an already deployed contract.
func bindIDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDeposit *IDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDeposit.Contract.IDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDeposit *IDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDeposit.Contract.IDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDeposit *IDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDeposit.Contract.IDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDeposit *IDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDeposit *IDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDeposit *IDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDeposit.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_IDeposit *IDepositCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDeposit.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_IDeposit *IDepositSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _IDeposit.Contract.BalanceOf(&_IDeposit.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_IDeposit *IDepositCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _IDeposit.Contract.BalanceOf(&_IDeposit.CallOpts, _account)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_IDeposit *IDepositCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IDeposit.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_IDeposit *IDepositSession) IsShutdown() (bool, error) {
	return _IDeposit.Contract.IsShutdown(&_IDeposit.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_IDeposit *IDepositCallerSession) IsShutdown() (bool, error) {
	return _IDeposit.Contract.IsShutdown(&_IDeposit.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IDeposit *IDepositCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _IDeposit.contract.Call(opts, &out, "poolInfo", arg0)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IDeposit *IDepositSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IDeposit.Contract.PoolInfo(&_IDeposit.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IDeposit *IDepositCallerSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IDeposit.Contract.PoolInfo(&_IDeposit.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDeposit *IDepositCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDeposit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDeposit *IDepositSession) TotalSupply() (*big.Int, error) {
	return _IDeposit.Contract.TotalSupply(&_IDeposit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IDeposit *IDepositCallerSession) TotalSupply() (*big.Int, error) {
	return _IDeposit.Contract.TotalSupply(&_IDeposit.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 , address ) returns(bool)
func (_IDeposit *IDepositTransactor) ClaimRewards(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IDeposit.contract.Transact(opts, "claimRewards", arg0, arg1)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 , address ) returns(bool)
func (_IDeposit *IDepositSession) ClaimRewards(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IDeposit.Contract.ClaimRewards(&_IDeposit.TransactOpts, arg0, arg1)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 , address ) returns(bool)
func (_IDeposit *IDepositTransactorSession) ClaimRewards(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IDeposit.Contract.ClaimRewards(&_IDeposit.TransactOpts, arg0, arg1)
}

// RewardArbitrator is a paid mutator transaction binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() returns(address)
func (_IDeposit *IDepositTransactor) RewardArbitrator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDeposit.contract.Transact(opts, "rewardArbitrator")
}

// RewardArbitrator is a paid mutator transaction binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() returns(address)
func (_IDeposit *IDepositSession) RewardArbitrator() (*types.Transaction, error) {
	return _IDeposit.Contract.RewardArbitrator(&_IDeposit.TransactOpts)
}

// RewardArbitrator is a paid mutator transaction binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() returns(address)
func (_IDeposit *IDepositTransactorSession) RewardArbitrator() (*types.Transaction, error) {
	return _IDeposit.Contract.RewardArbitrator(&_IDeposit.TransactOpts)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 , address , uint256 ) returns()
func (_IDeposit *IDepositTransactor) RewardClaimed(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IDeposit.contract.Transact(opts, "rewardClaimed", arg0, arg1, arg2)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 , address , uint256 ) returns()
func (_IDeposit *IDepositSession) RewardClaimed(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IDeposit.Contract.RewardClaimed(&_IDeposit.TransactOpts, arg0, arg1, arg2)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 , address , uint256 ) returns()
func (_IDeposit *IDepositTransactorSession) RewardClaimed(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IDeposit.Contract.RewardClaimed(&_IDeposit.TransactOpts, arg0, arg1, arg2)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_IDeposit *IDepositTransactor) WithdrawTo(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IDeposit.contract.Transact(opts, "withdrawTo", arg0, arg1, arg2)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_IDeposit *IDepositSession) WithdrawTo(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IDeposit.Contract.WithdrawTo(&_IDeposit.TransactOpts, arg0, arg1, arg2)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 , uint256 , address ) returns()
func (_IDeposit *IDepositTransactorSession) WithdrawTo(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IDeposit.Contract.WithdrawTo(&_IDeposit.TransactOpts, arg0, arg1, arg2)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeDistroMetaData contains all meta data concerning the IFeeDistro contract.
var IFeeDistroMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4e71d92d": "claim()",
		"fc0c546a": "token()",
	},
}

// IFeeDistroABI is the input ABI used to generate the binding from.
// Deprecated: Use IFeeDistroMetaData.ABI instead.
var IFeeDistroABI = IFeeDistroMetaData.ABI

// Deprecated: Use IFeeDistroMetaData.Sigs instead.
// IFeeDistroFuncSigs maps the 4-byte function signature to its string representation.
var IFeeDistroFuncSigs = IFeeDistroMetaData.Sigs

// IFeeDistro is an auto generated Go binding around an Ethereum contract.
type IFeeDistro struct {
	IFeeDistroCaller     // Read-only binding to the contract
	IFeeDistroTransactor // Write-only binding to the contract
	IFeeDistroFilterer   // Log filterer for contract events
}

// IFeeDistroCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeDistroCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeDistroTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeDistroTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeDistroFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeDistroFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeDistroSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeDistroSession struct {
	Contract     *IFeeDistro       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeDistroCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeDistroCallerSession struct {
	Contract *IFeeDistroCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IFeeDistroTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeDistroTransactorSession struct {
	Contract     *IFeeDistroTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IFeeDistroRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeDistroRaw struct {
	Contract *IFeeDistro // Generic contract binding to access the raw methods on
}

// IFeeDistroCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeDistroCallerRaw struct {
	Contract *IFeeDistroCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeDistroTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeDistroTransactorRaw struct {
	Contract *IFeeDistroTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeDistro creates a new instance of IFeeDistro, bound to a specific deployed contract.
func NewIFeeDistro(address common.Address, backend bind.ContractBackend) (*IFeeDistro, error) {
	contract, err := bindIFeeDistro(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeDistro{IFeeDistroCaller: IFeeDistroCaller{contract: contract}, IFeeDistroTransactor: IFeeDistroTransactor{contract: contract}, IFeeDistroFilterer: IFeeDistroFilterer{contract: contract}}, nil
}

// NewIFeeDistroCaller creates a new read-only instance of IFeeDistro, bound to a specific deployed contract.
func NewIFeeDistroCaller(address common.Address, caller bind.ContractCaller) (*IFeeDistroCaller, error) {
	contract, err := bindIFeeDistro(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeDistroCaller{contract: contract}, nil
}

// NewIFeeDistroTransactor creates a new write-only instance of IFeeDistro, bound to a specific deployed contract.
func NewIFeeDistroTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeDistroTransactor, error) {
	contract, err := bindIFeeDistro(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeDistroTransactor{contract: contract}, nil
}

// NewIFeeDistroFilterer creates a new log filterer instance of IFeeDistro, bound to a specific deployed contract.
func NewIFeeDistroFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeDistroFilterer, error) {
	contract, err := bindIFeeDistro(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeDistroFilterer{contract: contract}, nil
}

// bindIFeeDistro binds a generic wrapper to an already deployed contract.
func bindIFeeDistro(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeDistroABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeDistro *IFeeDistroRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeDistro.Contract.IFeeDistroCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeDistro *IFeeDistroRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeDistro.Contract.IFeeDistroTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeDistro *IFeeDistroRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeDistro.Contract.IFeeDistroTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeDistro *IFeeDistroCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeDistro.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeDistro *IFeeDistroTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeDistro.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeDistro *IFeeDistroTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeDistro.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IFeeDistro *IFeeDistroCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFeeDistro.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IFeeDistro *IFeeDistroSession) Token() (common.Address, error) {
	return _IFeeDistro.Contract.Token(&_IFeeDistro.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IFeeDistro *IFeeDistroCallerSession) Token() (common.Address, error) {
	return _IFeeDistro.Contract.Token(&_IFeeDistro.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IFeeDistro *IFeeDistroTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeDistro.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IFeeDistro *IFeeDistroSession) Claim() (*types.Transaction, error) {
	return _IFeeDistro.Contract.Claim(&_IFeeDistro.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IFeeDistro *IFeeDistroTransactorSession) Claim() (*types.Transaction, error) {
	return _IFeeDistro.Contract.Claim(&_IFeeDistro.TransactOpts)
}

// IMinterMetaData contains all meta data concerning the IMinter contract.
var IMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6a627842": "mint(address)",
	},
}

// IMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMinterMetaData.ABI instead.
var IMinterABI = IMinterMetaData.ABI

// Deprecated: Use IMinterMetaData.Sigs instead.
// IMinterFuncSigs maps the 4-byte function signature to its string representation.
var IMinterFuncSigs = IMinterMetaData.Sigs

// IMinter is an auto generated Go binding around an Ethereum contract.
type IMinter struct {
	IMinterCaller     // Read-only binding to the contract
	IMinterTransactor // Write-only binding to the contract
	IMinterFilterer   // Log filterer for contract events
}

// IMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMinterSession struct {
	Contract     *IMinter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMinterCallerSession struct {
	Contract *IMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMinterTransactorSession struct {
	Contract     *IMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMinterRaw struct {
	Contract *IMinter // Generic contract binding to access the raw methods on
}

// IMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMinterCallerRaw struct {
	Contract *IMinterCaller // Generic read-only contract binding to access the raw methods on
}

// IMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMinterTransactorRaw struct {
	Contract *IMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMinter creates a new instance of IMinter, bound to a specific deployed contract.
func NewIMinter(address common.Address, backend bind.ContractBackend) (*IMinter, error) {
	contract, err := bindIMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMinter{IMinterCaller: IMinterCaller{contract: contract}, IMinterTransactor: IMinterTransactor{contract: contract}, IMinterFilterer: IMinterFilterer{contract: contract}}, nil
}

// NewIMinterCaller creates a new read-only instance of IMinter, bound to a specific deployed contract.
func NewIMinterCaller(address common.Address, caller bind.ContractCaller) (*IMinterCaller, error) {
	contract, err := bindIMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMinterCaller{contract: contract}, nil
}

// NewIMinterTransactor creates a new write-only instance of IMinter, bound to a specific deployed contract.
func NewIMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMinterTransactor, error) {
	contract, err := bindIMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMinterTransactor{contract: contract}, nil
}

// NewIMinterFilterer creates a new log filterer instance of IMinter, bound to a specific deployed contract.
func NewIMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMinterFilterer, error) {
	contract, err := bindIMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMinterFilterer{contract: contract}, nil
}

// bindIMinter binds a generic wrapper to an already deployed contract.
func bindIMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinter *IMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMinter.Contract.IMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinter *IMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinter.Contract.IMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinter *IMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinter.Contract.IMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMinter *IMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMinter *IMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMinter *IMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMinter.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address ) returns()
func (_IMinter *IMinterTransactor) Mint(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IMinter.contract.Transact(opts, "mint", arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address ) returns()
func (_IMinter *IMinterSession) Mint(arg0 common.Address) (*types.Transaction, error) {
	return _IMinter.Contract.Mint(&_IMinter.TransactOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address ) returns()
func (_IMinter *IMinterTransactorSession) Mint(arg0 common.Address) (*types.Transaction, error) {
	return _IMinter.Contract.Mint(&_IMinter.TransactOpts, arg0)
}

// IPoolsMetaData contains all meta data concerning the IPools contract.
var IPoolsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stashVersion\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolM\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"shutdownPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7e29d6c2": "addPool(address,address,uint256)",
		"cb0d5b52": "gaugeMap(address)",
		"1526fe27": "poolInfo(uint256)",
		"081e3eda": "poolLength()",
		"7aef6715": "setPoolManager(address)",
		"60cafe84": "shutdownPool(uint256)",
	},
}

// IPoolsABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolsMetaData.ABI instead.
var IPoolsABI = IPoolsMetaData.ABI

// Deprecated: Use IPoolsMetaData.Sigs instead.
// IPoolsFuncSigs maps the 4-byte function signature to its string representation.
var IPoolsFuncSigs = IPoolsMetaData.Sigs

// IPools is an auto generated Go binding around an Ethereum contract.
type IPools struct {
	IPoolsCaller     // Read-only binding to the contract
	IPoolsTransactor // Write-only binding to the contract
	IPoolsFilterer   // Log filterer for contract events
}

// IPoolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolsSession struct {
	Contract     *IPools           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolsCallerSession struct {
	Contract *IPoolsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolsTransactorSession struct {
	Contract     *IPoolsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolsRaw struct {
	Contract *IPools // Generic contract binding to access the raw methods on
}

// IPoolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolsCallerRaw struct {
	Contract *IPoolsCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolsTransactorRaw struct {
	Contract *IPoolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPools creates a new instance of IPools, bound to a specific deployed contract.
func NewIPools(address common.Address, backend bind.ContractBackend) (*IPools, error) {
	contract, err := bindIPools(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPools{IPoolsCaller: IPoolsCaller{contract: contract}, IPoolsTransactor: IPoolsTransactor{contract: contract}, IPoolsFilterer: IPoolsFilterer{contract: contract}}, nil
}

// NewIPoolsCaller creates a new read-only instance of IPools, bound to a specific deployed contract.
func NewIPoolsCaller(address common.Address, caller bind.ContractCaller) (*IPoolsCaller, error) {
	contract, err := bindIPools(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolsCaller{contract: contract}, nil
}

// NewIPoolsTransactor creates a new write-only instance of IPools, bound to a specific deployed contract.
func NewIPoolsTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolsTransactor, error) {
	contract, err := bindIPools(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolsTransactor{contract: contract}, nil
}

// NewIPoolsFilterer creates a new log filterer instance of IPools, bound to a specific deployed contract.
func NewIPoolsFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolsFilterer, error) {
	contract, err := bindIPools(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolsFilterer{contract: contract}, nil
}

// bindIPools binds a generic wrapper to an already deployed contract.
func bindIPools(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPools *IPoolsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPools.Contract.IPoolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPools *IPoolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPools.Contract.IPoolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPools *IPoolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPools.Contract.IPoolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPools *IPoolsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPools.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPools *IPoolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPools.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPools *IPoolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPools.Contract.contract.Transact(opts, method, params...)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_IPools *IPoolsCaller) GaugeMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IPools.contract.Call(opts, &out, "gaugeMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_IPools *IPoolsSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _IPools.Contract.GaugeMap(&_IPools.CallOpts, arg0)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_IPools *IPoolsCallerSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _IPools.Contract.GaugeMap(&_IPools.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IPools *IPoolsCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	var out []interface{}
	err := _IPools.contract.Call(opts, &out, "poolInfo", arg0)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IPools *IPoolsSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IPools.Contract.PoolInfo(&_IPools.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address, address, address, address, address, bool)
func (_IPools *IPoolsCallerSession) PoolInfo(arg0 *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, bool, error) {
	return _IPools.Contract.PoolInfo(&_IPools.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_IPools *IPoolsCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPools.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_IPools *IPoolsSession) PoolLength() (*big.Int, error) {
	return _IPools.Contract.PoolLength(&_IPools.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_IPools *IPoolsCallerSession) PoolLength() (*big.Int, error) {
	return _IPools.Contract.PoolLength(&_IPools.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_IPools *IPoolsTransactor) AddPool(opts *bind.TransactOpts, _lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _IPools.contract.Transact(opts, "addPool", _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_IPools *IPoolsSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _IPools.Contract.AddPool(&_IPools.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_IPools *IPoolsTransactorSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _IPools.Contract.AddPool(&_IPools.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_IPools *IPoolsTransactor) SetPoolManager(opts *bind.TransactOpts, _poolM common.Address) (*types.Transaction, error) {
	return _IPools.contract.Transact(opts, "setPoolManager", _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_IPools *IPoolsSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _IPools.Contract.SetPoolManager(&_IPools.TransactOpts, _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_IPools *IPoolsTransactorSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _IPools.Contract.SetPoolManager(&_IPools.TransactOpts, _poolM)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_IPools *IPoolsTransactor) ShutdownPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _IPools.contract.Transact(opts, "shutdownPool", _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_IPools *IPoolsSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _IPools.Contract.ShutdownPool(&_IPools.TransactOpts, _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_IPools *IPoolsTransactorSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _IPools.Contract.ShutdownPool(&_IPools.TransactOpts, _pid)
}

// IRegistryMetaData contains all meta data concerning the IRegistry contract.
var IRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"gauge_controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"get_gauges\",\"outputs\":[{\"internalType\":\"address[10]\",\"name\":\"\",\"type\":\"address[10]\"},{\"internalType\":\"uint128[10]\",\"name\":\"\",\"type\":\"uint128[10]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"get_lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d8b9a018": "gauge_controller()",
		"493f4f74": "get_address(uint256)",
		"56059ffb": "get_gauges(address)",
		"37951049": "get_lp_token(address)",
		"a262904b": "get_registry()",
	},
}

// IRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IRegistryMetaData.ABI instead.
var IRegistryABI = IRegistryMetaData.ABI

// Deprecated: Use IRegistryMetaData.Sigs instead.
// IRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IRegistryFuncSigs = IRegistryMetaData.Sigs

// IRegistry is an auto generated Go binding around an Ethereum contract.
type IRegistry struct {
	IRegistryCaller     // Read-only binding to the contract
	IRegistryTransactor // Write-only binding to the contract
	IRegistryFilterer   // Log filterer for contract events
}

// IRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRegistrySession struct {
	Contract     *IRegistry        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRegistryCallerSession struct {
	Contract *IRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRegistryTransactorSession struct {
	Contract     *IRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRegistryRaw struct {
	Contract *IRegistry // Generic contract binding to access the raw methods on
}

// IRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRegistryCallerRaw struct {
	Contract *IRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRegistryTransactorRaw struct {
	Contract *IRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistry creates a new instance of IRegistry, bound to a specific deployed contract.
func NewIRegistry(address common.Address, backend bind.ContractBackend) (*IRegistry, error) {
	contract, err := bindIRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistry{IRegistryCaller: IRegistryCaller{contract: contract}, IRegistryTransactor: IRegistryTransactor{contract: contract}, IRegistryFilterer: IRegistryFilterer{contract: contract}}, nil
}

// NewIRegistryCaller creates a new read-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryCaller(address common.Address, caller bind.ContractCaller) (*IRegistryCaller, error) {
	contract, err := bindIRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryCaller{contract: contract}, nil
}

// NewIRegistryTransactor creates a new write-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistryTransactor, error) {
	contract, err := bindIRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryTransactor{contract: contract}, nil
}

// NewIRegistryFilterer creates a new log filterer instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistryFilterer, error) {
	contract, err := bindIRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistryFilterer{contract: contract}, nil
}

// bindIRegistry binds a generic wrapper to an already deployed contract.
func bindIRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.IRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transact(opts, method, params...)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_IRegistry *IRegistryCaller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "gauge_controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_IRegistry *IRegistrySession) GaugeController() (common.Address, error) {
	return _IRegistry.Contract.GaugeController(&_IRegistry.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_IRegistry *IRegistryCallerSession) GaugeController() (common.Address, error) {
	return _IRegistry.Contract.GaugeController(&_IRegistry.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x493f4f74.
//
// Solidity: function get_address(uint256 _id) view returns(address)
func (_IRegistry *IRegistryCaller) GetAddress(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "get_address", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x493f4f74.
//
// Solidity: function get_address(uint256 _id) view returns(address)
func (_IRegistry *IRegistrySession) GetAddress(_id *big.Int) (common.Address, error) {
	return _IRegistry.Contract.GetAddress(&_IRegistry.CallOpts, _id)
}

// GetAddress is a free data retrieval call binding the contract method 0x493f4f74.
//
// Solidity: function get_address(uint256 _id) view returns(address)
func (_IRegistry *IRegistryCallerSession) GetAddress(_id *big.Int) (common.Address, error) {
	return _IRegistry.Contract.GetAddress(&_IRegistry.CallOpts, _id)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address ) view returns(address[10], uint128[10])
func (_IRegistry *IRegistryCaller) GetGauges(opts *bind.CallOpts, arg0 common.Address) ([10]common.Address, [10]*big.Int, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "get_gauges", arg0)

	if err != nil {
		return *new([10]common.Address), *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)
	out1 := *abi.ConvertType(out[1], new([10]*big.Int)).(*[10]*big.Int)

	return out0, out1, err

}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address ) view returns(address[10], uint128[10])
func (_IRegistry *IRegistrySession) GetGauges(arg0 common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _IRegistry.Contract.GetGauges(&_IRegistry.CallOpts, arg0)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address ) view returns(address[10], uint128[10])
func (_IRegistry *IRegistryCallerSession) GetGauges(arg0 common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _IRegistry.Contract.GetGauges(&_IRegistry.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address ) view returns(address)
func (_IRegistry *IRegistryCaller) GetLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "get_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address ) view returns(address)
func (_IRegistry *IRegistrySession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _IRegistry.Contract.GetLpToken(&_IRegistry.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address ) view returns(address)
func (_IRegistry *IRegistryCallerSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _IRegistry.Contract.GetLpToken(&_IRegistry.CallOpts, arg0)
}

// GetRegistry is a free data retrieval call binding the contract method 0xa262904b.
//
// Solidity: function get_registry() view returns(address)
func (_IRegistry *IRegistryCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "get_registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0xa262904b.
//
// Solidity: function get_registry() view returns(address)
func (_IRegistry *IRegistrySession) GetRegistry() (common.Address, error) {
	return _IRegistry.Contract.GetRegistry(&_IRegistry.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0xa262904b.
//
// Solidity: function get_registry() view returns(address)
func (_IRegistry *IRegistryCallerSession) GetRegistry() (common.Address, error) {
	return _IRegistry.Contract.GetRegistry(&_IRegistry.CallOpts)
}

// IRewardFactoryMetaData contains all meta data concerning the IRewardFactory contract.
var IRewardFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CreateCrvRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CreateTokenRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"activeRewardCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addActiveReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeActiveReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"58cbfd45": "CreateCrvRewards(uint256,address)",
		"f8d6122e": "CreateTokenRewards(address,address,address)",
		"0d5843f7": "activeRewardCount(address)",
		"b7f927b1": "addActiveReward(address,uint256)",
		"ef9126ad": "removeActiveReward(address,uint256)",
		"b84614a5": "setAccess(address,bool)",
	},
}

// IRewardFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardFactoryMetaData.ABI instead.
var IRewardFactoryABI = IRewardFactoryMetaData.ABI

// Deprecated: Use IRewardFactoryMetaData.Sigs instead.
// IRewardFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IRewardFactoryFuncSigs = IRewardFactoryMetaData.Sigs

// IRewardFactory is an auto generated Go binding around an Ethereum contract.
type IRewardFactory struct {
	IRewardFactoryCaller     // Read-only binding to the contract
	IRewardFactoryTransactor // Write-only binding to the contract
	IRewardFactoryFilterer   // Log filterer for contract events
}

// IRewardFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardFactorySession struct {
	Contract     *IRewardFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardFactoryCallerSession struct {
	Contract *IRewardFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IRewardFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardFactoryTransactorSession struct {
	Contract     *IRewardFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IRewardFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardFactoryRaw struct {
	Contract *IRewardFactory // Generic contract binding to access the raw methods on
}

// IRewardFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardFactoryCallerRaw struct {
	Contract *IRewardFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardFactoryTransactorRaw struct {
	Contract *IRewardFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardFactory creates a new instance of IRewardFactory, bound to a specific deployed contract.
func NewIRewardFactory(address common.Address, backend bind.ContractBackend) (*IRewardFactory, error) {
	contract, err := bindIRewardFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardFactory{IRewardFactoryCaller: IRewardFactoryCaller{contract: contract}, IRewardFactoryTransactor: IRewardFactoryTransactor{contract: contract}, IRewardFactoryFilterer: IRewardFactoryFilterer{contract: contract}}, nil
}

// NewIRewardFactoryCaller creates a new read-only instance of IRewardFactory, bound to a specific deployed contract.
func NewIRewardFactoryCaller(address common.Address, caller bind.ContractCaller) (*IRewardFactoryCaller, error) {
	contract, err := bindIRewardFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardFactoryCaller{contract: contract}, nil
}

// NewIRewardFactoryTransactor creates a new write-only instance of IRewardFactory, bound to a specific deployed contract.
func NewIRewardFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardFactoryTransactor, error) {
	contract, err := bindIRewardFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardFactoryTransactor{contract: contract}, nil
}

// NewIRewardFactoryFilterer creates a new log filterer instance of IRewardFactory, bound to a specific deployed contract.
func NewIRewardFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardFactoryFilterer, error) {
	contract, err := bindIRewardFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardFactoryFilterer{contract: contract}, nil
}

// bindIRewardFactory binds a generic wrapper to an already deployed contract.
func bindIRewardFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRewardFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardFactory *IRewardFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardFactory.Contract.IRewardFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardFactory *IRewardFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardFactory.Contract.IRewardFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardFactory *IRewardFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardFactory.Contract.IRewardFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardFactory *IRewardFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardFactory *IRewardFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardFactory *IRewardFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardFactory.Contract.contract.Transact(opts, method, params...)
}

// ActiveRewardCount is a free data retrieval call binding the contract method 0x0d5843f7.
//
// Solidity: function activeRewardCount(address ) view returns(uint256)
func (_IRewardFactory *IRewardFactoryCaller) ActiveRewardCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardFactory.contract.Call(opts, &out, "activeRewardCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveRewardCount is a free data retrieval call binding the contract method 0x0d5843f7.
//
// Solidity: function activeRewardCount(address ) view returns(uint256)
func (_IRewardFactory *IRewardFactorySession) ActiveRewardCount(arg0 common.Address) (*big.Int, error) {
	return _IRewardFactory.Contract.ActiveRewardCount(&_IRewardFactory.CallOpts, arg0)
}

// ActiveRewardCount is a free data retrieval call binding the contract method 0x0d5843f7.
//
// Solidity: function activeRewardCount(address ) view returns(uint256)
func (_IRewardFactory *IRewardFactoryCallerSession) ActiveRewardCount(arg0 common.Address) (*big.Int, error) {
	return _IRewardFactory.Contract.ActiveRewardCount(&_IRewardFactory.CallOpts, arg0)
}

// CreateCrvRewards is a paid mutator transaction binding the contract method 0x58cbfd45.
//
// Solidity: function CreateCrvRewards(uint256 , address ) returns(address)
func (_IRewardFactory *IRewardFactoryTransactor) CreateCrvRewards(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.contract.Transact(opts, "CreateCrvRewards", arg0, arg1)
}

// CreateCrvRewards is a paid mutator transaction binding the contract method 0x58cbfd45.
//
// Solidity: function CreateCrvRewards(uint256 , address ) returns(address)
func (_IRewardFactory *IRewardFactorySession) CreateCrvRewards(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.Contract.CreateCrvRewards(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// CreateCrvRewards is a paid mutator transaction binding the contract method 0x58cbfd45.
//
// Solidity: function CreateCrvRewards(uint256 , address ) returns(address)
func (_IRewardFactory *IRewardFactoryTransactorSession) CreateCrvRewards(arg0 *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.Contract.CreateCrvRewards(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// CreateTokenRewards is a paid mutator transaction binding the contract method 0xf8d6122e.
//
// Solidity: function CreateTokenRewards(address , address , address ) returns(address)
func (_IRewardFactory *IRewardFactoryTransactor) CreateTokenRewards(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.contract.Transact(opts, "CreateTokenRewards", arg0, arg1, arg2)
}

// CreateTokenRewards is a paid mutator transaction binding the contract method 0xf8d6122e.
//
// Solidity: function CreateTokenRewards(address , address , address ) returns(address)
func (_IRewardFactory *IRewardFactorySession) CreateTokenRewards(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.Contract.CreateTokenRewards(&_IRewardFactory.TransactOpts, arg0, arg1, arg2)
}

// CreateTokenRewards is a paid mutator transaction binding the contract method 0xf8d6122e.
//
// Solidity: function CreateTokenRewards(address , address , address ) returns(address)
func (_IRewardFactory *IRewardFactoryTransactorSession) CreateTokenRewards(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _IRewardFactory.Contract.CreateTokenRewards(&_IRewardFactory.TransactOpts, arg0, arg1, arg2)
}

// AddActiveReward is a paid mutator transaction binding the contract method 0xb7f927b1.
//
// Solidity: function addActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactoryTransactor) AddActiveReward(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.contract.Transact(opts, "addActiveReward", arg0, arg1)
}

// AddActiveReward is a paid mutator transaction binding the contract method 0xb7f927b1.
//
// Solidity: function addActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactorySession) AddActiveReward(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.Contract.AddActiveReward(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// AddActiveReward is a paid mutator transaction binding the contract method 0xb7f927b1.
//
// Solidity: function addActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactoryTransactorSession) AddActiveReward(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.Contract.AddActiveReward(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// RemoveActiveReward is a paid mutator transaction binding the contract method 0xef9126ad.
//
// Solidity: function removeActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactoryTransactor) RemoveActiveReward(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.contract.Transact(opts, "removeActiveReward", arg0, arg1)
}

// RemoveActiveReward is a paid mutator transaction binding the contract method 0xef9126ad.
//
// Solidity: function removeActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactorySession) RemoveActiveReward(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.Contract.RemoveActiveReward(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// RemoveActiveReward is a paid mutator transaction binding the contract method 0xef9126ad.
//
// Solidity: function removeActiveReward(address , uint256 ) returns(bool)
func (_IRewardFactory *IRewardFactoryTransactorSession) RemoveActiveReward(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewardFactory.Contract.RemoveActiveReward(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// SetAccess is a paid mutator transaction binding the contract method 0xb84614a5.
//
// Solidity: function setAccess(address , bool ) returns()
func (_IRewardFactory *IRewardFactoryTransactor) SetAccess(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IRewardFactory.contract.Transact(opts, "setAccess", arg0, arg1)
}

// SetAccess is a paid mutator transaction binding the contract method 0xb84614a5.
//
// Solidity: function setAccess(address , bool ) returns()
func (_IRewardFactory *IRewardFactorySession) SetAccess(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IRewardFactory.Contract.SetAccess(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// SetAccess is a paid mutator transaction binding the contract method 0xb84614a5.
//
// Solidity: function setAccess(address , bool ) returns()
func (_IRewardFactory *IRewardFactoryTransactorSession) SetAccess(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IRewardFactory.Contract.SetAccess(&_IRewardFactory.TransactOpts, arg0, arg1)
}

// IRewardsMetaData contains all meta data concerning the IRewards contract.
var IRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addExtraReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5e43c47b": "addExtraReward(address)",
		"b42652e9": "exit(address)",
		"c00007b0": "getReward(address)",
		"3c6b16ab": "notifyRewardAmount(uint256)",
		"590a41f5": "queueNewRewards(uint256)",
		"adc9772e": "stake(address,uint256)",
		"2ee40908": "stakeFor(address,uint256)",
		"72f702f3": "stakingToken()",
		"f3fef3a3": "withdraw(address,uint256)",
	},
}

// IRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardsMetaData.ABI instead.
var IRewardsABI = IRewardsMetaData.ABI

// Deprecated: Use IRewardsMetaData.Sigs instead.
// IRewardsFuncSigs maps the 4-byte function signature to its string representation.
var IRewardsFuncSigs = IRewardsMetaData.Sigs

// IRewards is an auto generated Go binding around an Ethereum contract.
type IRewards struct {
	IRewardsCaller     // Read-only binding to the contract
	IRewardsTransactor // Write-only binding to the contract
	IRewardsFilterer   // Log filterer for contract events
}

// IRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardsSession struct {
	Contract     *IRewards         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardsCallerSession struct {
	Contract *IRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardsTransactorSession struct {
	Contract     *IRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardsRaw struct {
	Contract *IRewards // Generic contract binding to access the raw methods on
}

// IRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardsCallerRaw struct {
	Contract *IRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardsTransactorRaw struct {
	Contract *IRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewards creates a new instance of IRewards, bound to a specific deployed contract.
func NewIRewards(address common.Address, backend bind.ContractBackend) (*IRewards, error) {
	contract, err := bindIRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewards{IRewardsCaller: IRewardsCaller{contract: contract}, IRewardsTransactor: IRewardsTransactor{contract: contract}, IRewardsFilterer: IRewardsFilterer{contract: contract}}, nil
}

// NewIRewardsCaller creates a new read-only instance of IRewards, bound to a specific deployed contract.
func NewIRewardsCaller(address common.Address, caller bind.ContractCaller) (*IRewardsCaller, error) {
	contract, err := bindIRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsCaller{contract: contract}, nil
}

// NewIRewardsTransactor creates a new write-only instance of IRewards, bound to a specific deployed contract.
func NewIRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardsTransactor, error) {
	contract, err := bindIRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsTransactor{contract: contract}, nil
}

// NewIRewardsFilterer creates a new log filterer instance of IRewards, bound to a specific deployed contract.
func NewIRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardsFilterer, error) {
	contract, err := bindIRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardsFilterer{contract: contract}, nil
}

// bindIRewards binds a generic wrapper to an already deployed contract.
func bindIRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewards *IRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewards.Contract.IRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewards *IRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewards.Contract.IRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewards *IRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewards.Contract.IRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewards *IRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewards *IRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewards *IRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewards.Contract.contract.Transact(opts, method, params...)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address ) returns()
func (_IRewards *IRewardsTransactor) AddExtraReward(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "addExtraReward", arg0)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address ) returns()
func (_IRewards *IRewardsSession) AddExtraReward(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.AddExtraReward(&_IRewards.TransactOpts, arg0)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address ) returns()
func (_IRewards *IRewardsTransactorSession) AddExtraReward(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.AddExtraReward(&_IRewards.TransactOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xb42652e9.
//
// Solidity: function exit(address ) returns()
func (_IRewards *IRewardsTransactor) Exit(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "exit", arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xb42652e9.
//
// Solidity: function exit(address ) returns()
func (_IRewards *IRewardsSession) Exit(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.Exit(&_IRewards.TransactOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xb42652e9.
//
// Solidity: function exit(address ) returns()
func (_IRewards *IRewardsTransactorSession) Exit(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.Exit(&_IRewards.TransactOpts, arg0)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address ) returns()
func (_IRewards *IRewardsTransactor) GetReward(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "getReward", arg0)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address ) returns()
func (_IRewards *IRewardsSession) GetReward(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.GetReward(&_IRewards.TransactOpts, arg0)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address ) returns()
func (_IRewards *IRewardsTransactorSession) GetReward(arg0 common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.GetReward(&_IRewards.TransactOpts, arg0)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 ) returns()
func (_IRewards *IRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "notifyRewardAmount", arg0)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 ) returns()
func (_IRewards *IRewardsSession) NotifyRewardAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.NotifyRewardAmount(&_IRewards.TransactOpts, arg0)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 ) returns()
func (_IRewards *IRewardsTransactorSession) NotifyRewardAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.NotifyRewardAmount(&_IRewards.TransactOpts, arg0)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 ) returns()
func (_IRewards *IRewardsTransactor) QueueNewRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "queueNewRewards", arg0)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 ) returns()
func (_IRewards *IRewardsSession) QueueNewRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.QueueNewRewards(&_IRewards.TransactOpts, arg0)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 ) returns()
func (_IRewards *IRewardsTransactorSession) QueueNewRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.QueueNewRewards(&_IRewards.TransactOpts, arg0)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address , uint256 ) returns()
func (_IRewards *IRewardsTransactor) Stake(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "stake", arg0, arg1)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address , uint256 ) returns()
func (_IRewards *IRewardsSession) Stake(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.Stake(&_IRewards.TransactOpts, arg0, arg1)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address , uint256 ) returns()
func (_IRewards *IRewardsTransactorSession) Stake(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.Stake(&_IRewards.TransactOpts, arg0, arg1)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) returns()
func (_IRewards *IRewardsTransactor) StakeFor(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "stakeFor", arg0, arg1)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) returns()
func (_IRewards *IRewardsSession) StakeFor(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.StakeFor(&_IRewards.TransactOpts, arg0, arg1)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) returns()
func (_IRewards *IRewardsTransactorSession) StakeFor(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.StakeFor(&_IRewards.TransactOpts, arg0, arg1)
}

// StakingToken is a paid mutator transaction binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() returns(address)
func (_IRewards *IRewardsTransactor) StakingToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "stakingToken")
}

// StakingToken is a paid mutator transaction binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() returns(address)
func (_IRewards *IRewardsSession) StakingToken() (*types.Transaction, error) {
	return _IRewards.Contract.StakingToken(&_IRewards.TransactOpts)
}

// StakingToken is a paid mutator transaction binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() returns(address)
func (_IRewards *IRewardsTransactorSession) StakingToken() (*types.Transaction, error) {
	return _IRewards.Contract.StakingToken(&_IRewards.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) returns()
func (_IRewards *IRewardsTransactor) Withdraw(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "withdraw", arg0, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) returns()
func (_IRewards *IRewardsSession) Withdraw(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.Withdraw(&_IRewards.TransactOpts, arg0, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address , uint256 ) returns()
func (_IRewards *IRewardsTransactorSession) Withdraw(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.Withdraw(&_IRewards.TransactOpts, arg0, arg1)
}

// IStakerMetaData contains all meta data concerning the IStaker contract.
var IStakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOfPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimCrv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"createLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increaseAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increaseTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setStashAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteGaugeWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b0f63794": "balanceOfPool(address)",
		"3fe9bc06": "claimCrv(address)",
		"2dbfa735": "claimFees(address,address)",
		"ef5cfb8c": "claimRewards(address)",
		"b52c05fe": "createLock(uint256,uint256)",
		"f9609f08": "deposit(address,address)",
		"b61d27f6": "execute(address,uint256,bytes)",
		"15456eba": "increaseAmount(uint256)",
		"3c9a2a1a": "increaseTime(uint256)",
		"570ca735": "operator()",
		"86d1a69f": "release()",
		"fa3964b2": "setStashAccess(address,bool)",
		"e2cdd42a": "vote(uint256,address,bool)",
		"5d7e9bcb": "voteGaugeWeight(address,uint256)",
		"51cff8d9": "withdraw(address)",
		"d9caed12": "withdraw(address,address,uint256)",
		"09cae2c8": "withdrawAll(address,address)",
	},
}

// IStakerABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakerMetaData.ABI instead.
var IStakerABI = IStakerMetaData.ABI

// Deprecated: Use IStakerMetaData.Sigs instead.
// IStakerFuncSigs maps the 4-byte function signature to its string representation.
var IStakerFuncSigs = IStakerMetaData.Sigs

// IStaker is an auto generated Go binding around an Ethereum contract.
type IStaker struct {
	IStakerCaller     // Read-only binding to the contract
	IStakerTransactor // Write-only binding to the contract
	IStakerFilterer   // Log filterer for contract events
}

// IStakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakerSession struct {
	Contract     *IStaker          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakerCallerSession struct {
	Contract *IStakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IStakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakerTransactorSession struct {
	Contract     *IStakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IStakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakerRaw struct {
	Contract *IStaker // Generic contract binding to access the raw methods on
}

// IStakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakerCallerRaw struct {
	Contract *IStakerCaller // Generic read-only contract binding to access the raw methods on
}

// IStakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakerTransactorRaw struct {
	Contract *IStakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaker creates a new instance of IStaker, bound to a specific deployed contract.
func NewIStaker(address common.Address, backend bind.ContractBackend) (*IStaker, error) {
	contract, err := bindIStaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaker{IStakerCaller: IStakerCaller{contract: contract}, IStakerTransactor: IStakerTransactor{contract: contract}, IStakerFilterer: IStakerFilterer{contract: contract}}, nil
}

// NewIStakerCaller creates a new read-only instance of IStaker, bound to a specific deployed contract.
func NewIStakerCaller(address common.Address, caller bind.ContractCaller) (*IStakerCaller, error) {
	contract, err := bindIStaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakerCaller{contract: contract}, nil
}

// NewIStakerTransactor creates a new write-only instance of IStaker, bound to a specific deployed contract.
func NewIStakerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakerTransactor, error) {
	contract, err := bindIStaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakerTransactor{contract: contract}, nil
}

// NewIStakerFilterer creates a new log filterer instance of IStaker, bound to a specific deployed contract.
func NewIStakerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakerFilterer, error) {
	contract, err := bindIStaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakerFilterer{contract: contract}, nil
}

// bindIStaker binds a generic wrapper to an already deployed contract.
func bindIStaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaker *IStakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaker.Contract.IStakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaker *IStakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaker.Contract.IStakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaker *IStakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaker.Contract.IStakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaker *IStakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaker *IStakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaker *IStakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaker.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfPool is a free data retrieval call binding the contract method 0xb0f63794.
//
// Solidity: function balanceOfPool(address ) view returns(uint256)
func (_IStaker *IStakerCaller) BalanceOfPool(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStaker.contract.Call(opts, &out, "balanceOfPool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfPool is a free data retrieval call binding the contract method 0xb0f63794.
//
// Solidity: function balanceOfPool(address ) view returns(uint256)
func (_IStaker *IStakerSession) BalanceOfPool(arg0 common.Address) (*big.Int, error) {
	return _IStaker.Contract.BalanceOfPool(&_IStaker.CallOpts, arg0)
}

// BalanceOfPool is a free data retrieval call binding the contract method 0xb0f63794.
//
// Solidity: function balanceOfPool(address ) view returns(uint256)
func (_IStaker *IStakerCallerSession) BalanceOfPool(arg0 common.Address) (*big.Int, error) {
	return _IStaker.Contract.BalanceOfPool(&_IStaker.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IStaker *IStakerCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStaker.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IStaker *IStakerSession) Operator() (common.Address, error) {
	return _IStaker.Contract.Operator(&_IStaker.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_IStaker *IStakerCallerSession) Operator() (common.Address, error) {
	return _IStaker.Contract.Operator(&_IStaker.CallOpts)
}

// ClaimCrv is a paid mutator transaction binding the contract method 0x3fe9bc06.
//
// Solidity: function claimCrv(address ) returns(uint256)
func (_IStaker *IStakerTransactor) ClaimCrv(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "claimCrv", arg0)
}

// ClaimCrv is a paid mutator transaction binding the contract method 0x3fe9bc06.
//
// Solidity: function claimCrv(address ) returns(uint256)
func (_IStaker *IStakerSession) ClaimCrv(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimCrv(&_IStaker.TransactOpts, arg0)
}

// ClaimCrv is a paid mutator transaction binding the contract method 0x3fe9bc06.
//
// Solidity: function claimCrv(address ) returns(uint256)
func (_IStaker *IStakerTransactorSession) ClaimCrv(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimCrv(&_IStaker.TransactOpts, arg0)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x2dbfa735.
//
// Solidity: function claimFees(address , address ) returns()
func (_IStaker *IStakerTransactor) ClaimFees(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "claimFees", arg0, arg1)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x2dbfa735.
//
// Solidity: function claimFees(address , address ) returns()
func (_IStaker *IStakerSession) ClaimFees(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimFees(&_IStaker.TransactOpts, arg0, arg1)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x2dbfa735.
//
// Solidity: function claimFees(address , address ) returns()
func (_IStaker *IStakerTransactorSession) ClaimFees(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimFees(&_IStaker.TransactOpts, arg0, arg1)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address ) returns()
func (_IStaker *IStakerTransactor) ClaimRewards(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "claimRewards", arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address ) returns()
func (_IStaker *IStakerSession) ClaimRewards(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimRewards(&_IStaker.TransactOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address ) returns()
func (_IStaker *IStakerTransactorSession) ClaimRewards(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.ClaimRewards(&_IStaker.TransactOpts, arg0)
}

// CreateLock is a paid mutator transaction binding the contract method 0xb52c05fe.
//
// Solidity: function createLock(uint256 , uint256 ) returns()
func (_IStaker *IStakerTransactor) CreateLock(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "createLock", arg0, arg1)
}

// CreateLock is a paid mutator transaction binding the contract method 0xb52c05fe.
//
// Solidity: function createLock(uint256 , uint256 ) returns()
func (_IStaker *IStakerSession) CreateLock(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.CreateLock(&_IStaker.TransactOpts, arg0, arg1)
}

// CreateLock is a paid mutator transaction binding the contract method 0xb52c05fe.
//
// Solidity: function createLock(uint256 , uint256 ) returns()
func (_IStaker *IStakerTransactorSession) CreateLock(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.CreateLock(&_IStaker.TransactOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address , address ) returns()
func (_IStaker *IStakerTransactor) Deposit(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "deposit", arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address , address ) returns()
func (_IStaker *IStakerSession) Deposit(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.Deposit(&_IStaker.TransactOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address , address ) returns()
func (_IStaker *IStakerTransactorSession) Deposit(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.Deposit(&_IStaker.TransactOpts, arg0, arg1)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_IStaker *IStakerTransactor) Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "execute", _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_IStaker *IStakerSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IStaker.Contract.Execute(&_IStaker.TransactOpts, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_IStaker *IStakerTransactorSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IStaker.Contract.Execute(&_IStaker.TransactOpts, _to, _value, _data)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 ) returns()
func (_IStaker *IStakerTransactor) IncreaseAmount(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "increaseAmount", arg0)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 ) returns()
func (_IStaker *IStakerSession) IncreaseAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.IncreaseAmount(&_IStaker.TransactOpts, arg0)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x15456eba.
//
// Solidity: function increaseAmount(uint256 ) returns()
func (_IStaker *IStakerTransactorSession) IncreaseAmount(arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.IncreaseAmount(&_IStaker.TransactOpts, arg0)
}

// IncreaseTime is a paid mutator transaction binding the contract method 0x3c9a2a1a.
//
// Solidity: function increaseTime(uint256 ) returns()
func (_IStaker *IStakerTransactor) IncreaseTime(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "increaseTime", arg0)
}

// IncreaseTime is a paid mutator transaction binding the contract method 0x3c9a2a1a.
//
// Solidity: function increaseTime(uint256 ) returns()
func (_IStaker *IStakerSession) IncreaseTime(arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.IncreaseTime(&_IStaker.TransactOpts, arg0)
}

// IncreaseTime is a paid mutator transaction binding the contract method 0x3c9a2a1a.
//
// Solidity: function increaseTime(uint256 ) returns()
func (_IStaker *IStakerTransactorSession) IncreaseTime(arg0 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.IncreaseTime(&_IStaker.TransactOpts, arg0)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_IStaker *IStakerTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_IStaker *IStakerSession) Release() (*types.Transaction, error) {
	return _IStaker.Contract.Release(&_IStaker.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_IStaker *IStakerTransactorSession) Release() (*types.Transaction, error) {
	return _IStaker.Contract.Release(&_IStaker.TransactOpts)
}

// SetStashAccess is a paid mutator transaction binding the contract method 0xfa3964b2.
//
// Solidity: function setStashAccess(address , bool ) returns()
func (_IStaker *IStakerTransactor) SetStashAccess(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "setStashAccess", arg0, arg1)
}

// SetStashAccess is a paid mutator transaction binding the contract method 0xfa3964b2.
//
// Solidity: function setStashAccess(address , bool ) returns()
func (_IStaker *IStakerSession) SetStashAccess(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IStaker.Contract.SetStashAccess(&_IStaker.TransactOpts, arg0, arg1)
}

// SetStashAccess is a paid mutator transaction binding the contract method 0xfa3964b2.
//
// Solidity: function setStashAccess(address , bool ) returns()
func (_IStaker *IStakerTransactorSession) SetStashAccess(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _IStaker.Contract.SetStashAccess(&_IStaker.TransactOpts, arg0, arg1)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 , address , bool ) returns()
func (_IStaker *IStakerTransactor) Vote(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address, arg2 bool) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "vote", arg0, arg1, arg2)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 , address , bool ) returns()
func (_IStaker *IStakerSession) Vote(arg0 *big.Int, arg1 common.Address, arg2 bool) (*types.Transaction, error) {
	return _IStaker.Contract.Vote(&_IStaker.TransactOpts, arg0, arg1, arg2)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 , address , bool ) returns()
func (_IStaker *IStakerTransactorSession) Vote(arg0 *big.Int, arg1 common.Address, arg2 bool) (*types.Transaction, error) {
	return _IStaker.Contract.Vote(&_IStaker.TransactOpts, arg0, arg1, arg2)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0x5d7e9bcb.
//
// Solidity: function voteGaugeWeight(address , uint256 ) returns()
func (_IStaker *IStakerTransactor) VoteGaugeWeight(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "voteGaugeWeight", arg0, arg1)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0x5d7e9bcb.
//
// Solidity: function voteGaugeWeight(address , uint256 ) returns()
func (_IStaker *IStakerSession) VoteGaugeWeight(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.VoteGaugeWeight(&_IStaker.TransactOpts, arg0, arg1)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0x5d7e9bcb.
//
// Solidity: function voteGaugeWeight(address , uint256 ) returns()
func (_IStaker *IStakerTransactorSession) VoteGaugeWeight(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.VoteGaugeWeight(&_IStaker.TransactOpts, arg0, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address ) returns()
func (_IStaker *IStakerTransactor) Withdraw(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address ) returns()
func (_IStaker *IStakerSession) Withdraw(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.Withdraw(&_IStaker.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address ) returns()
func (_IStaker *IStakerTransactorSession) Withdraw(arg0 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.Withdraw(&_IStaker.TransactOpts, arg0)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address , address , uint256 ) returns()
func (_IStaker *IStakerTransactor) Withdraw0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "withdraw0", arg0, arg1, arg2)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address , address , uint256 ) returns()
func (_IStaker *IStakerSession) Withdraw0(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.Withdraw0(&_IStaker.TransactOpts, arg0, arg1, arg2)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address , address , uint256 ) returns()
func (_IStaker *IStakerTransactorSession) Withdraw0(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _IStaker.Contract.Withdraw0(&_IStaker.TransactOpts, arg0, arg1, arg2)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address , address ) returns()
func (_IStaker *IStakerTransactor) WithdrawAll(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.contract.Transact(opts, "withdrawAll", arg0, arg1)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address , address ) returns()
func (_IStaker *IStakerSession) WithdrawAll(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.WithdrawAll(&_IStaker.TransactOpts, arg0, arg1)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x09cae2c8.
//
// Solidity: function withdrawAll(address , address ) returns()
func (_IStaker *IStakerTransactorSession) WithdrawAll(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _IStaker.Contract.WithdrawAll(&_IStaker.TransactOpts, arg0, arg1)
}

// IStashMetaData contains all meta data concerning the IStash contract.
var IStashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processStash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stashRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"372500ab": "claimRewards()",
		"ca8b0176": "processStash()",
		"b87bd481": "stashRewards()",
	},
}

// IStashABI is the input ABI used to generate the binding from.
// Deprecated: Use IStashMetaData.ABI instead.
var IStashABI = IStashMetaData.ABI

// Deprecated: Use IStashMetaData.Sigs instead.
// IStashFuncSigs maps the 4-byte function signature to its string representation.
var IStashFuncSigs = IStashMetaData.Sigs

// IStash is an auto generated Go binding around an Ethereum contract.
type IStash struct {
	IStashCaller     // Read-only binding to the contract
	IStashTransactor // Write-only binding to the contract
	IStashFilterer   // Log filterer for contract events
}

// IStashCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStashSession struct {
	Contract     *IStash           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStashCallerSession struct {
	Contract *IStashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IStashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStashTransactorSession struct {
	Contract     *IStashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStashRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStashRaw struct {
	Contract *IStash // Generic contract binding to access the raw methods on
}

// IStashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStashCallerRaw struct {
	Contract *IStashCaller // Generic read-only contract binding to access the raw methods on
}

// IStashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStashTransactorRaw struct {
	Contract *IStashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStash creates a new instance of IStash, bound to a specific deployed contract.
func NewIStash(address common.Address, backend bind.ContractBackend) (*IStash, error) {
	contract, err := bindIStash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStash{IStashCaller: IStashCaller{contract: contract}, IStashTransactor: IStashTransactor{contract: contract}, IStashFilterer: IStashFilterer{contract: contract}}, nil
}

// NewIStashCaller creates a new read-only instance of IStash, bound to a specific deployed contract.
func NewIStashCaller(address common.Address, caller bind.ContractCaller) (*IStashCaller, error) {
	contract, err := bindIStash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStashCaller{contract: contract}, nil
}

// NewIStashTransactor creates a new write-only instance of IStash, bound to a specific deployed contract.
func NewIStashTransactor(address common.Address, transactor bind.ContractTransactor) (*IStashTransactor, error) {
	contract, err := bindIStash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStashTransactor{contract: contract}, nil
}

// NewIStashFilterer creates a new log filterer instance of IStash, bound to a specific deployed contract.
func NewIStashFilterer(address common.Address, filterer bind.ContractFilterer) (*IStashFilterer, error) {
	contract, err := bindIStash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStashFilterer{contract: contract}, nil
}

// bindIStash binds a generic wrapper to an already deployed contract.
func bindIStash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStash *IStashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStash.Contract.IStashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStash *IStashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStash.Contract.IStashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStash *IStashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStash.Contract.IStashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStash *IStashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStash *IStashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStash *IStashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStash.Contract.contract.Transact(opts, method, params...)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns(bool)
func (_IStash *IStashTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStash.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns(bool)
func (_IStash *IStashSession) ClaimRewards() (*types.Transaction, error) {
	return _IStash.Contract.ClaimRewards(&_IStash.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns(bool)
func (_IStash *IStashTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _IStash.Contract.ClaimRewards(&_IStash.TransactOpts)
}

// ProcessStash is a paid mutator transaction binding the contract method 0xca8b0176.
//
// Solidity: function processStash() returns(bool)
func (_IStash *IStashTransactor) ProcessStash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStash.contract.Transact(opts, "processStash")
}

// ProcessStash is a paid mutator transaction binding the contract method 0xca8b0176.
//
// Solidity: function processStash() returns(bool)
func (_IStash *IStashSession) ProcessStash() (*types.Transaction, error) {
	return _IStash.Contract.ProcessStash(&_IStash.TransactOpts)
}

// ProcessStash is a paid mutator transaction binding the contract method 0xca8b0176.
//
// Solidity: function processStash() returns(bool)
func (_IStash *IStashTransactorSession) ProcessStash() (*types.Transaction, error) {
	return _IStash.Contract.ProcessStash(&_IStash.TransactOpts)
}

// StashRewards is a paid mutator transaction binding the contract method 0xb87bd481.
//
// Solidity: function stashRewards() returns(bool)
func (_IStash *IStashTransactor) StashRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStash.contract.Transact(opts, "stashRewards")
}

// StashRewards is a paid mutator transaction binding the contract method 0xb87bd481.
//
// Solidity: function stashRewards() returns(bool)
func (_IStash *IStashSession) StashRewards() (*types.Transaction, error) {
	return _IStash.Contract.StashRewards(&_IStash.TransactOpts)
}

// StashRewards is a paid mutator transaction binding the contract method 0xb87bd481.
//
// Solidity: function stashRewards() returns(bool)
func (_IStash *IStashTransactorSession) StashRewards() (*types.Transaction, error) {
	return _IStash.Contract.StashRewards(&_IStash.TransactOpts)
}

// IStashFactoryMetaData contains all meta data concerning the IStashFactory contract.
var IStashFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CreateStash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"99cb12de": "CreateStash(uint256,address,address,uint256)",
	},
}

// IStashFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IStashFactoryMetaData.ABI instead.
var IStashFactoryABI = IStashFactoryMetaData.ABI

// Deprecated: Use IStashFactoryMetaData.Sigs instead.
// IStashFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IStashFactoryFuncSigs = IStashFactoryMetaData.Sigs

// IStashFactory is an auto generated Go binding around an Ethereum contract.
type IStashFactory struct {
	IStashFactoryCaller     // Read-only binding to the contract
	IStashFactoryTransactor // Write-only binding to the contract
	IStashFactoryFilterer   // Log filterer for contract events
}

// IStashFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStashFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStashFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStashFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStashFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStashFactorySession struct {
	Contract     *IStashFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStashFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStashFactoryCallerSession struct {
	Contract *IStashFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IStashFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStashFactoryTransactorSession struct {
	Contract     *IStashFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IStashFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStashFactoryRaw struct {
	Contract *IStashFactory // Generic contract binding to access the raw methods on
}

// IStashFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStashFactoryCallerRaw struct {
	Contract *IStashFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IStashFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStashFactoryTransactorRaw struct {
	Contract *IStashFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStashFactory creates a new instance of IStashFactory, bound to a specific deployed contract.
func NewIStashFactory(address common.Address, backend bind.ContractBackend) (*IStashFactory, error) {
	contract, err := bindIStashFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStashFactory{IStashFactoryCaller: IStashFactoryCaller{contract: contract}, IStashFactoryTransactor: IStashFactoryTransactor{contract: contract}, IStashFactoryFilterer: IStashFactoryFilterer{contract: contract}}, nil
}

// NewIStashFactoryCaller creates a new read-only instance of IStashFactory, bound to a specific deployed contract.
func NewIStashFactoryCaller(address common.Address, caller bind.ContractCaller) (*IStashFactoryCaller, error) {
	contract, err := bindIStashFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStashFactoryCaller{contract: contract}, nil
}

// NewIStashFactoryTransactor creates a new write-only instance of IStashFactory, bound to a specific deployed contract.
func NewIStashFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IStashFactoryTransactor, error) {
	contract, err := bindIStashFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStashFactoryTransactor{contract: contract}, nil
}

// NewIStashFactoryFilterer creates a new log filterer instance of IStashFactory, bound to a specific deployed contract.
func NewIStashFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IStashFactoryFilterer, error) {
	contract, err := bindIStashFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStashFactoryFilterer{contract: contract}, nil
}

// bindIStashFactory binds a generic wrapper to an already deployed contract.
func bindIStashFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStashFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStashFactory *IStashFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStashFactory.Contract.IStashFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStashFactory *IStashFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStashFactory.Contract.IStashFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStashFactory *IStashFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStashFactory.Contract.IStashFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStashFactory *IStashFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStashFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStashFactory *IStashFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStashFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStashFactory *IStashFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStashFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateStash is a paid mutator transaction binding the contract method 0x99cb12de.
//
// Solidity: function CreateStash(uint256 , address , address , uint256 ) returns(address)
func (_IStashFactory *IStashFactoryTransactor) CreateStash(opts *bind.TransactOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _IStashFactory.contract.Transact(opts, "CreateStash", arg0, arg1, arg2, arg3)
}

// CreateStash is a paid mutator transaction binding the contract method 0x99cb12de.
//
// Solidity: function CreateStash(uint256 , address , address , uint256 ) returns(address)
func (_IStashFactory *IStashFactorySession) CreateStash(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _IStashFactory.Contract.CreateStash(&_IStashFactory.TransactOpts, arg0, arg1, arg2, arg3)
}

// CreateStash is a paid mutator transaction binding the contract method 0x99cb12de.
//
// Solidity: function CreateStash(uint256 , address , address , uint256 ) returns(address)
func (_IStashFactory *IStashFactoryTransactorSession) CreateStash(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _IStashFactory.Contract.CreateStash(&_IStashFactory.TransactOpts, arg0, arg1, arg2, arg3)
}

// ITokenFactoryMetaData contains all meta data concerning the ITokenFactory contract.
var ITokenFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CreateDepositToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"114a899c": "CreateDepositToken(address)",
	},
}

// ITokenFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenFactoryMetaData.ABI instead.
var ITokenFactoryABI = ITokenFactoryMetaData.ABI

// Deprecated: Use ITokenFactoryMetaData.Sigs instead.
// ITokenFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFactoryFuncSigs = ITokenFactoryMetaData.Sigs

// ITokenFactory is an auto generated Go binding around an Ethereum contract.
type ITokenFactory struct {
	ITokenFactoryCaller     // Read-only binding to the contract
	ITokenFactoryTransactor // Write-only binding to the contract
	ITokenFactoryFilterer   // Log filterer for contract events
}

// ITokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenFactorySession struct {
	Contract     *ITokenFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenFactoryCallerSession struct {
	Contract *ITokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ITokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenFactoryTransactorSession struct {
	Contract     *ITokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ITokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenFactoryRaw struct {
	Contract *ITokenFactory // Generic contract binding to access the raw methods on
}

// ITokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenFactoryCallerRaw struct {
	Contract *ITokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenFactoryTransactorRaw struct {
	Contract *ITokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenFactory creates a new instance of ITokenFactory, bound to a specific deployed contract.
func NewITokenFactory(address common.Address, backend bind.ContractBackend) (*ITokenFactory, error) {
	contract, err := bindITokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenFactory{ITokenFactoryCaller: ITokenFactoryCaller{contract: contract}, ITokenFactoryTransactor: ITokenFactoryTransactor{contract: contract}, ITokenFactoryFilterer: ITokenFactoryFilterer{contract: contract}}, nil
}

// NewITokenFactoryCaller creates a new read-only instance of ITokenFactory, bound to a specific deployed contract.
func NewITokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*ITokenFactoryCaller, error) {
	contract, err := bindITokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenFactoryCaller{contract: contract}, nil
}

// NewITokenFactoryTransactor creates a new write-only instance of ITokenFactory, bound to a specific deployed contract.
func NewITokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenFactoryTransactor, error) {
	contract, err := bindITokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenFactoryTransactor{contract: contract}, nil
}

// NewITokenFactoryFilterer creates a new log filterer instance of ITokenFactory, bound to a specific deployed contract.
func NewITokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFactoryFilterer, error) {
	contract, err := bindITokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFactoryFilterer{contract: contract}, nil
}

// bindITokenFactory binds a generic wrapper to an already deployed contract.
func bindITokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenFactory *ITokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenFactory.Contract.ITokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenFactory *ITokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenFactory.Contract.ITokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenFactory *ITokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenFactory.Contract.ITokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenFactory *ITokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenFactory *ITokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenFactory *ITokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateDepositToken is a paid mutator transaction binding the contract method 0x114a899c.
//
// Solidity: function CreateDepositToken(address ) returns(address)
func (_ITokenFactory *ITokenFactoryTransactor) CreateDepositToken(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ITokenFactory.contract.Transact(opts, "CreateDepositToken", arg0)
}

// CreateDepositToken is a paid mutator transaction binding the contract method 0x114a899c.
//
// Solidity: function CreateDepositToken(address ) returns(address)
func (_ITokenFactory *ITokenFactorySession) CreateDepositToken(arg0 common.Address) (*types.Transaction, error) {
	return _ITokenFactory.Contract.CreateDepositToken(&_ITokenFactory.TransactOpts, arg0)
}

// CreateDepositToken is a paid mutator transaction binding the contract method 0x114a899c.
//
// Solidity: function CreateDepositToken(address ) returns(address)
func (_ITokenFactory *ITokenFactoryTransactorSession) CreateDepositToken(arg0 common.Address) (*types.Transaction, error) {
	return _ITokenFactory.Contract.CreateDepositToken(&_ITokenFactory.TransactOpts, arg0)
}

// ITokenMinterMetaData contains all meta data concerning the ITokenMinter contract.
var ITokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"40c10f19": "mint(address,uint256)",
	},
}

// ITokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMinterMetaData.ABI instead.
var ITokenMinterABI = ITokenMinterMetaData.ABI

// Deprecated: Use ITokenMinterMetaData.Sigs instead.
// ITokenMinterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMinterFuncSigs = ITokenMinterMetaData.Sigs

// ITokenMinter is an auto generated Go binding around an Ethereum contract.
type ITokenMinter struct {
	ITokenMinterCaller     // Read-only binding to the contract
	ITokenMinterTransactor // Write-only binding to the contract
	ITokenMinterFilterer   // Log filterer for contract events
}

// ITokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMinterSession struct {
	Contract     *ITokenMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMinterCallerSession struct {
	Contract *ITokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMinterTransactorSession struct {
	Contract     *ITokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMinterRaw struct {
	Contract *ITokenMinter // Generic contract binding to access the raw methods on
}

// ITokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMinterCallerRaw struct {
	Contract *ITokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMinterTransactorRaw struct {
	Contract *ITokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMinter creates a new instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinter(address common.Address, backend bind.ContractBackend) (*ITokenMinter, error) {
	contract, err := bindITokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMinter{ITokenMinterCaller: ITokenMinterCaller{contract: contract}, ITokenMinterTransactor: ITokenMinterTransactor{contract: contract}, ITokenMinterFilterer: ITokenMinterFilterer{contract: contract}}, nil
}

// NewITokenMinterCaller creates a new read-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterCaller(address common.Address, caller bind.ContractCaller) (*ITokenMinterCaller, error) {
	contract, err := bindITokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterCaller{contract: contract}, nil
}

// NewITokenMinterTransactor creates a new write-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMinterTransactor, error) {
	contract, err := bindITokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterTransactor{contract: contract}, nil
}

// NewITokenMinterFilterer creates a new log filterer instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMinterFilterer, error) {
	contract, err := bindITokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterFilterer{contract: contract}, nil
}

// bindITokenMinter binds a generic wrapper to an already deployed contract.
func bindITokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.ITokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterTransactor) Burn(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "burn", arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterTransactor) Mint(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "mint", arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, arg0, arg1)
}

// IVestedEscrowMetaData contains all meta data concerning the IVestedEscrow contract.
var IVestedEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipient\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"fund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b1e56f6b": "fund(address[],uint256[])",
	},
}

// IVestedEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use IVestedEscrowMetaData.ABI instead.
var IVestedEscrowABI = IVestedEscrowMetaData.ABI

// Deprecated: Use IVestedEscrowMetaData.Sigs instead.
// IVestedEscrowFuncSigs maps the 4-byte function signature to its string representation.
var IVestedEscrowFuncSigs = IVestedEscrowMetaData.Sigs

// IVestedEscrow is an auto generated Go binding around an Ethereum contract.
type IVestedEscrow struct {
	IVestedEscrowCaller     // Read-only binding to the contract
	IVestedEscrowTransactor // Write-only binding to the contract
	IVestedEscrowFilterer   // Log filterer for contract events
}

// IVestedEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVestedEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVestedEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVestedEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVestedEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVestedEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVestedEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVestedEscrowSession struct {
	Contract     *IVestedEscrow    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVestedEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVestedEscrowCallerSession struct {
	Contract *IVestedEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IVestedEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVestedEscrowTransactorSession struct {
	Contract     *IVestedEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IVestedEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVestedEscrowRaw struct {
	Contract *IVestedEscrow // Generic contract binding to access the raw methods on
}

// IVestedEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVestedEscrowCallerRaw struct {
	Contract *IVestedEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// IVestedEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVestedEscrowTransactorRaw struct {
	Contract *IVestedEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVestedEscrow creates a new instance of IVestedEscrow, bound to a specific deployed contract.
func NewIVestedEscrow(address common.Address, backend bind.ContractBackend) (*IVestedEscrow, error) {
	contract, err := bindIVestedEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVestedEscrow{IVestedEscrowCaller: IVestedEscrowCaller{contract: contract}, IVestedEscrowTransactor: IVestedEscrowTransactor{contract: contract}, IVestedEscrowFilterer: IVestedEscrowFilterer{contract: contract}}, nil
}

// NewIVestedEscrowCaller creates a new read-only instance of IVestedEscrow, bound to a specific deployed contract.
func NewIVestedEscrowCaller(address common.Address, caller bind.ContractCaller) (*IVestedEscrowCaller, error) {
	contract, err := bindIVestedEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVestedEscrowCaller{contract: contract}, nil
}

// NewIVestedEscrowTransactor creates a new write-only instance of IVestedEscrow, bound to a specific deployed contract.
func NewIVestedEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*IVestedEscrowTransactor, error) {
	contract, err := bindIVestedEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVestedEscrowTransactor{contract: contract}, nil
}

// NewIVestedEscrowFilterer creates a new log filterer instance of IVestedEscrow, bound to a specific deployed contract.
func NewIVestedEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*IVestedEscrowFilterer, error) {
	contract, err := bindIVestedEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVestedEscrowFilterer{contract: contract}, nil
}

// bindIVestedEscrow binds a generic wrapper to an already deployed contract.
func bindIVestedEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVestedEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVestedEscrow *IVestedEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVestedEscrow.Contract.IVestedEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVestedEscrow *IVestedEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.IVestedEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVestedEscrow *IVestedEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.IVestedEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVestedEscrow *IVestedEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVestedEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVestedEscrow *IVestedEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVestedEscrow *IVestedEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.contract.Transact(opts, method, params...)
}

// Fund is a paid mutator transaction binding the contract method 0xb1e56f6b.
//
// Solidity: function fund(address[] _recipient, uint256[] _amount) returns(bool)
func (_IVestedEscrow *IVestedEscrowTransactor) Fund(opts *bind.TransactOpts, _recipient []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _IVestedEscrow.contract.Transact(opts, "fund", _recipient, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xb1e56f6b.
//
// Solidity: function fund(address[] _recipient, uint256[] _amount) returns(bool)
func (_IVestedEscrow *IVestedEscrowSession) Fund(_recipient []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.Fund(&_IVestedEscrow.TransactOpts, _recipient, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xb1e56f6b.
//
// Solidity: function fund(address[] _recipient, uint256[] _amount) returns(bool)
func (_IVestedEscrow *IVestedEscrowTransactorSession) Fund(_recipient []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _IVestedEscrow.Contract.Fund(&_IVestedEscrow.TransactOpts, _recipient, _amount)
}

// IVotingMetaData contains all meta data concerning the IVoting contract.
var IVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote_for_gauge_weights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5a55c1f0": "getVote(uint256)",
		"df133bca": "vote(uint256,bool,bool)",
		"d7136328": "vote_for_gauge_weights(address,uint256)",
	},
}

// IVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use IVotingMetaData.ABI instead.
var IVotingABI = IVotingMetaData.ABI

// Deprecated: Use IVotingMetaData.Sigs instead.
// IVotingFuncSigs maps the 4-byte function signature to its string representation.
var IVotingFuncSigs = IVotingMetaData.Sigs

// IVoting is an auto generated Go binding around an Ethereum contract.
type IVoting struct {
	IVotingCaller     // Read-only binding to the contract
	IVotingTransactor // Write-only binding to the contract
	IVotingFilterer   // Log filterer for contract events
}

// IVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVotingSession struct {
	Contract     *IVoting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVotingCallerSession struct {
	Contract *IVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVotingTransactorSession struct {
	Contract     *IVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVotingRaw struct {
	Contract *IVoting // Generic contract binding to access the raw methods on
}

// IVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVotingCallerRaw struct {
	Contract *IVotingCaller // Generic read-only contract binding to access the raw methods on
}

// IVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVotingTransactorRaw struct {
	Contract *IVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVoting creates a new instance of IVoting, bound to a specific deployed contract.
func NewIVoting(address common.Address, backend bind.ContractBackend) (*IVoting, error) {
	contract, err := bindIVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVoting{IVotingCaller: IVotingCaller{contract: contract}, IVotingTransactor: IVotingTransactor{contract: contract}, IVotingFilterer: IVotingFilterer{contract: contract}}, nil
}

// NewIVotingCaller creates a new read-only instance of IVoting, bound to a specific deployed contract.
func NewIVotingCaller(address common.Address, caller bind.ContractCaller) (*IVotingCaller, error) {
	contract, err := bindIVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVotingCaller{contract: contract}, nil
}

// NewIVotingTransactor creates a new write-only instance of IVoting, bound to a specific deployed contract.
func NewIVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*IVotingTransactor, error) {
	contract, err := bindIVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVotingTransactor{contract: contract}, nil
}

// NewIVotingFilterer creates a new log filterer instance of IVoting, bound to a specific deployed contract.
func NewIVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*IVotingFilterer, error) {
	contract, err := bindIVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVotingFilterer{contract: contract}, nil
}

// bindIVoting binds a generic wrapper to an already deployed contract.
func bindIVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoting *IVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoting.Contract.IVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoting *IVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoting.Contract.IVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoting *IVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoting.Contract.IVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoting *IVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoting *IVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoting *IVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoting.Contract.contract.Transact(opts, method, params...)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 ) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_IVoting *IVotingCaller) GetVote(opts *bind.CallOpts, arg0 *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	var out []interface{}
	err := _IVoting.contract.Call(opts, &out, "getVote", arg0)

	if err != nil {
		return *new(bool), *new(bool), *new(uint64), *new(uint64), *new(uint64), *new(uint64), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(uint64)).(*uint64)
	out5 := *abi.ConvertType(out[5], new(uint64)).(*uint64)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 ) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_IVoting *IVotingSession) GetVote(arg0 *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	return _IVoting.Contract.GetVote(&_IVoting.CallOpts, arg0)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 ) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_IVoting *IVotingCallerSession) GetVote(arg0 *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	return _IVoting.Contract.GetVote(&_IVoting.CallOpts, arg0)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 , bool , bool ) returns()
func (_IVoting *IVotingTransactor) Vote(opts *bind.TransactOpts, arg0 *big.Int, arg1 bool, arg2 bool) (*types.Transaction, error) {
	return _IVoting.contract.Transact(opts, "vote", arg0, arg1, arg2)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 , bool , bool ) returns()
func (_IVoting *IVotingSession) Vote(arg0 *big.Int, arg1 bool, arg2 bool) (*types.Transaction, error) {
	return _IVoting.Contract.Vote(&_IVoting.TransactOpts, arg0, arg1, arg2)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 , bool , bool ) returns()
func (_IVoting *IVotingTransactorSession) Vote(arg0 *big.Int, arg1 bool, arg2 bool) (*types.Transaction, error) {
	return _IVoting.Contract.Vote(&_IVoting.TransactOpts, arg0, arg1, arg2)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address , uint256 ) returns()
func (_IVoting *IVotingTransactor) VoteForGaugeWeights(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IVoting.contract.Transact(opts, "vote_for_gauge_weights", arg0, arg1)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address , uint256 ) returns()
func (_IVoting *IVotingSession) VoteForGaugeWeights(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IVoting.Contract.VoteForGaugeWeights(&_IVoting.TransactOpts, arg0, arg1)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address , uint256 ) returns()
func (_IVoting *IVotingTransactorSession) VoteForGaugeWeights(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IVoting.Contract.VoteForGaugeWeights(&_IVoting.TransactOpts, arg0, arg1)
}

// IWalletCheckerMetaData contains all meta data concerning the IWalletChecker contract.
var IWalletCheckerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c23697a8": "check(address)",
	},
}

// IWalletCheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use IWalletCheckerMetaData.ABI instead.
var IWalletCheckerABI = IWalletCheckerMetaData.ABI

// Deprecated: Use IWalletCheckerMetaData.Sigs instead.
// IWalletCheckerFuncSigs maps the 4-byte function signature to its string representation.
var IWalletCheckerFuncSigs = IWalletCheckerMetaData.Sigs

// IWalletChecker is an auto generated Go binding around an Ethereum contract.
type IWalletChecker struct {
	IWalletCheckerCaller     // Read-only binding to the contract
	IWalletCheckerTransactor // Write-only binding to the contract
	IWalletCheckerFilterer   // Log filterer for contract events
}

// IWalletCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWalletCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWalletCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWalletCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWalletCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWalletCheckerSession struct {
	Contract     *IWalletChecker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWalletCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWalletCheckerCallerSession struct {
	Contract *IWalletCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWalletCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWalletCheckerTransactorSession struct {
	Contract     *IWalletCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWalletCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWalletCheckerRaw struct {
	Contract *IWalletChecker // Generic contract binding to access the raw methods on
}

// IWalletCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWalletCheckerCallerRaw struct {
	Contract *IWalletCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// IWalletCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWalletCheckerTransactorRaw struct {
	Contract *IWalletCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWalletChecker creates a new instance of IWalletChecker, bound to a specific deployed contract.
func NewIWalletChecker(address common.Address, backend bind.ContractBackend) (*IWalletChecker, error) {
	contract, err := bindIWalletChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWalletChecker{IWalletCheckerCaller: IWalletCheckerCaller{contract: contract}, IWalletCheckerTransactor: IWalletCheckerTransactor{contract: contract}, IWalletCheckerFilterer: IWalletCheckerFilterer{contract: contract}}, nil
}

// NewIWalletCheckerCaller creates a new read-only instance of IWalletChecker, bound to a specific deployed contract.
func NewIWalletCheckerCaller(address common.Address, caller bind.ContractCaller) (*IWalletCheckerCaller, error) {
	contract, err := bindIWalletChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWalletCheckerCaller{contract: contract}, nil
}

// NewIWalletCheckerTransactor creates a new write-only instance of IWalletChecker, bound to a specific deployed contract.
func NewIWalletCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*IWalletCheckerTransactor, error) {
	contract, err := bindIWalletChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWalletCheckerTransactor{contract: contract}, nil
}

// NewIWalletCheckerFilterer creates a new log filterer instance of IWalletChecker, bound to a specific deployed contract.
func NewIWalletCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*IWalletCheckerFilterer, error) {
	contract, err := bindIWalletChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWalletCheckerFilterer{contract: contract}, nil
}

// bindIWalletChecker binds a generic wrapper to an already deployed contract.
func bindIWalletChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWalletCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWalletChecker *IWalletCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWalletChecker.Contract.IWalletCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWalletChecker *IWalletCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWalletChecker.Contract.IWalletCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWalletChecker *IWalletCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWalletChecker.Contract.IWalletCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWalletChecker *IWalletCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWalletChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWalletChecker *IWalletCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWalletChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWalletChecker *IWalletCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWalletChecker.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address ) view returns(bool)
func (_IWalletChecker *IWalletCheckerCaller) Check(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IWalletChecker.contract.Call(opts, &out, "check", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address ) view returns(bool)
func (_IWalletChecker *IWalletCheckerSession) Check(arg0 common.Address) (bool, error) {
	return _IWalletChecker.Contract.Check(&_IWalletChecker.CallOpts, arg0)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address ) view returns(bool)
func (_IWalletChecker *IWalletCheckerCallerSession) Check(arg0 common.Address) (bool, error) {
	return _IWalletChecker.Contract.Check(&_IWalletChecker.CallOpts, arg0)
}

// MathUtilMetaData contains all meta data concerning the MathUtil contract.
var MathUtilMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e6c7bc52ef8486439d0f59a7c7fef4323b4240c5172f722a821ebf1e935b65da64736f6c634300060c0033",
}

// MathUtilABI is the input ABI used to generate the binding from.
// Deprecated: Use MathUtilMetaData.ABI instead.
var MathUtilABI = MathUtilMetaData.ABI

// MathUtilBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathUtilMetaData.Bin instead.
var MathUtilBin = MathUtilMetaData.Bin

// DeployMathUtil deploys a new Ethereum contract, binding an instance of MathUtil to it.
func DeployMathUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathUtil, error) {
	parsed, err := MathUtilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathUtil{MathUtilCaller: MathUtilCaller{contract: contract}, MathUtilTransactor: MathUtilTransactor{contract: contract}, MathUtilFilterer: MathUtilFilterer{contract: contract}}, nil
}

// MathUtil is an auto generated Go binding around an Ethereum contract.
type MathUtil struct {
	MathUtilCaller     // Read-only binding to the contract
	MathUtilTransactor // Write-only binding to the contract
	MathUtilFilterer   // Log filterer for contract events
}

// MathUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathUtilSession struct {
	Contract     *MathUtil         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathUtilCallerSession struct {
	Contract *MathUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MathUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathUtilTransactorSession struct {
	Contract     *MathUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MathUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathUtilRaw struct {
	Contract *MathUtil // Generic contract binding to access the raw methods on
}

// MathUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathUtilCallerRaw struct {
	Contract *MathUtilCaller // Generic read-only contract binding to access the raw methods on
}

// MathUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathUtilTransactorRaw struct {
	Contract *MathUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMathUtil creates a new instance of MathUtil, bound to a specific deployed contract.
func NewMathUtil(address common.Address, backend bind.ContractBackend) (*MathUtil, error) {
	contract, err := bindMathUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MathUtil{MathUtilCaller: MathUtilCaller{contract: contract}, MathUtilTransactor: MathUtilTransactor{contract: contract}, MathUtilFilterer: MathUtilFilterer{contract: contract}}, nil
}

// NewMathUtilCaller creates a new read-only instance of MathUtil, bound to a specific deployed contract.
func NewMathUtilCaller(address common.Address, caller bind.ContractCaller) (*MathUtilCaller, error) {
	contract, err := bindMathUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathUtilCaller{contract: contract}, nil
}

// NewMathUtilTransactor creates a new write-only instance of MathUtil, bound to a specific deployed contract.
func NewMathUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*MathUtilTransactor, error) {
	contract, err := bindMathUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathUtilTransactor{contract: contract}, nil
}

// NewMathUtilFilterer creates a new log filterer instance of MathUtil, bound to a specific deployed contract.
func NewMathUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*MathUtilFilterer, error) {
	contract, err := bindMathUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathUtilFilterer{contract: contract}, nil
}

// bindMathUtil binds a generic wrapper to an already deployed contract.
func bindMathUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathUtil *MathUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathUtil.Contract.MathUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathUtil *MathUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathUtil.Contract.MathUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathUtil *MathUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathUtil.Contract.MathUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathUtil *MathUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathUtil *MathUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathUtil *MathUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathUtil.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209cbfbff0443881ecf0593f31afc8180f336de30d7fc67049280cbdb0059bb4d464736f6c634300060c0033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122066d6034da2af8ebe34c24d3b4eb315119f9fb9aa2ed15e7cd2acd951ce27783a64736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
