package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type Post struct {
	Id        int64  `gorm:"column:id"`
	Title     string `gorm:"column:title"`
	Body      string `gorm:"column:body"`
	UserId    int64  `gorm:"column:user_id"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

type PostWithUsername struct {
	Id        int64  `gorm:"column:id"`
	Title     string `gorm:"column:title"`
	Body      string `gorm:"column:body"`
	UserId    int64  `gorm:"column:user_id"`
	Username  string `gorm:"column:username"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) GetAll(limit, offset int64) ([]*entities.Post, error) {
	var posts []*PostWithUsername
	if err := r.Conn.Table("posts").Select("posts.*, users.name as username").Joins("JOIN users ON posts.user_id = users.id").
		Order("posts.id").Limit(int(limit)).Offset(int(offset)).Scan(&posts).Error; err != nil {
		return nil, err
	}

	var result []*entities.Post
	for _, post := range posts {
		result = append(result, convertPostWithUsernameToEntity(post))
	}

	return result, nil
}

func (r *PostRepository) GetById(postId int64) (*entities.Post, error) {
	var post PostWithUsername
	if err := r.Conn.Table("posts").Select("posts.*, users.name as username").Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.id = ?", postId).Scan(&post).Error; err != nil {
		return nil, err
	}

	if len(post.Username) == 0 {
		return nil, nil
	}
	return convertPostWithUsernameToEntity(&post), nil
}

func (r *PostRepository) Delete(postId int64) error {
	return r.Conn.Delete(&Post{}, postId).Error
}

func convertPostWithUsernameToEntity(v *PostWithUsername) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		Title:     v.Title,
		Body:      v.Body,
		UserId:    v.UserId,
		Username:  v.Username,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
