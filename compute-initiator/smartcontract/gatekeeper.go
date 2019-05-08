// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontract

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

// GatekeeperABI is the input ABI used to generate the binding from.
const GatekeeperABI = "[{\"name\":\"Query\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true},{\"type\":\"string\",\"name\":\"clientReference\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"queryType\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"queryId\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"identifier\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"attribute\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"submit_calc_query\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"clientReference\"},{\"type\":\"int128\",\"name\":\"identifier\"},{\"type\":\"uint256\",\"name\":\"attribute\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51146},{\"name\":\"submit_info_query\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"clientReference\"},{\"type\":\"int128\",\"name\":\"identifier\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51167},{\"name\":\"allowed_clients\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715}]"

// Gatekeeper is an auto generated Go binding around an Ethereum contract.
type Gatekeeper struct {
	GatekeeperCaller     // Read-only binding to the contract
	GatekeeperTransactor // Write-only binding to the contract
	GatekeeperFilterer   // Log filterer for contract events
}

// GatekeeperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatekeeperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatekeeperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatekeeperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatekeeperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatekeeperSession struct {
	Contract     *Gatekeeper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatekeeperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatekeeperCallerSession struct {
	Contract *GatekeeperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GatekeeperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatekeeperTransactorSession struct {
	Contract     *GatekeeperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GatekeeperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatekeeperRaw struct {
	Contract *Gatekeeper // Generic contract binding to access the raw methods on
}

// GatekeeperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatekeeperCallerRaw struct {
	Contract *GatekeeperCaller // Generic read-only contract binding to access the raw methods on
}

// GatekeeperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatekeeperTransactorRaw struct {
	Contract *GatekeeperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatekeeper creates a new instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeper(address common.Address, backend bind.ContractBackend) (*Gatekeeper, error) {
	contract, err := bindGatekeeper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gatekeeper{GatekeeperCaller: GatekeeperCaller{contract: contract}, GatekeeperTransactor: GatekeeperTransactor{contract: contract}, GatekeeperFilterer: GatekeeperFilterer{contract: contract}}, nil
}

// NewGatekeeperCaller creates a new read-only instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperCaller(address common.Address, caller bind.ContractCaller) (*GatekeeperCaller, error) {
	contract, err := bindGatekeeper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatekeeperCaller{contract: contract}, nil
}

// NewGatekeeperTransactor creates a new write-only instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperTransactor(address common.Address, transactor bind.ContractTransactor) (*GatekeeperTransactor, error) {
	contract, err := bindGatekeeper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatekeeperTransactor{contract: contract}, nil
}

// NewGatekeeperFilterer creates a new log filterer instance of Gatekeeper, bound to a specific deployed contract.
func NewGatekeeperFilterer(address common.Address, filterer bind.ContractFilterer) (*GatekeeperFilterer, error) {
	contract, err := bindGatekeeper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatekeeperFilterer{contract: contract}, nil
}

// bindGatekeeper binds a generic wrapper to an already deployed contract.
func bindGatekeeper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatekeeperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gatekeeper *GatekeeperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gatekeeper.Contract.GatekeeperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gatekeeper *GatekeeperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gatekeeper.Contract.GatekeeperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gatekeeper *GatekeeperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gatekeeper.Contract.GatekeeperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gatekeeper *GatekeeperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gatekeeper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gatekeeper *GatekeeperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gatekeeper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gatekeeper *GatekeeperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gatekeeper.Contract.contract.Transact(opts, method, params...)
}

// AllowedClients is a free data retrieval call binding the contract method 0x98f0e140.
//
// Solidity: function allowed_clients(address arg0) constant returns(bool out)
func (_Gatekeeper *GatekeeperCaller) AllowedClients(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gatekeeper.contract.Call(opts, out, "allowed_clients", arg0)
	return *ret0, err
}

// AllowedClients is a free data retrieval call binding the contract method 0x98f0e140.
//
// Solidity: function allowed_clients(address arg0) constant returns(bool out)
func (_Gatekeeper *GatekeeperSession) AllowedClients(arg0 common.Address) (bool, error) {
	return _Gatekeeper.Contract.AllowedClients(&_Gatekeeper.CallOpts, arg0)
}

// AllowedClients is a free data retrieval call binding the contract method 0x98f0e140.
//
// Solidity: function allowed_clients(address arg0) constant returns(bool out)
func (_Gatekeeper *GatekeeperCallerSession) AllowedClients(arg0 common.Address) (bool, error) {
	return _Gatekeeper.Contract.AllowedClients(&_Gatekeeper.CallOpts, arg0)
}

