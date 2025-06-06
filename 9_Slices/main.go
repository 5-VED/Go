package main

import (
	"fmt"
	"slices"
)

// Slice -> Most used Construct
func main() {

	var num []int
	fmt.Println(num == nil)

	// Create a slice

	var slice = make([]int, 0)
	slice = append(slice, 10)
	fmt.Println(slice)

	var slice1 = make([]int, len(slice))

	// Copy Slice
	copy(slice1, slice)
	fmt.Println(slice1)

	//Slice operator
	var slice2 = []int{1, 2, 3}
	fmt.Println(slice2[0:2])

	// It also has inbuild slice package
	var slice3 = []int{1, 2}
	var slice4 = []int{1, 2}

	fmt.Println(slices.Equal(slice3, slice4))

	// 2D Slices
	var slece5 = [][]int{{1, 2}, {1, 2}}
	fmt.Println(slece5)

}
