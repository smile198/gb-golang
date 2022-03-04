package main

import (
	"fmt"
	"math"
)

type Squarer interface {
	GetSquare() float64
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (t triangle) GetSquare() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

type circle struct {
	r float64
}

func (c circle) GetSquare() float64 {
	return math.Pi * c.r * c.r
}

func printSquare(s Squarer) {
	var tpe string
	switch s.(type) {
	case circle:
		tpe = "кргу"
	case triangle:
		tpe = "треугольник"
	default:
		tpe = "не известно что"
	}
	fmt.Printf("Это %s, площадь равена: %.2f\n", tpe, s.GetSquare())
}

func main() {
	t := triangle{
		a: 4,
		b: 5,
		c: 6,
	}

	printSquare(t)

	c := circle{r: 5}

	printSquare(c)

	b := Squarer(t)

	castedT, okT := b.(triangle)

	fmt.Println(okT, castedT.b)

	castedC, okC := b.(circle)

	fmt.Println(okC, castedC.r)
}
