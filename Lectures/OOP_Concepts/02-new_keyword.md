## Keywords in Go

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

There are several useful method naming conventions in Go, in addition to the "new" keyword convention for constructor-like methods. These conventions help make your code more idiomatic, readable, and maintainable. 

1. **MakeType**: The "MakeType" convention is similar to "NewType" but is often used when creating instances of slices, maps, or channels. For example:

   ```go
   func MakeSlice(length, capacity int) []int {
       return make([]int, length, capacity)
   }
   ```

   This convention indicates that the function is creating and initializing a composite data type.

2. **TypeFromX**: This convention is used when creating instances of a type from another type. For example, if you have a method that creates a `Person` from a `User`, you might name it `PersonFromUser`.

   ```go
   func PersonFromUser(user User) Person {
       // Convert and initialize a Person from a User
   }
   ```

   This naming convention is clear about the source type used for creating the target type.

3. **Init**: The "Init" convention is used for methods that perform initialization of an existing instance. It's often used for complex initialization logic. For example:

   ```go
   func (p *Person) Init(name string, age int) {
       p.Name = name
       p.Age = age
   }
   ```

   This naming convention emphasizes that the method initializes an existing object.

4. **IsX or HasX**: These conventions are used for methods that return boolean values indicating whether an object possesses certain characteristics or properties. For example:

   ```go
   func (p *Person) IsAdult() bool {
       return p.Age >= 18
   }
   ```

   This naming convention is often used for query methods that return true or false.

5. **SetX and GetX**: For methods that set or get specific fields or properties of an object, the "SetX" and "GetX" conventions are common. For example:

   ```go
   func (p *Person) SetName(name string) {
       p.Name = name
   }

   func (p *Person) GetName() string {
       return p.Name
   }
   ```

   These conventions make it clear whether a method is performing a set or get operation.

6. **WithX**: This convention is used for methods that return a modified copy of an object with specific changes. It is often used for method chaining and immutability. For example:

   ```go
   func (p *Person) WithAge(age int) *Person {
       pCopy := *p // Create a copy of the original person
       pCopy.Age = age
       return &pCopy
   }
   ```

   This naming convention is common for fluent interfaces.

These are some of the common method naming conventions in Go. Choosing an appropriate convention depends on the specific use case and the intended semantics of the method. Consistency in naming conventions within your codebase is essential for readability and maintainability.

[Previous: 01-not_oop.md](./01-not_oop.md) | [Next: 03-slices_of_structs.md](./03-slices_of_structs.md)
