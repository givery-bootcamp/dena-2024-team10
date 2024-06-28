package controllers

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/entities"
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	postRepository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(postRepository)

	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	request := schema.PostRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	result, err := usecase.Execute(request, userId.(int64))

	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func GetAllPosts(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewGetAllPostsUsecase(repository)

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "20"), 10, 64)
	if err != nil {
		ctx.Error(exception.ErrInvalidQuery)
		return
	}

	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		ctx.Error(exception.ErrInvalidQuery)
		return
	}

	result, err := usecase.Execute(limit, offset)
	if err != nil {
		ctx.Error(err)
	} else if result == nil {
		// result が nil の場合は空の配列を返す
		ctx.JSON(http.StatusOK, []*entities.Post{})
	} else {
		ctx.JSON(http.StatusOK, result)
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
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func UpdatePost(ctx *gin.Context) {
	postRepository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewUpdatePostUsecase(postRepository)

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		// postId を int64 に変換できない場合は 404 Not Found
		ctx.Error(exception.ErrNotFound)
		return
	}

	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.Error(exception.ErrUnauthorized)
		return
	}

	request := schema.PostRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	result, err := usecase.Execute(request, userId.(int64), postIdInt64)

	if err != nil {
		ctx.Error(err)
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
		ctx.JSON(http.StatusNoContent, nil)
	}
}
