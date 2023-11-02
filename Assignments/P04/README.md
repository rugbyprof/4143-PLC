## Program 4 - Concurrent Image Downloader
#### Due: 11-17-2023 (Friday @ 11:59 p.m.) 

#### Overview:
The goal of this assignment is to understand and implement basic concurrency in Go. You will write a program that concurrently downloads a set of images from given URLs and saves them to disk. By comparing the time taken to download images sequentially vs. concurrently, you will observe the benefits of concurrency for I/O-bound tasks.

#### Requirements:

1. Your program should read a list of image URLs from a file. 
2. I'm going to try and create three different sets so you don't have to hammer the same servers all the time.
3. You must implement two versions of the downloader:
   - Sequential Version: Downloads and saves each image one after the other.
   - Concurrent Version: Downloads and saves images concurrently using goroutines.
4. Save the downloaded images to disk with unique names.
5. Ensure proper error handling for failed downloads or any other issues that might arise.
6. Measure and print out the time taken by each version to complete the downloads and include this information in your README for this project. 

#### Starter Code:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
    // TODO: Implement sequential download logic
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
    // TODO: Implement concurrent download logic
}

func main() {
    urls := []string{
        "https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
        "https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640"
        // Add more image URLs
    }

    // Sequential download
    start := time.Now()
    downloadImagesSequential(urls)
    fmt.Printf("Sequential download took: %v\n", time.Since(start))

    // Concurrent download
    start = time.Now()
    downloadImagesConcurrent(urls)
    fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
    // TODO: Implement download logic
    return nil
}
```

### Basic Idea:

- For the sequential version, loop through the URLs and download each one using the `downloadImage` helper function.
- For the concurrent version, use goroutines to initiate downloads in parallel. You can use a wait group to wait for all downloads to complete.
- Use channels to handle errors that might occur in goroutines.
- The `downloadImage` function should handle the actual downloading and saving of the image. Use standard Go libraries like `net/http` for making requests and `os` for file operations.
- Here is a replit that downloads an image: https://replit.com/@rugbyprof/4143httprequest#main.go

### Finding Images

There are several websites where you can find royalty-free images that your students can use for the assignment. Here are a few popular ones:

1. **Unsplash** (https://unsplash.com/): Offers a large collection of high-resolution photos that are free to use.
2. **Pixabay** (https://pixabay.com/): Provides a vast library of images, vector graphics, and videos available for free.
3. **Pexels** (https://www.pexels.com/): Another great source for free stock photos and videos.
4. **StockSnap.io** (https://stocksnap.io/): Offers a wide array of free, high-resolution stock photos.
5. **Burst by Shopify** (https://burst.shopify.com/): Free stock photos for websites and commercial use.

Everyone needs to find 5 images and post them on this: [google sheet](https://docs.google.com/spreadsheets/d/14QSmG84jRhEgzXb96Ie9tw9zS7cY4PGXX9NGys-kc-o/edit?usp=sharing)

- Try and make sure the images are royalty free, and don't get them all from the same site! 
- Spread your urls around between either those links I provided above, or find your own royalty free images. 
- Just remember, there will be 30+ students writing code and testing it. This means everyone will be attempting to download the same set of images possibly at the same time, over and over again. So we should spread the URL's around to different sites as a courtesy.
- Lastly, remember the focus is on concurrency, so don't get huge files! Lots of smallish images will better serve you to complete the assignment. 

## Resources

Go Tour on Channels: https://tour.golang.org/concurrency/2
Go Tour on Goroutines: https://tour.golang.org/concurrency/1
Go By Example: https://gobyexample.com/http-clients

## Deliverables

- Add your `P04` folder with all its files to your assignments folder.
- Update your repo on github. 
- Make sure you have a README file describing this project and all its files as described [HERE](../../Resources/03-Readmees/README.md)


