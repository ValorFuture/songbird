package core

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ava-labs/coreth/core/vm"
	"github.com/ethereum/go-ethereum/common"
)

// Define a mock structure to spy and mock values for keeper calls
type MockEVMCallerData struct {
	callCalls            int
	addBalanceCalls      int
	blockNumber          big.Int
	gasLimit             uint64
	mintRequestReturn    big.Int
	lastAddBalanceAddr   common.Address
	lastAddBalanceAmount *big.Int
}

// Set up default mock method calls
func defautCall(e *MockEVMCallerData, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	e.callCalls++

	buffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return e.mintRequestReturn.FillBytes(buffer), 0, nil
}

func defaultGetBlockNumber(e *MockEVMCallerData) *big.Int {
	return &e.blockNumber
}

func defaultGetGasLimit(e *MockEVMCallerData) uint64 {
	return e.gasLimit
}

func defaultAddBalance(e *MockEVMCallerData, addr common.Address, amount *big.Int) {
	e.addBalanceCalls++
	e.lastAddBalanceAddr = addr
	e.lastAddBalanceAmount = amount
}

// Define the default EVM mock and define default mock receiver functions
type DefaultEVMMock struct {
	mockEVMCallerData MockEVMCallerData
}

func (e *DefaultEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	return defautCall(&e.mockEVMCallerData, caller, addr, input, gas, value)
}

func (e *DefaultEVMMock) GetBlockNumber() *big.Int {
	return defaultGetBlockNumber(&e.mockEVMCallerData)
}

func (e *DefaultEVMMock) GetGasLimit() uint64 {
	return defaultGetGasLimit(&e.mockEVMCallerData)
}

func (e *DefaultEVMMock) AddBalance(addr common.Address, amount *big.Int) {
	defaultAddBalance(&e.mockEVMCallerData, addr, amount)
}

func TestKeeperTriggerShouldReturnMintRequest(t *testing.T) {
	mintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *mintRequestReturn,
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	mintRequest, _ := triggerKeeper(defaultEVMMock)

	if mintRequest.Cmp(mintRequestReturn) != 0 {
		t.Errorf("got %s want %q", mintRequest.Text(10), "50000000000000000000000000")
	}
}

func TestKeeperTriggerShouldNotLetMintRequestOverflow(t *testing.T) {
	var mintRequestReturn big.Int
	// TODO: Compact with exponent?
	buffer := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	mintRequestReturn.SetBytes(buffer)

	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: mintRequestReturn,
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	mintRequest, mintRequestError := triggerKeeper(defaultEVMMock)

	if mintRequestError != nil {
		t.Errorf("received unexpected error %s", mintRequestError)
	}

	if mintRequest.Sign() < 1 {
		t.Errorf("unexpected negative")
	}
}

// Define a bad mint request return size mock
type BadMintReturnSizeEVMMock struct {
	mockEVMCallerData MockEVMCallerData
}

func (e *BadMintReturnSizeEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	e.mockEVMCallerData.callCalls++
	// Should be size 32 bytes
	buffer := []byte{0}
	return e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, nil
}

func (e *BadMintReturnSizeEVMMock) GetBlockNumber() *big.Int {
	return defaultGetBlockNumber(&e.mockEVMCallerData)
}

func (e *BadMintReturnSizeEVMMock) GetGasLimit() uint64 {
	return defaultGetGasLimit(&e.mockEVMCallerData)
}

func (e *BadMintReturnSizeEVMMock) AddBalance(addr common.Address, amount *big.Int) {
	defaultAddBalance(&e.mockEVMCallerData, addr, amount)
}

func TestKeeperTriggerValidatesMintRequestReturnValueSize(t *testing.T) {
	var mintRequestReturn big.Int
	// TODO: Compact with exponent?
	buffer := []byte{255}
	mintRequestReturn.SetBytes(buffer)

	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: mintRequestReturn,
	}
	badMintReturnSizeEVMMock := &BadMintReturnSizeEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}
	// Call to return less than 32 bytes
	_, err := triggerKeeper(badMintReturnSizeEVMMock)

	if err != nil {
		if err, ok := err.(*ErrInvalidKeeperData); !ok {
			want := &ErrInvalidKeeperData{}
			t.Errorf("got '%s' want '%s'", err.Error(), want.Error())
		}
	} else {
		t.Errorf("no error returned as expected")
	}
}

// Define a mock to simulate keeper trigger returning an error from Call
type BadTriggerCallEVMMock struct {
	mockEVMCallerData MockEVMCallerData
}

func (e *BadTriggerCallEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	e.mockEVMCallerData.callCalls++

	buffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, errors.New("Call error happened")
}

func (e *BadTriggerCallEVMMock) GetBlockNumber() *big.Int {
	return defaultGetBlockNumber(&e.mockEVMCallerData)
}

