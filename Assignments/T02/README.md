## Test 2 - Intro PLC Concepts
#### Due: 11-09-2023 (Thursday @ 9:30 a.m.) 

## Study Material

[02 Introduction To Programming Language Concepts](../../Lectures/02_Intro-Programming-Languages/README.md)

---

## Clarifications

### Imperative V Procedural

The difference is not as staggering as one might have thought. Really it's based on choice. If you use functions to organize code, it's procedural, if you don't and you just go top-down, that's the imperative part. 

**Imperative Programming** is a programming paradigm that uses statements to change a program's state. It consists of a series of commands for the computer to perform. This paradigm is about telling the 'how' – it outlines a sequence of steps to get to a result. Most languages that are considered to be procedural are also imperative because they follow a sequence of steps. Examples include C, Fortran, and even Python when used in a certain way.

**Procedural Programming** is a type of imperative programming that relies on well-organized sequence of procedures or routines (functions). The main focus is on the procedure in which the instructions are executed. Procedural programming takes the approach of breaking down the task into a collection of variables, data structures, and subroutines. So, while all procedural programming is imperative, not all imperative programming is procedural. For instance, you can write imperative code in a language like Python without necessarily using procedures or functions.

Here's a very basic example to illustrate the difference:

```python
# Imperative Approach
x = 10
y = 20
sum = x + y
print(sum)  # This is imperative but not necessarily procedural

# Procedural Approach
def add_numbers(x, y):
    return x + y

print(add_numbers(10, 20))  # This is both imperative and procedural
```

In the first snippet, we're giving imperative commands one after the other. In the second snippet, we encapsulate the logic into a procedure (function) called `add_numbers`, which is a hallmark of procedural programming.

When trying to determine if a piece of code is imperative or procedural, look for the following:

- **Imperative**: Does the code consist of a series of state changes or commands? If yes, it is imperative.
- **Procedural**: Are there functions or procedures that encapsulate part of the imperative logic? If yes, it is procedural.

- A language that supports both paradigms allows you to write in an `imperative` style without using functions (simply by changing states), 
- or in a `procedural` style where you encapsulate logic into procedures or functions.

In many cases, the distinction isn’t clear-cut because these paradigms are more about how you use the language rather than strict language features. **For example, Python supports multiple paradigms - it's object-oriented, but you can also write Python in an imperative or procedural style.** 

---

### Reflection 

Dylan brought up this topic when we were discussing `Attributes`,  `Annotations` and `Decorators`. I've included it since it really is another topic (like `hungarian notation`) which I understood, but never heard it phrases this way: `Reflection`....

Reflection in programming languages is a feature that enables a program to examine and modify its own structure and behavior at runtime. Reflection can be used for a variety of tasks, such as inspecting objects at runtime, invoking methods, or accessing fields that are not known at compile time.

### Purpose of Reflection:

1. **Introspection**: Reflection allows a program to inspect its own structure. For example, a program can list the methods and properties of an object, determine the type of an object, or find out what interfaces it implements.

2. **Dynamic Operation**: It enables a program to perform operations on objects that it only knows about at runtime. For example, it can instantiate classes, or call methods whose names are not known until runtime.

3. **Flexibility and Generalization**: Using reflection, libraries can operate on a wide variety of objects without knowing their specific types, making the code more flexible and general.

4. **Debugging and Testing**: Reflection can be useful for debugging, as it allows for examining the state of objects at runtime, and for testing, by making it easier to mock objects or inspect private members during test execution.

**Examples in Python:**

Python has a rich set of reflection capabilities through its built-in functions and standard libraries:

1. **Type Checking**:
   ```python
   x = [1, 2, 3]
   print(type(x))  # <class 'list'>
   print(isinstance(x, list))  # True
   ```

