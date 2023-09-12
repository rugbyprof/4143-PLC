---
marp: true
theme: gradient
paginate: true
---

# Functions in GoLang, Python, and C++
## Building Blocks of Reusable Code

---

# What are Functions?
- Functions are blocks of code that perform a specific task.
- They are reusable and can be called from various parts of a program.
- Functions help in modularizing code for better organization and readability.

---

# Function Declaration in Go

- Functions are declared using the `func` keyword.
- `sayHello` is the function name.

#### Python
```python
def say_hello():
    print("Hello, World!")
```

#### C++
```cpp
void sayHello() {
    cout << "Hello, World!" << endl;
}
```

```go
func sayHello() {
    fmt.Println("Hello, World!")
}
```

---

# Function Invocation

#### Python
```python
say_hello() # Call the say_hello function
```
#### C++
```cpp
sayHello(); // Call the sayHello function
```
#### Go
```go
sayHello() // Call the sayHello function
```
---

# Function Parameters

#### Python
```python
def greet(name):
    print(f"Hello, {name}!")
```
#### C++
```cpp
#include <iostream>
using namespace std;

void greet(string name) {
    cout << "Hello, " << name << "!" << endl;
}
```
#### Go
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```
---



# Function Arguments in Python
```python
greet("Alice") # Pass "Alice" as an argument
```

- Python function arguments are straightforward.

---

# Function Arguments in C++
```cpp
greet("Alice"); // Pass "Alice" as an argument
```

- C++ function arguments use parentheses.

---

# Function Arguments in Go
```go
greet("Alice") // Pass "Alice" as an argument
```

- Values passed to a function are called arguments.

---



# Return Values in Python
```python
def add(a, b):
    return a + b
```

- Python functions use `return` for returning values.

---

# Return Values in C++
```cpp
int add(int a, int b) {
    return a + b;
}
```

- C++ functions specify return type.

---

# Return Values in Go
```go
func add(a, b int) int {
    return a + b
}
```

- Functions can return values, specified after the parameter list.

---



# Calling Functions with Return Values in Python
```python
result = add(3, 5)
print("3 + 5 =", result)
```

- Python captures return values in variables.

---

# Calling Functions with Return Values in C++
```cpp
int result = add(3, 5);
cout << "3 + 5 = " << result << endl;
```

- C++ captures and uses return values similarly.

---

# Calling Functions with Return Values in Go
```go
result := add(3, 5)
fmt.Println("3 + 5 =", result)
```

- Capture the return value of a function when calling it.

---



# Multiple Return Values in Python
```python
def swap(a, b):
    return b, a
```

- Python functions can return multiple values as well.

---

# Multiple Return Values in C++
```cpp
#include <iostream>
using namespace std;

void swap(int &a, int &b) {
    int temp = a;
    a = b;
    b = temp;
}
```

- In C++, multiple values can be returned through reference parameters.

---

# Multiple Return Values in Go
```go
func swap(a, b int) (int, int) {
    return b, a
}
```

- Functions can return multiple values, separated by commas.

---


# Calling Functions with Multiple Return Values in Python
```python
x, y = swap(10, 20)
print(f"Swapped: x={x}, y={y}")
```

- Python assigns multiple return values to multiple variables.

---


# Calling Functions with Multiple Return Values in Go
```go
x, y := swap(10, 20)
fmt.Printf("Swapped: x=%d, y=%d\n", x, y)
```

- Capture multiple return values using multiple variables.

---

# Named Return Values in Go
```go
func divide(dividend, divisor float64) (result float64, err error) {
    if divisor == 0 {
        err = errors.New("division by zero")
        return
    }
    result = dividend / divisor
    return
}
```

- You can name return values in the function signature.

---

# Named Return Values in Python
```python
def divide(dividend, divisor):
    if divisor == 0:
        return None, "division by zero"
    return dividend