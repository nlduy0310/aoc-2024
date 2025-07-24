package labmap

import (
	"fmt"
	"slices"

	"github.com/nlduy0310/aoc-2024/utils"
)

type cellState int

const (
	cellEmpty cellState = iota
	cellBlocked
)

var validCellStates []cellState = []cellState{cellEmpty, cellBlocked}

func (s cellState) String() string {

	utils.Assert(
		slices.Contains(validCellStates, s),
		fmt.Sprintf("invalid cell state enum: %d", s),
	)

	switch s {
	case cellEmpty:
		return "."
	case cellBlocked:
		return "#"
	// never happens, just to satisfy the compiler
	default:
		return ""
	}
}
