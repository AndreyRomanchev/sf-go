package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	fmt.Print("Input the number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Cannot read the number:", err)
	}

	fmt.Println(hashint64(n))

	var s string
	fmt.Print("Input the string: ")
	_, err = fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Cannot read the string:", err)
	}

	fmt.Println(hashstr(s))
}

func hashint64(val int64) uint64 {
	return uint64(val) % 1000
}

func hashstr(val string) uint64 {
	var total uint64
	for k, v := range val {
		total += uint64(v) * uint64(math.Pow(10, float64(len(val)-k-1)))
	}
	return total % 1000
}
