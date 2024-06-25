package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type getPostRepositoryInput struct {
	postId int64
}

type getPostRepositoryOutput struct {
	post *entities.Post
	err  error
}

func TestGetPost(t *testing.T) {

	// create test cases
	testcases := []struct {
		testName         string
		repositoryInput  *getPostRepositoryInput
		repositoryOutput *getPostRepositoryOutput
		wantsFail        bool
	}{
		{
			"Success get post",
			&getPostRepositoryInput{
				postId: 1,
			},
			&getPostRepositoryOutput{
				&entities.Post{
					UserId: 1,
					Title:  "モックタイトル",
					Body:   "モック本文",
				},
				nil,
			},
			false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockPostRepository := mock_interfaces.NewMockPostRepository(mockCtrl)
			mockPostRepository.EXPECT().GetById(tc.repositoryInput.postId).Return(
				tc.repositoryOutput.post,
				tc.repositoryOutput.err,
			)

			usecase := usecases.NewGetPostUsecase(mockPostRepository)
			post, err := usecase.Execute(tc.repositoryInput.postId)

			if !tc.wantsFail {
				// check post
				if post == nil {
					t.Errorf("User is nil (success case)")
				}

				if post != nil && post != tc.repositoryOutput.post {
					t.Errorf("Post is not equal to expected user (success case)")
				}

				// check error
				if err != nil {
					t.Errorf("Error is not nil (success case)")
				}
			} else {
				// check user
				if post != nil {
					t.Errorf("Post is not nil (error case)")
				}

				// check error
				if err == nil {
					t.Errorf("Error is nil (error case)")
				}
			}
		})
	}
}
