package position

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day12/direction"
)

type Position struct {
	Row, Col int
}

func (p Position) String() string {

	return fmt.Sprintf("Position[row=%d, col=%d]", p.Row, p.Col)
}

func NewPosition(row, col int) Position {

	return Position{Row: row, Col: col}
}

func (p Position) Up() Position {

	return NewPosition(p.Row-1, p.Col)
}

func (p Position) Down() Position {

	return NewPosition(p.Row+1, p.Col)
}

func (p Position) Left() Position {

	return NewPosition(p.Row, p.Col-1)
}

func (p Position) Right() Position {

	return NewPosition(p.Row, p.Col+1)
}

func (p Position) Move(d direction.Direction) Position {

	switch d {
	case direction.North:
		return p.Up()
	case direction.South:
		return p.Down()
	case direction.West:
		return p.Left()
	case direction.East:
		return p.Right()
	default:
		panic(fmt.Sprintf("invalid direction %v", d))
	}
}
