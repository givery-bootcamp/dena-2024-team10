package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type CommentRepository struct {
	Conn *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) *CommentRepository {
	return &CommentRepository{
		Conn: conn,
	}
}

func (r *CommentRepository) Create(postId int64, body string, userId int64) (*entities.Comment, error) {
	panic("not implemented")
}