func (e *BadTriggerCallEVMMock) GetGasLimit() uint64 {
	return defaultGetGasLimit(&e.mockEVMCallerData)
}

func (e *BadTriggerCallEVMMock) AddBalance(addr common.Address, amount *big.Int) {
	defaultAddBalance(&e.mockEVMCallerData, addr, amount)
}

func TestKeeperTriggerReturnsCallError(t *testing.T) {
	mockEVMCallerData := &MockEVMCallerData{}
	badTriggerCallEVMMock := &BadTriggerCallEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}
	// Call to return less than 32 bytes
	_, err := triggerKeeper(badTriggerCallEVMMock)

	if err == nil {
		t.Errorf("no error received")
	} else {
		if err.Error() != "Call error happened" {
			t.Errorf("did not get expected error")
		}
	}
}

// Define a mock to simulate keeper trigger returning nil for mint request
type ReturnNilMintRequestEVMMock struct {
	mockEVMCallerData MockEVMCallerData
}

func (e *ReturnNilMintRequestEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	e.mockEVMCallerData.callCalls++

	return nil, 0, nil
}

func (e *ReturnNilMintRequestEVMMock) GetBlockNumber() *big.Int {
	return defaultGetBlockNumber(&e.mockEVMCallerData)
}

func (e *ReturnNilMintRequestEVMMock) GetGasLimit() uint64 {
	return defaultGetGasLimit(&e.mockEVMCallerData)
}

func (e *ReturnNilMintRequestEVMMock) AddBalance(addr common.Address, amount *big.Int) {
	defaultAddBalance(&e.mockEVMCallerData, addr, amount)
}

func TestKeeperTriggerHandlesNilMintRequest(t *testing.T) {
	mockEVMCallerData := &MockEVMCallerData{}
	returnNilMintRequestEVMMock := &ReturnNilMintRequestEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}
	// Call to return less than 32 bytes
	_, err := triggerKeeper(returnNilMintRequestEVMMock)

	if err != nil {
		if err, ok := err.(*ErrKeeperDataEmpty); !ok {
			want := &ErrKeeperDataEmpty{}
			t.Errorf("got '%s' want '%s'", err.Error(), want.Error())
		}
	} else {
		t.Errorf("no error returned as expected")
	}
}

func TestKeeperTriggerShouldNotMintMoreThanMax(t *testing.T) {
	mintRequest, _ := new(big.Int).SetString("50000000000000000000000001", 10)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *big.NewInt(0),
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	err := mint(defaultEVMMock, mintRequest)

	if err != nil {
		if err, ok := err.(*ErrMaxMintExceeded); !ok {
			want := &ErrMaxMintExceeded{
				mintRequest: mintRequest,
				mintMax:     GetMaximumMintRequest(big.NewInt(0)),
			}
			t.Errorf("got '%s' want '%s'", err.Error(), want.Error())
		}
	} else {
		t.Errorf("no error returned as expected")
	}
}

func TestKeeperTriggerShouldNotMintNegative(t *testing.T) {
	mintRequest := big.NewInt(-1)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *big.NewInt(0),
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	err := mint(defaultEVMMock, mintRequest)

	if err != nil {
		if err, ok := err.(*ErrMintNegative); !ok {
			want := &ErrMintNegative{}
			t.Errorf("got '%s' want '%s'", err.Error(), want.Error())
		}
	} else {
		t.Errorf("no error returned as expected")
	}
}

func TestKeeperTriggerShouldMint(t *testing.T) {
	// Assemble
	mintRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *big.NewInt(0),
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	// Act
	err := mint(defaultEVMMock, mintRequest)

	// Assert
	if err == nil {
		if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 1 {
			t.Errorf("AddBalance not called as expected")
		}
		if defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr.String() != GetSystemTriggerContractAddr(big.NewInt(0)) {
			t.Errorf("wanted addr %s; got addr %s", GetSystemTriggerContractAddr(big.NewInt(0)), defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr)
		}
		if defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Cmp(mintRequest) != 0 {
			t.Errorf("wanted amount %s; got amount %s", mintRequest.Text(10), defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Text(10))
		}
	} else {
		t.Errorf("unexpected error returned; was = %s", err.Error())
	}
}

func TestKeeperTriggerShouldNotErrorMintingZero(t *testing.T) {
	// Assemble
	mintRequest := big.NewInt(0)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *big.NewInt(0),
	}
	defaultEVMMock := &DefaultEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
	}

	// Act
	err := mint(defaultEVMMock, mintRequest)

	// Assert
	if err == nil {
		if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 0 {
			t.Errorf("AddBalance called unexpectedly")
		}
	} else {
		t.Errorf("unexpected error returned; was %s", err.Error())
	}
}
