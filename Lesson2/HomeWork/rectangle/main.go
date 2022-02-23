package main

import "fmt"

func main() {
	fmt.Println("Расчет площади прямоугольника.")

	var a int
	var b int

	fmt.Print(`Введите длину стороны "a": `)
	_, err := fmt.Scanln(&a)

	if err != nil {
		panic(err)
	}

	fmt.Print(`Введите длину стороны "b": `)
	_, err = fmt.Scanln(&b)

	if err != nil {
		panic(err)
	}

	fmt.Println("Площадь прямоугольника:", a*b)
}
