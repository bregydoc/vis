// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rgy

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

// RgyABI is the input ABI used to generate the binding from.
const RgyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAvailableShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"how\",\"type\":\"uint256\"}],\"name\":\"sellShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSoldShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableToTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_shares\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_how\",\"type\":\"uint256\"}],\"name\":\"ShareSold\",\"type\":\"event\"}]"

// RgyBin is the compiled bytecode used for deploying new contracts.
var RgyBin = "0x60806040526000600660006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5060405161098b38038061098b8339818101604052604081101561004e57600080fd5b81019080805164010000000081111561006657600080fd5b8281019050602081018481111561007c57600080fd5b815185600182028301116401000000008211171561009957600080fd5b505092919060200180519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816003908051906020019061010492919061011a565b50806004819055508060058190555050506101bf565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015b57805160ff1916838001178555610189565b82800160010185558215610189579182015b8281111561018857825182559160200191906001019061016d565b5b509050610196919061019a565b5090565b6101bc91905b808211156101b85760008160009055506001016101a0565b5090565b90565b6107bd806101ce6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b51d05341161005b578063b51d053414610141578063ba0aef15146101a3578063d57d0d39146101c1578063e3d670d7146101e35761007d565b806306fdde03146100825780631a43bcb514610105578063817bc0cb14610123575b600080fd5b61008a61023b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100ca5780820151818401526020810190506100af565b50505050905090810190601f1680156100f75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010d6102d9565b6040518082815260200191505060405180910390f35b61012b6102e3565b6040518082815260200191505060405180910390f35b61018d6004803603604081101561015757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506102ed565b6040518082815260200191505060405180910390f35b6101ab6106b7565b6040518082815260200191505060405180910390f35b6101c961075d565b604051808215151515815260200191505060405180910390f35b610225600480360360208110156101f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610770565b6040518082815260200191505060405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102d15780601f106102a6576101008083540402835291602001916102d1565b820191906000526020600020905b8154815290600101906020018083116102b457829003601f168201915b505050505081565b6000600454905090565b6000600554905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f796f7520617265206e6f742052656e766573746779000000000000000000000081525060200191505060405180910390fd5b816005541015610429576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f696e7375666963656e7420666f756e647300000000000000000000000000000081525060200191505060405180910390fd5b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600560008282540392505081905550600060055410156104ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f696e76616c69642073656c6c000000000000000000000000000000000000000081525060200191505060405180910390fd5b600060646004546005548161051057fe5b040290506014811015610539576001600660006101000a81548160ff0219169083151502179055505b600080905060008090505b6002805490508110156105d2578573ffffffffffffffffffffffffffffffffffffffff166002828154811061057557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156105c557600191506105d2565b8080600101915050610544565b508061063f5760028590806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b7fe8bb4edfb49cb224db3d7d51f707c053865322c038a54cfcfd81969d0eb1f0138585604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16005549250505092915050565b60008060008090505b6002805490508110156107555760016000600283815481106106de57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548201915080806001019150506106c0565b508091505090565b600660009054906101000a900460ff1681565b6001602052806000526040600020600091509050548156fea265627a7a72305820a1e4a1aef7c116b365e0d179de9315dcc7490c3ffc31f01e3dc30f69872042cd64736f6c634300050a0032"

// DeployRgy deploys a new Ethereum contract, binding an instance of Rgy to it.
func DeployRgy(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _shares *big.Int) (common.Address, *types.Transaction, *Rgy, error) {
	parsed, err := abi.JSON(strings.NewReader(RgyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RgyBin), backend, _name, _shares)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rgy{RgyCaller: RgyCaller{contract: contract}, RgyTransactor: RgyTransactor{contract: contract}, RgyFilterer: RgyFilterer{contract: contract}}, nil
}

// Rgy is an auto generated Go binding around an Ethereum contract.
type Rgy struct {
	RgyCaller     // Read-only binding to the contract
	RgyTransactor // Write-only binding to the contract
	RgyFilterer   // Log filterer for contract events
}

