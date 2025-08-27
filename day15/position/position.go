package position

import (
	"fmt"
	"log"

	"github.com/nlduy0310/aoc-2024/day15/direction"
)

type Position struct {
	Row, Col int
}

func NewPosition(row, col int) Position {

	return Position{Row: row, Col: col}
}

func (p Position) String() string {

	return fmt.Sprintf("Position[row=%d, col=%d]", p.Row, p.Col)
}

func (p *Position) MoveUp(steps int) {

	p.Row -= steps
}

func (p *Position) MoveDown(steps int) {

	p.Row += steps
}

func (p *Position) MoveLeft(steps int) {

	p.Col -= steps
}

func (p *Position) MoveRight(steps int) {

	p.Col += steps
}

func (p Position) GetUp(steps int) Position {

	return NewPosition(p.Row-steps, p.Col)
}

func (p Position) GetDown(steps int) Position {

	return NewPosition(p.Row+steps, p.Col)
}

func (p Position) GetLeft(steps int) Position {

	return NewPosition(p.Row, p.Col-steps)
}

func (p Position) GetRight(steps int) Position {

	return NewPosition(p.Row, p.Col+steps)
}

func (p Position) Moved(d direction.Direction, steps int) Position {

	switch d {
	case direction.Up:
		return p.GetUp(steps)
	case direction.Down:
		return p.GetDown(steps)
	case direction.Left:
		return p.GetLeft(steps)
	case direction.Right:
		return p.GetRight(steps)
	default:
		log.Fatalf("invalid direction enum value: %d", d)
		panic("just to satisfy compiler")
	}
}

func (p *Position) Move(d direction.Direction, steps int) {

	switch d {
	case direction.Up:
		p.MoveUp(steps)
	case direction.Down:
		p.MoveDown(steps)
	case direction.Left:
		p.MoveLeft(steps)
	case direction.Right:
		p.MoveRight(steps)
	default:
		log.Fatalf("invalid direction enum value: %d", d)
		panic("just to satisfy compiler")
	}
}
