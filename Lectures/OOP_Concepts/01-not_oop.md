## Not OOP

While Go is not a traditional object-oriented language like `Java` or `Python`, it allows you to apply some object-oriented principles in a different way. Let's explore how you can use Go in an object-oriented programming (OOP) manner, including using structs instead of classes and the `New()` function instead of constructors.

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

- We define a `NewCharacter()` function that takes parameters to initialize a `Character` struct and returns an instance of it.
- We use the `NewCharacter()` function to create a D&D character (`gandalf`).
- The `New()` function is a common Go convention for struct initialization, allowing you to perform any setup logic as needed.

While Go doesn't have constructors like in traditional OOP languages, creating `New()` functions is a common way to achieve structured and controlled initialization of structs. This approach aligns with Go's philosophy of simplicity and explicitness.

 | [Next: 02-new_keyword.md](02-new_keyword.md)
