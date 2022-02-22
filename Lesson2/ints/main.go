package main

import "fmt"

func main() {
	// Классическое объявление переменной
	var a int
	a = 5

	// Сократить
	var b int = 2

	// Сократить
	var c = 6

	// Сократить
	d := 7

	fmt.Println("a + b + c + d =", a+b+c+d)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
}
