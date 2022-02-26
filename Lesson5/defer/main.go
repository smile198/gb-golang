package main

import "fmt"

func main() {
	i := 0
	for i = 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()

		panic(2)
	}

	panic(1)

	i = 22
}
