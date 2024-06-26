//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../test/mock/mock_$GOPACKAGE/$GOFILE
package interfaces

import "myapp/internal/entities"

type CommentRepository interface {
	Create(postId int64, body string, userId int64) (*entities.Comment, error)
	GetById(commentId int64) (*entities.Comment, error)
	Update(comment *entities.Comment) (*entities.Comment, error)
	Delete(commentId int64) error
	GetAllByPostId(postId, limit, offset int64) ([]*entities.Comment, error)
}
