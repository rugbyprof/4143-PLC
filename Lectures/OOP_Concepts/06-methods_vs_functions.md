## Methods V Functions

There is a fundamental difference between a `method` and a `function`, although both are used to define blocks of code that can be executed. Here's the distinction:

### Function:

   - A function in Go is a self-contained block of code that performs a specific task or operation.
   - Functions are defined outside of any type and are not associated with a particular data structure or object.
   - They are typically declared using the `func` keyword, followed by a name, parameter list, return type (if any), and a code block.
   - Functions can be called independently, and they operate on their input arguments to produce results.
   - They can be either exported (capitalized) or unexported (lowercase) based on their visibility within and outside the package.

   Example of a function:
   ```go
   func Add(a, b int) int {
       return a + b
   }
   ```

### Method:

   - A method, in Go, is a function that is associated with a specific type (a struct type in most cases) and can operate on the data of that type.
   - Methods are defined as part of a type's definition and can access and modify the fields of that type.
   - They are declared with a receiver, which specifies the type on which the method operates. The receiver appears before the method name in the method declaration.
   - Methods are used to encapsulate behavior related to a specific type, and they allow you to define operations that are specific to that type.

   Example of a method:
   ```go
   type Rectangle struct {
       Width  float64
       Height float64
   }

   // Method with a receiver of type Rectangle
   func (r Rectangle) Area() float64 {
       return r.Width * r.Height
   }
   ```

   Here, `Area` is a method of the `Rectangle` type, and it operates on the `Rectangle` data.

- In summary, the key difference between a `function` and a `method` in `Go` is that a function is a standalone block of code, while a method is associated with a specific type and can access and manipulate the data of that type. 
- Methods provide a way to define behavior that is closely tied to the characteristics of a particular data structure, enhancing encapsulation and code organization.

[Previous: 05-struct_vs_interface.md](05-struct_vs_interface.md) | [Next: 07-first_class_func.md](07-first_class_func.md)