2. **Attribute Access and Manipulation**:
   ```python
   class MyClass:
       def __init__(self):
           self.x = 10
           self._y = 20
       
   obj = MyClass()
   print(hasattr(obj, 'x'))  # True
   setattr(obj, 'z', 30)  # Set new attribute 'z'
   print(getattr(obj, 'z'))  # 30
   print(obj.__dict__)  # {'x': 10, '_y': 20, 'z': 30}
   ```

3. **Method Invocation**:
   ```python
   class MyClass:
       def say_hello(self, name):
           return f"Hello, {name}"
   
   obj = MyClass()
   method = getattr(obj, 'say_hello')
   print(method('Terry'))  # Hello, Terry
   ```

4. **Introspection of Modules and Functions**:
   ```python
   import math
   import inspect
   
   print(inspect.getmembers(math, inspect.isfunction))  # List of math functions
   ```

5. **Working with Docstrings**:
   ```python
   def my_function():
       """This is a docstring for my_function."""
       pass
   
   print(my_function.__doc__)  # This is a docstring for my_function.
   ```

These examples show how Python's reflection capabilities can be used to dynamically interact with objects and classes, query their attributes, and invoke methods, all of which can be determined at runtime.

---

### Attributes & Annotations

Hey all. Look at the `my_decorator` function to see how decorators might work.  Other than that cool example, attributes and annotations are summed up below:

**Attributes** in Python are values associated with an object. This could be any data or functions (methods) that belong to an object. Think of attributes as variables belonging to an object.

For example, if we had a class `Dog`, an attribute might be `name` or `breed`:

```python
class Dog:
    def __init__(self, name, breed):
        self.name = name  # Attribute 'name'
        self.breed = breed  # Attribute 'breed'

fido = Dog("Fido", "Labrador")
print(fido.name)  # Outputs: Fido
```

**Annotations** in Python, introduced in Python 3.0 (PEP 3107), are a way of associating metadata with function parameters and return values. These annotations can be used for documentation, type checking, and more, but they don't affect the program's execution. They're just hints that can be used by various libraries or tools.

An example of a function with annotations might look like this:

```python
def greet(name: str) -> str:
    return 'Hello ' + name

# The annotation : str after `name` indicates that `name` should be a string
# The -> str indicates that the function should return a string
```

Annotations are not just for type hints; they can be any expression at all:

```python
def volume(level: 'int between 0 and 100') -> 'cubic meters':
    return level * 1000
```

Decorators, like you mentioned, can sometimes utilize annotations for additional functionality, such as validating input types or enforcing access control. They're like wrappers that you can place around a function to modify its behavior without changing the function's actual code.

Here's a bare-bones decorator example:

```python
def my_decorator(func):
    def wrapper(*args, **kwargs):
        # Do something before the function call
        result = func(*args, **kwargs)
        # Do something after the function call
        return result
    return wrapper

@my_decorator
def say_hello():
    return "Hello!"
```

```python
def response(func):
    def wrapper(*args, **kwargs):
        # get the size of the response
        # create forward and backward links for pagination
        # package those with the response
        return result
    return wrapper

@response
def api_call():
    # access database
    # get info
    # return info
```


So, in a nutshell, attributes are properties that objects carry around, while annotations are extra info about the types used by functions that tools and libraries can use to do cool stuff. They're both a kind of metadata, but they're used in different ways. Decorators can work with annotations but are more about doing things with functions, like modifying or enhancing them.

| Concept     | Description                                                                                                                                                                                                                                               |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Attributes  | Values associated with an object, which can be data or functions (methods). For example, in a `Dog` class, `name` and `breed` would be attributes representing the dog's name and breed respectively. They are like variables that belong to an object.   |
| Annotations | Metadata associated with function parameters and return values, introduced in Python 3.0. They can be used for documentation, type checking, etc., but don't affect program execution. Annotations can be any expression and are typically used as hints. |



### Goals

**Expect a matching question for the `Goals` and the `Golang` goals.**

These are overall for all languages:

