package middleware

import (
	"myapp/internal/exception"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// Check if there are any errors
		if len(ctx.Errors) == 0 {
			return
		}

		// Handle error
		var e *exception.Exception
		for _, err := range ctx.Errors {
			if castedErr, ok := err.Err.(*exception.Exception); ok {
				e = castedErr
				break
			}
		}
		if e == nil {
			e = exception.ErrInternalServerError
		}

		ctx.JSON(e.Status, e)
		ctx.Abort()
	}
}
