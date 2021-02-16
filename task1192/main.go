package main

import "fmt"

const n = 10
var weatherArray = [n]float32{10,11,12,13,14,15,16,17,18,19.1}

func avgT(array [n]float32) float32 {
	var sum float32
	for _, item := range array {
		sum += item
	}
	return sum / float32(len(array))
}

func printT(t float32) {
	fmt.Println("Current average temperature is", t)
}

func main() {
	printT(avgT(weatherArray))
	weatherArray[5] = 11.2
	printT(avgT(weatherArray))
	weatherArray[8] = 14.4
	printT(avgT(weatherArray))
	weatherArray[0] = 100
	printT(avgT(weatherArray))
}
