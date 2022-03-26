// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/spacesvm/chain (interfaces: UnsignedTransaction)

package chain

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	tdata "github.com/ava-labs/spacesvm/tdata"
	gomock "github.com/golang/mock/gomock"
)

// MockUnsignedTransaction is a mock of UnsignedTransaction interface.
type MockUnsignedTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockUnsignedTransactionMockRecorder
}

// MockUnsignedTransactionMockRecorder is the mock recorder for MockUnsignedTransaction.
type MockUnsignedTransactionMockRecorder struct {
	mock *MockUnsignedTransaction
}

// NewMockUnsignedTransaction creates a new mock instance.
func NewMockUnsignedTransaction(ctrl *gomock.Controller) *MockUnsignedTransaction {
	mock := &MockUnsignedTransaction{ctrl: ctrl}
	mock.recorder = &MockUnsignedTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsignedTransaction) EXPECT() *MockUnsignedTransactionMockRecorder {
	return m.recorder
}

// Activity mocks base method.
func (m *MockUnsignedTransaction) Activity() *Activity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activity")
	ret0, _ := ret[0].(*Activity)
	return ret0
}

// Activity indicates an expected call of Activity.
func (mr *MockUnsignedTransactionMockRecorder) Activity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activity", reflect.TypeOf((*MockUnsignedTransaction)(nil).Activity))
}

// Copy mocks base method.
func (m *MockUnsignedTransaction) Copy() UnsignedTransaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(UnsignedTransaction)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockUnsignedTransactionMockRecorder) Copy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockUnsignedTransaction)(nil).Copy))
}

// Execute mocks base method.
func (m *MockUnsignedTransaction) Execute(arg0 *TransactionContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockUnsignedTransactionMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUnsignedTransaction)(nil).Execute), arg0)
}

// ExecuteBase mocks base method.
func (m *MockUnsignedTransaction) ExecuteBase(arg0 *Genesis) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteBase", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteBase indicates an expected call of ExecuteBase.
func (mr *MockUnsignedTransactionMockRecorder) ExecuteBase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteBase", reflect.TypeOf((*MockUnsignedTransaction)(nil).ExecuteBase), arg0)
}

// FeeUnits mocks base method.
func (m *MockUnsignedTransaction) FeeUnits(arg0 *Genesis) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeeUnits", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// FeeUnits indicates an expected call of FeeUnits.
func (mr *MockUnsignedTransactionMockRecorder) FeeUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeeUnits", reflect.TypeOf((*MockUnsignedTransaction)(nil).FeeUnits), arg0)
}

// GetBlockID mocks base method.
func (m *MockUnsignedTransaction) GetBlockID() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockID")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// GetBlockID indicates an expected call of GetBlockID.
func (mr *MockUnsignedTransactionMockRecorder) GetBlockID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockID", reflect.TypeOf((*MockUnsignedTransaction)(nil).GetBlockID))
}

// GetMagic mocks base method.
func (m *MockUnsignedTransaction) GetMagic() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMagic")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetMagic indicates an expected call of GetMagic.
func (mr *MockUnsignedTransactionMockRecorder) GetMagic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMagic", reflect.TypeOf((*MockUnsignedTransaction)(nil).GetMagic))
}

// GetPrice mocks base method.
func (m *MockUnsignedTransaction) GetPrice() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockUnsignedTransactionMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockUnsignedTransaction)(nil).GetPrice))
}

// LoadUnits mocks base method.
func (m *MockUnsignedTransaction) LoadUnits(arg0 *Genesis) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUnits", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LoadUnits indicates an expected call of LoadUnits.
func (mr *MockUnsignedTransactionMockRecorder) LoadUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUnits", reflect.TypeOf((*MockUnsignedTransaction)(nil).LoadUnits), arg0)
}

// SetBlockID mocks base method.
func (m *MockUnsignedTransaction) SetBlockID(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBlockID", arg0)
}

// SetBlockID indicates an expected call of SetBlockID.
func (mr *MockUnsignedTransactionMockRecorder) SetBlockID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBlockID", reflect.TypeOf((*MockUnsignedTransaction)(nil).SetBlockID), arg0)
}

// SetMagic mocks base method.
func (m *MockUnsignedTransaction) SetMagic(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMagic", arg0)
}

// SetMagic indicates an expected call of SetMagic.
func (mr *MockUnsignedTransactionMockRecorder) SetMagic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMagic", reflect.TypeOf((*MockUnsignedTransaction)(nil).SetMagic), arg0)
}

// SetPrice mocks base method.
func (m *MockUnsignedTransaction) SetPrice(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPrice", arg0)
}

// SetPrice indicates an expected call of SetPrice.
func (mr *MockUnsignedTransactionMockRecorder) SetPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrice", reflect.TypeOf((*MockUnsignedTransaction)(nil).SetPrice), arg0)
}

// TypedData mocks base method.
func (m *MockUnsignedTransaction) TypedData() *tdata.TypedData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TypedData")
	ret0, _ := ret[0].(*tdata.TypedData)
	return ret0
}

// TypedData indicates an expected call of TypedData.
func (mr *MockUnsignedTransactionMockRecorder) TypedData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TypedData", reflect.TypeOf((*MockUnsignedTransaction)(nil).TypedData))
}
