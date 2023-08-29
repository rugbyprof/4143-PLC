# Overview of Programming Language Concepts

Concepts, terminology, definitions, and more help us organize our thoughts and opinions in such a way that we can better discuss the world and its topics. Maybe argue, maybe agree, but mostly converse.

## The Main Topics

Why do I bring up terms such as "argue" or "disagree"? Well, for example, OOP is huge but some people hate it. It's an overkill solution most of the time, where a collection of functions will suffice. And then, some claim that `functional programming` will be the death of OOP. I don't have a clue about the future, I just know that this is one argument in our CS world. So, not everyone is always on the same page. However, this is a good thing! It opens up a dialogue and hopefully a peaceful discussion about critical topics in CS. And when I say peaceful, I mean bullshit. Someone's getting punched in the mouth. 

My point is that perhaps none of us will agree on *the* set of main topics in the study of programming languages, but I'm going to list the topics that we will discuss in this course. And if anyone is getting punched in the mouth, let me know, so I can hide. 

1. **Names, Binding, and Scope (Declarations)**  
   - How do we give names to entities? And when we encounter a name, how do we know which entity it refers to?

2. **Evaluation (Expressions)**
   - How do we express computations, using values and operators?

3. **Execution (Control Flow)** 
   - How to we organize computation in time? What actions or effects can we produce?

4. **Types**  
   - How do we classify values so that they may behave in certain, predictable ways?

5. **Functional Abstraction (Subroutines and Coroutines)** 
   - How can we abstract computations into chunks so that we can invoke them whenever we need them?

6. **Data Abstraction (Objects and Modules)**
   - How do we make little bundles of data together with behavior?

7. **Concurrency**  
   - How do we arrange to do different computations at the “same” time (safely)?

8. **Metaprogramming**  
   - How can our programs know about themselves? How can we answer questions about the code itself?

> Are there others?  
>   
> What about things like performance? readability? maintainability?
> security?  
>   
> These are higher-level, qualitative concepts that are all important to
> be sure, but they are more cross-cutting concerns that apply to a
> language as a whole. The list above is focused on more technical,
> academic concerns of what can be expressed and how.  
>   
> But that said, read on to the end, because there’s something
> interesting at the end of these notes.

## Details

Let’s look a little deeper into each of these areas.

### Names

We have to be able to name things. Naming is indispensible even in human
languages.

-   `Binding`: associating names with entities
-   `Scope`: the region of text in which a binding is active
-   `Visibility`: sometimes you can see a name, sometimes you can’t
-   `Extent`: the lifetime of a binding
-   `Storage class`: “Where” are certain named (entities) stored?
-   `Linkage`: How, if at all, are names accessible from the outside?
-   `Defining` vs. `referencing` occurrences
-   `Overloading` and `aliasing`
-   `Polymorphism` (static/dynamic)
-   `Bound` vs. `free names`
-   `Shallow` vs. `deep binding`
-   `Static` vs. `dynamic scope`

### Expression Evaluation

An `expression` is either a value, or one or more
expressions joined by operators. Expressions are
`evaluated` to produce a value.

-   Operators (Precedence, assocativity, fixity, arity)
-   Evaluation order (defined, undefined, short-circuit)
-   Side effects
    -   Lvalues vs. Rvalues
    -   Initialization vs. Assignment
    -   Assignables?
-   Eager vs. Lazy Evaluation
-   Macros

### Control Flow

Computations often are required to occur in time. Some constructs
produce `actions`, or
`effects`. Some are triggered by
`events`. Many languages feature
`statements` (a.k.a.
`commands`) that are
`executed`.

-   `Sequencing` (semicolon, comma)
-   `Selection` (if, switch, case)
-   `Iteration` (for, while, forEach, takeWhile)
-   `Recursion`
-   `Procedural/Functional abstraction`
-   `Non-determinism`
-   `Concurrency`
-   `Disruption` (return, break, continue, throw)

### Types

A value’s type constrains the way it may be used in a program. Types
impose constraints on what we can and cannot say.

-   Fixed set of types or can you create new ones?
-   Are types extra-lingual or are they values you can manipulate
    (first-class types)?
-   Types vs. Classes vs. Typeclasses
-   Type Expressions
-   Equivalence, compatibility, checking, coercion
-   Inference
-   Spectra: static vs. dynamic, strong vs. weak, manifest vs.
    inferential
-   Parameterized types (generics), covariance, contravariance,
    invariance
-   Dependent types
-   Boolean, numeric, string, ...
-   Collections (arrays, lists, set, dictionaries)
-   Sum and product types
-   Optionals
-   Regular Expressions and other pattern types

### Subroutines

Subroutines, a.k.a. routines, subprograms, procedures, or functions, may
just be the most important building block of good code. They are the
most basic kind of code abstraction.

-   `Procedures` (abstractions for commands)

-   `Functions` (abstractions for expressions)

-   `Signatures`
    -   Names, modes, types, patterns, defaults, rest parameters

-   `Parameters and Arguments`
    -   Exactly one parameter? Or zero to many?
    -   Exactly one return value? Or zero to many?

-   `Parameter Association`

    - `Positional` vs. `Named`
    - Value, value-result, reference, name

-   `Higher-order functions`

-   `Closures`

### Modules

“Bob Barton, the main designer of the B5000 and a professor at Utah had
said in one of his talks a few days earlier: "The basic principle of
recursive design is to make the parts have the same power as the whole."
For the first time I thought of the whole as the entire computer and
wondered why anyone would want to divide it up into weaker things called
data structures and procedures. Why not divide it up into little
computers, as time sharing was starting to? But not in dozens. Why not
thousands of them, each simulating a useful structure?” — Alan Kay

-   Modules, packages, namespaces, classes, units, etc.
-   Import and export
-   Open vs. Closed
-   Object orientation
-   Separate compilation (the physical organization of a system)

### Concurrency

In the real world, things don’t happen in a single sequence. Things can
happen at anytime. Different things can happen at the same time. What is
time, anyway?

-   Threads vs. processes
-   Events
-   Single-threaded event queue vs. preemptive threading
-   Shared resources and synchronization
-   Scheduling
-   Asychronicity: callbacks, promises, async/await
-   Coroutines (yield)
-   Channels vs. Processes
-   Buffered vs. Unbuffered communication

### Metaprogramming

You might write a HR program that can answer questions like “What is the
average salaray of employees in the marketing department?” But can you
also ask it ”How many local variables are you made up of, and what are
their types?” Metaprogramming is writing programs that manipulate
programs (including themselves).

-   Introspection
-   Reflection
-   Eval
-   Decorators
-   Metaclasses

## Going Beyond Concepts

A lot of these concepts feel technical, dry, and even somewhat
language-independent. But you can’t really *know* a language without
knowing its history (who created it and why? When? What problems were
being solved? Was it built as a reaction to something?), and its vision.

A good vision helps to make a language learnable, usable, and maybe even
successful. Here’s a fantastic talk in which vision is mentioned a
couple times:

## Summary

We’ve covered:

-   Eight major concepts in the study of programming languages
-   Higher-level ideas


Sources: 

https://cs.lmu.edu/~ray/notes/plconcepts/
Chat GPT