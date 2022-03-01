package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 5
	fmt.Println(unsafe.Sizeof(x))
}
