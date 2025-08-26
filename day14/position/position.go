package position

import (
	"fmt"
	"log"
)

type Position struct {
	X, Y int
}

func NewPosition(x, y int) Position {

	return Position{
		X: x,
		Y: y,
	}
}

func (p Position) String() string {

	return fmt.Sprintf("Position[x=%d, y=%d]", p.X, p.Y)
}

func (p *Position) ClampX(xMin, xMax int) {

	clampedX, err := clampValue(p.X, xMin, xMax)
	if err != nil {
		log.Fatal(err)
	}

	p.X = clampedX
}

func (p *Position) ClampY(yMin, yMax int) {

	clampedY, err := clampValue(p.Y, yMin, yMax)
	if err != nil {
		log.Fatal(err)
	}

	p.Y = clampedY
}

// Moves the position in-place
func (p *Position) Move(xVelocity, yVelocity int) {

	p.X += xVelocity
	p.Y += yVelocity
}

func (p Position) GetUp(steps int) Position {

	return NewPosition(p.X, p.Y-steps)
}

func (p Position) GetDown(steps int) Position {

	return NewPosition(p.X, p.Y+steps)
}

func (p Position) GetLeft(steps int) Position {

	return NewPosition(p.X-steps, p.Y)
}

func (p Position) GetRight(steps int) Position {

	return NewPosition(p.X+steps, p.Y)
}
