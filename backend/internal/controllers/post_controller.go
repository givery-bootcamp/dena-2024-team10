package controllers

import (
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

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

func GetPost(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecases := usecases.NewGetPostUsecase(repository)

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		// postId を int64 に変換できない場合は 404 Not Found
		ctx.Error(exception.ErrNotFound)
		return
	}
	result, err := usecases.Execute(postIdInt64)
	if err != nil {
		ctx.Error(err)
	} else if result == nil {
		ctx.Error(exception.ErrNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func DeletePost(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewDeletePostUsecase(repository)

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		// postId を int64 に変換できない場合は 404 Not Found
		ctx.Error(exception.ErrNotFound)
		return
	}

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

	err = usecase.Execute(postIdInt64, userIdInt64)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(204, nil)
	}
}
