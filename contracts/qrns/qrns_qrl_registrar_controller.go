// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package qrns

import (
	"errors"
	"math/big"
	"strings"

	qrl "github.com/theQRL/go-zond"
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
	_ = qrl.NotFound
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

// QRNSQRLRegistrarControllerMetaData contains all meta data concerning the QRNSQRLRegistrarController contract.
var QRNSQRLRegistrarControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractIPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"contractReverseRegistrar\",\"name\":\"_reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"_nameWrapper\",\"type\":\"address\"},{\"internalType\":\"contractQRNS\",\"name\":\"_qrns\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooNew\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResolverRequiredWhenDataSupplied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"UnexpiredCommitmentExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameWrapper\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"internalType\":\"contractReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QRNSQRLRegistrarControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSQRLRegistrarControllerMetaData.ABI instead.
var QRNSQRLRegistrarControllerABI = QRNSQRLRegistrarControllerMetaData.ABI

// QRNSQRLRegistrarController is an auto generated Go binding around a QRL contract.
type QRNSQRLRegistrarController struct {
	QRNSQRLRegistrarControllerCaller     // Read-only binding to the contract
	QRNSQRLRegistrarControllerTransactor // Write-only binding to the contract
	QRNSQRLRegistrarControllerFilterer   // Log filterer for contract events
}

// QRNSQRLRegistrarControllerCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSQRLRegistrarControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSQRLRegistrarControllerTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSQRLRegistrarControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSQRLRegistrarControllerFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSQRLRegistrarControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSQRLRegistrarControllerSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSQRLRegistrarControllerSession struct {
	Contract     *QRNSQRLRegistrarController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QRNSQRLRegistrarControllerCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSQRLRegistrarControllerCallerSession struct {
	Contract *QRNSQRLRegistrarControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// QRNSQRLRegistrarControllerTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSQRLRegistrarControllerTransactorSession struct {
	Contract     *QRNSQRLRegistrarControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// QRNSQRLRegistrarControllerRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSQRLRegistrarControllerRaw struct {
	Contract *QRNSQRLRegistrarController // Generic contract binding to access the raw methods on
}

// QRNSQRLRegistrarControllerCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSQRLRegistrarControllerCallerRaw struct {
	Contract *QRNSQRLRegistrarControllerCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSQRLRegistrarControllerTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSQRLRegistrarControllerTransactorRaw struct {
	Contract *QRNSQRLRegistrarControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSQRLRegistrarController creates a new instance of QRNSQRLRegistrarController, bound to a specific deployed contract.
func NewQRNSQRLRegistrarController(address common.Address, backend bind.ContractBackend) (*QRNSQRLRegistrarController, error) {
	contract, err := bindQRNSQRLRegistrarController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarController{QRNSQRLRegistrarControllerCaller: QRNSQRLRegistrarControllerCaller{contract: contract}, QRNSQRLRegistrarControllerTransactor: QRNSQRLRegistrarControllerTransactor{contract: contract}, QRNSQRLRegistrarControllerFilterer: QRNSQRLRegistrarControllerFilterer{contract: contract}}, nil
}

// NewQRNSQRLRegistrarControllerCaller creates a new read-only instance of QRNSQRLRegistrarController, bound to a specific deployed contract.
func NewQRNSQRLRegistrarControllerCaller(address common.Address, caller bind.ContractCaller) (*QRNSQRLRegistrarControllerCaller, error) {
	contract, err := bindQRNSQRLRegistrarController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerCaller{contract: contract}, nil
}

// NewQRNSQRLRegistrarControllerTransactor creates a new write-only instance of QRNSQRLRegistrarController, bound to a specific deployed contract.
func NewQRNSQRLRegistrarControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSQRLRegistrarControllerTransactor, error) {
	contract, err := bindQRNSQRLRegistrarController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerTransactor{contract: contract}, nil
}

// NewQRNSQRLRegistrarControllerFilterer creates a new log filterer instance of QRNSQRLRegistrarController, bound to a specific deployed contract.
func NewQRNSQRLRegistrarControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSQRLRegistrarControllerFilterer, error) {
	contract, err := bindQRNSQRLRegistrarController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerFilterer{contract: contract}, nil
}

