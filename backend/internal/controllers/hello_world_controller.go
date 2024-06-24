package controllers

import (
	"myapp/internal/exception"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "ja")
	if err := validateHelloWorldParameters(lang); err != nil {
		ctx.Error(err)
		return
	}
	repository := repositories.NewHelloWorldRepository(DB(ctx))
	usecase := usecases.NewHelloWorldUsecase(repository)
	result, err := usecase.Execute(lang)
	if err != nil {
		ctx.Error(err)
	} else if result == nil {
		ctx.Error(exception.ErrNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func validateHelloWorldParameters(lang string) error {
	if len(lang) != 2 {
		return exception.ErrInvalidQuery
	}
	return nil
}
