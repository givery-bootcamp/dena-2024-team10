package entities

import "time"

type Comment struct {
	Id        int64
	PostId    int64
	UserId    int64
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
