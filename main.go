package main

import (
	"example.com/comments-api/controllers"
	"example.com/comments-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// Set up the router
	r := gin.Default()

	// Routes
	r.POST("/comments", controllers.CreateComment)
	r.GET("/comments", controllers.GetComments)
	r.DELETE("/comments/:id", controllers.DeleteComment)

	// Start the server
	r.Run(":8080")
}
