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

	fmt.Println(ar)
	ar = mergeSort(ar)
	fmt.Println(ar)
}

func mergeSort(ar []int) []int {
	l := len(ar)
	if l == 1 {
		return ar
	}

	ar1 := mergeSort(ar[:l/2])
	ar2 := mergeSort(ar[l/2:])
	resAr := make([]int, 0)
	i := 0
	j := 0
	for i < len(ar1) && j < len(ar2) {
		if ar1[i] < ar2[j] {
			resAr = append(resAr, ar1[i])
			i++
		} else {
			resAr = append(resAr, ar2[j])
			j++
		}
	}
	if i < len(ar1) {
		resAr = append(resAr, ar1[i:]...)
	} else if j < len(ar2) {
		resAr = append(resAr, ar2[j:]...)
	}
	return resAr
}
