// Code generated by MockGen. DO NOT EDIT.
// Source: comment_repository_interface.go
//
// Generated by this command:
//
//	mockgen -source=comment_repository_interface.go -package=mock_interfaces -destination=../../test/mock/mock_interfaces/comment_repository_interface.go
//

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	entities "myapp/internal/entities"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommentRepository) Create(postId int64, body string, userId int64) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", postId, body, userId)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentRepositoryMockRecorder) Create(postId, body, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentRepository)(nil).Create), postId, body, userId)
}

// Delete mocks base method.
func (m *MockCommentRepository) Delete(commentId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", commentId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentRepositoryMockRecorder) Delete(commentId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentRepository)(nil).Delete), commentId)
}

// GetById mocks base method.
func (m *MockCommentRepository) GetById(commentId int64) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", commentId)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCommentRepositoryMockRecorder) GetById(commentId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCommentRepository)(nil).GetById), commentId)
}

// Update mocks base method.
func (m *MockCommentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", comment)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentRepositoryMockRecorder) Update(comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentRepository)(nil).Update), comment)
}
