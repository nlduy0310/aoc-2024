package direction

import (
	"fmt"
	"slices"

	utils "github.com/nlduy0310/aoc-2024/utils"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var validDirections []Direction = []Direction{
	Up,
	Down,
	Left,
	Right,
}

func (d Direction) isValid() bool {

	return slices.Contains(validDirections, d)
}

func ensureValid(d Direction) {

	utils.Assert(d.isValid(), fmt.Sprintf("invalid direction enum: %d", d))
}

func (d Direction) String() string {

	ensureValid(d)

	switch d {
	case Up:
		return "Direction[Up]"
	case Down:
		return "Direction[Down]"
	case Left:
		return "Direction[Left]"
	case Right:
		return "Direction[Right]"
	// never happens, just to satisfy the compiler
	default:
		return ""
	}
}
