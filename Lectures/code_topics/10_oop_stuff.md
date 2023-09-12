---
marp: true
theme: gradient
paginate: true
---

# Object-Oriented Programming in Go
## Encapsulation, Polymorphism, and Abstraction

---

# OOP Concepts in Go
- Go supports OOP principles through struct types and interfaces.
- While Go doesn't have traditional classes, it achieves OOP-like behavior.

---


# Encapsulation in C++
- In C++, encapsulation is achieved using access specifiers (`public`, `private`, `protected`) within classes.

```cpp
class Rectangle {
public:
    float Length; // Public field
private:
    float Width; // Private field
};
```

---

# Encapsulation in Go
- Encapsulation is achieved by exporting or unexporting fields in a struct.
- Capitalized fields are exported, while uncapitalized ones are not.

```go
package shapes

type Rectangle struct {
    Length  float64 // Exported field
    width   float64 // Unexported field
}
```

---



# Polymorphism in C++
- In C++, polymorphism is achieved through classes and inheritance.
- Base and derived classes are used for polymorphic behavior.

```cpp
class Shape {
public:
    virtual float Area() const = 0; // Pure virtual function
};

class Rectangle : public Shape {
    // Implement Area() in Rectangle
};
```

---

# Polymorphism in Go
- Polymorphism is achieved through interfaces in Go.
- Types implicitly implement an interface if they define its methods.

```go
type Shape interface {
    Area() float64
}
```

---



# Abstraction in C++
- In C++, abstraction is achieved using abstract base classes and pure virtual functions.

```cpp
class Reader {
public:
    virtual int Read(char* buffer, int size) = 0;
};

class FileReader : public Reader {
    // Implement Read() in FileReader
};
```

---

# Abstraction in Go
- Abstraction is achieved through interfaces in Go.
- Interfaces define a set of methods without specifying the implementation.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

---

# Comparison: Encapsulation

- C++ uses access specifiers (public, private, protected) within classes.
- Go uses capitalized field names for exported fields.

```cpp
class Rectangle {
public:
    float Length; // Public field
private:
    float Width; // Private field
};
```

```go
type Rectangle struct {
    Length  float64 // Exported field
    Width   float64 // Unexported field
}
```

---

# Comparison: Polymorphism
- C++ uses classes and inheritance with virtual functions.
- Go uses interfaces to define polymorphic behavior.

```cpp
class Shape {
public:
    virtual float Area() const = 0; // Pure virtual function
};

class Rectangle : public Shape {
    // Implement Area() in Rectangle
};
```
```go
type Shape interface {
    Area() float64
}
```


---

# Comparison: Abstraction

- C++ uses abstract base classes with pure virtual functions.
- Go uses interfaces to define abstract behavior.



```cpp
class Reader {
public:
    virtual int Read(char* buffer, int size) = 0;
};

class FileReader : public Reader {
    // Implement Read() in FileReader
};
```

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

---

# Summary
- Go achieves OOP-like behavior through structs and interfaces.
- Encapsulation, polymorphism, and abstraction are supported.
- C++ uses classes and inheritance for OOP.

