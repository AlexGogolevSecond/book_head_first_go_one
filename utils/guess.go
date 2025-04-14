package utils

import (
	"fmt"
	"math/rand"
)

func Guess() {
	rnd := rand.Intn(100) + 1
	fmt.Printf("rnd %d\n", rnd)
}