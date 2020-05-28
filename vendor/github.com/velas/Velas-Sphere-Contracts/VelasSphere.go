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

// EthdepositcontractABI is the input ABI used to generate the binding from.
const EthdepositcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"banNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"proposePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"depositWithNodes\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pull\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_places\",\"type\":\"uint256\"}],\"name\":\"changePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"staking_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mining_addr\",\"type\":\"address\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keepPerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_writePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_GPUTPerCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CPUTtPerCycle\",\"type\":\"uint256\"}],\"name\":\"changeNodePricing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"old_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"new_addr\",\"type\":\"address\"}],\"name\":\"changeMiningAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthdepositcontractBin is the compiled bytecode used for deploying new contracts.
var EthdepositcontractBin = "0x608060405264174876e8006000556032600155600560065534801561002357600080fd5b5060016007600001819055506001600760010181905550600160076002018190555060016007600301819055506000600760040160006101000a81548160ff021916908315150217905550603e600381905550610e15806100856000396000f3fe6080604052600436106100915760003560e01c806388e66a541161005957806388e66a54146102345780638a99c70a146102a55780638e8b45c614610300578063b095853214610338578063e708013f1461039357610091565b80632790fb771461009b5780633296972f146100e05780633faba59e146101395780634f1c58e31461018a57806351cff8d9146101e3575b6100996103f7565b005b3480156100a757600080fd5b506100de600480360360408110156100be57600080fd5b810190808035906020019092919080359060200190929190505050610507565b005b3480156100ec57600080fd5b506101376004803603608081101561010357600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610584565b005b34801561014557600080fd5b506101886004803603602081101561015c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061061e565b005b34801561019657600080fd5b506101e1600480360360808110156101ad57600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610744565b005b3480156101ef57600080fd5b506102326004803603602081101561020657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107e0565b005b34801561024057600080fd5b506102a36004803603604081101561025757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108ef565b005b3480156102b157600080fd5b506102fe600480360360408110156102c857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061097a565b005b6103366004803603604081101561031657600080fd5b810190808035906020019092919080359060200190929190505050610a31565b005b34801561034457600080fd5b506103916004803603604081101561035b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a62565b005b6103f5600480360360408110156103a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c76565b005b6000600f60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000341161044757600080fd5b348160080160008282540192505081905550600015158160090160009054906101000a900460ff16151514156104e757600781600001600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a900460ff168160040160006101000a81548160ff02191690831515021790555090505060016004600082825401925050819055505b60018160090160006101000a81548160ff02191690831515021790555050565b6000600f60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905081816006016001018190555082816006016000018190555060018160050160006101000a81548160ff021916908315150217905550505050565b6000600f60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905084816000016000018190555083816000016001018190555082816000016002018190555081816000016003018190555060018160000160040160006101000a81548160ff0219169083151502179055505050505050565b6000600f60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600a0160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506001600e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282540192505081905550600654600e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002050505050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600115158160030160009054906101000a900460ff161515146107a957600080fd5b8481600401600001819055508381600401600101819055508281600401600201819055508181600401600301819055505050505050565b6000600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600201541161083457600080fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201549081150290604051600060405180830381858888f193505050501580156108e0573d6000803e3d6000fd5b50600081600201819055505050565b6000600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050818160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146109b257600080fd5b6000601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600e015414610a0657600080fd5b81816002018190555060018160030160006101000a81548160ff021916908315150217905550505050565b610a396103f7565b600082148015610a495750600081145b15610a5357610a5e565b610a5d8282610507565b5b5050565b6000600d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060030160009054906101000a900460ff16610ac057600080fd5b600654600e60008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015410610b3457600080fd5b6000600f60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600a0160008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015414610bec57600080fd5b6000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060030160009054906101000a900460ff16610c4a57600080fd5b600181600e01600082825401925050819055508381600d01600082825401925050819055505050505050565b6000600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600015158160030160009054906101000a900460ff16151514610cdb57600080fd5b60018160030160006101000a81548160ff0219169083151502179055506000543414610d0657600080fd5b828160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d51610d80565b816009016001018190555060055481600901600001819055506001600260008282540192505081905550505050565b600080600c600060055481526020019081526020016000209050605e816001015410610db85760016005600082825401925050819055505b600081600101546001901b90506001826001016000828254019250508190555080925050509056fea265627a7a72315820316bd7c63bee38b0744b825ebd54548facfa551db9083bd9b7c8a36958edce1a64736f6c63430005100032"

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

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) BanNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "banNode", _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// BanNode is a paid mutator transaction binding the contract method 0x3faba59e.
//
// Solidity: function banNode(address _node) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) BanNode(_node common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.BanNode(&_Ethdepositcontract.TransactOpts, _node)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) ChangeMiningAddr(opts *bind.TransactOpts, old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "changeMiningAddr", old_addr, new_addr)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) ChangeMiningAddr(old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeMiningAddr(&_Ethdepositcontract.TransactOpts, old_addr, new_addr)
}

// ChangeMiningAddr is a paid mutator transaction binding the contract method 0x88e66a54.
//
// Solidity: function changeMiningAddr(address old_addr, address new_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) ChangeMiningAddr(old_addr common.Address, new_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.ChangeMiningAddr(&_Ethdepositcontract.TransactOpts, old_addr, new_addr)
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

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) CreateInvoice(opts *bind.TransactOpts, user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "createInvoice", user, price)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractSession) CreateInvoice(user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, user, price)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xb0958532.
//
// Solidity: function createInvoice(address user, uint256 price) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) CreateInvoice(user common.Address, price *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.CreateInvoice(&_Ethdepositcontract.TransactOpts, user, price)
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

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) OpenInvoice(opts *bind.TransactOpts, addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "openInvoice", addr, deadline)
}

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractSession) OpenInvoice(addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.OpenInvoice(&_Ethdepositcontract.TransactOpts, addr, deadline)
}

// OpenInvoice is a paid mutator transaction binding the contract method 0x8a99c70a.
//
// Solidity: function openInvoice(address addr, uint256 deadline) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) OpenInvoice(addr common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.OpenInvoice(&_Ethdepositcontract.TransactOpts, addr, deadline)
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

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactor) RegisterNode(opts *bind.TransactOpts, staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.contract.Transact(opts, "registerNode", staking_addr, mining_addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) returns()
func (_Ethdepositcontract *EthdepositcontractSession) RegisterNode(staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, staking_addr, mining_addr)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xe708013f.
//
// Solidity: function registerNode(address staking_addr, address mining_addr) returns()
func (_Ethdepositcontract *EthdepositcontractTransactorSession) RegisterNode(staking_addr common.Address, mining_addr common.Address) (*types.Transaction, error) {
	return _Ethdepositcontract.Contract.RegisterNode(&_Ethdepositcontract.TransactOpts, staking_addr, mining_addr)
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
