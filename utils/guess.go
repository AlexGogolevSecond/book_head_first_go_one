package utils

import (
	"fmt"
	"math/rand"
)

func Guess() {
	// // работает и без time, но в книге используется
	// seconds := time.Now().Unix()
	// rand.Seed(seconds)  //lalala
	//rnd := rand.Intn(100) + 1  // интересная реализация - тут когда +1, то диапазон от 1 до 100, иначе от 0 до 99
	rnd := rand.Intn(100) + 1
	// fmt.Printf("rnd %d\n", rnd)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(rnd)
}
