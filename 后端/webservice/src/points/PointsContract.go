// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package points

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// EshouPointABI is the input ABI used to generate the binding from.
const EshouPointABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceReduced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAccountExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reduceBalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EshouPointBin is the compiled bytecode used for deploying new contracts.
var EshouPointBin = "0x608060405234801561001057600080fd5b506000600181905550610cec806100286000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b47279061161005b578063b47279061461013b578063c496a4e11461016b578063e89b0e1e1461019b578063f8b2cb4f146101b757610088565b806314b4a4f91461008d57806321e5383a146100bd57806327e235e3146100ed57806353c004241461011d575b600080fd5b6100a760048036038101906100a2919061086a565b6101e7565b6040516100b491906108b2565b60405180910390f35b6100d760048036038101906100d29190610903565b610293565b6040516100e491906108b2565b60405180910390f35b6101076004803603810190610102919061086a565b610454565b6040516101149190610952565b60405180910390f35b61012561046c565b6040516101329190610952565b60405180910390f35b61015560048036038101906101509190610903565b610472565b60405161016291906108b2565b60405180910390f35b6101856004803603810190610180919061096d565b6105ea565b60405161019291906109a9565b60405180910390f35b6101b560048036038101906101b0919061086a565b610629565b005b6101d160048036038101906101cc919061086a565b610777565b6040516101de9190610952565b60405180910390f35b600080600090505b600154811015610288578273ffffffffffffffffffffffffffffffffffffffff1660028281548110610224576102236109c4565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561027557600191505061028e565b808061028090610a22565b9150506101ef565b50600090505b919050565b600061029e836101e7565b6102dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d490610ac8565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546103669190610ae8565b10156103a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039e90610b8a565b60405180910390fd5b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103f59190610ae8565b925050819055508273ffffffffffffffffffffffffffffffffffffffff167fe96dd7a15a3974e8e7d5eb80de2cb6fd69907a1ee089170cd976cfffaedfa8ac836040516104429190610952565b60405180910390a26001905092915050565b60006020528060005260406000206000915090505481565b60015481565b600061047d836101e7565b6104bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b390610ac8565b60405180910390fd5b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561053d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053490610bf6565b60405180910390fd5b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461058b9190610c16565b925050819055508273ffffffffffffffffffffffffffffffffffffffff167f2e9bf8d4a8402929da26de77a79494626b184ddae2e3e0c076d6dfa10cd2a1d9836040516105d89190610952565b60405180910390a26001905092915050565b600281815481106105fa57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610632816101e7565b15610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066990610c96565b60405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160008154809291906106c990610a22565b91905055506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f8f42195a0bbfa58954be4349deb9efc38bdb9c298e529f705f8bc1e38bce039960405160405180910390a250565b6000610782826101e7565b6107c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b890610ac8565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108378261080c565b9050919050565b6108478161082c565b811461085257600080fd5b50565b6000813590506108648161083e565b92915050565b6000602082840312156108805761087f610807565b5b600061088e84828501610855565b91505092915050565b60008115159050919050565b6108ac81610897565b82525050565b60006020820190506108c760008301846108a3565b92915050565b6000819050919050565b6108e0816108cd565b81146108eb57600080fd5b50565b6000813590506108fd816108d7565b92915050565b6000806040838503121561091a57610919610807565b5b600061092885828601610855565b9250506020610939858286016108ee565b9150509250929050565b61094c816108cd565b82525050565b60006020820190506109676000830184610943565b92915050565b60006020828403121561098357610982610807565b5b6000610991848285016108ee565b91505092915050565b6109a38161082c565b82525050565b60006020820190506109be600083018461099a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610a2d826108cd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610a6057610a5f6109f3565b5b600182019050919050565b600082825260208201905092915050565b7f4163636f756e7420646f6573206e6f7420657869737400000000000000000000600082015250565b6000610ab2601683610a6b565b9150610abd82610a7c565b602082019050919050565b60006020820190508181036000830152610ae181610aa5565b9050919050565b6000610af3826108cd565b9150610afe836108cd565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610b3357610b326109f3565b5b828201905092915050565b7f496e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b6000610b74600e83610a6b565b9150610b7f82610b3e565b602082019050919050565b60006020820190508181036000830152610ba381610b67565b9050919050565b7f496e73756666696369656e742062616c616e6365000000000000000000000000600082015250565b6000610be0601483610a6b565b9150610beb82610baa565b602082019050919050565b60006020820190508181036000830152610c0f81610bd3565b9050919050565b6000610c21826108cd565b9150610c2c836108cd565b925082821015610c3f57610c3e6109f3565b5b828203905092915050565b7f4163636f756e7420616c72656164792065786973747300000000000000000000600082015250565b6000610c80601683610a6b565b9150610c8b82610c4a565b602082019050919050565b60006020820190508181036000830152610caf81610c73565b905091905056fea2646970667358221220791f11316b5ead1060b167ee125d7cbeb5dd680da080df1cd1c5f624f0d195c764736f6c634300080b0033"

