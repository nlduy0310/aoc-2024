package point

import "fmt"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {

	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) String() string {

	return fmt.Sprintf("Point[X=%d, Y=%d]", p.X, p.Y)
}
