package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите трехзначное число: ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		panic(err)
	}

	if (number > 999) || (number < 100) {
		panic("Укажите трехзначное число (от 100 до 999)")
	}

	ones := number % 10
	tens := (number / 10) % 10
	hundreds := number / 100

	fmt.Println(number, "=", hundreds, "сотни,", tens, "десятка,", ones, "единицы")
}
