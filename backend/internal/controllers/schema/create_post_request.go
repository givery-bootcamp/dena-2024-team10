package schema

type CreatePostRequest struct {
	Title string `json:"title" max:"100"`
	Body  string `json:"body"`
}
