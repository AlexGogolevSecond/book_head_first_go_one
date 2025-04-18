package utils

import (
	"one/math"
	"fmt"
)

var metersPerLiter float64

func Print(name string) string {
	s := math.Add(3, 2)
	// fmt.Println("Hello " + name + "!" + fmt.Sprint(s) + "!!!")
	return "Hello " + name + "!" + fmt.Sprint(s) + "!!!"
}

func area(width float64, height float64) float64 {
	return width * height
	// return int(width * height)
	
	// fmt.Printf("area: %.2f\n", area)
}

func calcPaintAreaNeed(area float64) float64 {
	return area / metersPerLiter
}

func PaintNeeded(width float64, height float64) float64 {
	area := area(width, height)
	metersPerLiter = 10.0
	fmt.Printf("%.2f liters needed\n", calcPaintAreaNeed(area))
	return area	
}

