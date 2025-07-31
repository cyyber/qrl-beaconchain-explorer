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

// QRNSNameWrapperMetaData contains all meta data concerning the QRNSNameWrapper contract.
var QRNSNameWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"_qrns\",\"type\":\"address\"},{\"internalType\":\"contractIBaseRegistrar\",\"name\":\"_registrar\",\"type\":\"address\"},{\"internalType\":\"contractIMetadataService\",\"name\":\"_metadataService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotUpgrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleParent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"IncorrectTargetOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedLabelhash\",\"type\":\"bytes32\"}],\"name\":\"LabelMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LabelTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameIsNotWrapped\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"OperationProhibited\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorised\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"ControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"ExpiryExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"}],\"name\":\"FusesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"NameWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokqrns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"fuseMask\",\"type\":\"uint32\"}],\"name\":\"allFusesBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canExtendSubnames\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canModifyName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qrns\",\"outputs\":[{\"internalType\":\"contractQRNS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"extendExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataService\",\"outputs\":[{\"internalType\":\"contractIMetadataService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onZRC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"registerAndWrapZOND2LD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"registrarExpiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"internalType\":\"contractIBaseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setChildFuses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"setFuses\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetadataService\",\"name\":\"_metadataService\",\"type\":\"address\"}],\"name\":\"setMetadataService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINameWrapperUpgrade\",\"name\":\"_upgradeAddress\",\"type\":\"address\"}],\"name\":\"setUpgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"unwrapZOND2LD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContract\",\"outputs\":[{\"internalType\":\"contractINameWrapperUpgrade\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"wrapZOND2LD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QRNSNameWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use QRNSNameWrapperMetaData.ABI instead.
var QRNSNameWrapperABI = QRNSNameWrapperMetaData.ABI

// QRNSNameWrapper is an auto generated Go binding around a QRL contract.
type QRNSNameWrapper struct {
	QRNSNameWrapperCaller     // Read-only binding to the contract
	QRNSNameWrapperTransactor // Write-only binding to the contract
	QRNSNameWrapperFilterer   // Log filterer for contract events
}

