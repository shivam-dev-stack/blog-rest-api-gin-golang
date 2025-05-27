package routes

import (
	"gin-blog-api/handlers"
	"gin-blog-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	router.GET("/posts", handlers.GetPosts)
	router.GET("/posts/:id", handlers.GetPostByID)

	auth := router.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.POST("/posts", handlers.CreatePost)
		auth.PUT("/posts/:id", handlers.UpdatePost)
		auth.DELETE("/posts/:id", handlers.DeletePost)
		auth.POST("/logout", handlers.Logout)
		auth.GET("/protected", handlers.ProtectedHandler)
	}
}
