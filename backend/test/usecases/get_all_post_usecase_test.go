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

type getAllPostsInput struct {
	limit  int64
	offset int64
}

type responseFromPostRepositoryGetAll struct {
	posts []*entities.Post
	err   error
}

func TestGetAllPosts(t *testing.T) {
	testcases := []struct {
		testName           string
		input              *getAllPostsInput
		responseFromGetAll *responseFromPostRepositoryGetAll
		expectedPosts      []*entities.Post
		expectedError      error
	}{
		{
			"Success get all posts",
			&getAllPostsInput{
				1,
				0,
			},
			&responseFromPostRepositoryGetAll{
				[]*entities.Post{
					{
						Id:        1,
						Title:     "Mock",
						Body:      "Mockやで",
						UserId:    2,
						Username:  "Mocker",
						CreatedAt: time.Date(2024, 5, 28, 13, 52, 55, 0, time.Local),
						UpdatedAt: time.Date(2024, 5, 28, 13, 52, 55, 0, time.Local),
					},
				},
				nil,
			},
			[]*entities.Post{
				{
					Id:        1,
					Title:     "Mock",
					Body:      "Mockやで",
					UserId:    2,
					Username:  "Mocker",
					CreatedAt: time.Date(2024, 5, 28, 13, 52, 55, 0, time.Local),
					UpdatedAt: time.Date(2024, 5, 28, 13, 52, 55, 0, time.Local),
				},
			},
			nil,
		},
		// TODO: Add test cases for limit and offset
		{
			"Fail with error from GetAll",
			&getAllPostsInput{
				1,
				0,
			},
			&responseFromPostRepositoryGetAll{
				nil,
				fmt.Errorf("error from GetAll"),
			},
			nil,
			fmt.Errorf("error from GetAll"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostRepository := mock_interfaces.NewMockPostRepository(mockCtrl)
			mockPostRepository.EXPECT().GetAll(
				tc.input.limit,
				tc.input.offset,
			).Return(
				tc.responseFromGetAll.posts,
				tc.responseFromGetAll.err,
			)

			usecase := usecases.NewGetAllPostsUsecase(mockPostRepository)
			posts, err := usecase.Execute(tc.input.limit, tc.input.offset)
			assert.Equal(t, tc.expectedError, err)
			assert.Len(t, posts, len(tc.expectedPosts))
			for i := range posts {
				assert.Equal(t, tc.expectedPosts[i], posts[i])
			}
		})
	}
}
