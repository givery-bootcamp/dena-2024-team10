package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignIn
// Check password and username contained in the request body
// If the password and username are correct, set JWT token in the Cookie
func SignIn(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewSignInUsecase(repository)

	body := SignInRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"failed to sign in": err.Error()})
		return
	}

	_, _, err := usecase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"failed to sign in": err.Error()})
		return
	}
}
