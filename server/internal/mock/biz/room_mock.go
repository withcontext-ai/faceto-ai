// Code generated by MockGen. DO NOT EDIT.
// Source: room.go

// Package biz is a generated GoMock package.
package biz

import (
	context "context"
	biz "faceto-ai/internal/biz"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoomRepo is a mock of RoomRepo interface.
type MockRoomRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepoMockRecorder
}

// MockRoomRepoMockRecorder is the mock recorder for MockRoomRepo.
type MockRoomRepoMockRecorder struct {
	mock *MockRoomRepo
}

// NewMockRoomRepo creates a new mock instance.
func NewMockRoomRepo(ctrl *gomock.Controller) *MockRoomRepo {
	mock := &MockRoomRepo{ctrl: ctrl}
	mock.recorder = &MockRoomRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepo) EXPECT() *MockRoomRepoMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockRoomRepo) GetByID(ctx context.Context, ID uint64) (*biz.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ID)
	ret0, _ := ret[0].(*biz.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRoomRepoMockRecorder) GetByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRoomRepo)(nil).GetByID), ctx, ID)
}

// GetByName mocks base method.
func (m *MockRoomRepo) GetByName(ctx context.Context, name string) (*biz.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*biz.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockRoomRepoMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRoomRepo)(nil).GetByName), ctx, name)
}

// GetBySID mocks base method.
func (m *MockRoomRepo) GetBySID(ctx context.Context, sid string) (*biz.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySID", ctx, sid)
	ret0, _ := ret[0].(*biz.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySID indicates an expected call of GetBySID.
func (mr *MockRoomRepoMockRecorder) GetBySID(ctx, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySID", reflect.TypeOf((*MockRoomRepo)(nil).GetBySID), ctx, sid)
}

// GetByUUID mocks base method.
func (m *MockRoomRepo) GetByUUID(ctx context.Context, uuid string) (*biz.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", ctx, uuid)
	ret0, _ := ret[0].(*biz.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockRoomRepoMockRecorder) GetByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockRoomRepo)(nil).GetByUUID), ctx, uuid)
}

// List mocks base method.
func (m *MockRoomRepo) List(ctx context.Context, page uint) ([]*biz.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, page)
	ret0, _ := ret[0].([]*biz.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoomRepoMockRecorder) List(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoomRepo)(nil).List), ctx, page)
}

// Save mocks base method.
func (m *MockRoomRepo) Save(ctx context.Context, room *biz.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRoomRepoMockRecorder) Save(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRoomRepo)(nil).Save), ctx, room)
}

// UpdateStatus mocks base method.
func (m *MockRoomRepo) UpdateStatus(ctx context.Context, sid string, fromStatus []uint8, toStatus uint8, up *biz.Room) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, sid, fromStatus, toStatus, up)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockRoomRepoMockRecorder) UpdateStatus(ctx, sid, fromStatus, toStatus, up interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockRoomRepo)(nil).UpdateStatus), ctx, sid, fromStatus, toStatus, up)
}
