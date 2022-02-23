package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getFilledSlice(size int) []int {
	rand.Seed(time.Now().Unix())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}

	return slice
}

func main() {
	startArr := getFilledSlice(10)
	fmt.Println(startArr)

	sortedArr := bubbleSort(startArr)
	fmt.Println(sortedArr)
}

func bubbleSort(arr []int) []int {
	for {
		swapped := false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}
