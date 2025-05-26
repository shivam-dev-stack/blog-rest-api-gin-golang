package main

import (
	"gin-blog-api/config"
	"gin-blog-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Connect to the database
	config.ConnectDatabase()
	// Register routes
	routes.RegisterRoutes(r)
	// Start the server
	// Listen and serve on
	// http://localhost:8080

	r.Run(":8080")
}
