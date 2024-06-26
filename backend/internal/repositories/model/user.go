package model

import "myapp/internal/entities"

type User struct {
	Id       int64
	Name     string
	Password string
}

func ConvertUserModelToEntity(v *User) *entities.User {
	return &entities.User{
		Id:       v.Id,
		Username: v.Name,
		Password: v.Password,
	}
}
