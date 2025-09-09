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

// DNSSECRRSetWithSignature is an auto generated low-level Go binding around an user-defined struct.
type DNSSECRRSetWithSignature struct {
	Rrset []byte
	Sig   []byte
}

// QRNSDNSRegistrarMetaData contains all meta data concerning the QRNSDNSRegistrar contract.
var QRNSDNSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousRegistrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"contractDNSSEC\",\"name\":\"_dnssec\",\"type\":\"address\"},{\"internalType\":\"contractPublicSuffixList\",\"name\":\"_suffixes\",\"type\":\"address\"},{\"internalType\":\"contractQRNS\",\"name\":\"_qrns\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"InvalidPublicSuffix\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOwnerRecordFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreconditionNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dnsname\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"inception\",\"type\":\"uint32\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"suffixes\",\"type\":\"address\"}],\"name\":\"NewPublicSuffixList\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"domain\",\"type\":\"bytes\"}],\"name\":\"enableNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qrns\",\"outputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inceptions\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractDNSSEC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousRegistrar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rrset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structDNSSEC.RRSetWithSignature[]\",\"name\":\"input\",\"type\":\"tuple[]\"}],\"name\":\"proveAndClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rrset\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structDNSSEC.RRSetWithSignature[]\",\"name\":\"input\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"proveAndClaimWithResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPublicSuffixList\",\"name\":\"_suffixes\",\"type\":\"address\"}],\"name\":\"setPublicSuffixList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suffixes\",\"outputs\":[{\"internalType\":\"contractPublicSuffixList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// QRNSDNSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSDNSRegistrarMetaData.ABI instead.
var QRNSDNSRegistrarABI = QRNSDNSRegistrarMetaData.ABI

// QRNSDNSRegistrar is an auto generated Go binding around a QRL contract.
type QRNSDNSRegistrar struct {
	QRNSDNSRegistrarCaller     // Read-only binding to the contract
	QRNSDNSRegistrarTransactor // Write-only binding to the contract
	QRNSDNSRegistrarFilterer   // Log filterer for contract events
}

// QRNSDNSRegistrarCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSDNSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSDNSRegistrarTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSDNSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSDNSRegistrarFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSDNSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSDNSRegistrarSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSDNSRegistrarSession struct {
	Contract     *QRNSDNSRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QRNSDNSRegistrarCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSDNSRegistrarCallerSession struct {
	Contract *QRNSDNSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// QRNSDNSRegistrarTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSDNSRegistrarTransactorSession struct {
	Contract     *QRNSDNSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QRNSDNSRegistrarRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSDNSRegistrarRaw struct {
	Contract *QRNSDNSRegistrar // Generic contract binding to access the raw methods on
}

// QRNSDNSRegistrarCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSDNSRegistrarCallerRaw struct {
	Contract *QRNSDNSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSDNSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSDNSRegistrarTransactorRaw struct {
	Contract *QRNSDNSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSDNSRegistrar creates a new instance of QRNSDNSRegistrar, bound to a specific deployed contract.
func NewQRNSDNSRegistrar(address common.Address, backend bind.ContractBackend) (*QRNSDNSRegistrar, error) {
	contract, err := bindQRNSDNSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrar{QRNSDNSRegistrarCaller: QRNSDNSRegistrarCaller{contract: contract}, QRNSDNSRegistrarTransactor: QRNSDNSRegistrarTransactor{contract: contract}, QRNSDNSRegistrarFilterer: QRNSDNSRegistrarFilterer{contract: contract}}, nil
}

// NewQRNSDNSRegistrarCaller creates a new read-only instance of QRNSDNSRegistrar, bound to a specific deployed contract.
func NewQRNSDNSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*QRNSDNSRegistrarCaller, error) {
	contract, err := bindQRNSDNSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrarCaller{contract: contract}, nil
}

// NewQRNSDNSRegistrarTransactor creates a new write-only instance of QRNSDNSRegistrar, bound to a specific deployed contract.
func NewQRNSDNSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSDNSRegistrarTransactor, error) {
	contract, err := bindQRNSDNSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrarTransactor{contract: contract}, nil
}

// NewQRNSDNSRegistrarFilterer creates a new log filterer instance of QRNSDNSRegistrar, bound to a specific deployed contract.
func NewQRNSDNSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSDNSRegistrarFilterer, error) {
	contract, err := bindQRNSDNSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrarFilterer{contract: contract}, nil
}

// bindQRNSDNSRegistrar binds a generic wrapper to an already deployed contract.
func bindQRNSDNSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSDNSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSDNSRegistrar.Contract.QRNSDNSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.QRNSDNSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.QRNSDNSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSDNSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Hyperion: function inceptions(bytes32 ) view returns(uint32)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) Inceptions(opts *bind.CallOpts, arg0 [32]byte) (uint32, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "inceptions", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Hyperion: function inceptions(bytes32 ) view returns(uint32)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) Inceptions(arg0 [32]byte) (uint32, error) {
	return _QRNSDNSRegistrar.Contract.Inceptions(&_QRNSDNSRegistrar.CallOpts, arg0)
}

