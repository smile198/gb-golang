package triangle

import (
	"errors"
	"fmt"
	"math"
)

var ErrImpossibleTriangle = errors.New("Impossible triangle")

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Square() (float64, error) {
	if t.SideA+t.SideB <= t.SideC || t.SideA+t.SideC <= t.SideB || t.SideC+t.SideB <= t.SideA {
		return 0, ErrImpossibleTriangle
	}

	p := float64(t.SideA+t.SideB+t.SideC) / 2
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC)), nil
}

func (t Triangle) String() string {
	square, err := t.Square()
	if err != nil {
		fmt.Sprintf("Невозможный треугольник со сторонами %.2f, %.2f и %.2f", t.SideA, t.SideB, t.SideC)
	}

	return fmt.Sprintf("Треугольник со сторонами %.2f, %.2f и %.2f и площадью %.2f", t.SideA, t.SideB, t.SideC, square)
}
