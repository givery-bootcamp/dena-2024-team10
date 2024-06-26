package model

import (
	"myapp/internal/entities"
	"time"
)

type Post struct {
	Id        int64     `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	UserId    int64     `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func ConvertPostModelToEntity(v *Post) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		Title:     v.Title,
		Body:      v.Body,
		UserId:    v.UserId,
		Username:  v.User.Name,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