// Inceptions is a free data retrieval call binding the contract method 0x25916d41.
//
// Hyperion: function inceptions(bytes32 ) view returns(uint32)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) Inceptions(arg0 [32]byte) (uint32, error) {
	return _QRNSDNSRegistrar.Contract.Inceptions(&_QRNSDNSRegistrar.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Hyperion: function oracle() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Hyperion: function oracle() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) Oracle() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Oracle(&_QRNSDNSRegistrar.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Hyperion: function oracle() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) Oracle() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Oracle(&_QRNSDNSRegistrar.CallOpts)
}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Hyperion: function previousRegistrar() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) PreviousRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "previousRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Hyperion: function previousRegistrar() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) PreviousRegistrar() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.PreviousRegistrar(&_QRNSDNSRegistrar.CallOpts)
}

// PreviousRegistrar is a free data retrieval call binding the contract method 0xab14ec59.
//
// Hyperion: function previousRegistrar() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) PreviousRegistrar() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.PreviousRegistrar(&_QRNSDNSRegistrar.CallOpts)
}

// Qrns is a free data retrieval call binding the contract method 0x0a105d31.
//
// Hyperion: function qrns() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) Qrns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "qrns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Qrns is a free data retrieval call binding the contract method 0x0a105d31.
//
// Hyperion: function qrns() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) Qrns() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Qrns(&_QRNSDNSRegistrar.CallOpts)
}

// Qrns is a free data retrieval call binding the contract method 0x0a105d31.
//
// Hyperion: function qrns() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) Qrns() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Qrns(&_QRNSDNSRegistrar.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Hyperion: function resolver() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Hyperion: function resolver() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) Resolver() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Resolver(&_QRNSDNSRegistrar.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Hyperion: function resolver() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) Resolver() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Resolver(&_QRNSDNSRegistrar.CallOpts)
}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Hyperion: function suffixes() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) Suffixes(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "suffixes")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Hyperion: function suffixes() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) Suffixes() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Suffixes(&_QRNSDNSRegistrar.CallOpts)
}

// Suffixes is a free data retrieval call binding the contract method 0x30349ebe.
//
// Hyperion: function suffixes() view returns(address)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) Suffixes() (common.Address, error) {
	return _QRNSDNSRegistrar.Contract.Suffixes(&_QRNSDNSRegistrar.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _QRNSDNSRegistrar.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSDNSRegistrar.Contract.SupportsInterface(&_QRNSDNSRegistrar.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Hyperion: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _QRNSDNSRegistrar.Contract.SupportsInterface(&_QRNSDNSRegistrar.CallOpts, interfaceID)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Hyperion: function enableNode(bytes domain) returns(bytes32 node)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactor) EnableNode(opts *bind.TransactOpts, domain []byte) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.contract.Transact(opts, "enableNode", domain)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Hyperion: function enableNode(bytes domain) returns(bytes32 node)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) EnableNode(domain []byte) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.EnableNode(&_QRNSDNSRegistrar.TransactOpts, domain)
}

// EnableNode is a paid mutator transaction binding the contract method 0x6f951221.
//
// Hyperion: function enableNode(bytes domain) returns(bytes32 node)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorSession) EnableNode(domain []byte) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.EnableNode(&_QRNSDNSRegistrar.TransactOpts, domain)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Hyperion: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactor) ProveAndClaim(opts *bind.TransactOpts, name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.contract.Transact(opts, "proveAndClaim", name, input)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Hyperion: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) ProveAndClaim(name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.ProveAndClaim(&_QRNSDNSRegistrar.TransactOpts, name, input)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0x29d56630.
//
// Hyperion: function proveAndClaim(bytes name, (bytes,bytes)[] input) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorSession) ProveAndClaim(name []byte, input []DNSSECRRSetWithSignature) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.ProveAndClaim(&_QRNSDNSRegistrar.TransactOpts, name, input)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Hyperion: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactor) ProveAndClaimWithResolver(opts *bind.TransactOpts, name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.contract.Transact(opts, "proveAndClaimWithResolver", name, input, resolver, addr)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Hyperion: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) ProveAndClaimWithResolver(name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.ProveAndClaimWithResolver(&_QRNSDNSRegistrar.TransactOpts, name, input, resolver, addr)
}

// ProveAndClaimWithResolver is a paid mutator transaction binding the contract method 0x06963218.
//
// Hyperion: function proveAndClaimWithResolver(bytes name, (bytes,bytes)[] input, address resolver, address addr) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorSession) ProveAndClaimWithResolver(name []byte, input []DNSSECRRSetWithSignature, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.ProveAndClaimWithResolver(&_QRNSDNSRegistrar.TransactOpts, name, input, resolver, addr)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Hyperion: function setPublicSuffixList(address _suffixes) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactor) SetPublicSuffixList(opts *bind.TransactOpts, _suffixes common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.contract.Transact(opts, "setPublicSuffixList", _suffixes)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Hyperion: function setPublicSuffixList(address _suffixes) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarSession) SetPublicSuffixList(_suffixes common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.SetPublicSuffixList(&_QRNSDNSRegistrar.TransactOpts, _suffixes)
}

