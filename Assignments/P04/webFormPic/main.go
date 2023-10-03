package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Result  string `json:"result"`
	URL     string `json:"url"`
	Name    string `json:"name"`
}

// DownloadImage downloads an image from the given URL and returns the image data.
func DownloadImage(imageURL string) ([]byte, error) {
	response, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch image. Status code %d", response.StatusCode)
	}

	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

// func SaveImage(fileName string,c *gin.Context) ([]byte, error) {

// }

func main() {
	// Create a Gin router
	r := gin.Default()

	// Serve static files (e.g., CSS, JavaScript) from a directory named "static"
	r.Static("/static", "./static/")

	// Define a route to display the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Define a route to handle image downloads
	r.POST("/download", func(c *gin.Context) {
		// Get the URL entered by the user
		imageURL := c.PostForm("imageURL")
		imageName := c.PostForm("imageName")

		fmt.Println(string(imageURL))
		fmt.Println(string(imageName))

		imageData, err := DownloadImage(imageURL)
		if err != nil {
			c.String(http.StatusBadRequest, "Error downloading image: %s", err.Error())
			return
		}

		// Create a new file to save the image
		outFile, err := os.Create("./static/downloads/" + imageName) // Change the filename and extension as needed
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating file: %s", err.Error())
			return
		}
		defer outFile.Close()

		// Copy the image data to the file
		// _, err = io.Copy(outFile, bytes.NewReader(imageData))
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, "Error writing image data to file: %s", err.Error())
		// 	return
		// }

		// Copy the image data to the file
		responseData := ResponseData{
			Error:   "",
			Message: "",
			Result:  "",
			URL:     imageURL,
			Name:    imageName,
		}
		_, err = io.Copy(outFile, bytes.NewReader(imageData))
		if err != nil {
			responseData.Error = fmt.Sprintf("Error writing image data to file: %s", err.Error())
			responseData.Result = "Failed"

		} else {
			responseData.Result = "Success"
		}

		jsonData, _ := json.Marshal(responseData)

		fmt.Println(string(jsonData))

		c.JSON(http.StatusInternalServerError, jsonData)

		// You can add your image downloading logic here
		// For simplicity, we'll just return the entered URL for now
		// c.String(http.StatusOK, "Image URL entered: %s\n\n", imageURL)
		// c.String(http.StatusOK, "Image Name entered: %s", imageName)

		// // Serve the image as a response
		// c.Data(http.StatusOK, "image/jpeg", imageData)
	})

	// Run the web server on port 8080
	r.Run(":8080")
}
