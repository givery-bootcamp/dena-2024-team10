package controllers

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	postRepository := repositories.NewPostRepository(DB(ctx))
	userRepository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(postRepository, userRepository)

	username, exists := ctx.Get("username")
	if !exists {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	request := schema.CreatePostRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	result, err := usecase.Execute(request, username.(string))

	if err != nil {
		ctx.Error(err)
	} else if result == nil {
		ctx.Error(exception.ErrNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func GetAllPosts(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewGetAllPostsUsecase(repository)
	result, err := usecase.Execute()
	if err != nil {
		ctx.Error(err)
	} else if result == nil {
		ctx.Error(exception.ErrNotFound)
	} else {
		ctx.JSON(200, result)
	}
}
