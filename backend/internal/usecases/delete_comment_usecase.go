package usecases

import (
	"myapp/internal/interfaces"
)

type DeleteCommentUsecase struct {
	commentRepository interfaces.CommentRepository
}

func NewDeleteCommentUsecase(cr interfaces.CommentRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{
		commentRepository: cr,
	}
}

func (u *DeleteCommentUsecase) Execute(postId int64, commentId int64, userId int64) error {
	panic("not implemented") // TODO: Implement
	// // Get the comment by ID
	// comment, err := u.commentRepository.GetById(commentId)
	// if err != nil {
	// 	return err
	// } else if comment == nil {
	// 	return exception.ErrCommentNotFound
	// }

	// // Check if the user is authorized to delete the comment
	// if comment.UserId != userId {
	// 	return exception.ErrUnauthorizedToDeleteComment
	// }

	// // Check if the comment belongs to the post
	// if comment.PostId != postId {
	// 	return exception.ErrInvalidPostId
	// }

	// return u.commentRepository.Delete(commentId)
}
