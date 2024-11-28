package controllers

import (
	// "fmt"
	"log"
	"net/http"

	"example.com/comments-api/database"
	"example.com/comments-api/models"
	"github.com/gin-gonic/gin"
)

// CreateComment handles the creation of a comment
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error binding JSON:", err)
		return
	}
	database.DB.Create(&comment)
	log.Printf("Comment Created: %+v\n", comment)
	c.JSON(http.StatusCreated, comment)
}

// GetComments fetches all comments
func GetComments(c *gin.Context) {
	var comments []models.Comment
	database.DB.Find(&comments)
	log.Printf("Fetched %d Comments\n", len(comments))
	for _, comment := range comments {
		log.Printf("Fetched Comment: %+v\n", comment)
	}
	c.JSON(http.StatusOK, comments)
}

// DeleteComment handles deletion of a comment by ID
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Comment{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		log.Printf("Comment with ID %s not found\n", id)
		return
	}
	log.Printf("Comment with ID %s deleted successfully\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
