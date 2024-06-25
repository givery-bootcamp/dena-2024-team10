package usecases

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type deletePostUsecaseInput struct {
	postId string
	userId int64
}

type responseFromPostRepositoryGetById struct {
	post *entities.Post
	err  error
}

func TestDeletePost(t *testing.T) {
	testcases := []struct {
		testName           string
		input              *deletePostUsecaseInput
		responseGetByID    *responseFromPostRepositoryGetById
		responseFromDelete error
		expected           error
	}{
		{
			"Success",
			&deletePostUsecaseInput{
				"test_post_id",
				1,
			},
			&responseFromPostRepositoryGetById{
				&entities.Post{
					UserId: 1,
				},
				nil,
			},
			nil,
			nil,
		},
		{
			"Fail with post not found",
			&deletePostUsecaseInput{
				"test_post_id",
				1,
			},
			&responseFromPostRepositoryGetById{
				nil,
				nil,
			},
			nil,
			exception.ErrNotFound,
		},
		{
			"Fail with error from GetById",
			&deletePostUsecaseInput{
				"test_post_id",
				1,
			},
			&responseFromPostRepositoryGetById{
				nil,
				fmt.Errorf("error from GetById"),
			},
			nil,
			fmt.Errorf("error from GetById"),
		},
		{
			"Fail with unauthorized user",
			&deletePostUsecaseInput{
				"test_post_id",
				1,
			},
			&responseFromPostRepositoryGetById{
				&entities.Post{
					UserId: 2,
				},
				nil,
			},
			nil,
			exception.ErrUnauthorizedToDeletePost,
		},
		{
			"Fail with error from Delete",
			&deletePostUsecaseInput{
				"test_post_id",
				1,
			},
			&responseFromPostRepositoryGetById{
				&entities.Post{
					UserId: 1,
				},
				nil,
			},
			fmt.Errorf("error from Delete"),
			fmt.Errorf("error from Delete"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostRepository := mock_interfaces.NewMockPostRepository(mockCtrl)
			mockPostRepository.EXPECT().GetById(tc.input.postId).Return(tc.responseGetByID.post, tc.responseGetByID.err)

			// Delete is called only if the user is authorized to delete the post
			mockPostRepository.EXPECT().Delete(tc.input.postId).Return(tc.responseFromDelete).AnyTimes()

			usecase := usecases.NewDeletePostUsecase(mockPostRepository)
			err := usecase.Execute(tc.input.postId, tc.input.userId)
			assert.Equal(t, tc.expected, err)
		})
	}
}
