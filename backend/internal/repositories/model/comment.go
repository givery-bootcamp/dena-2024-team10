package model

import (
	"myapp/internal/entities"
	"time"
)

type Comment struct {
	Id        int64      `gorm:"column:id;primary_key"`
	PostId    int64      `gorm:"column:post_id"`
	UserId    int64      `gorm:"column:user_id"`
	Body      string     `gorm:"column:body"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func ConvertCommentModelToEntity(comment *Comment) *entities.Comment {
	return &entities.Comment{
		Id:        comment.Id,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
