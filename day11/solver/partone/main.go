package partone

import (
	"github.com/nlduy0310/aoc-2024/day11/parser"
	"github.com/nlduy0310/aoc-2024/day11/stone"
)

type Solver struct {
	inputFile string
	stones    []stone.Stone
}

func MustInitSolver(inputFile string) Solver {

	stones := parser.MustParseFromFile(inputFile)

	return Solver{
		inputFile: inputFile,
		stones:    stones,
	}
}

func (s Solver) Solve() int {

	return s.solveStonesRecursively(s.stones, 0)
}

func (s Solver) solveStonesRecursively(stones []stone.Stone, depth int) int {

	if depth == 25 {
		return 1 * len(stones)
	}

	ret := 0

	for _, stone := range stones {
		childStones := stone.Change()
		ret += s.solveStonesRecursively(childStones, depth+1)
	}

	return ret
}
