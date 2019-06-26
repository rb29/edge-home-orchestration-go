// Code generated by MockGen. DO NOT EDIT.
// Source: common/networkhelper/detector (interfaces: Detector)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDetector is a mock of Detector interface
type MockDetector struct {
	ctrl     *gomock.Controller
	recorder *MockDetectorMockRecorder
}

// MockDetectorMockRecorder is the mock recorder for MockDetector
type MockDetectorMockRecorder struct {
	mock *MockDetector
}

// NewMockDetector creates a new mock instance
func NewMockDetector(ctrl *gomock.Controller) *MockDetector {
	mock := &MockDetector{ctrl: ctrl}
	mock.recorder = &MockDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDetector) EXPECT() *MockDetectorMockRecorder {
	return m.recorder
}

// AddrSubscribe mocks base method
func (m *MockDetector) AddrSubscribe(arg0 chan<- bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddrSubscribe", arg0)
}

// AddrSubscribe indicates an expected call of AddrSubscribe
func (mr *MockDetectorMockRecorder) AddrSubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrSubscribe", reflect.TypeOf((*MockDetector)(nil).AddrSubscribe), arg0)
}
