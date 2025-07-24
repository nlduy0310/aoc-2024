package guard

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day06/direction"
	"github.com/nlduy0310/aoc-2024/day06/position"
)

type Guard struct {
	Position  position.Position
	Direction direction.Direction
}

func NewGuard(position position.Position, direction direction.Direction) Guard {

	return Guard{
		Position:  position,
		Direction: direction,
	}
}

func (g Guard) String() string {

	return fmt.Sprintf("Guard[position=%s, direction=%s]", g.Position, g.Direction)
}

func (g *Guard) MoveForward() {

	g.Position.MoveInDirection(g.Direction)
}

func (g *Guard) TurnRight() {

	g.Direction = g.Direction.AfterTurn(direction.Right)
}
