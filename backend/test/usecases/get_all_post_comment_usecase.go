package usecases

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type getAllPostCommentInput struct {
	postId int64
	limit  int64
	offset int64
}

type responceFromCommentRepositoryGetAll struct {
	comments []*entities.Comment
	err      error
}

func TestGetAllPostComment(t *testing.T) {
	testcases := []struct {
		testName         string
		usecaseInput     *getAllPostCommentInput
		repositoryOutput *responceFromCommentRepositoryGetAll
		expectedComment  []*entities.Comment
		expectedError    error
	}{
		{
			"Success get all posts",
			&getAllPostCommentInput{
				postId: 1,
				limit:  1,
				offset: 0,
			},
			&responceFromCommentRepositoryGetAll{
				[]*entities.Comment{
					{
						Id:        1,
						PostId:    1,
						UserId:    1,
						Body:      "mock comment",
						CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				nil,
			},
			[]*entities.Comment{
				{
					Id:        1,
					PostId:    1,
					UserId:    1,
					Body:      "mock comment",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			nil,
		},
		{
			"Success get all comments with empty posts",
			&getAllPostCommentInput{
				postId: 1,
				limit:  1,
				offset: 0,
			},
			&responceFromCommentRepositoryGetAll{
				[]*entities.Comment{},
				nil,
			},
			[]*entities.Comment{},
			nil,
		},
		{
			"Fail with error from GetAllCommentByPostId",
			&getAllPostCommentInput{
				1,
				1,
				0,
			},
			&responceFromCommentRepositoryGetAll{
				nil,
				fmt.Errorf("error from GetAllCommentByPostId"),
			},
			nil,
			fmt.Errorf("error from GetAllCommentByPostId"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCommentRepository := mock_interfaces.NewMockCommentRepository(mockCtrl)
			mockCommentRepository.EXPECT().
				GetAllByPostId(tc.usecaseInput.postId, tc.usecaseInput.limit, tc.usecaseInput.offset).
				Return(tc.repositoryOutput.comments, tc.repositoryOutput.err)
			usecase := usecases.NewGetAllPostCommentUsecase(mockCommentRepository)
			comments, err := usecase.Execute(tc.usecaseInput.postId, tc.usecaseInput.limit, tc.usecaseInput.offset)

			assert.Equal(t, tc.expectedError, err)
			assert.Len(t, comments, len(tc.expectedComment))
			for i, comment := range comments {
				assert.Equal(t, tc.expectedComment[i], comment)
			}
		})
	}
}