// DeployEshouPoint deploys a new contract, binding an instance of EshouPoint to it.
func DeployEshouPoint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *EshouPoint, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouPointABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EshouPointBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &EshouPoint{EshouPointCaller: EshouPointCaller{contract: contract}, EshouPointTransactor: EshouPointTransactor{contract: contract}, EshouPointFilterer: EshouPointFilterer{contract: contract}}, nil
}

func AsyncDeployEshouPoint(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouPointABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(EshouPointBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// EshouPoint is an auto generated Go binding around a Solidity contract.
type EshouPoint struct {
	EshouPointCaller     // Read-only binding to the contract
	EshouPointTransactor // Write-only binding to the contract
	EshouPointFilterer   // Log filterer for contract events
}

// EshouPointCaller is an auto generated read-only Go binding around a Solidity contract.
type EshouPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouPointTransactor is an auto generated write-only Go binding around a Solidity contract.
type EshouPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouPointFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EshouPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EshouPointSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EshouPointSession struct {
	Contract     *EshouPoint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EshouPointCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EshouPointCallerSession struct {
	Contract *EshouPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EshouPointTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EshouPointTransactorSession struct {
	Contract     *EshouPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EshouPointRaw is an auto generated low-level Go binding around a Solidity contract.
type EshouPointRaw struct {
	Contract *EshouPoint // Generic contract binding to access the raw methods on
}

// EshouPointCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EshouPointCallerRaw struct {
	Contract *EshouPointCaller // Generic read-only contract binding to access the raw methods on
}

// EshouPointTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EshouPointTransactorRaw struct {
	Contract *EshouPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEshouPoint creates a new instance of EshouPoint, bound to a specific deployed contract.
func NewEshouPoint(address common.Address, backend bind.ContractBackend) (*EshouPoint, error) {
	contract, err := bindEshouPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EshouPoint{EshouPointCaller: EshouPointCaller{contract: contract}, EshouPointTransactor: EshouPointTransactor{contract: contract}, EshouPointFilterer: EshouPointFilterer{contract: contract}}, nil
}

// NewEshouPointCaller creates a new read-only instance of EshouPoint, bound to a specific deployed contract.
func NewEshouPointCaller(address common.Address, caller bind.ContractCaller) (*EshouPointCaller, error) {
	contract, err := bindEshouPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EshouPointCaller{contract: contract}, nil
}

// NewEshouPointTransactor creates a new write-only instance of EshouPoint, bound to a specific deployed contract.
func NewEshouPointTransactor(address common.Address, transactor bind.ContractTransactor) (*EshouPointTransactor, error) {
	contract, err := bindEshouPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EshouPointTransactor{contract: contract}, nil
}

// NewEshouPointFilterer creates a new log filterer instance of EshouPoint, bound to a specific deployed contract.
func NewEshouPointFilterer(address common.Address, filterer bind.ContractFilterer) (*EshouPointFilterer, error) {
	contract, err := bindEshouPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EshouPointFilterer{contract: contract}, nil
}

// bindEshouPoint binds a generic wrapper to an already deployed contract.
func bindEshouPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EshouPointABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EshouPoint *EshouPointRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EshouPoint.Contract.EshouPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EshouPoint *EshouPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.EshouPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EshouPoint *EshouPointRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.EshouPointTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EshouPoint *EshouPointCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EshouPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EshouPoint *EshouPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EshouPoint *EshouPointTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Account is a free data retrieval call binding the contract method 0xc496a4e1.
//
// Solidity: function Account(uint256 ) constant returns(address)
func (_EshouPoint *EshouPointCaller) Account(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EshouPoint.contract.Call(opts, out, "Account", arg0)
	return *ret0, err
}

// Account is a free data retrieval call binding the contract method 0xc496a4e1.
//
// Solidity: function Account(uint256 ) constant returns(address)
func (_EshouPoint *EshouPointSession) Account(arg0 *big.Int) (common.Address, error) {
	return _EshouPoint.Contract.Account(&_EshouPoint.CallOpts, arg0)
}

// Account is a free data retrieval call binding the contract method 0xc496a4e1.
//
// Solidity: function Account(uint256 ) constant returns(address)
func (_EshouPoint *EshouPointCallerSession) Account(arg0 *big.Int) (common.Address, error) {
	return _EshouPoint.Contract.Account(&_EshouPoint.CallOpts, arg0)
}

// AccNum is a free data retrieval call binding the contract method 0x53c00424.
//
// Solidity: function accNum() constant returns(uint256)
func (_EshouPoint *EshouPointCaller) AccNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EshouPoint.contract.Call(opts, out, "accNum")
	return *ret0, err
}

// AccNum is a free data retrieval call binding the contract method 0x53c00424.
//
// Solidity: function accNum() constant returns(uint256)
func (_EshouPoint *EshouPointSession) AccNum() (*big.Int, error) {
	return _EshouPoint.Contract.AccNum(&_EshouPoint.CallOpts)
}

// AccNum is a free data retrieval call binding the contract method 0x53c00424.
//
// Solidity: function accNum() constant returns(uint256)
func (_EshouPoint *EshouPointCallerSession) AccNum() (*big.Int, error) {
	return _EshouPoint.Contract.AccNum(&_EshouPoint.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_EshouPoint *EshouPointCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EshouPoint.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_EshouPoint *EshouPointSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EshouPoint.Contract.Balances(&_EshouPoint.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_EshouPoint *EshouPointCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EshouPoint.Contract.Balances(&_EshouPoint.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_EshouPoint *EshouPointCaller) GetBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EshouPoint.contract.Call(opts, out, "getBalance", account)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_EshouPoint *EshouPointSession) GetBalance(account common.Address) (*big.Int, error) {
	return _EshouPoint.Contract.GetBalance(&_EshouPoint.CallOpts, account)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) constant returns(uint256)
func (_EshouPoint *EshouPointCallerSession) GetBalance(account common.Address) (*big.Int, error) {
	return _EshouPoint.Contract.GetBalance(&_EshouPoint.CallOpts, account)
}

// IsAccountExist is a free data retrieval call binding the contract method 0x14b4a4f9.
//
// Solidity: function isAccountExist(address account) constant returns(bool)
func (_EshouPoint *EshouPointCaller) IsAccountExist(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EshouPoint.contract.Call(opts, out, "isAccountExist", account)
	return *ret0, err
}

// IsAccountExist is a free data retrieval call binding the contract method 0x14b4a4f9.
//
// Solidity: function isAccountExist(address account) constant returns(bool)
func (_EshouPoint *EshouPointSession) IsAccountExist(account common.Address) (bool, error) {
	return _EshouPoint.Contract.IsAccountExist(&_EshouPoint.CallOpts, account)
}

// IsAccountExist is a free data retrieval call binding the contract method 0x14b4a4f9.
//
// Solidity: function isAccountExist(address account) constant returns(bool)
func (_EshouPoint *EshouPointCallerSession) IsAccountExist(account common.Address) (bool, error) {
	return _EshouPoint.Contract.IsAccountExist(&_EshouPoint.CallOpts, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns()
func (_EshouPoint *EshouPointTransactor) AddAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _EshouPoint.contract.TransactWithResult(opts, out, "addAccount", account)
	return transaction, receipt, err
}

func (_EshouPoint *EshouPointTransactor) AsyncAddAccount(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EshouPoint.contract.AsyncTransact(opts, handler, "addAccount", account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns()
func (_EshouPoint *EshouPointSession) AddAccount(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.AddAccount(&_EshouPoint.TransactOpts, account)
}

func (_EshouPoint *EshouPointSession) AsyncAddAccount(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncAddAccount(handler, &_EshouPoint.TransactOpts, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns()
func (_EshouPoint *EshouPointTransactorSession) AddAccount(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.AddAccount(&_EshouPoint.TransactOpts, account)
}

func (_EshouPoint *EshouPointTransactorSession) AsyncAddAccount(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncAddAccount(handler, &_EshouPoint.TransactOpts, account)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointTransactor) AddBalance(opts *bind.TransactOpts, account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _EshouPoint.contract.TransactWithResult(opts, out, "addBalance", account, amount)
	return *ret0, transaction, receipt, err
}

func (_EshouPoint *EshouPointTransactor) AsyncAddBalance(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.contract.AsyncTransact(opts, handler, "addBalance", account, amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointSession) AddBalance(account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.AddBalance(&_EshouPoint.TransactOpts, account, amount)
}

func (_EshouPoint *EshouPointSession) AsyncAddBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncAddBalance(handler, &_EshouPoint.TransactOpts, account, amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointTransactorSession) AddBalance(account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.AddBalance(&_EshouPoint.TransactOpts, account, amount)
}

func (_EshouPoint *EshouPointTransactorSession) AsyncAddBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncAddBalance(handler, &_EshouPoint.TransactOpts, account, amount)
}

// ReduceBalance is a paid mutator transaction binding the contract method 0xb4727906.
//
// Solidity: function reduceBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointTransactor) ReduceBalance(opts *bind.TransactOpts, account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _EshouPoint.contract.TransactWithResult(opts, out, "reduceBalance", account, amount)
	return *ret0, transaction, receipt, err
}

func (_EshouPoint *EshouPointTransactor) AsyncReduceBalance(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.contract.AsyncTransact(opts, handler, "reduceBalance", account, amount)
}

// ReduceBalance is a paid mutator transaction binding the contract method 0xb4727906.
//
// Solidity: function reduceBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointSession) ReduceBalance(account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.ReduceBalance(&_EshouPoint.TransactOpts, account, amount)
}

func (_EshouPoint *EshouPointSession) AsyncReduceBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncReduceBalance(handler, &_EshouPoint.TransactOpts, account, amount)
}

// ReduceBalance is a paid mutator transaction binding the contract method 0xb4727906.
//
// Solidity: function reduceBalance(address account, uint256 amount) returns(bool)
func (_EshouPoint *EshouPointTransactorSession) ReduceBalance(account common.Address, amount *big.Int) (bool, *types.Transaction, *types.Receipt, error) {
	return _EshouPoint.Contract.ReduceBalance(&_EshouPoint.TransactOpts, account, amount)
}

func (_EshouPoint *EshouPointTransactorSession) AsyncReduceBalance(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EshouPoint.Contract.AsyncReduceBalance(handler, &_EshouPoint.TransactOpts, account, amount)
}

// EshouPointAccountAdded represents a AccountAdded event raised by the EshouPoint contract.
type EshouPointAccountAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchAccountAdded is a free log subscription operation binding the contract event 0x8f42195a0bbfa58954be4349deb9efc38bdb9c298e529f705f8bc1e38bce0399.
//
// Solidity: event AccountAdded(address indexed account)
func (_EshouPoint *EshouPointFilterer) WatchAccountAdded(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "AccountAdded", account)
}

func (_EshouPoint *EshouPointFilterer) WatchAllAccountAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "AccountAdded")
}

// ParseAccountAdded is a log parse operation binding the contract event 0x8f42195a0bbfa58954be4349deb9efc38bdb9c298e529f705f8bc1e38bce0399.
//
// Solidity: event AccountAdded(address indexed account)
func (_EshouPoint *EshouPointFilterer) ParseAccountAdded(log types.Log) (*EshouPointAccountAdded, error) {
	event := new(EshouPointAccountAdded)
	if err := _EshouPoint.contract.UnpackLog(event, "AccountAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchAccountAdded is a free log subscription operation binding the contract event 0x8f42195a0bbfa58954be4349deb9efc38bdb9c298e529f705f8bc1e38bce0399.
//
// Solidity: event AccountAdded(address indexed account)
func (_EshouPoint *EshouPointSession) WatchAccountAdded(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.Contract.WatchAccountAdded(fromBlock, handler, account)
}

func (_EshouPoint *EshouPointSession) WatchAllAccountAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.Contract.WatchAllAccountAdded(fromBlock, handler)
}

// ParseAccountAdded is a log parse operation binding the contract event 0x8f42195a0bbfa58954be4349deb9efc38bdb9c298e529f705f8bc1e38bce0399.
//
// Solidity: event AccountAdded(address indexed account)
func (_EshouPoint *EshouPointSession) ParseAccountAdded(log types.Log) (*EshouPointAccountAdded, error) {
	return _EshouPoint.Contract.ParseAccountAdded(log)
}

// EshouPointBalanceAdded represents a BalanceAdded event raised by the EshouPoint contract.
type EshouPointBalanceAdded struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchBalanceAdded is a free log subscription operation binding the contract event 0xe96dd7a15a3974e8e7d5eb80de2cb6fd69907a1ee089170cd976cfffaedfa8ac.
//
// Solidity: event BalanceAdded(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointFilterer) WatchBalanceAdded(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "BalanceAdded", account)
}

func (_EshouPoint *EshouPointFilterer) WatchAllBalanceAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "BalanceAdded")
}

// ParseBalanceAdded is a log parse operation binding the contract event 0xe96dd7a15a3974e8e7d5eb80de2cb6fd69907a1ee089170cd976cfffaedfa8ac.
//
// Solidity: event BalanceAdded(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointFilterer) ParseBalanceAdded(log types.Log) (*EshouPointBalanceAdded, error) {
	event := new(EshouPointBalanceAdded)
	if err := _EshouPoint.contract.UnpackLog(event, "BalanceAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchBalanceAdded is a free log subscription operation binding the contract event 0xe96dd7a15a3974e8e7d5eb80de2cb6fd69907a1ee089170cd976cfffaedfa8ac.
//
// Solidity: event BalanceAdded(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointSession) WatchBalanceAdded(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.Contract.WatchBalanceAdded(fromBlock, handler, account)
}

func (_EshouPoint *EshouPointSession) WatchAllBalanceAdded(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.Contract.WatchAllBalanceAdded(fromBlock, handler)
}

// ParseBalanceAdded is a log parse operation binding the contract event 0xe96dd7a15a3974e8e7d5eb80de2cb6fd69907a1ee089170cd976cfffaedfa8ac.
//
// Solidity: event BalanceAdded(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointSession) ParseBalanceAdded(log types.Log) (*EshouPointBalanceAdded, error) {
	return _EshouPoint.Contract.ParseBalanceAdded(log)
}

// EshouPointBalanceReduced represents a BalanceReduced event raised by the EshouPoint contract.
type EshouPointBalanceReduced struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// WatchBalanceReduced is a free log subscription operation binding the contract event 0x2e9bf8d4a8402929da26de77a79494626b184ddae2e3e0c076d6dfa10cd2a1d9.
//
// Solidity: event BalanceReduced(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointFilterer) WatchBalanceReduced(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "BalanceReduced", account)
}

