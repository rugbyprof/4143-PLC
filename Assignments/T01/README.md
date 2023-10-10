## Test 1 - Go's OOP Concepts
#### Due: 10-10-2023 (Tuesday @ 9:30 a.m.) 

First test content is below. We will discuss in class and I will create some example questions as well.

|   #   | Name                                                                                     |
| :---: | :--------------------------------------------------------------------------------------- |
|   1   | [01 Not OOP](../../Lectures/OOP_Concepts/01-not_oop.md)                                  |
|   2   | [02 New Keyword](../../Lectures/OOP_Concepts/02-new_keyword.md)                          |
|   3   | [03 Slices of Structs](../../Lectures/OOP_Concepts/03-slices_of_structs.md)              |
|   4   | [04 Maps](../../Lectures/OOP_Concepts/04-maps.md)                                        |
|   5   | [05 Struct vs Interface](../../Lectures/OOP_Concepts/05-struct_vs_interface.md)          |
|   6   | [06 Methods vs Functions](../../Lectures/OOP_Concepts/06-methods_vs_functions.md)        |
|   7   | [07 Functions](../../Lectures/OOP_Concepts/07-functions.md)                              |
|   8   | [08 Interfaces](../../Lectures/OOP_Concepts/08-interfaces.md)                            |
|   9   | [09 Composition and Pckages](../../Lectures/OOP_Concepts/09-composition_and_packages.md) |
|  10   | [10 Packages in Depth](../../Lectures/OOP_Concepts/10-packages_in_depth.md)              |
|  11   | [11 Composition Over Inheritance](../../Lectures/OOP_Concepts/11-comp_over_inherit.md)   |


# Study Guide

## Is Go an OOP Language?

   While Go has features that allow developers to adopt some OOP-style patterns, it's not purely an object-oriented language. Here's why:

   1. **No Classes**: Go doesn't have the concept of "classes". Instead, it uses "structs" to define and group data. While you can attach methods to structs, the approach is different from traditional class-based OOP languages.

   2. **No Inheritance**: Go does not support the classic OOP inheritance model. You can't create a new struct (or type) that inherits the properties and methods of another struct. Instead, Go promotes composition over inheritance.

   3. **No Polymorphism through Subtype**: In many OOP languages, you can pass an instance of a subclass wherever a superclass is expected (often called Liskov substitution). In Go, polymorphism is achieved through interfaces, not subtyping.

   4. **No "Objects" per se**: While you can create instances of structs, these aren't "objects" in the classical OOP sense, because they don't carry with them any metadata or behavior definitions as classes do in languages like Java or C++.

   5. **Explicit over Implicit**: Go prefers clarity over the magic that sometimes happens in OOP languages. For instance, while OOP languages might have hidden behaviors due to constructors, destructors, or operator overloading, Go tends to be more transparent.

   ## Motivations for Using Go:

   1. **Concurrency**: One of Go's standout features is its built-in support for concurrent programming using goroutines and channels. It's designed for systems that run on modern hardware, where concurrent operations are more common.

   2. **Simplicity & Readability**: Go promotes writing simple, clear code. The language is designed to be straightforward, avoiding complex abstractions and features that can make code harder to understand.

   3. **Performance**: Go combines the ease of a dynamically-typed interpreted language with the efficiency of a statically-typed compiled language. It's optimized for performance and can achieve performance close to C or C++.

   4. **Cross-platform**: Go has a cross-compilation feature that makes it easy to build executables for different OS and architectures from a single codebase.

   5. **Standard Library**: Go's extensive standard library offers a range of utilities and supports many common tasks, reducing the need for third-party libraries.

   6. **Ecosystem & Community**: The Go community is vibrant, with a growing ecosystem of tools, libraries, and frameworks.

   7. **Designed for the Cloud**: Go is becoming a language of choice for many cloud-based applications, including those in the Kubernetes ecosystem.

   8. **Garbage Collection**: Go manages memory allocation behind the scenes, helping reduce the potential for memory leaks.

   Overall, while Go takes inspiration from OOP concepts, it merges them with procedural and functional programming paradigms to create a unique blend tailored for modern software development, especially in cloud and concurrent environments.

### Questions
- In Go, methods can be associated with what? (structs, classes, data types, etc. ??)

