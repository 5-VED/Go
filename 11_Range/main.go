package main

import "fmt"

// Range --> Iterate over data structures
func main() {

	// For Slice
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)

	for i := 0; i < len(slice); i++ {
		fmt.Println("---- slice =--->", slice[i])
	}

	var sum int = 0

	// Using Range
	for _, element := range slice {
		sum = sum + element
		fmt.Println(element)
	}

	fmt.Println(sum)

	// For Maps
	m := make(map[string]string)
	m["first_name"] = "Ved"
	m["last_name"] = "Parmar"
	m["email"] = "vedparmar@appscrip.co"

	// Ranges are for each functions
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	// For Strings
	var str string = "ab"

	// It converts it into unicode code point rune
	for j, c := range str {
		fmt.Println(j, c)
	}

}
