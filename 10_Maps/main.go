package main

import (
	"fmt"
	"maps"
)

// Maps --> They are like objects in JS.
func main() {

	// Creating a Map
	m := make(map[string]string)
	n := make(map[string]string)
	m["first_name"] = "Ved"
	m["last_name"] = "Parmar"
	m["email"] = "vedparmar@gmail.com"

	// Deleting a specific field from map
	delete(m, "email")
	fmt.Println(len(m))

	// Function to remove all the elements fro  the map
	// clear(m)
	fmt.Println((m))

	//Funcion to check if property exist in out map or not
	v, ok := m["email"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// Check if two maps are equal
	fmt.Println(maps.Equal(m, n))
}