- Structs and interfaces. Is it a one to one relationship? Or can a struct implement more than one interface?

- Does go use special keywords to facilitate inheritance?

- What is the primary way in which Go provides code reuse and polymorphism if it doesn't support traditional class-based inheritance?

- In Go, how do you make a field or method of a struct "private" to its package, and how do you make it "public" or accessible from outside its package?

- How can you associate a function with a struct in Go to make it behave like a method of that struct?


## Go's Naming Conventions

Naming conventions in Go help make the code more readable and are an integral part of Go's philosophy of clarity and simplicity. Here's a brief overview of Go's naming conventions:


1. **Exported Names**: If an identifier (like a variable, type, function, or method name) starts with an uppercase letter, then it is exported, meaning it can be accessed from outside the package it's defined in. This serves as Go's version of "public" in many OOP languages.
   - Example: `Person`, `PrintMessage`

2. **Unexported Names**: If an identifier starts with a lowercase letter, it's unexported, or "private", and can't be accessed from outside its package.
   - Example: `personDetail`, `printDetail`

3. **CamelCase**: Go prefers CamelCase (or mixedCaps) for multi-word names. This applies to both exported and unexported identifiers.
   - Example: `calculateTotalAmount`, `ServerAddress`

4. **Package Names**: Package names are always lowercase, without underscores or mixed caps. They should be short, concise, and evocative of the package's purpose.
   - Example: `fmt`, `net`, `http`

5. **Getters**: Go does not use the `get` prefix for getter methods. If you have an attribute named `length`, the getter method should just be named `Length` (and not `GetLength`).
   - Example: If there's a `name` attribute, the getter is `Name()`.

6. **Avoid Stutter**: Go encourages avoiding repetitive naming. If a package name already gives context, avoid using that context in the identifier names.
   - Example: Prefer `cache.New()` over `cache.NewCache()`.

7. **Acronyms**: For acronyms, Go tends to keep the acronym uppercase. So, if you have an identifier with words like `URL`, `ID`, `HTTP`, etc., keep them uppercase.
   - Example: `HTTPHandler`, `UserID`

8. **Errors**: By convention, errors usually have names ending in `Error`.
   - Example: `PermissionError`, `TimeoutError`

Go's naming conventions are straightforward, emphasizing clarity and easy readability. They play a vital role in making Go code consistent and understandable across the community.

### Questions

- What is a Go convention for exported names?
- How do you mimic privacy in Go when naming functions?
- What is the recommended naming convention for packages in Go?
- Given a struct attribute named `size`, what would be the conventional getter method name in Go?
- For an acronym like `XML`, how should it be represented in a Go identifier when combined with another word, e.g., "parser"?
- What is the significance of an identifier starting with an uppercase letter in Go?
- How does Go handle the naming of getter methods for struct attributes? Provide an example.
- Why should Go developers avoid using repetitive naming in context with package names? 
- What is the Go naming convention for multi-word identifiers, and how does it differ from other conventions like snake_case?
- What is the conventional way to allocate memory and create a new instance of a struct in Go?
- How does Go determine if a field or method of a struct is exported outside of its package?
- How does Go's approach to memory allocation differ from many other languages when creating a new instance of a type, and why might this be beneficial?
- Describe the significance of naming conventions in Go, particularly concerning exported and unexported identifiers.
- Why doesn't Go use explicit access specifiers like `public` or `private` as seen in many other languages?

## Slices Overview

A **slice** is a dynamic, flexible view into the elements of an array. They're more common than arrays in Go due to their versatility.

**Characteristics of Slices**:

1. **Dynamic Size**: Unlike arrays, slices don't have a predefined size. They can grow or shrink as needed.

2. **Reference Semantics**: Slices reference an underlying array. Multiple slices can refer to the same array, making slice operations more memory efficient but also meaning changes in one slice can affect others.

3. **Components**: A slice is described by three components: a pointer to the first element, a length, and a capacity. 

4. **Creation**: Slices can be created using the built-in `make` function, by slicing an existing array or slice, or by using a literal without specifying a size.

   ```go
   s := make([]int, 3, 5) // type, length, and capacity
   ```

5. **Appending**: The `append` function can be used to add elements to a slice. It can increase the slice's capacity when needed.

