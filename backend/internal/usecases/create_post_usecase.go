package usecases

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CreatePostUsecase struct {
	postRepository interfaces.PostRepository
	userRepository interfaces.UserRepository
}

func NewCreatePostUsecase(
	postRepository interfaces.PostRepository,
	userRepository interfaces.UserRepository,
) *CreatePostUsecase {
	return &CreatePostUsecase{
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

func (u *CreatePostUsecase) Execute(request schema.CreatePostRequest, username string) ([]*entities.Post, error) {
	user, err := u.userRepository.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	post, err := u.postRepository.CreatePost(request.Title, request.Body, user.Id)

	if err != nil {
		return nil, err
	}

	return post, nil
}
