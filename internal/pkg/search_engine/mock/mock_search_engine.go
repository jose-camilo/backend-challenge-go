// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jose-camilo/backend-challenge-go/internal/pkg/search_engine (interfaces: SearchEngine)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/jose-camilo/backend-challenge-go/internal/model"
	reflect "reflect"
)

// MockSearchEngine is a mock of SearchEngine interface
type MockSearchEngine struct {
	ctrl     *gomock.Controller
	recorder *MockSearchEngineMockRecorder
}

// MockSearchEngineMockRecorder is the mock recorder for MockSearchEngine
type MockSearchEngineMockRecorder struct {
	mock *MockSearchEngine
}

// NewMockSearchEngine creates a new mock instance
func NewMockSearchEngine(ctrl *gomock.Controller) *MockSearchEngine {
	mock := &MockSearchEngine{ctrl: ctrl}
	mock.recorder = &MockSearchEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchEngine) EXPECT() *MockSearchEngineMockRecorder {
	return m.recorder
}

// Index mocks base method
func (m *MockSearchEngine) Index(arg0 context.Context, arg1 model.TokenDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Index indicates an expected call of Index
func (mr *MockSearchEngineMockRecorder) Index(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockSearchEngine)(nil).Index), arg0, arg1)
}

// Search mocks base method
func (m *MockSearchEngine) Search(arg0 context.Context, arg1, arg2, arg3 string) ([]model.TokenDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]model.TokenDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockSearchEngineMockRecorder) Search(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearchEngine)(nil).Search), arg0, arg1, arg2, arg3)
}
