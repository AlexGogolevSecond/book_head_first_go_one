// pass_fail сообщает, сдал ли пользователь экзамен.
package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PassFail() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)   // Создать «буферизованное средство чтения» для получения текста с клавиатуры
	input, err := reader.ReadString('\n') //Возвращает текст, введенный пользователем до нажатия клавиши Enter

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)  // удаляем с краёв пробелы и всякие служебные символы типа табуляции и переноса строки
	grade, err := strconv.ParseFloat(input, 64)  // преобразуем строку в число float

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println(status)

	// s := "\t lalala tru-la-la\n"
	// fmt.Println(strings.TrimSpace(s))
	// fmt.Println(s)

}
