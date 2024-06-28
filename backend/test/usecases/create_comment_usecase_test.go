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

type createCommentUsecaseInput struct {
	postId int64
	body   string
	userId int64
}

type responseFromCommentRepositoryCreate struct {
	comment *entities.Comment
	err     error
}

func TestCreateComennt(t *testing.T) {
	testcases := []struct {
		testName           string
		input              *createCommentUsecaseInput
		responseFromCreate *responseFromCommentRepositoryCreate
		expectedComment    *entities.Comment
		expectedError      error
	}{
		{
			"Success",
			&createCommentUsecaseInput{
				1,
				"body",
				1,
			},
			&responseFromCommentRepositoryCreate{
				&entities.Comment{
					Id:        1,
					PostId:    1,
					Body:      "body",
					UserId:    1,
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				nil,
			},
			&entities.Comment{
				Id:        1,
				PostId:    1,
				Body:      "body",
				UserId:    1,
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			nil,
		},
		{
			"Fail with post not found",
			&createCommentUsecaseInput{
				1,
				"body",
				1,
			},
			&responseFromCommentRepositoryCreate{
				nil,
				errors.New("post or user not found"),
			},
			nil,
			exception.ErrPostNotFound,
		},
		{
			"Fail with repository error",
			&createCommentUsecaseInput{
				1,
				"body",
				1,
			},
			&responseFromCommentRepositoryCreate{
				nil,
				errors.New("repository error"),
			},
			nil,
			errors.New("repository error"),
		},
	}

	for _, tc := range testcases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCommentRepository := mock_interfaces.NewMockCommentRepository(mockCtrl)
		mockCommentRepository.EXPECT().Create(tc.input.postId, tc.input.body, tc.input.userId).Return(tc.responseFromCreate.comment, tc.responseFromCreate.err)

		usecase := usecases.NewCreateCommentUsecase(mockCommentRepository)
		comment, err := usecase.Execute(tc.input.postId, tc.input.body, tc.input.userId)
		assert.Equal(t, tc.expectedComment, comment)
		assert.Equal(t, tc.expectedError, err)
	}
}
