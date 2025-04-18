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


	// a, err := utils.PaintNeeded(3.22, 2)
	// fmt.Println(a, err)
	// b, err := utils.PaintNeeded(4.2, -3)
	// fmt.Println(b, err)
	// c, err := utils.PaintNeeded(25.2, 33.5)
	// fmt.Println(c, err)
	
	n := 5
	fmt.Printf("n=%d\n", n)
	n = utils.Double(n)
	fmt.Printf("n=%d\n", n)
}
