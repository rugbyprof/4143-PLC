## Go Maps

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

[Previous: 03-slices_of_structs.md](03-slices_of_structs.md) | [Next: 05-struct_vs_interface.md](05-struct_vs_interface.md)