func (_EshouPoint *EshouPointFilterer) WatchAllBalanceReduced(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.contract.WatchLogs(fromBlock, handler, "BalanceReduced")
}

// ParseBalanceReduced is a log parse operation binding the contract event 0x2e9bf8d4a8402929da26de77a79494626b184ddae2e3e0c076d6dfa10cd2a1d9.
//
// Solidity: event BalanceReduced(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointFilterer) ParseBalanceReduced(log types.Log) (*EshouPointBalanceReduced, error) {
	event := new(EshouPointBalanceReduced)
	if err := _EshouPoint.contract.UnpackLog(event, "BalanceReduced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchBalanceReduced is a free log subscription operation binding the contract event 0x2e9bf8d4a8402929da26de77a79494626b184ddae2e3e0c076d6dfa10cd2a1d9.
//
// Solidity: event BalanceReduced(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointSession) WatchBalanceReduced(fromBlock *uint64, handler func(int, []types.Log), account common.Address) (string, error) {
	return _EshouPoint.Contract.WatchBalanceReduced(fromBlock, handler, account)
}

func (_EshouPoint *EshouPointSession) WatchAllBalanceReduced(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EshouPoint.Contract.WatchAllBalanceReduced(fromBlock, handler)
}

// ParseBalanceReduced is a log parse operation binding the contract event 0x2e9bf8d4a8402929da26de77a79494626b184ddae2e3e0c076d6dfa10cd2a1d9.
//
// Solidity: event BalanceReduced(address indexed account, uint256 amount)
func (_EshouPoint *EshouPointSession) ParseBalanceReduced(log types.Log) (*EshouPointBalanceReduced, error) {
	return _EshouPoint.Contract.ParseBalanceReduced(log)
}
