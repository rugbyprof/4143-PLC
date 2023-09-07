---
marp: true
theme: gradient
paginate: true
---

# Variable Declaration in Go
## Implicit and Explicit Creation

---

# Variable Declaration in Go
- Go is statically typed, but it offers flexibility in variable creation.
- You can declare variables explicitly or let Go infer the type.

```go
var age int          // Explicit declaration
name := "Alice"      // Implicit declaration
```

- `var` keyword for explicit, `:=` for implicit.

---

# Implicit Variable Declaration in Go
- Go uses type inference to determine the variable's type.

```go
name := "Alice"       // string
age := 30             // int
average := 87.5       // float64
isStudent := true     // bool
```

- Type is inferred from the value.

---

# Explicit Variable Declaration in Go
- Declare variables with a specified type.

```go
var name string = "Alice"
var age int = 30
var average float64 = 87.5
var isStudent bool = true
```

- Type is explicitly mentioned.

---

# Equivalent in Python
- Python uses dynamic typing, so variable types are determined at runtime.

```python
name = "Alice"      # str
age = 30            # int
average = 87.5      # float
is_student = True   # bool
```

- No type declarations needed.

---

# Equivalent in C++
- C++ is statically typed, similar to Go.

```cpp
std::string name = "Alice";
int age = 30;
double average = 87.5;
bool isStudent = true;
```

- Type is explicitly declared.

---

# Short Variable Declaration in Go
- Go allows short variable declaration inside functions.

```go
func main() {
    count := 42
    message := "Hello"
}
```

- `:=` is used within functions.

---

# Global Variable Declaration in Go
- Global variables are declared using `var`.

```go
var globalVar int = 100
```

- Global scope, explicit type.

---

# Constants in Go
- Go allows constant declarations using `const`.

```go
const pi = 3.14159265
const companyName string = "Acme Inc."
```

- Constants are immutable.

---

# Comparison: Implicit Declaration
- Go uses `:=` for implicit declaration.
- Python is dynamically typed; no explicit declarations.
- C++ requires explicit type declarations.

---

# Comparison: Explicit Declaration
- Go uses `var variableName typeName` for explicit declarations.
- Python doesn't need explicit type declarations.
- C++ requires explicit type declarations.

---

# Comparison: Constants
- Go uses `const` for constants.
- Python can use variables as constants.
- C++ uses `const` as well.

---

# Summary
- Go offers both implicit and explicit variable declarations.
- Python uses dynamic typing.
- C++ is statically typed like Go.
- Constants are declared using `const` in Go.

