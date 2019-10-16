package main

import (
	"fmt"

	"./calculator"
)

//Функция с ДЗ
func help() {
	fmt.Println("Вы должны слитно ввести выражение, которое хотите посчитать.")
	fmt.Println("Вы можете писать сложные выражения используя скобки.")
	fmt.Println("Также вам доступны операторы sin, cos, tg, ctg и т.д.")
}

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			help()
		} else if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
