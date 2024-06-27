package usecases

import (
	"myapp/internal/entities"
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

	panic("not implemented")

	// comment, err := u.CommentRepository.Get(commentId)
	// if err != nil {
	// 	if err.Error() == "comment not found" {
	// 		return nil, exception.ErrCommentNotFound
	// 	}
	// 	return nil, err
	// }
	// return comment, nil
}
