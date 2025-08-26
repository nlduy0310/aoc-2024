package robot

import (
	"fmt"

	"github.com/nlduy0310/aoc-2024/day14/position"
)

type Robot struct {
	InitialPosition position.Position
	CurrentPosition position.Position
	velocityX       int
	velocityY       int
}

func NewRobot(initialPosition position.Position, velocityX, velocityY int) Robot {

	return Robot{
		InitialPosition: initialPosition,
		CurrentPosition: initialPosition,
		velocityX:       velocityX,
		velocityY:       velocityY,
	}
}

func (r Robot) String() string {

	return fmt.Sprintf("Robot[initialPosition=%s, currentPosition=%s, velocityX=%d, velocityY=%d]", r.InitialPosition.String(), r.CurrentPosition.String(), r.velocityX, r.velocityY)
}
