package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Serve static files (e.g., CSS, JavaScript) from a directory named "static"
	r.Static("/static", "./static")

	// Define a route to display the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Define a route to handle image downloads
	r.POST("/download", func(c *gin.Context) {
		// Get the URL entered by the user
		imageURL := c.PostForm("imageURL")

		// You can add your image downloading logic here
		// For simplicity, we'll just return the entered URL for now
		c.String(http.StatusOK, "Image URL entered: %s", imageURL)
	})

	// Run the web server on port 8080
	r.Run(":8080")
}
