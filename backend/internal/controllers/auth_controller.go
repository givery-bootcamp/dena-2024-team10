package controllers

import (
	"myapp/internal/config"
	"myapp/internal/controllers/schema"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignIn
// Check password and username contained in the request body
// If the password and username are correct, set JWT token in the Cookie
func SignIn(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewSignInUsecase(repository)

	body := schema.SignInRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := usecase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie(config.CookieNameForJWT, token, 0, "/", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})
}
