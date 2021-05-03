package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				fmt.Println("Goroutine number", i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Goroutines are done")
}
