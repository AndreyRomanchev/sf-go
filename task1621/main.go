package main

import (
	"fmt"
	"time"
)

func main() {

	var n int
	fmt.Print("Input number of goroutines: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Invalid number:", n)
	}

	for i := 0; i < n; i++ {
		go func(id int) {
			for {
				time.Sleep(time.Second)
				fmt.Println("Goroutine ID:", id)
			}

		}(i)
	}

	fmt.Println("Press the Enter Key to stop")
	_, _ = fmt.Scanln()
}
