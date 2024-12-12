// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cryft-labs/cryftgo/vms/components/cryft (interfaces: TransferableOut)
//
// Generated by this command:
//
//	mockgen -package=cryft -destination=vms/components/cryft/mock_transferable_out.go github.com/cryft-labs/cryftgo/vms/components/cryft TransferableOut
//

// Package cryft is a generated GoMock package.
package cryft

import (
	reflect "reflect"

	snow "github.com/cryft-labs/cryftgo/snow"
	verify "github.com/cryft-labs/cryftgo/vms/components/verify"
	gomock "go.uber.org/mock/gomock"
)

// MockTransferableOut is a mock of TransferableOut interface.
type MockTransferableOut struct {
	verify.IsState

	ctrl     *gomock.Controller
	recorder *MockTransferableOutMockRecorder
}

// MockTransferableOutMockRecorder is the mock recorder for MockTransferableOut.
type MockTransferableOutMockRecorder struct {
	mock *MockTransferableOut
}

// NewMockTransferableOut creates a new mock instance.
func NewMockTransferableOut(ctrl *gomock.Controller) *MockTransferableOut {
	mock := &MockTransferableOut{ctrl: ctrl}
	mock.recorder = &MockTransferableOutMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferableOut) EXPECT() *MockTransferableOutMockRecorder {
	return m.recorder
}

// Amount mocks base method.
func (m *MockTransferableOut) Amount() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Amount")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Amount indicates an expected call of Amount.
func (mr *MockTransferableOutMockRecorder) Amount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Amount", reflect.TypeOf((*MockTransferableOut)(nil).Amount))
}

// InitCtx mocks base method.
func (m *MockTransferableOut) InitCtx(arg0 *snow.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCtx", arg0)
}

// InitCtx indicates an expected call of InitCtx.
func (mr *MockTransferableOutMockRecorder) InitCtx(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCtx", reflect.TypeOf((*MockTransferableOut)(nil).InitCtx), arg0)
}

// Verify mocks base method.
func (m *MockTransferableOut) Verify() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify")
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockTransferableOutMockRecorder) Verify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTransferableOut)(nil).Verify))
}

// isState mocks base method.
func (m *MockTransferableOut) isState() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isState")
}

// isState indicates an expected call of isState.
func (mr *MockTransferableOutMockRecorder) isState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isState", reflect.TypeOf((*MockTransferableOut)(nil).isState))
}
