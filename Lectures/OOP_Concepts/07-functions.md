## Functions

Let's explore the idea of `first-class functions` in `Go` within the context of our scary Dungeons & Dragons (D&D) theme. This overview will go through the following topics: 
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

user-defined function types, often referred to as "function types" or "function signatures," allow you to declare custom types that represent functions with specific parameter types and return types. These user-defined function types can be useful for creating flexible and reusable code, especially in scenarios where functions are passed as arguments or returned from other functions. Here's some more information about user-defined function types in Go:

1. **Defining a User-Defined Function Type**:

   You can define a user-defined function type by specifying the parameter types and return type using the `type` keyword. For example:

   ```go
   type MyFunc func(int, int) int
   ```

   In this example, `MyFunc` is a user-defined function type that represents a function taking two `int` parameters and returning an `int`.

2. **Using Function Types as Variables**:

   You can use function types to declare variables that can hold functions with matching signatures. For example:

   ```go
   var add MyFunc
   ```

   This `add` variable can now hold any function that matches the `MyFunc` signature, which means it must accept two `int` parameters and return an `int`.

3. **Assigning Functions to Variables**:

   You can assign functions to variables of user-defined function types. For example:

   ```go
   add = func(a, b int) int {
       return a + b
   }
   ```

   Here, we're assigning an anonymous function that adds two integers to the `add` variable.

4. **Using Function Variables**:

   Once a function is assigned to a variable with a specific function type, you can call that function using the variable as if it were a regular function. For example:

   ```go
   result := add(3, 4)
   ```

   In this case, `add` is used as if it were a function taking two `int` parameters and returning an `int`.

5. **Passing Function Types as Arguments**:

   You can pass variables of user-defined function types as arguments to functions. This enables you to define functions that accept functions as parameters, allowing for more flexible behavior. For example:

   ```go
   func calculate(a, b int, operation MyFunc) int {
       return operation(a, b)
   }

   sum := calculate(5, 3, add) // Using the 'add' function
   ```

   Here, `calculate` is a function that takes two integers and a function of type `MyFunc` as arguments.

6. **Returning Function Types from Functions**:

   Functions can also return user-defined function types. This is often used in scenarios where a function needs to create and return different functions based on some conditions or configurations.

   ```go
   func getOperation(opType string) MyFunc {
       if opType == "add" {
           return func(a, b int) int {
               return a + b
           }
       } else {
           return func(a, b int) int {
               return a - b
           }
       }
   }
   ```

   In this example, `getOperation` returns a function of type `MyFunc` based on the `opType` parameter.

User-defined function types in Go provide a way to define functions as first-class citizens, allowing for more flexible and powerful code structures, such as callbacks, higher-order functions, and function factories. They are particularly useful in scenarios where you want to create generic or configurable code that works with functions of various types.


### Higher-Order Functions:

In Go, you can define higher-order functions by creating functions that accept other functions as arguments or return functions as results.

1. **Accepting Functions as Arguments**:

   In Go, you can create higher-order functions that accept functions as arguments, often referred to as callbacks. These callbacks are functions that you pass to the higher-order function to customize its behavior.

   ```go
   package main

   import "fmt"

   // Higher-order function that accepts a callback function
   func applyCallback(data []int, callback func(int) int) []int {
       result := make([]int, len(data))
       for i, value := range data {
           result[i] = callback(value)
       }
       return result
   }

   // Callback function to double a value
   func double(x int) int {
       return x * 2
   }

   func main() {
       numbers := []int{1, 2, 3, 4, 5}

       // Use the higher-order function with the 'double' callback
       doubledNumbers := applyCallback(numbers, double)

       fmt.Println(doubledNumbers) // Output: [2 4 6 8 10]
   }
   ```

   In this example, `applyCallback` is a higher-order function that accepts a callback function `callback` as an argument and applies it to each element of the `data` slice.

2. **Returning Functions as Results**:

   You can also create higher-order functions in Go that return functions as results. This is useful for creating functions with specific behaviors based on the arguments passed to the higher-order function.

   ```go
   package main

   import "fmt"

   // Higher-order function that returns a function
   func createMultiplier(factor int) func(int) int {
       return func(x int) int {
           return x * factor
       }
   }

   func main() {
       double := createMultiplier(2)
       triple := createMultiplier(3)

       result1 := double(5) // Result: 10
       result2 := triple(5) // Result: 15

       fmt.Println(result1, result2)
   }
   ```

   Here, `createMultiplier` is a higher-order function that returns a function. When you call `createMultiplier(2)`, it returns a function that multiplies its argument by `2`. Similarly, calling `createMultiplier(3)` returns a function that multiplies its argument by `3`.

Let's now look at an example that incorporates the D&D theme.

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

In this example, `CreateSpell` is a higher-order function that returns a `SpellCaster` function. We can then pass these `SpellCaster` functions to `CastSpell` which accepts a `SpellCaster` type as an argument. 

Functions in Go can return other functions, allowing for dynamic behavior. In D&D terms, it's like finding a magical item that grants you new abilities

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

- Creating customizable behavior: You can pass functions as arguments to customize the behavior of higher-order functions.
- Implementing strategies: You can use first-class functions to implement strategies for different situations.
- Handling callbacks: Functions can be passed as callbacks to handle events or asynchronous operations.
- Achieving modularity: Functions can be swapped or extended easily to achieve modularity and code reusability.

#### Full working example using all of the above concepts
```go
package main

import "fmt"

// User defined function types
type Spell func() string
type SpellCaster func() string

// CLOSURE - counter for number of spells cast
func SpellsCast(num int) func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

// higher order function, returns a function
// returns a function with the type SpellCaster
func CreateSpell(spellName string) Spell {
	return func() string {
		return "Casting " + spellName + " spell!"
	}
}

// higher order function, returns a function
// returns a function with the type SpellCaster
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

// uses the SpellCaster type as params, can be used to display
// both CreateSpell and CreateSpellcaster
func CastSpell(caster SpellCaster, spell Spell, spell_count func() int) {
  caster()
  fmt.Println(caster())
  spell_count() // everytime a spell is cast, we increment our spell count by 1
}

func main() {
  num_spells := SpellsCast(0)

  wizard := CreateSpellcaster("Wizard")
  sorcerer := CreateSpellcaster("Sorcerer")

  fireball := CreateSpell("Fireball")
  lightningBolt := CreateSpell("Lightning Bolt")

	for i := 1; i <= 10; i++ {
		CastSpell(wizard, fireball, num_spells)      // Cast the Fireball spell
		CastSpell(wizard, lightningBolt, num_spells) // Cast the Lightning Bolt spell

		CastSpell(sorcerer, fireball, num_spells)      // Cast the Fireball spell
		CastSpell(sorcerer, lightningBolt, num_spells) // Cast the Lightning Bolt spell
	}

  // minus one from the called value when displaying. since every call increments count by 1
  fmt.Printf("The total number of spells cast was %d\n", num_spells()-1)
}


```

Link to Golang docs: [Functions - Go Docs](https://golangdocs.com/functions-in-golang)

[Previous: 06-methods_vs_functions.md](06-methods_vs_functions.md) | [Next: 08-interfaces.md](08-interfaces.md)
