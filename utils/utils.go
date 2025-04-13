package utils

import (
	"one/math"
	"fmt"
)

func Print(name string) string {
	s := math.Add(3, 2)
	// fmt.Println("Hello " + name + "!" + fmt.Sprint(s) + "!!!")
	return "Hello " + name + "!" + fmt.Sprint(s) + "!!!"
}
