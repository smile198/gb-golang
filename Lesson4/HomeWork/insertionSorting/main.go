package main

import "fmt"

func main() {
	inputArray := []int{5, 2, 4, 6, 1, 3, 8, 7, 9, 10}
	fmt.Println("Входной массив:", inputArray)

	sortByInsertion(inputArray)
	fmt.Println("Отсортированный массив:", inputArray)
	// Отсортированный массив: [1 2 3 4 5 6 7 8 9 10]

	inputArray2 := []int{1, 2, 1, 2, 4, 3, 3, 3}
	fmt.Println("Входной массив:", inputArray2)

	sortByInsertion(inputArray2)
	fmt.Println("Отсортированный массив:", inputArray2)
	// Отсортированный массив: [1 1 2 2 3 3 3 4]
}

func sortByInsertion(array []int) {
	for currentIndex, currentValue := range array {
		prevIndex := currentIndex - 1

		for prevIndex >= 0 && array[prevIndex] > currentValue {
			array[prevIndex+1] = array[prevIndex]
			prevIndex--
		}

		array[prevIndex+1] = currentValue
	}
}
