package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type getCommentRepositoryOutput struct {
	comment *entities.Comment
	err     error
}

func TestGetComment(t *testing.T) {
	testcases := []struct {
		testName         string
		commentId        int64
		repositoryOutput *getCommentRepositoryOutput
		expectedComment  *entities.Comment
		expectedError    error
	}{
		{
			"Success",
			1,
			&getCommentRepositoryOutput{
				&entities.Comment{
					Id:        1,
					PostId:    1,
					UserId:    1,
					Body:      "mock comment",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				nil,
			},
			&entities.Comment{
				Id:        1,
				PostId:    1,
				UserId:    1,
				Body:      "mock comment",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			nil,
		},
		{
			"Success with unexisting comment (return nil)",
			100,
			&getCommentRepositoryOutput{
				nil,
				nil,
			},
			nil,
			exception.ErrNotFound,
		},
		{
			"Fails with unknown repository error",
			1,
			&getCommentRepositoryOutput{
				nil,
				errors.New("unknown error from repository"),
			},
			nil,
			errors.New("unknown error from repository"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCommentRepository := mock_interfaces.NewMockCommentRepository(mockCtrl)
			mockCommentRepository.EXPECT().
				GetById(tc.commentId).
				Return(tc.repositoryOutput.comment, tc.repositoryOutput.err)

			usecase := usecases.NewGetCommentUsecase(mockCommentRepository)
			comment, err := usecase.Execute(tc.commentId)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}
