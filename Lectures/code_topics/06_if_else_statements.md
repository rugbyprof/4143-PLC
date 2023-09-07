# If Else Statements
## Making Decisions in Code

Comparing and contrasting `C++`, `Python`, and `GoLang`

---

# If Else Statements

**Go**
```go
age := 18

if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are a minor.")
}
```
**Python**
```python
age = 18

if age >= 18:
    print("You are an adult.")
else:
    print("You are a minor.")
```
**C++**
```cpp
int age = 18;

if (age >= 18) {
    cout << "You are an adult." << endl;
} else {
    cout << "You are a minor." << endl;
}
```
---

# Nested If Else
<style scoped>
pre {
   font-size: 10px;
}
</style>
**GO**
```go
temperature := 25

if temperature > 30 {
    fmt.Println("It's hot outside.")
} else if temperature >= 20 {
    fmt.Println("It's a pleasant day.")
} else {
    fmt.Println("It's chilly.")
}
```
**Python**
```python
temperature = 25

if temperature > 30:
    print("It's hot outside.")
elif temperature >= 20:
    print("It's a pleasant day.")
else:
    print("It's chilly.")
```
**C++**
```cpp
int temperature = 25;

if (temperature > 30) {
    cout << "It's hot outside." << endl;
} else if (temperature >= 20) {
    cout << "It's a pleasant day." << endl;
} else {
    cout << "It's chilly." << endl;
}
```
---

# Ternary Operator in Go
**Go**
```go
//Go
value := 42
result := ""

if value >= 40 {
    result = "High"
} else {
    result = "Low"
}
```
**Python**
```python
value = 42
result = "High" if value >= 40 else "Low"
```
**C++**
```cpp
int value = 42;
string result = (value >= 40) ? "High" : "Low";
```
---

# Summary
- Conditional decisions are essential in programming.
- `if`, `else`, and `elif` statements in Go, Python, and C++.
- Ternary operators for concise assignments.
