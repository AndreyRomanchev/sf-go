package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(ar)
	selectionSortBiDimensional(ar)
	fmt.Println(ar)
}

func selectionSort(ar []int) {
	l := len(ar)
	for i := 0; i < l; i++ {
		min := ar[i]
		mini := i
		for j := i; j < l; j++ {
			if ar[j] < min {
				min = ar[j]
				mini = j
			}
		}
		ar[i], ar[mini] = ar[mini], ar[i]
	}
}

func selectionSortByMax(ar []int) {
	l := len(ar)
	for i := l - 1; i >= 0; i-- {
		max := ar[i]
		maxi := i
		for j := i; j >= 0; j-- {
			if ar[j] > max {
				max = ar[j]
				maxi = j
			}
		}
		ar[i], ar[maxi] = ar[maxi], ar[i]
	}
}

func selectionSortBiDimensional(ar []int) {
	l := len(ar)
	for i := 0; i < l/2; i++ {
		min, max := ar[i], ar[i]
		mini, maxi := i, i
		for j := i; j < l-i; j++ {
			if ar[j] < min {
				min = ar[j]
				mini = j
			}
			if ar[j] > max {
				max = ar[j]
				maxi = j
			}
		}
		ar[i], ar[mini] = ar[mini], ar[i]
		if i == maxi {
			maxi = mini
		}
		ar[l-i-1], ar[maxi] = ar[maxi], ar[l-i-1]
	}
}
