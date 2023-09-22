## Interfaces and Inheritance

#### Q:
How would we have an attack method for each character type if we cannot overload methods (same name, different param lists)? 
#### A:
You would create different packages for each character type, and import the base character package.

#### Q:
And if we did have separate packages for each character type, how would an `Attack` interface be implemented and/or enforced for all character types

#### A:
Having a common method name like `Attack` across different character types can simplify your code and provide consistency. You can achieve this by defining an interface within the generic character types that other characters must implement. There is an example provided below.


## Example

Let's assume you have the following directory structure:

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

### Define the Character Interface:

Create an `attacker.go` file in the `character` package to define the `Attacker` interface:

```go
// character/attacker.go

package character

// Attacker is an interface for characters that can attack.
type Attacker interface {
    Attack() int
}
```

### Implement the Interface in Character Types:

In the `wizard` and `warrior` packages, you can implement the `Attacker` interface for the respective character types.

#### Wizard

```go
// wizard/wizard.go

package wizard

import (
    "DandD/character"
)

// Wizard represents a wizard character.
type Wizard struct {
    // Fields specific to wizards
}

// Implement the Attack method for wizards.
func (w Wizard) Attack() int {
    // Wizard-specific attack logic
}
```

#### Warrior

```go
// warrior/warrior.go

package warrior

import (
    "DandD/character"
)

// Warrior represents a warrior character.
type Warrior struct {
    // Fields specific to warriors
}

// Implement the Attack method for warriors.
func (w Warrior) Attack() int {
    // Warrior-specific attack logic
}
```

### Use the Interface in Your Main Code:

In your `main.go` or other main code, you can work with character instances through the `Attacker` interface, allowing you to call the `Attack` method regardless of the character's actual type.

```go
// main.go

package main

import (
    "DandD/character"
    "DandD/wizard"
    "DandD/warrior"
)

func main() {
    // Create character instances
    wizardCharacter := wizard.Wizard{}
    warriorCharacter := warrior.Warrior{}

    // Put them in a slice to demonstrate polymorphism
    characters := []character.Attacker{wizardCharacter, warriorCharacter}

    for _, c := range characters {
        // Call the Attack method on characters of different types
        damage := c.Attack()
        // Handle the damage or other actions
    }
}
```

- By defining an interface (`Attacker`) in the `character` package and implementing it in the `wizard` and `warrior` packages, you can enforce a common behavior (in the `Attack` method) across different character types. 
- This approach allows you to work with characters uniformly in your main code while still leveraging the unique implementations of the `Attack` method for each character type.



[Previous: 08-inheritance.md](08-inheritance.md) | [Next: 10-interfaces.md](10-interfaces.md)
