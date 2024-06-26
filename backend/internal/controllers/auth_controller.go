package controllers

import (
	"myapp/internal/config"
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
func SignIn(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewSignInUsecase(repository)

	body := schema.SignInRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	user, token, err := usecase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie(config.CookieNameForJWT, token, 0, "/", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})
}

// SignUp
// Check password and username contained in the request body
// If the username is not unique, reject request
// If the username is valid, create new user
func SignUp(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	signUpUsecase := usecases.NewSignUpUsecase(repository)

	body := schema.SignUpRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	user, err := signUpUsecase.Execute(body.Username, body.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	signinUsecase := usecases.NewSignInUsecase(repository)

	user, token, err := signinUsecase.Execute(user.Username, user.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie(config.CookieNameForJWT, token, 0, "/", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})

	ctx.JSON(http.StatusOK, schema.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	})
}

// SignOut
// Delete JWT token in the Cookie
func SignOut(ctx *gin.Context) {
	ctx.SetCookie(config.CookieNameForJWT, "", -1, "/", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusNoContent, nil)
}
