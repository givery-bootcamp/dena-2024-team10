package usecases

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
	"myapp/internal/utils"
	"time"
)

type SignInUsecase struct {
	repository interfaces.UserRepository
}

func NewSignInUsecase(r interfaces.UserRepository) *SignInUsecase {
	return &SignInUsecase{
		repository: r,
	}
}

// Execute is the method to execute SignInUsecase
// Check password and username contained in the request body
// If the password and username are correct, return User entity and JWT token
func (u *SignInUsecase) Execute(username, password string) (*entities.User, string, error) {
	// Get user by username
	user, err := u.repository.GetByUsername(username)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, "", exception.ErrSigninFailed
		}
		return nil, "", fmt.Errorf("failed to get user by username: %w", err)
	}

	// Check password
	if user.Password != password {
		return nil, "", exception.ErrSigninFailed
	}

	// Create JWT token
	timeToExpire := time.Now().Add(time.Hour * 24).Unix()
	token, err := utils.CreateToken(username, timeToExpire)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create token: %w", err)
	}

	return user, token, nil
}
