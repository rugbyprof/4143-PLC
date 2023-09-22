## First Class Functions

Let's explore `first-class functions` in `Go` within the context of our scary Dungeons & Dragons (D&D) theme. This overview will: 
- first-class functions
- anonymous functions
- user-defined function types
- higher-order functions
- returning functions from other functions
- closures
- and sone other practical use cases

### What are First-Class Functions?

In Go, functions are `first-class citizens`, which means they can be treated like any other value, such as integers or strings. First-class functions can be assigned to variables, passed as arguments to other functions, and returned from other functions.

### Anonymous Functions:

`Anonymous functions`, also known as function literals or lambda functions, are functions without a defined name. In D&D terms, think of them as scrolls with magical spells that can be used without naming them.

```go
spell := func() {
    fmt.Println("Casting a fireball spell!")
}

spell() // Cast the fireball spell
```

In this example, we define an anonymous function and assign it to the variable `spell`. We can then call it like a regular function.

### User-Defined Function Types:

In Go, you can declare your own function types. Imagine creating custom classes for D&D character abilities:

```go
type Spell func() string

func Fireball() string {
    return "Fireball spell cast!"
}

func LightningBolt() string {
    return "Lightning bolt spell cast!"
}

func main() {
    var spell Spell
    spell = Fireball
    fmt.Println(spell()) // Cast the Fireball spell

    spell = LightningBolt
    fmt.Println(spell()) // Cast the Lightning Bolt spell
}
```

Here, we define a `Spell` type, which is a function that takes no arguments and returns a string. We then create functions (`Fireball` and `LightningBolt`) that conform to this type.

### Higher-Order Functions:

`Higher-order functions` are functions that either take other functions as arguments or return functions as results. In D&D, think of them as characters who can cast spells and characters who can create spells:

```go
type SpellCaster func() string

func CreateSpell(spellName string) SpellCaster {
    return func() string {
        return "Casting " + spellName + " spell!"
    }
}

func CastSpell(caster SpellCaster) {
    fmt.Println(caster())
}

func main() {
    fireball := CreateSpell("Fireball")
    lightningBolt := CreateSpell("Lightning Bolt")

    CastSpell(fireball)      // Cast the Fireball spell
    CastSpell(lightningBolt) // Cast the Lightning Bolt spell
}
```

In this example, `CreateSpell` is a higher-order function that returns a `SpellCaster` function. We can then pass these `SpellCaster` functions to `CastSpell`.

### Returning Functions from Other Functions:

Functions in Go can return other functions, allowing for dynamic behavior. In D&D terms, it's like finding a magical item that grants you new abilities:

```go
func CreateSpellcaster(class string) SpellCaster {
    if class == "Wizard" {
        return func() string {
            return "Casting a wizard spell!"
        }
    } else if class == "Sorcerer" {
        return func() string {
            return "Casting a sorcerer spell!"
        }
    }
    return nil
}

func main() {
    wizardCaster := CreateSpellcaster("Wizard")
    sorcererCaster := CreateSpellcaster("Sorcerer")

    CastSpell(wizardCaster)   // Cast a wizard spell
    CastSpell(sorcererCaster) // Cast a sorcerer spell
}
```

Here, `CreateSpellcaster` is a function that returns different `SpellCaster` functions based on the character's class.

### Closures:

A closure is a function value that captures variables from its surrounding lexical scope. In D&D, think of it as a magical aura that keeps certain variables alive:

```go
func MakeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := MakeCounter()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

Here, the inner function captures the `count` variable from the outer function's scope, creating a closure. The `count` variable persists between calls to the inner function.

### Practical Use of First-Class Functions:

First-class functions are used in Go for various purposes, including:

- Creating customizable behavior: You can pass functions as arguments to customize the behavior of higher-order functions.
- Implementing strategies: You can use first-class functions to implement strategies for different situations.
- Handling callbacks: Functions can be passed as callbacks to handle events or asynchronous operations.
- Achieving modularity: Functions can be swapped or extended easily to achieve modularity and code reusability.


### Summary

First-class functions in Go empower you to create flexible, dynamic, and modular code, just like characters in a D&D world who can learn and cast different spells or use various abilities as needed.

[Previous: 06-methods_vs_functions.md](06-methods_vs_functions.md) | [Next: 08-inheritance.md](08-inheritance.md)
