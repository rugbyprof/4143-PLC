### Explanation of packages and more ways to organize them in a project


1. **Package Declaration and Visibility**:

   ```go
   // packageexample.go
   package mypackage

   import "fmt"

   // Exported variable
   var ExportedVar = 42

   // Unexported variable (only accessible within the package)
   var unexportedVar = "unexported"

   // Exported function
   func ExportedFunction() {
       fmt.Println("Exported function")
   }

   // Unexported function (only accessible within the package)
   func unexportedFunction() {
       fmt.Println("Unexported function")
   }
   ```

   In this example, `ExportedVar` and `ExportedFunction` are exported and can be used outside the package, while `unexportedVar` and `unexportedFunction` are unexported and only accessible within the package.

2. **Package Imports**:

   ```go
   // main.go
   package main

   import (
       "fmt"
       "mypackage"
   )

   func main() {
       fmt.Println(mypackage.ExportedVar) // Accessing an exported variable
       mypackage.ExportedFunction()       // Calling an exported function
   }
   ```

   In this example, the `mypackage` package is imported and its exported identifiers are used in the `main` package.

3. **Main Package and Entry Point**:

   ```go
   // main.go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, world!")
   }
   ```

   The `main` package serves as the entry point for the Go program, and it contains the `main` function, which is executed when the program runs.

4. **Creating Your Own Packages**:

   Let's say you have a directory structure like this:

   ```
   myproject/
   ├── main.go
   ├── mypackage/
   │   ├── packageexample.go
   │   └── anotherpackage.go
   ```

   You can organize related functionality into separate files within the `mypackage` package directory, and each file should have the same package declaration, e.g., `package mypackage`.

5. **Package Initialization**:

   ```go
   // mypackage/packageexample.go
   package mypackage

   import "fmt"

   // Initialization code
   func init() {
       fmt.Println("Initialization of mypackage")
   }
   ```

   The `init` function in a package is automatically called when the package is imported. This is often used for package-level setup.

6. **Package Documentation**:

   ```go
   // mypackage/packageexample.go
   package mypackage

   // ExportedFunction is an example of an exported function.
   func ExportedFunction() {
       // ...
   }
   ```

   You can add comments like this to document your exported identifiers. To generate documentation, you can use the `godoc` tool.

7. **Package Versioning**:

   Go modules provide a way to manage package dependencies and versions. You can create a `go.mod` file to specify your module's dependencies, including their versions. Here's a simplified example:

   ```go
   module mymodule

   require (
       github.com/somepackage v1.2.3
   )
   ```

   This example specifies a dependency on `github.com/somepackage` at version `v1.2.3`.

### More ways of organizing your code with a packages code example. 
Let's assume you have the following directory structure shown below, 

```txt
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

This is different from the last slide where you were shown the Wizard, Warrior and Character structs all in the same package. Now they are all different packages. With this comes implications though... Run through the following code. 

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

[Previous: 09-composition_and_packages.md](./09-composition_and_packages.md) | [Next: 11-comp_over_inherit.md](./11-comp_over_inherit.md)





