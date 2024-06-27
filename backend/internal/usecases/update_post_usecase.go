package usecases

import (
	"errors"
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type UpdatePostUsecase struct {
	postRepository interfaces.PostRepository
}

func NewUpdatePostUsecase(
	postRepository interfaces.PostRepository,
) *UpdatePostUsecase {
	return &UpdatePostUsecase{
		postRepository: postRepository,
	}
}

func (u *UpdatePostUsecase) Execute(request schema.PostRequest, userId int64, postId int64) (*entities.Post, error) {
	post, err := u.postRepository.GetById(postId)
	if post == nil || err == errors.New("post not found") {
		return nil, exception.ErrPostNotFound
	} else if err != nil {
		return nil, err
	}

	// Check if the user is authorized to delete the post
	if post.UserId != userId {
		return nil, exception.ErrUnauthorizedToUpdatePost
	}

	return u.postRepository.UpdatePost(request.Title, request.Body, userId)
}
