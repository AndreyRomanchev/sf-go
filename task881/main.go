package main

import (
	"fmt"
	"github.com/AndreyRomanchev/sf-go/task881/electronic"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Println("Brand:", p.Brand())
	fmt.Println("Model", p.Model())
	fmt.Println("Type:", p.Type())
	switch p := p.(type) {
	case electronic.SmartPhone:
		fmt.Println("OS:", p.OS())
	case electronic.StationPhone:
		fmt.Println("Number of buttons:", p.ButtonsCount())
	}
}

func main() {
	applePhone := electronic.NewApplePhone("iPhone 6", "iOS 8")
	androidPhone := electronic.NewAndroidPhone("Google", "Pixel 3a", "Android 10")
	stationPhone := electronic.NewRadioPhone("Panasonic", "KX-TG1612", 18)

	printCharacteristics(applePhone)
	printCharacteristics(androidPhone)
	printCharacteristics(stationPhone)
}
