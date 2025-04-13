// pass_fail сообщает, сдал ли пользователь экзамен.
package utils

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func PassFail() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) // Создать «буферизованное средство чтения» для получения текста с клавиатуры
	input, err := reader.ReadString('\n') //Возвращает текст, введенный пользователем до нажатия клавиши Enter
	if err != nil {
		log.Fatal(err)	
	}
	
	fmt.Println(input)
}

