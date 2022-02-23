package main

import "fmt"

var primeNumbers []int

func main() {
	fmt.Println("Поиск простых чисел в диапазоне от 0 до N.")

	var max int
	fmt.Print("Введи N: ")
	_, err := fmt.Scanln(&max)

	if err != nil {
		panic(err)
	}

	if max < 0 {
		panic("Введите число больше 0.")
	}

	for number := 1; number <= max; number++ {
		if number != 1 {
			checkNumber(number)
		}
	}

	fmt.Println("Натуральные числа в диапазоне от 0 до ", max)
	fmt.Println(primeNumbers)
}

func checkNumber(number int) {
	for _, primeNumber := range primeNumbers {
		if number%primeNumber == 0 {
			return
		}
	}

	primeNumbers = append(primeNumbers, number)
}
