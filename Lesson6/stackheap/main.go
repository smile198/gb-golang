package main

func myFunc() *int {
	var a = 5
	println(a)

	return &a
}

func main() {
	var x int = 5
	println(x)

	myFunc()
	// go run -gcflags="-m -l" main.go
}
