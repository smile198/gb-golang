package main

import "os"

func main() {
	//if err := os.WriteFile("my-file.txt", []byte("Hello world!"), 0777); err != nil {
	//	panic(err)
	//}

	file, err := os.OpenFile("my-file.txt", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteAt([]byte("Another Hello"), 12)
	if err != nil {
		panic(err)
	}
}
