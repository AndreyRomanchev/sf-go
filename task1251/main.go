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
	ar := make([]int, 5)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	l := len(ar)
	for i := 1; i < l; i++ {
		elem := ar[i]
		j := i - 1
		for j >= 0 && elem < ar[j] {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = elem
	}
}
