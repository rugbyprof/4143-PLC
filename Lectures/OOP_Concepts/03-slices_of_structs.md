## Slices Of Structs

In the Go programming language, a "slice of structs" refers to a slice that contains elements, and each element in the slice is a struct. Let's break down what this means:

### Structs: 

In Go, a struct is a composite data type that groups together variables (fields) under a single type. Each struct can have multiple fields of different data types. For example:

   ```go
   type Person struct {
       FirstName string
       LastName  string
       Age       int
   }
   ```

   Here, `Person` is a struct with three fields: `FirstName`, `LastName`, and `Age`.

### Slice: 

A slice in Go is a dynamically sized, flexible view of an array. Slices allow you to work with a variable-length sequence of elements. For example:

   ```go
   names := []string{"Alice", "Bob", "Charlie"}
   ```

   Here, `names` is a slice of strings, containing three elements.

Now, when we talk about a "slice of structs" in the Go world, it means we have a slice where each element is an instance of a struct. For instance:

```go
people := []Person{
    {FirstName: "Alice", LastName: "Johnson", Age: 28},
    {FirstName: "Bob", LastName: "Smith", Age: 35},
    {FirstName: "Charlie", LastName: "Brown", Age: 42},
}
```

In this example, `people` is a slice of `Person` structs. Each element in the `people` slice is a `Person` struct with its own `FirstName`, `LastName`, and `Age` fields.

Using slices of structs is common in `Go` for organizing and working with collections of structured data. It allows you to manage and manipulate groups of related data entities efficiently.



[Previous: 02-new_keyword.md](02-new_keyword.md) | [Next: 04-maps.md](04-maps.md)
