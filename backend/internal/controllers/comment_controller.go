package controllers

import (
	"myapp/internal/controllers/schema"
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

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

	req := &schema.CommentRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(exception.ErrInvalidRequest)
		return
	}

	comment, err := usecase.Execute(req.PostId, req.Body, userIdInt64)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(200, comment)
	}
}
