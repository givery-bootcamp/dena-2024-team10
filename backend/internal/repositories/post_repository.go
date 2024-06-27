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
	var posts []*model.Post
	if err := r.Conn.Preload("User").Limit(int(limit)).Offset(int(offset)).Order("id").Find(&posts).Error; err != nil {
		return nil, err
	}

	var result []*entities.Post
	for _, post := range posts {
		result = append(result, model.ConvertPostModelToEntity(post))
	}

	return result, nil
}

func (r *PostRepository) CreatePost(title string, body string, userId int64) (*entities.Post, error) {
	post := model.Post{
		Title:  title,
		Body:   body,
		UserId: userId,
	}

	if err := r.Conn.Create(&post).Error; err != nil {
		return nil, err
	}

	return model.ConvertPostModelToEntity(&post), nil
}

func (r *PostRepository) GetById(postId int64) (*entities.Post, error) {
	var post model.Post
	if err := r.Conn.Preload("User").First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return model.ConvertPostModelToEntity(&post), nil
}

func (r *PostRepository) Delete(postId int64) error {
	return r.Conn.Delete(&model.Post{}, postId).Error
}

func (r *PostRepository) UpdatePost(title string, body string, postId int64) (*entities.Post, error) {
	post := model.Post{
		Title: title,
		Body:  body,
	}

	if err := r.Conn.Where("id = ?", postId).Updates(&post).Error; err != nil {
		return nil, err
	}

	return model.ConvertPostModelToEntity(&post), nil
}
