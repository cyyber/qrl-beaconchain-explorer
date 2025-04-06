// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zns

import (
	"errors"
	"math/big"
	"strings"

	zond "github.com/theQRL/go-zond"
	"github.com/theQRL/go-zond/accounts/abi"
	"github.com/theQRL/go-zond/accounts/abi/bind"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/core/types"
	"github.com/theQRL/go-zond/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = zond.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IPriceOraclePrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceOraclePrice struct {
	Base    *big.Int
	Premium *big.Int
}

// ZNSZONDRegistrarControllerMetaData contains all meta data concerning the ZNSZONDRegistrarController contract.
var ZNSZONDRegistrarControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractIPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"contractReverseRegistrar\",\"name\":\"_reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"_nameWrapper\",\"type\":\"address\"},{\"internalType\":\"contractZNS\",\"name\":\"_zns\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooNew\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResolverRequiredWhenDataSupplied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"UnexpiredCommitmentExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameWrapper\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"internalType\":\"contractReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZNSZONDRegistrarControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ZNSZONDRegistrarControllerMetaData.ABI instead.
var ZNSZONDRegistrarControllerABI = ZNSZONDRegistrarControllerMetaData.ABI

// ZNSZONDRegistrarController is an auto generated Go binding around a Zond contract.
type ZNSZONDRegistrarController struct {
	ZNSZONDRegistrarControllerCaller     // Read-only binding to the contract
	ZNSZONDRegistrarControllerTransactor // Write-only binding to the contract
	ZNSZONDRegistrarControllerFilterer   // Log filterer for contract events
}

