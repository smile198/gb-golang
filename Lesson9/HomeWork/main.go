package main

import (
	"GeekBrains/Lesson9/HomeWork/config"
	"flag"
	"fmt"
)

func main() {
	// go run main.go --file=config.json
	filePath := flag.String("file", "", "Ссылка на json файл")

	flag.Parse()

	configFile, err := config.FromJsonFile(*filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(configFile)
}
