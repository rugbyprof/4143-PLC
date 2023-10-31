## Setters Getters

In Go, there are no strict naming conventions for setters and getters like you might find in some other languages. Go follows a more straightforward approach where the naming of fields, methods, and functions should be clear and idiomatic.

Here are some common practices for naming setters and getters in Go:

1. **Use Direct Access for Simple Fields**:
   - In Go, you often directly access fields (variables) within a struct if they are publicly visible (start with an uppercase letter). This is because getters and setters for simple fields are not required.

2. **Use Capitalized Field Names**:
   - If a field is intended to be accessed from outside the package, give it a name that starts with an uppercase letter (e.g., `FieldName` instead of `fieldName`).

3. **Use Methods for Complex Logic**:
   - If a field requires more complex logic when getting or setting its value, use methods instead of direct access. The method names should be clear and follow Go's naming conventions.

Here's an example of how you might define getters and setters for a struct in Go:

```go
package mypackage

type MyStruct struct {
    // Public fields
    Name string
    Age  int
}

// Getter for Name
func (m *MyStruct) GetName() string {
    return m.Name
}

// Setter for Name
func (m *MyStruct) SetName(newName string) {
    m.Name = newName
}

// Getter for Age
func (m *MyStruct) GetAge() int {
    return m.Age
}

// Setter for Age
func (m *MyStruct) SetAge(newAge int) {
    m.Age = newAge
}
```

In the example above, `GetName`, `SetName`, `GetAge`, and `SetAge` are used as method names for getting and setting the values of the fields `Name` and `Age`. Note that these methods are only necessary if you need to encapsulate more logic than simple field access.

However, for simple fields, it's common in Go to directly access them using the field name (e.g., `myStruct.Name`) if they are publicly visible, rather than using explicit getter and setter methods. This aligns with Go's philosophy of simplicity and readability.