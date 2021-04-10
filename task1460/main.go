package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Print("Enter first array size: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Cannot read the number:", err)
	}
	nslice := make([]int, n)
	fmt.Print("Enter second array size: ")
	_, err = fmt.Scanln(&m)
	if err != nil {
		fmt.Println("Cannot read the number:", err)
	}
	mslice := make([]int, m)

	fmt.Println("Enter first array: ")
	for i := 0; i < n; i++ {
		nslice = append(nslice)
	}

	fmt.Println("Enter second array: ")
	for i := 0; i < m; i++ {
		mslice = append(mslice)
	}

	comp := make(map[int]int)

	for _, v1 := range nslice {
		for _, v2 := range mslice {
			if nslice[v1] == mslice[v2] {
				comp[v1] = v2
			}
		}
	}

}
