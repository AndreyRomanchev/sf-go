package main

import (
	"fmt"
)

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch t {
		case Inch:
			value = value / 2.54
		case CM:
			value = value * 2.54
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type dimensions struct {
	length Unit
	width  Unit
	height Unit
}

type DimensionsInch struct {
	dimensions
}

type DimensionsCM struct {
	dimensions
}

func (d DimensionsInch) Length() Unit {
	return Unit{Value: d.length.Get(Inch), T: Inch}
}

func (d DimensionsInch) Width() Unit {
	return Unit{Value: d.width.Get(Inch), T: Inch}
}

func (d DimensionsInch) Height() Unit {
	return Unit{Value: d.height.Get(Inch), T: Inch}
}

func (d DimensionsCM) Length() Unit {
	return Unit{Value: d.length.Get(CM), T: CM}
}

func (d DimensionsCM) Width() Unit {
	return Unit{Value: d.width.Get(CM), T: CM}
}

func (d DimensionsCM) Height() Unit {
	return Unit{Value: d.height.Get(CM), T: CM}
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type BMW struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

type Mercedes struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

type Dodge struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (a BMW) Brand() string {
	return a.brand
}

func (a BMW) Model() string {
	return a.model
}

func (a BMW) Dimensions() Dimensions {
	return DimensionsCM{dimensions{
		length: Unit{a.dimensions.Length().Get(CM), CM},
		width:  Unit{a.dimensions.Width().Get(CM), CM},
		height: Unit{a.dimensions.Height().Get(CM), CM},
	},
	}
}

func (a BMW) MaxSpeed() int {
	return a.maxSpeed
}

func (a BMW) EnginePower() int {
	return a.enginePower
}

func (a Mercedes) Brand() string {
	return a.brand
}

func (a Mercedes) Model() string {
	return a.model
}

func (a Mercedes) Dimensions() Dimensions {
	return DimensionsCM{dimensions{
		length: Unit{a.dimensions.Length().Get(CM), CM},
		width:  Unit{a.dimensions.Width().Get(CM), CM},
		height: Unit{a.dimensions.Height().Get(CM), CM},
	},
	}
}

func (a Mercedes) MaxSpeed() int {
	return a.maxSpeed
}

func (a Mercedes) EnginePower() int {
	return a.enginePower
}

func (a Dodge) Brand() string {
	return a.brand
}

func (a Dodge) Model() string {
	return a.model
}

func (a Dodge) Dimensions() Dimensions {
	return DimensionsInch{dimensions{
		length: Unit{a.dimensions.Length().Get(Inch), Inch},
		width:  Unit{a.dimensions.Width().Get(Inch), Inch},
		height: Unit{a.dimensions.Height().Get(Inch), Inch},
	},
	}
}

func (a Dodge) MaxSpeed() int {
	return a.maxSpeed
}

func (a Dodge) EnginePower() int {
	return a.enginePower
}

func printAuto(a Auto) {
	fmt.Println("Brand:", a.Brand())
	fmt.Println("Model", a.Model())
	fmt.Println("Dimensions:", a.Dimensions().Height(), a.Dimensions().Length(), a.Dimensions().Width())
	fmt.Println("Max Speed", a.MaxSpeed())
	fmt.Println("Engine Power:", a.EnginePower())
}

func main() {
	a1 := BMW{brand: "BMW", model: "X5", dimensions: DimensionsCM{dimensions{
		length: Unit{100, CM},
		width:  Unit{100, CM},
		height: Unit{100, CM},
	}}, maxSpeed: 100, enginePower: 100}
	a2 := Mercedes{brand: "Mercedes", model: "S600", dimensions: DimensionsCM{dimensions{
		length: Unit{200, CM},
		width:  Unit{200, CM},
		height: Unit{200, CM},
	}}, maxSpeed: 200, enginePower: 200}
	a3 := Dodge{brand: "Dodge", model: "RAM", dimensions: DimensionsInch{dimensions{
		length: Unit{300, CM},
		width:  Unit{300, CM},
		height: Unit{300, CM},
	}}, maxSpeed: 300, enginePower: 300}

	printAuto(a1)
	printAuto(a2)
	printAuto(a3)
}
