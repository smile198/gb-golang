package main

import "fmt"

func main() {
	var y *int = new(int)
	fmt.Println(*y)

	//var x = new(map[int]string)
	//(*x)[1] = "one"
	//fmt.Println(*x)

	var z = &map[int]string{}
	(*z)[1] = "one"
	fmt.Println(z)
}
