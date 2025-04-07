package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// create a Gin router
	router := gin.Default()

	// Add a basic GET route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mesage": "🚀 Welcome to the Gin API!",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
