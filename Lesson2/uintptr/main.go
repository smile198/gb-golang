package main

import "fmt"

func main() {
	a := 5
	var b uintptr = uintptr(a)

	fmt.Println(&a)
	fmt.Println(&b)
}
