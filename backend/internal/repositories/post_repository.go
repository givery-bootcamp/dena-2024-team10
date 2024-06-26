package repositories

import (
	"myapp/internal/entities"
	"myapp/internal/repositories/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) GetAll(limit, offset int64) ([]*entities.Post, error) {
	var posts []*model.PostWithUsername
	if err := r.Conn.Table("posts").Select("posts.*, users.name as username").Joins("JOIN users ON posts.user_id = users.id").
		Order("posts.id").Limit(int(limit)).Offset(int(offset)).Scan(&posts).Error; err != nil {
		return nil, err
	}

	var result []*entities.Post
	for _, post := range posts {
		result = append(result, model.ConvertPostWithUsernameToEntity(post))
	}

	return result, nil
}

func (r *PostRepository) GetById(postId int64) (*entities.Post, error) {
	var post model.PostWithUsername
	if err := r.Conn.Table("posts").Select("posts.*, users.name as username").Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.id = ?", postId).Scan(&post).Error; err != nil {
		return nil, err
	}

	if len(post.Username) == 0 {
		return nil, nil
	}
	return model.ConvertPostWithUsernameToEntity(&post), nil
}

func (r *PostRepository) Delete(postId int64) error {
	return r.Conn.Delete(&model.Post{}, postId).Error
}
