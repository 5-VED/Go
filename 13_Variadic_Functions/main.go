package main

import "fmt"

// Variadic Functions are function which can take n numbers of parameters
func add(nums ...int) int {
	var total int = 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

var slice []int = []int{1, 2, 3, 4, 5}

func main() {
	var total int = add(1, 2)
	fmt.Println("==== Total ======>", total)

	// Sum using spread operator
	var sum int = add(slice...)
	fmt.Println("==== Sum ======>", sum)
}
