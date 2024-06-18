package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"myapp/test/mock/mock_interfaces"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type SignUpUsecaseInput struct {
	username string
	password string
}

type SignUpRepositoryOutput struct {
	user *entities.User
	err  error
}

func TestSignUp(t *testing.T) {
	testcases := []struct {
		testName string
		input    *SignUpUsecaseInput
		output   *SignUpRepositoryOutput
		wantsErr bool
	}{
		{
			"Success sign up",
			&SignUpUsecaseInput{
				"test_username",
				"test_password",
			},
			&SignUpRepositoryOutput{
				&entities.User{Username: "test_username", Password: "test_password"},
				nil,
			},
			false,
		},
		{
			"Failed to create user",
			&SignUpUsecaseInput{
				"test_username",
				"test_password",
			},
			&SignUpRepositoryOutput{
				nil,
				errors.New("this username is already exists"),
			},
			true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUserRepository := mock_interfaces.NewMockUserRepository(mockCtrl)
			mockUserRepository.EXPECT().CreateUser(tc.input.username, tc.input.password).Return(
				tc.output.user,
				tc.output.err,
			)

			usecase := usecases.NewSignUpUsecase(mockUserRepository)
			user, err := usecase.Execute(tc.input.username, tc.input.password)

			if !tc.wantsErr {
				// check user
				if user == nil {
					t.Errorf("User is nil (success case)")
				}
				if user != nil && user != tc.output.user {
					t.Errorf("User is not equal to expected user (success case)")
				}

				// check error
				if err != nil {
					t.Errorf("Error is not nil (success case)")
				}
			} else {
				// check user
				if user != nil {
					t.Errorf("User is not nil (error case)")
				}

				// check error
				if err == nil {
					t.Errorf("Error is nil (error case)")
				}
			}
		})
	}
}
