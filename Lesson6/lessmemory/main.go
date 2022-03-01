package main

import (
	"fmt"
	"unsafe"
)

func myFunc(arr [10000]string) {
	fmt.Println(unsafe.Pointer(&arr[0]))
	fmt.Println(unsafe.Sizeof(arr))
}

func myFuncPointer(arr *[10000]string) {
	fmt.Println(unsafe.Pointer(&arr[0]))
	fmt.Println(unsafe.Sizeof(arr))
}

func main() {
	arr := [10000]string{}
	fmt.Println(unsafe.Pointer(&arr[0]))
	fmt.Println(unsafe.Sizeof(arr))
	myFunc(arr)
	myFuncPointer(&arr)
}
