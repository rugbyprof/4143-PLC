## Inheritance

 In Go there are not traditional classes along with inheritance as in traditional object-oriented languages like Java or C++. But you can achieve similar concepts using `struct`s embedded with other structs and also the `interface` type the enforce certain behaviours.

 Below is an example of a base character extended to be a `Wizard` and a `Warrior` in `Go`, while emphasizing the principles of Polymorphism, Encapsulation, and Abstraction.

 Let's assume the following directory structure:

 ```
DandD/
├── main.go
├── character/
│   ├── character.go
│   └── character.go (implementation for common character-related logic)
├── wizard/
│   ├── wizard.go
│   └── wizard.go (implementation for wizards)
├── warrior/
│   ├── warrior.go
│   └── warrior.go (implementation for warriors)
└── ...
```

### Example in folder DandD 

#### Folder character 


```go
// File character.go
package main

import (
    "fmt"
)

// Character is the base character type.
type Character struct {
    Name     string
    HitPoints int
}

// Attack represents the attack action for a character.
type Attacker interface {
    Attack() int
}
```

#### Folder wizard
```go
// file wizard.go
// Wizard is a type of character that can cast spells.
type Wizard struct {
    Character
    ManaPoints int
}

```

#### Folder warrior
```go
// file warrior.go
// Warrior is a type of character that excels in melee combat.
type Warrior struct {
    Character
    Strength int
}

```

#### Folder DandD
```go
// file main.go
// NewCharacter creates a new base character.
func NewCharacter(name string, hp int) Character {
    return Character{
        Name:     name,
        HitPoints: hp,
    }
}

// NewWizard creates a new wizard character.
func NewWizard(name string, hp int, mp int) Wizard {
    return Wizard{
        Character:  NewCharacter(name, hp),
        ManaPoints: mp,
    }
}

// NewWarrior creates a new warrior character.
func NewWarrior(name string, hp int, strength int) Warrior {
    return Warrior{
        Character: NewCharacter(name, hp),
        Strength:  strength,
    }
}

// Attack implements the Attack method for a Wizard.
func (w Wizard) Attack() int {
    return w.ManaPoints * 2
}

// Attack implements the Attack method for a Warrior.
func (w Warrior) Attack() int {
    return w.Strength * 3
}

func main() {
    // Create a base character
    character := NewCharacter("Alice", 100)

    // Create a wizard character
    wizard := NewWizard("Gandalf", 80, 150)

    // Create a warrior character
    warrior := NewWarrior("Conan", 120, 30)

    // Polymorphism: Characters can attack regardless of their specific type
    characters := []Attacker{character, wizard, warrior}

    for _, c := range characters {
        fmt.Printf("%s attacks for %d damage\n", c.(Character).Name, c.Attack())
    }
}
```

In this code:

1. We define a `Character` struct to represent the base character. Both `Wizard` and `Warrior` embed the `Character` struct, which allows them to inherit its fields.

2. We define an `Attacker` interface, which has an `Attack` method. Both `Wizard` and `Warrior` implement this interface, which demonstrates Polymorphism.

3. We use constructor functions like `NewCharacter`, `NewWizard`, and `NewWarrior` to create instances of characters while encapsulating the initialization logic.

4. The `Attack` method is implemented differently for `Wizard` and `Warrior`, showing how they can have their unique behaviors.

5. In the `main` function, we demonstrate Polymorphism by creating a slice of `Attacker` interfaces and having characters of different types (Wizard, Warrior, and Character) attack, encapsulating their specific behaviors.

This code illustrates how Go can achieve concepts of `OOP` like `Polymorphism`, `Encapsulation`, and `Abstraction` through the use of `struct embedding` and `interfaces`.

[Previous: 07-first_class_func.md](07-first_class_func.md) | [Next: 09-interfaces_and_inheritance.md](09-interfaces_and_inheritance.md)
