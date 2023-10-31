## Go Maps

Maps are collections of key-value pairs. 

Lets work through some of the basic ideas with maps, starting off with small code snippets of usage of maps as a data structure in Go. Then we will move into using maps within slices, and also using maps to hold structs. 


1. **Creating a Map**:

   To create a map in Go, you use the `make` function with the `map` keyword and specify the types for keys and values. Here's an example of creating a map that maps strings to integers:

   ```go
   // Create a map with string keys and integer values
   myMap := make(map[string]int)
   ```

2. **Adding Key-Value Pairs**:

   You can add key-value pairs to a map using the assignment operator (`=`) or the `make` function. For example:

   ```go
   // Adding key-value pairs to the map
   myMap["apple"] = 3
   myMap["banana"] = 2
   ```

   This code assigns the value `3` to the key `"apple"` and `2` to the key `"banana"` in the `myMap` map.

3. **Accessing Values**:

   You can retrieve the value associated with a key by using the key in square brackets. For example:

   ```go
   // Accessing values in the map
   fmt.Println(myMap["apple"]) // Output: 3
   ```

   If the key doesn't exist in the map, it returns the zero value for the value type. In this case, it would be `0` for an integer.

4. **Updating Values**:

   To update the value associated with a key, simply assign a new value to that key:

   ```go
   // Updating a value in the map
   myMap["apple"] = 5
   ```

   This code updates the value associated with the key `"apple"` to `5`.

5. **Checking for Key Existence**:

   You can check if a key exists in a map using a two-value assignment. The second value indicates whether the key exists in the map or not:

   ```go
   // Check if a key exists in the map
   value, exists := myMap["apple"]
   if exists {
       fmt.Printf("Value for apple: %d\n", value)
   } else {
       fmt.Println("Key 'apple' does not exist in the map")
   }
   ```

6. **Deleting Key-Value Pairs**:

   To remove a key-value pair from a map, you can use the `delete` function:

   ```go
   // Delete a key-value pair from the map
   delete(myMap, "banana")
   ```

   This code removes the key `"banana"` and its associated value from the `myMap` map.

7. **Iterating Over a Map**:

   You can iterate over the key-value pairs in a map using a `for` loop:

   ```go
   // Iterate over key-value pairs in the map
   for key, value := range myMap {
       fmt.Printf("Key: %s, Value: %d\n", key, value)
   }
   ```

   The `range` keyword is used to iterate over the map, and it returns both the key and value in each iteration.

8. **Map Literals**:

   You can use map literals to declare and initialize maps in a single line:

   ```go
   // Map literal to declare and initialize a map
   fruitCount := map[string]int{
       "apple": 3,
       "banana": 2,
   }
   ```
   
Let's work with a couple examples to solidify what maps are good for.

Here's an example of a slice of maps representing data for a list of people, where each person's information is stored as a map with keys for "Name," "Age," and "City":
```go
package main

import (
    "fmt"
)

func main() {
    // Create a slice of maps representing people's data
    people := []map[string]interface{}{
        {"Name": "Alice", "Age": 30, "City": "New York"},
        {"Name": "Bob", "Age": 25, "City": "Los Angeles"},
        {"Name": "Charlie", "Age": 35, "City": "Chicago"},
        {"Name": "David", "Age": 28, "City": "San Francisco"},
    }

    // Print the data for each person
    for i, person := range people {
        fmt.Printf("Person %d:\n", i+1)
        fmt.Printf("Name: %s\n", person["Name"])
        fmt.Printf("Age: %d\n", person["Age"])
        fmt.Printf("City: %s\n", person["City"])
        fmt.Println()
    }
}
```
In the above example, you can clearly see how maps could be very easily used to store JSON files, Database objects, etc. 

Here's an example using a map in Go with a Dungeons & Dragons (D&D) theme, including initialization, updating, searching for a key, printing the map, deleting an item, and placing character structs in a map.

```go
package main

import (
    "fmt"
)

// Define a struct to represent a D&D character
type Character struct {
    Name      string
    Class     string
    Level     int
    HitPoints int
}

func main() {
    // Initialize an empty map with string keys and Character values
    characters := make(map[string]Character)

    // Add character information to the map
    characters["Gandalf"] = Character{Name: "Gandalf", Class: "Wizard", Level: 10, HitPoints: 80}
    characters["Conan"] = Character{Name: "Conan", Class: "Warrior", Level: 8, HitPoints: 120}
    characters["Elrond"] = Character{Name: "Elrond", Class: "Elf", Level: 12, HitPoints: 95}

    // Print the entire map
    fmt.Println("Characters:")
    for name, character := range characters {
        fmt.Printf("%s (%s, Level %d) - HP: %d\n", name, character.Class, character.Level, character.HitPoints)
    }

    // Search for a character by name and check if it exists
    name := "Conan"
    character, exists := characters[name]
    if exists {
        fmt.Printf("%s (%s, Level %d) - HP: %d\n", name, character.Class, character.Level, character.HitPoints)
    } else {
        fmt.Printf("%s not found in the character map\n", name)
    }

    // Delete a character from the map
    delete(characters, "Gandalf")

    // Check if Gandalf is still in the map
    _, exists = characters["Gandalf"]
    if !exists {
        fmt.Println("Gandalf has been removed from the character map")
    }
}
```

In this D&D-themed example:

- We initialize an empty map `characters` with string keys and `Character` struct values using `make`.
- We add character information for Gandalf, Conan, and Elrond to the map.
- We print the entire map, displaying the characters' names, classes, levels, and hit points.
- We search for a character by name (`Conan`) and check if it exists in the map.
- We delete a character (`Gandalf`) from the map.
- We check if Gandalf is still in the map after deletion.

This example demonstrates the use of maps with D&D character data, including searching, deleting, and printing the map contents.

[Previous: 03-slices_of_structs.md](./03-slices_of_structs.md) | [Next: 05-struct_vs_interface.md](./05-struct_vs_interface.md)
