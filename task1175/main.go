package main

import (
	"fmt"
	"github.com/AndreyRomanchev/sf-go/task1175/queue"
)

func main() {

	q := queue.NewQueue()
	q.Print()
	q.Queue(1)
	q.Queue(2)
	q.Queue(3)
	q.Print()
	el1, _ := q.Dequeue()
	fmt.Println(el1)
	el2, _ := q.Dequeue()
	fmt.Println(el2)
	el3, _ := q.Dequeue()
	fmt.Println(el3)
	_, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
}