// RgyCaller is an auto generated read-only Go binding around an Ethereum contract.
type RgyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RgyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RgyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RgyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RgyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RgySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RgySession struct {
	Contract     *Rgy              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RgyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RgyCallerSession struct {
	Contract *RgyCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RgyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RgyTransactorSession struct {
	Contract     *RgyTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RgyRaw is an auto generated low-level Go binding around an Ethereum contract.
type RgyRaw struct {
	Contract *Rgy // Generic contract binding to access the raw methods on
}

// RgyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RgyCallerRaw struct {
	Contract *RgyCaller // Generic read-only contract binding to access the raw methods on
}

// RgyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RgyTransactorRaw struct {
	Contract *RgyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRgy creates a new instance of Rgy, bound to a specific deployed contract.
func NewRgy(address common.Address, backend bind.ContractBackend) (*Rgy, error) {
	contract, err := bindRgy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rgy{RgyCaller: RgyCaller{contract: contract}, RgyTransactor: RgyTransactor{contract: contract}, RgyFilterer: RgyFilterer{contract: contract}}, nil
}

// NewRgyCaller creates a new read-only instance of Rgy, bound to a specific deployed contract.
func NewRgyCaller(address common.Address, caller bind.ContractCaller) (*RgyCaller, error) {
	contract, err := bindRgy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RgyCaller{contract: contract}, nil
}

// NewRgyTransactor creates a new write-only instance of Rgy, bound to a specific deployed contract.
func NewRgyTransactor(address common.Address, transactor bind.ContractTransactor) (*RgyTransactor, error) {
	contract, err := bindRgy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RgyTransactor{contract: contract}, nil
}

// NewRgyFilterer creates a new log filterer instance of Rgy, bound to a specific deployed contract.
func NewRgyFilterer(address common.Address, filterer bind.ContractFilterer) (*RgyFilterer, error) {
	contract, err := bindRgy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RgyFilterer{contract: contract}, nil
}

// bindRgy binds a generic wrapper to an already deployed contract.
func bindRgy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RgyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rgy *RgyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rgy.Contract.RgyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rgy *RgyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rgy.Contract.RgyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rgy *RgyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rgy.Contract.RgyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rgy *RgyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rgy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rgy *RgyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rgy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rgy *RgyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rgy.Contract.contract.Transact(opts, method, params...)
}

// AvailableToTransfer is a free data retrieval call binding the contract method 0xd57d0d39.
//
// Solidity: function availableToTransfer() constant returns(bool)
func (_Rgy *RgyCaller) AvailableToTransfer(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "availableToTransfer")
	return *ret0, err
}

// AvailableToTransfer is a free data retrieval call binding the contract method 0xd57d0d39.
//
// Solidity: function availableToTransfer() constant returns(bool)
func (_Rgy *RgySession) AvailableToTransfer() (bool, error) {
	return _Rgy.Contract.AvailableToTransfer(&_Rgy.CallOpts)
}

// AvailableToTransfer is a free data retrieval call binding the contract method 0xd57d0d39.
//
// Solidity: function availableToTransfer() constant returns(bool)
func (_Rgy *RgyCallerSession) AvailableToTransfer() (bool, error) {
	return _Rgy.Contract.AvailableToTransfer(&_Rgy.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) constant returns(uint256)
func (_Rgy *RgyCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "balance", arg0)
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) constant returns(uint256)
func (_Rgy *RgySession) Balance(arg0 common.Address) (*big.Int, error) {
	return _Rgy.Contract.Balance(&_Rgy.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) constant returns(uint256)
func (_Rgy *RgyCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _Rgy.Contract.Balance(&_Rgy.CallOpts, arg0)
}

// GetAvailableShares is a free data retrieval call binding the contract method 0x817bc0cb.
//
// Solidity: function getAvailableShares() constant returns(uint256)
func (_Rgy *RgyCaller) GetAvailableShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "getAvailableShares")
	return *ret0, err
}

// GetAvailableShares is a free data retrieval call binding the contract method 0x817bc0cb.
//
// Solidity: function getAvailableShares() constant returns(uint256)
func (_Rgy *RgySession) GetAvailableShares() (*big.Int, error) {
	return _Rgy.Contract.GetAvailableShares(&_Rgy.CallOpts)
}

// GetAvailableShares is a free data retrieval call binding the contract method 0x817bc0cb.
//
// Solidity: function getAvailableShares() constant returns(uint256)
func (_Rgy *RgyCallerSession) GetAvailableShares() (*big.Int, error) {
	return _Rgy.Contract.GetAvailableShares(&_Rgy.CallOpts)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(uint256)
func (_Rgy *RgyCaller) GetGenesis(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "getGenesis")
	return *ret0, err
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(uint256)
func (_Rgy *RgySession) GetGenesis() (*big.Int, error) {
	return _Rgy.Contract.GetGenesis(&_Rgy.CallOpts)
}

// GetGenesis is a free data retrieval call binding the contract method 0x1a43bcb5.
//
// Solidity: function getGenesis() constant returns(uint256)
func (_Rgy *RgyCallerSession) GetGenesis() (*big.Int, error) {
	return _Rgy.Contract.GetGenesis(&_Rgy.CallOpts)
}

// GetSoldShares is a free data retrieval call binding the contract method 0xba0aef15.
//
// Solidity: function getSoldShares() constant returns(uint256)
func (_Rgy *RgyCaller) GetSoldShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "getSoldShares")
	return *ret0, err
}

