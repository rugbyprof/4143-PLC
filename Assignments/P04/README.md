# ON HOLD!!!!

Here's a walkthrough on how to structure your Go project incrementally while focusing on object-oriented principles:

### Step 1: Project Setup
1. Create a new Go module for your project:

   ```bash
   go mod init webdata
   ```

2. Set up your project directory structure. You can start with a basic structure like this:

   ```
   webdata/
   ├── main.go
   ├── weather/
   │   ├── weather.go
   │   └── weather_test.go
   ├── sports/
   │   ├── sports.go
   │   └── sports_test.go
   ├── data/
   │   ├── data.go
   │   └── data_test.go
   └── api/
       ├── api.go
       └── api_test.go
   ```

### Step 2: Creating Data Packages
1. Start with the `data` package to encapsulate common data handling functionality. This package might include structures and functions for managing data types like JSON or XML.

2. Create separate packages (`weather` and `sports`) for fetching and handling specific types of data. Implement object-oriented principles such as encapsulation by defining methods for these packages.

### Step 3: Abstraction and Inheritance
1. Introduce an abstract interface, let's call it `DataFetcher`, that defines common methods like `FetchData()`.

   ```go
   // data/data.go
   package data

   type DataFetcher interface {
       FetchData() ([]byte, error)
   }
   ```

2. Implement this interface in both the `weather` and `sports` packages. Each package can have its own data source and data processing logic while adhering to the common interface.

### Step 4: Web Interface and API
1. Create an `api` package that will serve as the web interface for your project. Use a Go web framework like "gin-gonic" or "chi" to build your API.

   ```go
   // api/api.go
   package api

   import (
       "github.com/gin-gonic/gin"
   )

   func StartAPI() {
       router := gin.Default()
       // Define routes and handlers for interacting with data packages
       // e.g., router.GET("/weather", weatherHandler)
       // e.g., router.GET("/sports", sportsHandler)
       router.Run(":8080")
   }
   ```

2. Implement handlers for your API routes. These handlers should utilize the functionality provided by your `weather` and `sports` packages to serve data.

### Step 5: Incremental Improvement
1. Continue to add more data packages as needed, like fetching news data or stock market data, while ensuring they implement the `DataFetcher` interface.

2. Extend your API to support these new data packages.

3. Encourage students to explore how to apply object-oriented principles like composition and polymorphism in Go.

### Step 6: Testing
1. Write comprehensive unit tests for your packages (`weather`, `sports`, and `data`) to ensure their correctness.

2. Consider using Go's testing framework along with external libraries like "testify" for assertion methods.

This incremental approach will help students understand how to structure a Go project with object-oriented principles, even in a language that is not inherently object-oriented. It will also provide them with valuable experience in building a web interface and API using Go.


===

Sites categories with available data: 

1. **OpenWeatherMap API:** OpenWeatherMap provides weather data via a RESTful API. You can use Go's standard library to make HTTP requests to their API and retrieve weather information.

2. **Weather Underground API:** Weather Underground, now part of The Weather Channel, offers a weather data API that you can access using Go's HTTP client.

3. **Dark Sky API (Now Apple Weather API):** Dark Sky (now part of Apple) offers hyperlocal weather data. You can access this data using Go and the HTTP client.

4. **ClimaCell (Tomorrow.io) API:** ClimaCell, now known as Tomorrow.io, provides weather data with a focus on hyper-local, minute-by-minute forecasts. They offer an API that you can integrate into your Go project.

5. **Sports Open Data API:** Sports Open Data provides various sports-related data, including results, fixtures, and more. You can make HTTP requests to their API in your Go project.

6. **TheSportsDB API:** TheSportsDB offers sports-related data, including information about teams, players, and events. You can use the Go HTTP client to access their API.

7. **ESPN API:** While ESPN doesn't officially provide a public API, some unofficial wrappers and scrapers have been created by the community to access sports data from ESPN. Be aware that the availability and reliability of unofficial APIs may vary.

8. **Government Open Data Portals:** Many governments around the world provide open data portals where you can find a wide range of datasets on topics such as demographics, economics, transportation, and health. Examples include data.gov (U.S.), data.gov.uk (U.K.), and data.gouv.fr (France).

9. **Social Media APIs:** Social media platforms like Twitter and Reddit offer APIs that allow you to access public data, such as tweets or Reddit posts. You can use these APIs to gather data for social media analysis or sentiment analysis projects.

10. **Public APIs:** Explore public APIs that offer free access to data on various topics. For example, the NASA API provides access to space-related data, and the Google Books API offers access to book-related information.

11. **Data Aggregator Websites:** Some websites curate and aggregate datasets from various sources. Websites like Kaggle, Data.gov, and Awesome-Public-Datasets on GitHub provide links to a wide range of datasets.

12. **Educational Institutions:** Universities and research institutions often publish datasets related to their research. These datasets are typically freely available for educational and research purposes.

13. **Nonprofit Organizations:** Nonprofit organizations and research institutes may offer datasets related to social issues, health, and other research areas. These datasets are often available for free.

14. **Weather APIs:** Beyond weather data, some weather APIs provide historical weather data for free. This data can be used for climate analysis or weather-related projects.

15. **Sports Data APIs:** In addition to the mentioned sports data APIs, some sports organizations offer free access to certain types of sports data, such as schedules and scores.

