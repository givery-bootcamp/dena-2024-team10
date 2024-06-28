package schema

type PostRequest struct {
	Title string `json:"title" max:"100" min:"1" required:"true"`
	Body  string `json:"body" min:"1" required:"true"`
}
