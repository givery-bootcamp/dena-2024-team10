package middleware

import (
	"myapp/internal/config"
	"myapp/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(ctx *gin.Context) {
	jwtToken, err := ctx.Cookie(config.CookieNameForJWT)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	parsedToken, err := utils.ParseToken(jwtToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	username, err := utils.GetUsernameFromParsedToken(parsedToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	ctx.Set("username", username)

	ctx.Next()
}