package direction

import (
	"slices"

	"github.com/nlduy0310/aoc-2024/utils"
)

var rightTurnMap map[Direction]Direction = map[Direction]Direction{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

var leftTurnMap map[Direction]Direction = map[Direction]Direction{
	Up:    Left,
	Left:  Down,
	Down:  Right,
	Right: Left,
}

var validTurnDirections []Direction = []Direction{Left, Right}

func (d Direction) AfterTurn(turnDirection Direction) Direction {

	ensureValid(d)
	ensureValid(turnDirection)
	utils.Assert(
		slices.Contains(validTurnDirections, turnDirection),
		sliceString(validDirections, ", "),
	)

	switch turnDirection {
	case Left:
		return leftTurnMap[d]
	case Right:
		return rightTurnMap[d]
	// never happens, just to satisfy the compiler
	default:
		return -1
	}
}
