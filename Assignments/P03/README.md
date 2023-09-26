## Program 3 - Image Getter
#### Due: 10-03-2023 (Tuesday @ 9:30 a.m.) 

## Background

### Project Structure

- We are are going to keep using the same project structure as the last program. 
- But we are going to add a little functionality and give you guys a little challenge!
- Remember `projectFolder` is your `module`, and modules use `packages` to add functionality.

```
projectFolder/
├── main.go
├── go.mod
├── somePackage/
│   ├── package1.go
│   └── package2.go
```

## Overview

I want to use the code from P02 and add functionality that goLang is supposed to make `easy`. However I think all of us are realizing how important dev environments are and learning the intricacies of a new language. The previous program just opened an image and drew a rectangle on said image. Lets add some convenience and robustness to our code:

- Auto download an image from the web
- Run a small web server that provides a form to enter our url. 

### Future

- Async requests
- Parallel downloads


### Http Requests Tutorial
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