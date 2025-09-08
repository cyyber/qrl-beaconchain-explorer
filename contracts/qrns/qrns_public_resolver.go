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

// QRNSPublicResolverMetaData contains all meta data concerning the QRNSPublicResolver contract.
var QRNSPublicResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"_qrns\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"wrapperAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedQRLController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedReverseRegistrar\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"record\",\"type\":\"bytes\"}],\"name\":\"DNSRecordChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"DNSRecordDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lastzonehash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"zonehash\",\"type\":\"bytes\"}],\"name\":\"DNSZonehashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVersion\",\"type\":\"uint64\"}],\"name\":\"VersionChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"interfaceImplementer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"isApprovedFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodehash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallWithNodeCheck\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"recordVersions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"setInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setZonehash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"zonehash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// QRNSPublicResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSPublicResolverMetaData.ABI instead.
var QRNSPublicResolverABI = QRNSPublicResolverMetaData.ABI

// QRNSPublicResolver is an auto generated Go binding around a QRL contract.
type QRNSPublicResolver struct {
	QRNSPublicResolverCaller     // Read-only binding to the contract
	QRNSPublicResolverTransactor // Write-only binding to the contract
	QRNSPublicResolverFilterer   // Log filterer for contract events
}

// QRNSPublicResolverCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSPublicResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSPublicResolverTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSPublicResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSPublicResolverFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSPublicResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSPublicResolverSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSPublicResolverSession struct {
	Contract     *QRNSPublicResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QRNSPublicResolverCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSPublicResolverCallerSession struct {
	Contract *QRNSPublicResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// QRNSPublicResolverTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSPublicResolverTransactorSession struct {
	Contract     *QRNSPublicResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// QRNSPublicResolverRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSPublicResolverRaw struct {
	Contract *QRNSPublicResolver // Generic contract binding to access the raw methods on
}

// QRNSPublicResolverCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSPublicResolverCallerRaw struct {
	Contract *QRNSPublicResolverCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSPublicResolverTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSPublicResolverTransactorRaw struct {
	Contract *QRNSPublicResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSPublicResolver creates a new instance of QRNSPublicResolver, bound to a specific deployed contract.
func NewQRNSPublicResolver(address common.Address, backend bind.ContractBackend) (*QRNSPublicResolver, error) {
	contract, err := bindQRNSPublicResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolver{QRNSPublicResolverCaller: QRNSPublicResolverCaller{contract: contract}, QRNSPublicResolverTransactor: QRNSPublicResolverTransactor{contract: contract}, QRNSPublicResolverFilterer: QRNSPublicResolverFilterer{contract: contract}}, nil
}

// NewQRNSPublicResolverCaller creates a new read-only instance of QRNSPublicResolver, bound to a specific deployed contract.
func NewQRNSPublicResolverCaller(address common.Address, caller bind.ContractCaller) (*QRNSPublicResolverCaller, error) {
	contract, err := bindQRNSPublicResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverCaller{contract: contract}, nil
}

// NewQRNSPublicResolverTransactor creates a new write-only instance of QRNSPublicResolver, bound to a specific deployed contract.
func NewQRNSPublicResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSPublicResolverTransactor, error) {
	contract, err := bindQRNSPublicResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverTransactor{contract: contract}, nil
}

// NewQRNSPublicResolverFilterer creates a new log filterer instance of QRNSPublicResolver, bound to a specific deployed contract.
func NewQRNSPublicResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSPublicResolverFilterer, error) {
	contract, err := bindQRNSPublicResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverFilterer{contract: contract}, nil
}

