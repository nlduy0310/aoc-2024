package state

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day16/direction"
	"github.com/nlduy0310/aoc-2024/day16/position"
)

type State struct {
	Position  position.Position
	Direction direction.Direction
}

func NewState(pos position.Position, dir direction.Direction) State {
	return State{
		Position:  pos,
		Direction: dir,
	}
}

func (s State) MovedForward() State {
	return NewState(s.Position.MovedInDirection(s.Direction, 1), s.Direction)
}

func (s State) TurnedLeft() State {
	return NewState(s.Position, s.Direction.TurnedLeft())
}

func (s State) TurnedRight() State {
	return NewState(s.Position, s.Direction.TurnedRight())
}

func (s State) String() string {
	return fmt.Sprintf("State[Position=%s, Direction=%s]", s.Position, s.Direction.Name())
}
