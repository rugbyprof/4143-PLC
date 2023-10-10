# Struct vs Interface

In Go, `structs` and `interfaces` are two fundamental constructs, each serving a distinct purpose. Let's explore the differences between structs and interfaces and relate them to a D&D (Dungeons & Dragons) example.

## Struct

1. **Purpose:** 
   - A `struct` in `Go` is used to define a `composite data type` that groups together variables (fields) under a single name. 
   - It represents a concrete data structure with specific fields.

2. **Data and Behavior:** 
   - Structs primarily store data (fields) and can include methods, which are functions associated with the struct type. 
   - These methods operate on the data stored in the struct.

3. **Example (D&D-related):** 
   - In the context of D&D, you can use structs to define character types with specific attributes, such as a `Character` struct with fields like `Name`, `HitPoints`, `ManaPoints`, and `Strength`. 
   - You can also have methods specific to characters, like `Attack()` or `CastSpell()`.

4. **Concrete Types:** 
   - Structs are `concrete types`, meaning you can `create instances (objects)` of a struct with specific data values.

## Interface

1. **Purpose:** 
   - An `interface` in Go is used to define a `set of method signatures `that a type must implement. 
   - It represents behavior rather than concrete data.

2. **Data and Behavior:** 
   - Interfaces don't contain data (fields); instead, they specify a contract for methods that types must adhere to. 
   - A type that implements all the methods of an interface is said to satisfy that interface.

3. **Example (D&D-related):** 
   - In D&D, you can define an `Attacker` interface with a method `Attack() int`. 
   - Any character type, such as `Wizard` or `Warrior`, can implement this interface by providing their own `Attack` method. 
   - This allows characters of different types to perform attacks, even if their implementations differ.

4. **Abstraction and Polymorphism:** 
   - `Interfaces` enable `abstraction` by defining behavior without specifying the implementation details. 
   - They also enable `polymorphism`, allowing different types to be used interchangeably if they satisfy the same interface.

Here's some code that relates to a D&D character:

```go
package main

import "fmt"

// This interface requires a charactor to implement the Attack method,
// but also allows differing character types to tailor attacks to their specific class.
type Attacker interface {
	Attack() int
	GetName() string
}

// Struct representing a generic character (Like a base class)
type Character struct {
	Name string
}

// Struct representing a wizard character
// Shows composisition by using an instance of a character type
type Wizard struct {
	c          Character
	ManaPoints int
}

// Method to attack for a wizard character
// Shows polymorphism since this attack method is tailored for a wizard.
func (w Wizard) Attack() int {
	return w.ManaPoints * 2
}

func (w Wizard) GetName() string {
	return w.c.Name
}

// Struct representing a warrior character
// Shows composisition by using an instance of a character type
type Warrior struct {
	c        Character
	Strength int
}

// Method to attack for a warrior character
// Shows polymorphism since this attack method is tailored for a warrior.
func (w Warrior) Attack() int {
	return w.Strength * 3
}

func (w Warrior) GetName() string {
	return w.c.Name
}

func _main() {
	// Create instances of characters
	//genericCharacter := &Character{Name: "Generic Character"} // can't use because it does not implement 
	wizardCharacter := &Wizard{c: Character{Name: "Wizard Character"}, ManaPoints: 20}
	warriorCharacter := &Warrior{c: Character{Name: "Warrior Character"}, Strength: 30}

	// Create a slice of attackers
	attackers := []Attacker{wizardCharacter, warriorCharacter}

	// Iterate through the attackers and display their names and attacks
	for _, attacker := range attackers {
		fmt.Printf("%s attacks for %d damage.\n", attacker.GetName(), attacker.Attack())
	}
}

```

- In this example, we define an `Attacker` interface, and both `Wizard` and `Warrior` implement it by providing their own `Attack` methods. 
- This allows us to abstract the concept of attacking and apply polymorphism, where we can call the `Attack` method on characters of different types.
- We then use a slice of structs that implement all of the requirements of the attacker interface. 

#### Summary 
- `Structs` are used for defining data structures with fields and methods
- `Interfaces` are used for defining behavior through method signatures that structs must adhere to. 
- `Slices of Structs` that implement the rules of an interface. This allows you to mix and match structs in a slice if the structs in that slice all fulfill the obligations of a common interface. As shown above.
- `Output` we then output using the methods the Attacker interface requires. 
- `Explicit` we can begin to see why Go is a very explicit language in the way that we must write our methods to perform any intended action on the struct. You must explicitly code that functionality. GetName() and Attack() are good at visualizing this.  
- In the context of D&D, `structs` represent character types with attributes, while `interfaces` represent shared actions like attacking.

[Previous: 04-maps.md](./04-maps.md) | [Next: 06-methods_vs_functions.md](./06-methods_vs_functions.md)
