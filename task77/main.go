package main

import (
	"fmt"
	"github.com/AndreyRomanchev/sf-go/task77/calc"
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

	c := calc.NewCalculator(firstNum, secondNum)
	res, err := c.Calculate(operation)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}
