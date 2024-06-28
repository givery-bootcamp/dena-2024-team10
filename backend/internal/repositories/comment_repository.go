package repositories

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories/model"

	"github.com/go-sql-driver/mysql"
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1452:
				return nil, errors.New("post or user not found")
			default:
				return nil, err
			}
		}
		return nil, err
	}

	return model.ConvertCommentModelToEntity(comment), nil
}

func (r *CommentRepository) GetById(commentId int64) (*entities.Comment, error) {
	comment := &model.Comment{}
	if err := r.Conn.Preload("User").First(comment, commentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return model.ConvertCommentModelToEntity(comment), nil
}

func (r *CommentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	newComment := model.Comment{
		Id:     comment.Id,
		PostId: comment.PostId,
		UserId: comment.UserId,
		Body:   comment.Body,
	}

	result := r.Conn.Model(&newComment).
		Where("id = ?", newComment.Id).
		Updates(&newComment)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("comment not found")
	} else if result.RowsAffected > 1 {
		return nil, errors.New("more than one comment updated")
	}

	return model.ConvertCommentModelToEntity(&newComment), nil
}

func (r *CommentRepository) Delete(commentId int64) error {
	return r.Conn.Delete(&model.Comment{}, commentId).Error
}

func (r *CommentRepository) GetAllByPostId(postId, limit, offset int64) ([]*entities.Comment, error) {
	var modelComments []*model.Comment
	if err := r.Conn.
		Preload("User").
		Where("post_id = ?", postId).
		Limit(int(limit)).
		Offset(int(offset)).
		Order("id").
		Find(&modelComments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	result := []*entities.Comment{}
	for _, comment := range modelComments {
		result = append(result, model.ConvertCommentModelToEntity(comment))
	}
	return result, nil
}
