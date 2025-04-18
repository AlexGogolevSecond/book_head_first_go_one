package main

import (
	"fmt"
	// "one/math"
	"one/utils"
)

func main() {
	fmt.Println("---")
	// fmt.Println(utils.Print("Alice"))
	// fmt.Println(math.Add(5, 3))
	// fmt.Println(math.Subtract(10, 4))
	// fmt.Println(utils.GetTime())
	// utils.Replace()
	// utils.PassFail()
	// utils.Erunda()
	// utils.Guess()
	// utils.Display1()
	a := utils.PaintNeeded(3.22, 2)
	fmt.Println(a)
	utils.PaintNeeded(4.2, 3)
	utils.PaintNeeded(5.2, 3.5)
}
