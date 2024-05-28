package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type Post struct {
	Id        int64
	Title     string
	Body      string
	UserId    int64
	UserName  string
	CreatedAt string
	UpdatedAt string
}

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) GetAll() ([]*entities.Post, error) {
	var posts []*Post
	if err := r.Conn.Find(&posts).Error; err != nil {
		return nil, err
	}

	var result []*entities.Post
	for _, post := range posts {
		result = append(result, convertPostRepositoryModelToEntity(post))
	}

	return result, nil
}

func convertPostRepositoryModelToEntity(v *Post) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		Title:     v.Title,
		Body:      v.Body,
		UserId:    v.UserId,
		UserName:  v.UserName,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
