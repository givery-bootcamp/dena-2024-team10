package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	mock_interfaces "myapp/test/mock"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Input struct {
	username string
	password string
}

type RepositoryOutput struct {
	user *entities.User
	err  error
}

func TestSignIn(t *testing.T) {
	testcases := []struct {
		testName string
		input    *Input
		output   *RepositoryOutput
		wantsErr bool
	}{
		{
			"Success sign in",
			&Input{
				"test_username",
				"test_password",
			},
			&RepositoryOutput{
				&entities.User{Username: "test_username", Password: "test_password"},
				nil,
			},
			false,
		},
		{
			"Failed to get user",
			&Input{
				"test_username",
				"test_password",
			},
			&RepositoryOutput{
				nil,
				errors.New("user not found"),
			},
			true,
		},
		{
			"Incorrect password",
			&Input{
				"test_username",
				"incorrect_password",
			},
			&RepositoryOutput{
				&entities.User{Username: "test_username", Password: "test_password"},
				nil,
			},
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUserRepository := mock_interfaces.NewMockUserRepository(mockCtrl)
			mockUserRepository.EXPECT().GetByUsername(tc.input.username).Return(
				tc.output.user,
				tc.output.err,
			)

			usecase := usecases.NewSignInUsecase(mockUserRepository)
			user, token, err := usecase.Execute(tc.input.username, tc.input.password)

			if !tc.wantsErr {
				// check user
				if user == nil {
					t.Errorf("User is nil (success case)")
				}
				if user != nil && user != tc.output.user {
					t.Errorf("User is not equal to expected user (success case)")
				}

				// check token
				if token == "" {
					t.Errorf("Token is empty (success case)")
				}
				// token の検証を行いたいが、CreateToken がモック化できていないため、検証ができない

				// check error
				if err != nil {
					t.Errorf("Error is not nil (success case)")
				}
			} else {
				// check user
				if user != nil {
					t.Errorf("User is not nil (error case)")
				}

				// check token
				if token != "" {
					t.Errorf("Token is not empty (error case)")
				}

				// check error
				if err == nil {
					t.Errorf("Error is nil (error case)")
				}
			}
		})
	}
}
