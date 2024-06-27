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

func (u *GetAllPostsUsecase) Execute(limit int64, offset int64) ([]*entities.Post, error) {
	// Validate limit
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	// Validate offset
	if offset < 0 {
		offset = 0
	}

	result, err := u.repository.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return result, nil
}
