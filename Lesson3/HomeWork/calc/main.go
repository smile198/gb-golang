package main

import (
	"fmt"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		return
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
