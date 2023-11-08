## Test 2 - Intro PLC Concepts
#### Due: 11-09-2023 (Thursday @ 9:30 a.m.) 

## Study Material

[02 Introduction To Programming Language Concepts](../../Lectures/02_Intro-Programming-Languages/README.md)

What does dynamic typing prioritize over static typing?


## Imperative V Procedural

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


<!-- 
decorators attributes / 

paradigm clarification

goals (summarize and filter)

compiled vs interpreted -->