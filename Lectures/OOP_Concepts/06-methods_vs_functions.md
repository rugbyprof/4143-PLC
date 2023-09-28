## Methods V Functions

In Go, functions and methods are both used to define blocks of code that perform a specific task, but they have distinct characteristics and use cases. Here's a detailed comparison of functions and methods in Go:

1. **Functions**:

   - **Definition**:
     - Functions in Go are standalone blocks of code that can be defined at the package level or within other functions.
     - They are declared using the `func` keyword, followed by the function name, a list of parameters (if any), a return type (if any), and the function body.

   - **Usage**:
     - Functions can be called from anywhere within the package where they are defined or from other packages after they are exported (start with an uppercase letter).
     - They are invoked by calling the function name followed by parentheses and providing arguments if required.

   - **Example**:
     ```go
     package main

     import "fmt"

     // Function definition
     func add(a, b int) int {
         return a + b
     }

     func main() {
         result := add(3, 4)
         fmt.Println(result) // Output: 7
     }
     ```

   - **Receiver**:
     - Functions do not have a receiver. They operate on the arguments passed to them.

2. **Methods**:

   - **Definition**:
     - Methods in Go are functions associated with a particular type, and they are defined within the scope of a type.
     - They are declared similarly to functions but with an additional receiver parameter that specifies the type on which the method operates.

   - **Usage**:
     - Methods are called on instances of the type they are defined for.
     - They are invoked using the dot notation on an instance of the type.

   - **Example**:
     ```go
     package main

     import "fmt"

     // Define a custom type
     type Rectangle struct {
         Width  float64
         Height float64
     }

     // Method definition for the Rectangle type
     func (r Rectangle) Area() float64 {
         return r.Width * r.Height
     }

     func main() {
         rect := Rectangle{Width: 3.0, Height: 4.0}
         area := rect.Area()
         fmt.Println(area) // Output: 12.0
     }
     ```

   - **Receiver**:
     - Methods have a receiver parameter, which specifies the type on which the method is defined.
     - The receiver parameter is placed in parentheses before the method name, and it allows the method to access and manipulate the fields of the receiver type.

3. **Key Differences**:

   - Functions are standalone, while methods are associated with a type.
   - Functions can be called from anywhere, while methods are called on instances of their associated types.
   - Methods have access to the fields and methods of the receiver type, while functions operate only on their parameters.
   - Functions can have any name, but methods must have unique names within their type.
   - Methods enable you to define behavior specific to a type, leading to more structured and organized code.

Functions in Go are standalone and operate on their parameters, while methods are tied to specific types and operate on instances of those types. Methods enable you to define behavior that is closely related to a type, making your code more organized and expressive.

[Previous: 05-struct_vs_interface.md](05-struct_vs_interface.md) | [Next: 07-functions.md](07-functions.md)
