package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello" + " " + "world!"
	fmt.Println(s)

	// Длина строки
	str := "ABCDЁ"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