// QRNSNameWrapperCaller is an auto generated read-only Go binding around a QRL contract.
type QRNSNameWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSNameWrapperTransactor is an auto generated write-only Go binding around a QRL contract.
type QRNSNameWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSNameWrapperFilterer is an auto generated log filtering Go binding around a QRL contract events.
type QRNSNameWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QRNSNameWrapperSession is an auto generated Go binding around a QRL contract,
// with pre-set call and transact options.
type QRNSNameWrapperSession struct {
	Contract     *QRNSNameWrapper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QRNSNameWrapperCallerSession is an auto generated read-only Go binding around a QRL contract,
// with pre-set call options.
type QRNSNameWrapperCallerSession struct {
	Contract *QRNSNameWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// QRNSNameWrapperTransactorSession is an auto generated write-only Go binding around a QRL contract,
// with pre-set transact options.
type QRNSNameWrapperTransactorSession struct {
	Contract     *QRNSNameWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// QRNSNameWrapperRaw is an auto generated low-level Go binding around a QRL contract.
type QRNSNameWrapperRaw struct {
	Contract *QRNSNameWrapper // Generic contract binding to access the raw methods on
}

// QRNSNameWrapperCallerRaw is an auto generated low-level read-only Go binding around a QRL contract.
type QRNSNameWrapperCallerRaw struct {
	Contract *QRNSNameWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// QRNSNameWrapperTransactorRaw is an auto generated low-level write-only Go binding around a QRL contract.
type QRNSNameWrapperTransactorRaw struct {
	Contract *QRNSNameWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQRNSNameWrapper creates a new instance of QRNSNameWrapper, bound to a specific deployed contract.
func NewQRNSNameWrapper(address common.Address, backend bind.ContractBackend) (*QRNSNameWrapper, error) {
	contract, err := bindQRNSNameWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapper{QRNSNameWrapperCaller: QRNSNameWrapperCaller{contract: contract}, QRNSNameWrapperTransactor: QRNSNameWrapperTransactor{contract: contract}, QRNSNameWrapperFilterer: QRNSNameWrapperFilterer{contract: contract}}, nil
}

// NewQRNSNameWrapperCaller creates a new read-only instance of QRNSNameWrapper, bound to a specific deployed contract.
func NewQRNSNameWrapperCaller(address common.Address, caller bind.ContractCaller) (*QRNSNameWrapperCaller, error) {
	contract, err := bindQRNSNameWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperCaller{contract: contract}, nil
}

// NewQRNSNameWrapperTransactor creates a new write-only instance of QRNSNameWrapper, bound to a specific deployed contract.
func NewQRNSNameWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*QRNSNameWrapperTransactor, error) {
	contract, err := bindQRNSNameWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperTransactor{contract: contract}, nil
}

// NewQRNSNameWrapperFilterer creates a new log filterer instance of QRNSNameWrapper, bound to a specific deployed contract.
func NewQRNSNameWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*QRNSNameWrapperFilterer, error) {
	contract, err := bindQRNSNameWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperFilterer{contract: contract}, nil
}

// bindQRNSNameWrapper binds a generic wrapper to an already deployed contract.
func bindQRNSNameWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QRNSNameWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSNameWrapper *QRNSNameWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSNameWrapper.Contract.QRNSNameWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSNameWrapper *QRNSNameWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.QRNSNameWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSNameWrapper *QRNSNameWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.QRNSNameWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QRNSNameWrapper *QRNSNameWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QRNSNameWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QRNSNameWrapper *QRNSNameWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QRNSNameWrapper *QRNSNameWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.contract.Transact(opts, method, params...)
}

// Tokqrns is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokqrns(uint256 ) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Tokqrns(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "_tokqrns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokqrns is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokqrns(uint256 ) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Tokqrns(arg0 *big.Int) (*big.Int, error) {
	return _QRNSNameWrapper.Contract.Tokqrns(&_QRNSNameWrapper.CallOpts, arg0)
}

// Tokqrns is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokqrns(uint256 ) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Tokqrns(arg0 *big.Int) (*big.Int, error) {
	return _QRNSNameWrapper.Contract.Tokqrns(&_QRNSNameWrapper.CallOpts, arg0)
}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) AllFusesBurned(opts *bind.CallOpts, node [32]byte, fuseMask uint32) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "allFusesBurned", node, fuseMask)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) AllFusesBurned(node [32]byte, fuseMask uint32) (bool, error) {
	return _QRNSNameWrapper.Contract.AllFusesBurned(&_QRNSNameWrapper.CallOpts, node, fuseMask)
}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) AllFusesBurned(node [32]byte, fuseMask uint32) (bool, error) {
	return _QRNSNameWrapper.Contract.AllFusesBurned(&_QRNSNameWrapper.CallOpts, node, fuseMask)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _QRNSNameWrapper.Contract.BalanceOf(&_QRNSNameWrapper.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _QRNSNameWrapper.Contract.BalanceOf(&_QRNSNameWrapper.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_QRNSNameWrapper *QRNSNameWrapperCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_QRNSNameWrapper *QRNSNameWrapperSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _QRNSNameWrapper.Contract.BalanceOfBatch(&_QRNSNameWrapper.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _QRNSNameWrapper.Contract.BalanceOfBatch(&_QRNSNameWrapper.CallOpts, accounts, ids)
}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) CanExtendSubnames(opts *bind.CallOpts, node [32]byte, addr common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "canExtendSubnames", node, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) CanExtendSubnames(node [32]byte, addr common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.CanExtendSubnames(&_QRNSNameWrapper.CallOpts, node, addr)
}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) CanExtendSubnames(node [32]byte, addr common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.CanExtendSubnames(&_QRNSNameWrapper.CallOpts, node, addr)
}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) CanModifyName(opts *bind.CallOpts, node [32]byte, addr common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "canModifyName", node, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) CanModifyName(node [32]byte, addr common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.CanModifyName(&_QRNSNameWrapper.CallOpts, node, addr)
}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) CanModifyName(node [32]byte, addr common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.CanModifyName(&_QRNSNameWrapper.CallOpts, node, addr)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Controllers(arg0 common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.Controllers(&_QRNSNameWrapper.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.Controllers(&_QRNSNameWrapper.CallOpts, arg0)
}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Qrns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "qrns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Qrns() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Qrns(&_QRNSNameWrapper.CallOpts)
}

// Qrns is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function qrns() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Qrns() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Qrns(&_QRNSNameWrapper.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) GetApproved(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "getApproved", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_QRNSNameWrapper *QRNSNameWrapperSession) GetApproved(id *big.Int) (common.Address, error) {
	return _QRNSNameWrapper.Contract.GetApproved(&_QRNSNameWrapper.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) GetApproved(id *big.Int) (common.Address, error) {
	return _QRNSNameWrapper.Contract.GetApproved(&_QRNSNameWrapper.CallOpts, id)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) GetData(opts *bind.CallOpts, id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "getData", id)

	outstruct := new(struct {
		Owner  common.Address
		Fuses  uint32
		Expiry uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fuses = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Expiry = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperSession) GetData(id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	return _QRNSNameWrapper.Contract.GetData(&_QRNSNameWrapper.CallOpts, id)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) GetData(id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	return _QRNSNameWrapper.Contract.GetData(&_QRNSNameWrapper.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.IsApprovedForAll(&_QRNSNameWrapper.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _QRNSNameWrapper.Contract.IsApprovedForAll(&_QRNSNameWrapper.CallOpts, account, operator)
}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) IsWrapped(opts *bind.CallOpts, parentNode [32]byte, labelhash [32]byte) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "isWrapped", parentNode, labelhash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) IsWrapped(parentNode [32]byte, labelhash [32]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.IsWrapped(&_QRNSNameWrapper.CallOpts, parentNode, labelhash)
}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) IsWrapped(parentNode [32]byte, labelhash [32]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.IsWrapped(&_QRNSNameWrapper.CallOpts, parentNode, labelhash)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) IsWrapped0(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "isWrapped0", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) IsWrapped0(node [32]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.IsWrapped0(&_QRNSNameWrapper.CallOpts, node)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) IsWrapped0(node [32]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.IsWrapped0(&_QRNSNameWrapper.CallOpts, node)
}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) MetadataService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "metadataService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperSession) MetadataService() (common.Address, error) {
	return _QRNSNameWrapper.Contract.MetadataService(&_QRNSNameWrapper.CallOpts)
}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) MetadataService() (common.Address, error) {
	return _QRNSNameWrapper.Contract.MetadataService(&_QRNSNameWrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Name() (string, error) {
	return _QRNSNameWrapper.Contract.Name(&_QRNSNameWrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Name() (string, error) {
	return _QRNSNameWrapper.Contract.Name(&_QRNSNameWrapper.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Names(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Names(arg0 [32]byte) ([]byte, error) {
	return _QRNSNameWrapper.Contract.Names(&_QRNSNameWrapper.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Names(arg0 [32]byte) ([]byte, error) {
	return _QRNSNameWrapper.Contract.Names(&_QRNSNameWrapper.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Owner() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Owner(&_QRNSNameWrapper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Owner() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Owner(&_QRNSNameWrapper.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_QRNSNameWrapper *QRNSNameWrapperSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _QRNSNameWrapper.Contract.OwnerOf(&_QRNSNameWrapper.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _QRNSNameWrapper.Contract.OwnerOf(&_QRNSNameWrapper.CallOpts, id)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Registrar() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Registrar(&_QRNSNameWrapper.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Registrar() (common.Address, error) {
	return _QRNSNameWrapper.Contract.Registrar(&_QRNSNameWrapper.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.SupportsInterface(&_QRNSNameWrapper.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QRNSNameWrapper.Contract.SupportsInterface(&_QRNSNameWrapper.CallOpts, interfaceId)
}

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) UpgradeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "upgradeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperSession) UpgradeContract() (common.Address, error) {
	return _QRNSNameWrapper.Contract.UpgradeContract(&_QRNSNameWrapper.CallOpts)
}

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) UpgradeContract() (common.Address, error) {
	return _QRNSNameWrapper.Contract.UpgradeContract(&_QRNSNameWrapper.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _QRNSNameWrapper.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Uri(tokenId *big.Int) (string, error) {
	return _QRNSNameWrapper.Contract.Uri(&_QRNSNameWrapper.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_QRNSNameWrapper *QRNSNameWrapperCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _QRNSNameWrapper.Contract.Uri(&_QRNSNameWrapper.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Approve(&_QRNSNameWrapper.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Approve(&_QRNSNameWrapper.TransactOpts, to, tokenId)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) ExtendExpiry(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "extendExpiry", parentNode, labelhash, expiry)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_QRNSNameWrapper *QRNSNameWrapperSession) ExtendExpiry(parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.ExtendExpiry(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, expiry)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) ExtendExpiry(parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.ExtendExpiry(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, expiry)
}

// OnZRC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onZRC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) OnZRC721Received(opts *bind.TransactOpts, to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "onZRC721Received", to, arg1, tokenId, data)
}

// OnZRC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onZRC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_QRNSNameWrapper *QRNSNameWrapperSession) OnZRC721Received(to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.OnZRC721Received(&_QRNSNameWrapper.TransactOpts, to, arg1, tokenId, data)
}

// OnZRC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onZRC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) OnZRC721Received(to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.OnZRC721Received(&_QRNSNameWrapper.TransactOpts, to, arg1, tokenId, data)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RecoverFunds(&_QRNSNameWrapper.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RecoverFunds(&_QRNSNameWrapper.TransactOpts, _token, _to, _amount)
}

// RegisterAndWrapZOND2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapZOND2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) RegisterAndWrapZOND2LD(opts *bind.TransactOpts, label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "registerAndWrapZOND2LD", label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// RegisterAndWrapZOND2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapZOND2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_QRNSNameWrapper *QRNSNameWrapperSession) RegisterAndWrapZOND2LD(label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RegisterAndWrapZOND2LD(&_QRNSNameWrapper.TransactOpts, label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// RegisterAndWrapZOND2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapZOND2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) RegisterAndWrapZOND2LD(label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RegisterAndWrapZOND2LD(&_QRNSNameWrapper.TransactOpts, label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) Renew(opts *bind.TransactOpts, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "renew", tokenId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_QRNSNameWrapper *QRNSNameWrapperSession) Renew(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Renew(&_QRNSNameWrapper.TransactOpts, tokenId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) Renew(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Renew(&_QRNSNameWrapper.TransactOpts, tokenId, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RenounceOwnership(&_QRNSNameWrapper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.RenounceOwnership(&_QRNSNameWrapper.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SafeBatchTransferFrom(&_QRNSNameWrapper.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SafeBatchTransferFrom(&_QRNSNameWrapper.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SafeTransferFrom(&_QRNSNameWrapper.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SafeTransferFrom(&_QRNSNameWrapper.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetApprovalForAll(&_QRNSNameWrapper.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetApprovalForAll(&_QRNSNameWrapper.TransactOpts, operator, approved)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetChildFuses(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setChildFuses", parentNode, labelhash, fuses, expiry)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetChildFuses(parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetChildFuses(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, fuses, expiry)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetChildFuses(parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetChildFuses(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, fuses, expiry)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetController(opts *bind.TransactOpts, controller common.Address, active bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setController", controller, active)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetController(controller common.Address, active bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetController(&_QRNSNameWrapper.TransactOpts, controller, active)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetController(controller common.Address, active bool) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetController(&_QRNSNameWrapper.TransactOpts, controller, active)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetFuses(opts *bind.TransactOpts, node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setFuses", node, ownerControlledFuses)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetFuses(node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetFuses(&_QRNSNameWrapper.TransactOpts, node, ownerControlledFuses)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetFuses(node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetFuses(&_QRNSNameWrapper.TransactOpts, node, ownerControlledFuses)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetMetadataService(opts *bind.TransactOpts, _metadataService common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setMetadataService", _metadataService)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetMetadataService(_metadataService common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetMetadataService(&_QRNSNameWrapper.TransactOpts, _metadataService)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetMetadataService(_metadataService common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetMetadataService(&_QRNSNameWrapper.TransactOpts, _metadataService)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetRecord(&_QRNSNameWrapper.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetRecord(&_QRNSNameWrapper.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetResolver(&_QRNSNameWrapper.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetResolver(&_QRNSNameWrapper.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetSubnodeOwner(opts *bind.TransactOpts, parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setSubnodeOwner", parentNode, label, owner, fuses, expiry)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetSubnodeOwner(parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetSubnodeOwner(&_QRNSNameWrapper.TransactOpts, parentNode, label, owner, fuses, expiry)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetSubnodeOwner(parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetSubnodeOwner(&_QRNSNameWrapper.TransactOpts, parentNode, label, owner, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetSubnodeRecord(opts *bind.TransactOpts, parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setSubnodeRecord", parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetSubnodeRecord(parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetSubnodeRecord(&_QRNSNameWrapper.TransactOpts, parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetSubnodeRecord(parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetSubnodeRecord(&_QRNSNameWrapper.TransactOpts, parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetTTL(&_QRNSNameWrapper.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetTTL(&_QRNSNameWrapper.TransactOpts, node, ttl)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) SetUpgradeContract(opts *bind.TransactOpts, _upgradeAddress common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "setUpgradeContract", _upgradeAddress)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) SetUpgradeContract(_upgradeAddress common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetUpgradeContract(&_QRNSNameWrapper.TransactOpts, _upgradeAddress)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) SetUpgradeContract(_upgradeAddress common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.SetUpgradeContract(&_QRNSNameWrapper.TransactOpts, _upgradeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.TransferOwnership(&_QRNSNameWrapper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.TransferOwnership(&_QRNSNameWrapper.TransactOpts, newOwner)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) Unwrap(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "unwrap", parentNode, labelhash, controller)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) Unwrap(parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Unwrap(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, controller)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) Unwrap(parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Unwrap(&_QRNSNameWrapper.TransactOpts, parentNode, labelhash, controller)
}

// UnwrapZOND2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapZOND2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) UnwrapZOND2LD(opts *bind.TransactOpts, labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "unwrapZOND2LD", labelhash, registrant, controller)
}

// UnwrapZOND2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapZOND2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) UnwrapZOND2LD(labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.UnwrapZOND2LD(&_QRNSNameWrapper.TransactOpts, labelhash, registrant, controller)
}

// UnwrapZOND2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapZOND2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) UnwrapZOND2LD(labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.UnwrapZOND2LD(&_QRNSNameWrapper.TransactOpts, labelhash, registrant, controller)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) Upgrade(opts *bind.TransactOpts, name []byte, extraData []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "upgrade", name, extraData)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) Upgrade(name []byte, extraData []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Upgrade(&_QRNSNameWrapper.TransactOpts, name, extraData)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) Upgrade(name []byte, extraData []byte) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Upgrade(&_QRNSNameWrapper.TransactOpts, name, extraData)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) Wrap(opts *bind.TransactOpts, name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "wrap", name, wrappedOwner, resolver)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperSession) Wrap(name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Wrap(&_QRNSNameWrapper.TransactOpts, name, wrappedOwner, resolver)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) Wrap(name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.Wrap(&_QRNSNameWrapper.TransactOpts, name, wrappedOwner, resolver)
}

// WrapZOND2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapZOND2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperTransactor) WrapZOND2LD(opts *bind.TransactOpts, label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.contract.Transact(opts, "wrapZOND2LD", label, wrappedOwner, ownerControlledFuses, resolver)
}

// WrapZOND2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapZOND2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperSession) WrapZOND2LD(label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.WrapZOND2LD(&_QRNSNameWrapper.TransactOpts, label, wrappedOwner, ownerControlledFuses, resolver)
}

// WrapZOND2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapZOND2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperTransactorSession) WrapZOND2LD(label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _QRNSNameWrapper.Contract.WrapZOND2LD(&_QRNSNameWrapper.TransactOpts, label, wrappedOwner, ownerControlledFuses, resolver)
}

// QRNSNameWrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperApprovalIterator struct {
	Event *QRNSNameWrapperApproval // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperApproval)
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
		it.Event = new(QRNSNameWrapperApproval)
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
func (it *QRNSNameWrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperApproval represents a Approval event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*QRNSNameWrapperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperApprovalIterator{contract: _QRNSNameWrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperApproval)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseApproval(log types.Log) (*QRNSNameWrapperApproval, error) {
	event := new(QRNSNameWrapperApproval)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperApprovalForAllIterator struct {
	Event *QRNSNameWrapperApprovalForAll // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperApprovalForAll)
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
		it.Event = new(QRNSNameWrapperApprovalForAll)
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
func (it *QRNSNameWrapperApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperApprovalForAll represents a ApprovalForAll event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*QRNSNameWrapperApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperApprovalForAllIterator{contract: _QRNSNameWrapper.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperApprovalForAll)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseApprovalForAll(log types.Log) (*QRNSNameWrapperApprovalForAll, error) {
	event := new(QRNSNameWrapperApprovalForAll)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperControllerChangedIterator struct {
	Event *QRNSNameWrapperControllerChanged // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperControllerChanged)
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
		it.Event = new(QRNSNameWrapperControllerChanged)
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
func (it *QRNSNameWrapperControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperControllerChanged represents a ControllerChanged event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperControllerChanged struct {
	Controller common.Address
	Active     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterControllerChanged(opts *bind.FilterOpts, controller []common.Address) (*QRNSNameWrapperControllerChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperControllerChangedIterator{contract: _QRNSNameWrapper.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperControllerChanged, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperControllerChanged)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseControllerChanged(log types.Log) (*QRNSNameWrapperControllerChanged, error) {
	event := new(QRNSNameWrapperControllerChanged)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperExpiryExtendedIterator is returned from FilterExpiryExtended and is used to iterate over the raw logs and unpacked data for ExpiryExtended events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperExpiryExtendedIterator struct {
	Event *QRNSNameWrapperExpiryExtended // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperExpiryExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperExpiryExtended)
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
		it.Event = new(QRNSNameWrapperExpiryExtended)
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
func (it *QRNSNameWrapperExpiryExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperExpiryExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperExpiryExtended represents a ExpiryExtended event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperExpiryExtended struct {
	Node   [32]byte
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpiryExtended is a free log retrieval operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterExpiryExtended(opts *bind.FilterOpts, node [][32]byte) (*QRNSNameWrapperExpiryExtendedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "ExpiryExtended", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperExpiryExtendedIterator{contract: _QRNSNameWrapper.contract, event: "ExpiryExtended", logs: logs, sub: sub}, nil
}

// WatchExpiryExtended is a free log subscription operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchExpiryExtended(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperExpiryExtended, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "ExpiryExtended", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperExpiryExtended)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "ExpiryExtended", log); err != nil {
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

// ParseExpiryExtended is a log parse operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseExpiryExtended(log types.Log) (*QRNSNameWrapperExpiryExtended, error) {
	event := new(QRNSNameWrapperExpiryExtended)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "ExpiryExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperFusesSetIterator is returned from FilterFusesSet and is used to iterate over the raw logs and unpacked data for FusesSet events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperFusesSetIterator struct {
	Event *QRNSNameWrapperFusesSet // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperFusesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperFusesSet)
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
		it.Event = new(QRNSNameWrapperFusesSet)
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
func (it *QRNSNameWrapperFusesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperFusesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperFusesSet represents a FusesSet event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperFusesSet struct {
	Node  [32]byte
	Fuses uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFusesSet is a free log retrieval operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterFusesSet(opts *bind.FilterOpts, node [][32]byte) (*QRNSNameWrapperFusesSetIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "FusesSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperFusesSetIterator{contract: _QRNSNameWrapper.contract, event: "FusesSet", logs: logs, sub: sub}, nil
}

// WatchFusesSet is a free log subscription operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchFusesSet(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperFusesSet, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "FusesSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperFusesSet)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "FusesSet", log); err != nil {
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

// ParseFusesSet is a log parse operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseFusesSet(log types.Log) (*QRNSNameWrapperFusesSet, error) {
	event := new(QRNSNameWrapperFusesSet)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "FusesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperNameUnwrappedIterator is returned from FilterNameUnwrapped and is used to iterate over the raw logs and unpacked data for NameUnwrapped events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperNameUnwrappedIterator struct {
	Event *QRNSNameWrapperNameUnwrapped // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperNameUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperNameUnwrapped)
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
		it.Event = new(QRNSNameWrapperNameUnwrapped)
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
func (it *QRNSNameWrapperNameUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperNameUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperNameUnwrapped represents a NameUnwrapped event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperNameUnwrapped struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNameUnwrapped is a free log retrieval operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterNameUnwrapped(opts *bind.FilterOpts, node [][32]byte) (*QRNSNameWrapperNameUnwrappedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "NameUnwrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperNameUnwrappedIterator{contract: _QRNSNameWrapper.contract, event: "NameUnwrapped", logs: logs, sub: sub}, nil
}

// WatchNameUnwrapped is a free log subscription operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchNameUnwrapped(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperNameUnwrapped, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "NameUnwrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperNameUnwrapped)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "NameUnwrapped", log); err != nil {
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

// ParseNameUnwrapped is a log parse operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseNameUnwrapped(log types.Log) (*QRNSNameWrapperNameUnwrapped, error) {
	event := new(QRNSNameWrapperNameUnwrapped)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "NameUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperNameWrappedIterator is returned from FilterNameWrapped and is used to iterate over the raw logs and unpacked data for NameWrapped events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperNameWrappedIterator struct {
	Event *QRNSNameWrapperNameWrapped // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperNameWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperNameWrapped)
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
		it.Event = new(QRNSNameWrapperNameWrapped)
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
func (it *QRNSNameWrapperNameWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperNameWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperNameWrapped represents a NameWrapped event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperNameWrapped struct {
	Node   [32]byte
	Name   []byte
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNameWrapped is a free log retrieval operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterNameWrapped(opts *bind.FilterOpts, node [][32]byte) (*QRNSNameWrapperNameWrappedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "NameWrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperNameWrappedIterator{contract: _QRNSNameWrapper.contract, event: "NameWrapped", logs: logs, sub: sub}, nil
}

// WatchNameWrapped is a free log subscription operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchNameWrapped(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperNameWrapped, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "NameWrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperNameWrapped)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "NameWrapped", log); err != nil {
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

// ParseNameWrapped is a log parse operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseNameWrapped(log types.Log) (*QRNSNameWrapperNameWrapped, error) {
	event := new(QRNSNameWrapperNameWrapped)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "NameWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperOwnershipTransferredIterator struct {
	Event *QRNSNameWrapperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperOwnershipTransferred)
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
		it.Event = new(QRNSNameWrapperOwnershipTransferred)
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
func (it *QRNSNameWrapperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperOwnershipTransferred represents a OwnershipTransferred event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QRNSNameWrapperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperOwnershipTransferredIterator{contract: _QRNSNameWrapper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperOwnershipTransferred)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseOwnershipTransferred(log types.Log) (*QRNSNameWrapperOwnershipTransferred, error) {
	event := new(QRNSNameWrapperOwnershipTransferred)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperTransferBatchIterator struct {
	Event *QRNSNameWrapperTransferBatch // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperTransferBatch)
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
		it.Event = new(QRNSNameWrapperTransferBatch)
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
func (it *QRNSNameWrapperTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperTransferBatch represents a TransferBatch event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*QRNSNameWrapperTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperTransferBatchIterator{contract: _QRNSNameWrapper.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperTransferBatch)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseTransferBatch(log types.Log) (*QRNSNameWrapperTransferBatch, error) {
	event := new(QRNSNameWrapperTransferBatch)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperTransferSingleIterator struct {
	Event *QRNSNameWrapperTransferSingle // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperTransferSingle)
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
		it.Event = new(QRNSNameWrapperTransferSingle)
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
func (it *QRNSNameWrapperTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperTransferSingle represents a TransferSingle event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*QRNSNameWrapperTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperTransferSingleIterator{contract: _QRNSNameWrapper.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperTransferSingle)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseTransferSingle(log types.Log) (*QRNSNameWrapperTransferSingle, error) {
	event := new(QRNSNameWrapperTransferSingle)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QRNSNameWrapperURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the QRNSNameWrapper contract.
type QRNSNameWrapperURIIterator struct {
	Event *QRNSNameWrapperURI // Event containing the contract specifics and raw log

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
func (it *QRNSNameWrapperURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QRNSNameWrapperURI)
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
		it.Event = new(QRNSNameWrapperURI)
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
func (it *QRNSNameWrapperURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QRNSNameWrapperURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QRNSNameWrapperURI represents a URI event raised by the QRNSNameWrapper contract.
type QRNSNameWrapperURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*QRNSNameWrapperURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &QRNSNameWrapperURIIterator{contract: _QRNSNameWrapper.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *QRNSNameWrapperURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _QRNSNameWrapper.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QRNSNameWrapperURI)
				if err := _QRNSNameWrapper.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_QRNSNameWrapper *QRNSNameWrapperFilterer) ParseURI(log types.Log) (*QRNSNameWrapperURI, error) {
	event := new(QRNSNameWrapperURI)
	if err := _QRNSNameWrapper.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
