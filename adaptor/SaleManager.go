// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adaptor

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SaleManagerABI is the input ABI used to generate the binding from.
const SaleManagerABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Sale\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"refundMany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"releaseMany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSaleAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_sale\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"Fail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SaleManager is an auto generated Go binding around an Ethereum contract.
type SaleManager struct {
	SaleManagerCaller     // Read-only binding to the contract
	SaleManagerTransactor // Write-only binding to the contract
	SaleManagerFilterer   // Log filterer for contract events
}

// SaleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleManagerSession struct {
	Contract     *SaleManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleManagerCallerSession struct {
	Contract *SaleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SaleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleManagerTransactorSession struct {
	Contract     *SaleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SaleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleManagerRaw struct {
	Contract *SaleManager // Generic contract binding to access the raw methods on
}

// SaleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleManagerCallerRaw struct {
	Contract *SaleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SaleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleManagerTransactorRaw struct {
	Contract *SaleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleManager creates a new instance of SaleManager, bound to a specific deployed contract.
func NewSaleManager(address common.Address, backend bind.ContractBackend) (*SaleManager, error) {
	contract, err := bindSaleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleManager{SaleManagerCaller: SaleManagerCaller{contract: contract}, SaleManagerTransactor: SaleManagerTransactor{contract: contract}, SaleManagerFilterer: SaleManagerFilterer{contract: contract}}, nil
}

// NewSaleManagerCaller creates a new read-only instance of SaleManager, bound to a specific deployed contract.
func NewSaleManagerCaller(address common.Address, caller bind.ContractCaller) (*SaleManagerCaller, error) {
	contract, err := bindSaleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleManagerCaller{contract: contract}, nil
}

// NewSaleManagerTransactor creates a new write-only instance of SaleManager, bound to a specific deployed contract.
func NewSaleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleManagerTransactor, error) {
	contract, err := bindSaleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleManagerTransactor{contract: contract}, nil
}

// NewSaleManagerFilterer creates a new log filterer instance of SaleManager, bound to a specific deployed contract.
func NewSaleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleManagerFilterer, error) {
	contract, err := bindSaleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleManagerFilterer{contract: contract}, nil
}

// bindSaleManager binds a generic wrapper to an already deployed contract.
func bindSaleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleManager *SaleManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleManager.Contract.SaleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleManager *SaleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleManager.Contract.SaleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleManager *SaleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleManager.Contract.SaleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleManager *SaleManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleManager *SaleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleManager *SaleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleManager.Contract.contract.Transact(opts, method, params...)
}

// Sale is a free data retrieval call binding the contract method 0xb78f9de7.
//
// Solidity: function Sale() constant returns(address)
func (_SaleManager *SaleManagerCaller) Sale(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleManager.contract.Call(opts, out, "Sale")
	return *ret0, err
}

// Sale is a free data retrieval call binding the contract method 0xb78f9de7.
//
// Solidity: function Sale() constant returns(address)
func (_SaleManager *SaleManagerSession) Sale() (common.Address, error) {
	return _SaleManager.Contract.Sale(&_SaleManager.CallOpts)
}

// Sale is a free data retrieval call binding the contract method 0xb78f9de7.
//
// Solidity: function Sale() constant returns(address)
func (_SaleManager *SaleManagerCallerSession) Sale() (common.Address, error) {
	return _SaleManager.Contract.Sale(&_SaleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleManager *SaleManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleManager *SaleManagerSession) Owner() (common.Address, error) {
	return _SaleManager.Contract.Owner(&_SaleManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaleManager *SaleManagerCallerSession) Owner() (common.Address, error) {
	return _SaleManager.Contract.Owner(&_SaleManager.CallOpts)
}

// RefundMany is a paid mutator transaction binding the contract method 0xcee829ea.
//
// Solidity: function refundMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerTransactor) RefundMany(opts *bind.TransactOpts, _addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.contract.Transact(opts, "refundMany", _addrs)
}

// RefundMany is a paid mutator transaction binding the contract method 0xcee829ea.
//
// Solidity: function refundMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerSession) RefundMany(_addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.RefundMany(&_SaleManager.TransactOpts, _addrs)
}

// RefundMany is a paid mutator transaction binding the contract method 0xcee829ea.
//
// Solidity: function refundMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerTransactorSession) RefundMany(_addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.RefundMany(&_SaleManager.TransactOpts, _addrs)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerTransactor) ReleaseMany(opts *bind.TransactOpts, _addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.contract.Transact(opts, "releaseMany", _addrs)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerSession) ReleaseMany(_addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.ReleaseMany(&_SaleManager.TransactOpts, _addrs)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_SaleManager *SaleManagerTransactorSession) ReleaseMany(_addrs []common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.ReleaseMany(&_SaleManager.TransactOpts, _addrs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SaleManager *SaleManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SaleManager *SaleManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _SaleManager.Contract.RenounceOwnership(&_SaleManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SaleManager *SaleManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SaleManager.Contract.RenounceOwnership(&_SaleManager.TransactOpts)
}

