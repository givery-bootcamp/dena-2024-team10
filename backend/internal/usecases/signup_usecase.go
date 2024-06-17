package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
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
	user, err := u.repository.CreateUser(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
