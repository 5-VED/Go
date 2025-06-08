package main

import "fmt"

// Passed by value
func swapNumByRef(num *int) {
	*num = 5
	fmt.Println("====== Number before call ======>", *num)
}

// Passed by value
func swapNum(num int) {
	num = 5
	fmt.Println("====== Number before call ======>", num)
}

func main() {
	num := 10
	swapNum(num)
	fmt.Println("====== Number after change ======>", num)
}
