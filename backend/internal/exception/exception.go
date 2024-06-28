package exception

import "net/http"

type Exception struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e Exception) Error() string {
	return e.Message
}

func new(status int, message string) *Exception {
	return &Exception{Status: status, Message: message}
}

var (
	// 4xx: Client Error
	ErrDuplicateUser               = new(http.StatusBadRequest, "Duplicate User")
	ErrInvalidQuery                = new(http.StatusBadRequest, "Invalid Query")
	ErrInvalidRequest              = new(http.StatusBadRequest, "Invalid Request")
	ErrInvalidPostId               = new(http.StatusBadRequest, "Invalid Post Id")
	ErrPostNotFound                = new(http.StatusBadRequest, "Post Not Found")
	ErrCommentNotFound             = new(http.StatusBadRequest, "Comment Not Found")
	ErrSigninFailed                = new(http.StatusBadRequest, "Signin Failed")
	ErrUnauthorized                = new(http.StatusUnauthorized, "Unauthorized")
	ErrUnauthorizedToDeletePost    = new(http.StatusUnauthorized, "Unauthorized to delete this post")
	ErrUnauthorizedToUpdatePost    = new(http.StatusUnauthorized, "Unauthorized to update this post")
	ErrUnauthorizedToUpdateComment = new(http.StatusUnauthorized, "Unauthorized to update this comment")
	ErrNotFound                    = new(http.StatusNotFound, "Not Found")
	// 5xx: Server Error
	ErrInternalServerError = new(500, "Internal Server Error")
)
