// Code generated by MockGen. DO NOT EDIT.
// Source: driven.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	todo "github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
	events "github.com/anuragaryan/ddd-with-hex-go/internal/application/events"
	gomock "github.com/golang/mock/gomock"
)

// MockStoragePort is a mock of StoragePort interface.
type MockStoragePort struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePortMockRecorder
}

// MockStoragePortMockRecorder is the mock recorder for MockStoragePort.
type MockStoragePortMockRecorder struct {
	mock *MockStoragePort
}

// NewMockStoragePort creates a new mock instance.
func NewMockStoragePort(ctrl *gomock.Controller) *MockStoragePort {
	mock := &MockStoragePort{ctrl: ctrl}
	mock.recorder = &MockStoragePortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePort) EXPECT() *MockStoragePortMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStoragePort) Add(list todo.List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", list)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStoragePortMockRecorder) Add(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStoragePort)(nil).Add), list)
}

// AddItem mocks base method.
func (m *MockStoragePort) AddItem(id string, item todo.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", id, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddItem indicates an expected call of AddItem.
func (mr *MockStoragePortMockRecorder) AddItem(id, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockStoragePort)(nil).AddItem), id, item)
}

// Delete mocks base method.
func (m *MockStoragePort) Delete(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", id)
}

// Delete indicates an expected call of Delete.
func (mr *MockStoragePortMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStoragePort)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockStoragePort) GetAll() ([]todo.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]todo.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStoragePortMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStoragePort)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockStoragePort) GetByID(id string) (todo.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(todo.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockStoragePortMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockStoragePort)(nil).GetByID), id)
}

// ListItem mocks base method.
func (m *MockStoragePort) ListItem(id string) ([]todo.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItem", id)
	ret0, _ := ret[0].([]todo.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItem indicates an expected call of ListItem.
func (mr *MockStoragePortMockRecorder) ListItem(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItem", reflect.TypeOf((*MockStoragePort)(nil).ListItem), id)
}

// MarkItemDone mocks base method.
func (m *MockStoragePort) MarkItemDone(id, itemID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkItemDone", id, itemID)
}

// MarkItemDone indicates an expected call of MarkItemDone.
func (mr *MockStoragePortMockRecorder) MarkItemDone(id, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkItemDone", reflect.TypeOf((*MockStoragePort)(nil).MarkItemDone), id, itemID)
}

// MockEventHandlerPort is a mock of EventHandlerPort interface.
type MockEventHandlerPort struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerPortMockRecorder
}

// MockEventHandlerPortMockRecorder is the mock recorder for MockEventHandlerPort.
type MockEventHandlerPortMockRecorder struct {
	mock *MockEventHandlerPort
}

// NewMockEventHandlerPort creates a new mock instance.
func NewMockEventHandlerPort(ctrl *gomock.Controller) *MockEventHandlerPort {
	mock := &MockEventHandlerPort{ctrl: ctrl}
	mock.recorder = &MockEventHandlerPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandlerPort) EXPECT() *MockEventHandlerPortMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockEventHandlerPort) Notify(event events.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockEventHandlerPortMockRecorder) Notify(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockEventHandlerPort)(nil).Notify), event)
}

// Subscribe mocks base method.
func (m *MockEventHandlerPort) Subscribe(handler events.EventHandler, events ...events.Event) {
	m.ctrl.T.Helper()
	varargs := []interface{}{handler}
	for _, a := range events {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Subscribe", varargs...)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockEventHandlerPortMockRecorder) Subscribe(handler interface{}, events ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{handler}, events...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventHandlerPort)(nil).Subscribe), varargs...)
}
