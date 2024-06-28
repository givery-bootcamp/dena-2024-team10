package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type CreateCommentUsecase struct {
	CommentRepository interfaces.CommentRepository
}

func NewCreateCommentUsecase(cr interfaces.CommentRepository) *CreateCommentUsecase {
	return &CreateCommentUsecase{
		CommentRepository: cr,
	}
}

func (u *CreateCommentUsecase) Execute(postId int64, body string, userId int64) (*entities.Comment, error) {
	comment, err := u.CommentRepository.Create(postId, body, userId)
	if err != nil {
		if err.Error() == "post or user not found" {
			return nil, exception.ErrPostNotFound
		}
		return nil, err
	}

	return comment, nil
}
