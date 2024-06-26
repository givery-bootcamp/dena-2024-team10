package schema

type CommentRequest struct {
	PostId int64  `json:"postId" binding:"required"`
	Body   string `json:"content" binding:"required"`
}
