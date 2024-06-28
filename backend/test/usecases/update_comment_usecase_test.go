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

type updateCommentUsecaseInput struct {
	userId    int64
	postId    int64
	commentId int64
	body      string
}

type responseFromCommentRepositoryGetById struct {
	comment *entities.Comment
	err     error
}

type responseFromCommentRepositoryUpdate struct {
	comment *entities.Comment
	err     error
}

func TestUpdateComment(t *testing.T) {
	testcases := []struct {
		testName            string
		input               *updateCommentUsecaseInput
		responseFromGetById *responseFromCommentRepositoryGetById
		responseFromUpdate  *responseFromCommentRepositoryUpdate
		expectedComment     *entities.Comment
		expectedError       error
	}{
		{
			"Success",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Body:   "old body",
				},
				nil,
			},
			&responseFromCommentRepositoryUpdate{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Body:   "new body",
				},
				nil,
			},
			&entities.Comment{
				Id:     1,
				UserId: 1,
				PostId: 1,
				Body:   "new body",
			},
			nil,
		},
		{
			"Fail with comment not found",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				nil,
				nil,
			},
			&responseFromCommentRepositoryUpdate{
				nil,
				nil,
			},
			nil,
			exception.ErrCommentNotFound,
		},
		{
			"Fail with error from GetById",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				nil,
				errors.New("error from GetById"),
			},
			&responseFromCommentRepositoryUpdate{
				nil,
				nil,
			},
			nil,
			errors.New("error from GetById"),
		},
		{
			"Fail with invalid user",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 2,
					PostId: 1,
					Body:   "old body",
				},
				nil,
			},
			&responseFromCommentRepositoryUpdate{
				nil,
				nil,
			},
			nil,
			exception.ErrUnauthorizedToUpdateComment,
		},
		{
			"Fail with invalid post_id",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 2,
					Body:   "old body",
				},
				nil,
			},
			&responseFromCommentRepositoryUpdate{
				nil,
				nil,
			},
			nil,
			exception.ErrInvalidPostId,
		},
		{
			"Fail with error from Update",
			&updateCommentUsecaseInput{
				1,
				1,
				1,
				"new body",
			},
			&responseFromCommentRepositoryGetById{
				&entities.Comment{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Body:   "old body",
				},
				nil,
			},
			&responseFromCommentRepositoryUpdate{
				nil,
				errors.New("error from Update"),
			},
			nil,
			errors.New("error from Update"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCommentRepository := mock_interfaces.NewMockCommentRepository(mockCtrl)
			mockCommentRepository.EXPECT().
				GetById(tc.input.commentId).
				Return(tc.responseFromGetById.comment, tc.responseFromGetById.err)

			newComment := &entities.Comment{}
			if tc.responseFromGetById.comment != nil {
				newComment = &entities.Comment{
					Id:     tc.responseFromGetById.comment.Id,
					UserId: tc.responseFromGetById.comment.UserId,
					Body:   tc.input.body,
				}
			}
			mockCommentRepository.EXPECT().
				Update(newComment).
				Return(tc.responseFromUpdate.comment, tc.responseFromUpdate.err).
				AnyTimes()

			usecase := usecases.NewUpdateCommentUsecase(mockCommentRepository)
			comment, err := usecase.Execute(tc.input.userId, tc.input.postId, tc.input.commentId, tc.input.body)
			assert.Equal(t, tc.expectedError, err)
			if comment != nil {
				assert.Equal(t, tc.expectedComment.Id, comment.Id)
				assert.Equal(t, tc.expectedComment.UserId, comment.UserId)
				assert.Equal(t, tc.expectedComment.Body, comment.Body)
				// Do not compare the UpdatedAt and DeletedAt fields
			}
		})
	}
}
