package main

import (
	"fmt"
	"sync"
)

func goroutine(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goroutine(i, &wg)
	}

	wg.Wait()
}
