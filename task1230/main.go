package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}

	bubbleSort(ar)
	fmt.Println(ar)

	bubbleSortReversed(ar)
	fmt.Println(ar)

	bubbleSortRecursive(ar, len(ar))
	fmt.Println(ar)
}

func bubbleSort(ar []int) {
	l := len(ar)
	swapFlag := false
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swapFlag = true
			}
		}
		if !swapFlag {
			break
		}
	}
}

func bubbleSortReversed(ar []int) {
	l := len(ar)
	swapFlag := false
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if ar[j] < ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swapFlag = true
			}
		}
		if !swapFlag {
			break
		}
	}
}

func bubbleSortRecursive(ar []int, l int) {
	if l == 1 {
		return
	}
	for j := 0; j < l-1; j++ {
		if ar[j] > ar[j+1] {
			ar[j], ar[j+1] = ar[j+1], ar[j]
		}
	}
	bubbleSortRecursive(ar, l-1)
}
