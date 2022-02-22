package main

import "fmt"

func main() {
	var a int = 7
	var b bool = a > 7
	fmt.Println(b)

	c := 6

	if !(a > c) {
		fmt.Println("Больше")
	} else {
		fmt.Println("Меньше")
	}
}
