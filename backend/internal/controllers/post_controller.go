package controllers

import (
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

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
