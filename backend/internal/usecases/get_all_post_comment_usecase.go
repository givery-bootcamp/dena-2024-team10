package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type GetAllPostCommentUsecase struct {
	CommentRepository interfaces.CommentRepository
}

func NewGetAllPostCommentUsecase(cr interfaces.CommentRepository) *GetAllPostCommentUsecase {
	return &GetAllPostCommentUsecase{
		CommentRepository: cr,
	}
}

func (u *GetAllPostCommentUsecase) Execute(postId, limit, offset int64) ([]*entities.Comment, error) {
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

	comments, err := u.CommentRepository.GetAllByPostId(postId, limit, offset)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
