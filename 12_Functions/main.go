package main

import "fmt"

// Declare add Funcion
func add(a int, b int) int {
	return a + b
}

func languages() (string, string, bool) {
	return "Hello", "World", true
}

func fxn(fn func(a int) int) int {
	return fn(1)
}

func main() {
	var result int = add(5, 1)
	fmt.Println("===== Sum =====>", result)

	// Go functions can return multiple languages
	lang1, lang2, lang3 := languages()
	fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int {
		return 2
	}

	var response int = fxn(fn)
	fmt.Println(response)
}
