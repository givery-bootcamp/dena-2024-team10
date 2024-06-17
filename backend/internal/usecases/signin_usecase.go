package usecases

import (
	"errors"
	"myapp/internal/entities"
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
		return nil, "", errors.New("failed to get user: " + err.Error())
	}

	// Check password
	if user.Password != password {
		return nil, "", errors.New("password is incorrect")
	}

	// Create JWT token
	timeToExpire := time.Now().Add(time.Hour * 24).Unix()
	token, err := utils.CreateToken(username, timeToExpire)
	if err != nil {
		return nil, "", errors.New("failed to create token: " + err.Error())
	}

	return user, token, nil
}
