package schema

type CommentRequest struct {
	Body string `json:"body" binding:"required"`
}