// bindQRNSPublicResolver binds a generic wrapper to an already deployed contract.
func bindQRNSPublicResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSPublicResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSPublicResolver *QRNSPublicResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSPublicResolver.Contract.QRNSPublicResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSPublicResolver *QRNSPublicResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.QRNSPublicResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSPublicResolver *QRNSPublicResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.QRNSPublicResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSPublicResolver *QRNSPublicResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSPublicResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSPublicResolver *QRNSPublicResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSPublicResolver *QRNSPublicResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Hyperion: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "ABI", node, contentTypes)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Hyperion: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_QRNSPublicResolver *QRNSPublicResolverSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _QRNSPublicResolver.Contract.ABI(&_QRNSPublicResolver.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Hyperion: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _QRNSPublicResolver.Contract.ABI(&_QRNSPublicResolver.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Hyperion: function addr(bytes32 node) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Hyperion: function addr(bytes32 node) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _QRNSPublicResolver.Contract.Addr(&_QRNSPublicResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Hyperion: function addr(bytes32 node) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _QRNSPublicResolver.Contract.Addr(&_QRNSPublicResolver.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Hyperion: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "addr0", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Hyperion: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Addr0(&_QRNSPublicResolver.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Hyperion: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Addr0(&_QRNSPublicResolver.CallOpts, node, coinType)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Hyperion: function contenthash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Hyperion: function contenthash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Contenthash(node [32]byte) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Contenthash(&_QRNSPublicResolver.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Hyperion: function contenthash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Contenthash(&_QRNSPublicResolver.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Hyperion: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "dnsRecord", node, name, resource)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Hyperion: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _QRNSPublicResolver.Contract.DnsRecord(&_QRNSPublicResolver.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Hyperion: function dnsRecord(bytes32 node, bytes32 name, uint16 resource) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _QRNSPublicResolver.Contract.DnsRecord(&_QRNSPublicResolver.CallOpts, node, name, resource)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Hyperion: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) HasDNSRecords(opts *bind.CallOpts, node [32]byte, name [32]byte) (bool, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "hasDNSRecords", node, name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Hyperion: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _QRNSPublicResolver.Contract.HasDNSRecords(&_QRNSPublicResolver.CallOpts, node, name)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Hyperion: function hasDNSRecords(bytes32 node, bytes32 name) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) HasDNSRecords(node [32]byte, name [32]byte) (bool, error) {
	return _QRNSPublicResolver.Contract.HasDNSRecords(&_QRNSPublicResolver.CallOpts, node, name)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Hyperion: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) InterfaceImplementer(opts *bind.CallOpts, node [32]byte, interfaceID [4]byte) (common.Address, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "interfaceImplementer", node, interfaceID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Hyperion: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _QRNSPublicResolver.Contract.InterfaceImplementer(&_QRNSPublicResolver.CallOpts, node, interfaceID)
}

// InterfaceImplementer is a free data retrieval call binding the contract method 0x124a319c.
//
// Hyperion: function interfaceImplementer(bytes32 node, bytes4 interfaceID) view returns(address)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) InterfaceImplementer(node [32]byte, interfaceID [4]byte) (common.Address, error) {
	return _QRNSPublicResolver.Contract.InterfaceImplementer(&_QRNSPublicResolver.CallOpts, node, interfaceID)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Hyperion: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) IsApprovedFor(opts *bind.CallOpts, owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "isApprovedFor", owner, node, delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Hyperion: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverSession) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _QRNSPublicResolver.Contract.IsApprovedFor(&_QRNSPublicResolver.CallOpts, owner, node, delegate)
}

// IsApprovedFor is a free data retrieval call binding the contract method 0xa9784b3e.
//
// Hyperion: function isApprovedFor(address owner, bytes32 node, address delegate) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) IsApprovedFor(owner common.Address, node [32]byte, delegate common.Address) (bool, error) {
	return _QRNSPublicResolver.Contract.IsApprovedFor(&_QRNSPublicResolver.CallOpts, owner, node, delegate)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Hyperion: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Hyperion: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _QRNSPublicResolver.Contract.IsApprovedForAll(&_QRNSPublicResolver.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Hyperion: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _QRNSPublicResolver.Contract.IsApprovedForAll(&_QRNSPublicResolver.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Hyperion: function name(bytes32 node) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Hyperion: function name(bytes32 node) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Name(node [32]byte) (string, error) {
	return _QRNSPublicResolver.Contract.Name(&_QRNSPublicResolver.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Hyperion: function name(bytes32 node) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Name(node [32]byte) (string, error) {
	return _QRNSPublicResolver.Contract.Name(&_QRNSPublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Hyperion: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Hyperion: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _QRNSPublicResolver.Contract.Pubkey(&_QRNSPublicResolver.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Hyperion: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _QRNSPublicResolver.Contract.Pubkey(&_QRNSPublicResolver.CallOpts, node)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Hyperion: function recordVersions(bytes32 ) view returns(uint64)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) RecordVersions(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "recordVersions", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Hyperion: function recordVersions(bytes32 ) view returns(uint64)
func (_QRNSPublicResolver *QRNSPublicResolverSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _QRNSPublicResolver.Contract.RecordVersions(&_QRNSPublicResolver.CallOpts, arg0)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Hyperion: function recordVersions(bytes32 ) view returns(uint64)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _QRNSPublicResolver.Contract.RecordVersions(&_QRNSPublicResolver.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSPublicResolver.Contract.SupportsInterface(&_QRNSPublicResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSPublicResolver.Contract.SupportsInterface(&_QRNSPublicResolver.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Hyperion: function text(bytes32 node, string key) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Hyperion: function text(bytes32 node, string key) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Text(node [32]byte, key string) (string, error) {
	return _QRNSPublicResolver.Contract.Text(&_QRNSPublicResolver.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Hyperion: function text(bytes32 node, string key) view returns(string)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Text(node [32]byte, key string) (string, error) {
	return _QRNSPublicResolver.Contract.Text(&_QRNSPublicResolver.CallOpts, node, key)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Hyperion: function zonehash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCaller) Zonehash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _QRNSPublicResolver.contract.Call(opts, &out, "zonehash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Hyperion: function zonehash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Zonehash(node [32]byte) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Zonehash(&_QRNSPublicResolver.CallOpts, node)
}

// Zonehash is a free data retrieval call binding the contract method 0x5c98042b.
//
// Hyperion: function zonehash(bytes32 node) view returns(bytes)
func (_QRNSPublicResolver *QRNSPublicResolverCallerSession) Zonehash(node [32]byte) ([]byte, error) {
	return _QRNSPublicResolver.Contract.Zonehash(&_QRNSPublicResolver.CallOpts, node)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Hyperion: function approve(bytes32 node, address delegate, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) Approve(opts *bind.TransactOpts, node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "approve", node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Hyperion: function approve(bytes32 node, address delegate, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.Approve(&_QRNSPublicResolver.TransactOpts, node, delegate, approved)
}

// Approve is a paid mutator transaction binding the contract method 0xa4b91a01.
//
// Hyperion: function approve(bytes32 node, address delegate, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) Approve(node [32]byte, delegate common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.Approve(&_QRNSPublicResolver.TransactOpts, node, delegate, approved)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Hyperion: function clearRecords(bytes32 node) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) ClearRecords(opts *bind.TransactOpts, node [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "clearRecords", node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Hyperion: function clearRecords(bytes32 node) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.ClearRecords(&_QRNSPublicResolver.TransactOpts, node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Hyperion: function clearRecords(bytes32 node) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) ClearRecords(node [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.ClearRecords(&_QRNSPublicResolver.TransactOpts, node)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Hyperion: function multicall(bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Hyperion: function multicall(bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.Multicall(&_QRNSPublicResolver.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Hyperion: function multicall(bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.Multicall(&_QRNSPublicResolver.TransactOpts, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Hyperion: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) MulticallWithNodeCheck(opts *bind.TransactOpts, nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "multicallWithNodeCheck", nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Hyperion: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverSession) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.MulticallWithNodeCheck(&_QRNSPublicResolver.TransactOpts, nodehash, data)
}

// MulticallWithNodeCheck is a paid mutator transaction binding the contract method 0xe32954eb.
//
// Hyperion: function multicallWithNodeCheck(bytes32 nodehash, bytes[] data) returns(bytes[] results)
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) MulticallWithNodeCheck(nodehash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.MulticallWithNodeCheck(&_QRNSPublicResolver.TransactOpts, nodehash, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Hyperion: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Hyperion: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetABI(&_QRNSPublicResolver.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Hyperion: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetABI(&_QRNSPublicResolver.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Hyperion: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Hyperion: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetAddr(&_QRNSPublicResolver.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Hyperion: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetAddr(&_QRNSPublicResolver.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Hyperion: function setAddr(bytes32 node, address a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Hyperion: function setAddr(bytes32 node, address a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetAddr0(&_QRNSPublicResolver.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Hyperion: function setAddr(bytes32 node, address a) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetAddr0(&_QRNSPublicResolver.TransactOpts, node, a)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Hyperion: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Hyperion: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetApprovalForAll(&_QRNSPublicResolver.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Hyperion: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetApprovalForAll(&_QRNSPublicResolver.TransactOpts, operator, approved)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Hyperion: function setContenthash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Hyperion: function setContenthash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetContenthash(&_QRNSPublicResolver.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Hyperion: function setContenthash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetContenthash(&_QRNSPublicResolver.TransactOpts, node, hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Hyperion: function setDNSRecords(bytes32 node, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetDNSRecords(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setDNSRecords", node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Hyperion: function setDNSRecords(bytes32 node, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetDNSRecords(&_QRNSPublicResolver.TransactOpts, node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Hyperion: function setDNSRecords(bytes32 node, bytes data) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetDNSRecords(&_QRNSPublicResolver.TransactOpts, node, data)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Hyperion: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetInterface(opts *bind.TransactOpts, node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setInterface", node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Hyperion: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetInterface(&_QRNSPublicResolver.TransactOpts, node, interfaceID, implementer)
}

// SetInterface is a paid mutator transaction binding the contract method 0xe59d895d.
//
// Hyperion: function setInterface(bytes32 node, bytes4 interfaceID, address implementer) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetInterface(node [32]byte, interfaceID [4]byte, implementer common.Address) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetInterface(&_QRNSPublicResolver.TransactOpts, node, interfaceID, implementer)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Hyperion: function setName(bytes32 node, string newName) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, newName string) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setName", node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Hyperion: function setName(bytes32 node, string newName) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetName(&_QRNSPublicResolver.TransactOpts, node, newName)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Hyperion: function setName(bytes32 node, string newName) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetName(node [32]byte, newName string) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetName(&_QRNSPublicResolver.TransactOpts, node, newName)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Hyperion: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Hyperion: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetPubkey(&_QRNSPublicResolver.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Hyperion: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetPubkey(&_QRNSPublicResolver.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Hyperion: function setText(bytes32 node, string key, string value) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Hyperion: function setText(bytes32 node, string key, string value) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetText(&_QRNSPublicResolver.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Hyperion: function setText(bytes32 node, string key, string value) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetText(&_QRNSPublicResolver.TransactOpts, node, key, value)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Hyperion: function setZonehash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactor) SetZonehash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.contract.Transact(opts, "setZonehash", node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Hyperion: function setZonehash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverSession) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetZonehash(&_QRNSPublicResolver.TransactOpts, node, hash)
}

// SetZonehash is a paid mutator transaction binding the contract method 0xce3decdc.
//
// Hyperion: function setZonehash(bytes32 node, bytes hash) returns()
func (_QRNSPublicResolver *QRNSPublicResolverTransactorSession) SetZonehash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _QRNSPublicResolver.Contract.SetZonehash(&_QRNSPublicResolver.TransactOpts, node, hash)
}

// QRNSPublicResolverABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverABIChangedIterator struct {
	Event *QRNSPublicResolverABIChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverABIChanged)
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
		it.Event = new(QRNSPublicResolverABIChanged)
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
func (it *QRNSPublicResolverABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverABIChanged represents a ABIChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Hyperion: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*QRNSPublicResolverABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverABIChangedIterator{contract: _QRNSPublicResolver.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Hyperion: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverABIChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Hyperion: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseABIChanged(log types.Log) (*QRNSPublicResolverABIChanged, error) {
	event := new(QRNSPublicResolverABIChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverAddrChangedIterator struct {
	Event *QRNSPublicResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverAddrChanged)
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
		it.Event = new(QRNSPublicResolverAddrChanged)
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
func (it *QRNSPublicResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverAddrChanged represents a AddrChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Hyperion: event AddrChanged(bytes32 indexed node, address a)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverAddrChangedIterator{contract: _QRNSPublicResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Hyperion: event AddrChanged(bytes32 indexed node, address a)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverAddrChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Hyperion: event AddrChanged(bytes32 indexed node, address a)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseAddrChanged(log types.Log) (*QRNSPublicResolverAddrChanged, error) {
	event := new(QRNSPublicResolverAddrChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverAddressChangedIterator struct {
	Event *QRNSPublicResolverAddressChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverAddressChanged)
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
		it.Event = new(QRNSPublicResolverAddressChanged)
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
func (it *QRNSPublicResolverAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverAddressChanged represents a AddressChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverAddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Hyperion: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverAddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverAddressChangedIterator{contract: _QRNSPublicResolver.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Hyperion: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverAddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverAddressChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Hyperion: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseAddressChanged(log types.Log) (*QRNSPublicResolverAddressChanged, error) {
	event := new(QRNSPublicResolverAddressChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverApprovalForAllIterator struct {
	Event *QRNSPublicResolverApprovalForAll // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverApprovalForAll)
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
		it.Event = new(QRNSPublicResolverApprovalForAll)
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
func (it *QRNSPublicResolverApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverApprovalForAll represents a ApprovalForAll event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Hyperion: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*QRNSPublicResolverApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverApprovalForAllIterator{contract: _QRNSPublicResolver.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Hyperion: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverApprovalForAll)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Hyperion: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseApprovalForAll(log types.Log) (*QRNSPublicResolverApprovalForAll, error) {
	event := new(QRNSPublicResolverApprovalForAll)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverApprovedIterator struct {
	Event *QRNSPublicResolverApproved // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverApproved)
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
		it.Event = new(QRNSPublicResolverApproved)
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
func (it *QRNSPublicResolverApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverApproved represents a Approved event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverApproved struct {
	Owner    common.Address
	Node     [32]byte
	Delegate common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Hyperion: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterApproved(opts *bind.FilterOpts, node [][32]byte, delegate []common.Address, approved []bool) (*QRNSPublicResolverApprovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverApprovedIterator{contract: _QRNSPublicResolver.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Hyperion: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverApproved, node [][32]byte, delegate []common.Address, approved []bool) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "Approved", nodeRule, delegateRule, approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverApproved)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0xf0ddb3b04746704017f9aa8bd728fcc2c1d11675041205350018915f5e4750a0.
//
// Hyperion: event Approved(address owner, bytes32 indexed node, address indexed delegate, bool indexed approved)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseApproved(log types.Log) (*QRNSPublicResolverApproved, error) {
	event := new(QRNSPublicResolverApproved)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverContenthashChangedIterator struct {
	Event *QRNSPublicResolverContenthashChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverContenthashChanged)
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
		it.Event = new(QRNSPublicResolverContenthashChanged)
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
func (it *QRNSPublicResolverContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverContenthashChanged represents a ContenthashChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Hyperion: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverContenthashChangedIterator{contract: _QRNSPublicResolver.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Hyperion: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverContenthashChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Hyperion: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseContenthashChanged(log types.Log) (*QRNSPublicResolverContenthashChanged, error) {
	event := new(QRNSPublicResolverContenthashChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverDNSRecordChangedIterator is returned from FilterDNSRecordChanged and is used to iterate over the raw logs and unpacked data for DNSRecordChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSRecordChangedIterator struct {
	Event *QRNSPublicResolverDNSRecordChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverDNSRecordChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverDNSRecordChanged)
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
		it.Event = new(QRNSPublicResolverDNSRecordChanged)
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
func (it *QRNSPublicResolverDNSRecordChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverDNSRecordChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverDNSRecordChanged represents a DNSRecordChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSRecordChanged struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Record   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordChanged is a free log retrieval operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Hyperion: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterDNSRecordChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverDNSRecordChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverDNSRecordChangedIterator{contract: _QRNSPublicResolver.contract, event: "DNSRecordChanged", logs: logs, sub: sub}, nil
}

// WatchDNSRecordChanged is a free log subscription operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Hyperion: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchDNSRecordChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverDNSRecordChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "DNSRecordChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverDNSRecordChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
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

// ParseDNSRecordChanged is a log parse operation binding the contract event 0x52a608b3303a48862d07a73d82fa221318c0027fbbcfb1b2329bface3f19ff2b.
//
// Hyperion: event DNSRecordChanged(bytes32 indexed node, bytes name, uint16 resource, bytes record)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseDNSRecordChanged(log types.Log) (*QRNSPublicResolverDNSRecordChanged, error) {
	event := new(QRNSPublicResolverDNSRecordChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSRecordChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverDNSRecordDeletedIterator is returned from FilterDNSRecordDeleted and is used to iterate over the raw logs and unpacked data for DNSRecordDeleted events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSRecordDeletedIterator struct {
	Event *QRNSPublicResolverDNSRecordDeleted // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverDNSRecordDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverDNSRecordDeleted)
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
		it.Event = new(QRNSPublicResolverDNSRecordDeleted)
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
func (it *QRNSPublicResolverDNSRecordDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverDNSRecordDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverDNSRecordDeleted represents a DNSRecordDeleted event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSRecordDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDNSRecordDeleted is a free log retrieval operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Hyperion: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterDNSRecordDeleted(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverDNSRecordDeletedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverDNSRecordDeletedIterator{contract: _QRNSPublicResolver.contract, event: "DNSRecordDeleted", logs: logs, sub: sub}, nil
}

// WatchDNSRecordDeleted is a free log subscription operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Hyperion: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchDNSRecordDeleted(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverDNSRecordDeleted, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "DNSRecordDeleted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverDNSRecordDeleted)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
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

// ParseDNSRecordDeleted is a log parse operation binding the contract event 0x03528ed0c2a3ebc993b12ce3c16bb382f9c7d88ef7d8a1bf290eaf35955a1207.
//
// Hyperion: event DNSRecordDeleted(bytes32 indexed node, bytes name, uint16 resource)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseDNSRecordDeleted(log types.Log) (*QRNSPublicResolverDNSRecordDeleted, error) {
	event := new(QRNSPublicResolverDNSRecordDeleted)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSRecordDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverDNSZonehashChangedIterator is returned from FilterDNSZonehashChanged and is used to iterate over the raw logs and unpacked data for DNSZonehashChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSZonehashChangedIterator struct {
	Event *QRNSPublicResolverDNSZonehashChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverDNSZonehashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverDNSZonehashChanged)
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
		it.Event = new(QRNSPublicResolverDNSZonehashChanged)
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
func (it *QRNSPublicResolverDNSZonehashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverDNSZonehashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverDNSZonehashChanged represents a DNSZonehashChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverDNSZonehashChanged struct {
	Node         [32]byte
	Lastzonehash []byte
	Zonehash     []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDNSZonehashChanged is a free log retrieval operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Hyperion: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterDNSZonehashChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverDNSZonehashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverDNSZonehashChangedIterator{contract: _QRNSPublicResolver.contract, event: "DNSZonehashChanged", logs: logs, sub: sub}, nil
}

// WatchDNSZonehashChanged is a free log subscription operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Hyperion: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchDNSZonehashChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverDNSZonehashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "DNSZonehashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverDNSZonehashChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
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

// ParseDNSZonehashChanged is a log parse operation binding the contract event 0x8f15ed4b723ef428f250961da8315675b507046737e19319fc1a4d81bfe87f85.
//
// Hyperion: event DNSZonehashChanged(bytes32 indexed node, bytes lastzonehash, bytes zonehash)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseDNSZonehashChanged(log types.Log) (*QRNSPublicResolverDNSZonehashChanged, error) {
	event := new(QRNSPublicResolverDNSZonehashChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "DNSZonehashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverInterfaceChangedIterator is returned from FilterInterfaceChanged and is used to iterate over the raw logs and unpacked data for InterfaceChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverInterfaceChangedIterator struct {
	Event *QRNSPublicResolverInterfaceChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverInterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverInterfaceChanged)
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
		it.Event = new(QRNSPublicResolverInterfaceChanged)
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
func (it *QRNSPublicResolverInterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverInterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverInterfaceChanged represents a InterfaceChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverInterfaceChanged struct {
	Node        [32]byte
	InterfaceID [4]byte
	Implementer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterfaceChanged is a free log retrieval operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Hyperion: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterInterfaceChanged(opts *bind.FilterOpts, node [][32]byte, interfaceID [][4]byte) (*QRNSPublicResolverInterfaceChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverInterfaceChangedIterator{contract: _QRNSPublicResolver.contract, event: "InterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchInterfaceChanged is a free log subscription operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Hyperion: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchInterfaceChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverInterfaceChanged, node [][32]byte, interfaceID [][4]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var interfaceIDRule []interface{}
	for _, interfaceIDItem := range interfaceID {
		interfaceIDRule = append(interfaceIDRule, interfaceIDItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "InterfaceChanged", nodeRule, interfaceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverInterfaceChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
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

// ParseInterfaceChanged is a log parse operation binding the contract event 0x7c69f06bea0bdef565b709e93a147836b0063ba2dd89f02d0b7e8d931e6a6daa.
//
// Hyperion: event InterfaceChanged(bytes32 indexed node, bytes4 indexed interfaceID, address implementer)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseInterfaceChanged(log types.Log) (*QRNSPublicResolverInterfaceChanged, error) {
	event := new(QRNSPublicResolverInterfaceChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "InterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverNameChangedIterator struct {
	Event *QRNSPublicResolverNameChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverNameChanged)
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
		it.Event = new(QRNSPublicResolverNameChanged)
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
func (it *QRNSPublicResolverNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverNameChanged represents a NameChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Hyperion: event NameChanged(bytes32 indexed node, string name)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverNameChangedIterator{contract: _QRNSPublicResolver.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Hyperion: event NameChanged(bytes32 indexed node, string name)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverNameChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Hyperion: event NameChanged(bytes32 indexed node, string name)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseNameChanged(log types.Log) (*QRNSPublicResolverNameChanged, error) {
	event := new(QRNSPublicResolverNameChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverPubkeyChangedIterator struct {
	Event *QRNSPublicResolverPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverPubkeyChanged)
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
		it.Event = new(QRNSPublicResolverPubkeyChanged)
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
func (it *QRNSPublicResolverPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverPubkeyChanged represents a PubkeyChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Hyperion: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverPubkeyChangedIterator{contract: _QRNSPublicResolver.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Hyperion: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverPubkeyChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Hyperion: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParsePubkeyChanged(log types.Log) (*QRNSPublicResolverPubkeyChanged, error) {
	event := new(QRNSPublicResolverPubkeyChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverTextChangedIterator struct {
	Event *QRNSPublicResolverTextChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverTextChanged)
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
		it.Event = new(QRNSPublicResolverTextChanged)
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
func (it *QRNSPublicResolverTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverTextChanged represents a TextChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Hyperion: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*QRNSPublicResolverTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverTextChangedIterator{contract: _QRNSPublicResolver.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Hyperion: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverTextChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0x448bc014f1536726cf8d54ff3d6481ed3cbc683c2591ca204274009afa09b1a1.
//
// Hyperion: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key, string value)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseTextChanged(log types.Log) (*QRNSPublicResolverTextChanged, error) {
	event := new(QRNSPublicResolverTextChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSPublicResolverVersionChangedIterator is returned from FilterVersionChanged and is used to iterate over the raw logs and unpacked data for VersionChanged events raised by the QRNSPublicResolver contract.
type QRNSPublicResolverVersionChangedIterator struct {
	Event *QRNSPublicResolverVersionChanged // Event containing the contract specifics and raw log

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
func (it *QRNSPublicResolverVersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSPublicResolverVersionChanged)
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
		it.Event = new(QRNSPublicResolverVersionChanged)
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
func (it *QRNSPublicResolverVersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSPublicResolverVersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSPublicResolverVersionChanged represents a VersionChanged event raised by the QRNSPublicResolver contract.
type QRNSPublicResolverVersionChanged struct {
	Node       [32]byte
	NewVersion uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionChanged is a free log retrieval operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Hyperion: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) FilterVersionChanged(opts *bind.FilterOpts, node [][32]byte) (*QRNSPublicResolverVersionChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.FilterLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSPublicResolverVersionChangedIterator{contract: _QRNSPublicResolver.contract, event: "VersionChanged", logs: logs, sub: sub}, nil
}

// WatchVersionChanged is a free log subscription operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Hyperion: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) WatchVersionChanged(opts *bind.WatchOpts, sink chan<- *QRNSPublicResolverVersionChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSPublicResolver.contract.WatchLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSPublicResolverVersionChanged)
				if err := _QRNSPublicResolver.contract.UnpackLog(event, "VersionChanged", log); err != nil {
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

// ParseVersionChanged is a log parse operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Hyperion: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_QRNSPublicResolver *QRNSPublicResolverFilterer) ParseVersionChanged(log types.Log) (*QRNSPublicResolverVersionChanged, error) {
	event := new(QRNSPublicResolverVersionChanged)
	if err := _QRNSPublicResolver.contract.UnpackLog(event, "VersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
