package main

import "fmt"

func main() {
	var a float32 = 358.99999
	var b float32 = 359.00001

	fmt.Println("a - b =", a-b)

	c := 358.99999
	d := 359.00001

	fmt.Println("c - d =", c-d)

	var e float32 = 359.9
	fmt.Println(e)

	var f float64 = float64(e)
	fmt.Println(f)
}
