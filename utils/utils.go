package utils

import (
	"one/math"
	"fmt"
	// "errors"
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

func PaintNeeded(width float64, height float64) (float64, error) {
	
	if height < 0 || width < 0 {
		// err := errors.New("height не может быть отрицательным")
		// // fmt.Println(err.Error())
		// fmt.Println(err)
		// err := fmt.Errorf("a height %.2f is invalid", height)
		// fmt.Println(err)
		// fmt.Println(err.Error())
		return 0, fmt.Errorf("width (%.2f) or height (%.2f) < 0", width, height)
	}
	
	area := area(width, height)
	metersPerLiter = 10.0
	fmt.Printf("%2.2f liters needed\n", calcPaintAreaNeed(area))
	return area, nil	
}

func Double(number int) int{
	number *= 2
	fmt.Printf("first %d\n", number)
	return number
}
