// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	// _ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TipscontractABI is the input ABI used to generate the binding from.
const TipscontractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"tip\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimTips\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TipscontractBin is the compiled bytecode used for deploying new contracts.
var TipscontractBin = "0x608060405234801561001057600080fd5b506105d5806100206000396000f3fe60806040526004361061003f5760003560e01c806306af3dfd1461004457806350382c1a1461005b57806357b50e8d146100e1578063cb56393c1461017b575b600080fd5b34801561005057600080fd5b506100596101f4565b005b34801561006757600080fd5b506100df6004803603602081101561007e57600080fd5b810190808035906020019064010000000081111561009b57600080fd5b8201836020820111156100ad57600080fd5b803590602001918460018302840111640100000000831117156100cf57600080fd5b9091929391929390505050610313565b005b3480156100ed57600080fd5b506101656004803603602081101561010457600080fd5b810190808035906020019064010000000081111561012157600080fd5b82018360208201111561013357600080fd5b8035906020019184600183028401116401000000008311171561015557600080fd5b9091929391929390505050610409565b6040518082815260200191505060405180910390f35b6101f26004803603602081101561019157600080fd5b81019080803590602001906401000000008111156101ae57600080fd5b8201836020820111156101c057600080fd5b803590602001918460018302840111640100000000831117156101e257600080fd5b909192939192939050505061043d565b005b60003390506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541161024557600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f193505050501580156102ca573d6000803e3d6000fd5b506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6000339050600073ffffffffffffffffffffffffffffffffffffffff1660008484604051808383808284378083019250505092505050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461039b57600080fd5b8060008484604051808383808284378083019250505092505050908152602001604051809103902060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000808383604051808383808284378083019250505092505050908152602001604051809103902060010154905092915050565b6000341161044a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008383604051808383808284378083019250505092505050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156104ce57600080fd5b34600083836040518083838082843780830192505050925050509081526020016040518091039020600101600082825401925050819055503460016000808585604051808383808284378083019250505092505050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505056fea265627a7a72315820379f1562adbe7d8dbca894ff981ffaab54cd47cbab26d59aa407eae917c138cb64736f6c63430005100032"

// DeployTipscontract deploys a new Ethereum contract, binding an instance of Tipscontract to it.
func DeployTipscontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tipscontract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipscontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TipscontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tipscontract{TipscontractCaller: TipscontractCaller{contract: contract}, TipscontractTransactor: TipscontractTransactor{contract: contract}, TipscontractFilterer: TipscontractFilterer{contract: contract}}, nil
}

// Tipscontract is an auto generated Go binding around an Ethereum contract.
type Tipscontract struct {
	TipscontractCaller     // Read-only binding to the contract
	TipscontractTransactor // Write-only binding to the contract
	TipscontractFilterer   // Log filterer for contract events
}

// TipscontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipscontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipscontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipscontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipscontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipscontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipscontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipscontractSession struct {
	Contract     *Tipscontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipscontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipscontractCallerSession struct {
	Contract *TipscontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TipscontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipscontractTransactorSession struct {
	Contract     *TipscontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TipscontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipscontractRaw struct {
	Contract *Tipscontract // Generic contract binding to access the raw methods on
}

// TipscontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipscontractCallerRaw struct {
	Contract *TipscontractCaller // Generic read-only contract binding to access the raw methods on
}

// TipscontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipscontractTransactorRaw struct {
	Contract *TipscontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipscontract creates a new instance of Tipscontract, bound to a specific deployed contract.
func NewTipscontract(address common.Address, backend bind.ContractBackend) (*Tipscontract, error) {
	contract, err := bindTipscontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tipscontract{TipscontractCaller: TipscontractCaller{contract: contract}, TipscontractTransactor: TipscontractTransactor{contract: contract}, TipscontractFilterer: TipscontractFilterer{contract: contract}}, nil
}

// NewTipscontractCaller creates a new read-only instance of Tipscontract, bound to a specific deployed contract.
func NewTipscontractCaller(address common.Address, caller bind.ContractCaller) (*TipscontractCaller, error) {
	contract, err := bindTipscontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipscontractCaller{contract: contract}, nil
}

// NewTipscontractTransactor creates a new write-only instance of Tipscontract, bound to a specific deployed contract.
func NewTipscontractTransactor(address common.Address, transactor bind.ContractTransactor) (*TipscontractTransactor, error) {
	contract, err := bindTipscontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipscontractTransactor{contract: contract}, nil
}

// NewTipscontractFilterer creates a new log filterer instance of Tipscontract, bound to a specific deployed contract.
func NewTipscontractFilterer(address common.Address, filterer bind.ContractFilterer) (*TipscontractFilterer, error) {
	contract, err := bindTipscontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipscontractFilterer{contract: contract}, nil
}

// bindTipscontract binds a generic wrapper to an already deployed contract.
func bindTipscontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipscontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tipscontract *TipscontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tipscontract.Contract.TipscontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tipscontract *TipscontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tipscontract.Contract.TipscontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tipscontract *TipscontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tipscontract.Contract.TipscontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tipscontract *TipscontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tipscontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tipscontract *TipscontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tipscontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tipscontract *TipscontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tipscontract.Contract.contract.Transact(opts, method, params...)
}

