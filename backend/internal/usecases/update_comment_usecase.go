package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/interfaces"
)

type UpdateCommentUsecase struct {
	CommentRepository interfaces.CommentRepository
}

func NewUpdateCommentUsecase(cr interfaces.CommentRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		CommentRepository: cr,
	}
}

func (u *UpdateCommentUsecase) Execute(userId int64, postId int64, commentId int64, body string) (*entities.Comment, error) {
	// Check if the comment exists
	comment, err := u.CommentRepository.GetById(commentId)
	if err != nil {
		return nil, err
	} else if comment == nil {
		return nil, exception.ErrCommentNotFound
	}

	// Check if the user is the owner of the comment
	if comment.UserId != userId {
		return nil, exception.ErrUnauthorizedToUpdateComment
	}

	// Check if the post_id is the same as the comment's post_id
	if comment.PostId != postId {
		return nil, exception.ErrInvalidPostId
	}

	// Update the comment
	newComment := &entities.Comment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		Body:      body, // Update the body
		CreatedAt: comment.CreatedAt,
	}

	updatedComment, err := u.CommentRepository.Update(newComment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}
