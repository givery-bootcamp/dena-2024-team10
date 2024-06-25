package entities

type Post struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    int64  `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
