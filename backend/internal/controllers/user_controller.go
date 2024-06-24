package controllers

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignIn
// Check password and username contained in the request body
// If the password and username are correct, set JWT token in the Cookie
func GetSignedInUser(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecases := usecases.NewGetUserByUsernameUsecase(repository)

	username, exists := ctx.Get("username")
	if !exists {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	user, err := usecases.Execute(username.(string))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})
}
