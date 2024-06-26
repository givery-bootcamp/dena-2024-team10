package usecases

import (
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
func (u *SignInUsecase) Execute(username, nonHashedPassword string) (*entities.User, string, error) {
	// Get user by username
	user, err := u.repository.GetByUsername(username)
	if err != nil {
		return nil, "", err
	}

	// Check if user is nil
	if user == nil {
		return nil, "", exception.ErrSigninFailed
	}

	err = utils.CheckPasswordHash(user.Password, nonHashedPassword)
	if err != nil {
		return nil, "", exception.ErrSigninFailed
	}

	// Create JWT token
	timeToExpire := time.Now().Add(time.Hour * 24).Unix()
	token, err := utils.CreateToken(user.Id, username, timeToExpire)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
