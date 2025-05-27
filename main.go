package main

import (
	"gin-blog-api/config"
	"gin-blog-api/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Environment variables loaded successfully")
	}

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
