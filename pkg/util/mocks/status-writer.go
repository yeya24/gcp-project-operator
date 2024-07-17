// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/controller-runtime/pkg/client (interfaces: StatusWriter)
//
// Generated by this command:
//
//	mockgen -destination ./status-writer.go -package mocks sigs.k8s.io/controller-runtime/pkg/client StatusWriter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockStatusWriter is a mock of StatusWriter interface.
type MockStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockStatusWriterMockRecorder
}

// MockStatusWriterMockRecorder is the mock recorder for MockStatusWriter.
type MockStatusWriterMockRecorder struct {
	mock *MockStatusWriter
}

// NewMockStatusWriter creates a new mock instance.
func NewMockStatusWriter(ctrl *gomock.Controller) *MockStatusWriter {
	mock := &MockStatusWriter{ctrl: ctrl}
	mock.recorder = &MockStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusWriter) EXPECT() *MockStatusWriterMockRecorder {
	return m.recorder
}

// Patch mocks base method.
func (m *MockStatusWriter) Patch(arg0 context.Context, arg1 client.Object, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockStatusWriterMockRecorder) Patch(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockStatusWriter)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockStatusWriter) Update(arg0 context.Context, arg1 client.Object, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStatusWriterMockRecorder) Update(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatusWriter)(nil).Update), varargs...)
}
