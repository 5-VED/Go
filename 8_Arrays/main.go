package main

import "fmt"

func main() {
	var nums [5]string
	fmt.Println(len(nums))
	fmt.Println((nums))

	nuber := [3]int{1, 2, 3}
	fmt.Println(nuber)

	// 2D array
	arr := [2][2]int{{2, 2}, {1, 1}}
	fmt.Println(arr)
}
