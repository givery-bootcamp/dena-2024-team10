package usecases

import (
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type DeletePostUsecase struct {
	postRepository interfaces.PostRepository
}

func NewDeletePostUsecase(pr interfaces.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{
		postRepository: pr,
	}
}

func (u *DeletePostUsecase) Execute(postId string, userId int64) error {
	// Get the post by ID
	post, err := u.postRepository.GetById(postId)
	if err != nil {
		return err
	} else if post == nil {
		return exception.ErrNotFound
	}

	// Check if the user is authorized to delete the post
	if post.UserId != userId {
		return exception.ErrUnauthorizedToDeletePost
	}

	err = u.postRepository.Delete(postId)
	if err != nil {
		return err
	}

	return nil
}