// ZNSZONDRegistrarControllerCaller is an auto generated read-only Go binding around a Zond contract.
type ZNSZONDRegistrarControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSZONDRegistrarControllerTransactor is an auto generated write-only Go binding around a Zond contract.
type ZNSZONDRegistrarControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSZONDRegistrarControllerFilterer is an auto generated log filtering Go binding around a Zond contract events.
type ZNSZONDRegistrarControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZNSZONDRegistrarControllerSession is an auto generated Go binding around a Zond contract,
// with pre-set call and transact options.
type ZNSZONDRegistrarControllerSession struct {
	Contract     *ZNSZONDRegistrarController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ZNSZONDRegistrarControllerCallerSession is an auto generated read-only Go binding around a Zond contract,
// with pre-set call options.
type ZNSZONDRegistrarControllerCallerSession struct {
	Contract *ZNSZONDRegistrarControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ZNSZONDRegistrarControllerTransactorSession is an auto generated write-only Go binding around a Zond contract,
// with pre-set transact options.
type ZNSZONDRegistrarControllerTransactorSession struct {
	Contract     *ZNSZONDRegistrarControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ZNSZONDRegistrarControllerRaw is an auto generated low-level Go binding around a Zond contract.
type ZNSZONDRegistrarControllerRaw struct {
	Contract *ZNSZONDRegistrarController // Generic contract binding to access the raw methods on
}

// ZNSZONDRegistrarControllerCallerRaw is an auto generated low-level read-only Go binding around a Zond contract.
type ZNSZONDRegistrarControllerCallerRaw struct {
	Contract *ZNSZONDRegistrarControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ZNSZONDRegistrarControllerTransactorRaw is an auto generated low-level write-only Go binding around a Zond contract.
type ZNSZONDRegistrarControllerTransactorRaw struct {
	Contract *ZNSZONDRegistrarControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZNSZONDRegistrarController creates a new instance of ZNSZONDRegistrarController, bound to a specific deployed contract.
func NewZNSZONDRegistrarController(address common.Address, backend bind.ContractBackend) (*ZNSZONDRegistrarController, error) {
	contract, err := bindZNSZONDRegistrarController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarController{ZNSZONDRegistrarControllerCaller: ZNSZONDRegistrarControllerCaller{contract: contract}, ZNSZONDRegistrarControllerTransactor: ZNSZONDRegistrarControllerTransactor{contract: contract}, ZNSZONDRegistrarControllerFilterer: ZNSZONDRegistrarControllerFilterer{contract: contract}}, nil
}

// NewZNSZONDRegistrarControllerCaller creates a new read-only instance of ZNSZONDRegistrarController, bound to a specific deployed contract.
func NewZNSZONDRegistrarControllerCaller(address common.Address, caller bind.ContractCaller) (*ZNSZONDRegistrarControllerCaller, error) {
	contract, err := bindZNSZONDRegistrarController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerCaller{contract: contract}, nil
}

// NewZNSZONDRegistrarControllerTransactor creates a new write-only instance of ZNSZONDRegistrarController, bound to a specific deployed contract.
func NewZNSZONDRegistrarControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZNSZONDRegistrarControllerTransactor, error) {
	contract, err := bindZNSZONDRegistrarController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerTransactor{contract: contract}, nil
}

// NewZNSZONDRegistrarControllerFilterer creates a new log filterer instance of ZNSZONDRegistrarController, bound to a specific deployed contract.
func NewZNSZONDRegistrarControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZNSZONDRegistrarControllerFilterer, error) {
	contract, err := bindZNSZONDRegistrarController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerFilterer{contract: contract}, nil
}

// bindZNSZONDRegistrarController binds a generic wrapper to an already deployed contract.
func bindZNSZONDRegistrarController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZNSZONDRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSZONDRegistrarController.Contract.ZNSZONDRegistrarControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.ZNSZONDRegistrarControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.ZNSZONDRegistrarControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZNSZONDRegistrarController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MINREGISTRATIONDURATION(&_ZNSZONDRegistrarController.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MINREGISTRATIONDURATION(&_ZNSZONDRegistrarController.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Available(name string) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.Available(&_ZNSZONDRegistrarController.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) Available(name string) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.Available(&_ZNSZONDRegistrarController.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.Commitments(&_ZNSZONDRegistrarController.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.Commitments(&_ZNSZONDRegistrarController.CallOpts, arg0)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "makeCommitment", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _ZNSZONDRegistrarController.Contract.MakeCommitment(&_ZNSZONDRegistrarController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _ZNSZONDRegistrarController.Contract.MakeCommitment(&_ZNSZONDRegistrarController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) MaxCommitmentAge() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MaxCommitmentAge(&_ZNSZONDRegistrarController.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MaxCommitmentAge(&_ZNSZONDRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) MinCommitmentAge() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MinCommitmentAge(&_ZNSZONDRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) MinCommitmentAge() (*big.Int, error) {
	return _ZNSZONDRegistrarController.Contract.MinCommitmentAge(&_ZNSZONDRegistrarController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) NameWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "nameWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) NameWrapper() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.NameWrapper(&_ZNSZONDRegistrarController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) NameWrapper() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.NameWrapper(&_ZNSZONDRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Owner() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.Owner(&_ZNSZONDRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) Owner() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.Owner(&_ZNSZONDRegistrarController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "prices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Prices() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.Prices(&_ZNSZONDRegistrarController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) Prices() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.Prices(&_ZNSZONDRegistrarController.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(IPriceOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _ZNSZONDRegistrarController.Contract.RentPrice(&_ZNSZONDRegistrarController.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _ZNSZONDRegistrarController.Contract.RentPrice(&_ZNSZONDRegistrarController.CallOpts, name, duration)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "reverseRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) ReverseRegistrar() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.ReverseRegistrar(&_ZNSZONDRegistrarController.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) ReverseRegistrar() (common.Address, error) {
	return _ZNSZONDRegistrarController.Contract.ReverseRegistrar(&_ZNSZONDRegistrarController.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.SupportsInterface(&_ZNSZONDRegistrarController.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.SupportsInterface(&_ZNSZONDRegistrarController.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZNSZONDRegistrarController.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Valid(name string) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.Valid(&_ZNSZONDRegistrarController.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerCallerSession) Valid(name string) (bool, error) {
	return _ZNSZONDRegistrarController.Contract.Valid(&_ZNSZONDRegistrarController.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Commit(&_ZNSZONDRegistrarController.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Commit(&_ZNSZONDRegistrarController.TransactOpts, commitment)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.RecoverFunds(&_ZNSZONDRegistrarController.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.RecoverFunds(&_ZNSZONDRegistrarController.TransactOpts, _token, _to, _amount)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "register", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Register(&_ZNSZONDRegistrarController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Register(&_ZNSZONDRegistrarController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Renew(&_ZNSZONDRegistrarController.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Renew(&_ZNSZONDRegistrarController.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.RenounceOwnership(&_ZNSZONDRegistrarController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.RenounceOwnership(&_ZNSZONDRegistrarController.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.TransferOwnership(&_ZNSZONDRegistrarController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.TransferOwnership(&_ZNSZONDRegistrarController.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerSession) Withdraw() (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Withdraw(&_ZNSZONDRegistrarController.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ZNSZONDRegistrarController.Contract.Withdraw(&_ZNSZONDRegistrarController.TransactOpts)
}

// ZNSZONDRegistrarControllerNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerNameRegisteredIterator struct {
	Event *ZNSZONDRegistrarControllerNameRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSZONDRegistrarControllerNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSZONDRegistrarControllerNameRegistered)
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
		it.Event = new(ZNSZONDRegistrarControllerNameRegistered)
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
func (it *ZNSZONDRegistrarControllerNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSZONDRegistrarControllerNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSZONDRegistrarControllerNameRegistered represents a NameRegistered event raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerNameRegistered struct {
	Name     string
	Label    [32]byte
	Owner    common.Address
	BaseCost *big.Int
	Premium  *big.Int
	Expires  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*ZNSZONDRegistrarControllerNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerNameRegisteredIterator{contract: _ZNSZONDRegistrarController.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *ZNSZONDRegistrarControllerNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSZONDRegistrarControllerNameRegistered)
				if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) ParseNameRegistered(log types.Log) (*ZNSZONDRegistrarControllerNameRegistered, error) {
	event := new(ZNSZONDRegistrarControllerNameRegistered)
	if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSZONDRegistrarControllerNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerNameRenewedIterator struct {
	Event *ZNSZONDRegistrarControllerNameRenewed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSZONDRegistrarControllerNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSZONDRegistrarControllerNameRenewed)
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
		it.Event = new(ZNSZONDRegistrarControllerNameRenewed)
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
func (it *ZNSZONDRegistrarControllerNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSZONDRegistrarControllerNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSZONDRegistrarControllerNameRenewed represents a NameRenewed event raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*ZNSZONDRegistrarControllerNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerNameRenewedIterator{contract: _ZNSZONDRegistrarController.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *ZNSZONDRegistrarControllerNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSZONDRegistrarControllerNameRenewed)
				if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) ParseNameRenewed(log types.Log) (*ZNSZONDRegistrarControllerNameRenewed, error) {
	event := new(ZNSZONDRegistrarControllerNameRenewed)
	if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZNSZONDRegistrarControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerOwnershipTransferredIterator struct {
	Event *ZNSZONDRegistrarControllerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  zond.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZNSZONDRegistrarControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZNSZONDRegistrarControllerOwnershipTransferred)
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
		it.Event = new(ZNSZONDRegistrarControllerOwnershipTransferred)
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
func (it *ZNSZONDRegistrarControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZNSZONDRegistrarControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZNSZONDRegistrarControllerOwnershipTransferred represents a OwnershipTransferred event raised by the ZNSZONDRegistrarController contract.
type ZNSZONDRegistrarControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZNSZONDRegistrarControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZNSZONDRegistrarControllerOwnershipTransferredIterator{contract: _ZNSZONDRegistrarController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZNSZONDRegistrarControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZNSZONDRegistrarController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZNSZONDRegistrarControllerOwnershipTransferred)
				if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZNSZONDRegistrarController *ZNSZONDRegistrarControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ZNSZONDRegistrarControllerOwnershipTransferred, error) {
	event := new(ZNSZONDRegistrarControllerOwnershipTransferred)
	if err := _ZNSZONDRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
