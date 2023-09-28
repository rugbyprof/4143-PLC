## Composition and Packages

In Go there are not traditional classes along with inheritance as in traditional object-oriented languages like Java or C++. But you can achieve similar concepts using `struct`s embedded with other structs (composition) and also the `interface` type that enforce certain methods.

Encouragement of Clear Intent: Composition makes it explicit which components a type depends on. This clarity of intent can lead to more self-documenting code, making it easier for other developers to understand and work with the codebase.

Simplicity and Readability: Go strives to keep its language and codebase simple and easy to read. Inheritance hierarchies can often become complex and difficult to understand, leading to code that is hard to maintain and debug. Composition provides a simpler and more straightforward way to achieve code reuse.

 Below is an example of a base character extended to be a `Wizard` and a `Warrior` in `Go`, while emphasizing the principles of Polymorphism, Encapsulation, and Abstraction.

 Let's assume the following directory structure:

 ```
DandD/
├── main.go
├─────  module
│   ├──── character.go
│   ├──── wizard.go
│   ├──── warrior.go
└── ...
```

### Example in folder DandD 

#### directory module/character 


```go
// File character.go
package module

// Character is the base character type.
type Character struct {
    Name      string
    HitPoints int
}

// NewCharacter creates a new base character.
func NewCharacter(name string, hp int) Character {
	return Character{
		Name:      name,
		HitPoints: hp,
	}
}

// Attack represents the attack action for a character.
type Attacker interface {
    Attack() int
    GetName() string
}
```

#### directory module/wizard
```go
// file wizard.go
// Wizard is a type of character that can cast spells.
package module

type Wizard struct {
    Character
    ManaPoints int
}

// NewWizard creates a new wizard character.
func NewWizard(name string, hp int, mp int) Wizard {
	return Wizard{
		Character:  NewCharacter(name, hp),
		ManaPoints: mp,
	}
}

func (w Wizard) GetName() string {
  return w.Character.Name
}

// Attack implements the Attack method for a Wizard.
func (w Wizard) Attack() int {
	return w.ManaPoints * 2
}

```

#### directory module/warrior
```go
// file warrior.go
// Warrior is a type of character that excels in melee combat.
package module

type Warrior struct {
    Character
    Strength int
}

// Attack implements the Attack method for a Warrior.
func (w Warrior) Attack() int {
	return w.Strength * 3
}

func (w Warrior) GetName() string {
  return w.Character.Name
}

// NewWarrior creates a new warrior character.
func NewWarrior(name string, hp int, strength int) Warrior {
	return Warrior{
		Character: NewCharacter(name, hp),
		Strength:  strength,
	}
}

```

#### main.py
```go
// file main.go
package main

import (
  "fmt"
	"main/module"
)

func main() {

	// Create a wizard character
	wizard := module.NewWizard("Gandalf", 80, 150)

	// Create a warrior character
	warrior := module.NewWarrior("Conan", 120, 30)

	// Polymorphism: Characters can attack regardless of their specific type
	characters := []module.Attacker{wizard, warrior}

	for _, c := range characters {
		fmt.Printf("%s attacks for %d damage\n", c.GetName(), c.Attack())
	}
}

```

In this code:

1. Composition in Go: The Wizard and Warrior types use composition by embedding a Character within themselves. This means they have a Character instance as a field and can access its fields and methods through that instance.

2. Inheritance in Go: Go does not support classical inheritance like some other object-oriented languages. Instead, it promotes composition and the use of interfaces for achieving similar functionality.

3. We use constructor functions like `NewCharacter`, `NewWizard`, and `NewWarrior` to create instances of characters while encapsulating the initialization logic.

4. The `Attack` method is implemented differently for `Wizard` and `Warrior`, showing how they can have their unique behaviors. but they both implement Attack() and GetName() as defined by the Attacker interface. 

5. In the `main` function, we demonstrate Polymorphism by creating a slice of `Attacker` interfaces and having characters of different types (Wizard, Warrior, and Character) attack, encapsulating their specific behaviors.

This code illustrates how Go can achieve concepts of `OOP` like `Polymorphism`, `Encapsulation`, and `Abstraction` through the use of `struct embedding`, 'composition' and `interfaces`.

[Previous: 08-interfaces.md](08-interfaces.md) | [Next: 10-packages_in_depth.md](10-packages_in_depth.md)
