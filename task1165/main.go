package main

import (
	"fmt"
	"github.com/AndreyRomanchev/sf-go/task1165/stack"
)

func main() {

	a := stack.NewStack()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	el1, _ := a.Pop()
	fmt.Println(el1)
	el2, _ := a.Pop()
	fmt.Println(el2)
	el3, _ := a.Pop()
	fmt.Println(el3)
	_, err := a.Pop()
	if err != nil {
		fmt.Println(err)
	}
}