## Structs in Go:

A **struct** is a composite data type that groups together variables (fields) under a single name. It's Go's way of creating more complex custom types.

**Characteristics of Structs**:

1. **Field Grouping**: Structs can contain fields of different data types, allowing you to bundle relevant data together.

2. **No Classes**: Go doesn’t have classes. Instead, structs are used to define and group data, and methods can be associated with them.

3. **Value Semantics**: By default, structs have value semantics, meaning when you assign a struct to a new variable or pass it to a function, the entire struct is copied.

4. **Anonymous Fields**: Go supports embedded or anonymous fields in structs. These fields don't have an explicit name, and their type becomes their implicit name.

5. **Initialization**: You can initialize a struct using named or unnamed field values.

   ```go
   person := Person{name: "John", age: 25}
   ```

### Questions

- Understand how the : syntax works with slices. 
- How do you append an element to a slice in Go?
- What is the difference between an array and a slice in Go?
- Explain the relationship between a slice, the underlying array, and how this relationship impacts memory usage in Go.
- What happens when you exceed the capacity of a slice while appending data, and how does it affect performance?
- What's the primary purpose of a struct?**
- Know how to initialize a struct.
- When you have a struct with an embedded struct, how do you access the embedded struct?
- How does Go handle "inheritance" with structs, given that it doesn't support traditional class-based OOP?
- Explain the difference between a value receiver and a pointer receiver when defining methods on structs in Go.
- Why might you choose to use named struct types over anonymous structs in Go?

## Anonymous vs Named Structs

### Anonymous Structs:

An anonymous struct is a struct that is defined without a specific name. These are typically used for one-off situations where you don't expect to reuse the struct type in multiple places.

#### Characteristics:

1. **In-place Definition**: Defined and instantiated in the same place.
2. **Limited Scope**: Best used for throwaway or one-off cases where the structure won't be reused elsewhere.
3. **No Method Association**: You can't associate methods with anonymous structs directly.

#### Example:

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}
fmt.Println(person.Name)  // Outputs: Alice
```

---

### Named Structs:

A named struct is a typical way of defining structs in Go where the struct type is given a specific name. Once defined, this type can be reused throughout your codebase.

#### Characteristics:

1. **Reusable**: You can use the same struct type in different parts of your program.
2. **Method Association**: You can define methods on named structs, making them more versatile.
3. **Exportable**: If the name starts with an uppercase letter, the struct type can be exported and used in other packages.

#### Example:

```go
type Person struct {
    Name string
    Age  int
}

// Method associated with the named struct
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

