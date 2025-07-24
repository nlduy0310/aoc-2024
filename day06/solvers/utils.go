package solvers

import (
	"github.com/nlduy0310/aoc-2024/day06/direction"
	"github.com/nlduy0310/aoc-2024/day06/guard"
	"github.com/nlduy0310/aoc-2024/day06/position"
)

// too lazy to make this efficient
func mustFindGuard(lines []string) guard.Guard {

	for rowIdx, line := range lines {
		for colIdx, cellRune := range line {
			cellValue := string(cellRune)
			if cellValue == "^" {
				return guard.NewGuard(
					position.NewPosition(rowIdx, colIdx),
					direction.Up,
				)
			}
		}
	}

	panic("can not find a guard symbol '^' in the map")
}
