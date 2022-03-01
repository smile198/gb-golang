package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println(x)

	var pointerOnX = &x
	fmt.Printf("%T %T\n", x, pointerOnX)
	fmt.Printf("type: %T, value: %+v\n", pointerOnX, pointerOnX)
	fmt.Printf("type: %T, value: %+v\n", pointerOnX, *pointerOnX)
}
