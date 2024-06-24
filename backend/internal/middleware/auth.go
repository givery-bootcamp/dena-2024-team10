package middleware

import (
	"myapp/internal/config"
	"myapp/internal/exception"
	"myapp/internal/utils"

	"github.com/gin-gonic/gin"
)

func Authenticated(ctx *gin.Context) {
	jwtToken, err := ctx.Cookie(config.CookieNameForJWT)
	if err != nil {
		ctx.Error(exception.ErrUnauthorized)
		ctx.Abort()
		return
	}

	parsedToken, err := utils.ParseToken(jwtToken)
	if err != nil {
		ctx.Error(exception.ErrUnauthorized)
		ctx.Abort()
		return
	}

	username, err := utils.GetUsernameFromParsedToken(parsedToken)
	if err != nil {
		ctx.Error(exception.ErrUnauthorized)
		ctx.Abort()
		return
	}

	ctx.Set("username", username)

	ctx.Next()
}
