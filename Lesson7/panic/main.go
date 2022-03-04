package main

import "fmt"

func a() {
	b()
}

func b() {
	c()
}

func c() {
	panic(123)
}

func main() {
	defer func() {
		err := recover()
		fmt.Println("Поймали панику", err)
	}()
	a()

	fmt.Println("not from defer")
}
