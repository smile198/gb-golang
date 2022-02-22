package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Расчет диаметра и длины окружности по заданной площади круга.")

	var area int
	fmt.Print("Введите площадь круга: ")
	_, err := fmt.Scanln(&area)

	if err != nil {
		panic(err)
	}

	radius := math.Sqrt(float64(area) / math.Pi)
	diameter := 2 * radius
	perimeter := 2 * math.Pi * radius

	fmt.Println("Диаметр круга: ", diameter)
	fmt.Println("Длина окружности: ", perimeter)
}
