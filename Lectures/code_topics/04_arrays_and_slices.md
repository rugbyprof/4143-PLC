---
marp: true
theme: gradient
paginate: true
---

# Slide Title: Arrays and Slices in GoLang
## Declaration, Usage, and Alteration

---

# Declaring Arrays in Go
```go
var myArray [5]int // Declaring an array of 5 integers
```

- Declare an array in Go with a specific size.

---

# Using Arrays in Go
```go
myArray[0] = 42
```

- Access elements using zero-based indexing.
- Set the value of the first element to 42.

---

# Altering Arrays in Go
```go
myArray[2] = 99
```

- Modify elements by assigning new values.
- Here, we set the value of the third element to 99.

---

# Python Equivalent: Lists
```python
myList = [0] * 5  # Declare a list of 5 zeros
myList[0] = 42    # Access and modify as in Go
myList[2] = 99
```

- In Python, we use lists, which are dynamic arrays.

---

# Introduction to Slices in Go
```go
mySlice := []int{1, 2, 3, 4, 5} // Declare and initialize a slice
```

- Slices are dynamic arrays in Go.

---

# Using Slices in Go
```go
mySlice = append(mySlice, 6) // Append an element to the slice
```

- Modify slices using the `append` function.

---

# Python Equivalent: Lists (Again)
```python
myList = [1, 2, 3, 4, 5]
myList.append(6)  # Append an element in Python
```

- Python lists also support dynamic resizing.

---

# Common Operations on Arrays and Slices
- Slicing
- Reslicing
- Copying

---

# Slicing Arrays and Slices in Go
```go
subArray := myArray[1:4]    // Slice an array
subSlice := mySlice[1:4]    // Slice a slice
```

- Slicing extracts a portion of an array or slice.

---

# Slicing Lists in Python
```python
subList = myList[1:4]  # Slice a list in Python
```

- Python uses a similar slicing syntax.

---

# Reslicing in Go
```go
resliced := mySlice[:2] // Reslice to include the first two elements
```

- Reslicing changes the start index of a slice.

---

# Reslicing in Python
```python
resliced = myList[:2]  # Reslice in Python
```

- Python also supports reslicing lists.

---

# Copying Arrays and Slices in Go
```go
copyArray := myArray     // Copy an array
copySlice := make([]int, len(mySlice))
copy(copySlice, mySlice) // Copy a slice
```

- Copying creates a new array/slice with the same values.

---

# Copying Lists in Python
```python
copyList = myList.copy()  # Copy a list in Python
```

- Python's `copy` method creates a copy of a list.

---

# Summary
- Arrays in Go have a fixed size.
- Slices are dynamic arrays.
- Common operations include slicing, reslicing, and copying.
