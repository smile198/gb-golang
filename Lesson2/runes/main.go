package main

import "fmt"

func main() {
	//var a rune = 'a'
	var original string = "MY STRING"
	var encrypted string = ""

	for i := 0; i < len(original); i++ {
		encrypted += string(original[i] + 2)
	}

	fmt.Println("orginal:", original)
	fmt.Println("encrypted:", encrypted)
}
