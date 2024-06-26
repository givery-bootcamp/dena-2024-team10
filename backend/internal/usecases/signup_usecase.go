package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
	"myapp/internal/utils"
)

type SignUpUsecase struct {
	repository interfaces.UserRepository
}

func NewSignUpUsecase(r interfaces.UserRepository) *SignUpUsecase {
	return &SignUpUsecase{
		repository: r,
	}
}

// Execute is the method to execute SignUpUsecase
// Check if username is unique
// If the username is valid, create user and return user entity
func (u *SignUpUsecase) Execute(username, password string) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.CreateUser(username, hashedPassword)
	if err != nil {
		if err.Error() == "user already exists" {
			return nil, exception.ErrDuplicateUser
		}
		return nil, err
	}
	return user, nil
}
