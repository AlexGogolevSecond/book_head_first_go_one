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
	success := false
	fmt.Println("I've chosen a random number between 1 and 100.")
	// fmt.Println(rnd)
	reader := bufio.NewReader(os.Stdin) // Создать «буферизованное средство чтения» для получения текста с клавиатуры
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

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
		
		// for x:=1; x < 5; x++ {
		// 	y := x + 10
		// 	fmt.Println("y" + strconv.Itoa(y))  // !!! пробразуем int в строку
		
		// fmt.Print(x) //x недоступен в текущей области видимости !!!
		// for x:=10; x > 0; x-=2 {
		// 	fmt.Println(x)
		// }
		
		if guess < rnd {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > rnd {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("You are win!!! You spent " + strconv.Itoa(guesses + 1) + " attempt. Target is " + strconv.Itoa(rnd))
			success = true
			break
		}		
	}
	//if success == false {
	if !success {
		fmt.Println("You lost. Target number is " + strconv.Itoa(rnd))
	}
}

