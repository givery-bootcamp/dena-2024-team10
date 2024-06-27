package usecases

import (
	"fmt"
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

type updatePostUsecaseInput struct {
	request schema.PostRequest
	userId  int64
	postId  int64
}

type responseFromUpdatePostRepositoryGetById struct {
	post *entities.Post
	err  error
}

type expectedPostUsecaseOutput struct {
	post *entities.Post
	err  error
}

func TestUpdatePost(t *testing.T) {
	testcases := []struct {
		testName           string
		input              *updatePostUsecaseInput
		responseGetByID    *responseFromUpdatePostRepositoryGetById
		responseFromUpdate *expectedPostUsecaseOutput
		expected           error
	}{
		{
			"Success",
			&updatePostUsecaseInput{
				request: schema.PostRequest{
					Title: "更新titleやで",
					Body:  "更新bodyやで",
				},
				userId: 1,
				postId: 1,
			},
			&responseFromUpdatePostRepositoryGetById{
				&entities.Post{
					Id:     1,
					UserId: 1,
					Title:  "title",
					Body:   "body",
				},
				nil,
			},
			&expectedPostUsecaseOutput{
				&entities.Post{
					Id:     1,
					UserId: 1,
					Title:  "更新titleやで",
					Body:   "更新bodyやで",
				},
				nil,
			},
			nil,
		},
		{
			"Fail with post not found",
			&updatePostUsecaseInput{
				request: schema.PostRequest{
					Title: "更新titleやで",
					Body:  "更新bodyやで",
				},
				userId: 1,
				postId: 0,
			},
			&responseFromUpdatePostRepositoryGetById{
				nil,
				exception.ErrPostNotFound,
			},
			&expectedPostUsecaseOutput{
				nil,
				exception.ErrPostNotFound,
			},
			exception.ErrPostNotFound,
		},
		{
			"Fail with error from GetById",
			&updatePostUsecaseInput{
				request: schema.PostRequest{
					Title: "更新titleやで",
					Body:  "更新bodyやで",
				},
				userId: 1,
				postId: 1,
			},
			&responseFromUpdatePostRepositoryGetById{
				nil,
				fmt.Errorf("error from GetById"),
			},
			nil,
			fmt.Errorf("error from GetById"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postRepository := mock_interfaces.NewMockPostRepository(ctrl)
			postRepository.
				EXPECT().
				GetById(tc.input.postId).
				Return(tc.responseGetByID.post, tc.responseGetByID.err)

			postRepository.
				EXPECT().
				UpdatePost(tc.input.request.Title, tc.input.request.Body, tc.input.postId).
				Return(tc.responseFromUpdate.post, tc.responseFromUpdate.err)

			usecase := usecases.NewUpdatePostUsecase(postRepository)
			post, err := usecase.Execute(tc.input.request, tc.input.userId, tc.input.postId)
			fmt.Println(err)

			assert.Equal(t, tc.expected, err)

			if post != nil {
				assert.Equal(t, tc.responseFromUpdate.post.Id, post.Id)
				assert.Equal(t, tc.responseFromUpdate.post.UserId, post.UserId)
				assert.Equal(t, tc.responseFromUpdate.post.Title, post.Title)
				assert.Equal(t, tc.responseFromUpdate.post.Body, post.Body)
			}
		})
	}
}
