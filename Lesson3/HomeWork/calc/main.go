package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)

	if op != "!" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на 0 невозможно.")
			return
		}
		res = a / b
	case "^":
		res = float32(math.Pow(float64(a), float64(b)))
	case "!":
		if (a < 0) || (a != float32(int32(a))) {
			fmt.Println("Расчитать факториал можно только с неотрицательными целыми числами")
			return
		}
		res = float32(calcFactorial(int(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		return
	}

	if res == float32(int32(res)) {
		fmt.Printf("Результат выполнения операции: %d\n", int32(res))
	} else {
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	}
}

func calcFactorial(number int) int {
	if number < 2 {
		return 1
	}
	return number * calcFactorial(number-1)
}
