package usecases

import (
	"errors"
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type CreatePostUsecaseInput struct {
	request  schema.CreatePostRequest
	username string
}

type CreatePostRepositoryOutput struct {
	user *entities.User
	post []*entities.Post
	err  error
}

func TestCreatePost(t *testing.T) {
	testcases := []struct {
		testName string
		input    *CreatePostUsecaseInput
		output   *CreatePostRepositoryOutput
		wantsErr bool
	}{
		{
			"Success create post",
			&CreatePostUsecaseInput{
				request: schema.CreatePostRequest{
					Title: "test_title",
					Body:  "test_body",
				},
				username: "test_user",
			},
			&CreatePostRepositoryOutput{
				user: &entities.User{Id: 1, Username: "test_user"},
				post: []*entities.Post{
					{Id: 1, Title: "test_title1", Body: "test_body1", UserId: 1},
				},
				err: nil,
			},
			false,
		},
		{
			"Failed to get user by username",
			&CreatePostUsecaseInput{
				request: schema.CreatePostRequest{
					Title: "test_title",
					Body:  "test_body",
				},
				username: "invalid_user",
			},
			&CreatePostRepositoryOutput{
				user: nil,
				post: nil,
				err:  errors.New("user not found"),
			},
			true,
		},
		{
			"Failed to create post",
			&CreatePostUsecaseInput{
				request: schema.CreatePostRequest{
					Title: "test_title",
					Body:  "test_body",
				},
				username: "test_user",
			},
			&CreatePostRepositoryOutput{
				user: &entities.User{Id: 1, Username: "test_user"},
				post: nil,
				err:  errors.New("failed to create post"),
			},
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostRepository := mock_interfaces.NewMockPostRepository(mockCtrl)
			mockUserRepository := mock_interfaces.NewMockUserRepository(mockCtrl)

			mockUserRepository.EXPECT().GetByUsername(tc.input.username).Return(
				tc.output.user,
				tc.output.err,
			)

			if tc.output.user != nil {
				mockPostRepository.EXPECT().CreatePost(tc.input.request.Title, tc.input.request.Body, tc.output.user.Id).Return(
					tc.output.post,
					tc.output.err,
				)
			}

			usecase := usecases.NewCreatePostUsecase(mockPostRepository, mockUserRepository)
			post, err := usecase.Execute(tc.input.request, tc.input.username)

			if !tc.wantsErr {
				if post == nil {
					t.Errorf("Post is nil (success case)")
				}
				if err != nil {
					t.Errorf("Error is not nil (success case): %v", err)
				}
			} else {
				if post != nil {
					t.Errorf("Post is not nil (error case)")
				}
				if err == nil {
					t.Errorf("Error is nil (error case)")
				}
			}
		})
	}
}
