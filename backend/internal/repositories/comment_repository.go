package repositories

import (
	"myapp/internal/entities"
	"myapp/internal/repositories/model"

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
	comment := &model.Comment{
		PostId: postId,
		UserId: userId,
		Body:   body,
	}

	if err := r.Conn.Create(comment).Error; err != nil {
		return nil, err
	}

	return model.ConvertCommentModelToEntity(comment), nil
}
