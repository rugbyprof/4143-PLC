package main

import (
    "fmt"
    "net/http"
    "os"
    "io"
)

func main() {
    // URL of the image you want to download
    //imageUrl := "https://images.pexels.com/photos/17929271/pexels-photo-17929271/free-photo-of-woman-standing-on-vineyard.jpeg"
    imageUrl := "https://s3e8p5g8.rocketcdn.me/wp-content/uploads/2020/11/midwestern-state-university2.jpg"

    // Create an HTTP GET request
    response, err := http.Get(imageUrl)
    if err != nil {
        fmt.Println("Error making the request:", err)
        return
    }
    defer response.Body.Close()

    // Check if the response status code is OK (200)
    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Status code", response.StatusCode)
        return
    }

    // Create a new file to save the image
    outputFile, err := os.Create("downloaded_image.jpg")
    if err != nil {
        fmt.Println("Error creating the file:", err)
        return
    }
    defer outputFile.Close()

    // Copy the HTTP response body to the file
    _, err = io.Copy(outputFile, response.Body)
    if err != nil {
        fmt.Println("Error saving the image:", err)
        return
    }

    fmt.Println("Image downloaded and saved as 'downloaded_image.jpg'")
}
