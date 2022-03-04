package main

import (
	"fmt"
	"sort"
)

type myType struct {
	a int
	b string
}

func main() {
	sliceOfInts := []int{1, 2, 5, 78, 4, 6, 7, 7}
	sliceOfStrings := []string{"ferf", "wegfer", "wefgwfg"}
	sliceOfTypes := []myType{
		{
			a: 1,
			b: "2",
		},
		{
			a: 2,
			b: "1",
		},
	}

	sort.Ints(sliceOfInts)
	sort.Strings(sliceOfStrings)

	sort.Slice(sliceOfTypes, func(i, j int) bool {
		return sliceOfTypes[i].b < sliceOfTypes[j].b
	})

	fmt.Println(sliceOfInts)
	fmt.Println(sliceOfStrings)
	fmt.Println(sliceOfTypes)
}
