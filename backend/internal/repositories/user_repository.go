package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type User struct {
	Id       int64
	Name     string
	Password string
}

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetByUsername(username string) (*entities.User, error) {
	var user User
	if err := r.Conn.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return convertUserModelToEntity(&user), nil
}

func convertUserModelToEntity(v *User) *entities.User {
	return &entities.User{
		Id:       v.Id,
		Username: v.Name,
		Password: v.Password,
	}
}
