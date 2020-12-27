package main

import (
	"fmt"
)

var firstNum float64
var operation string
var secondNum float64

func main() {
	_, err := fmt.Scanln(&firstNum)
	if err != nil {
		fmt.Println("Cannot read first number:", err)
	}

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Cannot read operation:", err)
	}

	_, err = fmt.Scanln(&secondNum)
	if err != nil {
		fmt.Println("Cannot read second number:", err)
	}

	switch operation {
	case "+":
		fmt.Println(firstNum + secondNum)
	case "-":
		fmt.Println(firstNum - secondNum)
	case "*":
		fmt.Println(firstNum * secondNum)
	case "/":
		fmt.Println(firstNum / secondNum)
	default:
		fmt.Println("Unknown operation:", operation)
	}
}
