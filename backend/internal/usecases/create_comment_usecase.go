package usecases

import (
	"myapp/internal/entities"
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
	panic("not implemented")
}
