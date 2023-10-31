Closures are a powerful concept in programming languages that allow functions to "remember" the lexical scope in which they were created. This means they can access variables even after the outer function has finished executing. Here are examples of closures in JavaScript and Python:

**JavaScript Closure Example:**
```javascript
function outer() {
  var outerVar = "I am from outer function";
  
  function inner() {
    console.log(outerVar); // inner can access outerVar even after outer function finished
  }
  
  return inner;
}

var closureFunction = outer();
closureFunction(); // Prints: "I am from outer function"
```

In this example, the `inner` function is returned from the `outer` function. The `inner` function forms a closure, retaining access to the `outerVar` variable even though `outer` has finished executing.

**Python Closure Example:**
```python
def outer():
    outer_var = "I am from outer function"
    
    def inner():
        print(outer_var) # inner can access outer_var even after outer function finished
        
    return inner

closure_function = outer()
closure_function() # Prints: "I am from outer function"
```

The Python example operates similarly to the JavaScript one. The `inner` function formed within the `outer` function retains access to `outer_var` after the `outer` function has completed its execution.

Both of these examples showcase how closures work by "closing over" variables from their containing scope, making them accessible even after the containing function has completed execution.