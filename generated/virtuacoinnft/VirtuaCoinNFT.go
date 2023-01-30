// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package virtuacoinnft

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

// VirtuacoinnftMetaData contains all meta data concerning the Virtuacoinnft contract.
var VirtuacoinnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metaDataURI\",\"type\":\"string\"}],\"name\":\"AssetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VIRTUACOIN_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VIRTUACOIN_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VIRTUACOIN_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"createAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"delegateAssetCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VirtuacoinnftABI is the input ABI used to generate the binding from.
// Deprecated: Use VirtuacoinnftMetaData.ABI instead.
var VirtuacoinnftABI = VirtuacoinnftMetaData.ABI

// Virtuacoinnft is an auto generated Go binding around an Ethereum contract.
type Virtuacoinnft struct {
	VirtuacoinnftCaller     // Read-only binding to the contract
	VirtuacoinnftTransactor // Write-only binding to the contract
	VirtuacoinnftFilterer   // Log filterer for contract events
}

// VirtuacoinnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type VirtuacoinnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtuacoinnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VirtuacoinnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtuacoinnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VirtuacoinnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtuacoinnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VirtuacoinnftSession struct {
	Contract     *Virtuacoinnft    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VirtuacoinnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VirtuacoinnftCallerSession struct {
	Contract *VirtuacoinnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VirtuacoinnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VirtuacoinnftTransactorSession struct {
	Contract     *VirtuacoinnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VirtuacoinnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type VirtuacoinnftRaw struct {
	Contract *Virtuacoinnft // Generic contract binding to access the raw methods on
}

// VirtuacoinnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VirtuacoinnftCallerRaw struct {
	Contract *VirtuacoinnftCaller // Generic read-only contract binding to access the raw methods on
}

// VirtuacoinnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VirtuacoinnftTransactorRaw struct {
	Contract *VirtuacoinnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVirtuacoinnft creates a new instance of Virtuacoinnft, bound to a specific deployed contract.
func NewVirtuacoinnft(address common.Address, backend bind.ContractBackend) (*Virtuacoinnft, error) {
	contract, err := bindVirtuacoinnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Virtuacoinnft{VirtuacoinnftCaller: VirtuacoinnftCaller{contract: contract}, VirtuacoinnftTransactor: VirtuacoinnftTransactor{contract: contract}, VirtuacoinnftFilterer: VirtuacoinnftFilterer{contract: contract}}, nil
}

// NewVirtuacoinnftCaller creates a new read-only instance of Virtuacoinnft, bound to a specific deployed contract.
func NewVirtuacoinnftCaller(address common.Address, caller bind.ContractCaller) (*VirtuacoinnftCaller, error) {
	contract, err := bindVirtuacoinnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftCaller{contract: contract}, nil
}

// NewVirtuacoinnftTransactor creates a new write-only instance of Virtuacoinnft, bound to a specific deployed contract.
func NewVirtuacoinnftTransactor(address common.Address, transactor bind.ContractTransactor) (*VirtuacoinnftTransactor, error) {
	contract, err := bindVirtuacoinnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftTransactor{contract: contract}, nil
}

// NewVirtuacoinnftFilterer creates a new log filterer instance of Virtuacoinnft, bound to a specific deployed contract.
func NewVirtuacoinnftFilterer(address common.Address, filterer bind.ContractFilterer) (*VirtuacoinnftFilterer, error) {
	contract, err := bindVirtuacoinnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftFilterer{contract: contract}, nil
}

// bindVirtuacoinnft binds a generic wrapper to an already deployed contract.
func bindVirtuacoinnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VirtuacoinnftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Virtuacoinnft *VirtuacoinnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Virtuacoinnft.Contract.VirtuacoinnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Virtuacoinnft *VirtuacoinnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.VirtuacoinnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Virtuacoinnft *VirtuacoinnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.VirtuacoinnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Virtuacoinnft *VirtuacoinnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Virtuacoinnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Virtuacoinnft *VirtuacoinnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Virtuacoinnft *VirtuacoinnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.DEFAULTADMINROLE(&_Virtuacoinnft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.DEFAULTADMINROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINADMINROLE is a free data retrieval call binding the contract method 0x81628385.
//
// Solidity: function VIRTUACOIN_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCaller) VIRTUACOINADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "VIRTUACOIN_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VIRTUACOINADMINROLE is a free data retrieval call binding the contract method 0x81628385.
//
// Solidity: function VIRTUACOIN_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftSession) VIRTUACOINADMINROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINADMINROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINADMINROLE is a free data retrieval call binding the contract method 0x81628385.
//
// Solidity: function VIRTUACOIN_ADMIN_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) VIRTUACOINADMINROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINADMINROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINCREATORROLE is a free data retrieval call binding the contract method 0x535d98be.
//
// Solidity: function VIRTUACOIN_CREATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCaller) VIRTUACOINCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "VIRTUACOIN_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VIRTUACOINCREATORROLE is a free data retrieval call binding the contract method 0x535d98be.
//
// Solidity: function VIRTUACOIN_CREATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftSession) VIRTUACOINCREATORROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINCREATORROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINCREATORROLE is a free data retrieval call binding the contract method 0x535d98be.
//
// Solidity: function VIRTUACOIN_CREATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) VIRTUACOINCREATORROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINCREATORROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINOPERATORROLE is a free data retrieval call binding the contract method 0x9d5519d6.
//
// Solidity: function VIRTUACOIN_OPERATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCaller) VIRTUACOINOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "VIRTUACOIN_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VIRTUACOINOPERATORROLE is a free data retrieval call binding the contract method 0x9d5519d6.
//
// Solidity: function VIRTUACOIN_OPERATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftSession) VIRTUACOINOPERATORROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINOPERATORROLE(&_Virtuacoinnft.CallOpts)
}

// VIRTUACOINOPERATORROLE is a free data retrieval call binding the contract method 0x9d5519d6.
//
// Solidity: function VIRTUACOIN_OPERATOR_ROLE() view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) VIRTUACOINOPERATORROLE() ([32]byte, error) {
	return _Virtuacoinnft.Contract.VIRTUACOINOPERATORROLE(&_Virtuacoinnft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Virtuacoinnft.Contract.BalanceOf(&_Virtuacoinnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Virtuacoinnft.Contract.BalanceOf(&_Virtuacoinnft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.GetApproved(&_Virtuacoinnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.GetApproved(&_Virtuacoinnft.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Virtuacoinnft.Contract.GetRoleAdmin(&_Virtuacoinnft.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Virtuacoinnft.Contract.GetRoleAdmin(&_Virtuacoinnft.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.GetRoleMember(&_Virtuacoinnft.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.GetRoleMember(&_Virtuacoinnft.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Virtuacoinnft.Contract.GetRoleMemberCount(&_Virtuacoinnft.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Virtuacoinnft.Contract.GetRoleMemberCount(&_Virtuacoinnft.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Virtuacoinnft.Contract.HasRole(&_Virtuacoinnft.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Virtuacoinnft.Contract.HasRole(&_Virtuacoinnft.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Virtuacoinnft.Contract.IsApprovedForAll(&_Virtuacoinnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Virtuacoinnft.Contract.IsApprovedForAll(&_Virtuacoinnft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftSession) Name() (string, error) {
	return _Virtuacoinnft.Contract.Name(&_Virtuacoinnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) Name() (string, error) {
	return _Virtuacoinnft.Contract.Name(&_Virtuacoinnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.OwnerOf(&_Virtuacoinnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Virtuacoinnft.Contract.OwnerOf(&_Virtuacoinnft.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftSession) Paused() (bool, error) {
	return _Virtuacoinnft.Contract.Paused(&_Virtuacoinnft.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) Paused() (bool, error) {
	return _Virtuacoinnft.Contract.Paused(&_Virtuacoinnft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Virtuacoinnft.Contract.SupportsInterface(&_Virtuacoinnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Virtuacoinnft.Contract.SupportsInterface(&_Virtuacoinnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftSession) Symbol() (string, error) {
	return _Virtuacoinnft.Contract.Symbol(&_Virtuacoinnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) Symbol() (string, error) {
	return _Virtuacoinnft.Contract.Symbol(&_Virtuacoinnft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Virtuacoinnft.Contract.TokenByIndex(&_Virtuacoinnft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Virtuacoinnft.Contract.TokenByIndex(&_Virtuacoinnft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Virtuacoinnft.Contract.TokenOfOwnerByIndex(&_Virtuacoinnft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Virtuacoinnft.Contract.TokenOfOwnerByIndex(&_Virtuacoinnft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Virtuacoinnft *VirtuacoinnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Virtuacoinnft.Contract.TokenURI(&_Virtuacoinnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Virtuacoinnft.Contract.TokenURI(&_Virtuacoinnft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Virtuacoinnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) TotalSupply() (*big.Int, error) {
	return _Virtuacoinnft.Contract.TotalSupply(&_Virtuacoinnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Virtuacoinnft.Contract.TotalSupply(&_Virtuacoinnft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Approve(&_Virtuacoinnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Approve(&_Virtuacoinnft.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Burn(&_Virtuacoinnft.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Burn(&_Virtuacoinnft.TransactOpts, tokenId)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x55b1f24e.
//
// Solidity: function createAsset(string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftTransactor) CreateAsset(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "createAsset", metadataURI)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x55b1f24e.
//
// Solidity: function createAsset(string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) CreateAsset(metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.CreateAsset(&_Virtuacoinnft.TransactOpts, metadataURI)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x55b1f24e.
//
// Solidity: function createAsset(string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) CreateAsset(metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.CreateAsset(&_Virtuacoinnft.TransactOpts, metadataURI)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x03e30209.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftTransactor) DelegateAssetCreation(opts *bind.TransactOpts, creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "delegateAssetCreation", creator, metadataURI)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x03e30209.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftSession) DelegateAssetCreation(creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.DelegateAssetCreation(&_Virtuacoinnft.TransactOpts, creator, metadataURI)
}

// DelegateAssetCreation is a paid mutator transaction binding the contract method 0x03e30209.
//
// Solidity: function delegateAssetCreation(address creator, string metadataURI) returns(uint256)
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) DelegateAssetCreation(creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.DelegateAssetCreation(&_Virtuacoinnft.TransactOpts, creator, metadataURI)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.GrantRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.GrantRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Virtuacoinnft *VirtuacoinnftSession) Pause() (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Pause(&_Virtuacoinnft.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) Pause() (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Pause(&_Virtuacoinnft.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.RenounceRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.RenounceRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.RevokeRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.RevokeRole(&_Virtuacoinnft.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SafeTransferFrom(&_Virtuacoinnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SafeTransferFrom(&_Virtuacoinnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SafeTransferFrom0(&_Virtuacoinnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SafeTransferFrom0(&_Virtuacoinnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SetApprovalForAll(&_Virtuacoinnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.SetApprovalForAll(&_Virtuacoinnft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.TransferFrom(&_Virtuacoinnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.TransferFrom(&_Virtuacoinnft.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Virtuacoinnft *VirtuacoinnftTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Virtuacoinnft.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Virtuacoinnft *VirtuacoinnftSession) Unpause() (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Unpause(&_Virtuacoinnft.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Virtuacoinnft *VirtuacoinnftTransactorSession) Unpause() (*types.Transaction, error) {
	return _Virtuacoinnft.Contract.Unpause(&_Virtuacoinnft.TransactOpts)
}

// VirtuacoinnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Virtuacoinnft contract.
type VirtuacoinnftApprovalIterator struct {
	Event *VirtuacoinnftApproval // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftApproval)
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
		it.Event = new(VirtuacoinnftApproval)
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
func (it *VirtuacoinnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftApproval represents a Approval event raised by the Virtuacoinnft contract.
type VirtuacoinnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VirtuacoinnftApprovalIterator, error) {

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

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftApprovalIterator{contract: _Virtuacoinnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftApproval)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseApproval(log types.Log) (*VirtuacoinnftApproval, error) {
	event := new(VirtuacoinnftApproval)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Virtuacoinnft contract.
type VirtuacoinnftApprovalForAllIterator struct {
	Event *VirtuacoinnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftApprovalForAll)
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
		it.Event = new(VirtuacoinnftApprovalForAll)
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
func (it *VirtuacoinnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftApprovalForAll represents a ApprovalForAll event raised by the Virtuacoinnft contract.
type VirtuacoinnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VirtuacoinnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftApprovalForAllIterator{contract: _Virtuacoinnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftApprovalForAll)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseApprovalForAll(log types.Log) (*VirtuacoinnftApprovalForAll, error) {
	event := new(VirtuacoinnftApprovalForAll)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftAssetCreatedIterator is returned from FilterAssetCreated and is used to iterate over the raw logs and unpacked data for AssetCreated events raised by the Virtuacoinnft contract.
type VirtuacoinnftAssetCreatedIterator struct {
	Event *VirtuacoinnftAssetCreated // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftAssetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftAssetCreated)
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
		it.Event = new(VirtuacoinnftAssetCreated)
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
func (it *VirtuacoinnftAssetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftAssetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftAssetCreated represents a AssetCreated event raised by the Virtuacoinnft contract.
type VirtuacoinnftAssetCreated struct {
	TokenID     *big.Int
	Creator     common.Address
	MetaDataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetCreated is a free log retrieval operation binding the contract event 0x49b4e4c6879cdc85b51d8c3b572252aaa7842edb863810854acb6d2ed420d962.
//
// Solidity: event AssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterAssetCreated(opts *bind.FilterOpts, creator []common.Address) (*VirtuacoinnftAssetCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "AssetCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftAssetCreatedIterator{contract: _Virtuacoinnft.contract, event: "AssetCreated", logs: logs, sub: sub}, nil
}

// WatchAssetCreated is a free log subscription operation binding the contract event 0x49b4e4c6879cdc85b51d8c3b572252aaa7842edb863810854acb6d2ed420d962.
//
// Solidity: event AssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchAssetCreated(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftAssetCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "AssetCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftAssetCreated)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "AssetCreated", log); err != nil {
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

// ParseAssetCreated is a log parse operation binding the contract event 0x49b4e4c6879cdc85b51d8c3b572252aaa7842edb863810854acb6d2ed420d962.
//
// Solidity: event AssetCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseAssetCreated(log types.Log) (*VirtuacoinnftAssetCreated, error) {
	event := new(VirtuacoinnftAssetCreated)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "AssetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Virtuacoinnft contract.
type VirtuacoinnftPausedIterator struct {
	Event *VirtuacoinnftPaused // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftPaused)
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
		it.Event = new(VirtuacoinnftPaused)
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
func (it *VirtuacoinnftPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftPaused represents a Paused event raised by the Virtuacoinnft contract.
type VirtuacoinnftPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterPaused(opts *bind.FilterOpts) (*VirtuacoinnftPausedIterator, error) {

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftPausedIterator{contract: _Virtuacoinnft.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftPaused) (event.Subscription, error) {

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftPaused)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParsePaused(log types.Log) (*VirtuacoinnftPaused, error) {
	event := new(VirtuacoinnftPaused)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleAdminChangedIterator struct {
	Event *VirtuacoinnftRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftRoleAdminChanged)
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
		it.Event = new(VirtuacoinnftRoleAdminChanged)
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
func (it *VirtuacoinnftRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftRoleAdminChanged represents a RoleAdminChanged event raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VirtuacoinnftRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftRoleAdminChangedIterator{contract: _Virtuacoinnft.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftRoleAdminChanged)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseRoleAdminChanged(log types.Log) (*VirtuacoinnftRoleAdminChanged, error) {
	event := new(VirtuacoinnftRoleAdminChanged)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleGrantedIterator struct {
	Event *VirtuacoinnftRoleGranted // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftRoleGranted)
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
		it.Event = new(VirtuacoinnftRoleGranted)
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
func (it *VirtuacoinnftRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftRoleGranted represents a RoleGranted event raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VirtuacoinnftRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftRoleGrantedIterator{contract: _Virtuacoinnft.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftRoleGranted)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseRoleGranted(log types.Log) (*VirtuacoinnftRoleGranted, error) {
	event := new(VirtuacoinnftRoleGranted)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleRevokedIterator struct {
	Event *VirtuacoinnftRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftRoleRevoked)
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
		it.Event = new(VirtuacoinnftRoleRevoked)
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
func (it *VirtuacoinnftRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftRoleRevoked represents a RoleRevoked event raised by the Virtuacoinnft contract.
type VirtuacoinnftRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VirtuacoinnftRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftRoleRevokedIterator{contract: _Virtuacoinnft.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftRoleRevoked)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseRoleRevoked(log types.Log) (*VirtuacoinnftRoleRevoked, error) {
	event := new(VirtuacoinnftRoleRevoked)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Virtuacoinnft contract.
type VirtuacoinnftTransferIterator struct {
	Event *VirtuacoinnftTransfer // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftTransfer)
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
		it.Event = new(VirtuacoinnftTransfer)
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
func (it *VirtuacoinnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftTransfer represents a Transfer event raised by the Virtuacoinnft contract.
type VirtuacoinnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VirtuacoinnftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftTransferIterator{contract: _Virtuacoinnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftTransfer)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseTransfer(log types.Log) (*VirtuacoinnftTransfer, error) {
	event := new(VirtuacoinnftTransfer)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtuacoinnftUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Virtuacoinnft contract.
type VirtuacoinnftUnpausedIterator struct {
	Event *VirtuacoinnftUnpaused // Event containing the contract specifics and raw log

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
func (it *VirtuacoinnftUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtuacoinnftUnpaused)
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
		it.Event = new(VirtuacoinnftUnpaused)
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
func (it *VirtuacoinnftUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtuacoinnftUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtuacoinnftUnpaused represents a Unpaused event raised by the Virtuacoinnft contract.
type VirtuacoinnftUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VirtuacoinnftUnpausedIterator, error) {

	logs, sub, err := _Virtuacoinnft.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VirtuacoinnftUnpausedIterator{contract: _Virtuacoinnft.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VirtuacoinnftUnpaused) (event.Subscription, error) {

	logs, sub, err := _Virtuacoinnft.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtuacoinnftUnpaused)
				if err := _Virtuacoinnft.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Virtuacoinnft *VirtuacoinnftFilterer) ParseUnpaused(log types.Log) (*VirtuacoinnftUnpaused, error) {
	event := new(VirtuacoinnftUnpaused)
	if err := _Virtuacoinnft.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
