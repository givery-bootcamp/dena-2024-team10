package usecases

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CreatePostUsecase struct {
	postRepository interfaces.PostRepository
}

func NewCreatePostUsecase(
	postRepository interfaces.PostRepository,
) *CreatePostUsecase {
	return &CreatePostUsecase{
		postRepository: postRepository,
	}
}

func (u *CreatePostUsecase) Execute(request schema.CreatePostRequest, userId int64) (*entities.Post, error) {
	post, err := u.postRepository.CreatePost(request.Title, request.Body, userId)

	if err != nil {
		return nil, err
	}

	return post, nil
}
