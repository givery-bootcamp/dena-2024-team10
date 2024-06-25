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
	Username  string
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

func (r *PostRepository) GetAll(limit, offset int64) ([]*entities.Post, error) {
	var posts []*Post
	if err := r.Conn.Limit(int(limit)).Offset(int(offset)).Find(&posts).Error; err != nil {
		return nil, err
	}

	var result []*entities.Post
	for _, post := range posts {
		result = append(result, convertPostRepositoryModelToEntity(post))
	}

	return result, nil
}

func (r *PostRepository) GetById(postId int64) (*entities.Post, error) {
	var post Post
	if err := r.Conn.First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return convertPostRepositoryModelToEntity(&post), nil
}

func (r *PostRepository) Delete(postId int64) error {
	return r.Conn.Delete(&Post{}, postId).Error
}

func convertPostRepositoryModelToEntity(v *Post) *entities.Post {
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
