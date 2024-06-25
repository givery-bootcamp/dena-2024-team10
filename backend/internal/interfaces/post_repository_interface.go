//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../test/mock/mock_$GOPACKAGE/$GOFILE
package interfaces

import (
	"myapp/internal/entities"
)

type PostRepository interface {
	GetAll() ([]*entities.Post, error)
	CreatePost(title string, body string, userId int64) ([]*entities.Post, error)
	GetById(postId int64) (*entities.Post, error)
	Delete(postId int64) error
}
