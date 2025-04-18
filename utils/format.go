package utils

import (
	"fmt"
)

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func Display1() {
	fmt.Printf("%.2f liters needed\n", 1.8199999999999998)
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
	// Erunda()  // !!!! пикольно - функции, которые определены в текущем проекте не нужно импортировать - они уже доступны
	// GetTime() // тут также
	// lalala()  // функции, у которых имя с маленькой буквы могут использоваться в других функциях только текущего пакета
	repeatLine("Hello, World!", 3)
}


