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
	authenticated := app.Group("/")
	authenticated.Use(Authenticated)
	{
		authenticated.GET("/user", controllers.GetSignedInUser)
		authenticated.POST("/signout", controllers.SignOut)
		authenticated.GET("/posts", controllers.GetAllPosts)
		authenticated.GET("/posts/:postId", controllers.GetPost)
		authenticated.DELETE("/posts/:postId", controllers.DeletePost)
	}
}
