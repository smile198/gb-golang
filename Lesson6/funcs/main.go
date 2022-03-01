package main

import "fmt"

type Counter struct {
	val int
}

func (c Counter) String() string {
	return fmt.Sprintf("%v", c.val)
}

func (c *Counter) increment() {
	c.val++
}

func (c *Counter) decrement() {
	c.val--
}

func main() {
	x := Counter{}
	x.increment()
	fmt.Println(x)
}
