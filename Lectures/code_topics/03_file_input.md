---
marp: true
theme: gradient
paginate: true
---


<!-- Slide 1: Title -->
# Getting File Input Data

---

<!-- Slide 2: Introduction -->
# Introduction

- Reading data from files is a common task in programming.
- C++ and Go provide different approaches to read data from files.

---

<!-- Slide 3: C++ File Input (ifstream) -->
# C++ File Input (ifstream)

- In C++, you use the `ifstream` class from the `<fstream>` library to read data from files.
  
---

<!-- Slide 3: C++ File Input (ifstream) -->
# C++ File Input (ifstream)

### Example (C++):

```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream inputFile("data.txt");
    if (!inputFile) {
        cerr << "Error opening file." << endl;
        return 1;
    }

    int num;
    inputFile >> num;
    cout << "Read from file: " << num << endl;
    inputFile.close();

    return 0;
}
```

---

<!-- Slide 4: Go File Input (os.Open) -->
# Go File Input (os.Open)

- In Go, you use the `os` package to open and read data from files.

---

<!-- Slide 4: Go File Input (os.Open) -->
# Go File Input (os.Open)
### Example (Go):

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var num int
    _, err = fmt.Fscanf(file, "%d", &num)
    if err != nil {
        fmt.Println("Error reading from file:", err)
        return
    }

    fmt.Println("Read from file:", num)
}
```

---

<!-- Slide 5: C++ File Input (getline) -->
# C++ File Input (getline)

- For reading lines from a file in C++, you can use the `getline` function.

--- 
# C++ File Input (getline)
### Example (C++):

```cpp
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main() {
    ifstream inputFile("data.txt");
    if (!inputFile) {
        cerr << "Error opening file." << endl;
        return 1;
    }

    string line;
    while (getline(inputFile, line)) {
        cout << "Read from file: " << line << endl;
    }

    inputFile.close();

    return 0;
}
```

---

<!-- Slide 6: Go File Input (bufio.NewScanner) -->
# Go File Input (bufio.NewScanner)

- For reading lines from a file in Go, you can use the `bufio` package and `Scanner`.

---

# Go File Input (bufio.NewScanner)
### Example (Go):

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println("Read from file:", line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }
}
```

---

<!-- Slide 7: Summary -->
# Summary

- In C++, `ifstream` is used for file input, and `getline` can be used for reading lines.
- In Go, the `os` package is used to open and read files, and `bufio.Scanner` is used for reading lines.
- Both languages provide mechanisms to read data from files.

---

<!-- Slide 8: Questions -->
# Questions

Any questions on getting file input data in C++ and Go?
