## Program 2 - Baby Steps
#### Due: 09-21-2023 (Thursday @ 9:30 a.m.) 

## Background

### Project Structure

- A Go project can have different structures, but the one below is what I will be discussing. What you see is a `projectFolder` that contains a `main.go` file and some packages in a sub-folder. 
- This structure will be turned into a module (explained next), where the module name doesn't have to match any of the other folder / file / package names held within. 
- Remember that Go organizes projects into packages and your `main.go` file is actually `package main` (the main package).
- Notice the lack of a `go.mod` file, since we haven't run `go mod init` yet which creates said file.

```
projectFolder/
├── main.go
├── somePackage/
│   ├── randpackage.go
│   └── other_package_files.go
```

### Module Initialization:

Initialize your project as a Go module using the `go mod init` command in the root directory of your project:

```bash
go mod init myproject
```

- This adds a `go.mod` folder to our project which is in control of dependancies.
- Note: the name of the module (`myproject` in this case) doesn't have to match any folder name.
- The module name matters, since its used in imports, where my initial reaction was to use the folder name for imports, which is wrong. This is explained more next.
  
```
projectFolder/
├── go.mod
├── main.go
├── somePackage/
│   ├── randpackage.go
│   └── other_package_files.go
```

### Imports:

In your Go code, use relative import paths for local packages. For example, if `main.go` and `randpackage.go` are in the same directory, you can import `randpackage` like this:

```go
import (
    "./randpackage"
)
```

Alternatively, if `randpackage` is a subdirectory within your project, you can use the module name followed by the relative path:

```go
import (
    "myproject/randpackage"
)
```

Ensure that your import paths in your Go code match the module name specified in your `go.mod` file. This helps Go locate and manage local packages correctly.

## Project Background

In this program we will try and show some advantages of Go when compared to C++ and Python. Every high level language has its strengths, and I think Golang's can be seen in its similarity to both aforementioned languages, but more so in its library availability. Not as strong as Python's but more so than C++'s.  I would like us to create Go's version of a class, that will do some rudimentary image manipulation.

## ImageManipulator

This is a basic walkthrough to create a `package` that contains Go's version of a `class` that will be used inside a `module` with a `main.go` file that uses it. A project structured at its most basic, but will implement basic image editing, which is not something that you see everyday in lower level C++ courses.

### Structure

First we need a project folder. This will be the home of our `module` and any `packages` we need to get the job done.

1. **Create module home**:

    Create a new directory for your module. Let's call it "imagemod."

    ```bash
    mkdir imagemod
    cd imagemod
    ```

2. **Create a Package**:
   
    Create a new directory for your package. Let's call it "imageManipulator."

    ```bash
    mkdir imageManipulator
    cd imageManipulator
    ```

3. **Create the Package File**:
   Inside the "imageManipulator" directory, create a Go source file, e.g., "imageManipulator.go." Define your `ImageManipulator` struct and methods in this file.

   ```go
   // imagemod/imageManipulator/imageManipulator.go

   package imageManipulator

   import (
       "github.com/fogleman/gg"
   )

   // ImageManipulator represents an image manipulation tool.
   type ImageManipulator struct {
       Image *gg.Context
   }

   // NewImageManipulator creates a new ImageManipulator instance.
   func NewImageManipulator(width, height int) *ImageManipulator {
       img := gg.NewContext(width, height)
       return &ImageManipulator{Image: img}
   }

   // SaveToFile saves the manipulated image to a file.
   func (im *ImageManipulator) SaveToFile(filename string) error {
       return im.Image.SavePNG(filename)
   }

   // DrawRectangle draws a rectangle on the image.
   func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
       im.Image.DrawRectangle(x, y, width, height)
       im.Image.Stroke()
   }
   ```

4. **Create a `main` Function**:
   In the same module directory, you can create another Go source file, e.g., "main.go," that contains the `main` function. In this file, you can use the `imagemod` package that you've just created.

   ```go
   // imageManipulator/main.go

   package main

   import (
       "myimageapp/imagemod"
   )

   func main() {
       // Create an ImageManipulator instance
       im := imageManipulator.NewImageManipulator(800, 600)

       // Draw a rectangle
       im.DrawRectangle(150, 50, 560, 411)

       // Save the image to a file
       im.SaveToFile("mustangs.png")
   }
   ```

5. **Initialize a `go.mod` File**:
   You can initialize a Go module for your project by running the following command in the root directory of your project (where "imagemanipulator" and "main.go" are located):

   ```bash
   go mod init myimageapp
   ```

   This command will create a `go.mod` file that keeps track of your project's dependencies.

6. **Build and Run**:
   You can now build and run your program using the `go run` command:

   ```bash
   go run main.go
   ```

   This will execute the `main` function in "main.go," which uses the `imageManipulator` package to manipulate and save an image.

By structuring your code in this way, you create a reusable package (`imageManipulator`) that can be imported and used in different Go files. It allows you to keep your code organized and promotes code reuse and modularity.

## Adding another Constructor

In Go, you can't have overloaded constructors like you can in some other languages (such as Java or C++), where different constructors can take different sets of parameters. However, you can achieve similar functionality by using variadic functions and providing optional parameters. In your case, you want to create a constructor that accepts a path to an existing image. Here's how you can do it:

1. **Update the `ImageManipulator` Struct**:

   Add a field to store the path of the existing image:

   ```go
   package imagemod

   import (
       "github.com/fogleman/gg"
   )

   // ImageManipulator represents an image manipulation tool.
   type ImageManipulator struct {
       Image *gg.Context
       ImagePath string // Add a field to store the image path
   }
   ```

2. **Add a New Constructor Function**:

   Create a new constructor function that accepts the image path and initializes the `ImageManipulator` with the existing image:

   ```go
   // NewImageManipulatorWithImage loads an existing image and creates an ImageManipulator instance.
   func NewImageManipulatorWithImage(imagePath string) (*ImageManipulator, error) {
       img, err := gg.LoadImage(imagePath)
       if err != nil {
           return nil, err
       }

       ctx := gg.NewContextForImage(img)
       return &ImageManipulator{Image: ctx, ImagePath: imagePath}, nil
   }
   ```

3. **Update the `main` Function**:

   In your `main` function, you can now use this new constructor to load an existing image:

   ```go
   package main

   import (
       "fmt"
       "myimageapp/imagemod"
   )

   func main() {
       // Create an ImageManipulator instance with an existing image
       im, err := imagemod.NewImageManipulatorWithImage("mustangs.jpg")
       if err != nil {
           fmt.Println("Error:", err)
           return
       }

       // Draw a rectangle
       im.DrawRectangle(150, 50, 560, 411)

       // Save the modified image to a file
       im.SaveToFile("mustangs.png")
   }
   ```



## Deliverables

- Add your `P02` folder with all its files to your assignments folder.
- Update your repo on github. 
- Make sure you have a README file describing this project and all its files.
