// Code generated by MockGen. DO NOT EDIT.
// Source: room_vod.go

// Package biz is a generated GoMock package.
package biz

import (
	context "context"
	biz "faceto-ai/internal/biz"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoomVodRepo is a mock of RoomVodRepo interface.
type MockRoomVodRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRoomVodRepoMockRecorder
}

// MockRoomVodRepoMockRecorder is the mock recorder for MockRoomVodRepo.
type MockRoomVodRepoMockRecorder struct {
	mock *MockRoomVodRepo
}

// NewMockRoomVodRepo creates a new mock instance.
func NewMockRoomVodRepo(ctrl *gomock.Controller) *MockRoomVodRepo {
	mock := &MockRoomVodRepo{ctrl: ctrl}
	mock.recorder = &MockRoomVodRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomVodRepo) EXPECT() *MockRoomVodRepoMockRecorder {
	return m.recorder
}

// GetByEgressID mocks base method.
func (m *MockRoomVodRepo) GetByEgressID(ctx context.Context, egressID string) (*biz.RoomVod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEgressID", ctx, egressID)
	ret0, _ := ret[0].(*biz.RoomVod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEgressID indicates an expected call of GetByEgressID.
func (mr *MockRoomVodRepoMockRecorder) GetByEgressID(ctx, egressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEgressID", reflect.TypeOf((*MockRoomVodRepo)(nil).GetByEgressID), ctx, egressID)
}

// GetBySid mocks base method.
func (m *MockRoomVodRepo) GetBySid(ctx context.Context, sid string) (*biz.RoomVod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySid", ctx, sid)
	ret0, _ := ret[0].(*biz.RoomVod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySid indicates an expected call of GetBySid.
func (mr *MockRoomVodRepoMockRecorder) GetBySid(ctx, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySid", reflect.TypeOf((*MockRoomVodRepo)(nil).GetBySid), ctx, sid)
}

// Save mocks base method.
func (m *MockRoomVodRepo) Save(ctx context.Context, vod *biz.RoomVod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, vod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRoomVodRepoMockRecorder) Save(ctx, vod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRoomVodRepo)(nil).Save), ctx, vod)
}

// UpdateStatus mocks base method.
func (m *MockRoomVodRepo) UpdateStatus(ctx context.Context, egressID string, fromStatus []uint8, toStatus uint8, vod *biz.RoomVod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, egressID, fromStatus, toStatus, vod)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockRoomVodRepoMockRecorder) UpdateStatus(ctx, egressID, fromStatus, toStatus, vod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockRoomVodRepo)(nil).UpdateStatus), ctx, egressID, fromStatus, toStatus, vod)
}
