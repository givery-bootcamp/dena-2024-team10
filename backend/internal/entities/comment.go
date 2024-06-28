package entities

import "time"

type Comment struct {
	Id        int64     `json:"id"`
	PostId    int64     `json:"post_id"`
	UserId    int64     `json:"user_id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
