package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type deleteCommentUsecaseInput struct {
	postId    int64
	commentId int64
	userId    int64
}

func TestDeleteComment(t *testing.T) {
	testcases := []struct {
		testName           string
		input              *deleteCommentUsecaseInput
		responseGetByID    *responseFromCommentRepositoryGetById
		responseFromDelete error
		expectedError      error
	}{
		{
			"Success",
			&deleteCommentUsecaseInput{
				1,
				1,
				1,
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Body:   "body",
				},
				nil,
			},
			nil,
			nil,
		},
		{
			"Fail with comment not found",
			&deleteCommentUsecaseInput{
				1,
				1,
				1,
			},
			&responseFromCommentRepositoryGetById{
				nil,
				nil,
			},
			nil,
			exception.ErrCommentNotFound,
		},
		{
			"Fail with unauthorized to delete comment",
			&deleteCommentUsecaseInput{
				1,
				1,
				1,
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 2,
					PostId: 1,
					Body:   "body",
				},
				nil,
			},
			nil,
			exception.ErrUnauthorizedToDeleteComment,
		},
		{
			"Fail with invalid post ID",
			&deleteCommentUsecaseInput{
				1,
				1,
				1,
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 2,
					Body:   "body",
				},
				nil,
			},
			nil,
			exception.ErrInvalidPostId,
		},
		{
			"Fail with error from Delete",
			&deleteCommentUsecaseInput{
				1,
				1,
				1,
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Body:   "body",
				},
				nil,
			},
			errors.New("error from Delete"),
			errors.New("error from Delete"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCommentRepository := mock_interfaces.NewMockCommentRepository(gomock.NewController(t))
			mockCommentRepository.EXPECT().
				GetById(tc.input.commentId).
				Return(tc.responseGetByID.comment, tc.responseGetByID.err)
			mockCommentRepository.EXPECT().
				Delete(tc.input.commentId).
				Return(tc.responseFromDelete).AnyTimes()

			// create usecase
			usecase := usecases.NewDeleteCommentUsecase(mockCommentRepository)
			err := usecase.Execute(tc.input.postId, tc.input.commentId, tc.input.userId)

			// assert
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
