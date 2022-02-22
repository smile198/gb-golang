package main

import "fmt"

func main() {
	a := 3
	b := 5

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
}
