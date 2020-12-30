package main

import "fmt"

func main() {
	var output string

	for i := 1; i <= 100; i++ {
		output = ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		fmt.Println(i, output)
	}
}
