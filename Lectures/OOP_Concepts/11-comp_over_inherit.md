## Composition Over Inheritance

In Go, you typically favor `composition` over `inheritance`, which means you build complex types by `composing` (or `embedding`) simpler types rather than `inheriting` behavior from other types. Let's dive into a discussion about both of them and Go's approach to solving this decades long discussion. 

### Inheritance:

Inheritance is a mechanism in OOP where a new class (subclass or derived class) inherits properties and behaviors (attributes and methods) from an existing class (base class or parent class). Inheritance creates an "is-a" relationship between classes, where the subclass is a specialized version of the parent class.

**Advantages of Inheritance:**

1. **Code Reusability:** Inheritance allows you to reuse code from existing classes, reducing duplication and promoting a modular code structure.

2. **Polymorphism:** Inheritance facilitates polymorphism, where objects of derived classes can be treated as objects of the base class. This enables dynamic dispatch and runtime method binding.

3. **Structural Hierarchies:** Inheritance helps create hierarchical relationships between classes, making it easier to represent real-world concepts and entities.

**Disadvantages of Inheritance:**

1. **Tight Coupling:** Inheritance can lead to tight coupling between classes, making the code harder to maintain and less flexible to changes.

2. **Inflexibility:** Subclasses are bound to the implementation details of their parent class, making it challenging to modify or extend behavior without affecting subclasses.

3. **Inheritance Depth:** Deep inheritance hierarchies can become complex and difficult to understand, leading to the "diamond problem" and other issues.

### Composition:

Composition is a design principle that encourages building complex objects by combining simpler objects (components) rather than through inheritance. In composition, an object "has-a" relationship with its components, and it delegates some of its responsibilities to them.

**Advantages of Composition:**

1. **Flexibility:** Composition provides greater flexibility than inheritance. You can change an object's behavior by swapping or modifying its components without affecting other parts of the code.

2. **Low Coupling:** Components are loosely coupled, promoting a cleaner and more maintainable codebase.

3. **Reusability:** You can reuse components in multiple classes, fostering code reusability.

4. **Dynamic Composition:** Components can be added or removed at runtime, allowing for dynamic behavior.

**Disadvantages of Composition:**

1. **Complexity:** Complex systems built with composition may have a large number of components, making the codebase harder to manage.

2. **Increased Code:** Composition can lead to more lines of code as you create and manage objects and their components.

3. **Delegation Overhead:** The need to delegate responsibilities to components can introduce some overhead compared to direct method calls.

### Go's Approach:

Go takes a different approach, emphasizing composition over inheritance. It doesn't support traditional inheritance like some other OOP languages (e.g., Java or C++). Instead, Go encourages code reuse through composition, interfaces, and embedding.

- **Composition in Go:** You create objects by composing them from smaller structs (types). You can embed one struct type within another to reuse its fields and methods.

- **Interfaces in Go:** Interfaces define behavior, and types implement interfaces. This promotes polymorphism and code reuse without relying on inheritance.

Go's approach simplifies the language, reduces complexity, and encourages a more flexible and modular design while avoiding many of the pitfalls associated with traditional inheritance.

# Examples:
### Composition by Embedding Structs:

Let's explore this concept of composition using the goto Dungeons & Dragons (D&D) theme as an example. This will include composing structs by embedding other structs within them. and ultimately embedding slices of structs. Composition in Go is achieved by embedding structs within other structs. This allows you to reuse and compose behavior by including instances of other structs as fields. In the context of D&D, let's consider character classes as an example:

```go
package main

import "fmt"

// BaseCharacter represents basic character attributes.
type BaseCharacter struct {
    Name    string
    Level   int
    HitDice int
}

// Wizard represents a D&D wizard class.
type Wizard struct {
    BaseCharacter // Embed the BaseCharacter to reuse its fields
    Spellbook     []string // Additional field for wizard-specific data
}

// Create a new Wizard character
func NewWizard(name string, level, hitDice int, spells []string) Wizard {
    return Wizard{
        BaseCharacter: BaseCharacter{
            Name:    name,
            Level:   level,
            HitDice: hitDice,
        },
        Spellbook: spells,
    }
}

func main() {
    gandalf := NewWizard("Gandalf", 10, 6, []string{"Fireball", "Lightning Bolt"})

    // Access fields from both BaseCharacter and Wizard
    fmt.Printf("%s is a level %d wizard with %d hit dice.\n", gandalf.Name, gandalf.Level, gandalf.HitDice)
    fmt.Println("Spells in Spellbook:", gandalf.Spellbook)
}
```

In this example:

- We define a `BaseCharacter` struct representing basic character attributes and reuse it by embedding it within the `Wizard` struct.
- The `Wizard` struct adds an additional field `Spellbook` to store wizard-specific data.
- We create a new `Wizard` character using the `NewWizard` function, which initializes both the `BaseCharacter` and `Spellbook` fields.
- By embedding the `BaseCharacter`, the `Wizard` type inherits its fields and methods, allowing us to access character attributes like `Name`, `Level`, and `HitDice` directly.

### Embedding Slices of Structs:
You can also use composition to represent groups of related objects, such as a party of D&D characters:

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

// Party represents a group of D&D characters.
type Party struct {
    Members []Character // Embed a slice of Character structs
}

// Create a new D&D party
func NewParty(members []Character) Party {
    return Party{Members: members}
}

func main() {
    gandalf := Character{Name: "Gandalf", Class: "Wizard", Level: 10, HitDice: 6}
    conan := Character{Name: "Conan", Class: "Warrior", Level: 8, HitDice: 10}

    adventurers := []Character{gandalf, conan}

    party := NewParty(adventurers)

    fmt.Printf("Party Members:\n")
    for _, member := range party.Members {
        fmt.Printf("%s - %s, Level %d\n", member.Name, member.Class, member.Level)
    }
}
```

In this example:

- We define a `Character` struct representing individual D&D characters.
- The `Party` struct embeds a slice of `Character` structs to represent a group of characters.
- We create a new D&D party using the `NewParty` function, which takes a slice of character members.
- The `Party` type allows us to group and manage multiple character instances in a convenient way.

This demonstrates how composition can be used to model relationships between objects in Go, allowing you to create more flexible and modular code, which is a core principle of the language.

[Previous: 10-packages_in_depth.md](10-packages_in_depth.md) | [Next: README.md](README.md)