// SetPublicSuffixList is a paid mutator transaction binding the contract method 0x1ecfc411.
//
// Hyperion: function setPublicSuffixList(address _suffixes) returns()
func (_QRNSDNSRegistrar *QRNSDNSRegistrarTransactorSession) SetPublicSuffixList(_suffixes common.Address) (*types.Transaction, error) {
	return _QRNSDNSRegistrar.Contract.SetPublicSuffixList(&_QRNSDNSRegistrar.TransactOpts, _suffixes)
}

// QRNSDNSRegistrarClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the QRNSDNSRegistrar contract.
type QRNSDNSRegistrarClaimIterator struct {
	Event *QRNSDNSRegistrarClaim // Event containing the contract specifics and raw log

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
func (it *QRNSDNSRegistrarClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSDNSRegistrarClaim)
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
		it.Event = new(QRNSDNSRegistrarClaim)
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
func (it *QRNSDNSRegistrarClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSDNSRegistrarClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSDNSRegistrarClaim represents a Claim event raised by the QRNSDNSRegistrar contract.
type QRNSDNSRegistrarClaim struct {
	Node      [32]byte
	Owner     common.Address
	Dnsname   []byte
	Inception uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Hyperion: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) FilterClaim(opts *bind.FilterOpts, node [][32]byte, owner []common.Address) (*QRNSDNSRegistrarClaimIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _QRNSDNSRegistrar.contract.FilterLogs(opts, "Claim", nodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrarClaimIterator{contract: _QRNSDNSRegistrar.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Hyperion: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *QRNSDNSRegistrarClaim, node [][32]byte, owner []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _QRNSDNSRegistrar.contract.WatchLogs(opts, "Claim", nodeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSDNSRegistrarClaim)
				if err := _QRNSDNSRegistrar.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x87db02a0e483e2818060eddcbb3488ce44e35aff49a70d92c2aa6c8046cf01e2.
//
// Hyperion: event Claim(bytes32 indexed node, address indexed owner, bytes dnsname, uint32 inception)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) ParseClaim(log types.Log) (*QRNSDNSRegistrarClaim, error) {
	event := new(QRNSDNSRegistrarClaim)
	if err := _QRNSDNSRegistrar.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSDNSRegistrarNewPublicSuffixListIterator is returned from FilterNewPublicSuffixList and is used to iterate over the raw logs and unpacked data for NewPublicSuffixList events raised by the QRNSDNSRegistrar contract.
type QRNSDNSRegistrarNewPublicSuffixListIterator struct {
	Event *QRNSDNSRegistrarNewPublicSuffixList // Event containing the contract specifics and raw log

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
func (it *QRNSDNSRegistrarNewPublicSuffixListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSDNSRegistrarNewPublicSuffixList)
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
		it.Event = new(QRNSDNSRegistrarNewPublicSuffixList)
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
func (it *QRNSDNSRegistrarNewPublicSuffixListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSDNSRegistrarNewPublicSuffixListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSDNSRegistrarNewPublicSuffixList represents a NewPublicSuffixList event raised by the QRNSDNSRegistrar contract.
type QRNSDNSRegistrarNewPublicSuffixList struct {
	Suffixes common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPublicSuffixList is a free log retrieval operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Hyperion: event NewPublicSuffixList(address suffixes)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) FilterNewPublicSuffixList(opts *bind.FilterOpts) (*QRNSDNSRegistrarNewPublicSuffixListIterator, error) {

	logs, sub, err := _QRNSDNSRegistrar.contract.FilterLogs(opts, "NewPublicSuffixList")
	if err != nil {
		return nil, err
	}
	return &QRNSDNSRegistrarNewPublicSuffixListIterator{contract: _QRNSDNSRegistrar.contract, event: "NewPublicSuffixList", logs: logs, sub: sub}, nil
}

// WatchNewPublicSuffixList is a free log subscription operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Hyperion: event NewPublicSuffixList(address suffixes)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) WatchNewPublicSuffixList(opts *bind.WatchOpts, sink chan<- *QRNSDNSRegistrarNewPublicSuffixList) (event.Subscription, error) {

	logs, sub, err := _QRNSDNSRegistrar.contract.WatchLogs(opts, "NewPublicSuffixList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSDNSRegistrarNewPublicSuffixList)
				if err := _QRNSDNSRegistrar.contract.UnpackLog(event, "NewPublicSuffixList", log); err != nil {
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

// ParseNewPublicSuffixList is a log parse operation binding the contract event 0x9176b7f47e4504df5e5516c99d90d82ac7cbd49cc77e7f22ba2ac2f2e3a3eba8.
//
// Hyperion: event NewPublicSuffixList(address suffixes)
func (_QRNSDNSRegistrar *QRNSDNSRegistrarFilterer) ParseNewPublicSuffixList(log types.Log) (*QRNSDNSRegistrarNewPublicSuffixList, error) {
	event := new(QRNSDNSRegistrarNewPublicSuffixList)
	if err := _QRNSDNSRegistrar.contract.UnpackLog(event, "NewPublicSuffixList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
