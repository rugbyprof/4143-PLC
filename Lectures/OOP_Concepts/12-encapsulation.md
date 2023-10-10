# Encapsulation 

Encapsulation in Go is supported through even more naming conventions within Structs. For examples, if you name your struct field in the uppercase it will be a public field, but if you name it in the lowercase it will be private. 

### For example
```go
// public struct
type Book struct{

        // public field
        Name string
        // private field, only
        // available in gfg package
        author string
}
```

The above example showcases the usage of private and public properties within a Golang Struct. This allows the practice of encapsulation and information hiding found in C++ using the public and private keywords. 

Let's showcase the usage of public methods. The ToUpper() function is defined to operate on strings and changes all characters to be uppercase. This is a built in function like python's .upper() function

```go
// Go program to illustrate
// the concept of encapsulation
// using exported function
package main
  
import (
    "fmt"
    "strings"
)
  
// Main function
func main() {
  
    // Creating a slice of strings
    slc := []string{"Garrett", "Is", "A", "Boss", "Programmer"}
  
    // Convert the case of the
    // elements of the given slice
    // Using ToUpper() function
    for x := 0; x < len(slc); x++ {
  
        // Exported Method
        res := strings.ToUpper(slc[x])
  
        // Exported Method
        fmt.Println(res)
    }
}
```
output: 
GARRETT 
IS 
A 
BOSS 
PROGRAMMER

### But what if we did lowercase .toUpper()?

It would error because anything beginning with a lowercase character is unexported.

```
res := strings.toUpper(slc[x])

Output:
./prog.go:22:9: cannot refer to unexported name strings.toUpper
./prog.go:22:9: undefined: strings.toUpper
```

## Why is this knowledge useful?

- Because when we work with packages and modules in Go, we often will declare structs to use throughout our project and may want to expose some functionality within that struct to be useable throughout the project. 
- If the method is simply used for internal setup of the data type then you may not want to export it so you can hide the implementation details. 

```
EXPORTED:         func Attack(w wizard) int
NOT EXPORTED:     func attack(w Wizard) int
```

This file is intended to clarify some more object oriented concepts that Go implements in it's own fashion. 

[Previous: 11-comp_over_inherit.md](./11-comp_over_inherit.md) | [Next: README.md](./README.md)



