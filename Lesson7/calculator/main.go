package main

import (
	"errors"
	"fmt"
)

func main() {
	var a float64
	var b float64
	var op string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&op)

	result, err := calc(a, b, op)
	if err != nil {
		if errors.Is(err, ErrorDivideByZero) {
			fmt.Println("Вы пытаетесь делить на ноль.")
			return
		}
		fmt.Printf("Вы допустили ошибку: %s\n", err.Error())
		return
	}

	fmt.Println(result)
}

type ErrorDivideByZero struct{}

func (e ErrorDivideByZero) Error() string {
	panic("can't divide by 0")
}

func calc(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "/":
		if b == 0 {
			return 0, ErrorDivideByZero{}
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
