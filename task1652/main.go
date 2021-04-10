package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var n int
	fmt.Print("Input number of for program to run: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Invalid number:", n)
	}
	var m sync.Mutex

	var counter int
	go func() {
		for {
			time.Sleep(1 * time.Second)
			m.Lock()
			counter++
			m.Unlock()
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Println(counter)
		}
	}()

	time.Sleep(time.Duration(n) * time.Second)
}