// GetSoldShares is a free data retrieval call binding the contract method 0xba0aef15.
//
// Solidity: function getSoldShares() constant returns(uint256)
func (_Rgy *RgySession) GetSoldShares() (*big.Int, error) {
	return _Rgy.Contract.GetSoldShares(&_Rgy.CallOpts)
}

// GetSoldShares is a free data retrieval call binding the contract method 0xba0aef15.
//
// Solidity: function getSoldShares() constant returns(uint256)
func (_Rgy *RgyCallerSession) GetSoldShares() (*big.Int, error) {
	return _Rgy.Contract.GetSoldShares(&_Rgy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Rgy *RgyCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Rgy.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Rgy *RgySession) Name() (string, error) {
	return _Rgy.Contract.Name(&_Rgy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Rgy *RgyCallerSession) Name() (string, error) {
	return _Rgy.Contract.Name(&_Rgy.CallOpts)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address to, uint256 how) returns(uint256)
func (_Rgy *RgyTransactor) SellShares(opts *bind.TransactOpts, to common.Address, how *big.Int) (*types.Transaction, error) {
	return _Rgy.contract.Transact(opts, "sellShares", to, how)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address to, uint256 how) returns(uint256)
func (_Rgy *RgySession) SellShares(to common.Address, how *big.Int) (*types.Transaction, error) {
	return _Rgy.Contract.SellShares(&_Rgy.TransactOpts, to, how)
}

// SellShares is a paid mutator transaction binding the contract method 0xb51d0534.
//
// Solidity: function sellShares(address to, uint256 how) returns(uint256)
func (_Rgy *RgyTransactorSession) SellShares(to common.Address, how *big.Int) (*types.Transaction, error) {
	return _Rgy.Contract.SellShares(&_Rgy.TransactOpts, to, how)
}

// RgyShareSoldIterator is returned from FilterShareSold and is used to iterate over the raw logs and unpacked data for ShareSold events raised by the Rgy contract.
type RgyShareSoldIterator struct {
	Event *RgyShareSold // Event containing the contract specifics and raw log

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
func (it *RgyShareSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RgyShareSold)
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
		it.Event = new(RgyShareSold)
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
func (it *RgyShareSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RgyShareSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RgyShareSold represents a ShareSold event raised by the Rgy contract.
type RgyShareSold struct {
	To  common.Address
	How *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShareSold is a free log retrieval operation binding the contract event 0xe8bb4edfb49cb224db3d7d51f707c053865322c038a54cfcfd81969d0eb1f013.
//
// Solidity: event ShareSold(address _to, uint256 _how)
func (_Rgy *RgyFilterer) FilterShareSold(opts *bind.FilterOpts) (*RgyShareSoldIterator, error) {

	logs, sub, err := _Rgy.contract.FilterLogs(opts, "ShareSold")
	if err != nil {
		return nil, err
	}
	return &RgyShareSoldIterator{contract: _Rgy.contract, event: "ShareSold", logs: logs, sub: sub}, nil
}

// WatchShareSold is a free log subscription operation binding the contract event 0xe8bb4edfb49cb224db3d7d51f707c053865322c038a54cfcfd81969d0eb1f013.
//
// Solidity: event ShareSold(address _to, uint256 _how)
func (_Rgy *RgyFilterer) WatchShareSold(opts *bind.WatchOpts, sink chan<- *RgyShareSold) (event.Subscription, error) {

	logs, sub, err := _Rgy.contract.WatchLogs(opts, "ShareSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RgyShareSold)
				if err := _Rgy.contract.UnpackLog(event, "ShareSold", log); err != nil {
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

// ParseShareSold is a log parse operation binding the contract event 0xe8bb4edfb49cb224db3d7d51f707c053865322c038a54cfcfd81969d0eb1f013.
//
// Solidity: event ShareSold(address _to, uint256 _how)
func (_Rgy *RgyFilterer) ParseShareSold(log types.Log) (*RgyShareSold, error) {
	event := new(RgyShareSold)
	if err := _Rgy.contract.UnpackLog(event, "ShareSold", log); err != nil {
		return nil, err
	}
	return event, nil
}