16. **Machine Learning Datasets:** Platforms like UCI Machine Learning Repository and OpenAI's GPT-3 Playground provide datasets for machine learning and natural language processing projects.

17. **Web Scraping:** You can write web scrapers using libraries like BeautifulSoup (for Python) or GoQuery (for Go) to extract data from websites that do not offer APIs. Be sure to respect website terms of use and robots.txt rules when scraping.

18. **Educational Datasets:** Websites like UCI's Education Datasets provide educational-related datasets, which can be useful for educational research or creating educational tools.

When working with freely available datasets, it's essential to consider the licensing terms and any usage restrictions associated with the data. Always adhere to data usage policies and provide proper attribution when required.

====

Example Http Request for Kaggle

1. **HTTP Requests**: You can use Go's standard library package `net/http` to make HTTP requests to Kaggle's website or API endpoints.

2. **API Authentication**: Kaggle provides an API that requires authentication. To access Kaggle datasets programmatically, you'll need to create a Kaggle account, generate an API token, and use it to authenticate your requests.

3. **Parsing Data**: Once you retrieve data from Kaggle, you may need to parse it, depending on the format (e.g., CSV, JSON, etc.). You can use Go's standard library packages like `encoding/csv` or `encoding/json` for this purpose.

Here's a simplified example of how you might use Go to download a Kaggle dataset:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // Kaggle dataset URL
    datasetURL := "https://www.kaggle.com/someuser/somedataset/download"

    // Create an HTTP client
    client := http.Client{}

    // Make a GET request to the dataset URL
    resp, err := client.Get(datasetURL)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Create a file to save the downloaded dataset
    file, err := os.Create("dataset.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Copy the response body (dataset) to the file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Dataset downloaded successfully.")
}
```






### Http Requests Tutorial

This is just a couple of sources for getting items from the web (http requests)

- https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go
- https://zetcode.com/golang/getpostrequest/


# Ascii Art Project

#### Step 1: Project Setup

1. Create a new directory for your project, e.g., "image-to-asciiart."

2. Inside your project directory, create three subdirectories: "cmd," "imageutil," and "asciiart." Your project structure should look like this:

   ```
   image-to-asciiart/
   ├── cmd/
   ├── imageutil/
   └── asciiart/
   ```

#### Step 2: Initialize the Project and Modules

1. Open your terminal and navigate to your project directory, e.g., `cd path/to/image-to-asciiart`.

2. Run the following command to initialize your project as a Go module:

   ```bash
   go mod init image-to-asciiart
   ```

   This creates a `go.mod` file that will help manage your project's dependencies.

#### Step 3: Create Packages

In your "imageutil" and "asciiart" directories, create Go files (`imageutil.go` and `asciiart.go`) with the functions you need. For example, in "imageutil/imageutil.go":

```go
package imageutil

import (
    "fmt"
)

func DownloadImage(url, filename string) error {
    // Your code to download an image goes here
    fmt.Printf("Downloading image from %s to %s...\n", url, filename)
    return nil
}
```

And in "asciiart/asciiart.go":

```go
package asciiart

import (
    "fmt"
)

func ConvertToASCII(imagePath string, cols, rows int) (string, error) {
    // Your code to convert an image to ASCII art goes here
    fmt.Printf("Converting image at %s to ASCII art (%d columns, %d rows)...\n", imagePath, cols, rows)
    return "ASCII art result", nil
}
```

#### Step 4: Create the Main Program

In your "cmd" directory, create a `main.go` file where you will use the functions from your packages:

```go
package main

import (
    "fmt"
    "image-to-asciiart/imageutil"
    "image-to-asciiart/asciiart"
)

func main() {
    // Your program logic here

    // Example usage of the imageutil package
    imageURL := "URL_OF_THE_IMAGE"
    err := imageutil.DownloadImage(imageURL, "image.jpg")
    if err != nil {
        fmt.Println("Error downloading image:", err)
        return
    }

    // Example usage of the asciiart package
    asciiArt, err := asciiart.ConvertToASCII("image.jpg", 100, 50)
    if err != nil {
        fmt.Println("Error converting to ASCII art:", err)
        return
    }

    // Print the ASCII art to the console
    fmt.Println(asciiArt)
}
```

#### Step 5: Run Your Program

Now you can run your program by navigating to the "cmd" directory and executing `go run main.go`:

```bash
cd path/to/image-to-asciiart/cmd
go run main.go
```

This will execute your program, which uses the functions from the "imageutil" and "asciiart" packages.

#### GitHub Option (Optional):

>You can read [this Chat GPT exchange](../../Lectures/module_conversation.md) if you want to know more about modules on github.

If you want to host your packages on GitHub, you can do the following:

1. Create separate GitHub repositories for each package (e.g., "imageutil" and "asciiart").

2. Push the package code to their respective GitHub repositories.

3. To use these packages in your "image-to-asciiart" project, you can modify your import paths in the `main.go` file to point to the GitHub repositories:

   ```go
   import (
       "fmt"
       "github.com/yourusername/imageutil" // Replace with your GitHub username
       "github.com/yourusername/asciiart"  // Replace with your GitHub username
   )
   ```

   Then run `go get` to fetch the packages:

   ```bash
   go get github.com/yourusername/imageutil
   go get github.com/yourusername/asciiart
   ```

   Finally, you can use them as described in the previous steps.

I hope this step-by-step guide clarifies the process of setting up a Go project with local packages and, if needed, GitHub-hosted packages.