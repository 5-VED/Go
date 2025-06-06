package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Default Case")
	}

	// Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's WeekEnd")
	default:
		fmt.Println("Its working day")
	}

	// Type Switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Bool")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI(true)

}
