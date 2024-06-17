package usecases

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/interfaces"

	"github.com/go-sql-driver/mysql"
)

type SignUpUsecase struct {
	repository interfaces.UserRepository
}

func NewSignUpUsecase(r interfaces.UserRepository) *SignUpUsecase {
	return &SignUpUsecase{
		repository: r,
	}
}

// Execute is the method to execute SignUpUsecase
// Check if username is unique
// If the username is valid, create user and return user entity
func (u *SignUpUsecase) Execute(username, password string) (*entities.User, error) {
	// Get user by username
	user, err := u.repository.CreateUser(username, password)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, errors.New("this username is already exists")
		}
		return nil, errors.New("failed to create user: " + err.Error())
	}

	return user, nil
}
