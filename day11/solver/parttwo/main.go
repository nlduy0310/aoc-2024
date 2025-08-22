package parttwo

import (
	"github.com/nlduy0310/aoc-2024/day11/lookup"
	"github.com/nlduy0310/aoc-2024/day11/parser"
	"github.com/nlduy0310/aoc-2024/day11/stone"
)

const maximumDepth int = 75

type Solver struct {
	inputFile string
	stones    []stone.Stone
	lookup    lookup.Lookup
}

func MustInitSolver(inputFile string) Solver {

	stones := parser.MustParseFromFile(inputFile)
	return Solver{
		inputFile: inputFile,
		stones:    stones,
		lookup:    lookup.NewLookup(),
	}
}

func (s Solver) Solve() int {

	return s.solveRecursively(s.stones, 0)
}

func (s Solver) solveRecursively(stones []stone.Stone, depth int) int {

	if depth == maximumDepth {
		for _, stone := range stones {
			s.lookup.Mark(stone.Val, maximumDepth-depth, 1)
		}
		return 1 * len(stones)
	}

	ret := 0

	for _, stone := range stones {
		cachedResult, ok := s.lookup.TryGet(stone.Val, maximumDepth-depth)
		if ok {
			ret += cachedResult
		} else {
			childStones := stone.Change()
			result := s.solveRecursively(childStones, depth+1)
			s.lookup.Mark(stone.Val, maximumDepth-depth, result)
			ret += result
		}
	}

	return ret
}