alice := Person{Name: "Alice", Age: 30}
alice.Greet()  // Outputs: Hello, my name is Alice
```

---

### Comparison & Use Cases:

1. **Short-term vs. Long-term**: 
   - **Anonymous Struct**: Useful for quick, temporary data structures that won't be reused—like reading a specific JSON structure from an API response.
   - **Named Struct**: Better for long-term, reusable data structures.

2. **Methods**:
   - **Anonymous Struct**: Cannot have methods directly associated.
   - **Named Struct**: Can have methods associated, providing more functionality.

3. **Clarity and Readability**:
   - **Anonymous Struct**: Useful when the context is clear, and the structure is simple. Overusing them, especially for complex structures, can reduce readability.
   - **Named Struct**: Provides better clarity, especially for complex or widely used data structures.

---

In practice, you'll likely use `named structs` most of the time due to their versatility and clarity. However, `anonymous structs` are a handy tool to have in your Go toolbox for those situations where you need a quick and uncomplicated data structure.

## Maps


### Maps in Go:

A **map** is a built-in associative data type that pairs unique keys to values (a key-value store). In other languages, maps might be called dictionaries, hashes, or associative arrays.

**Characteristics of Maps**:

1. **Dynamic Size**: Unlike arrays or slices, maps can grow as we add new pairs to them.

2. **Key Uniqueness**: Every key in a map is unique, and it pairs with exactly one value.

3. **Any Data Type**: Keys and values in a map don't need to be of the same type. However, all keys must be of the same type, and all values must be of the same type within a single map.

4. **Initialization**: Maps can be created using the built-in `make` function or using map literals.

   ```go
   m := make(map[string]int)
   m2 := map[string]int{"one": 1, "two": 2}
   ```

5. **Zero Value**: The zero value of a map is `nil`. A `nil` map has no keys and can't be assigned keys.

6. **Accessing Values**: You can access the value for a given key using the `map[key]` syntax. If the key doesn't exist, you'll get the zero value for the value type.

   ```go
   value := m["key"]
   ```

7. **Deleting Pairs**: The `delete` function removes a key-value pair from a map.

   ```go
   delete(m, "key")
   ```

8. **Checking Key Existence**: When retrieving a value, you can also get a second return value which is a boolean indicating if the key was present in the map.

   ```go
   value, exists := m["key"]
   ```

---

Maps are an essential part of Go's data structures, allowing for flexible and dynamic storage of paired data. They offer fast lookups, additions, and deletions, making them a popular choice for many programming scenarios in Go.

### Questions:

- Describe a map in Go?
- How do you declare a new empty map with string keys and int values?
- What is the result of accessing a key that doesn't exist in a map?
- How can you safely check if a key exists in a map?
- Describe the difference between maps and slices in Go.
- How can you delete an entry from a map in Go?
- Are the keys in a Go map guaranteed to be in any specific order?
- How would you handle scenarios where you need to differentiate between a zero value and a key that doesn't exist in a map?

## Interfaces in Go:

An **interface** in Go specifies a method set, representing a behavior. An interface doesn't describe data or implement the methods itself but instead specifies what methods must be implemented.

**Characteristics of Interfaces**:

1. **Method Set**: Interfaces are defined as a set of method signatures (name and type).

   ```go
   type Writer interface {
       Write([]byte) (int, error)
   }
   ```

2. **Implicit Implementation**: Unlike many other languages, in Go, you don't explicitly declare that a type implements an interface. If a type has methods that match the interface's method set, it implicitly implements the interface.

3. **Interface Values**: An interface value can hold any value that implements the necessary methods. It consists of two components: the value's concrete type and its value.

4. **Empty Interface**: An interface that doesn't specify any methods is known as the empty interface (`interface{}`). It can hold any value, making it analogous to `Object` in other languages.

   ```go
   var anyType interface{}
   ```

5. **Type Assertion**: To retrieve the underlying value from an interface, you can use type assertion.

   ```go
   value, ok := interfaceValue.(Type)
   ```

6. **Type Switches**: A type switch helps you determine the type of the concrete value held by an interface.

   ```go
   switch v := interfaceValue.(type) {
   case Type1:
       // ...
   case Type2:
       // ...
   }
   ```

7. **Composing Interfaces**: Interfaces can be composed of other interfaces, which is a way to create a new interface that embeds the method sets of multiple interfaces.

   ```go
   type ReaderWriter interface {
       Reader
       Writer
   }
   ```

---

Interfaces in Go are a powerful tool for abstraction and are central to Go's type system. They allow you to define the behavior that types must fulfill, making the system flexible and facilitating decoupled and modular code. They're especially key in writing generic functions and components that can operate on various data types as long as they satisfy certain behaviors.

### Questions 

- Describes how a type implements an interface in Go.
- What is an empty interface in Go?
- How does a type in Go indicate that it implements an interface?
- Describe the components of an interface value in Go.
- What is the purpose of type switches in the context of interfaces in Go?
- How can you compose multiple interfaces into a new interface in Go? Provide a brief example.
- Describes an interface in Go.
- What happens if a type does not implement all methods of an interface?
- What does an empty interface look like in Go?
- Explain the concept of "implicit implementation" of interfaces in Go.
- Why is the empty interface (`interface{}`) significant in Go?
- Can an interface in Go have fields? If not, why?
- How does Go handle situations where a type might need to implement multiple behaviors or functionalities?
  
## Functions in Go:

**Functions** are central to Go programming. They promote code reuse, abstraction, and clarity.

**Characteristics of Functions**:

1. **Definition**: Functions are defined using the `func` keyword. They can take zero or more parameters and return zero or more values.

   ```go
   func greet(name string) string {
       return "Hello, " + name
   }
   ```

2. **Multiple Return Values**: Go supports multiple return values, which is often used for returning a result and an error.

   ```go
   func divide(a, b float64) (float64, error) {
       if b == 0 {
           return 0, errors.New("cannot divide by zero")
       }
       return a / b, nil
   }
   ```

3. **Named Return Values**: Functions in Go can return named values. Once named, they can be treated as variables defined at the top of the function.

   ```go
   func rectangleArea(width, height float64) (area float64) {
       area = width * height
       return
   }
   ```

4. **Variadic Functions**: These can take an arbitrary number of arguments of a specified type using the `...` syntax.

   ```go
   func sum(numbers ...int) int {
       total := 0
       for _, num := range numbers {
           total += num
       }
       return total
   }
   ```

5. **First-Class Citizens**: Functions in Go are first-class citizens. They can be assigned to variables, passed as arguments, and returned from other functions.

   ```go
   func operation() func(int, int) int {
       return func(a, b int) int {
           return a + b
       }
   }
   ```

6. **Anonymous Functions**: Go supports anonymous functions (or lambda functions). They can also form closures.

   ```go
   add := func(a, b int) int {
       return a + b
   }
   ```

7. **Defer**: The `defer` keyword postpones the execution of a function until the surrounding function returns, typically used for cleanup actions.

   ```go
   func readFile(filename string) {
       file, err := os.Open(filename)
       defer file.Close()
       // ... read file
   }
   ```

---

Functions in Go provide the building blocks for structuring and organizing code. Their flexibility, from multiple return values to first-class citizenship, offers developers a range of patterns and conventions that can be tailored to various scenarios and requirements.

### Questions:

- Write a function that recieves two integers an returns an integer.
- In Go, how can you return multiple values from a function?
- What does the `defer` keyword do in a function?
- What is a closure in Go? 
- Describe the purpose and a use case for variadic functions in Go.
- What's the difference between passing by value and passing by reference in Go? 
- How does Go handle recursive functions, and what is a tail call?**
- What are first-class functions in Go, and how do they influence functional programming patterns in the language?

## Composition in Go:

In object-oriented programming, the phrase "favor composition over inheritance" is a guiding principle. While Go doesn't have classes and inheritance as in traditional OOP languages, it strongly emphasizes the use of composition.

**Characteristics of Composition in Go**:

1. **Embedding**: Instead of inheriting from a base or parent "class", Go allows you to embed types within other types. This is done by including the type you want to embed without a field name.

   ```go
   type Address struct {
       Street, City, State, Zip string
   }

   type Person struct {
       Name    string
       Address // Embedding the Address type
   }
   ```

2. **Accessing Methods**: When one type is embedded within another, methods defined on the embedded type become accessible on the encompassing type. This allows for a kind of "pseudo-inheritance" of methods.

   ```go
   func (a Address) FullAddress() string {
       return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
   }

   // Person type can now use FullAddress() method due to embedded Address.
   ```

3. **Overriding Methods**: While methods from the embedded type are accessible from the encompassing type, they can also be overridden by defining a method with the same name on the encompassing type.

4. **Interfaces and Composition**: Composition plays a big role in how Go uses interfaces. You can embed multiple interfaces into a new interface. This allows you to build up complex behaviors using smaller, more focused interfaces.

   ```go
   type Reader interface {
       Read(p []byte) (n int, err error)
   }

   type Writer interface {
       Write(p []byte) (n int, err error)
   }

   type ReadWriter interface {
       Reader // Embedding the Reader interface
       Writer // Embedding the Writer interface
   }
   ```

5. **Reusability**: By preferring composition, Go encourages the creation of small, reusable pieces that can be easily combined in various ways without the rigid structure and pitfalls of inheritance hierarchies.


Composition in Go promotes flexibility, reusability, and a clear way to build up complex structures and behaviors from simpler parts. It's a foundational concept that's crucial for writing idiomatic and effective Go code.

### Questions:

- In Go, how is composition typically achieved?
- Why might you prefer composition over inheritance in Go?
- What is the significance of an embedded struct in Go being "anonymous"?
- Describe the relationship between embedding and composition in Go.
- How does composition help in achieving the principle of "composition over inheritance"?
- Given a scenario where you have multiple types of vehicles, like `Car`, `Truck`, and `Bicycle`, all needing common functionalities like `Drive` and `Stop`, how would you design the system using composition in Go?
- Can an embedded struct in Go have its own methods? If so, how do they interact with the outer struct's methods of the same name?

## Packages vs. Modules in Go:

#### Packages:

A **package** is a way to group related Go source files together. Each file belongs to exactly one package, and the package scope allows for organizing and encapsulating code. Go encourages organizing code by their purpose rather than the functionality they offer.

**Characteristics of Packages**:

1. Defined with the `package` keyword at the top of each Go source file.
2. The package `main` is special in Go. A Go program needs a `main` package to produce an executable.
3. Packages provide code encapsulation, helping in organizing and reusing code.
4. To use a function, variable, or type from another package, its name must start with a capital letter (public visibility).

**Example**:

```go
// In a file named mathfuncs.go
package mathfuncs

