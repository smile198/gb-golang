package calc

import (
	"errors"
	"math"
)

var (
	DivisionByZeroError    = errors.New("деление на 0 невозможно")
	NegativeFactorialError = errors.New("расчитать факториал можно только с неотрицательными целыми числами")
	UnknownOperatorError   = errors.New("неизвестный оператор")
)

func Calc(a, b float32, op string) (float32, error) {
	var res float32

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			return 0, DivisionByZeroError
		}
		res = a / b
	case "^":
		res = float32(math.Pow(float64(a), float64(b)))
	case "!":
		if (a < 0) || (a != float32(int32(a))) {
			return 0, NegativeFactorialError
		}
		res = float32(calcFactorial(int(a)))
	default:
		return 0, UnknownOperatorError
	}

	return res, nil
}

func calcFactorial(number int) int {
	if number < 2 {
		return 1
	}

	return number * calcFactorial(number-1)
}
