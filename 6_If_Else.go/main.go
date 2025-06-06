package main

import "fmt"

func main() {
	var age int = 18

	if age >= 18 {
		fmt.Println("Person is adult")
	} else {
		fmt.Println("Person is not adult")
	}

	if age := 16; age > 12 {
		fmt.Println("Person is an adult")
	}

	// No ternary operator in Go for now.

}