// GetTips is a free data retrieval call binding the contract method 0x57b50e8d.
//
// Solidity: function getTips(string key) constant returns(uint256)
func (_Tipscontract *TipscontractCaller) GetTips(opts *bind.CallOpts, key string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tipscontract.contract.Call(opts, out, "getTips", key)
	return *ret0, err
}

// GetTips is a free data retrieval call binding the contract method 0x57b50e8d.
//
// Solidity: function getTips(string key) constant returns(uint256)
func (_Tipscontract *TipscontractSession) GetTips(key string) (*big.Int, error) {
	return _Tipscontract.Contract.GetTips(&_Tipscontract.CallOpts, key)
}

// GetTips is a free data retrieval call binding the contract method 0x57b50e8d.
//
// Solidity: function getTips(string key) constant returns(uint256)
func (_Tipscontract *TipscontractCallerSession) GetTips(key string) (*big.Int, error) {
	return _Tipscontract.Contract.GetTips(&_Tipscontract.CallOpts, key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Tipscontract *TipscontractTransactor) AddKey(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Tipscontract.contract.Transact(opts, "addKey", key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Tipscontract *TipscontractSession) AddKey(key string) (*types.Transaction, error) {
	return _Tipscontract.Contract.AddKey(&_Tipscontract.TransactOpts, key)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string key) returns()
func (_Tipscontract *TipscontractTransactorSession) AddKey(key string) (*types.Transaction, error) {
	return _Tipscontract.Contract.AddKey(&_Tipscontract.TransactOpts, key)
}

// ClaimTips is a paid mutator transaction binding the contract method 0x06af3dfd.
//
// Solidity: function claimTips() returns()
func (_Tipscontract *TipscontractTransactor) ClaimTips(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tipscontract.contract.Transact(opts, "claimTips")
}

// ClaimTips is a paid mutator transaction binding the contract method 0x06af3dfd.
//
// Solidity: function claimTips() returns()
func (_Tipscontract *TipscontractSession) ClaimTips() (*types.Transaction, error) {
	return _Tipscontract.Contract.ClaimTips(&_Tipscontract.TransactOpts)
}

// ClaimTips is a paid mutator transaction binding the contract method 0x06af3dfd.
//
// Solidity: function claimTips() returns()
func (_Tipscontract *TipscontractTransactorSession) ClaimTips() (*types.Transaction, error) {
	return _Tipscontract.Contract.ClaimTips(&_Tipscontract.TransactOpts)
}

// Tip is a paid mutator transaction binding the contract method 0xcb56393c.
//
// Solidity: function tip(string key) returns()
func (_Tipscontract *TipscontractTransactor) Tip(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Tipscontract.contract.Transact(opts, "tip", key)
}

// Tip is a paid mutator transaction binding the contract method 0xcb56393c.
//
// Solidity: function tip(string key) returns()
func (_Tipscontract *TipscontractSession) Tip(key string) (*types.Transaction, error) {
	return _Tipscontract.Contract.Tip(&_Tipscontract.TransactOpts, key)
}

// Tip is a paid mutator transaction binding the contract method 0xcb56393c.
//
// Solidity: function tip(string key) returns()
func (_Tipscontract *TipscontractTransactorSession) Tip(key string) (*types.Transaction, error) {
	return _Tipscontract.Contract.Tip(&_Tipscontract.TransactOpts, key)
}
