package main

import (
	"fmt"
	"os"
)

func main() {
	myEnv := os.Getenv("MY_ENV")
	fmt.Println("my env is:", myEnv)

	myEnv, ok := os.LookupEnv("MY_ENV")
	if !ok {
		fmt.Println("Переменная не задана")
	} else {
		fmt.Println("my env is:", myEnv)
	}
}