| #   | Goal                          | Description                                                                                   |
| --- | ----------------------------- | --------------------------------------------------------------------------------------------- |
| 1   | Readability & Maintainability | Aim for code to be easily understood; vital for maintenance and debugging.                    |
| 2   | Efficiency                    | Enable writing optimized code; important for both system-level and high-level languages.      |
| 3   | Portability                   | Code can run on multiple platforms; achieved through VMs or interpreters like in Java/Python. |
| 4   | Scalability                   | Support development of all scales, from scripts to large applications.                        |
| 5   | Abstraction                   | Provide appropriate levels of detail abstraction for different tasks.                         |
| 6   | Expressiveness                | Allow developers to convey intentions clearly, reducing verbosity.                            |
| 7   | Safety & Security             | Prevent errors and vulnerabilities; some languages focus heavily on this like Rust.           |
| 8   | Community & Ecosystem         | A robust community and rich ecosystem are crucial for success; Python is a prime example.     |
| 9   | Interoperability              | Compatibility with other languages; scripting in C/C++ apps or C# with .NET.                  |
| 10  | Conciseness                   | Less verbose code to minimize errors and enhance understanding; seen in Python and Ruby.      |
| 11  | Performance                   | Control over resources for systems programming and computing; prioritize execution speed.     |
| 12  | Paradigm Support              | Tailor features to support programming paradigms like OOP, functional, or concurrent.         |
| 13  | Standardization               | Maintain consistency and prevent fragmentation across platforms and implementations.          |
| 14  | Ease of Learning              | Designed for beginners with simple syntax and a smooth learning curve.                        |
| 15  | Specialization                | Created for specific domains, like SQL for databases and MATLAB for numerical computing.      |



### GoLang Goals

This is for goloang.

| Goal                                 | Description                                                                            |
| ------------------------------------ | -------------------------------------------------------------------------------------- |
| Simplicity and Clarity               | Designed for straightforward, readable code with a clean syntax.                       |
| Efficiency                           | Aims for fast compilation and execution, with a garbage collector to minimize latency. |
| Concurrent and Parallel Programming  | Supports concurrent and parallel programming with goroutines and channels.             |
| Safety                               | Emphasizes memory safety and garbage collection to reduce programming errors.          |
| Robust Standard Library              | Includes a comprehensive standard library for various applications.                    |
| Strong Typing and Static Compilation | Features a strong, static type system to catch errors early.                           |
| Built-in Concurrency Primitives      | Provides concurrency primitives like goroutines and channels for efficient coding.     |
| Tooling                              | Offers a robust toolchain for effective development practices.                         |
| Cross-Platform Compatibility         | Facilitates writing code on one platform and compiling for others.                     |
| Garbage Collection                   | Uses garbage collection for automatic memory management.                               |
| Orthogonality                        | Encourages orthogonal code for cleaner, more maintainable structures.                  |
| Open Source and Community-Driven     | Is an open-source language with community-driven development.                          |



<!-- 

paradigm clarification

compiled vs interpreted -->

## Questions

Know the bare bones basics about the following languages: (static or dynamic and major paradigm)
1. Javascript
2. Python
3. Go
4. Java
5. Kotlin
6. PHP
7. C#
8. Swift
9. R
10. Ruby
11. C and C++
12. Matlab
13. TypeScript
14. Scala
15. SQL
16. HTML
17. CSS
18. NoSQL
19. Rust
20. Perl

   
Out of the list above know which are dynamic, static, weak, or strong.


Why are programming languages essential in technology and computer science?

Which programming languages are crucial for data analysis and machine learning?

What is the role of programming languages in web development?

   
What drives innovation in the field of programming languages?

   
What opportunities arise from proficiency in programming languages?


Know the difference between a syntax error vs a semantic (logic) error?

   
What could be the result of a semantic error in your code?

   
Think about statically-typed languages and what they give us vs dynamically-typed languages

   
Discuss the relationship between syntax and semantics in programming languages and why both are crucial for writing reliable and functional software.

   
In compiled languages, what is the role of a compiler?

