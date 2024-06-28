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
		authenticated.POST("/posts", controllers.CreatePost)
		authenticated.GET("/posts/:postId", controllers.GetPost)
		authenticated.PUT("/posts/:postId", controllers.UpdatePost)
		authenticated.DELETE("/posts/:postId", controllers.DeletePost)
		authenticated.POST("/posts/:postId/comments", controllers.CreateComment)
		authenticated.PUT("/posts/:postId/comments/:commentId", controllers.UpdateComment)
		authenticated.DELETE("/posts/:postId/comments/:commentId", controllers.DeleteComment)
		authenticated.GET("/comments/:commentId", controllers.GetComment)
		authenticated.GET("/posts/:postId/comments", controllers.GetComments)
	}
}
