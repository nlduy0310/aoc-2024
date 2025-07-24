package position

import (
	"fmt"

	direction "github.com/nlduy0310/aoc-2024/day06/direction"
)

type Position struct {
	Row, Col int // allowing negative ints to represent invalid positions
}

func NewPosition(row, col int) Position {

	return Position{
		Row: row,
		Col: col,
	}
}

func (p Position) Copy() Position {

	return Position{
		Row: p.Row,
		Col: p.Col,
	}
}

func (p Position) String() string {

	return fmt.Sprintf("Position[row=%d, col=%d]", p.Row, p.Col)
}

func (p *Position) MoveInDirection(d direction.Direction) {

	switch d {
	case direction.Up:
		p.Row -= 1
	case direction.Down:
		p.Row += 1
	case direction.Left:
		p.Col -= 1
	case direction.Right:
		p.Col += 1
	default:
		panic(fmt.Sprintf("invalid direction: %d", d))
	}
}
