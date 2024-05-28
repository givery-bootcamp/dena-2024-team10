package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type GetAllPostsUsecase struct {
	repository interfaces.PostRepository
}

func NewGetAllPostsUsecase(r interfaces.PostRepository) *GetAllPostsUsecase {
	return &GetAllPostsUsecase{
		repository: r,
	}
}

func (u *GetAllPostsUsecase) Execute() ([]*entities.Post, error) {
	result, err := u.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}
