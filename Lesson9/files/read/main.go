package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("my-file.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	stat, _ := file.Stat()
	fmt.Printf("%+v", stat)

	fromFile := make([]byte, stat.Size())

	read, err := file.Read(fromFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("I have read %d bytes: %s, len: %d\n", read, string(fromFile), len(fromFile))

	content, err := os.ReadFile("my-file.txt")
	fmt.Println(string(content))
}
