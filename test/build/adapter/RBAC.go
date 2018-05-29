// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RBACABI is the input ABI used to generate the binding from.
const RBACABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// RBACBin is the compiled bytecode used for deploying new contracts.
const RBACBin = `0x608060405234801561001057600080fd5b5061029c806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610050578063217fe6c6146100c6575b600080fd5b34801561005c57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100c495833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375094975061014e9650505050505050565b005b3480156100d257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261013a95833573ffffffffffffffffffffffffffffffffffffffff169536956044949193909101919081908401838280828437509497506101bc9650505050505050565b604080519115158252519081900360200190f35b6101b8826000836040518082805190602001908083835b602083106101845780518252601f199092019160209182019101610165565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061022f565b5050565b6000610228836000846040518082805190602001908083835b602083106101f45780518252601f1990920191602091820191016101d5565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610244565b9392505050565b6102398282610244565b15156101b857600080fd5b73ffffffffffffffffffffffffffffffffffffffff166000908152602091909152604090205460ff16905600a165627a7a72305820aa2014d59e025dd1798005f631eafd97956a5292df972b5f9c84326f914de7820029`

// DeployRBAC deploys a new Ethereum contract, binding an instance of RBAC to it.
func DeployRBAC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RBAC, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// RBAC is an auto generated Go binding around an Ethereum contract.
type RBAC struct {
	RBACCaller     // Read-only binding to the contract
	RBACTransactor // Write-only binding to the contract
	RBACFilterer   // Log filterer for contract events
}

// RBACCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RBACFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACSession struct {
	Contract     *RBAC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACCallerSession struct {
	Contract *RBACCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RBACTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACTransactorSession struct {
	Contract     *RBACTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACRaw struct {
	Contract *RBAC // Generic contract binding to access the raw methods on
}

// RBACCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACCallerRaw struct {
	Contract *RBACCaller // Generic read-only contract binding to access the raw methods on
}

// RBACTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACTransactorRaw struct {
	Contract *RBACTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBAC creates a new instance of RBAC, bound to a specific deployed contract.
func NewRBAC(address common.Address, backend bind.ContractBackend) (*RBAC, error) {
	contract, err := bindRBAC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// NewRBACCaller creates a new read-only instance of RBAC, bound to a specific deployed contract.
func NewRBACCaller(address common.Address, caller bind.ContractCaller) (*RBACCaller, error) {
	contract, err := bindRBAC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RBACCaller{contract: contract}, nil
}

// NewRBACTransactor creates a new write-only instance of RBAC, bound to a specific deployed contract.
func NewRBACTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACTransactor, error) {
	contract, err := bindRBAC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RBACTransactor{contract: contract}, nil
}

// NewRBACFilterer creates a new log filterer instance of RBAC, bound to a specific deployed contract.
func NewRBACFilterer(address common.Address, filterer bind.ContractFilterer) (*RBACFilterer, error) {
	contract, err := bindRBAC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RBACFilterer{contract: contract}, nil
}

// bindRBAC binds a generic wrapper to an already deployed contract.
func bindRBAC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.RBACCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transact(opts, method, params...)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACCaller) CheckRole(opts *bind.CallOpts, addr common.Address, roleName string) error {
	var ()
	out := &[]interface{}{}
	err := _RBAC.contract.Call(opts, out, "checkRole", addr, roleName)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACSession) CheckRole(addr common.Address, roleName string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, addr, roleName)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_RBAC *RBACCallerSession) CheckRole(addr common.Address, roleName string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACCaller) HasRole(opts *bind.CallOpts, addr common.Address, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBAC.contract.Call(opts, out, "hasRole", addr, roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_RBAC *RBACCallerSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, addr, roleName)
}

// RBACRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the RBAC contract.
type RBACRoleAddedIterator struct {
	Event *RBACRoleAdded // Event containing the contract specifics and raw log

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
func (it *RBACRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleAdded)
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
		it.Event = new(RBACRoleAdded)
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
func (it *RBACRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleAdded represents a RoleAdded event raised by the RBAC contract.
type RBACRoleAdded struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: event RoleAdded(addr address, roleName string)
func (_RBAC *RBACFilterer) FilterRoleAdded(opts *bind.FilterOpts) (*RBACRoleAddedIterator, error) {

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return &RBACRoleAddedIterator{contract: _RBAC.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: event RoleAdded(addr address, roleName string)
func (_RBAC *RBACFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *RBACRoleAdded) (event.Subscription, error) {

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleAdded)
				if err := _RBAC.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// RBACRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the RBAC contract.
type RBACRoleRemovedIterator struct {
	Event *RBACRoleRemoved // Event containing the contract specifics and raw log

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
func (it *RBACRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleRemoved)
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
		it.Event = new(RBACRoleRemoved)
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
func (it *RBACRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleRemoved represents a RoleRemoved event raised by the RBAC contract.
type RBACRoleRemoved struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: event RoleRemoved(addr address, roleName string)
func (_RBAC *RBACFilterer) FilterRoleRemoved(opts *bind.FilterOpts) (*RBACRoleRemovedIterator, error) {

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return &RBACRoleRemovedIterator{contract: _RBAC.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: event RoleRemoved(addr address, roleName string)
func (_RBAC *RBACFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *RBACRoleRemoved) (event.Subscription, error) {

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleRemoved)
				if err := _RBAC.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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
