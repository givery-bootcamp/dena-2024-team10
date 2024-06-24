package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: delete after implementing error handling to post_controller.go
type ErrorResponse struct {
	Message string "json:`message`"
}

// TODO: delete after implementing dependency injection
func DB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}

// TODO: delete after implementing error handling to post_controller.go
func handleError(ctx *gin.Context, status int, err error) {
	res := ErrorResponse{
		Message: err.Error(),
	}
	ctx.JSON(status, &res)
}
