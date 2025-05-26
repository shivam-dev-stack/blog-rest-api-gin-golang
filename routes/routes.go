package routes

import (
	"gin-blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	router.GET("/posts", handlers.GetPosts)
	router.GET("/posts/:id", handlers.GetPostByID)
	router.POST("/posts", handlers.CreatePost)
	router.PUT("/posts/:id", handlers.UpdatePost)
	router.DELETE("/posts/:id", handlers.DeletePost)
}
