package controllers

import (
	"myapp/internal/controllers/schema"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
		return
	}
	user, err := usecases.Execute(username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})
}
