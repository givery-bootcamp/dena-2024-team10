package usecases

import "myapp/internal/interfaces"

type UpdateCommentUsecase struct {
	CommentRepository interfaces.CommentRepository
}

func NewUpdateCommentUsecase(cr interfaces.CommentRepository) *UpdateCommentUsecase {
	return &UpdateCommentUsecase{
		CommentRepository: cr,
	}
}

func (u *UpdateCommentUsecase) Execute(commentId int64, body string) error {
	panic("not implemented") // TODO: Implement
	// _, err := u.CommentRepository.Update(commentId, body)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
