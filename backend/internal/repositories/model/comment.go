package model

import "time"

type Comment struct {
	Id        int64      `gorm:"column:id;primary_key"`
	PostId    int64      `gorm:"column:post_id"`
	UserId    int64      `gorm:"column:user_id"`
	Body      string     `gorm:"column:body"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
