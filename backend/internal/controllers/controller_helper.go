package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: delete after implementing dependency injection
func DB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}
