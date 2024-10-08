// Перед вами неполный код программы, которая считает, сколько строк пользователь ввёл в консоль,
// и после ввода каждой новой строки выводит общее количество на экран.
// Напишите реализацию функции f.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func funct(count *int) {
	// *count = *count + 1
	(*count)++
}

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		funct(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}