// SetSaleAddress is a paid mutator transaction binding the contract method 0xf8fb491f.
//
// Solidity: function setSaleAddress(_addr address) returns()
func (_SaleManager *SaleManagerTransactor) SetSaleAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SaleManager.contract.Transact(opts, "setSaleAddress", _addr)
}

// SetSaleAddress is a paid mutator transaction binding the contract method 0xf8fb491f.
//
// Solidity: function setSaleAddress(_addr address) returns()
func (_SaleManager *SaleManagerSession) SetSaleAddress(_addr common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.SetSaleAddress(&_SaleManager.TransactOpts, _addr)
}

// SetSaleAddress is a paid mutator transaction binding the contract method 0xf8fb491f.
//
// Solidity: function setSaleAddress(_addr address) returns()
func (_SaleManager *SaleManagerTransactorSession) SetSaleAddress(_addr common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.SetSaleAddress(&_SaleManager.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleManager *SaleManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SaleManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleManager *SaleManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.TransferOwnership(&_SaleManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SaleManager *SaleManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleManager.Contract.TransferOwnership(&_SaleManager.TransactOpts, newOwner)
}

// SaleManagerFailIterator is returned from FilterFail and is used to iterate over the raw logs and unpacked data for Fail events raised by the SaleManager contract.
type SaleManagerFailIterator struct {
	Event *SaleManagerFail // Event containing the contract specifics and raw log

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
func (it *SaleManagerFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleManagerFail)
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
		it.Event = new(SaleManagerFail)
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
func (it *SaleManagerFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleManagerFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleManagerFail represents a Fail event raised by the SaleManager contract.
type SaleManagerFail struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFail is a free log retrieval operation binding the contract event 0x21d0324cd6cce67ffe52c8955137310671bcef292844385858e00841a18d60c0.
//
// Solidity: event Fail(_addr indexed address)
func (_SaleManager *SaleManagerFilterer) FilterFail(opts *bind.FilterOpts, _addr []common.Address) (*SaleManagerFailIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _SaleManager.contract.FilterLogs(opts, "Fail", _addrRule)
	if err != nil {
		return nil, err
	}
	return &SaleManagerFailIterator{contract: _SaleManager.contract, event: "Fail", logs: logs, sub: sub}, nil
}

// WatchFail is a free log subscription operation binding the contract event 0x21d0324cd6cce67ffe52c8955137310671bcef292844385858e00841a18d60c0.
//
// Solidity: event Fail(_addr indexed address)
func (_SaleManager *SaleManagerFilterer) WatchFail(opts *bind.WatchOpts, sink chan<- *SaleManagerFail, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _SaleManager.contract.WatchLogs(opts, "Fail", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleManagerFail)
				if err := _SaleManager.contract.UnpackLog(event, "Fail", log); err != nil {
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

// SaleManagerOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the SaleManager contract.
type SaleManagerOwnershipRenouncedIterator struct {
	Event *SaleManagerOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *SaleManagerOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleManagerOwnershipRenounced)
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
		it.Event = new(SaleManagerOwnershipRenounced)
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
func (it *SaleManagerOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleManagerOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleManagerOwnershipRenounced represents a OwnershipRenounced event raised by the SaleManager contract.
type SaleManagerOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_SaleManager *SaleManagerFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*SaleManagerOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SaleManager.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SaleManagerOwnershipRenouncedIterator{contract: _SaleManager.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_SaleManager *SaleManagerFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *SaleManagerOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SaleManager.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleManagerOwnershipRenounced)
				if err := _SaleManager.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// SaleManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SaleManager contract.
type SaleManagerOwnershipTransferredIterator struct {
	Event *SaleManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SaleManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleManagerOwnershipTransferred)
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
		it.Event = new(SaleManagerOwnershipTransferred)
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
func (it *SaleManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleManagerOwnershipTransferred represents a OwnershipTransferred event raised by the SaleManager contract.
type SaleManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SaleManager *SaleManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SaleManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SaleManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SaleManagerOwnershipTransferredIterator{contract: _SaleManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SaleManager *SaleManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SaleManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SaleManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleManagerOwnershipTransferred)
				if err := _SaleManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
