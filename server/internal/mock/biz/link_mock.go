// Code generated by MockGen. DO NOT EDIT.
// Source: link.go

// Package biz is a generated GoMock package.
package biz

import (
	context "context"
	biz "faceto-ai/internal/biz"
	schema "faceto-ai/internal/data/schema"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLinkRepo is a mock of LinkRepo interface.
type MockLinkRepo struct {
	ctrl     *gomock.Controller
	recorder *MockLinkRepoMockRecorder
}

// MockLinkRepoMockRecorder is the mock recorder for MockLinkRepo.
type MockLinkRepoMockRecorder struct {
	mock *MockLinkRepo
}

// NewMockLinkRepo creates a new mock instance.
func NewMockLinkRepo(ctrl *gomock.Controller) *MockLinkRepo {
	mock := &MockLinkRepo{ctrl: ctrl}
	mock.recorder = &MockLinkRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkRepo) EXPECT() *MockLinkRepoMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockLinkRepo) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockLinkRepoMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockLinkRepo)(nil).Count), ctx)
}

// GetLinkByName mocks base method.
func (m *MockLinkRepo) GetLinkByName(ctx context.Context, roomName string) (*biz.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinkByName", ctx, roomName)
	ret0, _ := ret[0].(*biz.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLinkByName indicates an expected call of GetLinkByName.
func (mr *MockLinkRepoMockRecorder) GetLinkByName(ctx, roomName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinkByName", reflect.TypeOf((*MockLinkRepo)(nil).GetLinkByName), ctx, roomName)
}

// Save mocks base method.
func (m *MockLinkRepo) Save(ctx context.Context, view *biz.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, view)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockLinkRepoMockRecorder) Save(ctx, view interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockLinkRepo)(nil).Save), ctx, view)
}

// SetConfigByUUID mocks base method.
func (m *MockLinkRepo) SetConfigByUUID(ctx context.Context, uuid string, config *schema.RoomConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfigByUUID", ctx, uuid, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfigByUUID indicates an expected call of SetConfigByUUID.
func (mr *MockLinkRepoMockRecorder) SetConfigByUUID(ctx, uuid, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigByUUID", reflect.TypeOf((*MockLinkRepo)(nil).SetConfigByUUID), ctx, uuid, config)
}

// SetRoomVoiceID mocks base method.
func (m *MockLinkRepo) SetRoomVoiceID(ctx context.Context, roomName, voiceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRoomVoiceID", ctx, roomName, voiceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRoomVoiceID indicates an expected call of SetRoomVoiceID.
func (mr *MockLinkRepoMockRecorder) SetRoomVoiceID(ctx, roomName, voiceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoomVoiceID", reflect.TypeOf((*MockLinkRepo)(nil).SetRoomVoiceID), ctx, roomName, voiceID)
}
