package interfaces

import (
	"myapp/internal/entities"
)

type UserRepository interface {
	GetByUsername(username string) (*entities.User, error)
}
