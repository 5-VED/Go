package main

import "fmt"

// for --> Only construct for looping in go
func main() {
	// While loop using for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Infinite loog
	// for {
	// 	fmt.Println(1)
	// }

	// Classic For loop
	for i := 0; i < 3; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	// Range
	for i := range 3 {
		fmt.Println(i)
	}

}
