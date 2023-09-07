---
marp: true
theme: gradient
paginate: true
---

# Loops in GoLang
## Iteration and Repetition

Looking at various looping constructs in Go as compared to Python and C++

---

# For Loops in Go

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

- Use `for` loops for fixed iterations.
- Print numbers from 0 to 4.

---

# For Loops in Python
```python
for i in range(5):
    print(i)
```

- Python's `for` loop with `range()` for iteration.

---

# For Loops in C++
```cpp
for (int i = 0; i < 5; i++) {
    cout << i << endl;
}
```

- C++ `for` loop for similar behavior.

---

# While Loops in Go
```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

- Use `for` as a `while` loop in Go.

---

# While Loops in Python
```python
count = 0
while count < 5:
    print(count)
    count += 1
```

- Python's `while` loop for condition-based iteration.

---

# While Loops in C++
```cpp
int count = 0;
while (count < 5) {
    cout << count << endl;
    count++;
}
```

- C++ `while` loop for similar behavior.

---

# Looping Over Collections in Go
```go
fruits := []string{"apple", "banana", "cherry"}
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

- Use `for range` to iterate over collections in Go.

---

# Looping Over Collections in Python
```python
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(fruit)
```

- Python's `for` loop simplifies collection iteration.

---

# Looping Over Collections in C++
```cpp
vector<string> fruits = {"apple", "banana", "cherry"};
for (const string& fruit : fruits) {
    cout << fruit << endl;
}
```

- C++ `for-each` loop for collection iteration.

---

# Infinite Loops in Go
```go
for {
    fmt.Println("This is infinite")
    // Break or return to exit
}
```

- Use `for` without a condition for infinite loops.

---

# Infinite Loops in Python
```python
while True:
    print("This is infinite")
    # Break or return to exit
```

- Python's `while True` for infinite loops.

---

# Infinite Loops in C++
```cpp
while (true) {
    cout << "This is infinite" << endl;
    // Break or return to exit
}
```

- C++ `while (true)` for infinite loops.

---

# Summary
- **GoLang**, **Python**, and **C++** offer versatile loop constructs.
- `for`, `while`, and `for-each` variations.
- Be cautious with infinite loops.
