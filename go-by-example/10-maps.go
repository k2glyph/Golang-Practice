package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	fmt.Println("Maps: ", m)
	// Getting value from map
	v1 := m["key1"]
	fmt.Println("Key1:", v1)
	// Find the length of map
	fmt.Println("Map length: ", len(m))
	// Deleting key from map
	delete(m, "key1")
	fmt.Println("Maps: ", m)
	_, key1Delete := m["key1"]
	fmt.Println("Retriving deleted item:", key1Delete)

	// Initializing map in same declaration
	m2 := map[string]string{"key1": "value1", "key2": "value2"}
	fmt.Println("Map2", m2)
}