// SubmitCalcQuery is a paid mutator transaction binding the contract method 0x0d305afa.
//
// Solidity: function submit_calc_query(string clientReference, int128 identifier, uint256 attribute) returns()
func (_Gatekeeper *GatekeeperTransactor) SubmitCalcQuery(opts *bind.TransactOpts, clientReference string, identifier *big.Int, attribute *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "submit_calc_query", clientReference, identifier, attribute)
}

// SubmitCalcQuery is a paid mutator transaction binding the contract method 0x0d305afa.
//
// Solidity: function submit_calc_query(string clientReference, int128 identifier, uint256 attribute) returns()
func (_Gatekeeper *GatekeeperSession) SubmitCalcQuery(clientReference string, identifier *big.Int, attribute *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.SubmitCalcQuery(&_Gatekeeper.TransactOpts, clientReference, identifier, attribute)
}

// SubmitCalcQuery is a paid mutator transaction binding the contract method 0x0d305afa.
//
// Solidity: function submit_calc_query(string clientReference, int128 identifier, uint256 attribute) returns()
func (_Gatekeeper *GatekeeperTransactorSession) SubmitCalcQuery(clientReference string, identifier *big.Int, attribute *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.SubmitCalcQuery(&_Gatekeeper.TransactOpts, clientReference, identifier, attribute)
}

// SubmitInfoQuery is a paid mutator transaction binding the contract method 0xbe133d01.
//
// Solidity: function submit_info_query(string clientReference, int128 identifier) returns()
func (_Gatekeeper *GatekeeperTransactor) SubmitInfoQuery(opts *bind.TransactOpts, clientReference string, identifier *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.contract.Transact(opts, "submit_info_query", clientReference, identifier)
}

// SubmitInfoQuery is a paid mutator transaction binding the contract method 0xbe133d01.
//
// Solidity: function submit_info_query(string clientReference, int128 identifier) returns()
func (_Gatekeeper *GatekeeperSession) SubmitInfoQuery(clientReference string, identifier *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.SubmitInfoQuery(&_Gatekeeper.TransactOpts, clientReference, identifier)
}

// SubmitInfoQuery is a paid mutator transaction binding the contract method 0xbe133d01.
//
// Solidity: function submit_info_query(string clientReference, int128 identifier) returns()
func (_Gatekeeper *GatekeeperTransactorSession) SubmitInfoQuery(clientReference string, identifier *big.Int) (*types.Transaction, error) {
	return _Gatekeeper.Contract.SubmitInfoQuery(&_Gatekeeper.TransactOpts, clientReference, identifier)
}

// GatekeeperQueryIterator is returned from FilterQuery and is used to iterate over the raw logs and unpacked data for Query events raised by the Gatekeeper contract.
type GatekeeperQueryIterator struct {
	Event *GatekeeperQuery // Event containing the contract specifics and raw log

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
func (it *GatekeeperQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatekeeperQuery)
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
		it.Event = new(GatekeeperQuery)
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
func (it *GatekeeperQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatekeeperQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatekeeperQuery represents a Query event raised by the Gatekeeper contract.
type GatekeeperQuery struct {
	Sender          common.Address
	ClientReference string
	QueryType       *big.Int
	QueryId         *big.Int
	Identifier      *big.Int
	Attribute       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterQuery is a free log retrieval operation binding the contract event 0x6c3a59d71d35fc27c453f5602f1cdbbbccd5bb00dbcb062a2db991a9e36d1baa.
//
// Solidity: event Query(address indexed sender, string clientReference, uint256 queryType, int128 queryId, int128 identifier, uint256 attribute)
func (_Gatekeeper *GatekeeperFilterer) FilterQuery(opts *bind.FilterOpts, sender []common.Address) (*GatekeeperQueryIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gatekeeper.contract.FilterLogs(opts, "Query", senderRule)
	if err != nil {
		return nil, err
	}
	return &GatekeeperQueryIterator{contract: _Gatekeeper.contract, event: "Query", logs: logs, sub: sub}, nil
}

// WatchQuery is a free log subscription operation binding the contract event 0x6c3a59d71d35fc27c453f5602f1cdbbbccd5bb00dbcb062a2db991a9e36d1baa.
//
// Solidity: event Query(address indexed sender, string clientReference, uint256 queryType, int128 queryId, int128 identifier, uint256 attribute)
func (_Gatekeeper *GatekeeperFilterer) WatchQuery(opts *bind.WatchOpts, sink chan<- *GatekeeperQuery, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gatekeeper.contract.WatchLogs(opts, "Query", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatekeeperQuery)
				if err := _Gatekeeper.contract.UnpackLog(event, "Query", log); err != nil {
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
