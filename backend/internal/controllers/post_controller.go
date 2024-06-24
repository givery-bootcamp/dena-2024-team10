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

func DeletePost(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewDeletePostUsecase(repository)

	postId := ctx.Param("postId")

	// Get user ID from ctx
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	userIdInt64, ok := userId.(int64)
	if !ok {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	err := usecase.Execute(postId, userIdInt64)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(204, nil)
	}
}
