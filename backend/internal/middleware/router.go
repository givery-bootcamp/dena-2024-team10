package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", controllers.HelloWorld)
	app.POST("/signin", controllers.SignIn)
	app.POST("/signup", controllers.SignUp)
	app.POST("/signout", controllers.SignOut)
	app.GET("/posts", controllers.GetAllPosts)
}
