package controllers

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
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

// FetchAndSavePosts fetches posts from an external API and saves them to the database
func FetchAndSavePosts(c *gin.Context) {
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	// Fetch data from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		log.Println("Error fetching posts:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		log.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var posts []models.Comment
	if err := json.Unmarshal(body, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	// Save posts to the database
	result := database.DB.Create(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save posts"})
		log.Println("Error saving posts:", result.Error)
		return
	}

	log.Printf("Successfully saved %d posts\n", len(posts))
	c.JSON(http.StatusCreated, gin.H{"message": "Posts fetched and saved successfully", "count": len(posts)})
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
	id := c.Query("id") // Get the 'id' from query parameters
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID query parameter is required"})
		log.Println("ID query parameter missing")
		return
	}

	// Delete the comment with the specified ID
	result := database.DB.Delete(&models.Comment{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		log.Printf("Comment with ID %s not found\n", id)
		return
	}

	log.Printf("Comment with ID %s deleted successfully\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
