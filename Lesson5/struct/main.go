package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (point Point) DistanceTo(anotherPoint Point) float64 {
	return math.Sqrt(math.Pow(point.x-anotherPoint.x, 2) + math.Pow(point.y-anotherPoint.y, 2))
}

type Circle struct {
	certer Point
	radius float64
}

func (circle Circle) DistanceTo(anoutherCircle Circle) float64 {
	return math.Abs(circle.certer.DistanceTo(anoutherCircle.certer) - circle.radius - anoutherCircle.radius)
}

func CalcDistBetweenPoints(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow(point2.x-point1.x, 2) + math.Pow(point2.y-point1.y, 2))
}

func CalcDistBetweenPointsSpead(points ...Point) []float64 {
	if len(points) == 1 {
		return []float64{0}
	}

	result := make([]float64, 0, len(points))

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if j != i {
				result = append(result, math.Sqrt(math.Pow(points[i].x-points[j].x, 2)+math.Pow(points[i].y-points[j].y, 2)))
			}
		}
	}

	return result
}

func main() {
	points := []Point{{x: 3, y: 3}, {x: 6, y: 6}, {x: 9, y: 9}}

	dist1 := CalcDistBetweenPoints(points[0], points[1])
	dist2 := CalcDistBetweenPoints(points[0], points[2])
	dist3 := CalcDistBetweenPoints(points[2], points[1])
	distSpread := CalcDistBetweenPointsSpead(points...)
	fmt.Println(dist1, dist2, dist3, distSpread, points[0].DistanceTo(points[1]))

	circle1 := Circle{
		certer: Point{
			x: 1,
			y: 2,
		},
		radius: 10,
	}
	circle2 := Circle{
		certer: Point{
			x: 5,
			y: 7,
		},
		radius: 6,
	}

	fmt.Println(circle1.DistanceTo(circle2))
}
