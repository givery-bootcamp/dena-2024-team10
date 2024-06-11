package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := usecase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("JWT", token, 0, "/", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