What languages are known cross-platform portability due to the lack of a separate compilation step?

   
What is a major advantage of interpreted languages over compiled languages?

   
What optimizations might a compiler perform?

Which virtual machines uses Just-In-Time compilation for performance?

Compiled languages like \fillin[Go (Golang) and C++] are preferred for \fillin[performance-critical] applications.

Compiled languages like ______________________ are preferred for  ____________________ applications.
   
Interpreted languages such as \fillin[Python and Perl] are often used for their \fillin[portability and ease of development].

Interpreted languages such as ________________________ are often used for their _________________________________________.

Briefly explain the difference between semantic errors and syntax errors in the context of compiled versus interpreted languages.
   
In statically typed languages, what is required for variables?
   
What does dynamic typing prioritize over static typing?

Static and Dynamic typing do checks when respictively?
   
In a weakly typed language you can ... ?
In a strongly typed language you cannot ... ?
   
Explain in your own words what the difference is between static typing and dynamic typing.
   
Which of the following code samples is an example of procedural programming?
In terms of programming paradigms, how would you classify this code?
>Note: Given a code snippet identify if it is imperitive or procedural or oop.
   
What is the main characteristic that differentiates procedural programming from imperative programming?

   
What are attributes in Python?
   
Which of the following is a correct example of a function with annotated parameters in Python?
   
What are annotations in Python primarily used for?

   
What is the purpose of a decorator in Python?
 
What is reflection in programming languages primarily used for?

   
Which Python function would you use to check if an object has an attribute?

   
In Python, how can you get a list of all function objects in a module?

   
What is the purpose of the  \mintinline{python}{__doc__} attribute in Python classes and functions?


What does dynamic typing prioritize over static typing?

What is the main characteristic that differentiates procedural programming from imperative programming?

Is it possible for a single programming language to support multiple programming paradigms? Give examples of paradigm and language.


## Annotations v Decorators

### Annotations

Annotations in Python are used primarily for type hinting. They're a way to explicitly state the expected types of variables and function returns. This was a big deal when it was introduced in Python 3, as Python is a dynamically typed language where you don't normally declare variable types.

Here's a simple example of using annotations:

```python
def add_numbers(a: int, b: int) -> int:
    return a + b
```

In this function, `a: int` and `b: int` are annotations indicating that `a` and `b` should be integers, and `-> int` suggests that the function returns an integer.

Annotations don't affect the runtime behavior of your code. They are mainly for readability and for tools like type checkers, linters, and IDEs. Python's flexibility also allows annotations to be used for other purposes, like providing metadata, but that's more of an advanced, non-standard use case.

### Decorators

Now, decorators are a whole different story. They're used to modify or extend the behavior of functions or methods, without changing their actual code. Decorators are a very powerful feature in Python, allowing for aspects of functional programming like higher-order functions.

Here's an example of a decorator:

```python
def my_decorator(func):
    def wrapper():
        print("Something is happening before the function is called.")
        func()
        print("Something is happening after the function is called.")
    return wrapper

@my_decorator
def say_hello():
    print("Hello!")

say_hello()
```

In this example, `@my_decorator` is used to "decorate" `say_hello()`. The `say_hello()` function gets wrapped inside the `wrapper()` function, allowing for code to be executed before and after `say_hello()`.

### Key Differences

1. **Purpose**: Annotations are for type hinting and providing metadata. Decorators are for altering behavior.
   
2. **Impact on Code**: Annotations don't change how your code runs. Decorators can significantly change how your functions behave.

3. **Syntax**: Annotations use a colon (`:`) and arrow (`->`) syntax in function definitions. Decorators use the `@` symbol before a function definition.

4. **Usage Context**: Annotations are used within function definitions and variable declarations. Decorators wrap around a function or method, affecting its invocation.

In summary, annotations are like notes in your code, specifying types or other info, while decorators are like magic spells, changing how your functions work without altering their internal code. Both are awesome tools in Python's toolbox, serving different but equally valuable purposes!