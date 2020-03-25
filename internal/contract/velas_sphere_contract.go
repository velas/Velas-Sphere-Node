// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthdepositcontractABI is the input ABI used to generate the binding from.
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"changePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height_end\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"changeNodePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e800600055603260015534801561001e57600080fd5b5060016006600001819055506001600660010181905550600160066002018190555060016006600301819055506000600660040160006101000a81548160ff021916908315150217905550603e600381905550610db8806100806000396000f3fe6080604052600436106100705760003560e01c80634f1c58e31161004e5780634f1c58e31461018457806351cff8d9146101dd5780638e8b45c61461022e57806396ea86301461026657610070565b80632790fb771461007a5780632eefc412146100bf5780633296972f1461012b575b6100786102f3565b005b34801561008657600080fd5b506100bd6004803603604081101561009d57600080fd5b810190808035906020019092919080359060200190929190505050610403565b005b610129600480360360a08110156100d557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610480565b005b34801561013757600080fd5b506101826004803603608081101561014e57600080fd5b810190808035906020019092919080359060200190929190803590602001909291908035906020019092919050505061057a565b005b34801561019057600080fd5b506101db600480360360808110156101a757600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610614565b005b3480156101e957600080fd5b5061022c6004803603602081101561020057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b0565b005b6102646004803603604081101561024457600080fd5b81019080803590602001909291908035906020019092919050505061079b565b005b34801561027257600080fd5b506102f1600480360360e081101561028957600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506107cc565b005b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000341161034357600080fd5b348160080160008282540192505081905550600015158160090160009054906101000a900460ff16151514156103e357600681600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555090505060016004600082825401925050819055505b60018160090160006101000a81548160ff02191690831515021790555050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905081816006016001018190555082816006016000018190555060018160050160006101000a81548160ff021916908315150217905550505050565b6000600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600015158160020160009054906101000a900460ff161515146104e557600080fd5b60018160020160006101000a81548160ff021916908315150217905550600054341461051057600080fd5b610518610b59565b816008016001018190555060055481600801600001819055508481600301600001819055508381600301600101819055508281600301600201819055508181600301600301819055506001600260008282540192505081905550505050505050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905084816000016000018190555083816000016001018190555082816000016002018190555081816000016003018190555060018160000160040160006101000a81548160ff0219169083151502179055505050505050565b6000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160020160009054906101000a900460ff1615151461067957600080fd5b8481600301600001819055508381600301600101819055508281600301600201819055508181600301600301819055505050505050565b6000600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161070457600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc600c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549081150290604051600060405180830381858888f1935050505015801561078c573d6000803e3d6000fd5b50600081600101819055505050565b6107a36102f3565b6000821480156107b35750600081145b156107bd576107c8565b6107c78282610403565b5b5050565b6000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160009054906101000a900460ff1661082a57600080fd5b6000600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160050160009054906101000a900460ff16151514156108a65781600801600001548160060160000154146108a557600080fd5b5b6000600e60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600d0154141561098057898160000181905550888160010181905550878160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508681600801600001819055508581600801600101819055508481600801600201819055508381600801600301819055506109ed565b8981600001541461099057600080fd5b888160010154146109a057600080fd5b868160080160000154146109b357600080fd5b858160080160010154146109c657600080fd5b848160080160020154146109d957600080fd5b838160080160030154146109ec57600080fd5b5b600181600d01600082825401925050819055506000610a128360000189898989610bb9565b90506001548a014310610b335760035482600d015410610a3f57610a368982610bf5565b50505050610b50565b600e60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d820160009055505050505050610b50565b605e82600d01541415610b4b57610b4a8982610bf5565b5b505050505b50505050505050565b600080600b600060055481526020019081526020016000209050605e816001015410610b915760016005600082825401925050819055505b600081600101546001901b905060018260010160008282540192505081905550809250505090565b60008085876000015402810190508487600101540281019050828760030154028101905083876002015402810190508091505095945050505050565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080154811115610c4457600080fd5b80600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080160008282540392505081905550600e60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382016000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff02191690555050600882016000808201600090556001820160009055600282016000905560038201600090555050600c820160009055600d8201600090555050505056fea265627a7a72315820ac2e8b2a7091cb55b5c52a32c7c9ab1ac0f97c99d7a5a8a0e2a7b3b0ee1b13c664736f6c63430005100032"

// DeployEthdepositcontract deploys a new Ethereum contract, binding an instance of Ethdepositcontract to it.
func DeployEthdepositcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethdepositcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthdepositcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// Ethdepositcontract is an auto generated Go binding around an Ethereum contract.
type Ethdepositcontract struct {
	EthdepositcontractCaller     // Read-only binding to the contract
	EthdepositcontractTransactor // Write-only binding to the contract
	EthdepositcontractFilterer   // Log filterer for contract events
}

// EthdepositcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthdepositcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthdepositcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthdepositcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthdepositcontractSession struct {
	Contract     *Ethdepositcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthdepositcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthdepositcontractCallerSession struct {
	Contract *EthdepositcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EthdepositcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthdepositcontractTransactorSession struct {
	Contract     *EthdepositcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EthdepositcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthdepositcontractRaw struct {
	Contract *Ethdepositcontract // Generic contract binding to access the raw methods on
}

// EthdepositcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthdepositcontractCallerRaw struct {
	Contract *EthdepositcontractCaller // Generic read-only contract binding to access the raw methods on
}

// EthdepositcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthdepositcontractTransactorRaw struct {
	Contract *EthdepositcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthdepositcontract creates a new instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontract(address common.Address, backend bind.ContractBackend) (*Ethdepositcontract, error) {
	contract, err := bindEthdepositcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethdepositcontract{EthdepositcontractCaller: EthdepositcontractCaller{contract: contract}, EthdepositcontractTransactor: EthdepositcontractTransactor{contract: contract}, EthdepositcontractFilterer: EthdepositcontractFilterer{contract: contract}}, nil
}

// NewEthdepositcontractCaller creates a new read-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractCaller(address common.Address, caller bind.ContractCaller) (*EthdepositcontractCaller, error) {
	contract, err := bindEthdepositcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractCaller{contract: contract}, nil
}

// NewEthdepositcontractTransactor creates a new write-only instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*EthdepositcontractTransactor, error) {
	contract, err := bindEthdepositcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractTransactor{contract: contract}, nil
}

// NewEthdepositcontractFilterer creates a new log filterer instance of Ethdepositcontract, bound to a specific deployed contract.
func NewEthdepositcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*EthdepositcontractFilterer, error) {
	contract, err := bindEthdepositcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthdepositcontractFilterer{contract: contract}, nil
}

// bindEthdepositcontract binds a generic wrapper to an already deployed contract.
func bindEthdepositcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthdepositcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.EthdepositcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.EthdepositcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethdepositcontract *EthdepositcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethdepositcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethdepositcontract *EthdepositcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.contract.Transact(opts, method, params...)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangeNodePricing(opts *bind.TransactOpts, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changeNodePricing", _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangeNodePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeNodePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangeNodePricing is a paid mutator transaction binding the contract method 0x4f1c58e3.
//
// Solidity: function changeNodePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangeNodePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeNodePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangePool(opts *bind.TransactOpts, _pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changePool", _pull, _places)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangePool(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangePool(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// ChangePool is a paid mutator transaction binding the contract method 0x2790fb77.
//
// Solidity: function changePool(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangePool(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangePool(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x96ea8630.
//
// Solidity: function createInvoice(uint256 height_start, uint256 height_end, address user, uint256 keepPerByte, uint256 writePerByte, uint256 GPUTPerCycle, uint256 CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(height_start *big.Int, height_end *big.Int, user common.Address, keepPerByte *big.Int, writePerByte *big.Int, GPUTPerCycle *big.Int, CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, height_start, height_end, user, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) DepositWithNodes(opts *bind.TransactOpts, _pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "depositWithNodes", _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// DepositWithNodes is a paid mutator transaction binding the contract method 0x8e8b45c6.
//
// Solidity: function depositWithNodes(uint256 _pull, uint256 _places) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) DepositWithNodes(_pull *big.Int, _places *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.DepositWithNodes(&_Ethdepositcontract.TransactOpts, _pull, _places)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ProposePricing(opts *bind.TransactOpts, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "proposePricing", _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// ProposePricing is a paid mutator transaction binding the contract method 0x3296972f.
//
// Solidity: function proposePricing(uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ProposePricing(_keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ProposePricing(&_Ethdepositcontract.TransactOpts, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterNode(opts *bind.TransactOpts, addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerNode", addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterNode(addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x2eefc412.
//
// Solidity: function registerNode(address addr, uint256 _keepPerByte, uint256 _writePerByte, uint256 _GPUTPerCycle, uint256 _CPUTtPerCycle) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterNode(addr common.Address, _keepPerByte *big.Int, _writePerByte *big.Int, _GPUTPerCycle *big.Int, _CPUTtPerCycle *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, addr, _keepPerByte, _writePerByte, _GPUTPerCycle, _CPUTtPerCycle)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) Withdraw(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "withdraw", addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) Withdraw(addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.Withdraw(&_Ethdepositcontract.TransactOpts, addr)
}
