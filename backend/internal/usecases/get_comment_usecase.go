package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type GetCommentUsecase struct {
	CommentRepository interfaces.CommentRepository
}

func NewGetCommentUsecase(cr interfaces.CommentRepository) *GetCommentUsecase {
	return &GetCommentUsecase{
		CommentRepository: cr,
	}
}

func (u *GetCommentUsecase) Execute(commentId int64) (*entities.Comment, error) {
	comment, err := u.CommentRepository.GetById(commentId)
	if err != nil {
		return nil, err
	} else if comment == nil {
		return nil, exception.ErrNotFound
	}
	return comment, nil
}
