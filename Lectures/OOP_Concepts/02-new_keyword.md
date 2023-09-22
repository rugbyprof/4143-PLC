## The New Keyword in Go

In Go, naming conventions such as using `New` in front of a function name and capitalizing method names have specific meanings and significance.

### Naming Convention with the `New` Prefix:

   In Go, it's a common convention to use the `New` prefix in front of a function name to create and initialize a new instance of a type. This is often used as a constructor function to encapsulate the details of creating an object.

   - Example: `NewCharacter`, `NewWizard`, and `NewWarrior` in the provided code.

   This convention is useful for the following reasons:
   
   - It clearly indicates that a new instance of a type is being created.
   - It abstracts the initialization logic and allows you to change it internally without affecting the calling code.
   - It helps maintainers of your code understand the purpose of the function.

### Capitalized Method Names:

   In Go, method names (functions associated with a type) are capitalized if they are intended to be exported (i.e., accessible from outside the package). If a method name is lowercase, it is treated as unexported and can only be accessed within the same package.

   - Example: `Attack` method in the provided code is capitalized, making it an exported method.

   The significance of capitalizing method names:
   
   - Exported methods are part of the type's public API, so they are accessible to other packages.
   - Unexported methods are private to the package and cannot be used outside of it. This provides encapsulation by restricting access to internal implementation details.

- These naming conventions contribute to code clarity, maintainability, and adherence to Go's principles of simplicity and readability. - They help developers understand the intended use of functions and methods while also ensuring proper encapsulation and visibility control.

[Previous: 01-not_oop.md](01-not_oop.md) | [Next: 03-slices_of_structs.md](03-slices_of_structs.md)
