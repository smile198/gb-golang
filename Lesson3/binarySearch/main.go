package main

import "fmt"

func main() {
	arr := []int{-1, 0, 1, 5, 10}

	index := binarySearch(arr, 5)
	fmt.Printf("Число %d на позиции %d\n", 5, index)

	index = binarySearch(arr, 2)
	fmt.Printf("Число %d на позиции %d\n", 2, index)
}

func binarySearch(arr []int, needle int) int {
	left := 0
	right := len(arr)

	for {
		middle := countMiddle(left, right)
		fmt.Println(left, right, middle)

		if needle < arr[middle] {
			right = middle - 1
		} else if needle > arr[middle] {
			left = middle + 1
		} else {
			return middle
		}

		if left > right {
			return -1
		}
	}
}

func countMiddle(left int, right int) int {
	return (left + right) / 2
}
