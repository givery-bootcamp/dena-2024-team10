package schema

type CommentRequest struct {
	PostId int64  `json:"post_id" binding:"required"`
	Body   string `json:"body" binding:"required"`
}
