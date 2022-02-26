package main

import "fmt"

type Operation func(int, int) float64

var operations = map[string]Operation{
	"+": func(a, b int) float64 {
		return float64(a + b)
	},
	"-": func(a, b int) float64 {
		return float64(a - b)
	},
	"/": func(a, b int) float64 {
		return float64(a) / float64(b)
	},
}

func OperationWithLog(operationFunc Operation, a, b int, operation string) {
	fmt.Printf("i'm gonna count %d %s %d\n", a, operation, b)
	res := operationFunc(a, b)
	fmt.Println("result is:", res)
}

func main() {
	a, b := 3, 4
	operation := "+"

	myOperation, exists := operations[operation]
	if !exists {
		panic(1)
	}

	OperationWithLog(myOperation, a, b, operation)
}
