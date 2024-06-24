package repositories

import (
	"fmt"
	"myapp/internal/entities"

	"github.com/go-sql-driver/mysql"
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
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return convertUserModelToEntity(&user), nil
}

func (r *UserRepository) CreateUser(username, password string) (*entities.User, error) {
	user := User{Name: username, Password: password}

	result := r.Conn.Create(&user)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return nil, fmt.Errorf("user already exists")
			default:
				return nil, fmt.Errorf("failed to create user due to mysql error: %w", result.Error)
			}
		} else {
			return nil, fmt.Errorf("failed to create user: %w", result.Error)
		}
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
