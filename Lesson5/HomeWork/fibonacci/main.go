package main

import "fmt"

type FibonacciFinder struct {
	cached map[uint]uint
}

func (ff *FibonacciFinder) getNumber(index uint) uint {
	//if ff.cached == nil {
	//	ff.cached = make(map[int]int)
	//}

	_, exists := ff.cached[index]
	if !exists {
		ff.cached[index] = ff.calcNumber(index)
	}

	return ff.cached[index]
}

func (ff FibonacciFinder) calcNumber(number uint) uint {
	if number <= 1 {
		return number
	}

	return ff.getNumber(number-1) + ff.getNumber(number-2)
}

func main() {
	fmt.Print("Сколько чисел из последовательности Фибоначчи вывести? ")

	var number uint
	_, err := fmt.Scanln(&number)
	if err != nil {
		panic(err)
	}

	defer fmt.Print("\n")
	findFibonacciSequence(number, func(fibonacciNumber uint) {
		fmt.Print(fibonacciNumber, " ")
	})
}

func findFibonacciSequence(length uint, callback func(uint)) {
	fibonacci := FibonacciFinder{
		cached: make(map[uint]uint, length),
	}

	fmt.Println(fibonacci)

	for index := uint(1); index <= length; index++ {
		callback(fibonacci.getNumber(index))
	}
}
