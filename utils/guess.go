package utils

import (
	"fmt"
	"math/rand"
)

func Guess() {
	rnd := rand.Intn(100) + 1  // интересная реализация - тут когда +1, то диапазон от 1 до 100, иначе от 0 до 99
	// fmt.Printf("rnd %d\n", rnd)
	fmt.Println(rnd)
}