## Not OOP

While Go is not a traditional object-oriented language like `Java`, `C++` or `Python`, it allows you to apply some object-oriented principles in a different way. Let's explore how you can use Go in an object-oriented programming (OOP) manner, including using structs instead of classes and the `New()` function instead of constructors.

### Structs Instead of Classes:

In Go, you use structs to define data structures that encapsulate both data and methods (functions) that operate on that data. While Go doesn't have classes, structs serve a similar purpose and allow you to define types with fields and methods.

In a Dungeons & Dragons (D&D) context, let's create a struct to represent a D&D character:

```go
package main

import "fmt"

// Character represents a D&D character.
type Character struct {
    Name    string
    Class   string
    Level   int
    HitDice int
}

// Method associated with the Character struct
func (c Character) PrintInfo() {
    fmt.Printf("Name: %s\nClass: %s\nLevel: %d\nHit Dice: %d\n", c.Name, c.Class, c.Level, c.HitDice)
}

func main() {
    // Create a D&D character using a struct literal
    gandalf := Character{
        Name:    "Gandalf",
        Class:   "Wizard",
        Level:   10,
        HitDice: 6,
    }

    // Call a method on the character
    gandalf.PrintInfo()
}
```

In this example:

- We define a `Character` struct with fields for the character's name, class, level, and hit dice.
- We define a method `PrintInfo()` associated with the `Character` struct to print the character's information.
- We create a D&D character (`gandalf`) using a struct literal and call the `PrintInfo()` method on it.

While Go doesn't have classes and inheritance in the traditional sense, structs provide a way to encapsulate data and behavior together, similar to objects in OOP.

### `New()` Function Instead of Constructors:

In Go, instead of constructors, it's a common practice to create a `New()` function to initialize and return an instance of a struct. This allows you to perform any necessary setup before returning the struct.

In our standard D&D example, let's create a `NewCharacter()` function:

```go
package main

import "fmt"

// Character represents a D&D character.
type Character struct {
    Name    string
    Class   string
    Level   int
    HitDice int
}

// Method associated with the Character struct
func (c Character) PrintInfo() {
    fmt.Printf("Name: %s\nClass: %s\nLevel: %d\nHit Dice: %d\n", c.Name, c.Class, c.Level, c.HitDice)
}

// NewCharacter creates and initializes a new Character instance.
func NewCharacter(name, class string, level, hitDice int) Character {
    return Character{
        Name:    name,
        Class:   class,
        Level:   level,
        HitDice: hitDice,
    }
}

func main() {
    // Create a D&D character using the NewCharacter() function
    gandalf := NewCharacter("Gandalf", "Wizard", 10, 6)

    // Call a method on the character
    gandalf.PrintInfo()
}
```

In this updated example:

- We define a `NewCharacter()` function that takes parameters to initialize a `Character` struct and returns an instance of it. NewCharacter() can be thought of as the 'constructor' of the 'Character' class in C++. But in Golang, you define the struct that represents your data, and then tie functions to it. 
- We use the `NewCharacter()` function to create a D&D character (`gandalf`).
- The `New()` function is a common Go convention for struct initialization, allowing you to perform any setup logic as needed.

While Go doesn't have constructors like in traditional OOP languages, creating `New()` functions is a common way to achieve structured and controlled initialization of structs. This approach aligns with Go's philosophy of simplicity and explicitness. 

There are really 11 main points to how Go differs from what we are used to in Python and C++. Here are those to be aware of as you go throughout the rest of this guide. 

1. **No Classes, Only Structs**:
   - In Go, you define data structures using structs. These structs can have fields and methods.
   - Methods are functions attached to a struct that can operate on instances of that struct.
   - While Python and C++ use classes to define objects and their behaviors, Go uses a more lightweight and flexible approach with structs and methods.

2. **No Inheritance Hierarchy**:
   - Go does not support traditional class-based inheritance. There are no parent-child class relationships with "is-a" semantics.
   - Instead of inheritance, Go encourages composition, where types are built by embedding other types (structs) or by implementing interfaces.
   - This simplifies code and avoids complex inheritance hierarchies.

3. **Interfaces for Polymorphism**:
   - Polymorphism in Go is achieved through interfaces, which define a set of method signatures that a type must implement.
   - Types in Go are not explicitly tied to interfaces but can satisfy interfaces implicitly based on method signatures.
   - This allows for more flexible and duck-typed polymorphism compared to statically declared inheritance hierarchies in C++.

4. **No Constructors or Destructors**:
   - Go does not have constructors or destructors like C++. Instead, you use factory functions to create instances and rely on garbage collection for resource management.
   - Initialization methods, such as `func NewType(...)` conventions, are often used in Go to initialize struct instances.

5. **No Operator Overloading**:
   - Go does not support operator overloading, unlike C++ where you can define custom behavior for operators.
   - Go prefers the use of standard functions and methods for operations, making code more explicit and readable.

6. **Strong Typing**:
   - Go enforces strong static typing, similar to C++. Type safety is checked at compile-time.
   - The type system helps catch errors early, making Go code less error-prone compared to dynamically typed languages like Python.

7. **Minimalist Syntax**:
   - Go's syntax is designed to be minimalistic and easy to read. It promotes readability and reduces verbosity.
   - This simplicity aids in maintaining clean and straightforward code, especially in large projects.

8. **Concurrency Model**:
   - Go introduces a unique concurrency model with goroutines and channels.
   - While not directly related to OOP, it's a powerful feature for writing concurrent and parallel code.

9. **No Method Overloading**:
    - Go does not support method overloading based on parameter types. Each method must have a unique name.
    - Overloading is achieved through different method names or by using variadic functions.

10. **Composition over Inheritance**:
    - Go promotes composition over inheritance. You can create complex types by embedding structs or by implementing interfaces.
    - This simplifies code structure, reduces complexity, and encourages code reuse.

11. **Garbage Collection**:
    - Go uses automatic garbage collection to manage memory, similar to Python. Developers don't need to manage memory manually.
    - This reduces the risk of memory-related bugs, making Go code more reliable than manual memory management in C++.

Go offers a unique approach to OOP that emphasizes simplicity, composition, strong typing, and interfaces for polymorphism. It avoids some of the complexities associated with class-based inheritance and introduces innovative features like goroutines and channels for concurrency. While Go differs from Python and C++ in several ways, it provides a practical and efficient toolset for building scalable and maintainable software.

 | [Next: 02-new_keyword.md](./02-new_keyword.md)
