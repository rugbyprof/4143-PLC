### Slice: 

A slice in Go is a dynamically sized, flexible view of an array. Slices allow you to work with a variable-length sequence of elements. For example:

   ```go
   names := []string{"Alice", "Bob", "Charlie"}
   ```

   Here, `names` is a slice of strings, containing three elements.

They are like arrays in C++. The type is explicitly declared, unlike Python's implicit allocation. 

## Slices Of Structs

In the Go programming language, a "slice of structs" refers to a slice that contains elements, and each element in the slice is a struct. Let's break down what this means:

   ```go
   type Person struct {
       FirstName string
       LastName  string
       Age       int
   }
   ```

   Here, `Person` is a struct with three fields: `FirstName`, `LastName`, and `Age`.

Now, when we talk about a "slice of structs" in the Go world, it means we have a slice where each element is an instance of a struct. For example:

```go
people := []Person{
    {FirstName: "Alice", LastName: "Johnson", Age: 28},
    {FirstName: "Bob", LastName: "Smith", Age: 35},
    {FirstName: "Charlie", LastName: "Brown", Age: 42},
}
```

Slices of structs are almost like JSON objects or python dictionaries. It's essentially a list of dictionaries in python. A C++ reference: vectors of objects or structs. 


Accessing an element

```go
firstPerson := people[0]
fmt.Println("First person:", firstPerson.Name)
```

Appending a new element

```go
newPerson := Person{"David", 40}
people = append(people, newPerson)
```

Modifying an element

```go
people[0].Age = 26
```

'slicing' the slice
```go
subset := people[1:3] // Creates a new slice containing elements at index 1 and 2
```

Slices have a length and capacity. The length is the number of elements currently in the slice, while the capacity is the maximum number of elements the slice can hold without reallocating memory.
You can use the built-in len and cap functions to get the length and capacity of a slice.
```go
length := len(people)
capacity := cap(people)
```



In this example, `people` is a slice of `Person` structs. Each element in the `people` slice is a `Person` struct with its own `FirstName`, `LastName`, and `Age` fields. We then show how to add "David" later on, and then show much greater detail behind the slice of structs. 

Using slices of structs is common in `Go` for organizing and working with collections of structured data. It allows you to manage and manipulate groups of related data entities efficiently.

For those inspired by this document here is the `Go Official Docs` link [Slices - Go docs](https://go.dev/blog/slices-intro)


[Previous: 02-new_keyword.md](02-new_keyword.md) | [Next: 04-maps.md](04-maps.md)
