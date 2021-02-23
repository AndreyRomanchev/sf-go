package main

import "math/rand"

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

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
