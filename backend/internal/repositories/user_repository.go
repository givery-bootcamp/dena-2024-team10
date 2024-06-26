package repositories

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/repositories/model"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetByUsername(username string) (*entities.User, error) {
	var user model.User
	if err := r.Conn.Where("name = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return model.ConvertUserModelToEntity(&user), nil
}

func (r *UserRepository) CreateUser(username, password string) (*entities.User, error) {
	user := model.User{Name: username, Password: password}

	if err := r.Conn.Create(&user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return nil, fmt.Errorf("user already exists")
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return model.ConvertUserModelToEntity(&user), nil
}
