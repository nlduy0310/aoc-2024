package position

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day16/direction"
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

func (p *Position) Move(x, y int) {
	p.X += x
	p.Y += y
}

func (p Position) Moved(x, y int) Position {
	return NewPosition(p.X+x, p.Y+y)
}

func (p *Position) MoveInDirection(d direction.Direction, offset int) {
	switch d {
	case direction.North:
		p.Y -= offset
	case direction.South:
		p.Y += offset
	case direction.West:
		p.X -= offset
	case direction.East:
		p.X += offset
	default:
		panic(fmt.Sprintf("unrecognized direction enum value: %d", d))
	}
}

func (p Position) MovedInDirection(d direction.Direction, offset int) Position {
	ret := p
	ret.MoveInDirection(d, offset)

	return ret
}

func (p Position) String() string {
	return fmt.Sprintf("Position[X=%d, Y=%d]", p.X, p.Y)
}
