package box

import (
	"fmt"
	"log"

	"github.com/nlduy0310/aoc-2024/day15/direction"
	"github.com/nlduy0310/aoc-2024/day15/position"
)

type Box struct {
	LeftPosition  position.Position
	RightPosition position.Position
}

func (b Box) String() string {

	return fmt.Sprintf("Box[left=%s, right=%s]", b.LeftPosition.String(), b.RightPosition.String())
}

func NewBox(leftPos, rightPos position.Position) Box {

	return Box{
		LeftPosition:  leftPos,
		RightPosition: rightPos,
	}
}

func (b Box) ContainsPosition(p position.Position) bool {

	return b.LeftPosition == p || b.RightPosition == p
}

func (b Box) GetBound(d direction.Direction) []position.Position {

	switch d {
	case direction.Up, direction.Down:
		return []position.Position{b.LeftPosition.Moved(d, 1), b.RightPosition.Moved(d, 1)}
	case direction.Left:
		return []position.Position{b.LeftPosition.Moved(d, 1)}
	case direction.Right:
		return []position.Position{b.RightPosition.Moved(d, 1)}
	default:
		log.Fatalf("invalid direction enum value: %d", d)
		panic("satisfy compiler")
	}
}

func (b *Box) Move(d direction.Direction, steps int) {

	b.LeftPosition.Move(d, steps)
	b.RightPosition.Move(d, steps)
}
