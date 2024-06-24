package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type GetUserByUsernameUsecase struct {
	repository interfaces.UserRepository
}

func NewGetUserByUsernameUsecase(r interfaces.UserRepository) *GetUserByUsernameUsecase {
	return &GetUserByUsernameUsecase{
		repository: r,
	}
}

// Execute is the method to execute GetUserByUsernameUsecase
func (u *GetUserByUsernameUsecase) Execute(username string) (*entities.User, error) {
	return u.repository.GetByUsername(username)
}
