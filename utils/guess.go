package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReader(os.Stdin) // Создать «буферизованное средство чтения» для получения текста с клавиатуры
	fmt.Print("Your guess: ")
	input, err := reader.ReadString('\n')  // читаем данные из консоли
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)  // удаляем символы справа/слева
	guess, err := strconv.Atoi(input)  // преобразуем строку в int
	if err != nil {
		log.Fatal(err)
	}
	
	for x:=1; x < 5; x++ {
		y := x + 10
		fmt.Println("y" + strconv.Itoa(y))  // !!! пробразуем int в строку
	}

	// for x:=10; x > 0; x-=2 {
	// 	fmt.Println(x)
	// }
	
	if guess < rnd {
		fmt.Println("Oops. Your guess was LOW")
	} else if guess > rnd {
		fmt.Println("Oops. Your guess was HIGH")
	}
}
