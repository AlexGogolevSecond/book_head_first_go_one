// pass_fail сообщает, сдал ли пользователь экзамен.
package utils

import (
	"fmt"
	"bufio"
	"os"
)

func PassFail() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) // Создать «буферизованное средство чтения» для получения текста с клавиатуры
	input, _ := reader.ReadString('\n') //Возвращает текст, введенный пользователем до нажатия клавиши Enter
	fmt.Println(input)
}