func Add(a int, b int) int {
    return a + b
}
```

#### Modules:

A **module** is a collection of related Go packages stored in a file tree with a `go.mod` file at its root. Modules introduce a way to define versions of packages, allowing for better dependency management.

**Characteristics of Modules**:

1. Introduced in Go 1.11 to improve dependency management.
2. A module is initialized using `go mod init <module-name>`.
3. The `go.mod` file defines the module's path and its dependencies.
4. Modules can be versioned, aiding in reproducible builds.

**Example**:

To start a new module:

```bash
go mod init example.com/mymodule
```

This creates a `go.mod` file like:

```
module example.com/mymodule

go 1.14
```

###  Questions:

- Which keyword is used at the beginning of a Go file to declare its package?
- What is the significance of the `main` package in Go?
- Which file is crucial for defining a module in Go?
- Which command initializes a new module in Go?
- Describe the difference between a package and a module in Go.
- Why is it important for modules to have versioning?
- How do you make a function in a package available for other packages to use?
- Why did the Go team introduce modules, and how do they enhance package management?

## Composition over Inheritance

**Composition Over Inheritance** is a design principle suggesting that it's more beneficial to compose what an object can do (its behaviors or capabilities) rather than defining what it is in a strict class hierarchy.

**Reasons for Favoring Composition**:

1. **Flexibility**: Composition allows for greater flexibility in the final capabilities of an object. You can easily compose new behaviors or change behaviors without altering existing class hierarchies.

2. **Simpler Code**: Inheritance often leads to deep class hierarchies which can become hard to navigate, understand, and maintain. Composition tends to lead to flatter and more modular code structures.

3. **Avoids Inheritance Pitfalls**: Inheritance can lead to issues like the diamond problem (a dilemma that arises from multiple inheritance), which doesn't exist in languages or paradigms that prioritize composition.

4. **Better Encapsulation**: By favoring composition, you can ensure better data encapsulation. Parts of an object can change independently of its composed counterparts.

---

### Questions: 

- Which design principle suggests building object capabilities from smaller, reusable pieces instead of strict class hierarchies?
- Why might a developer prefer composition over inheritance?
- Which problem is associated with multiple inheritance and can be avoided by using composition?
- Describe a situation where composition might be more beneficial than inheritance.
- What does the principle "composition over inheritance" prioritize when designing objects?

## Encapsulation


**Encapsulation** refers to restricting access to certain details of an object and only showing the necessary features. It's a cornerstone of object-oriented design, ensuring that data and behavior are bundled together and that internal states of objects are protected from unauthorized external changes.

In Go, encapsulation is achieved primarily through exported and unexported identifiers:

1. **Exported Identifiers**: In Go, if an identifier (like a variable, type, or function name) starts with an uppercase letter, it's exported, meaning it can be accessed from outside the package it's defined in.

2. **Unexported Identifiers**: If an identifier starts with a lowercase letter, it's unexported, meaning it's private to the package it's defined in. This is the primary way to encapsulate data in Go.

3. **Methods**: Just like variables and types, methods can also be exported or unexported based on their name. This allows for creating a public API for a type while keeping certain behaviors or actions private.

---

### Questions:

- In Go, how do you define an identifier that can be accessed from outside its package?
_ If you see a function named `processData` in a Go package, what can you infer?
- Explain the significance of exported and unexported identifiers in Go.
- How does Go achieve encapsulation without using traditional access modifiers like `private` or `protected`?