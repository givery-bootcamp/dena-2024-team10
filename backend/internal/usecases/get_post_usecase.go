package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type GetPostUsecase struct {
	repository interfaces.PostRepository
}

func NewGetPostUsecase(r interfaces.PostRepository) *GetPostUsecase {
	return &GetPostUsecase{
		repository: r,
	}
}

func (u *GetPostUsecase) Execute(postId int64) (*entities.Post, error) {
	post, err := u.repository.GetById(postId)
	if err != nil {
		return nil, err
	} else if post == nil {
		return nil, exception.ErrNotFound
	}
	return post, nil
}
