//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../test/mock/mock_$GOPACKAGE/$GOFILE
package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	GetByUsername(username string) (*entities.User, error)
}
