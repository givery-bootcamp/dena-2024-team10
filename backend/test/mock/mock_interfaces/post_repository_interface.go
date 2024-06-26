// Code generated by MockGen. DO NOT EDIT.
// Source: post_repository_interface.go
//
// Generated by this command:
//
//	mockgen -source=post_repository_interface.go -package=mock_interfaces -destination=../../test/mock/mock_interfaces/post_repository_interface.go
//

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	entities "myapp/internal/entities"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostRepository) CreatePost(title, body string, userId int64) (*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", title, body, userId)
	ret0, _ := ret[0].(*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostRepositoryMockRecorder) CreatePost(title, body, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostRepository)(nil).CreatePost), title, body, userId)
}

// Delete mocks base method.
func (m *MockPostRepository) Delete(postId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", postId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostRepositoryMockRecorder) Delete(postId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostRepository)(nil).Delete), postId)
}

// GetAll mocks base method.
func (m *MockPostRepository) GetAll(limit, offset int64) ([]*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", limit, offset)
	ret0, _ := ret[0].([]*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPostRepositoryMockRecorder) GetAll(limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPostRepository)(nil).GetAll), limit, offset)
}

// GetById mocks base method.
func (m *MockPostRepository) GetById(postId int64) (*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", postId)
	ret0, _ := ret[0].(*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPostRepositoryMockRecorder) GetById(postId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPostRepository)(nil).GetById), postId)
}

// UpdatePost mocks base method.
func (m *MockPostRepository) UpdatePost(title, body string, postId int64) (*entities.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", title, body, postId)
	ret0, _ := ret[0].(*entities.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRepositoryMockRecorder) UpdatePost(title, body, postId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRepository)(nil).UpdatePost), title, body, postId)
}