// bindQRNSQRLRegistrarController binds a generic wrapper to an already deployed contract.
func bindQRNSQRLRegistrarController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSQRLRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSQRLRegistrarController.Contract.QRNSQRLRegistrarControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.QRNSQRLRegistrarControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.QRNSQRLRegistrarControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSQRLRegistrarController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MINREGISTRATIONDURATION(&_QRNSQRLRegistrarController.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MINREGISTRATIONDURATION(&_QRNSQRLRegistrarController.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Available(name string) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.Available(&_QRNSQRLRegistrarController.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) Available(name string) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.Available(&_QRNSQRLRegistrarController.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.Commitments(&_QRNSQRLRegistrarController.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.Commitments(&_QRNSQRLRegistrarController.CallOpts, arg0)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "makeCommitment", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _QRNSQRLRegistrarController.Contract.MakeCommitment(&_QRNSQRLRegistrarController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _QRNSQRLRegistrarController.Contract.MakeCommitment(&_QRNSQRLRegistrarController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) MaxCommitmentAge() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MaxCommitmentAge(&_QRNSQRLRegistrarController.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MaxCommitmentAge(&_QRNSQRLRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) MinCommitmentAge() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MinCommitmentAge(&_QRNSQRLRegistrarController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) MinCommitmentAge() (*big.Int, error) {
	return _QRNSQRLRegistrarController.Contract.MinCommitmentAge(&_QRNSQRLRegistrarController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) NameWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "nameWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) NameWrapper() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.NameWrapper(&_QRNSQRLRegistrarController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) NameWrapper() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.NameWrapper(&_QRNSQRLRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Owner() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.Owner(&_QRNSQRLRegistrarController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) Owner() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.Owner(&_QRNSQRLRegistrarController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "prices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Prices() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.Prices(&_QRNSQRLRegistrarController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) Prices() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.Prices(&_QRNSQRLRegistrarController.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(IPriceOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _QRNSQRLRegistrarController.Contract.RentPrice(&_QRNSQRLRegistrarController.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _QRNSQRLRegistrarController.Contract.RentPrice(&_QRNSQRLRegistrarController.CallOpts, name, duration)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "reverseRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) ReverseRegistrar() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.ReverseRegistrar(&_QRNSQRLRegistrarController.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) ReverseRegistrar() (common.Address, error) {
	return _QRNSQRLRegistrarController.Contract.ReverseRegistrar(&_QRNSQRLRegistrarController.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.SupportsInterface(&_QRNSQRLRegistrarController.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.SupportsInterface(&_QRNSQRLRegistrarController.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _QRNSQRLRegistrarController.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Valid(name string) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.Valid(&_QRNSQRLRegistrarController.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerCallerSession) Valid(name string) (bool, error) {
	return _QRNSQRLRegistrarController.Contract.Valid(&_QRNSQRLRegistrarController.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Commit(&_QRNSQRLRegistrarController.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Commit(&_QRNSQRLRegistrarController.TransactOpts, commitment)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.RecoverFunds(&_QRNSQRLRegistrarController.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.RecoverFunds(&_QRNSQRLRegistrarController.TransactOpts, _token, _to, _amount)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "register", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Register(&_QRNSQRLRegistrarController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Register(&_QRNSQRLRegistrarController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Renew(&_QRNSQRLRegistrarController.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Renew(&_QRNSQRLRegistrarController.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.RenounceOwnership(&_QRNSQRLRegistrarController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.RenounceOwnership(&_QRNSQRLRegistrarController.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.TransferOwnership(&_QRNSQRLRegistrarController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.TransferOwnership(&_QRNSQRLRegistrarController.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerSession) Withdraw() (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Withdraw(&_QRNSQRLRegistrarController.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _QRNSQRLRegistrarController.Contract.Withdraw(&_QRNSQRLRegistrarController.TransactOpts)
}

// QRNSQRLRegistrarControllerNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerNameRegisteredIterator struct {
	Event *QRNSQRLRegistrarControllerNameRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log   // Log channel receiving the found contract events
	sub  qrl.Subscription // Subscription for errors, completion and termination
	done bool             // Whether the subscription completed delivering logs
	fail error            // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QRNSQRLRegistrarControllerNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSQRLRegistrarControllerNameRegistered)
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
		it.Event = new(QRNSQRLRegistrarControllerNameRegistered)
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
func (it *QRNSQRLRegistrarControllerNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSQRLRegistrarControllerNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSQRLRegistrarControllerNameRegistered represents a NameRegistered event raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerNameRegistered struct {
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
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*QRNSQRLRegistrarControllerNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerNameRegisteredIterator{contract: _QRNSQRLRegistrarController.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *QRNSQRLRegistrarControllerNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSQRLRegistrarControllerNameRegistered)
				if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) ParseNameRegistered(log types.Log) (*QRNSQRLRegistrarControllerNameRegistered, error) {
	event := new(QRNSQRLRegistrarControllerNameRegistered)
	if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSQRLRegistrarControllerNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerNameRenewedIterator struct {
	Event *QRNSQRLRegistrarControllerNameRenewed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log   // Log channel receiving the found contract events
	sub  qrl.Subscription // Subscription for errors, completion and termination
	done bool             // Whether the subscription completed delivering logs
	fail error            // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QRNSQRLRegistrarControllerNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSQRLRegistrarControllerNameRenewed)
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
		it.Event = new(QRNSQRLRegistrarControllerNameRenewed)
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
func (it *QRNSQRLRegistrarControllerNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSQRLRegistrarControllerNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSQRLRegistrarControllerNameRenewed represents a NameRenewed event raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*QRNSQRLRegistrarControllerNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerNameRenewedIterator{contract: _QRNSQRLRegistrarController.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *QRNSQRLRegistrarControllerNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSQRLRegistrarControllerNameRenewed)
				if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) ParseNameRenewed(log types.Log) (*QRNSQRLRegistrarControllerNameRenewed, error) {
	event := new(QRNSQRLRegistrarControllerNameRenewed)
	if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSQRLRegistrarControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerOwnershipTransferredIterator struct {
	Event *QRNSQRLRegistrarControllerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log   // Log channel receiving the found contract events
	sub  qrl.Subscription // Subscription for errors, completion and termination
	done bool             // Whether the subscription completed delivering logs
	fail error            // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QRNSQRLRegistrarControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSQRLRegistrarControllerOwnershipTransferred)
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
		it.Event = new(QRNSQRLRegistrarControllerOwnershipTransferred)
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
func (it *QRNSQRLRegistrarControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSQRLRegistrarControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSQRLRegistrarControllerOwnershipTransferred represents a OwnershipTransferred event raised by the QRNSQRLRegistrarController contract.
type QRNSQRLRegistrarControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QRNSQRLRegistrarControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSQRLRegistrarControllerOwnershipTransferredIterator{contract: _QRNSQRLRegistrarController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QRNSQRLRegistrarControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSQRLRegistrarController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSQRLRegistrarControllerOwnershipTransferred)
				if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QRNSQRLRegistrarController *QRNSQRLRegistrarControllerFilterer) ParseOwnershipTransferred(log types.Log) (*QRNSQRLRegistrarControllerOwnershipTransferred, error) {
	event := new(QRNSQRLRegistrarControllerOwnershipTransferred)
	if err := _QRNSQRLRegistrarController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
