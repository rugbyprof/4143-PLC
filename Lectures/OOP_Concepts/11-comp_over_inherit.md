## Composition Over Inheritance

In Go, you typically favor `composition` over `inheritance`, which means you build complex types by `composing` (or `embedding`) simpler types rather than `inheriting` behavior from other types. Let's explore this concept of composition using the goto Dungeons & Dragons (D&D) theme as an example. This will include composing structs by embedding other structs within them. and ultimately embedding slices of structs.

### Composition by Embedding Structs:

Composition in Go is achieved by embedding structs within other structs. This allows you to reuse and compose behavior by including instances of other structs as fields. In the context of D&D, let's consider character classes as an example:

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


[Previous: 10-interfaces.md](10-interfaces.md) | [Next: README.md](README.md)
