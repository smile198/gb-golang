package sorting

import "sort"

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

func Sort(arr []int) {
	//sortByInsertion(arr)
	sort.Ints(arr)
}
