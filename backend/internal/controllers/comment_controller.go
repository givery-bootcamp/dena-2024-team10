package controllers

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewCreateCommentUsecase(repository)

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

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrPostNotFound)
		return
	}

	req := &schema.CommentRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	comment, err := usecase.Execute(postIdInt64, req.Body, userIdInt64)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusOK, comment)
	}
}

func GetComment(ctx *gin.Context) {
	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewGetCommentUsecase(repository)

	commentId := ctx.Param("commentId")
	commentIdInt64, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrNotFound)
		return
	}

	comment, err := usecase.Execute(commentIdInt64)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func UpdateComment(ctx *gin.Context) {
	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewUpdateCommentUsecase(repository)

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

	commentId := ctx.Param("commentId")
	commentIdInt64, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrNotFound)
		return
	}

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrNotFound)
		return
	}

	req := &schema.CommentRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	comment, err := usecase.Execute(userIdInt64, postIdInt64, commentIdInt64, req.Body)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func DeleteComment(ctx *gin.Context) {
	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewDeleteCommentUsecase(repository)

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

	commentId := ctx.Param("commentId")
	commentIdInt64, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrNotFound)
		return
	}

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrInvalidPostId)
		return
	}

	err = usecase.Execute(postIdInt64, commentIdInt64, userIdInt64)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func GetComments(ctx *gin.Context) {
	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewGetAllPostCommentUsecase(repository)

	postId := ctx.Param("postId")
	postIdInt64, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		ctx.Error(exception.ErrNotFound)
		return
	}

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

	comments, err := usecase.Execute(postIdInt64, limit, offset)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
