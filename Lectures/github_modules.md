## Github Modules

Creating a Go module and publishing it on GitHub involves a few steps. Here's a quick rundown:

### 1. Create a GitHub Repository

1. Create a new repository on GitHub. Let's call it `my-module`.
2. Clone the repository locally.

    ```bash
    git clone https://github.com/student123/my-module.git
    cd my-module
    ```

### 2. Initialize Go Module

In your local repository folder, run the following command:

```bash
go mod init github.com/student123/my-module
```

This will create a `go.mod` file, declaring the module path.

### 3. Add Some Go Code

Create a new Go file (e.g., `main.go`) and write some Go code that you want to include in your module.

```go
// main.go

package main

import "fmt"

func HelloWorld() {
    fmt.Println("Hello, world!")
}
```

### 4. Commit and Push

Commit these changes to your local git repository and push them to GitHub.

```bash
git add .
git commit -m "Initial commit"
git push origin main
```

### 5. Tagging a Version

To set a version number, you tag your Git repository using semantic versioning. Here's how to tag a v1.0.0 version:

```bash
git tag v1.0.0
git push --tags
```

### 6. Updating the Module

For any future changes, update your code and push a new tag.

1. Make changes to your code.
2. Commit and push those changes.
3. Update the version tag.

    ```bash
    git tag v1.0.1
    git push --tags
    ```

This way, you can maintain versioning for your Go module. Users can specify the version they want to use or update to the latest version easily.

### 7. Importing Your Module

Now, you or someone else can import your module using the version tag.

In a different Go project:

```bash
go get github.com/student123/my-module@v1.0.0
```

Or in the `go.mod` file:

```go
require github.com/student123/my-module v1.0.0
```

```go
package main

import (
    "github.com/student123/my-module@v1.0.0"
)

func main() {
    HelloWorld()
}
```